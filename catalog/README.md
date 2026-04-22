# Plexon AI Catalog

This directory hosts the provider/model registry consumed by the Plexon AI
client at runtime. Moving this data out of the Go binary means:

- New models ship to users without a client release.
- Community contributions arrive via pull requests, not issues asking us
  to bump a Go file.
- A weekly CI job keeps third-party providers in sync with
  [`models.dev`](https://models.dev), the community-maintained model database.

## Layout

```
catalog/
├── providers/               # One YAML file per provider
│   ├── anthropic.yaml
│   ├── openai.yaml
│   ├── google.yaml
│   ├── …
│   ├── plexon.yaml          # hand-authored (in .sync-exclude)
│   └── claude-desktop.yaml  # hand-authored (in .sync-exclude)
├── snapshots/
│   └── latest.json          # bundled by CI, fetched by Plexon clients
├── scripts/
│   ├── sync_models_dev/     # go run ./scripts/sync_models_dev
│   └── build_snapshot/      # go run ./scripts/build_snapshot
├── .sync-exclude            # provider names the weekly sync must never touch
└── README.md
```

YAML is the authoring format (diffs scope to one provider per PR, bot vs
human don't collide). JSON is the wire / embedded format (clients never
parse YAML → removes the `yaml.v3` dep from the boot path and shrinks
startup latency). `scripts/build_snapshot` walks `providers/*.yaml` and
emits `snapshots/latest.json`.

## How Plexon clients consume this

The Plexon desktop app fetches
`https://raw.githubusercontent.com/hellenic-development/plexon.ai/main/catalog/snapshots/latest.json`
with an `If-None-Match` header, caches it for 24 hours at
`~/.plexon/catalog/catalog.json`, and falls back to an embedded snapshot
shipped in the binary if the network is unavailable. Details in
`docs/providers/DYNAMIC_CATALOG.md` inside the plexon repo.

## Adding or editing a provider

1. Fork this repo and edit `providers/<provider>.yaml`.
2. Run `go run ./catalog/scripts/build_snapshot` (or wait for CI) — this
   refreshes `snapshots/latest.json`.
3. Open a PR. The catalog sync workflow will validate the schema and block
   merge on errors.

### Subscription models

Subscription-based models (e.g. `kimi-for-coding`) are fully supported via
`pricing.is_subscription: true`. The client suppresses per-request cost
calculation for these models, falling back to your provider's quota.

### Latest-per-family sync rule

`scripts/sync_models_dev` pulls fresh data from `models.dev/api.json` every
Monday. For each provider not in `.sync-exclude`, it groups models by their
"family key" (the model ID with trailing date / version / preview suffixes
stripped) and keeps only the newest model per family (by `release_date`).
Pinned model IDs survive this pruning — see the `sync.pinned` array in each
provider YAML. Models explicitly listed in `sync.excluded` are always
dropped.

## Hand-authored providers (`.sync-exclude`)

Two providers are never touched by the weekly sync:

- `plexon` — the Plexon meta-provider. Its routing (key pool, tier-based
  model selection) lives in Plexon server code; the catalog entry only
  describes the three virtual models it exposes.
- `claude-desktop` — the local Claude Desktop CLI integration. Capabilities
  here track what the Claude Code binary supports, not models.dev data.

Add more names to `.sync-exclude` (one per line) if you introduce further
meta-providers.
