# 📘 hewd — Project Health, Documentation & Structure Toolkit

`hewd` is a fast, dependency‑free CLI for analyzing, comparing, and improving the overall health of software repositories.

It evaluates documentation, configuration, and structural conventions, produces machine‑readable reports, highlights regressions, and can automatically generate missing project files.  
It also integrates with GitHub Actions for CI‑driven health checks.

---

## ✨ Features

- 🔍 **Scan** repositories for documentation, configuration, structure, and metadata
- 🩺 **Run diagnostics** with `hewd doctor`
- 🧮 **Score** documentation, configuration, and structure
- 🔁 **Compare reports** via `hewd diff` (score deltas, new/resolved issues)
- 🚨 **Regression gating** for CI pipelines
- 🧾 **Machine‑readable JSON/YAML output**
- 📝 **Markdown reports** for GitHub PR comments
- 🧰 **Auto‑fix** missing documentation and CI files
- 🏷️ **SVG badge generation**
- 🤖 **GitHub Action** with PR comment updating
- 🎨 **Pretty‑printed output** with severity icons and color
- ❗ **Structured error system** with consistent messages and actionable hints

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

Pretty output includes:

- overall score
- category scores
- issues with severity icons
- fixable items

Generate Markdown:

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

Supports pretty, JSON, YAML, and Markdown output.

---

## 🧰 Auto‑Fix Missing Files

```bash
hewd fix        # dry-run
hewd fix --apply  # write new files
```

Creates:

- LICENSE
- CONTRIBUTING.md
- CHANGELOG.md
- docs/ directory
- CI workflow templates

(and more in future versions)

---

## 🏷️ Badge Generation

```bash
hewd badge --output badge.svg
```

Ideal for READMEs, dashboards, or CI artifacts.

---

## 🤖 GitHub Action

hewd ships with a first‑class GitHub Action enabling:

- PR comments (`hewd diff --md`)
- regression gating
- machine‑readable report uploads

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
- colorized deltas (`⬆️`, `⬇️`)
- consistent section formatting

Example:

    ===== OVERALL SCORE =====
    85 / 100

    ===== DOCUMENTATION ISSUES =====
    ⚠️ DOC_LICENSE_MISSING — LICENSE file is missing
    ℹ️ DOC_README_STALE — README.md may be outdated

---

## ❗ Structured Error System (Phase 1E — Fully Completed)

hewd provides clean, consistent error messages:

    error: cannot combine --json and --yaml
    hint: use only one machine-readable format at a time

Command-level errors are prefixed:

    error (hewd doctor): failed to read config
    hint: ensure .hewd/config.yaml is valid yaml

Internal code uses structured `HewdError` values everywhere.

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
✔ Phase 1 (A–E) is fully complete
➡ Starting Phase 2 — _v0.1.0 Release Preparation_

---

## 📄 License

To be determined.

---

## 🔗 Links

- Source: <https://github.com/jbakchr/hewd>
- Issues: <https://github.com/jbakchr/hewd/issues>
- Releases: <https://github.com/jbakchr/hewd/releases>
````
