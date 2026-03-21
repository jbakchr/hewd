# 🗺️ hewd Roadmap

_A clear guide to completed work, active milestones, and long‑term plans for the hewd CLI._

---

# ⭐ Phase 0 — Completed Foundations

These core building blocks are stable and form the bedrock for all future development.

### ✔ Core Capabilities

- Repository scanning
- Rule engine (documentation, config, structure)
- Scoring engine
- Machine-readable exports (JSON/YAML)
- Auto-fix system (safe, non-destructive)
- `.hewd/config.yaml` support
- SVG badge generator

### ✔ Diff Engine

- Score deltas
- Category deltas
- New/resolved issue detection
- Pretty, JSON, YAML, Markdown diff outputs
- Regression gating for CI

### ✔ GitHub Integration

- Composite GitHub Action
- PR comment updating
- Markdown diff output
- CI-friendly exit codes

---

# 🌱 Phase 1 — CLI UX & Developer Experience Polish

**Status: 100% complete (A–E)**

### ✔ Phase 1A — Helptext Extraction

- All helptexts moved to `internal/helptext`
- Commands simplified
- Unified tone and structure

### ✔ Phase 1B — Global Flag Consistency

- Unified flag names, ordering, and conflict rules
- Added `ValidateOutputFormatFlags()` helper

### ✔ Phase 1C — Helptext Harmonization

- Standardized formatting across all commands
- Parallel examples
- No stray HTML entities

### ✔ Phase 1D — Pretty Output Polish

- Unified color palette
- GitHub‑style severity icons
- Consistent section formatting
- Centralized severity icon logic

### ✔ Phase 1E — Structured Error System (Now **fully completed**)

- New structured error type: `HewdError{Msg, Hint}`
- New helpers: `Err`, `ErrHint`, `RootErr`
- Full cleanup across _all_ internal packages:
  - `config/`
  - `scan/`
  - `fix/`
  - `diff/` (JSON/YAML outputs updated)
- All raw `fmt.Errorf` removed
- Full propagation of structured errors through CLI
- Unified tone (lowercase, no trailing periods)
- Actionable, consistent hints across commands

---

# 🚀 Phase 2 — v0.1.0 Release Preparation

**Status: Next active milestone**

### 🔹 Release Engineering

- Produce `CHANGELOG.md`
- Draft `v0.1.0` release notes
- Build cross‑platform binaries (macOS/Linux/Windows)
- Validate CLI installation methods
- Tag `v0.1.0`
- _(Optional)_ Homebrew tap

### 🔹 Example Repository

Demonstrates:

- Typical repository structure
- Doctor + diff reports (pretty, MD, JSON/YAML)
- GitHub Action with PR comment
- Recommended `.hewd/config.yaml` usage

### 🔹 Final Output Polish

- Confirm spacing/alignment across commands
- Verify Markdown rendering
- Validate pretty output across small & large repos

---

# 🧪 Phase 3 — Stability, Testing & Reliability

### 🔹 Unit Tests

- Rule engine
- Scoring engine
- Fix detectors
- Diff computation

### 🔹 Integration Tests

- scan → doctor → export pipeline
- fix → apply → doctor
- diff gating in CI

### 🔹 Schema Stability

- Output schema versioning
- Backwards compatibility plan
- Example schemas in `docs/`

---

# 📈 Phase 4 — Feature Expansion

### 🔹 Extended Auto-Fixers

- README scaffolds
- SECURITY.md
- CODEOWNERS
- ADR templates

### 🔹 New Rule Types

- Repo “smell” heuristics
- Deprecated config detection
- Multi-language project recognition

### 🔹 New Output Formats

- Single-file HTML reports
- Raw text output
- Extended JSON metadata

---

# 🌐 Phase 5 — Ecosystem Integration

### 🔹 GitHub Action Enhancements

- Auto-upload diff artifacts
- Auto-detect base report
- Combined doctor + diff PR comment

### 🔹 GitLab Support

- Merge-request comments
- JUnit export
- Artifacts

### 🔹 VS Code Extension

- Inline documentation checks
- Editor hints
- Health dashboard pane

---

# 🧭 Phase 6 — Long-term Vision

### 🔮 Web Dashboard

- Upload reports
- Compare repos
- Metrics over time

### 🔮 Plugin System

- Org-specific rule packs
- Custom scoring modules
- Pluggable evaluation

### 🔮 Project Templates

- `hewd init --template go-service`
- `hewd init --template python-lib`

---

# 🎉 Summary

hewd has now completed **all of Phase 1**, including unified CLI UX, error systems, helptext polish, output formatting, and internal consistency.

Next up is **Phase 2**, preparing the first official release: **v0.1.0** — including changelog, release notes, binaries, and example repository.
