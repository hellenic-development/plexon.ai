package main

import (
	"bytes"
	"sort"

	"gopkg.in/yaml.v3"

	catalogtypes "github.com/hellenic-development/plexon.ai/catalog/scripts"
)

// orderLike returns next in the order the file already has: every id that
// existed keeps its place, new ids follow sorted. A provider file's row order
// is authored (text models before speech, the default first), and sorting it
// made every sync rewrite the file and flagged "changed" providers whose model
// list had not moved (minimax, every Monday from 2026-07 to 2026-10).
func orderLike(existing, next []catalogtypes.CatalogModel) []catalogtypes.CatalogModel {
	byID := make(map[string]catalogtypes.CatalogModel, len(next))
	for _, m := range next {
		byID[m.ID] = m
	}
	out := make([]catalogtypes.CatalogModel, 0, len(next))
	seen := make(map[string]struct{}, len(next))
	for _, m := range existing {
		if n, ok := byID[m.ID]; ok {
			out = append(out, n)
			seen[m.ID] = struct{}{}
		}
	}
	var added []catalogtypes.CatalogModel
	for _, m := range next {
		if _, ok := seen[m.ID]; !ok {
			added = append(added, m)
		}
	}
	sort.Slice(added, func(i, j int) bool { return added[i].ID < added[j].ID })
	return append(out, added...)
}

// marshalKeepingComments encodes next the way yaml.Marshal does, then copies
// the comments from the file's previous text onto the matching nodes: mapping
// keys by name, sequence items by `id` (or by value for scalar lists). A plain
// Marshal threw every authored note away, so a provider file's band notes and
// "why this flag is off" explanations lasted exactly until the next upstream
// change to that provider. Comments on a node that no longer exists (a removed
// model) are dropped with it.
func marshalKeepingComments(prev []byte, next any) ([]byte, error) {
	fresh, err := yaml.Marshal(next)
	if err != nil {
		return nil, err
	}
	var oldDoc, newDoc yaml.Node
	if err := yaml.Unmarshal(prev, &oldDoc); err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(fresh, &newDoc); err != nil {
		return nil, err
	}
	transplantComments(&oldDoc, &newDoc)
	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	if err := enc.Encode(&newDoc); err != nil {
		return nil, err
	}
	if err := enc.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func transplantComments(old, next *yaml.Node) {
	if old == nil || next == nil || old.Kind != next.Kind {
		return
	}
	copyComments(old, next)
	switch next.Kind {
	case yaml.DocumentNode:
		if len(old.Content) > 0 && len(next.Content) > 0 {
			transplantComments(old.Content[0], next.Content[0])
		}
	case yaml.MappingNode:
		for i := 0; i+1 < len(next.Content); i += 2 {
			key, value := next.Content[i], next.Content[i+1]
			oldKey, oldValue := mappingEntry(old, key.Value)
			if oldKey == nil {
				continue
			}
			copyComments(oldKey, key)
			transplantComments(oldValue, value)
		}
	case yaml.SequenceNode:
		for _, item := range next.Content {
			if match := sequenceItem(old, item); match != nil {
				transplantComments(match, item)
			}
		}
	}
}

func copyComments(from, to *yaml.Node) {
	if to.HeadComment == "" {
		to.HeadComment = from.HeadComment
	}
	if to.LineComment == "" {
		to.LineComment = from.LineComment
	}
	if to.FootComment == "" {
		to.FootComment = from.FootComment
	}
}

// mappingEntry returns the key and value nodes for key in a mapping node.
func mappingEntry(mapping *yaml.Node, key string) (*yaml.Node, *yaml.Node) {
	for i := 0; i+1 < len(mapping.Content); i += 2 {
		if mapping.Content[i].Value == key {
			return mapping.Content[i], mapping.Content[i+1]
		}
	}
	return nil, nil
}

// sequenceItem finds the item of seq that corresponds to want: a mapping with
// the same `id`, or a scalar with the same value.
func sequenceItem(seq *yaml.Node, want *yaml.Node) *yaml.Node {
	identity := func(n *yaml.Node) string {
		if n.Kind == yaml.MappingNode {
			if _, v := mappingEntry(n, "id"); v != nil {
				return "id:" + v.Value
			}
			return ""
		}
		return "value:" + n.Value
	}
	id := identity(want)
	if id == "" {
		return nil
	}
	for _, item := range seq.Content {
		if identity(item) == id {
			return item
		}
	}
	return nil
}

// yamlUnmarshal is yaml.Unmarshal under a name the tests can use without a
// second import of the package.
func yamlUnmarshal(data []byte, v any) error { return yaml.Unmarshal(data, v) }
