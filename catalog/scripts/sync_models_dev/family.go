package main

import (
	"regexp"
	"strings"
	"time"
)

// DefaultFamilyRegex collapses date stamps, version suffixes, and -preview /
// -beta / -latest markers to extract the family key. First capture group is
// the family.
//
// Handles:
//
//	claude-opus-4-7                     → claude-opus
//	claude-opus-4-1                     → claude-opus
//	kimi-k2-0905-preview                → kimi-k2
//	kimi-k2-0712                        → kimi-k2
//	gpt-4o-2025-08-01                   → gpt-4o
//	gemini-2.5-pro                      → gemini-2.5-pro  (no trailing noise)
//	grok-4.1-fast                       → grok-4.1-fast    (no trailing noise)
//
// Caller-provided overrides win (SyncPolicy.FamilyRegex) — this is just the
// default used when the provider YAML doesn't declare one.
var DefaultFamilyRegex = regexp.MustCompile(
	`^(.*?)(?:-\d{4}-\d{2}-\d{2}|-\d{4,8}|-v?\d+(?:[.-]\d+)+|-preview|-beta|-latest)+$`,
)

// familyKey returns the family group for a model ID.
//
// The input is first passed through `stripRe` (per-provider mid-string
// version stripper — optional) to collapse conventions like Google's
// `gemini-2.5-flash` / `gemini-3-flash` or OpenAI's `gpt-5.4-codex`. The
// cleaned string is then matched against `re` to strip trailing date /
// version / preview suffixes.
//
// If neither pass changes anything the full ID is returned, so pruning
// only removes exact duplicates rather than collapsing unrelated models.
func familyKey(re *regexp.Regexp, stripRe *regexp.Regexp, id string) string {
	cleaned := id
	if stripRe != nil {
		cleaned = stripRe.ReplaceAllString(cleaned, "")
	}
	if m := re.FindStringSubmatch(cleaned); len(m) >= 2 && m[1] != "" {
		return m[1]
	}
	return cleaned
}

// pickLatest picks the newest model from a family by release_date, falling
// back to last_updated, then lexicographically by ID so a family with no
// dates still has a deterministic winner. Returns the winning model.
func pickLatest(models []modelWithMeta) modelWithMeta {
	best := models[0]
	for _, m := range models[1:] {
		if isNewer(m, best) {
			best = m
		}
	}
	return best
}

type modelWithMeta struct {
	ID          string
	ReleaseDate string
	LastUpdated string
	Index       int // index in the raw provider data, tiebreak on equal dates
}

func isNewer(a, b modelWithMeta) bool {
	ar, aok := parseDate(a.ReleaseDate)
	br, bok := parseDate(b.ReleaseDate)
	if aok && bok && !ar.Equal(br) {
		return ar.After(br)
	}
	if aok != bok {
		return aok // dated beats undated
	}
	au, aok := parseDate(a.LastUpdated)
	bu, bok := parseDate(b.LastUpdated)
	if aok && bok && !au.Equal(bu) {
		return au.After(bu)
	}
	// Both undated → lexicographic stable order.
	return strings.Compare(a.ID, b.ID) > 0
}

func parseDate(s string) (time.Time, bool) {
	if s == "" {
		return time.Time{}, false
	}
	// models.dev uses YYYY-MM or YYYY-MM-DD.
	for _, layout := range []string{"2006-01-02", "2006-01"} {
		if t, err := time.Parse(layout, s); err == nil {
			return t, true
		}
	}
	return time.Time{}, false
}
