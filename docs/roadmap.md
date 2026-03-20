# 🗺️ **hewd Roadmap**

This roadmap outlines the ongoing development of **hewd**, a project health, documentation, and structure toolkit. It is organized into phases that reflect the current state of the project and the natural next steps toward a stable, public release.

---

# ⭐ Phase 0 — Completed Milestones

Core foundational features are finished and stable:

### ✔ Core Functionality

- Repository scanner
- Rule engine (documentation, config, structure)
- Scoring engine
- Auto‑fix system
- Machine-readable export (`hewd export`)
- SVG badge generator

### ✔ Diff Engine

- Score delta detection
- Category delta detection
- New/resolved issue detection
- Sorting & grouping
- Pretty terminal diff output
- JSON/YAML/Markdown diff output

### ✔ GitHub Integration

- Composite GitHub Action
- PR comment creation & updating
- Markdown diff comments
- Regression gating flags
- CI-friendly exit codes
- Diff-mode workflows

hewd is now stable, feature-rich, and ready for polishing.

---

# 🌱 Phase 1 — Final Pre‑Release Polish _(current phase)_

Goal: ensure hewd feels polished, professional, and cohesive.

### 🔹 Documentation polish

- `/docs/` folder created ✔
- Getting Started guide ✔
- Command-specific docs ✔
- GitHub Action docs ✔
- Configuration guide ✔
- MachineOutput schema docs ✔
- Roadmap rewrite ✔
- Polished README ✔

### 🔹 CLI UX polish

- Rewrite `--help` text for all commands ✔
- Provide consistent formatting & examples ✔
- Add clear command summaries ✔
- Improve error messages (remaining)
- Severity indicators in output (optional)

### 🔹 Output polish

- Improve Markdown diff layout
- Add compact summary section
- Add `<details>` collapsing for long issue lists
- Improve spacing in pretty output

This phase ensures hewd is easy to understand and pleasant for first‑time users.

---

# 🚀 Phase 2 — v0.1.0 Release Preparation

Goal: publish the first official release of hewd.

### 🔹 Release engineering

- Create `CHANGELOG.md`
- Write release notes
- Tag `v0.1.0` using release workflow
- Build binaries for Linux, macOS, Windows
- Validate installation experience

### 🔹 Installation support

- Provide `go install github.com/jbakchr/hewd/cmd/hewd@latest` instructions
- Confirm release artifacts function correctly
- (Optional) Add Homebrew tap

### 🔹 Example repository

Create a small demo showcasing:

- before/after reports
- GitHub Action PR comment
- diff output (JSON/Markdown)
- badge embedding

This helps new users understand hewd’s value at a glance.

---

# 🧪 Phase 3 — Stability & Testing

Goal: improve confidence and robustness.

### 🔹 Test coverage

- Unit tests for diff engine
- Unit tests for regression gating
- Unit tests for rule sorting & grouping
- Integration tests for scan → doctor → export
- Integration tests for fix mode

### 🔹 Schema stability

- Confirm versioning policy for `MachineOutput`
- Provide schema samples in `/docs`
- Document breaking-change policy

### 🔹 Robustness improvements

- More user-friendly error messages
- Graceful handling of missing files
- Validate `.hewd/config.yaml`

---

# 📈 Phase 4 — Feature Expansion

Next-wave enhancements to improve the toolkit.

### 🔹 Extended Auto-Fixers

- README scaffolding
- SECURITY.md template
- CODEOWNERS generator
- ADR template scaffolding

### 🔹 Additional rule types

- Repo smell detection
- Outdated documentation detection
- Build/config inconsistency checks
- Multi-language structure issues

### 🔹 New output formats

- HTML single-file report
- Raw text output mode
- Expanded JSON schema for rule-level metadata

---

# 🌐 Phase 5 — Ecosystem & Integrations

Broaden hewd’s usefulness beyond GitHub.

### 🔹 GitHub Action enhancements

- Automatically upload JSON/Markdown diff artifacts
- Auto-detect base report (no manual `diff-old`)
- Automatically run export before diff
- Optionally merge doctor + diff into a unified PR comment

### 🔹 GitLab integration

- GitLab MR comments
- JUnit-compatible output
- Artifact exports

### 🔹 VS Code integration

- Inline documentation completeness hints
- Side-pane project health dashboard

---

# 🧭 Phase 6 — Big Vision & Future Directions

Long-term ideas that evolve hewd into a full ecosystem.

### 🔮 hewd Web Dashboard

- Upload/export reports
- Visualize trends
- Multi-project score comparisons
- Historical views

### 🔮 Plugin system

Allow custom rules:

- organization-specific
- project-type-specific
- pluggable scoring

### 🔮 Project templates

Initialize new repositories:

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
- starter project layout

---

# 🎉 Summary

hewd is now in the **final polishing phase** before the first public release. Almost all core features are complete:

- rich diff engine
- export schema
- auto-fix system
- scoring
- regression gating
- GitHub Action

### Next milestones:

- Finish final output polish
- Create example repo
- Prepare v0.1.0 release (CHANGELOG + release notes)

After that, the project can grow into a powerful ecosystem for repository maturity, documentation quality, and structure enforcement.
