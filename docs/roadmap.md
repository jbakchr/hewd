# 🗺️ **hewd Roadmap (Updated 2026‑03, post-Phase 1E progress)**

```markdown
# 🗺️ hewd Roadmap

_A clear guide to completed work, current focus, and future direction for the hewd CLI._

---

# ⭐ Phase 0 — Completed Foundations

These core building blocks are stable, tested, and form the foundation for all later phases.

### ✔ Core Capabilities

- Repository scanner
- Rule engine (documentation, config, structure)
- Scoring engine
- Machine-readable export (JSON/YAML)
- Auto-fix system (safe, non-destructive)
- `.hewd/config.yaml` support
- SVG badge generator

### ✔ Diff Engine

- Score deltas
- Category deltas
- New/resolved issue detection
- Pretty, JSON, YAML, Markdown diff outputs
- Regression gating (for CI)

### ✔ GitHub Integration

- Composite GitHub Action
- PR comment updating
- Markdown diff output
- CI-friendly exit codes
```

---

# 🌱 **Phase 1 — CLI UX & Developer Experience Polish**

_(Nearly fully complete)_

```markdown
## ✔ Phase 1A — Helptext Extraction

- All helptexts moved to `internal/helptext`
- Command files simplified
- Unified tone and structure

## ✔ Phase 1B — Global Flag Consistency

- Unified flag names and ordering
- Added ValidateOutputFormatFlags() helper
- Shared conflict rules across commands

## ✔ Phase 1C — Helptext Harmonization

- Consistent formatting across all commands
- Parallel example blocks
- No HTML entities or formatting inconsistencies

## ✔ Phase 1D — Pretty Output Polish

- Consistent color palette (cyan headers, GitHub-style severity icons)
- Standardized spacing/indentation
- Unified pretty output for scan, doctor, diff
- Centralized severity icon logic

## ✔ Phase 1E — Error Message Standardization (now ~90% complete)

### Completed:

- New structured error type: `HewdError{Msg, Hint}`
- New helpers: `Err`, `ErrHint`, `RootErr`
- Updated `scan.go`, `doctor.go`, `diff.go`, `export.go`, `badge.go`, `fix.go`, `init.go`
- Updated `main.go` for top-level error prefixing
- Standardized error tone (lowercase, no trailing period)
- Hints shown on separate lines
- All fmt.Errorf removed from commands

### Remaining:

- Apply structured error handling to:
  - internal/config
  - internal/scan
  - internal/fix
  - internal/diff
- Minor final polish across helpers
- Optional: unify hint phrasing project-wide
```

---

# 🚀 **Phase 2 — v0.1.0 Release Preparation**

_(Begins immediately after Phase 1E finalization)_

```markdown
### 🔹 Release Engineering

- Produce CHANGELOG.md
- Draft v0.1.0 release notes
- Build cross-platform binaries (macOS/Linux/Windows)
- Validate installation methods
- Tag v0.1.0
- Optional: Homebrew tap

### 🔹 Example Repository

Demonstrates:

- Typical repository structure
- Doctor and diff reports (JSON/MD)
- GitHub PR comment example
- Recommended `.hewd/config.yaml` usage

### 🔹 Final Output Polish

- Spacing and alignment review
- Validate Markdown rendering for doctor/diff
- Confirm pretty output across small/large repos
```

---

# 🧪 **Phase 3 — Stability, Testing & Reliability**

```markdown
### 🔹 Unit Tests

- Rule engine
- Scoring engine
- Fix detectors
- Diff computation

### 🔹 Integration Tests

- scan → doctor → export pipeline
- fix → apply → doctor
- diff gating (fail-on regressions)

### 🔹 Schema Stability

- MachineOutput schema versioning
- Backwards compatibility plan
- Example schemas in docs
```

---

# 📈 **Phase 4 — Feature Expansion**

```markdown
### 🔹 Extended Auto-Fixers

- README scaffolds
- SECURITY.md
- CODEOWNERS
- ADR templates

### 🔹 New Rule Types

- Repo “smell” heuristics
- Deprecated configuration detection
- Build/config consistency rules
- Multi-language project recognition

### 🔹 New Output Formats

- Single-file HTML report
- Raw text output
- Extended JSON schema metadata
```

---

# 🌐 **Phase 5 — Ecosystem Integration**

```markdown
### 🔹 GitHub Action Enhancements

- Auto-upload diff artifacts
- Auto-detect base report
- Unified “doctor + diff” PR comment

### 🔹 GitLab Support

- MR comments
- JUnit export
- Artifacts

### 🔹 VS Code Extension

- Inline documentation checks
- Editor hints for missing files
- Sidebar project health dashboard
```

---

# 🧭 **Phase 6 — Long-term Vision**

```markdown
### 🔮 hewd Web Dashboard

- Upload reports
- Compare across repos
- Trends and metrics over time

### 🔮 Plugin System

- Organization-specific rule packs
- Custom scoring modules
- Pluggable rule evaluation

### 🔮 Project Templates

Examples:
hewd init --template go-service
hewd init --template python-lib

Includes:

- README
- CONTRIBUTING.md
- docs/
- CI workflows
- CHANGELOG
- Starter layout
```

---

# 🎉 Summary

```markdown
hewd has now completed nearly all of Phase 1, including full CLI polish, helptext consistency, pretty output standardization, and most of the new error-handling system.

Once the final internal helpers are aligned with the structured error system, hewd will enter Phase 2 — preparing for the first official release: v0.1.0.
```
