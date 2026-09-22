<div align="center">
  <img src="assets/icon.png" alt="Plexon AI" width="120" />
  <h1>Plexon AI</h1>
  <p><strong>Your assistant. Your apps. Your machine.</strong></p>
  <p>A desktop AI assistant for Windows, macOS, and Linux</p>

  <a href="https://plexon.ai"><img src="https://img.shields.io/badge/Website-plexon.ai-729ff0?style=flat-square" alt="Website" /></a>
  <a href="https://github.com/hellenic-development/plexon.ai/releases"><img src="https://img.shields.io/github/v/release/hellenic-development/plexon.ai?style=flat-square&color=729ff0&label=Download" alt="Download" /></a>
  <a href="https://github.com/hellenic-development/plexon.ai/issues"><img src="https://img.shields.io/github/issues/hellenic-development/plexon.ai?style=flat-square&color=729ff0" alt="Issues" /></a>
  <img src="https://img.shields.io/badge/Platform-Windows%20%7C%20macOS%20%7C%20Linux-729ff0?style=flat-square" alt="Platform" />
  <a href="https://github.com/hellenic-development/plexon.ai/releases/latest"><img src="https://img.shields.io/badge/VirusTotal-scanned%20every%20release-brightgreen?style=flat-square&logo=virustotal" alt="VirusTotal" /></a>
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

Out of the box it ships 12 personas, 82 app connectors, 192 agents, 150+ skills, 17 plugins, and 153 built-in tools.

## One marketplace, five kinds of thing

Skills, agents, plugins, personas, and MCP servers share one marketplace: browse everything, filter by kind, click install. It lists 25,157 items from 8 sources, and every card names its source so you know what you are installing.

A marketplace is a folder in a Git repository. No signup, no proprietary format, no central registry deciding what gets listed. Paste `owner/repo`, point at a `tree/branch/subdir` URL, drop a zip on the Import tab, or double-click a `.plxm` file a colleague emailed you. Plexon previews what is inside before anything reaches your disk, and nothing from a marketplace runs at install time.

Developer Mode builds one from inside the app. An 8-step wizard reads your installed library and fills in each item's name, description, and version for you; linked items are symlinked rather than copied, so edits to `SKILL.md` or a persona YAML go live the moment you save. One Sync button publishes your changes and pulls everyone else's. Conflicts resolve per file: keep mine, take theirs, or let the AI merge. Nobody types a git command.

That is also how a company keeps a team in step. Each department publishes its curated marketplace in a GitHub repo it owns, a new hire pastes one URL, and every install checks that repo on launch and applies what changed. A marketplace the company manages can be developed from inside Plexon too, and "Update from documents" turns a folder of source documents into changes you approve file by file. Adding sources and linking local folders is included from Pro.

Read more: [the marketplace](https://plexon.ai/features/marketplace/) and [Developer Mode](https://plexon.ai/features/marketplace-developer-mode/).

### For teams and companies

[Plexon for Business](https://plexon.ai/business/) white-labels the app with your branding, pre-loads your own agents and personas, and can ship secret items your clients can use but cannot read or copy.

- [Company Admin Panel](https://plexon.ai/features/company-admin/): approve people, put them in groups, see who holds a seat, and choose which marketplaces and buttons each group gets
- [Company Marketplaces](https://plexon.ai/features/company-marketplaces/): your own catalog of skills, agents, personas, and workflows, hosted by Plexon or in your own repository
- [Company Websites](https://plexon.ai/features/company-websites/): one inventory of every site your people published, per-group publish rights, and takedown
- [Persona Agreements](https://plexon.ai/features/persona-agreements/): a role that reaches sensitive data can require a written agreement first, recorded against the exact text shown
- Single sign-on over SAML 2.0 or OpenID Connect

<!-- Partner organisations run their own assistant on Plexon with a public page and, optionally, a sign-in-gated internal one for staff. [PNOĒ](https://plexon.ai/companies/pnoe/) was the first: a doctor co-pilot for breath-biomarker consultations that loads each patient's report before the visit. The rest are at [plexon.ai/companies](https://plexon.ai/companies/). -->

## What you can do

**Everyday**

- **Ask and research.** It searches the web out of the box, with no connector to install and no key to paste.
- **Write and edit.** Letters, posts, reports, CVs. It reads and writes real Word and Excel files without Office installed, shows edits as tracked changes while a document is open in Word, and saves a finished draft as a proper PDF.
- **Read anything.** PDFs, ebooks, slide decks, spreadsheets, and saved emails, with no conversion first.
- **Sort out your files.** Point it at a folder and it renames, sorts, and finds what takes up space. Removed files go to the recycle bin. Ask it to clear the screenshots out of a phone backup and it identifies each file by reading it, so an export that renamed everything to `IMG_9084.PNG` does not defeat it.
- **Make pictures, video, and music.** Describe an image, a clip, or a song. Narrated videos with captions render on your own computer, and Edit in Studio exports MP4 or WebM.
- **Talk instead of typing.** Plexon Voice handles dictation and spoken replies on every plan. Premium adds High fidelity and cloning your own voice from a short recording, and My Voice clones it locally in 23 languages.

**At work**

- **Work like a specialist.** Twelve personas, each swapping in its own server-hosted system prompt.
- **Reach into your apps.** 82 built-in connectors, signed in with your own account.
- **Put it on a schedule.** A Monday-morning report, an inbox summary every evening, a one-time reminder in your time zone. [Dynamic workflows](https://plexon.ai/features/dynamic-workflows/) run it on your computer.
- **Build software.** Five modes, language servers for Go, Python, TypeScript, and Rust, parallel sub-agents in isolated git worktrees, and a review queue holding every diff they produce.

## Twelve personas

Each persona swaps in its own server-hosted system prompt, so the AI genuinely thinks like that role, and pre-activates the agents, skills, plugins, and tools it needs. Build your own with an AI-assisted wizard, or install one from a marketplace.

| Persona | What it brings |
|---------|----------------|
| [Software Developer](https://plexon.ai/personas/software-developer/) | Correctness first, minimal diffs, match-the-file style, git safety by default. Can watch Jira for bugs assigned to you |
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

## 82 app connectors

<div align="center">
  <a href="https://plexon.ai/features/app-connectors/">
    <img src="assets/connectors.png" alt="The 82 app connectors built into Plexon AI: Slack, Discord, Telegram, WhatsApp, Twilio, Instagram, X (Twitter), Unipile, Notion, Google Drive Suite, Google Workspace, Microsoft 365, Todoist, Asana, Trello, Dropbox, WordPress, Linear, Atlassian, Salesforce, HubSpot, GoHighLevel, Twenty, Zendesk, Intercom, Mailchimp, Keragon, FitMetrics, GitHub, Vercel, Netlify, Cloudflare, Docker, Supabase, Sentry, x64dbg, GitLab, Datadog, Grafana, Kubernetes, Firebase, Buildkite, CircleCI, Railway, shadcn/ui, Langfuse, GoDaddy, Brave Search, Perplexity, Exa, n8n, Zapier, Pipedream, Make, Airtable, Database, BigQuery, PostHog, Redis, ClickHouse, Stripe, Shopify Dev, WooCommerce, Product Search, Shopify Shopping, Affiliate Buy and Feed Hub, eBay, Figma, Meta Ads, Spotify, Reddit, Higgsfield, ViewMax, HyperFrames (Hosted), Agent Browser, Skyvern, Playwright, Browser MCP, Chrome MCP, Desktop Automation, Code Graph, and Home Assistant" width="860" />
  </a>
</div>

Branded integrations you turn on once with your own account, so the AI works inside the software you already use: post to Slack, reply to mail, update your store, file issues, check a Grafana dashboard, edit DNS. Sign-in matches each app (a browser sign-in, a device code, a key you paste, a server URL, or a one-time QR scan for personal WhatsApp). Credentials are encrypted and stay on your machine, and activation is per workspace.

- **Messaging:** Slack, Discord, Telegram, WhatsApp, Twilio, Instagram, X (Twitter), Unipile
- **Productivity:** Notion, Google Drive Suite, Google Workspace, Microsoft 365, Todoist, Asana, Trello, Dropbox, WordPress
- **Project tracking:** Linear, Atlassian (Jira and Confluence)
- **CRM and support:** Salesforce, HubSpot, GoHighLevel, Twenty, Zendesk, Intercom, Mailchimp
- **Developer and cloud:** GitHub, Vercel, Netlify, Cloudflare, Docker, Supabase, Sentry, x64dbg, GitLab, Datadog, Grafana, Kubernetes, Firebase, Buildkite, CircleCI, Railway, shadcn/ui, Langfuse, Code Graph
- **Data:** Airtable, Database (PostgreSQL, MySQL, MariaDB, SQL Server, SQLite), BigQuery, PostHog, Redis, ClickHouse
- **Commerce:** Stripe, Shopify Dev, WooCommerce, Product Search, Shopify Shopping, Affiliate Buy and Feed Hub, eBay
- **Search:** Brave Search, Perplexity, Exa
- **Automation:** n8n, Zapier, Pipedream, Make, Desktop Automation
- **Browser:** Agent Browser, Skyvern, Playwright, Browser MCP, Chrome MCP
- **Media:** Spotify, Reddit, Higgsfield, ViewMax, HyperFrames (Hosted)
- **One apiece:** Keragon (healthcare), FitMetrics (fitness and coaching), GoDaddy (domains), Home Assistant (smart home), Figma (design), Meta Ads (advertising)

Anything without a tile still works: [API Connections](https://plexon.ai/features/api-connections/) turn any HTTP API into typed tools by OpenAPI import, Postman import, or by hand. Every connector has its own page with example prompts and how sign-in works, listed at [plexon.ai/features/app-connectors](https://plexon.ai/features/app-connectors/).

## Automation and memory

Plexon keeps working when you are not looking at it, and it remembers what happened.

| | What it does |
|---|---|
| [Dynamic workflows](https://plexon.ai/features/dynamic-workflows/) | A trigger wired to typed steps, or described in chat and drafted for you. Starts on a schedule, a one-time date, or an event in Jira, GitHub, or Linear. A step can call any built-in tool, open a chat, or ping your phone. A local engine snapshots every run and sends at most once |
| [Dashboards](https://plexon.ai/features/dashboards/) | Boards of live widgets fed by a connector, an AI query, a workflow, or static text. Connector and tool refreshes spend zero AI tokens |
| [Projects](https://plexon.ai/features/projects/) and [Checklist](https://plexon.ai/features/checklist/) | Task boards with subtasks, schedules, and owners, plus a checklist beside the chat whose steps can move to a project |
| [Flows](https://plexon.ai/features/flows/) | Repeat an action across many items in one chat card. Pause a job or retry what failed while other work continues |
| [Data Sources](https://plexon.ai/features/data-sources/) | Load a spreadsheet or CSV and use its rows in chat, workflows, and dashboards |
| [Orchestrator](https://plexon.ai/features/orchestrator/) and [Team Coordination](https://plexon.ai/features/teams/) | Split a large task across specialist agents working at the same time, sharing notes across conversations |
| [Knowledge graph](https://plexon.ai/features/knowledge-graph/) | Ask about a person, a client, or a project and get every fact with the note and line it came from. Derived from your markdown, so forgetting it loses nothing |
| [AutoDream](https://plexon.ai/features/autodream/) | Between sessions, removes duplicate and outdated memory notes and turns your feedback into instructions |
| [Self-improving skills](https://plexon.ai/features/self-improving-skills/) | A curator agent reviews each session and saves, refines, or retires skills. Pin an instruction to keep it out of reach. Off by default, because each run is a real AI call that costs tokens |
| [Quality Harness](https://plexon.ai/features/quality-harness/) | The model verifies facts, writes a spec for multi-step work, and reports pass or fail after each change. It checks for eleven completion shortcuts |
| [Hook events](https://plexon.ai/features/richer-hooks/) | Run your own script or agent on seven lifecycle events, from a tool call to a finished reply |

The [Telegram](https://plexon.ai/features/telegram/) and [Slack](https://plexon.ai/features/slack/) bots put the same assistant on your phone, each locked to your own account. Send files both ways, and have routines post their results back. Slack connects over Socket Mode, so nothing is exposed to the internet, and it works through Slack on Apple Watch.

## Around the chat

| | What it does |
|---|---|
| [Drafts](https://plexon.ai/features/drafts/) | Reports, pages, and videos beside the chat, with saved versions to compare and continue from |
| [Command Palette](https://plexon.ai/features/command-palette/) | Ctrl+K finds a chat, setting, tool, or panel |
| [Session Branching](https://plexon.ai/features/session-branching/) | Take a chat a different way from any message and keep the original |
| [Side Questions](https://plexon.ai/features/side-questions/) | Type `@btw` for a quick answer in a separate card while work continues |
| [Prompt Suggestions](https://plexon.ai/features/prompt-suggestions/) | A predicted next message after each reply. Press Tab to take it |
| [Pinned Messages](https://plexon.ai/features/pinned-messages/) and [Cross-Chat Messaging](https://plexon.ai/features/cross-chat-messaging/) | Pin what matters across projects, and let chats on the same project share decisions |
| [Dynamic UI Widgets](https://plexon.ai/features/dynamic-ui-widgets/) | Answers as charts, tables, tickable checklists, forms, and playable video |
| [Session Sharing](https://plexon.ai/features/session-sharing/) | Save a conversation as an encrypted `.plexshare` file someone else can open and continue |
| [Cloud Backup](https://plexon.ai/features/cloud-backup/) | Encrypted backups of your setup and chats to your own Google Drive, restorable on a new install |
| [Authenticator](https://plexon.ai/features/authenticator/) | Your two-factor codes, imported from Google Authenticator, Aegis, 2FAS, Bitwarden and others, or scanned with the webcam. Keys stay encrypted on your machine |
| [Plexonaki](https://plexon.ai/features/plexonaki/) | New in September: a small pet over your desktop that shows what the assistant is doing, answers out loud when you say its name, and takes work you drop on it. Off by default |
| [Customization](https://plexon.ai/features/settings/) | 13 themes, text size, density, and hiding any panel you never open |
| [Migrate from Claude Code](https://plexon.ai/features/claude-code-migration/) | Brings your Claude Code agents, hooks, plans, memory, and conversations across |

## For developers

| | What it does |
|---|---|
| [Code Intelligence](https://plexon.ai/features/code-intelligence/) | Language servers for Go, Python, TypeScript, and Rust. Definitions, references, and fresh diagnostics after each edit |
| [Parallel agents](https://plexon.ai/features/parallel-agents/) | Fan one prompt across agents in isolated git worktrees and merge per tile with fast-forward, squash, or cherry-pick |
| [Diff review queue](https://plexon.ai/features/diff-review-queue/) | Every tile waiting on review, in one panel, across sessions |
| [Graph-Aware Code Review](https://plexon.ai/features/code-review-graph/) | Reviews only the part of the codebase a diff touches, through a persistent Tree-sitter graph. Roughly 8x fewer review tokens |
| [GitHub Auto-Review](https://plexon.ai/features/github-auto-review/) | Watches issues and pull requests, and drafts replies, reviews, and tested patches into Approvals before anything is posted |
| [Embedded Browser](https://plexon.ai/features/embedded-browser/) | A browser docked beside the chat. Point at an element and the AI gets its selector, HTML, styles, and a screenshot. Sites that publish their own actions (WebMCP) can run them, one confirmation per call |
| [Publish Sites](https://plexon.ai/features/publish-sites/) | Preview a static site, share a temporary link, or publish to GitHub Pages. Plexon follows the build and checks the live address before calling it done |
| [Design Extract](https://plexon.ai/features/design-extract/) | A site's design tokens, typography, components, and WCAG scores, as 8 handoff artifacts for Tailwind, shadcn/ui, Figma, React, and CSS |
| [Shell commands](https://plexon.ai/features/shell-commands/) | Start a message with `!` to run it in the session directory. Commands past 90 seconds continue in the background |
| [Developer Tools](https://plexon.ai/features/developer/) | Source control, hooks, and prompt controls in the panel next to the chat |

## Modes and providers

Five modes, each with different permissions:

- **Code** reads, writes, and runs. The default for hands-on work.
- **Plan** explores read-only and writes one implementation plan. Approving it hands off inside the same conversation, with everything it read still in context.
- **Ask** is read-only question answering. No file writes.
- **Auto** runs long tasks without stopping to check in.
- **Spec** drives constitution, specify, plan, tasks, implement through living markdown contracts, compatible with GitHub Spec Kit.

Eleven providers. Sign in to Plexon for a managed one, or bring your own key:

| Provider | How you connect |
|----------|-----------------|
| Plexon | Included in every plan. Text, image, video, voice, transcription, and music, with no keys to manage |
| Anthropic | Your API key |
| OpenAI | Your API key |
| Google Gemini | Your API key |
| Moonshot AI | Your API key |
| Fireworks AI | Your API key, for open-weight models |
| OpenRouter | Your API key, 300+ models from every major lab |
| Alibaba Qwen | Your API key |
| Z.AI | Your API key |
| MiniMax | Your API key, including image, video, speech, and music models |
| Claude Desktop | Your local app, no API key |

The model list lives in [`catalog/`](catalog/) in this repository and the app syncs from it, so new models arrive without an app update. It holds 413 models today. Premium adds [Custom Providers](https://plexon.ai/features/custom-providers/): any OpenAI-compatible or Anthropic-compatible endpoint, including local models through Ollama, LM Studio, or vLLM.

## Security and privacy

- Your files and credentials stay on your machine, and connector secrets are encrypted at rest
- The server is stateless. It assembles the prompt, routes the request, and discards everything: no storage, no logging, no retention
- AES-256-GCM for local session data. Shell commands, browser launches, and connector installs each ask first, and destructive commands like `rm -rf` are blocked unless you allow them for a project
- [Session Lock and App Lock](https://plexon.ai/features/session-lock/) put a passcode on one chat or on the whole app, unlockable with Touch ID or Windows Hello. The app lock can come back after idle time, and you can cap wrong PINs so that reaching the limit wipes Plexon's local data. A forgotten chat passcode cannot be recovered by anyone, including us. Turning a lock on needs Pro; unlocking, changing, and removing one work on any plan
- [Live Share](https://plexon.ai/features/live-share/) connects two apps directly over WebRTC, with no Plexon server in the handshake. Because there is no relay, some network pairs cannot connect: the same network is the reliable case. Premium plan
- [Moving to another machine](https://plexon.ai/features/transfer-bundles/) exports your whole setup as one encrypted file, under a passphrase stored nowhere. No cloud, no account sync
- Every installer is uploaded to VirusTotal when the release is published, and each [release page](https://github.com/hellenic-development/plexon.ai/releases/latest) carries the results per file with a link to the full report. macOS builds are code-signed and the disk images are notarized by Apple

Details: [plexon.ai/security](https://plexon.ai/security/).

## Platform support

| Platform | Versions | Architectures |
|----------|----------|---------------|
| Windows | 10 and later | x64, ARM64 |
| macOS | 12 (Monterey) and later | Intel (x64), Apple Silicon (ARM64) |
| Linux | Ubuntu 22.04+, Fedora 38+ | x64, ARM64 (AppImage, .deb, tar.gz) |

Minimum 4 GB RAM and 500 MB of disk. My Voice needs about 5 GB more, or about 11.5 GB with an NVIDIA card. Auto-update is built in.

## Getting started

1. **Download** the latest release from [Releases](https://github.com/hellenic-development/plexon.ai/releases) or [plexon.ai/download](https://plexon.ai/download/)
2. **Install** and launch
3. **Sign in** with a Plexon plan
4. **Pick a persona**, turn on the connectors you use, and ask

| Plan | Monthly | Every 3 months | What you get |
|------|---------|----------------|--------------|
| Pro | $49 | $132 | Plexon's AI model, 2,000 messages every 5 hours, 10 images a day, Plexon Voice, all personas, agents, skills, and tools, marketplace sources and publishing, Session Lock and App Lock |
| Premium | $129 | $348 | Everything in Pro, on the priority tier: 4,500 messages every 5 hours, 50 images, 10 videos, and 30 music tracks a day, High fidelity voice and voice cloning, shared project boards, publishing websites, company personas, Custom Providers, running a model on your own machine, Claude Code Desktop integration, and Live Share |

Paying every 3 months saves 10%. There is no free tier. Both plans also work with your own account on any of the providers above. Full details at [plexon.ai/pricing](https://plexon.ai/pricing/).

New to it? [plexon.ai/learn](https://plexon.ai/learn/) is an eight-step path from install to publishing your own marketplace, 115 minutes end to end.

## This repository

This repository hosts **issue tracking, public releases, and the public model catalog**. The application source is maintained privately.

[`catalog/`](catalog/) holds the provider and model definitions the app syncs from: 11 providers and 413 models across text, image, video, speech, transcription, and music. Its own README covers the format and how to contribute. A [weekly workflow](.github/workflows/catalog-sync.yml) syncs it against models.dev and opens a pull request, and every change is checked against a vendor denylist before it can merge.

- Report bugs or request features via [Issues](https://github.com/hellenic-development/plexon.ai/issues)
- Security reports go to contact@plexon.ai

## Links

- **Website:** [plexon.ai](https://plexon.ai)
- **Download:** [plexon.ai/download](https://plexon.ai/download/)
- **Features:** [plexon.ai/features](https://plexon.ai/features/)
- **Use cases:** [plexon.ai/use-cases](https://plexon.ai/use-cases/)
- **Pricing:** [plexon.ai/pricing](https://plexon.ai/pricing/)
- **For business:** [plexon.ai/business](https://plexon.ai/business/)
- **Community:** [plexon.ai/community](https://plexon.ai/community/)
- **Changelog:** [plexon.ai/changelog](https://plexon.ai/changelog/)

## License

All rights reserved. Copyright &copy; 2026 [Hellenic Development](https://hellenic.dev).
