// Command build_snapshot walks providers/*.yaml and writes
// snapshots/latest.json — the bundled JSON Plexon clients fetch.
//
//	go run ./scripts/build_snapshot
//
// Exit codes: 0 on success, 1 on any parse / write error.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gopkg.in/yaml.v3"

	catalogtypes "github.com/hellenic-development/plexon.ai/catalog/scripts"
)

func main() {
	providersDir := flag.String("providers", "providers", "directory containing <name>.yaml files")
	out := flag.String("out", "snapshots/latest.json", "output path for the bundled snapshot")
	flag.Parse()

	entries, err := os.ReadDir(*providersDir)
	if err != nil {
		log.Fatalf("read %s: %v", *providersDir, err)
	}

	// Deterministic provider order → diffable snapshots.
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })

	var providers []catalogtypes.CatalogProvider
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".yaml" {
			continue
		}
		path := filepath.Join(*providersDir, e.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("read %s: %v", path, err)
		}
		var p catalogtypes.CatalogProvider
		if err := yaml.Unmarshal(data, &p); err != nil {
			log.Fatalf("parse %s: %v", path, err)
		}
		// A provider without a name is a schema bug — surface it.
		if p.Name == "" {
			log.Fatalf("%s: missing required field 'name'", path)
		}
		// Strip per-provider sync policy from the snapshot; it's authoring-
		// side metadata only and clients don't need to transmit it.
		p.Sync = nil
		providers = append(providers, p)
	}

	cat := catalogtypes.Catalog{
		Version:   catalogtypes.SchemaVersion,
		Generated: time.Now().UTC().Format(time.RFC3339),
		Providers: providers,
	}

	if err := os.MkdirAll(filepath.Dir(*out), 0o755); err != nil {
		log.Fatalf("mkdir: %v", err)
	}
	buf, err := json.MarshalIndent(&cat, "", "  ")
	if err != nil {
		log.Fatalf("marshal: %v", err)
	}
	if err := os.WriteFile(*out, append(buf, '\n'), 0o644); err != nil {
		log.Fatalf("write %s: %v", *out, err)
	}
	fmt.Printf("wrote %d provider(s) to %s\n", len(providers), *out)
}
