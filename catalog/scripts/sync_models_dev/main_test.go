package main

import (
	"reflect"
	"regexp"
	"strings"
	"testing"

	catalogtypes "github.com/hellenic-development/plexon.ai/catalog/scripts"
)

// fill sets every field reachable from v to a non-zero value, so a merge that
// forgets one field shows up as that field going back to zero.
func fill(t *testing.T, v reflect.Value) {
	t.Helper()
	switch v.Kind() {
	case reflect.Bool:
		v.SetBool(true)
	case reflect.String:
		v.SetString("x")
	case reflect.Int, reflect.Int64:
		v.SetInt(1)
	case reflect.Float64:
		v.SetFloat(1)
	case reflect.Ptr:
		v.Set(reflect.New(v.Type().Elem()))
		fill(t, v.Elem())
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fill(t, v.Field(i))
		}
	case reflect.Slice:
		elem := reflect.New(v.Type().Elem()).Elem()
		fill(t, elem)
		v.Set(reflect.Append(reflect.MakeSlice(v.Type(), 0, 1), elem))
	case reflect.Map:
		m := reflect.MakeMap(v.Type())
		k := reflect.New(v.Type().Key()).Elem()
		fill(t, k)
		e := reflect.New(v.Type().Elem()).Elem()
		fill(t, e)
		m.SetMapIndex(k, e)
		v.Set(m)
	default:
		t.Fatalf("fill: unhandled kind %s; teach fill about it so the merge guard keeps covering every field", v.Kind())
	}
}

// TestPreferLocalKeepsEveryAuthoredField is the guard for the merge rule the
// weekly sync relies on: a hand-authored field must never be blanked by an
// upstream entry that does not know it. It fills every field of CatalogModel by
// reflection, so a field added to types.go without a matching line in
// preferLocal fails here instead of vanishing from a provider YAML on the next
// Monday (thinking_always_on and rejects_forced_tool_choice did exactly that
// between 2026-09-22 and 2026-10-17).
func TestPreferLocalKeepsEveryAuthoredField(t *testing.T) {
	var local catalogtypes.CatalogModel
	fill(t, reflect.ValueOf(&local).Elem())
	upstream := catalogtypes.CatalogModel{ID: local.ID}

	got := preferLocal(upstream, local)

	gv, lv := reflect.ValueOf(got), reflect.ValueOf(local)
	for i := 0; i < lv.NumField(); i++ {
		name := lv.Type().Field(i).Name
		if !reflect.DeepEqual(gv.Field(i).Interface(), lv.Field(i).Interface()) {
			t.Errorf("preferLocal drops the authored %s: got %#v, want %#v", name, gv.Field(i).Interface(), lv.Field(i).Interface())
		}
	}
}

// TestPinWarnsWhenItHidesANewerSibling: a pinned model shadows its whole family,
// which is how an undated hand-authored id beats a dated upstream variant. The
// same rule silently hid claude-fable-5-1 behind pinned claude-fable-5 for six
// weeks, so the filter now says so when the hidden sibling is strictly newer.
func TestPinWarnsWhenItHidesANewerSibling(t *testing.T) {
	strip := regexp.MustCompile(`(-\d+)+`)
	models := []catalogtypes.CatalogModel{
		{ID: "claude-fable-5", ReleaseDate: "2026-06-09"},
		{ID: "claude-fable-5-1", ReleaseDate: "2026-09-01"},
	}

	kept, warnings := keepLatestPerFamily(models, []string{"claude-fable-5"}, DefaultFamilyRegex, strip)

	if len(kept) != 1 || kept[0].ID != "claude-fable-5" {
		t.Fatalf("kept = %v, want only the pinned claude-fable-5", kept)
	}
	if len(warnings) != 1 || !strings.Contains(warnings[0], "claude-fable-5 hides newer claude-fable-5-1 (2026-09-01)") {
		t.Fatalf("warnings = %q, want one naming the pin and the hidden successor", warnings)
	}
}

// TestPinStaysQuietForDatedVariantsAndTies covers the two shapes the shadow rule
// exists for: a dated upstream variant of the same release (claude-haiku-4-5 vs
// claude-haiku-4-5-20251001) and a sibling with no date at all. Neither is a
// successor, so neither is reported.
func TestPinStaysQuietForDatedVariantsAndTies(t *testing.T) {
	strip := regexp.MustCompile(`(-\d+)+`)
	models := []catalogtypes.CatalogModel{
		{ID: "claude-haiku-4-5", ReleaseDate: "2025-10-15"},
		{ID: "claude-haiku-4-5-20251001", ReleaseDate: "2025-10-15"},
		{ID: "claude-haiku-4-6"},
	}

	kept, warnings := keepLatestPerFamily(models, []string{"claude-haiku-4-5"}, DefaultFamilyRegex, strip)

	if len(kept) != 1 || kept[0].ID != "claude-haiku-4-5" {
		t.Fatalf("kept = %v, want only the pinned claude-haiku-4-5", kept)
	}
	if len(warnings) != 0 {
		t.Fatalf("warnings = %q, want none for a same-day variant or an undated sibling", warnings)
	}
}

// TestExcludedPrefixSparesPins: a trailing "*" retires a lineage by prefix, and
// a pinned id inside that prefix survives, which is what lets `gpt-5*` drop the
// GPT-5 generation while the pinned gpt-5.3-codex and gpt-5-image rows stay.
func TestExcludedPrefixSparesPins(t *testing.T) {
	policy := &catalogtypes.SyncPolicy{
		Pinned:   []string{"gpt-5.3-codex"},
		Excluded: []string{"gpt-5*", "o4-mini"},
	}
	for id, want := range map[string]bool{
		"gpt-5.6":       true,
		"gpt-5-mini":    true,
		"gpt-5.3-codex": false,
		"gpt-6-astra":   false,
		"o4-mini":       true,
		"o4-mini-high":  false,
	} {
		if got := isExcluded(policy, id); got != want {
			t.Errorf("isExcluded(%q) = %v, want %v", id, got, want)
		}
	}
}
