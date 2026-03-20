# 📘 _**hewd**_ — Project Health, Documentation & Structure Toolkit

_`hewd` is a fast, dependency‑free CLI for analyzing, comparing, and improving the overall health of software repositories._

It helps teams enforce documentation standards, maintain consistent structure, detect regressions, and automatically fix common issues — locally or in CI.

---

## ✨ Features

- 🔍 **Scan** repositories for documentation, configuration, structure, and metadata
- 🩺 **Run diagnostics** via `hewd doctor`
- 🧮 **Compute scores** for documentation, config, and structure
- 🔁 **Compare reports** with `hewd diff` (new issues, resolved issues, score deltas)
- 🚨 **Regression gating** for CI pipelines
- 🧾 **Machine‑readable JSON/YAML output**
- 📝 **Markdown reports** for PR comments
- 🧰 **Auto‑fix** missing docs and CI files
- 🏷️ **Badge generator** (SVG score badge)
- 🤖 **GitHub Action** with PR comment updating

Full documentation is available in the docs/ directory.

---

## 🚀 Quick Start

### Install (build from source)

```bash
git clone https://github.com/jbakchr/hewd
cd hewd
go build -o hewd ./cmd/hewd
sudo mv hewd /usr/local/bin/hewd
```

### Scan your project

```bash
hewd scan --pretty
```

### Run diagnostics

```bash
hewd doctor
```

Generate a Markdown report:

```bash
hewd doctor --md > health.md
```

### Export machine‑readable output

```bash
hewd export --output hewd.json
```

### Compare two reports (diff engine)

```bash
hewd diff old.json new.json
```

Pretty, JSON, YAML, and Markdown output modes available.

---

## 📊 Demo Output — What `hewd` Looks Like in Practice

### 🔍 Example: Running `hewd doctor`

```bash
===== OVERALL SCORE =====
  82 / 100

===== CATEGORY SCORES =====
  Documentation:   70
  Config:          90
  Structure:       85

===== DOCUMENTATION ISSUES =====
  - DOC_LICENSE_MISSING (warn): LICENSE file is missing.
  - DOC_CONTRIBUTING_MISSING (info): CONTRIBUTING.md not found.

===== CONFIG ISSUES =====
  (none)

===== STRUCTURE ISSUES =====
  - STR_DOCS_DIR_MISSING (warn): docs/ directory not found.
```

---

### 🔁 Example: Running hewd diff old.json new.json

```bash
===== OVERALL SCORE =====
Old: 78
New: 87
Change: +9 ↑

===== CATEGORY SCORES =====
Documentation:   65 → 75   (+10)
Config:          80 → 84   (+4)
Structure:       90 → 92   (+2)

===== NEW ISSUES =====
(none)

===== RESOLVED ISSUES =====
documentation
  - DOC_LICENSE_MISSING (warn)
  - DOC_README_STALE (info)
```

---

### 🤖 Example: GitHub PR Comment (Markdown)

```bash
# 📊 Hewd Diff Report

## 📈 Score Summary

| Metric         | Old  | New  | Δ     | Trend |
|----------------|------|------|-------|-------|
| Overall Score  |   78 |   87 |   +9  | 🟩⬆️   |
| Documentation  |   65 |   75 |  +10  | 🟩⬆️   |
| Config         |   80 |   84 |   +4  | 🟩⬆️   |
| Structure      |   90 |   92 |   +2  | 🟩⬆️   |

---

## 🆕 New Issues
_No new issues! 🎉_

## ✅ Resolved Issues
### documentation
- **DOC_LICENSE_MISSING** (warn)
- **DOC_README_STALE** (info)
```

---

## 🔁 Example: Markdown Diff Output

```bash
hewd diff old.json new.json --md > diff.md
```

This produces a GitHub‑friendly report with:

- Score changes
- Category score deltas
- New issues
- Resolved issues
- Grouped/sorted sections
- Emojis and trend indicators

Ideal for PR comments.

Learn more → docs/commands/diff.md

---

## 🤖 GitHub Action

`hewd` includes a first‑class GitHub Action that can:

- Run `hewd doctor` or `hewd diff`
- Post or update PR comments
- Run regression gating (`--fail-on-any-regression`)
- Export Markdown + JSON diff output

Minimal GitHub Action usage:

```bash
jobs:
  hewd:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: ./.github/actions/hewd-action
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          pr-comment: true
          md-report: true
```

Full documentation → docs/github-action.md

---

## 🧰 Auto‑Fix Mode

```bash
hewd fix
hewd fix --apply
```

Can generate:

- LICENSE
- CONTRIBUTING.md
- CHANGELOG.md
- .github/workflows/ci.yml
- docs/ directory

More fixers will arrive in future versions.

Learn more → docs/commands/fix.md

---

## 📚 Documentation

All extended/technical documentation is located in the **/docs** directory:

- 📘 Getting Started — docs/getting-started.md
- 🩺 Doctor Command — docs/commands/doctor.md
- 🔁 Diff Command — docs/commands/diff.md (to be added)
- 🔧 Fix Command — docs/commands/fix.md (to be added)
- 🧾 Machine Output — docs/machine-output.md (to be added)
- ⚙️ Configuration — docs/configuration.md (to be added)
- 🤖 GitHub Action — docs/github-action.md (to be added)
- 🗺️ Roadmap — docs/roadmap.md

This keeps the README clean while giving power users all the detail they need.

---

## 🏗 Roadmap

See the full roadmap here:
👉 docs/roadmap.md

`hewd` is currently in the **pre‑release polishing phase, preparing for v0.1.0**.

---

## 📄 License

To be determined.

---

🔗 Links

- Source: https://github.com/jbakchr/hewd
- Issues: https://github.com/jbakchr/hewd/issues
- Releases: https://github.com/jbakchr/hewd/releases

```bash
placeholder
```

```bash
placeholder
```

```bash
placeholder
```

## OLD STUFF - DELETE

## 📑 Table of Contents

- #-features
- #-installation
- #-quick-start
- #-commands
- #-machine-readable-output
- #-diff-engine-hewd-diff
- #-github-action-integration
- #-configuration
- #-roadmap
- #-license
- #-links

---

## ✨ Features

### 🔍 Project Scanning (`hewd scan`)

`hewd scans` your repository to detect:

- programming languages
- documentation files (README, LICENSE, CONTRIBUTING, CHANGELOG…)
- configuration files (CI workflows, package managers, Dockerfiles…)
- docs/ directory
- repo statistics

---

### 🩺 Diagnostic Engine (hewd doctor)

Runs a full ruleset and reports:

- missing/incomplete docs
- missing LICENSE/CHANGELOG/CONTRIBUTING
- missing CI workflows
- stale documentation
- structural problems
- category‑aware severity (info / warn / error)

Includes:

- overall score (0–100)
- documentation/config/structure category scores
- Markdown/JSON/YAML output
- CI‑friendly exit codes

---

### 🧮 Scoring

- Overall project score
- Documentation score
- Config score
- Structure score

---

### 🧰 Auto‑Fix (`hewd fix`)

Automatically generates missing assets such as:

- LICENSE
- CONTRIBUTING.md
- CHANGELOG.md
- .github/workflows/ci.yml
- docs/ directory

Supports dry‑run + --apply.

---

### 🏷️ Badge Generator (`hewd badge`)

Generates SVG badges (hewd badge --output badge.svg) showing project score.

---

### 🧾 Machine‑Readable Export

Outputs a versioned JSON schema containing:

- scores
- rule results
- fixable issues
- metadata
- timestamps
- version info

Perfect for CI, dashboards, or diff comparisons.

---

### 🔁 Diff Engine (hewd diff)

Compares two hewd JSON reports and shows:

- score delta and trend
- category deltas
- new issues
- resolved issues
- pretty terminal output
- Markdown output
- JSON/YAML output
- CI regression gating
- PR‑optimized formatting

### 🤖 GitHub Action Integration

- runs doctor on pull requests
- posts Markdown reports
- updates PR comments instead of spamming
- supports badges, diff, and scoring
- uses exit codes for CI enforcement

---

## 📦 Installation

### Manual Build

```bash
git clone https://github.com/jbakchr/hewd
cd hewd
go build -o hewd ./cmd/hewd
```

(Go install coming soon.)

---

## 🚀 Quick Start

### Scan your repo:

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

### Export machine-readable output:

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

Detects languages, docs, config files, structure, and statistics.

Supports:

```
--json
--yaml
--pretty
```

---

### `hewd doctor`

Produces a full diagnostic health report.

Options:

```
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

### 🔁 Diff Engine (`hewd diff`)

The diff engine **compares two exported hewd JSON reports**:

```bash
hewd diff old.json new.json
```

#### Features

- Overall score delta (with arrows/emojis)
- Documentation/config/structure deltas
- Sorted + grouped new/resolved issues
- Pretty terminal output
- Machine‑readable diff
  ```bash
  hewd diff old.json new.json --json
  hewd diff old.json new.json --yaml
  ```
- Markdown diff for PR comments:
  ```bash
  hewd diff old.json new.json --md
  ```
- Regression gating:
  ```bash
  --fail-on-score-drop=N
  --fail-on-new-errors
  --fail-on-any-regression
  ```

These allow CI to fail automatically if quality regresses.

---

## 🤖 GitHub Action Integration

hewd provides a full GitHub Action to automate repository health checks and diff comparisons inside PRs.

### ✔ Features

- Auto-run hewd doctor or hewd diff
- GitHub‑flavored Markdown comments
- Updates a single PR comment (no spam)
- Regression gating
- JSON/Markdown diff artifacts
- Works on forks
- Composite action (no Node or Docker required)

#### 📦 Example: Doctor Mode

```
jobs:
  hewd:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Run hewd doctor
        uses: ./.github/actions/hewd-action
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          pr-comment: true
          md-report: true
```

#### 🔁 Example: Diff Mode

```
jobs:
  hewd-diff:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Export base report
        run: hewd export --output old.json

      - name: Export PR report
        run: hewd export --output new.json

      - name: Run hewd diff
        uses: ./.github/actions/hewd-action
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          diff: true
          diff-old: old.json
          diff-new: new.json
          diff-pr-comment: true
```

The Action automatically posts or updates a PR comment beginning with:

```
📊 Hewd Diff Report
```

---

## 🧾 Machine-Readable Output

MachineOutput JSON/YAML includes:

- scores
- category scores
- rule results
- fixable items
- metadata
- timestamps
- schema version

This schema is used for:

- CI pipelines
- dashboards
- diff comparisons
- PR comments

---

## ⚙️ Configuration

hewd init creates:

```
.hewd/
  config.yaml
```

Supports:

- rule enabling/disabling
- severity overrides
- include/exclude paths

Example:

```
rules:
  DOC_README_MISSING: true

weights:
  DOC_LICENSE_MISSING: 3

scan:
  exclude:
    - node_modules
    - vendor
```

---

## 🗺️ Roadmap

- Extended auto-fix templates
- HTML reports
- Additional badges
- Repo structure smell detection
- Plugin system for custom rules
- GitHub Action artifact optimizations

---

## 📄 License

To be determined.

---

## 🔗 Links

- Source: https://github.com/jbakchr/hewd
- Issues: https://github.com/jbakchr/hewd/issues
- Releases: https://github.com/jbakchr/hewd/releases
