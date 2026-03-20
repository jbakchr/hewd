# 🗺️ `hewd` Roadmap

This document outlines the short‑term, mid‑term, and long‑term roadmap for **`hewd`**, based on the current capabilities of the tool and the natural direction that will strengthen its usefulness and maturity over time.

`hewd` already includes:

- project scanning
- diagnostics (doctor)
- scoring engine
- auto-fix system
- JSON/YAML/Markdown output
- diff engine
- regression gating
- GitHub Action integration
- PR comment updating
- badge generation

The roadmap below builds on this foundation.

---

## ⭐ Phase 0 — Completed Milestones

These major foundational tasks are already complete:

### ✔ Core Features

- Project scanner
- Rule engine (documentation/config/structure)
- Scoring engine
- Auto-fix system
- Machine-readable export
- SVG badge generator

### ✔ Diff Engine

- Score delta detection
- Category delta detection
- New/resolved issue detection
- Sorting + grouping
- JSON diff output
- YAML diff output
- Markdown diff output
- Pretty terminal diff output

### ✔ CI & GitHub Integration

- Regression gating
- Composite GitHub Action
- PR comment creation + updating
- Strict exit-code behavior
- Diff-mode CI workflows

`hewd` is now stable and usable by teams.

---

## 🌱 Phase 1 — Pre‑Release Polish (Current Phase)

These tasks help present hewd as a polished, mature CLI tool.

### 🔹 Documentation Cleanup (in progress)

- /docs/ folder structure
- Getting started guide
- Command‑specific documentation
- GitHub Action documentation
- Diff engine documentation
- Roadmap
- Configuration guide
- Machine-output schema documentation
- Updated README (short, skimmable, professional)

### 🔹 CLI UX Polish

- Review --help text for all commands
- Ensure examples are included for every command
- Consistent formatting in error messages
- Add emoji severity indicators (⚠️, 🔥, ℹ️)
- Make pretty output more visually balanced

### 🔹 Output Polish

- Improve formatting of diff Markdown tables
- Add `<details>` sections for collapsed issue lists
- Add compact summary section to Markdown diff

This polish makes the project feel “real” to new visitors.

---

## 🚀 Phase 2 — Initial Public Release (v0.1.0)

Tasks for the first tagged release:

### 🔹 Release Engineering

- Tag v0.1.0 using existing release workflow
- Generate binaries for Linux/macOS/Windows
- Write release notes

### 🔹 Improved Installation Support

- Provide `go install github.com/jbakchr/hewd/cmd/hewd@latest` instructions
- Add installation tests
- Consider providing a Homebrew tap (optional)

### 🔹 Example Repository

Provide a small demo repo showing:

- a before/after diff
- GitHub Action comments
- doctor JSON report
- badge usage

This helps new users understand the value instantly.

---

## 🧪 Phase 3 — Testing & Stability

These tasks solidify hewd’s reliability.

### 🔹 Test Coverage

- Unit tests for diff engine
- Unit tests for regression gating
- Tests for rule engine grouping/sorting
- Integration tests for doctor mode
- Integration tests for fix mode

### 🔹 Schema Stability

- Version MachineOutput format
- Document breaking-change policy for schemas
- Provide schema samples in /docs

### 🔹 Robustness

- Handle missing files gracefully
- More helpful error messages
- Validate config files (.hewd/config.yaml)

---

## 📈 Phase 4 — Feature Expansion

These tasks extend hewd’s usefulness.

### 🔹 Extended Auto-Fixers

- README scaffolding
- SECURITY.md template
- CODEOWNERS generation
- ADR template scaffolding

### 🔹 Additional Rules

- Repo smell detection (e.g., “many languages but no docs”)
- Dependency file inconsistencies
- Build file detection
- Outdated documentation detection (timestamps, mismatches)

### 🔹 New Output Formats

- HTML reports (single-file output)
- Raw text mode
- More detailed JSON schema for issues

---

## 🌐 Phase 5 — Ecosystem & Integrations

### 🔹 GitHub Action Enhancements

- Upload diff artifacts automatically
- Auto-detect base report (no need for diff-old)
- Auto-run export before diff
- Support merging doctor + diff into unified PR comment

### 🔹 GitLab CI Integration

Support incoming GitLab users with:

- JUnit output
- GitLab MR comments
- Artifact export tools

### 🔹 VSCode Extension (future)

- Inline documentation completeness hints
- Project health dashboard

---

## 🧭 Phase 6 — Big Vision & Long‑Term

These are forward-looking ideas that could greatly expand hewd’s impact.

### 🔮 hewd Web Dashboard (self-hosted)

- Upload/export reports
- View project history
- Visual score evolution
- Multi-project views

### 🔮 Plugin System

Let users write custom rules, e.g.:

- project-specific documentation rules
- organization-wide structure checks
- custom config checks

### 🔮 Project Templates

Generate project skeletons:

```bash
hewd init --template=go-service
hewd init --template=python-lib
```

Each template includes:

- README
- docs folder
- CI
- CONTRIBUTING
- CHANGELOG
- starter code layout

---

## 🎉 Summary

`hewd` is currently in the polishing phase, with all major features implemented:

- rich diff engine
- export schema
- auto-fix system
- scoring
- GitHub Action
- regression gating

Next immediate steps:

- polish output and docs
- prepare for v0.1.0 release
- create a small example repo

Beyond that, hewd has the potential to become a **general tool for repository maturity, documentation quality, and maintainability tracking** — a niche with real value and very few good tools.



