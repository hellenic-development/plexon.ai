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
and caches it for 24 hours at `~/.plexon/catalog/catalog.json`. Two
fetch modes:

- **Background (every 24h):** `If-None-Match` with the saved ETag. The
  server answers 304 when nothing changed → cheap no-op.
- **Manual (☁ Refresh Catalog button):** ETag is *deliberately not* sent
  and a unique query string + `Cache-Control: no-cache` is added, so the
  CDN edge can't answer 304 with stale content. This matters right after
  a push: `raw.githubusercontent.com` caches per-edge for up to 5 min,
  and a user hitting Refresh within that window used to see no change.

If the fetch fails entirely (offline / blocked / GitHub down), the
client falls back to an embedded snapshot shipped inside the binary, and
if even the daemon isn't up yet, the UI has a tiny hardcoded
**Plexon 1.0** placeholder so the Settings dropdown is never empty.

The client then sends every chat request through the Plexon cloud server
with `Provider.InlineModel` attached — i.e. the client is the source of
truth for what models exist and what they support. The server does
**not** run its own catalog loader; missing `InlineModel` returns HTTP
426 `client_outdated` with a `plexon.ai/download` upgrade link. Full
mechanics in `docs/providers/DYNAMIC_CATALOG.md` inside the plexon repo.

## Adding or editing a provider

1. Fork this repo and edit `providers/<provider>.yaml`.
2. Run `go run ./catalog/scripts/build_snapshot` (or wait for CI) — this
   refreshes `snapshots/latest.json`.
3. Open a PR. The catalog sync workflow will validate the schema and block
   merge on errors.

## Running the Go tools locally

All commands below are run from `catalog/`. The scripts live in their own
Go module (`catalog/go.mod`) so they don't depend on the main plexon.ai
site code.

### Rebuild the bundled snapshot

Walks `providers/*.yaml` and emits `snapshots/latest.json`. Run this after
every manual YAML edit before committing. Doubles as a schema validator.

```bash
cd catalog
go run ./scripts/build_snapshot
```

### Pull fresh data from models.dev (what the weekly CI runs)

Fetches `https://models.dev/api.json`, applies the latest-per-family
filter per provider, and merges into each `providers/*.yaml`. Preserves
hand-edited `headers`, `base_url`, `regions`, `builtin_tools`,
`default_model`, and anything under `sync:`. Skips providers listed in
`.sync-exclude`.

```bash
cd catalog

# Normal sync — writes back to providers/*.yaml
go run ./scripts/sync_models_dev

# Preview only — prints the diff, no writes
go run ./scripts/sync_models_dev -dry-run

# Use a local fixture instead of hitting models.dev
go run ./scripts/sync_models_dev -fixture ./testdata/models-dev.json
```

### Typical maintainer flow

```bash
cd catalog

# 1. Pull any upstream changes + preview
go run ./scripts/sync_models_dev -dry-run

# 2. Apply them
go run ./scripts/sync_models_dev

# 3. Rebuild the snapshot clients actually fetch
go run ./scripts/build_snapshot

# 4. Commit and push
git add providers/*.yaml snapshots/latest.json
git commit -m "chore(catalog): weekly models.dev sync"
git push
```

### Flag reference

| Script              | Flag           | Default              | Purpose                                                   |
| ------------------- | -------------- | -------------------- | --------------------------------------------------------- |
| `build_snapshot`    | `-providers`   | `providers`          | Directory of provider YAMLs to read                       |
| `build_snapshot`    | `-out`         | `snapshots/latest.json` | Output path for the bundled JSON                       |
| `sync_models_dev`   | `-providers`   | `providers`          | Directory of provider YAMLs to merge into                 |
| `sync_models_dev`   | `-exclude`     | `.sync-exclude`      | File listing provider names the sync must never touch    |
| `sync_models_dev`   | `-fixture`     | _(none)_             | Read from local JSON instead of hitting models.dev       |
| `sync_models_dev`   | `-dry-run`     | `false`              | Print changes without writing                             |

### Subscription models

Subscription-based models (e.g. `kimi-for-coding`) are fully supported via
`pricing.is_subscription: true`. The client suppresses per-request cost
calculation for these models, falling back to your provider's quota.

### Latest-per-family sync rule

`scripts/sync_models_dev` pulls fresh data from `models.dev/api.json` every
Monday. For each provider not in `.sync-exclude`, it groups models by their
"family key" (the model ID after optional `family_strip` removal, with
trailing date / version / preview suffixes stripped by `family_regex`) and
keeps only the newest model per family (by `release_date`).

Per-provider `sync:` block supports four knobs:

| Knob | Purpose |
| --- | --- |
| `family_strip` | Regex removed from the model ID **before** grouping. Handles providers with mid-string versions: google uses `-\d+(?:\.\d+)?` so `gemini-2.5-flash` and `gemini-3-flash-preview` collapse to the same `gemini-flash` family; moonshot uses `-k\d+(?:\.\d+)?` so `kimi-k2.5-thinking` becomes `kimi-thinking`. |
| `family_regex` | Override the default trailing-match regex (rarely needed). |
| `pinned` | Model IDs always kept, even if older. Useful when you've hand-authored a model upstream hasn't indexed yet (e.g. Kimi K2.6 when Moonshot wasn't a models.dev provider). |
| `excluded` | Model IDs always dropped. Curate these when a new flagship supersedes variant lineages — e.g. Kimi K2 Thinking/Turbo get excluded once K2.6 ships, until K2.6-thinking / K2.6-turbo ship. |

### Hand-authored models

When a provider releases a flagship before models.dev (or some other
upstream registry) adds it, seed the YAML manually and add the ID to
`sync.pinned`. Example — after the Kimi K2.6 launch on 2026-04-21:

```yaml
# providers/moonshot.yaml
models:
  - id: kimi-k2.6
    name: Kimi K2.6
    base_url: https://api.moonshot.ai/v1
    limit: 262144
    reserved_tokens: 32768
    supports_streaming: true
    supports_image: true
    supports_reasoning: true
    release_date: "2026-04-21"
    pricing: { input: 0.60, output: 2.50, cache_write: 0.60, cache_read: 0.15 }

sync:
  family_strip: -k\d+(?:\.\d+)?
  pinned: [kimi-k2.6]              # survives the filter even if upstream drops it
  excluded:                        # K2 variants deprecated by K2.6
    - kimi-k2-thinking
    - kimi-k2-thinking-turbo
    - kimi-k2-turbo
```

Once models.dev catches up, the next sync will merge upstream pricing /
capability data over the hand-authored entry (the merge preserves
hand-edited `headers`, `base_url`, `regions`, and the whole `sync:` block
— see the merge rule above).

## Hand-authored providers (`.sync-exclude`)

Two providers are never touched by the weekly sync:

- `plexon` — the Plexon meta-provider. Its routing (key pool, tier-based
  model selection) lives in Plexon server code; the catalog entry only
  describes the three virtual models it exposes.
- `claude-desktop` — the local Claude Desktop CLI integration. Capabilities
  here track what the Claude Code binary supports, not models.dev data.

Add more names to `.sync-exclude` (one per line) if you introduce further
meta-providers.
