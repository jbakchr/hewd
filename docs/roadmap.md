# 🗺️ **hewd Roadmap (Updated 2026)**

A clear overview of completed milestones, current work, and long-term direction.

---

# ⭐ Phase 0 — Completed Foundations

### ✔ Core Features

- Repository scanner
- Rule engine (documentation, config, structure)
- Scoring engine
- Auto‑fix system
- `.hewd/config.yaml` and configuration loader
- Machine-readable export (JSON/YAML)
- SVG badge generator

### ✔ Diff Engine

- Score deltas
- Category deltas
- New/resolved issues
- Pretty, JSON, YAML, Markdown diff outputs
- Regression gating

### ✔ GitHub Integration

- Composite GitHub Action
- PR comment posting/updating
- Markdown diff output
- CI exit codes for gating

These foundations make hewd feature-complete for its first public release.

---

# 🌱 Phase 1 — CLI UX & Documentation Polish _(Completed)_

### ✔ CLI Improvements

- Unified helptext format using `internal/helptext`
- Consistent command descriptions & examples
- Accurate flag conflict validation
- Clearer error messages
- Improved Markdown output
- Standardized timestamp formatting
- Cleaned command implementations (`scan`, `doctor`, `export`, `diff`, `fix`, `init`, `badge`)

### ✔ Documentation

- Updated README.md
- Updated roadmap.md
- Clean docs folder
- Consistent formatting throughout

This phase is fully complete.

---

# 🚀 Phase 2 — v0.1.0 Release Preparation _(Current Phase)_

### 🔹 Release Engineering

- Create `CHANGELOG.md`
- Write release notes
- Build binaries for Linux, macOS, and Windows
- Tag `v0.1.0`
- Validate installation experience (`go install`, binaries)
- (Optional) Homebrew tap

### 🔹 Example Repository

A minimal demo repo showing:

- before/after reports
- GitHub PR comment demo
- diff outputs
- badge embedding
- `.hewd/config.yaml` in action

### 🔹 Final Output Polish

- Improve pretty output alignment
- Enhance Markdown formatting (`<details>`, sections)
- Clean terminal summary sections

---

# 🧪 Phase 3 — Stability, Testing & Robustness

### 🔹 Test Coverage

- Unit tests for diff, scoring, rule engine, fix detection
- Integration tests (`scan → doctor → export`, `fix --apply`, diff gating)

### 🔹 Schema Stability

- Finalize `MachineOutput` versioning
- Add schema examples to docs
- Backward compatibility policy

### 🔹 Robustness Improvements

- Improved handling of missing files
- Additional config validation
- Friendlier error behavior

---

# 📈 Phase 4 — Feature Expansion

### 🔹 Extended Auto‑Fixers

- README scaffolding
- SECURITY.md template
- CODEOWNERS generator
- ADR templates

### 🔹 More Rule Types

- Repo smell detection
- Build/config consistency checks
- Outdated documentation detection
- Multi-language project structure rules

### 🔹 New Output Formats

- HTML single-file reports
- Raw text output
- Expanded JSON schema

---

# 🌐 Phase 5 — Ecosystem Integration

### 🔹 GitHub Action Enhancements

- Auto-upload JSON/Markdown artifacts
- Auto-run export before diff
- Automatic base-report detection
- Unified doctor + diff PR comment

### 🔹 GitLab Integration

- MR comments
- JUnit-compatible output
- Artifact uploads

### 🔹 VS Code Extension

- Inline rule feedback
- Documentation completeness hints
- Project health dashboard

---

# 🧭 Phase 6 — Long-Term Vision

### 🔮 hewd Web Dashboard

- Report uploads
- Trend visualizations
- Multi-project comparisons
- Historical timelines

### 🔮 Plugin System

- Custom rules
- Organization-specific rule packs
- Pluggable scoring systems

### 🔮 Project Templates

Example:

```bash
hewd init --template go-service
hewd init --template python-lib
```

Templates include:

- README
- docs/
- CI workflows
- CONTRIBUTING
- CHANGELOG
- starter layout

---

# 🎉 Summary

hewd has completed foundational development and full CLI polish.  

The project is now in **Phase 2**, preparing for its **v0.1.0 release** with final polishing, examples, and binaries.

After release, hewd will evolve into a full ecosystem of tools, integrations, and developer‑facing automation.