package main

import (
	"strings"
	"testing"

	catalogtypes "github.com/hellenic-development/plexon.ai/catalog/scripts"
)

// TestSyncWriteKeepsCommentsAndOrder: a sync rewrite must leave the authored
// notes and the authored row order where they were, and append what is new at
// the end. Before this, minimax.yaml lost 190 comment lines and had its rows
// re-sorted on every weekly run.
func TestSyncWriteKeepsCommentsAndOrder(t *testing.T) {
	prev := []byte(`name: p
# the default is the flagship on purpose
default_model: b
models:
    # b first: it is the default
    - id: b
      name: B
      pricing:
        input: 1 # base band only
    - id: a
      name: A
`)
	var prov catalogtypes.CatalogProvider
	if err := yamlUnmarshal(prev, &prov); err != nil {
		t.Fatal(err)
	}
	next := prov
	next.Models = orderLike(prov.Models, []catalogtypes.CatalogModel{
		{ID: "a", Name: "A"},
		{ID: "c", Name: "C"},
		{ID: "b", Name: "B", Pricing: catalogtypes.CatalogPricing{Input: 1}},
	})

	out, err := marshalKeepingComments(prev, &next)
	if err != nil {
		t.Fatal(err)
	}
	text := string(out)
	for _, want := range []string{"# the default is the flagship on purpose", "# b first: it is the default", "# base band only"} {
		if !strings.Contains(text, want) {
			t.Errorf("rewrite dropped the comment %q:\n%s", want, text)
		}
	}
	b, a, c := strings.Index(text, "id: b"), strings.Index(text, "id: a"), strings.Index(text, "id: c")
	if !(b < a && a < c) {
		t.Errorf("rows re-ordered: b=%d a=%d c=%d, want the authored order with c appended:\n%s", b, a, c, text)
	}
}
