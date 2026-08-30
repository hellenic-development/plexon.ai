<div align="center">
  <img src="assets/icon.png" alt="Plexon AI" width="120" />
  <h1>Plexon AI</h1>
  <p><strong>Your assistant. Your apps. Your machine.</strong></p>
  <p>A desktop AI assistant for Windows, macOS, and Linux</p>

  <a href="https://plexon.ai"><img src="https://img.shields.io/badge/Website-plexon.ai-729ff0?style=flat-square" alt="Website" /></a>
  <a href="https://github.com/hellenic-development/plexon.ai/releases"><img src="https://img.shields.io/github/v/release/hellenic-development/plexon.ai?style=flat-square&color=729ff0&label=Download" alt="Download" /></a>
  <a href="https://github.com/hellenic-development/plexon.ai/issues"><img src="https://img.shields.io/github/issues/hellenic-development/plexon.ai?style=flat-square&color=729ff0" alt="Issues" /></a>
  <img src="https://img.shields.io/badge/Platform-Windows%20%7C%20macOS%20%7C%20Linux-729ff0?style=flat-square" alt="Platform" />
</div>

<br />

<div align="center">
  <a href="https://www.youtube.com/watch?v=GM5KiArTZWw">
    <img src="https://img.youtube.com/vi/GM5KiArTZWw/maxresdefault.jpg" alt="Watch the Plexon AI walkthrough on YouTube" width="800" />
  </a>
  <p><em>A walkthrough of the desktop app: chat, personas, connectors, workflows, and dashboards.</em></p>
</div>

## What is Plexon?

Plexon answers questions, writes documents, sorts out your files, and does real work inside apps like Slack, Notion, Gmail, and GitHub. Pick a persona and it thinks like a developer, a designer, a dietitian, or a financial analyst, with that role's agents, skills, and tools already switched on. It runs from your computer. Your files, your API keys, and your chat history stay on it, and the server that routes each request keeps nothing.

Before you install a thing, it ships 12 personas, 64 app connectors, 156 agents, 200+ skills, 14 plugins, and 90+ built-in MCP tools.

## One marketplace, five kinds of thing

Skills, agents, plugins, personas, and MCP servers used to sit behind four different dialogs. Now they share one: browse everything, filter by kind, click install. Over 7,100 items are pre-bundled, and every card names its source so you know what you are installing.

A marketplace is a folder in a Git repository. No signup, no proprietary format, no central registry deciding what gets listed. Paste `owner/repo`, point at a `tree/branch/subdir` URL, drop a zip on the Import tab, or double-click a `.plxm` file a colleague emailed you. Plexon previews what is inside before anything reaches your disk, and nothing from a marketplace runs at install time.

Developer Mode builds one from inside the app. An 8-step wizard reads your installed library and fills in each item's name, description, and version for you; linked items are symlinked rather than copied, so edits to `SKILL.md` or a persona YAML go live the moment you save. One Sync button publishes your changes and pulls everyone else's. Conflicts resolve per file: keep mine, take theirs, or let the AI merge. Nobody types a git command.

That is also how a company keeps a team in step. Each department publishes its curated marketplace in a GitHub repo it owns, a new hire pastes one URL, and every install checks that repo on launch and applies what changed. Adding sources and linking local folders needs the Premium plan.

Read more: [the marketplace](https://plexon.ai/features/marketplace/) and [Developer Mode](https://plexon.ai/features/marketplace-developer-mode/).

### For teams and companies

[Plexon for Business](https://plexon.ai/business/) white-labels the app with your branding, pre-loads your own agents and personas, and can ship secret items your clients can use but cannot read or copy.

<!-- Partner organisations run their own assistant on Plexon with a public page and, optionally, a sign-in-gated internal one for staff. [PNOĒ](https://plexon.ai/companies/pnoe/) was the first: a doctor co-pilot for breath-biomarker consultations that loads each patient's report before the visit. The rest are at [plexon.ai/companies](https://plexon.ai/companies/). -->

## What you can do

**Everyday**

- **Ask and research.** It searches the web out of the box, with no connector to install and no key to paste.
- **Write and edit.** Letters, posts, reports, CVs. It reads and writes real Word and Excel files without Office installed, and prints a finished draft to a proper PDF.
- **Sort out your files.** Point it at a folder and it reads, renames, and summarises what is in there. Ask it to clear the screenshots out of a phone backup and it identifies each file by reading it, so an export that renamed everything to `IMG_9084.PNG` does not defeat it. Camera photos are recognised separately and left alone.
- **Talk instead of typing.** A hands-free voice loop, 400+ voices across 80+ languages, transcribed locally by Whisper.

**At work**

- **Work like a specialist.** Twelve personas, each swapping in its own server-hosted system prompt.
- **Reach into your apps.** 64 built-in connectors, signed in with your own account.
- **Put it on a schedule.** A Monday-morning report, an inbox summary every evening. Any sub-agent runs on a cron expression.
- **Build software.** Six modes, parallel sub-agents in isolated git worktrees, and a review queue holding every diff they produce.

## Twelve personas

Each persona swaps in its own server-hosted system prompt, so the AI genuinely thinks like that role, and pre-activates the agents, skills, plugins, and tools it needs. Build your own with an AI-assisted wizard, or install one from a marketplace.

| Persona | What it brings |
|---------|----------------|
| [Software Developer](https://plexon.ai/personas/software-developer/) | Correctness first, minimal diffs, match-the-file style, git safety by default |
| [Designer](https://plexon.ai/personas/designer/) | Clarity beats cleverness, hierarchy drives attention, accessibility non-negotiable |
| [Marketer](https://plexon.ai/personas/marketer/) | Positioning first, one audience per piece, specificity instead of superlatives |
| [Dietitian](https://plexon.ai/personas/dietitian/) | Evidence over trend, whole-diet patterns, supports practitioners and never diagnoses |
| [Financial Analyst](https://plexon.ai/personas/financial-analyst/) | Data integrity, conservative assumptions, sensitivity over point estimates |
| [Product Manager](https://plexon.ai/personas/product-manager/) | Problem before solution, outcomes over outputs, decisions written down |
| [Writer](https://plexon.ai/personas/writer/) | Voice first, show don't tell, cut ruthlessly, fact-check aggressively |
| [Stock Trader](https://plexon.ai/personas/stock-trader/) | Ticker lookup, live quotes, position sizing from a fixed risk budget. Read-only: it never places an order |
| [Greek Logistics](https://plexon.ai/personas/greek-logistics/) | Income tax and VAT, EFKA and payroll, gov.gr and TAXISnet, live AADE myDATA. Bilingual, current to 2026 |
| [Startup Founder](https://plexon.ai/personas/startup-founder/) | Idea validation, SAFEs and term sheets, cap tables and dilution, runway and unit economics |
| [Personal Shopper](https://plexon.ai/personas/personal-shopper/) | Finds products across stores and compares prices, then hands you a link. It never checks out for you |
| [Mirror](https://plexon.ai/personas/mirror/) | A private journal that stays on this machine. Not a licensed therapist, and it says so |

## 64 app connectors

<div align="center">
  <a href="https://plexon.ai/features/app-connectors/">
    <img src="assets/connectors.png" alt="The 64 app connectors built into Plexon AI: Slack, Discord, Telegram, WhatsApp, Twilio, Instagram, X (Twitter), Unipile, Notion, Google Drive Suite, Google Workspace, Microsoft 365, Todoist, Asana, Trello, Dropbox, WordPress, Linear, Atlassian, Salesforce, HubSpot, Twenty, Zendesk, Intercom, Mailchimp, Keragon, GitHub, Vercel, Netlify, Cloudflare, Docker, Supabase, Sentry, GoDaddy, Brave Search, Perplexity, Exa, n8n, Zapier, Pipedream, Make, Desktop Automation, Home Assistant, Airtable, Database, BigQuery, PostHog, Stripe, Shopify Dev, WooCommerce, Product Search, Shopify Shopping, Affiliate Buy and Feed Hub, eBay, Figma, Meta Ads, Spotify, Reddit, Higgsfield, Agent Browser, Skyvern, Playwright, Browser MCP, and Chrome MCP" width="860" />
  </a>
</div>

Branded integrations you turn on once with your own account, so the AI works inside the software you already use: post to Slack, reply to mail, update your store, file issues, edit DNS. Sign-in matches each app (a browser sign-in, a key you paste, a server URL, or a one-time QR scan for personal WhatsApp). Credentials are encrypted and stay on your machine, and activation is per workspace.

- **Messaging:** Slack, Discord, Telegram, WhatsApp, Twilio, Instagram, X (Twitter), Unipile
- **Productivity:** Notion, Google Drive Suite, Google Workspace, Microsoft 365, Todoist, Asana, Trello, Dropbox, WordPress
- **Project tracking:** Linear, Atlassian (Jira and Confluence)
- **CRM and support:** Salesforce, HubSpot, Twenty, Zendesk, Intercom, Mailchimp
- **Developer and cloud:** GitHub, Vercel, Netlify, Cloudflare, Docker, Supabase, Sentry
- **Data:** Airtable, Database (PostgreSQL, MySQL, MariaDB, SQL Server, SQLite), BigQuery, PostHog
- **Commerce:** Stripe, Shopify Dev, WooCommerce, Product Search, Shopify Shopping, Affiliate Buy and Feed Hub, eBay
- **Search:** Brave Search, Perplexity, Exa
- **Automation:** n8n, Zapier, Pipedream, Make, Desktop Automation
- **Browser:** Agent Browser, Skyvern, Playwright, Browser MCP, Chrome MCP
- **Media:** Spotify, Reddit, Higgsfield
- **One apiece:** Keragon (healthcare), GoDaddy (domains), Figma (design), Meta Ads (advertising), Home Assistant (smart home)

Anything without a tile still works: [API Connections](https://plexon.ai/features/api-connections/) turn any HTTP API into typed tools by OpenAPI import, Postman import, or by hand. Every connector has its own page with example prompts and how sign-in works, listed at [plexon.ai/features/app-connectors](https://plexon.ai/features/app-connectors/).

## Automation and memory

Plexon keeps working when you are not looking at it, and it remembers what happened.

| | What it does |
|---|---|
| [Dashboards](https://plexon.ai/features/dashboards/) | Boards of live widgets fed by a connector, an AI query, a workflow, or static text. Connector and tool refreshes spend zero AI tokens |
| [Dynamic workflows](https://plexon.ai/features/dynamic-workflows/) | A trigger wired to typed steps on a canvas, or described in chat and drafted for you. A local engine snapshots every run and sends at most once |
| [Schedule](https://plexon.ai/features/schedule/) | Any sub-agent on a cron expression or interval, with OS notifications. It disables itself after 3 consecutive failures rather than failing quietly forever |
| [Knowledge graph](https://plexon.ai/features/knowledge-graph/) | Ask about a person, a client, or a project and get every fact with the note and line it came from. Derived from your markdown, so forgetting it loses nothing |
| [Self-improving skills](https://plexon.ai/features/self-improving-skills/) | A curator agent reviews each session and saves, refines, or retires skills. Off by default, because each run is a real AI call that costs tokens |
| [Parallel agents](https://plexon.ai/features/parallel-agents/) | Fan one prompt across agents in isolated git worktrees, up to 3 at once, and merge per tile |
| [Diff review queue](https://plexon.ai/features/diff-review-queue/) | Every tile waiting on review, in one panel, across sessions |

The [Telegram](https://plexon.ai/features/telegram/) and [Slack](https://plexon.ai/features/slack/) bots put the same assistant on your phone, each locked to your own account. Slack connects over Socket Mode, so nothing is exposed to the internet. The full feature list is at [plexon.ai/features](https://plexon.ai/features/).

## Modes and providers

Six modes, each with different permissions:

- **Code** reads, writes, and runs. The default for hands-on work.
- **Plan** explores read-only and writes one implementation plan. Approving it hands off inside the same conversation, with everything it read still in context.
- **Ask** is read-only question answering. No file writes.
- **Auto** runs long tasks without stopping to check in.
- **Orchestrator** splits a task into 2 to 5 subtasks across parallel sub-agents, then synthesises.
- **Spec** drives constitution, specify, plan, tasks, implement through living markdown contracts, compatible with GitHub Spec Kit.

Ten providers. Sign in to Plexon for a managed one, or bring your own key:

| Provider | How you connect |
|----------|-----------------|
| Plexon | One subscription. Text, image, video, voice, and music, with no keys to manage |
| Anthropic | Your API key |
| OpenAI | Your API key |
| Google Gemini | Your API key |
| Moonshot | Your API key |
| Fireworks AI | Your API key, for open-weight models |
| OpenRouter | Your API key, 100+ models from every major lab |
| Qwen | Your API key |
| Z-AI | Your API key |
| Claude Desktop | Your local app, no API key |

The model list lives in [`catalog/`](catalog/) in this repository and the app syncs from it, so new models arrive without an app update. Premium adds [Custom Providers](https://plexon.ai/features/custom-providers/): any OpenAI-compatible or Anthropic-compatible endpoint, including local models through Ollama, LM Studio, or vLLM.

## Security and privacy

- Your files and credentials stay on your machine, and connector secrets are encrypted at rest
- The server is stateless. It assembles the prompt, routes the request, and discards everything: no storage, no logging, no retention. It can also be self-hosted
- AES-256-GCM for local session data, and shell commands, browser launches, and connector installs each ask first
- [Session Lock](https://plexon.ai/features/session-lock/) encrypts one chat under a passcode, unlockable with Touch ID or Windows Hello. A forgotten passcode cannot be recovered by anyone, including us. Turning a lock on needs Premium; unlocking, changing, and removing one never do
- [Live Share](https://plexon.ai/features/live-share/) connects two apps directly over WebRTC, with no Plexon server in the handshake. Because there is no relay, some network pairs cannot connect: the same network is the reliable case
- [Moving to another machine](https://plexon.ai/features/transfer-bundles/) exports your whole setup as one encrypted file, under a passphrase stored nowhere. No cloud, no account sync

Details: [plexon.ai/security](https://plexon.ai/security/).

## Platform support

| Platform | Versions | Architectures |
|----------|----------|---------------|
| Windows | 10 and later | x64, ARM64 |
| macOS | 12 and later | Intel (x64), Apple Silicon (ARM64) |
| Linux | Ubuntu 22.04+, Fedora 38+ | x64, ARM64 |

Minimum 4 GB RAM and 500 MB of disk. Auto-update is built in.

## Getting started

1. **Download** the latest release from [Releases](https://github.com/hellenic-development/plexon.ai/releases) or [plexon.ai/download](https://plexon.ai/download/)
2. **Install** and launch
3. **Sign in** with a Plexon plan, or paste your own provider API key
4. **Pick a persona**, turn on the connectors you use, and ask

Pro is $39/month and Premium is $110/month ([pricing](https://plexon.ai/pricing/)). There is no free tier. Both plans include bring-your-own-key access to every provider above.

New to it? [plexon.ai/learn](https://plexon.ai/learn/) is an eight-step path from install to publishing your own marketplace, about two hours end to end.

## This repository

This repository hosts **issue tracking, public releases, and the public model catalog**. The application source is maintained privately.

[`catalog/`](catalog/) holds the provider and model definitions the app syncs from, with its own README covering the format and how to contribute. A [weekly workflow](.github/workflows/catalog-sync.yml) syncs it against models.dev and opens a pull request.

- Report bugs or request features via [Issues](https://github.com/hellenic-development/plexon.ai/issues)
- Security reports go to contact@hellenic.dev

## Links

- **Website:** [plexon.ai](https://plexon.ai)
- **Download:** [plexon.ai/download](https://plexon.ai/download/)
- **Features:** [plexon.ai/features](https://plexon.ai/features/)
- **Pricing:** [plexon.ai/pricing](https://plexon.ai/pricing/)
- **For business:** [plexon.ai/business](https://plexon.ai/business/)
- **Community:** [plexon.ai/community](https://plexon.ai/community/)
- **Changelog:** [plexon.ai/changelog](https://plexon.ai/changelog/)

## License

All rights reserved. Copyright &copy; 2026 [Hellenic Development](https://hellenic.dev).
