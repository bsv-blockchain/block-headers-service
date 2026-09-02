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
       <a href="https://deepwiki.com/bsv-blockchain/block-headers-service"><img src="https://deepwiki.com/badge.svg" alt="Ask DeepWiki"></a>
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
       🤖&nbsp;<a href="#-ai-usage--assistant-guidelines"><code>AI&nbsp;Usage</code></a>
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
<summary><strong>Repository Features</strong></summary>
<br/>

This repository includes 25+ built-in features covering CI/CD, security, code quality, developer experience, and community tooling.

**[View the full Repository Features list →](.github/docs/repository-features.md)**

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

The system is configured via modular env files in [`.github/env/`](.github/env/README.md) and provides 17x faster execution than traditional Python-based pre-commit hooks. See the [complete documentation](http://github.com/mrz1836/go-pre-commit) for details.

</details>

<details>
<summary><strong>GitHub Workflows</strong></summary>
<br/>

All workflows are driven by modular configuration in [`.github/env/`](.github/env/README.md) — no YAML editing required.

**[View all workflows and the control center →](.github/docs/workflows.md)**

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

All unit tests and [examples](examples) run via [GitHub Actions](https://github.com/bsv-blockchain/block-headers-service/actions) and use [Go version 1.26.x](https://go.dev/doc/go1.26). View the [configuration file](.github/workflows/fortress.yml).

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

## 🤖 AI Usage & Assistant Guidelines
Read the [AI Usage & Assistant Guidelines](.github/tech-conventions/ai-compliance.md) for details on how AI is used in this project and how to interact with AI assistants.

<br/>

## 👥 Maintainers
| [<img src="https://github.com/icellan.png" height="50" alt="Siggi" />](https://github.com/icellan) | [<img src="https://github.com/galt-tr.png" height="50" alt="Galt" />](https://github.com/galt-tr) | [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) |
|:--------------------------------------------------------------------------------------------------:|:-------------------------------------------------------------------------------------------------:|:------------------------------------------------------------------------------------------------:|
|                                [Siggi](https://github.com/icellan)                                 |                                [Dylan](https://github.com/galt-tr)                                 |                                [MrZ](https://github.com/mrz1836)                                 |

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
