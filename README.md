# _**hewd**_

_A project health, documentation, and structure toolkit for modern software repositories._

hewd is a fast, dependency-free CLI that analyzes your repository’s documentation, configuration files, project structure, and code layout — then produces human‑friendly reports, machine‑readable output, badges, and even automated fixes.

It is designed to help teams maintain consistent repository quality, improve documentation health, track changes over time, and enforce standards in continuous integration environments.

---

## 📑 Table of Contents

- [Features](#-features)
- [Installation](#installation)
- [Quick-start](#-quick-start)
- [Commands](#-commands)
  - [hewd scan](#hewd-scan)
  - [hewd doctor](#hewd-doctor)
  - [hewd fix](#hewd-fix)
  - [hewd badge](#hewd-badge)
  - [hewd export](#hewd-export)
  - [hewd-diff](#hewd-diff)
  - [hewd init](#hewd-init)
- [Machine‑Readable Output](#machine-readable-output)
- [github-action](#-github-action-integration)
- [Configuration](#configuration)
- [Roadmap]
- [License](#license)

---

## ✨ Features

### 🔍 Project Scanning

hewd scans your repository to detect:

- programming languages
- documentation files
- configuration files
- CI workflows
- project metadata
- docs directories
- file and directory counts

### 🩺 Diagnostic Engine (`hewd doctor`)

Runs category‑aware rules to identify:

- missing or incomplete docs
- missing LICENSE/CHANGELOG/CONTRIBUTING
- missing or incomplete CI workflows
- missing docs/ structure
- stale documentation
- multi‑language repos without architecture docs

Includes:

- per‑category scoring
- overall health score
- severity levels (info/warn/error)

### 🧮 Scoring

- **Overall project score (0–100)**
- **Category scores:** documentation, config, structure
- **Machine‑readable scoring API**
- **Configurable severity overrides**

### 🧰 Auto‑Fix (`hewd fix`)

Automatically generates missing assets:

- LICENSE
- CONTRIBUTING.md
- CHANGELOG.md
- `.github/workflows/ci.yml`
- `docs/` directory

Supports:

- dry‑run (default)
- `--apply` (write changes)

### 🏷️ Badge Generator (`hewd badge`)

Generates clean SVG badges for:

- overall score
- category scores (future)

Perfect for READMEs and dashboards.

### 🧾 Machine‑Readable Export (`hewd export`)

Outputs a stable JSON schema containing:

- project scores
- all rule results
- fixable issues
- metadata
- timestamps
- version info

### 🔁 Diff Engine (`hewd diff`)

Compares two hewd JSON exports (e.g., main branch vs PR branch):

- score change
- category score changes
- new issues introduced
- resolved issues
- pretty‑printed diff
- arrows for direction (↑ ↓ ↔)
- section headers and separators

Perfect for CI regression detection.

### 🤖 GitHub Action Integration

- runs doctor on pull requests
- posts Markdown reports
- updates PR comments instead of spamming
- supports badges, diff, and scoring
- uses exit codes for CI enforcement

---

## 📦 Installation

_Installation methods are placeholders until releases are published._

### Go Install (future)

```bash
go install github.com/jbakchr/hewd/cmd/hewd@latest
```

### Manual Build

```bash
git clone https://github.com/jbakchr/hewd
cd hewd
go build -o hewd ./cmd/hewd
```

## 🚀 Quick Start

### Scan the project:

```bash
hewd scan --pretty
```

### Run diagnostics:

```bash
hewd doctor
```

### Generate a Markdown report:

```bash
hewd doctor --md > health.md
```

### Save a machine‑readable report:

```bash
hewd export --output hewd.json
```

### Compare two reports:

```bash
hewd diff old.json new.json
```

---

## 📚 Commands

### `hewd scan`

Scans the directory and outputs:

- languages
- documentation files
- configuration files
- project stats

Supports:

```
--json
--yaml
--pretty
```

---

### `hewd doctor`

Runs all diagnostics and produces:

- issues grouped by category
- scoring
- JSON/YAML/Markdown output
- CI‑friendly failures

Options:

```bash
--json
--yaml
--md
--only
--except
--score
--category-score
--fail-on=info|warn|error
```

---

### `hewd fix`

Dry-run by default:

```bash
hewd fix
```

Apply fixes:

```bash
hewd fix --apply
```

---

### `hewd badge`

Generate a score badge:

```bash
hewd badge --output badge.svg
```

---

### `hewd export`

Output machine‑readable project health:

```bash
hewd export --output hewd.json
```

---

### `hewd diff`

Compare two hewd reports:

```bash
hewd diff old.json new.json
```

Outputs:

- score delta
- category deltas
- new issues
- resolved issues
- structured and readable formatting

Future:

- --json, --yaml, --md diff output
- regression gating flags

---

### `hewd init`

Initializes:

- .hewd/ directory
- default project configuration

---

## 🧾 Machine‑Readable Output

All structured output uses the MachineOutput schema:

- schemaVersion
- hewdVersion
- generatedAt
- score
- categoryScores
- results
- fixable

Supports:

- JSON
- YAML
- Markdown (rendered from JSON)

These outputs power:

- dashboards
- GitHub Action PR comments
- diff comparisons
- trend tracking

---

## 🤖 GitHub Action

The repository includes a custom GitHub Action that:

- Runs hewd on PRs
- Builds the hewd binary
- Posts Markdown reports
- Updates PR comments (no duplication)
- Uploads reports as artifacts
- Supports automatic regression detection (future)

Example:

```bash
- uses: ./.github/actions/hewd-action
  with:
    github-token: ${{ secrets.GITHUB_TOKEN }}
    pr-comment: true
    md-report: true
```

---

## ⚙️ Configuration

`hewd init` generates:

```bash
.hewd/
  config.yaml
```

Supports:

- rule enabling/disabling
- severity weight overrides
- scanner include/exclude paths

Example:

```bash
rules:
  DOC_README_MISSING: true
weights:
  DOC_LICENSE_MISSING: 3
scan:
  include: []
  exclude:
    - node_modules
    - vendor
```

---

## 🗺️ Roadmap

- JSON/YAML/MD output for hewd diff
- Regression gating (--fail-on-score-drop=N)
- Extended auto‑fix (README scaffolding, ADR templates)
- Issue grouping by severity/category
- HTML reports
- Badges by category
- Project health dashboard
- Repo structure smell detection
- Plugin system for custom rules

---

## 📄 License

To be determined.

---

## 🔗 Project Links

- Source Code: https://github.com/jbakchr/hewd
- Issues: https://github.com/jbakchr/hewd/issues
- Releases: https://github.com/jbakchr/hewd/releases
