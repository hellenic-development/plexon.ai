// Command sync_models_dev pulls https://models.dev/api.json and merges it
// into the hand-authored providers/*.yaml files. For each provider not in
// .sync-exclude it applies a latest-per-family filter so we never accumulate
// long tails of deprecated model variants (e.g. kimi-k2-0712 beside
// kimi-k2-0905-preview — only the newer survives).
//
// Usage:
//
//	go run ./scripts/sync_models_dev                   # normal run
//	go run ./scripts/sync_models_dev -dry-run          # print diff, no writes
//	go run ./scripts/sync_models_dev -fixture path.json # use a local file
//
// The merge rule preserves hand-edited fields that models.dev doesn't know
// about: per-model `base_url` / `headers`, provider-level `builtin_tools` /
// `regions` / `default_model`, and anything under `sync:`.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	catalogtypes "github.com/hellenic-development/plexon.ai/catalog/scripts"
)

const modelsDevAPI = "https://models.dev/api.json"

func main() {
	providersDir := flag.String("providers", "providers", "directory containing hand-authored provider YAMLs")
	exclude := flag.String("exclude", ".sync-exclude", "file listing providers to never overwrite")
	fixture := flag.String("fixture", "", "if set, read from this local JSON file instead of models.dev")
	dryRun := flag.Bool("dry-run", false, "show changes without writing")
	flag.Parse()

	upstream, err := fetchUpstream(*fixture)
	if err != nil {
		log.Fatalf("fetch: %v", err)
	}
	excluded, err := readExclude(*exclude)
	if err != nil {
		log.Fatalf("read %s: %v", *exclude, err)
	}

	// For each provider already present in providers/ (except excluded ones),
	// merge models.dev data in. We do not create new provider files
	// automatically — adding a new provider is an explicit human act.
	entries, err := os.ReadDir(*providersDir)
	if err != nil {
		log.Fatalf("read %s: %v", *providersDir, err)
	}

	totalChanges := 0
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".yaml" {
			continue
		}
		name := strings.TrimSuffix(e.Name(), ".yaml")
		if slices.Contains(excluded, name) {
			fmt.Printf("skip  %-16s (in .sync-exclude)\n", name)
			continue
		}
		upstreamProv, ok := upstream[name]
		if !ok {
			fmt.Printf("keep  %-16s (not in models.dev — leaving untouched)\n", name)
			continue
		}
		path := filepath.Join(*providersDir, e.Name())
		changed, err := mergeOne(path, upstreamProv, *dryRun)
		if err != nil {
			log.Fatalf("merge %s: %v", path, err)
		}
		if changed {
			totalChanges++
		}
	}

	if *dryRun {
		fmt.Printf("\n[dry-run] %d provider(s) would change\n", totalChanges)
	} else {
		fmt.Printf("\nsynced %d provider(s)\n", totalChanges)
	}
}

// fetchUpstream returns the models.dev payload as a map keyed by provider id.
func fetchUpstream(fixture string) (map[string]upstreamProvider, error) {
	var body []byte
	if fixture != "" {
		b, err := os.ReadFile(fixture)
		if err != nil {
			return nil, err
		}
		body = b
	} else {
		client := &http.Client{Timeout: 30 * time.Second}
		resp, err := client.Get(modelsDevAPI)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("models.dev returned %s", resp.Status)
		}
		b, err := io.ReadAll(io.LimitReader(resp.Body, 16<<20))
		if err != nil {
			return nil, err
		}
		body = b
	}

	var raw map[string]upstreamProvider
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, fmt.Errorf("parse models.dev JSON: %w", err)
	}
	return raw, nil
}

// upstreamProvider mirrors the subset of models.dev's schema we care about.
// Only fields we use; extra keys are tolerated via encoding/json's default
// behaviour of ignoring unknown fields.
type upstreamProvider struct {
	ID     string                   `json:"id"`
	Name   string                   `json:"name"`
	Models map[string]upstreamModel `json:"models"`
}

type upstreamModel struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Attachment  bool                `json:"attachment"`
	Reasoning   bool                `json:"reasoning"`
	ToolCall    bool                `json:"tool_call"`
	Temperature bool                `json:"temperature"`
	ReleaseDate string              `json:"release_date"`
	LastUpdated string              `json:"last_updated"`
	Modalities  upstreamModalities  `json:"modalities"`
	Cost        upstreamCost        `json:"cost"`
	Limit       upstreamLimit       `json:"limit"`
}

type upstreamModalities struct {
	Input  []string `json:"input"`
	Output []string `json:"output"`
}

type upstreamCost struct {
	Input     float64 `json:"input"`
	Output    float64 `json:"output"`
	Reasoning float64 `json:"reasoning"`
	CacheRead float64 `json:"cache_read"`
	CacheWrite float64 `json:"cache_write"`
}

type upstreamLimit struct {
	Context int `json:"context"`
	Output  int `json:"output"`
}

// mergeOne reads the YAML at path, merges upstream data into it according to
// the provider's sync policy, and writes it back. Returns true if anything
// changed. In -dry-run mode the file is not rewritten but the computed
// change count is still reported.
func mergeOne(path string, up upstreamProvider, dryRun bool) (bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}
	var prov catalogtypes.CatalogProvider
	if err := yaml.Unmarshal(data, &prov); err != nil {
		return false, err
	}

	policy := prov.Sync
	if policy == nil {
		policy = &catalogtypes.SyncPolicy{Policy: "latest_per_family"}
	}

	next := applySync(prov, up, policy)
	if sameProvider(prov, next) {
		return false, nil
	}

	if dryRun {
		fmt.Printf("diff  %-16s (%d models → %d)\n", prov.Name, len(prov.Models), len(next.Models))
		return true, nil
	}

	out, err := yaml.Marshal(&next)
	if err != nil {
		return false, err
	}
	if err := os.WriteFile(path, out, 0o644); err != nil {
		return false, err
	}
	fmt.Printf("sync  %-16s (%d → %d models)\n", prov.Name, len(prov.Models), len(next.Models))
	return true, nil
}

// applySync returns a new CatalogProvider whose model list is derived from
// upstream per the policy, while preserving hand-edited per-model overrides
// (Headers / BaseURL) and provider-level fields models.dev doesn't supply.
func applySync(prov catalogtypes.CatalogProvider, up upstreamProvider, policy *catalogtypes.SyncPolicy) catalogtypes.CatalogProvider {
	existingByID := make(map[string]catalogtypes.CatalogModel, len(prov.Models))
	for _, m := range prov.Models {
		existingByID[m.ID] = m
	}

	var re *regexp.Regexp
	if policy.FamilyRegex != "" {
		re = regexp.MustCompile(policy.FamilyRegex)
	} else {
		re = DefaultFamilyRegex
	}
	var stripRe *regexp.Regexp
	if policy.FamilyStrip != "" {
		stripRe = regexp.MustCompile(policy.FamilyStrip)
	}

	// Convert upstream to our shape first (so manual pin/exclude can work on
	// the same data representation).
	converted := make([]catalogtypes.CatalogModel, 0, len(up.Models))
	for id, um := range up.Models {
		if slices.Contains(policy.Excluded, id) {
			continue
		}
		cm := convertUpstream(id, um)
		// Make sync additive: every non-zero field from the hand-authored
		// entry wins over upstream. Upstream only fills in where local is
		// silent, so a sync run can introduce new fields models.dev knows
		// about (release_date, last_updated, brand-new capability flags) but
		// can never overwrite curated names, pricing, or overrides.
		if existing, ok := existingByID[id]; ok {
			cm = preferLocal(cm, existing)
		}
		converted = append(converted, cm)
	}

	// Pinned IDs that upstream dropped — keep the existing entry so users
	// don't lose a provisional preview model just because models.dev updated.
	for _, pinned := range policy.Pinned {
		if existing, ok := existingByID[pinned]; ok {
			if !containsModel(converted, pinned) {
				converted = append(converted, existing)
			}
		}
	}

	// Apply latest-per-family when requested.
	kept := converted
	switch policy.Policy {
	case "manual":
		kept = filterToPinned(converted, policy.Pinned)
	default: // "" | "latest_per_family"
		kept = keepLatestPerFamily(converted, policy.Pinned, re, stripRe)
	}

	sort.Slice(kept, func(i, j int) bool { return kept[i].ID < kept[j].ID })
	prov.Models = kept
	return prov
}

func convertUpstream(id string, um upstreamModel) catalogtypes.CatalogModel {
	return catalogtypes.CatalogModel{
		ID:                        id,
		Name:                      um.Name,
		Limit:                     um.Limit.Context,
		ReservedTokens:            um.Limit.Output,
		SupportsStreaming:         true, // models.dev doesn't track this; default to true
		SupportsImage:             slices.Contains(um.Modalities.Input, "image"),
		SupportsVideo:             slices.Contains(um.Modalities.Input, "video"),
		SupportsReasoning:         um.Reasoning,
		SupportsToolSearch:        um.ToolCall,
		SupportsCustomTemperature: um.Temperature,
		ReleaseDate:               um.ReleaseDate,
		LastUpdated:               um.LastUpdated,
		Pricing: catalogtypes.CatalogPricing{
			Input:      um.Cost.Input,
			Output:     um.Cost.Output,
			CacheRead:  um.Cost.CacheRead,
			CacheWrite: um.Cost.CacheWrite,
		},
	}
}

// preferLocal returns a CatalogModel where every non-zero field of `local`
// shadows `upstream`. For booleans, only `true` is treated as "set" — local
// can never force a flag off, only on. Pricing is all-or-nothing: if local
// has any rate set (or marks subscription), the whole local Pricing block
// wins; otherwise upstream's rates come through. This keeps weekly sync
// strictly additive over hand-authored data.
func preferLocal(upstream, local catalogtypes.CatalogModel) catalogtypes.CatalogModel {
	result := upstream
	if local.Name != "" {
		result.Name = local.Name
	}
	if local.BaseURL != "" {
		result.BaseURL = local.BaseURL
	}
	if len(local.Headers) > 0 {
		result.Headers = local.Headers
	}
	if local.Limit != 0 {
		result.Limit = local.Limit
	}
	if local.ReservedTokens != 0 {
		result.ReservedTokens = local.ReservedTokens
	}
	if local.SupportsStreaming {
		result.SupportsStreaming = true
	}
	if local.SupportsImage {
		result.SupportsImage = true
	}
	if local.SupportsVideo {
		result.SupportsVideo = true
	}
	if local.SupportsReasoning {
		result.SupportsReasoning = true
	}
	if local.SupportsAdaptiveThinking {
		result.SupportsAdaptiveThinking = true
	}
	if local.SupportsIncludeReasoning {
		result.SupportsIncludeReasoning = true
	}
	if local.SupportsToolSearch {
		result.SupportsToolSearch = true
	}
	if local.SupportsCustomTemperature {
		result.SupportsCustomTemperature = true
	}
	if local.SupportsImageGeneration {
		result.SupportsImageGeneration = true
	}
	if local.SupportsVideoGeneration {
		result.SupportsVideoGeneration = true
	}
	if local.DefaultSelection {
		result.DefaultSelection = true
	}
	if local.GenerationType != "" {
		result.GenerationType = local.GenerationType
	}
	if local.PerUnitCost != 0 {
		result.PerUnitCost = local.PerUnitCost
	}
	if local.ReleaseDate != "" {
		result.ReleaseDate = local.ReleaseDate
	}
	if local.LastUpdated != "" {
		result.LastUpdated = local.LastUpdated
	}
	if !isZeroPricing(local.Pricing) {
		result.Pricing = local.Pricing
	}
	return result
}

func isZeroPricing(p catalogtypes.CatalogPricing) bool {
	return p.Input == 0 && p.Output == 0 && p.CacheRead == 0 && p.CacheWrite == 0 && !p.IsSubscription
}

func containsModel(ms []catalogtypes.CatalogModel, id string) bool {
	for _, m := range ms {
		if m.ID == id {
			return true
		}
	}
	return false
}

func filterToPinned(models []catalogtypes.CatalogModel, pinned []string) []catalogtypes.CatalogModel {
	pinSet := make(map[string]struct{}, len(pinned))
	for _, id := range pinned {
		pinSet[id] = struct{}{}
	}
	out := models[:0]
	for _, m := range models {
		if _, ok := pinSet[m.ID]; ok {
			out = append(out, m)
		}
	}
	return out
}

// keepLatestPerFamily groups models by their family key and keeps only the
// newest-dated one per group. Pinned IDs bypass the filter entirely.
// stripRe (optional) is applied to each ID before family-key extraction.
func keepLatestPerFamily(models []catalogtypes.CatalogModel, pinned []string, re *regexp.Regexp, stripRe *regexp.Regexp) []catalogtypes.CatalogModel {
	pinSet := make(map[string]struct{}, len(pinned))
	for _, id := range pinned {
		pinSet[id] = struct{}{}
	}

	families := make(map[string][]modelWithMeta)
	for i, m := range models {
		fam := familyKey(re, stripRe, m.ID)
		families[fam] = append(families[fam], modelWithMeta{
			ID:          m.ID,
			ReleaseDate: m.ReleaseDate,
			LastUpdated: m.LastUpdated,
			Index:       i,
		})
	}

	kept := make(map[string]struct{})
	for _, group := range families {
		// Pinned IDs shadow their family: if any member of the group is
		// pinned, every non-pinned sibling is dropped. This lets a hand-
		// authored undated ID (e.g. `claude-haiku-4-5`) win over a dated
		// upstream variant (e.g. `claude-haiku-4-5-20251001`) without
		// needing an exhaustive `excluded` list.
		anyPinned := false
		for _, m := range group {
			if _, ok := pinSet[m.ID]; ok {
				anyPinned = true
				kept[m.ID] = struct{}{}
			}
		}
		if anyPinned {
			continue
		}
		best := pickLatest(group)
		kept[best.ID] = struct{}{}
	}
	for id := range pinSet {
		kept[id] = struct{}{}
	}

	out := make([]catalogtypes.CatalogModel, 0, len(kept))
	for _, m := range models {
		if _, ok := kept[m.ID]; ok {
			out = append(out, m)
		}
	}
	return out
}

func sameProvider(a, b catalogtypes.CatalogProvider) bool {
	ab, _ := yaml.Marshal(&a)
	bb, _ := yaml.Marshal(&b)
	return string(ab) == string(bb)
}

func readExclude(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	var lines []string
	for _, line := range strings.Split(string(data), "\n") {
		if s := strings.TrimSpace(line); s != "" && !strings.HasPrefix(s, "#") {
			lines = append(lines, s)
		}
	}
	return lines, nil
}
