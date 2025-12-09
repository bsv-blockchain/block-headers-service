<div align="center">

# 🛡&nbsp;&nbsp;block-headers-service
> Formerly known as "Pulse"

**Go service for storing, indexing, and serving Bitcoin blockchain headers**

<br/>

<a href="https://github.com/bsv-blockchain/block-headers-service/releases"><img src="https://img.shields.io/github/release-pre/bsv-blockchain/block-headers-service?include_prereleases&style=flat-square&logo=github&color=black" alt="Release"></a>
<a href="https://golang.org/"><img src="https://img.shields.io/github/go-mod/go-version/bsv-blockchain/block-headers-service?style=flat-square&logo=go&color=00ADD8" alt="Go Version"></a>
<a href="https://github.com/bsv-blockchain/block-headers-service/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-OpenBSV-blue?style=flat-square" alt="License"></a>

<br/>

<table align="center" border="0">
  <tr>
    <td align="right">
       <code>CI / CD</code> &nbsp;&nbsp;
    </td>
    <td align="left">
       <a href="https://github.com/bsv-blockchain/block-headers-service/actions"><img src="https://img.shields.io/github/actions/workflow/status/bsv-blockchain/block-headers-service/fortress.yml?branch=main&label=build&logo=github&style=flat-square" alt="Build"></a>
       <a href="https://github.com/bsv-blockchain/block-headers-service/actions"><img src="https://img.shields.io/github/last-commit/bsv-blockchain/block-headers-service?style=flat-square&logo=git&logoColor=white&label=last%20update" alt="Last Commit"></a>
    </td>
    <td align="right">
       &nbsp;&nbsp;&nbsp;&nbsp; <code>Quality</code> &nbsp;&nbsp;
    </td>
    <td align="left">
       <a href="https://goreportcard.com/report/github.com/bsv-blockchain/block-headers-service"><img src="https://goreportcard.com/badge/github.com/bsv-blockchain/block-headers-service?style=flat-square" alt="Go Report"></a>
       <a href="https://codecov.io/gh/bsv-blockchain/block-headers-service"><img src="https://codecov.io/gh/bsv-blockchain/block-headers-service/branch/main/graph/badge.svg?style=flat-square" alt="Coverage"></a>
    </td>
  </tr>

  <tr>
    <td align="right">
       <code>Security</code> &nbsp;&nbsp;
    </td>
    <td align="left">
       <a href="https://scorecard.dev/viewer/?uri=github.com/bsv-blockchain/block-headers-service"><img src="https://api.scorecard.dev/projects/github.com/bsv-blockchain/block-headers-service/badge?style=flat-square" alt="Scorecard"></a>
       <a href=".github/SECURITY.md"><img src="https://img.shields.io/badge/policy-active-success?style=flat-square&logo=security&logoColor=white" alt="Security"></a>
    </td>
    <td align="right">
       &nbsp;&nbsp;&nbsp;&nbsp; <code>Community</code> &nbsp;&nbsp;
    </td>
    <td align="left">
       <a href="https://github.com/bsv-blockchain/block-headers-service/graphs/contributors"><img src="https://img.shields.io/github/contributors/bsv-blockchain/block-headers-service?style=flat-square&color=orange" alt="Contributors"></a>
       <a href="https://github.com/sponsors/bsv-blockchain"><img src="https://img.shields.io/badge/sponsor-BSV-181717.svg?logo=github&style=flat-square" alt="Sponsor"></a>
    </td>
  </tr>
</table>

</div>

<br/>
<br/>

<div align="center">

### <code>Project Navigation</code>

</div>

<table align="center">
  <tr>
    <td align="center" width="33%">
       📦&nbsp;<a href="#-about-the-project"><code>About&nbsp;the&nbsp;Project</code></a>
    </td>
    <td align="center" width="33%">
       📦&nbsp;<a href="#-installation"><code>Installation</code></a>
    </td>
    <td align="center" width="33%">
       📚&nbsp;<a href="#-documentation"><code>Documentation</code></a>
    </td>
  </tr>
  <tr>
    <td align="center">
       🤝&nbsp;<a href="#-contributing"><code>Contributing</code></a>
    </td>
    <td align="center">
       🛠️&nbsp;<a href="#-code-standards"><code>Code&nbsp;Standards</code></a>
    </td>
    <td align="center">
       ⚡&nbsp;<a href="#-benchmarks"><code>Benchmarks</code></a>
    </td>
  </tr>
  <tr>
    <td align="center">
       🤖&nbsp;<a href="#-ai-compliance"><code>AI&nbsp;Compliance</code></a>
    </td>
    <td align="center">
       📝&nbsp;<a href="#-license"><code>License</code></a>
    </td>
    <td align="center">
       👥&nbsp;<a href="#-maintainers"><code>Maintainers</code></a>
    </td>
  </tr>
</table>
<br/>

## 📘 About The Project
The Block Headers Service is a Go application that connects to the BSV P2P network to collect and serve information about blockchain headers—both historical and newly mined. It can run as a standalone application or as a module within a larger system.

The primary function is synchronizing with network peers to collect all block headers. On startup, the server initializes its components and connects to the BSV P2P network. Synchronization uses predefined checkpoints (specific known headers) to request headers in batches of 2,000 from peers. Each received header is stored in memory. Once fully synchronized, the server switches to listening mode, where it receives notifications from peers whenever a new block is mined.

For in-depth information and guidance, please refer to the [SPV Wallet Documentation](https://bsvblockchain.gitbook.io/docs).

<br/>

## 📦 Installation

**block-headers-service** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```shell script
git clone https://github.com/bsv-blockchain/block-headers-service
go run ./cmd/main.go
```

<br/>

## 📚 Documentation

<details>
<summary><strong><code>API Reference</code></strong></summary>
<br/>

- **Godoc** – Dive into the API documentation at [pkg.go.dev/github.com/bsv-blockchain/block-headers-service](https://pkg.go.dev/github.com/bsv-blockchain/block-headers-service)
- **Swagger** – Interactive API docs available at `http://localhost:8080/swagger/index.html` when running

</details>

<details>
<summary><strong><code>Quick Start</code></strong></summary>
<br/>

**Docker** (recommended):

Pull image from [Docker Hub](https://hub.docker.com/r/bsvb/block-headers-service):
```bash
docker pull bsvb/block-headers-service
```

Start a new instance:
```bash
docker run bsvb/block-headers-service:latest
```

**From Source**:

1. Install Go according to the [installation instructions](http://golang.org/doc/install)
2. Clone the repo and run:
```bash
git clone https://github.com/bsv-blockchain/block-headers-service
go run ./cmd/main.go
```

**Docker Compose**:
```bash
docker compose up --build
```

**Package Install**:
```bash
go get -u https://pkg.go.dev/github.com/bsv-blockchain/block-headers-service
go build -o block-headers-service
./block-headers-service
```

</details>

<details>
<summary><strong><code>Authentication</code></strong></summary>
<br/>

#### Enabled by Default

The default assumes you want to use Authentication. This requires a single environment variable:

```bash
BHS_HTTP_AUTH_TOKEN=replace_me_with_token_you_want_to_use_as_admin_token
```

#### Disabling Auth Requirement

To disable authentication exposing all endpoints openly, set:

```bash
BHS_HTTP_USE_AUTH=false
```

> **Warning:** We do not recommend exposing the server to the internet without authentication, as it would then be possible for anyone to prune your headers at will.

#### Authenticate with Admin Token

Add the following header to all requests:
```
Authorization: Bearer replace_me_with_token_you_want_to_use_as_admin_token
```

#### Additional Tokens

Create additional tokens:
```http
POST https://{{block-headers-service_url}}/api/v1/access
Authorization: Bearer replace_me_with_token_you_want_to_use_as_admin_token
```

Response:
```json
{
  "token": "some_token_created_by_server",
  "createdAt": "2023-05-11T10:20:16.227582Z",
  "isAdmin": false
}
```

Use the token in requests:
```
Authorization: Bearer some_token_created_by_server
```

Revoke a token:
```http
DELETE https://{{block-headers-service_url}}/api/v1/access/{{some_token_created_by_server}}
Authorization: Bearer replace_me_with_token_you_want_to_use_as_admin_token
```

</details>

<details>
<summary><strong><code>Configuration</code></strong></summary>
<br/>

> Every variable which is used and can be configured is described in [config.example.yaml](config.example.yaml)

#### Defaults

If you run block headers service without editing anything, it will use the default configuration from [defaults.go](/config/defaults.go). It is set up to use _sqlite_ database with enabled authorization (with default auth key) for the HTTP server.

#### Config Variables

Default config variables can be overridden by (in order of importance):

1. Flags (only the ones below)
2. ENV variables
3. Config file

#### Flags

```bash
  -C, --config_file string   custom config file path
  -h, --help                 show help
  -v, --version              show version
  -d, --dump_config          dump config to file, specified by config_file (-C) flag
  -e, --export_headers       export headers to file
```

Generate config file with defaults:
```bash
go run ./cmd/main.go -d
```

Use a custom config file:
```bash
go run ./cmd/main.go -C /my/config.yaml
```

#### Environment Variables

To override any config variable with ENV, use the `BHS_` prefix with the path using `_` as delimiter in uppercase.

Example from `config.example.yaml`:
```yaml
websocket:
  history_max: 300
  history_ttl: 10
```

Override with:
```bash
BHS_WEBSOCKET_HISTORY_MAX=300
```

Connect to Testnet:
```bash
BHS_P2P_CHAIN_NET_TYPE=testnet
```

</details>

<details>
<summary><strong><code>WebSocket Integration</code></strong></summary>
<br/>

Block headers service can notify clients via WebSockets when new headers are received.

#### Subscribing

Block headers service uses [centrifugal/centrifuge](https://github.com/centrifugal/centrifuge) to run a server. To integrate, choose a client library matching your programming language.

**Go Example:** See [./examples/ws-subscribe-to-new-headers/main.go](./examples/ws-subscribe-to-new-headers/main.go) for a complete example using [centrifugal/centrifuge-go](https://github.com/centrifugal/centrifuge-go).

</details>

<details>
<summary><strong><code>Webhooks</code></strong></summary>
<br/>

#### Creating a Webhook

```http
POST https://{{block-headers-service_url}}/api/v1/webhook
```

Request body:
```json
{
  "url": "<server_url>",
  "requiredAuth": {
    "type": "BEARER|CUSTOM_HEADER",
    "token": "<authorization_token>",
    "header": "<custom_header_name>"
  }
}
```

**Auth Types:**
- `BEARER` – Token placed in `Authorization: Bearer {{token}}` header
- `CUSTOM_HEADER` – Header built as `{{header}}: {{token}}`

**Notes:**
- URL must include `http://` or `https://` protocol
- If authorization is enabled, this request requires an `Authorization` header

Response:
```json
{
  "url": "https://example.com/api/v1/webhook/new-header",
  "createdAt": "2023-05-11T13:05:23.297808+02:00",
  "lastEmitStatus": "",
  "lastEmitTimestamp": "0001-01-01T00:00:00Z",
  "errorsCount": 0,
  "active": true
}
```

#### Check Webhook

```http
GET https://{{block-headers-service_url}}/api/v1/webhook?url={{webhook_url}}
```

#### Revoke Webhook

```http
DELETE https://{{block-headers-service_url}}/api/v1/webhook?url={{webhook_url}}
```

#### Refresh Webhook

If the number of failed requests exceeds `WEBHOOK_MAXTRIES`, the webhook will be set to inactive. Use the create endpoint again to refresh it.

</details>

<details>
<summary><strong><code>Database Management</code></strong></summary>
<br/>

#### Updating Predefined Database

When synchronization takes too long, it's recommended to export a fresh database with all headers using the `-e` flag:

```bash
go run ./cmd/main.go -e
```

> **Note:** Export feature works only with SQLite database.

This creates a new `.csv` file with all headers in the same directory as the database file. Commit your changes and create a pull request with the new database file.

</details>

<br/>

<details>
<summary><strong><code>Development Build Commands</code></strong></summary>
<br/>

Get the [MAGE-X](https://github.com/mrz1836/mage-x) build tool for development:
```shell script
go install github.com/mrz1836/mage-x/cmd/magex@latest
```

View all build commands

```bash script
magex help
```

</details>

<details>
<summary><strong><code>Repository Features</code></strong></summary>
<br/>

* **Continuous Integration on Autopilot** with [GitHub Actions](https://github.com/features/actions) – every push is built, tested, and reported in minutes.
* **Pull‑Request Flow That Merges Itself** thanks to [auto‑merge](.github/workflows/auto-merge-on-approval.yml) and hands‑free [Dependabot auto‑merge](.github/workflows/dependabot-auto-merge.yml).
* **One‑Command Builds** powered by battle‑tested [MAGE-X](https://github.com/mrz1836/mage-x) targets for linting, testing, releases, and more.
* **First‑Class Dependency Management** using native [Go Modules](https://github.com/golang/go/wiki/Modules).
* **Uniform Code Style** via [gofumpt](https://github.com/mvdan/gofumpt) plus zero‑noise linting with [golangci‑lint](https://github.com/golangci/golangci-lint).
* **Confidence‑Boosting Tests** with [testify](https://github.com/stretchr/testify), the Go [race detector](https://blog.golang.org/race-detector), crystal‑clear [HTML coverage](https://blog.golang.org/cover) snapshots, and automatic uploads to [Codecov](https://codecov.io/).
* **Hands‑Free Releases** delivered by [GoReleaser](https://github.com/goreleaser/goreleaser) whenever you create a [new Tag](https://git-scm.com/book/en/v2/Git-Basics-Tagging).
* **Relentless Dependency & Vulnerability Scans** via [Dependabot](https://dependabot.com), [Nancy](https://github.com/sonatype-nexus-community/nancy) and [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck).
* **Security Posture by Default** with [CodeQL](https://docs.github.com/en/github/finding-security-vulnerabilities-and-errors-in-your-code/about-code-scanning), [OpenSSF Scorecard](https://openssf.org) and secret‑leak detection via [gitleaks](https://github.com/gitleaks/gitleaks).
* **Automatic Syndication** to [pkg.go.dev](https://pkg.go.dev/) on every release for instant godoc visibility.
* **Polished Community Experience** using rich templates for [Issues & PRs](https://docs.github.com/en/communities/using-templates-to-encourage-useful-issues-and-pull-requests/configuring-issue-templates-for-your-repository).
* **All the Right Meta Files** (`LICENSE`, `CONTRIBUTING.md`, `CODE_OF_CONDUCT.md`, `SUPPORT.md`, `SECURITY.md`) pre‑filled and ready.
* **Code Ownership** clarified through a [CODEOWNERS](.github/CODEOWNERS) file, keeping reviews fast and focused.
* **Zero‑Noise Dev Environments** with tuned editor settings (`.editorconfig`) plus curated *ignore* files for [VS Code](.editorconfig), [Docker](.dockerignore), and [Git](.gitignore).
* **Label Sync Magic**: your repo labels stay in lock‑step with [.github/labels.yml](.github/labels.yml).
* **Friendly First PR Workflow** – newcomers get a warm welcome thanks to a dedicated [workflow](.github/workflows/pull-request-management.yml).
* **Standards‑Compliant Docs** adhering to the [standard‑readme](https://github.com/RichardLitt/standard-readme/blob/master/spec.md) spec.
* **Instant Cloud Workspaces** via [Gitpod](https://gitpod.io/) – spin up a fully configured dev environment with automatic linting and tests.
* **Out‑of‑the‑Box VS Code Happiness** with a preconfigured [Go](https://code.visualstudio.com/docs/languages/go) workspace and [`.vscode`](.vscode) folder with all the right settings.
* **Optional Release Broadcasts** to your community via [Slack](https://slack.com), [Discord](https://discord.com), or [Twitter](https://twitter.com) – plug in your webhook.
* **AI Compliance Playbook** – machine‑readable guidelines ([AGENTS.md](.github/AGENTS.md), [CLAUDE.md](.github/CLAUDE.md), [.cursorrules](.cursorrules), [sweep.yaml](.github/sweep.yaml)) keep ChatGPT, Claude, Cursor & Sweep aligned with your repo's rules.
* **Go-Pre-commit System** - [High-performance Go-native pre-commit hooks](https://github.com/mrz1836/go-pre-commit) with 17x faster execution—run the same formatting, linting, and tests before every commit, just like CI.
* **Zero Python Dependencies** - Pure Go implementation with environment-based configuration via [.env.base](.github/.env.base).
* **DevContainers for Instant Onboarding** – Launch a ready-to-code environment in seconds with [VS Code DevContainers](https://containers.dev/) and the included [.devcontainer.json](.devcontainer.json) config.

</details>

<details>
<summary><strong><code>Library Deployment</code></strong></summary>
<br/>

This project uses [goreleaser](https://github.com/goreleaser/goreleaser) for streamlined binary and library deployment to GitHub. To get started, install it via:

```bash
brew install goreleaser
```

The release process is defined in the [.goreleaser.yml](.goreleaser.yml) configuration file.


Then create and push a new Git tag using:

```bash
magex version:bump push=true bump=patch branch=main
```

This process ensures consistent, repeatable releases with properly versioned artifacts and citation metadata.

</details>

<details>
<summary><strong><code>Pre-commit Hooks</code></strong></summary>
<br/>

Set up the Go-Pre-commit System to run the same formatting, linting, and tests defined in [AGENTS.md](.github/AGENTS.md) before every commit:

```bash
go install github.com/mrz1836/go-pre-commit/cmd/go-pre-commit@latest
go-pre-commit install
```

The system is configured via [.env.base](.github/.env.base) and can be customized using also using [.env.custom](.github/.env.custom) and provides 17x faster execution than traditional Python-based pre-commit hooks. See the [complete documentation](http://github.com/mrz1836/go-pre-commit) for details.

</details>

<details>
<summary><strong><code>GitHub Workflows</code></strong></summary>
<br/>

### 🎛️ The Workflow Control Center

All GitHub Actions workflows in this repository are powered by a single configuration files – your one-stop shop for tweaking CI/CD behavior without touching a single YAML file! 🎯

**Configuration Files:**
- **[.env.base](.github/.env.base)** – Default configuration that works for most Go projects
- **[.env.custom](.github/.env.custom)** – Optional project-specific overrides

This magical file controls everything from:
- **⚙️ Go version matrix** (test on multiple versions or just one)
- **🏃 Runner selection** (Ubuntu or macOS, your wallet decides)
- **🔬 Feature toggles** (coverage, fuzzing, linting, race detection, benchmarks)
- **🛡️ Security tool versions** (gitleaks, nancy, govulncheck)
- **🤖 Auto-merge behaviors** (how aggressive should the bots be?)
- **🏷️ PR management rules** (size labels, auto-assignment, welcome messages)

<br/>

| Workflow Name                                                                      | Description                                                                                                            |
|------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------|
| [auto-merge-on-approval.yml](.github/workflows/auto-merge-on-approval.yml)         | Automatically merges PRs after approval and all required checks, following strict rules.                               |
| [codeql-analysis.yml](.github/workflows/codeql-analysis.yml)                       | Analyzes code for security vulnerabilities using [GitHub CodeQL](https://codeql.github.com/).                          |
| [dependabot-auto-merge.yml](.github/workflows/dependabot-auto-merge.yml)           | Automatically merges [Dependabot](https://github.com/dependabot) PRs that meet all requirements.                       |
| [fortress.yml](.github/workflows/fortress.yml)                                     | Runs the GoFortress security and testing workflow, including linting, testing, releasing, and vulnerability checks.    |
| [pull-request-management.yml](.github/workflows/pull-request-management.yml)       | Labels PRs by branch prefix, assigns a default user if none is assigned, and welcomes new contributors with a comment. |
| [scorecard.yml](.github/workflows/scorecard.yml)                                   | Runs [OpenSSF](https://openssf.org/) Scorecard to assess supply chain security.                                        |
| [stale.yml](.github/workflows/stale-check.yml)                                     | Warns about (and optionally closes) inactive issues and PRs on a schedule or manual trigger.                           |
| [sync-labels.yml](.github/workflows/sync-labels.yml)                               | Keeps GitHub labels in sync with the declarative manifest at [`.github/labels.yml`](./.github/labels.yml).             |

</details>

<details>
<summary><strong><code>Updating Dependencies</code></strong></summary>
<br/>

To update all dependencies (Go modules, linters, and related tools), run:

```bash
magex deps:update
```

This command ensures all dependencies are brought up to date in a single step, including Go modules and any tools managed by [MAGE-X](https://github.com/mrz1836/mage-x). It is the recommended way to keep your development environment and CI in sync with the latest versions.

</details>

<br/>

## 🧪 Examples & Tests

All unit tests and [examples](examples) run via [GitHub Actions](https://github.com/bsv-blockchain/block-headers-service/actions) and use [Go version 1.24.x](https://go.dev/doc/go1.24). View the [configuration file](.github/workflows/fortress.yml).

Run all tests (fast):

```bash script
magex test
```

Run all tests with race detector (slower):
```bash script
magex test:race
```

<br/>

## ⚡ Benchmarks

Run the Go benchmarks:

```bash script
magex bench
```

<br/>

## 🛠️ Code Standards
Read more about this Go project's [code standards](.github/CODE_STANDARDS.md).

<br/>

## 🤖 AI Compliance
This project documents expectations for AI assistants using a few dedicated files:

- [AGENTS.md](.github/AGENTS.md) — canonical rules for coding style, workflows, and pull requests used by [Codex](https://chatgpt.com/codex).
- [CLAUDE.md](.github/CLAUDE.md) — quick checklist for the [Claude](https://www.anthropic.com/product) agent.
- [.cursorrules](.cursorrules) — machine-readable subset of the policies for [Cursor](https://www.cursor.so/) and similar tools.
- [sweep.yaml](.github/sweep.yaml) — rules for [Sweep](https://github.com/sweepai/sweep), a tool for code review and pull request management.

Edit `AGENTS.md` first when adjusting these policies, and keep the other files in sync within the same pull request.

<br/>

## 👥 Maintainers
| [<img src="https://github.com/mrz1836.png" height="50" width="50" alt="MrZ" />](https://github.com/mrz1836) | [<img src="https://github.com/icellan.png" height="50" alt="Siggi" />](https://github.com/icellan) |
|:-----------------------------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------------------:|
|                                      [MrZ](https://github.com/mrz1836)                                      |                                [Siggi](https://github.com/icellan)                                 |

<br/>

## 🤝 Contributing
View the [contributing guidelines](.github/CONTRIBUTING.md) and please follow the [code of conduct](.github/CODE_OF_CONDUCT.md).

### How can I help?
All kinds of contributions are welcome :raised_hands:!
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:.

[![Stars](https://img.shields.io/github/stars/bsv-blockchain/block-headers-service?label=Please%20like%20us&style=social&v=1)](https://github.com/bsv-blockchain/block-headers-service/stargazers)

<br/>

## 📝 License

[![License](https://img.shields.io/badge/license-OpenBSV-blue?style=flat&logo=springsecurity&logoColor=white)](LICENSE)
