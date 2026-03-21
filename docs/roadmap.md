# 🗺️ **hewd Roadmap (Updated 2026‑03)**

A clear overview of completed milestones, the current development stage, and future directions for the hewd project health toolkit.

---

# ⭐ Phase 0 — Completed Foundations

These milestones represent the core hewd architecture, all of which are stable and fully implemented.

### ✔ Core Capabilities

- Repository scanner
- Rule engine (documentation, config, structure categories)
- Scoring engine
- Auto‑fix system (safe, non‑destructive fixes)
- `.hewd/config.yaml` support
- Machine‑readable export (JSON/YAML)
- SVG badge generation

### ✔ Diff Engine

- Score delta calculation
- Category delta calculation
- New/resolved issue detection
- Pretty diff output
- JSON/YAML/Markdown diff output
- Regression gating

### ✔ GitHub Integration

- Composite GitHub Action
- PR comment creation & updating
- Markdown diff for PRs
- CI exit‑codes for gating logic

### ✔ Architecture & Codebase Structure

- Clean package layout (`internal/`, `pkg/cmd/`, `docs/`)
- Consistent naming patterns
- Helper utilities for shared CLI logic

---

# 🌱 Phase 1 — CLI UX & Developer Experience Polish

_(This phase is now complete, with one optional optional micro‑polish remaining.)_

### ✔ Phase 1A — Helptext Extraction & Unification

- All `Use`, `Short`, `Long`, `Example` blocks moved out of command files
- New `internal/helptext` package
- Command files significantly cleaner
- Unified tone, structure, and voice across all commands

### ✔ Phase 1B — Global Flag Consistency Audit

- Standardized flag names (`--json`, `--yaml`, `--md`, `--pretty`, etc.)
- Unified conflict rules across commands
- Shared validation helper added (`ValidateOutputFormatFlags()`)

### ✔ Phase 1C — Helptext Harmonization

- All helptext files rewritten for consistency and professionalism
- Matched indentation, wording, formatting, and example structure
- Eliminated HTML entities and outdated phrasing

### ✔ Phase 1D — Pretty Output Polish

- Unified ANSI color palette (cyan headers, red/yellow/blue severities, green/red trend arrows)
- Consistent section headers across all commands
- Shared severity icon logic (`SeverityVisual()`)
- Standardized pretty‑output formatting for `scan`, `doctor`, and `diff`
- Category mapping standardized using `rules.CategoryForRule()`
- Clean, professional TTY output across the entire CLI

### ⚪ _Optional Micro‑Polish (not required before Phase 2)_

A tiny list of optional improvements, which can be done later if desired:

- Column alignment refinement
- Additional end‑summary blocks in pretty output
- Minor whitespace or stylistic tuning
- Monochrome fallback for terminals without color

These are explicitly **optional** and should not block release preparation.

---

# 🚀 Phase 1E — Error Message Standardization

_(This is the next active development phase.)_

### 🎯 Goals

Improve the clarity, consistency, predictability, and professionalism of all error messages.

### 🔹 Planned Work

- Normalize error phrasing (tone, sentence style, lower‑case starts, no trailing periods)
- Improve context in error messages (e.g., include file paths, rule names)
- Ensure consistent patterns for command‑specific errors
- Standardize common failure messages across commands
- Audit all `fmt.Errorf` calls for clarity
- Provide clearer guidance in CI‑failure scenarios
- Confirm markdown/JSON/YAML mode errors follow same structure
- Clearly distinguish:
  - user‑input errors
  - repo/content errors
  - internal logic errors

### 🔹 Expected Outcome

A consistent, predictable error system that strengthens UX and makes CI behavior more intelligible.

After Phase 1E is complete, hewd will be fully ready for Phase 2.

---

# 🚀 Phase 2 — v0.1.0 Release Preparation

_(This begins immediately after Phase 1E.)_

### 🔹 Release Engineering

- Create a `CHANGELOG.md`
- Publish first official release notes
- Build binaries for Linux/macOS/Windows
- Tag `v0.1.0`
- Verify installation pathways (`go install`, built binaries, GitHub Actions)
- Optional: Homebrew tap for macOS users

### 🔹 Example Repository

Create a demonstration repository showcasing:

- Before/after doctor reports
- GitHub Action PR comment examples
- Diff output examples (Markdown, JSON)
- Badge usage
- `.hewd/config.yaml` patterns

### 🔹 Final Output Polish

- Review spacing/indentation one last time
- Ensure Markdown output for doctor/diff is perfect
- Validate pretty output across varied repository sizes

---

# 🧪 Phase 3 — Stability, Testing & Reliability

### 🔹 Unit Tests

- Rule engine
- Scoring
- Diff engine
- Fix‑detection logic

### 🔹 Integration Tests

- `scan → doctor → export` pipeline
- `fix --apply`
- diff regression gating

### 🔹 Schema Stability

- Versioning strategy for `MachineOutput`
- Backwards‑compatible schema handling
- Clear documentation

---

# 📈 Phase 4 — Feature Expansion

### 🔹 Extended Auto‑Fixers

- README scaffolding
- SECURITY.md template
- CODEOWNERS generator
- ADR template generation

### 🔹 More Rule Types

- Repo smell detection
- Build/config consistency checks
- Outdated documentation heuristics
- Language‑specific structure checks

### 🔹 Additional Output Formats

- Single‑file HTML report
- Raw text
- Extended JSON schema with metadata

---

# 🌐 Phase 5 — Ecosystem Integrations

### 🔹 GitHub Action Enhancements

- Auto‑upload diff artifacts
- Auto‑detect base report
- Merge doctor+diff into unified PR comment

### 🔹 GitLab Integration

- MR comments
- Artifact creation
- JUnit‑style output

### 🔹 VS Code Extension

- Inline rule hints
- Project health dashboard
- Fix suggestions

---

# 🧭 Phase 6 — Future Vision

### 🔮 hewd Web Dashboard

- Upload reports
- Trend visualization
- Multi‑project comparisons
- Historical timelines

### 🔮 Plugin Architecture

- Custom rule packs
- Organization‑level policies
- Pluggable scoring systems

### 🔮 Project Templates

    hewd init --template go-service
    hewd init --template python-lib

Includes:

- README
- docs/
- CI workflows
- CONTRIBUTING
- CHANGELOG
- Basic scaffolding

---

# 🎉 Summary

hewd has now completed all major UX polish for commands, helptexts, and pretty output (Phase 1A–1D).

The project is currently entering **Phase 1E — Error Message Standardization**, the final step before moving to **Phase 2 — Release Engineering** and preparing the first tagged release (`v0.1.0`).
