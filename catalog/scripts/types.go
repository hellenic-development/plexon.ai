// Package catalogtypes defines the YAML / JSON shape shared by the sync and
// build-snapshot scripts. Kept deliberately small — every field here must
// round-trip unchanged between `providers/*.yaml` and `snapshots/latest.json`
// so that a Plexon client's `common/catalog.Catalog` unmarshalling stays in
// lockstep with the authored YAML.
//
// When the Go client's `common/catalog` schema changes, update this file in
// the same commit. The two packages are not wired together — they agree by
// convention on field names and shapes. `schema_version` lives at the
// top-level Catalog so a client can refuse to read a snapshot whose schema
// post-dates its binary.
package catalogtypes

// SchemaVersion is the schema version written into every snapshot and
// validated against the client's hard-coded expectation at load time.
const SchemaVersion = 1

// Catalog is the top-level container bundled in snapshots/latest.json.
type Catalog struct {
	Version   int                `json:"version" yaml:"version"`
	Generated string             `json:"generated,omitempty" yaml:"generated,omitempty"`
	Providers []CatalogProvider  `json:"providers" yaml:"providers,omitempty"`
}

// CatalogProvider describes one provider and its models.
type CatalogProvider struct {
	Name              string              `json:"name" yaml:"name"`
	DisplayName       string              `json:"display_name,omitempty" yaml:"display_name,omitempty"`
	Compatibility     string              `json:"compatibility" yaml:"compatibility"`
	RequiresAPIKey    bool                `json:"requires_api_key,omitempty" yaml:"requires_api_key,omitempty"`
	SupportsOAuth     bool                `json:"supports_oauth,omitempty" yaml:"supports_oauth,omitempty"`
	SupportsWebSearch bool                `json:"supports_web_search,omitempty" yaml:"supports_web_search,omitempty"`
	BaseURL           string              `json:"base_url,omitempty" yaml:"base_url,omitempty"`
	DefaultModel      string              `json:"default_model,omitempty" yaml:"default_model,omitempty"`
	BuiltinTools      map[string][]string `json:"builtin_tools,omitempty" yaml:"builtin_tools,omitempty"`
	Regions           []CatalogRegion     `json:"regions,omitempty" yaml:"regions,omitempty"`
	Sync              *SyncPolicy         `json:"sync,omitempty" yaml:"sync,omitempty"`
	Models            []CatalogModel      `json:"models" yaml:"models"`
}

// CatalogModel mirrors the PricingModel fields the Plexon client consumes.
type CatalogModel struct {
	ID                        string            `json:"id" yaml:"id"`
	Name                      string            `json:"name,omitempty" yaml:"name,omitempty"`
	BaseURL                   string            `json:"base_url,omitempty" yaml:"base_url,omitempty"`
	Headers                   map[string]string `json:"headers,omitempty" yaml:"headers,omitempty"`
	Limit                     int               `json:"limit,omitempty" yaml:"limit,omitempty"`
	ReservedTokens            int               `json:"reserved_tokens,omitempty" yaml:"reserved_tokens,omitempty"`
	SupportsStreaming         bool              `json:"supports_streaming,omitempty" yaml:"supports_streaming,omitempty"`
	SupportsImage             bool              `json:"supports_image,omitempty" yaml:"supports_image,omitempty"`
	SupportsVideo             bool              `json:"supports_video,omitempty" yaml:"supports_video,omitempty"`
	SupportsReasoning         bool              `json:"supports_reasoning,omitempty" yaml:"supports_reasoning,omitempty"`
	SupportsAdaptiveThinking  bool              `json:"supports_adaptive_thinking,omitempty" yaml:"supports_adaptive_thinking,omitempty"`
	SupportsIncludeReasoning  bool              `json:"supports_include_reasoning,omitempty" yaml:"supports_include_reasoning,omitempty"`
	SupportsToolSearch        bool              `json:"supports_tool_search,omitempty" yaml:"supports_tool_search,omitempty"`
	SupportsCustomTemperature bool              `json:"supports_custom_temperature,omitempty" yaml:"supports_custom_temperature,omitempty"`
	SupportsImageGeneration   bool              `json:"supports_image_generation,omitempty" yaml:"supports_image_generation,omitempty"`
	SupportsVideoGeneration   bool              `json:"supports_video_generation,omitempty" yaml:"supports_video_generation,omitempty"`
	GenerationType            string            `json:"generation_type,omitempty" yaml:"generation_type,omitempty"`
	PerUnitCost               float64           `json:"per_unit_cost,omitempty" yaml:"per_unit_cost,omitempty"`
	DefaultSelection          bool              `json:"default_selection,omitempty" yaml:"default_selection,omitempty"`
	ReleaseDate               string            `json:"release_date,omitempty" yaml:"release_date,omitempty"`
	LastUpdated               string            `json:"last_updated,omitempty" yaml:"last_updated,omitempty"`
	Pricing                   CatalogPricing    `json:"pricing,omitempty" yaml:"pricing,omitempty"`
}

// CatalogPricing is the per-1M-token rate card for a model.
type CatalogPricing struct {
	Input          float64 `json:"input,omitempty" yaml:"input,omitempty"`
	Output         float64 `json:"output,omitempty" yaml:"output,omitempty"`
	CacheWrite     float64 `json:"cache_write,omitempty" yaml:"cache_write,omitempty"`
	CacheRead      float64 `json:"cache_read,omitempty" yaml:"cache_read,omitempty"`
	IsSubscription bool    `json:"is_subscription,omitempty" yaml:"is_subscription,omitempty"`
}

// CatalogRegion defines a selectable regional endpoint for a provider.
type CatalogRegion struct {
	ID      string `json:"id" yaml:"id"`
	Label   string `json:"label" yaml:"label"`
	BaseURL string `json:"base_url" yaml:"base_url"`
}

// SyncPolicy is the per-provider override that the weekly models.dev sync
// consults. A nil value means "use package defaults" (latest-per-family,
// no pins, no exclusions).
type SyncPolicy struct {
	// Policy: "latest_per_family" (default) or "manual" (only pinned models kept).
	Policy string `yaml:"policy,omitempty"`
	// FamilyRegex overrides the default family-key extraction regex.
	// First capture group == family key. See scripts/sync_models_dev/family.go.
	FamilyRegex string `yaml:"family_regex,omitempty"`
	// FamilyStrip is a regex applied to each model ID BEFORE FamilyRegex.
	// Every match is removed, so mid-string version tokens can be collapsed.
	//
	// Examples:
	//   google    family_strip: '-\d+(\.\d+)?'   → gemini-2.5-flash  → gemini-flash
	//   openai    family_strip: '-\d+(\.\d+)?'   → gpt-5.1-codex      → gpt-codex
	//   moonshot  family_strip: '-k\d+(\.\d+)?'  → kimi-k2.5-thinking → kimi-thinking
	//
	// Without family_strip these providers leak every minor version into
	// the UI because the version sits mid-string, not at the end.
	FamilyStrip string `yaml:"family_strip,omitempty"`
	// Pinned model IDs are always kept, even if older than a family's latest.
	Pinned []string `yaml:"pinned,omitempty"`
	// Excluded model IDs are always dropped.
	Excluded []string `yaml:"excluded,omitempty"`
}
