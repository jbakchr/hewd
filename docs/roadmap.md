# 🗺️ **hewd Roadmap**

_A clear, structured roadmap for the future evolution of the hewd project._

hewd already includes:

- Project scanning
- Diagnostics (`hewd doctor`)
- Scoring engine
- Auto-fix system
- JSON/YAML/Markdown output
- Diff engine (new/resolved issues, score deltas, category deltas)
- Regression gating
- GitHub Action with PR comment updating
- Badge generation

hewd is now **fully functional** and ready for polishing, documentation, and its first public release.

---

# ⭐ Phase 0 — Completed Milestones

(_These foundational pieces are done and stable._)

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
- New/resolved issues
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

hewd is now stable and usable by teams.

---

# 🌱 Phase 1 — Pre‑Release Polish (Current Phase)

_Goal: polish the tool so it feels professional, consistent, and pleasant to use._

### 🔹 Documentation Polish

- Create `/docs/` folder structure ✔
- Getting started guide ✔
- Command-specific docs (doctor, diff, fix, scan, export, badge) ✔
- GitHub Action docs ✔
- Roadmap ✔
- Configuration guide ✔
- Machine-output schema ✔
- Updated README (light, skimmable, professional) ✔
- Add demo output to README ✔

### 🔹 CLI UX Polish (in progress)

- Rewrite `--help` text for all commands
- Add examples for each command
- Consistent error messages
- Severity emojis (`⚠️`, `🔥`, `ℹ️`)
- Improve alignment & pretty output spacing

### 🔹 Output Polish

- Improve diff Markdown tables
- Add `<details>` collapsing sections for long output
- Add compact summary section to diff MD output
- Improve terminal pretty output formatting

This makes hewd easy to understand, easy to adopt, and pleasant to use.

---

# 🚀 Phase 2 — Initial Public Release (`v0.1.0`)

_Goal: ship the first public version of hewd._

### 🔹 Release Engineering

- Tag `v0.1.0`
- Use existing GitHub Release workflow
- Generate binaries for Linux/macOS/Windows
- Write release notes

### 🔹 Installation Support

- Provide `go install github.com/jbakchr/hewd/cmd/hewd@latest` instructions
- Validate compilation via CI
- (Optional) Provide a Homebrew tap

### 🔹 Example Repository

Create a minimal repo demonstrating the tool:

- before/after diff
- GitHub Action PR comment
- doctor report
- badge usage

This helps new users immediately understand hewd’s value.

---

# 🧪 Phase 3 — Testing & Stability

_Goal: increase confidence and reliability._

### 🔹 Test Coverage

- Unit tests for diff engine
- Unit tests for regression gating
- Tests for rule grouping/sorting
- Integration tests for:
  - doctor
  - scan
  - fix
  - export
- End-to-end test: `scan → doctor → export → diff`

### 🔹 Schema Stability

- Confirm versioning strategy for `MachineOutput`
- Document breaking-change policy
- Provide schema samples under `/docs/`

### 🔹 Error Handling & Robustness

- More user-friendly error messages
- Confirm behavior on missing files
- Validate `.hewd/config.yaml`

---

# 📈 Phase 4 — Feature Expansion

_Goal: add more convenience, more rules, better outputs._

### 🔹 Extended Auto‑Fixers

- README scaffolding
- SECURITY.md template
- CODEOWNERS generation
- ADR template generation

### 🔹 Additional Rules

- Repo smell detection
- Dependency file inconsistencies
- Build file detection
- Outdated/stale documentation detection via timestamps

### 🔹 Output Formats

- HTML single-page report
- Raw text mode
- Expanded JSON schema for advanced tooling

---

# 🌐 Phase 5 — Ecosystem Integrations

_Goal: expand hewd into a multi-CI and multi-environment ecosystem._

### 🔹 GitHub Action Enhancements

- Upload JSON & MD diff artifacts
- Auto-detect base report (no manual `diff-old` needed)
- Auto-run export inside action
- Optionally merge doctor + diff into one PR comment

### 🔹 GitLab Integration

- GitLab MR comment support
- JUnit output
- GitLab artifact support

### 🔹 VSCode Extension

- Inline documentation completeness hints
- Side-panel project health dashboard

---

# 🧭 Phase 6 — Big Vision (Long‑Term)

_Goal: evolve hewd into a more comprehensive ecosystem._

### 🔮 hewd Web Dashboard

- Upload/export reports
- Visualize project score timelines
- Multi-repository view
- Compare repos side-by-side

### 🔮 Plugin System

Allow user-defined or org-defined rules:

- custom documentation rules
- organization-specific structure rules
- custom fixers
- extended export format

### 🔮 Project Templates

Generate project scaffolding:

```bash
hewd init --template go-service
hewd init --template python-lib
```

Templates include:

- README
- docs folder
- CI pipelines
- CONTRIBUTING
- CHANGELOG
- starter code layout

---

# 🎉 Summary

hewd is now in the **pre‑release polishing phase**, with all major features complete:

- diff engine
- export schema
- auto-fix system
- scoring
- GitHub Action
- regression gating

### **Next immediate steps:**

- CLI help text polish
- Output polish (pretty + Markdown)
- Prepare for `v0.1.0` release
- Create example repo

Beyond that, hewd is well-positioned to become a powerful ecosystem for repository maturity, documentation quality, and maintainability tracking — something very few tools provide today.

---

# ✔ Ready for next step?

If you'd like, I can now help you with:

- Polishing CLI help text (Option 1)
- Polishing diff MD output (Option 2)
- Creating an example repo
- Preparing a v0.1.0 release
- Adding badges to README
- Writing a changelog

Just tell me what you want to do next!
