# 📘 _**hewd**_ — Project Health, Documentation & Structure Toolkit

`hewd` is a fast, dependency‑free CLI for analyzing, comparing, and improving the overall health of software repositories.  
It helps teams enforce documentation standards, maintain consistent structure, detect regressions, and automatically fix common issues — locally or in CI.

---

## ✨ Features

- 🔍 **Scan** repositories for documentation, configuration, structure, and metadata
- 🩺 **Diagnose** project health via `hewd doctor`
- 🧮 **Score** documentation, config, and structure categories
- 🔁 **Compare reports** using `hewd diff` (new issues, resolved issues, score deltas)
- 🚨 **Regression gating** for CI pipelines
- 🧾 **Machine-readable JSON/YAML export**
- 📝 **Markdown reports** for GitHub PR comments
- 🧰 **Auto-fix** missing docs and CI files
- 🏷️ **SVG badge** generator
- 🤖 **GitHub Action** with PR comment updating

Full documentation is available in the `docs/` directory.

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

### Export machine-readable data

```bash
hewd export --output hewd.json
```

### Compare reports (diff engine)

```bash
hewd diff old.json new.json
```

Pretty, JSON, YAML, and Markdown output supported.

***

## 📊 Example Outputs

### 🔍 Running `hewd doctor`

```bash
===== OVERALL SCORE =====
 82 / 100

===== CATEGORY SCORES =====
 Documentation: 70
 Config:        90
 Structure:     85

===== DOCUMENTATION ISSUES =====
 - DOC_LICENSE_MISSING (warn): LICENSE file is missing.
 - DOC_CONTRIBUTING_MISSING (info): CONTRIBUTING.md not found.

===== CONFIG ISSUES =====
 (none)

===== STRUCTURE ISSUES =====
 - STR_DOCS_DIR_MISSING (warn): docs/ directory not found.
```

***

### 🔁 Diff Example

```bash
===== OVERALL SCORE =====
Old: 78
New: 87
Change: +9 ↑

===== CATEGORY SCORES =====
Documentation: 65 → 75 (+10)
Config:        80 → 84 (+4)
Structure:     90 → 92 (+2)

===== NEW ISSUES =====
(none)

===== RESOLVED ISSUES =====
documentation
 - DOC_LICENSE_MISSING (warn)
 - DOC_README_STALE (info)
```

***

### 🤖 GitHub PR Comment Example (`--md`)

Markdown reports include:

*   Score changes
*   Category score deltas
*   New issues / resolved issues
*   Trend indicators
*   Grouped sections for readability

***

## 🤖 GitHub Action

The included GitHub Action can:

*   Run `hewd doctor` or `hewd diff`
*   Post or update PR comments
*   Run regression gating (`--fail-on-any-regression`)
*   Export Markdown + JSON diff artifacts

Minimal setup:

```yaml
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

See `docs/github-action.md` for details.

***

## 🧰 Auto‑Fix Mode

```bash
hewd fix           # dry-run
hewd fix --apply   # apply fixes
```

Creates:

*   LICENSE
*   CONTRIBUTING.md
*   CHANGELOG.md
*   docs/ directory
*   Basic CI workflow templates

More fixers arriving in future releases.

***

## 📚 Documentation

All extended documentation lives in **`docs/`**:

*   📘 Getting Started — `docs/getting-started.md`
*   🩺 Doctor — `docs/commands/doctor.md`
*   🔁 Diff — `docs/commands/diff.md`
*   🔧 Fix — `docs/commands/fix.md`
*   🧾 Machine Output — `docs/machine-output.md`
*   ⚙️ Configuration — `docs/configuration.md`
*   🤖 GitHub Action — `docs/github-action.md`
*   🗺️ Roadmap — `docs/roadmap.md`

***

## 🏗 Roadmap

See the full roadmap:  
👉 `docs/roadmap.md`

*Current status: hewd is in **Phase 2 — v0.1.0 Release Preparation**.*

***

## 📄 License

To be determined.

***

## 🔗 Links

*   Source: <https://github.com/jbakchr/hewd>
*   Issues: <https://github.com/jbakchr/hewd/issues>
*   Releases: <https://github.com/jbakchr/hewd/releases>


