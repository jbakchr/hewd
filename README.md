# 📘 hewd — Project Health, Documentation & Structure Toolkit

`hewd` is a fast, dependency‑free CLI for analyzing, comparing, and improving the overall health of software repositories.  
It evaluates documentation, configuration, and structural conventions, produces machine‑readable reports, highlights regressions, and can automatically generate missing project files.

---

## ✨ Features

- 🔍 **Scan** repositories for documentation, configuration, structure, and metadata
- 🩺 **Run diagnostics** via `hewd doctor`
- 🧮 **Score** documentation, configuration, and structure
- 🔁 **Compare reports** using `hewd diff` (score deltas, new/resolved issues)
- 🚨 **Regression gating** for CI pipelines
- 🧾 **Machine‑readable JSON/YAML output**
- 📝 **Markdown reports** for GitHub PR comments
- 🧰 **Auto‑fix** missing documentation and CI files
- 🏷️ **SVG badge generator**
- 🤖 **GitHub Action** with PR comment updating
- 🎨 **Pretty‑printed terminal output** with severity icons and color
- ❗ **Standardized error system** with hints and clean formatting

Full documentation lives in the `docs/` directory.

---

## 🚀 Quick Start

### Install (build from source)

```bash
git clone https://github.com/jbakchr/hewd
cd hewd
go build -o hewd ./cmd/hewd
sudo mv hewd /usr/local/bin/hewd
```

---

## 🔍 Scan your project

```bash
hewd scan --pretty
```

---

## 🩺 Run diagnostics

```bash
hewd doctor
```

Pretty output will highlight:

- score
- category scores
- issues with severity icons
- fixable items

Generate a Markdown report:

```bash
hewd doctor --md > health.md
```

---

## 📤 Export machine‑readable data

```bash
hewd export --json --output hewd.json
```

Pretty JSON:

```bash
hewd export --json --pretty
```

YAML:

```bash
hewd export --yaml --output hewd.yaml
```

---

## 🔁 Compare reports (diff engine)

```bash
hewd diff old.json new.json
```

Markdown diff for PR comments:

```bash
hewd diff old.json new.json --md > diff.md
```

Supports pretty, JSON, YAML, and Markdown output modes.

---

## 🧰 Auto‑Fix Missing Files

```bash
hewd fix                  # dry-run
hewd fix --apply          # write new files
```

Creates:

- LICENSE
- CONTRIBUTING.md
- CHANGELOG.md
- docs/ directory
- CI workflow templates

More fixers coming in future releases.

---

## 🏷️ Badge Generation

```bash
hewd badge --output badge.svg
```

Useful for READMEs, dashboards, and CI artifacts.

---

## 🤖 GitHub Action

hewd includes a first‑class GitHub Action enabling:

- PR comments (`hewd diff --md`)
- regression gating
- machine‑readable artifacts

Minimal usage:

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

---

## 🎨 Pretty Output

hewd uses:

- cyan headers
- GitHub‑style severity icons (`ℹ️`, `⚠️`, `❌`)
- colorized score deltas (`↑`, `↓`)
- consistent section formatting

Example:

    ===== OVERALL SCORE =====
     85 / 100

    ===== DOCUMENTATION ISSUES =====
      ⚠️  DOC_LICENSE_MISSING — LICENSE file is missing
      ℹ️  DOC_README_STALE    — README.md may be outdated

---

## ❗ Standardized Error System

hewd uses a consistent, helpful error style:

    error: cannot combine --json and --yaml
    hint: use only one machine-readable format at a time

Top-level errors include the invoked command:

    error (hewd doctor): failed to read config
    hint: ensure .hewd/config.yaml is valid yaml

All internal commands use structured `HewdError` values.

---

## 📚 Documentation

Located in `docs/`:

- 📘 Getting Started — `docs/getting-started.md`
- 🩺 Doctor Command — `docs/commands/doctor.md`
- 🔁 Diff Command — `docs/commands/diff.md`
- 🔧 Fix Command — `docs/commands/fix.md`
- 🧾 Machine Output — `docs/machine-output.md`
- ⚙️ Configuration — `docs/configuration.md`
- 🤖 GitHub Action — `docs/github-action.md`
- 🗺️ Roadmap — `docs/roadmap.md`

---

## 🏗 Project Roadmap

See the full roadmap:
👉 `docs/roadmap.md`

**Current status:**
`hewd` has completed Phases **1A through 1D**, and Phase **1E** is actively being implemented.
After Phase 1E is complete, work begins on **Phase 2 — v0.1.0 Release Engineering**.

---

## 📄 License

To be determined.

---

## 🔗 Links

- Source: <https://github.com/jbakchr/hewd>
- Issues: <https://github.com/jbakchr/hewd/issues>
- Releases: <https://github.com/jbakchr/hewd/releases>

```

---

# 🎉 README.md is now fully updated and aligned
This version reflects:

- Your exact project state
- Phase 1E progress
- Error system rewrite
- Pretty output redesign
- Command grouping
- Unified flags
- Clean, professional tone
- Modern CLI UX


```
