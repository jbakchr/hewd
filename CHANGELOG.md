# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [v0.1.0] — 2026-03-21

### 🎉 Initial Public Release (Phase 1 Complete)

This is the first official release of **hewd**, delivering a fully polished
CLI experience, complete documentation, a stable machine‑readable output format,
and a modern structured‑error system.

---

## ✨ Added

### Core Features

- Repository scanner with support for:
  - language detection
  - documentation & configuration discovery
  - directory/file statistics
- Rule engine for documentation, configuration, and structure checks
- Scoring engine with category scoring
- Machine‑readable output (JSON and YAML)
- Pretty‑printed terminal output
- Markdown report generation
- SVG score badge generator
- Auto‑fix system (safe, non-destructive):
  - CONTRIBUTING.md
  - LICENSE
  - CHANGELOG.md
  - docs/ skeleton
  - GitHub CI workflow template

### Diff Engine

- Score deltas and category deltas
- New/resolved issue detection
- Pretty, JSON, YAML, and Markdown diff outputs
- Regression gating for CI pipelines

### GitHub Action

- Composite GitHub Action including:
  - PR comment updates
  - regression gating
  - machine‑readable artifact publishing

---

## 🛠 Improved

### CLI UX & Developer Experience (Phase 1A–1D)

- Unified helptext formatting across all commands
- Extracted helptext into `internal/helptext`
- Consistent global flag behaviors and validation
- Standardized pretty output:
  - color palette
  - severity icons
  - spacing and alignment
- Centralized severity icon logic in rule system

### Structured Error System (Phase 1E)

- Introduced new structured error type:  
  `HewdError{Msg, Hint}`
- Added helpers:  
  `Err`, `ErrHint`, `RootErr`
- Removed all `fmt.Errorf` and raw errors from:
  - `internal/config`
  - `internal/scan`
  - `internal/fix`
  - `internal/diff`
- Ensured all user-facing errors:
  - are lowercase
  - have no trailing period
  - provide actionable hints

---

## 🧹 Internal Cleanup

- Reorganized severity icon logic from `cliutils` → `rules`
- Removed all import cycles
- Strengthened config loading and error propagation
- Improved filesystem safety in auto‑fix module
- Ensured consistent JSON and YAML encoding behavior

---

## 📚 Documentation

- Updated README and roadmap to reflect completed Phase 1
- Improved command examples and usage instructions
- Added GitHub Action documentation
- Documented machine‑output schema

---

## 🔧 Development Quality of Life

- Improved internal consistency across packages
- Added clearer developer notes in `docs/development/`
- Prepared project structure for upcoming test suite (Phase 3)

---

## 🧭 Next Steps (Phase 2)

- Create example repository
- Final output polish
- Cross‑platform binary builds
- Tag stable v0.1.0 release
- Optional: Homebrew tap

---

## [Unreleased]

Planned:

- Additional auto-fixers
- New rule types
- HTML single‑file reports
- VS Code extension
- GitLab integration
- Web dashboard
