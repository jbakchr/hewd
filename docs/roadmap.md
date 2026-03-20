# 🗺️ **hewd Roadmap (Updated 2026)**

*A clear view of the completed milestones, current focus, and future direction for the hewd project health toolkit.*

***

# ⭐ Phase 0 — Completed Milestones

These foundations are fully implemented and stable.

### ✔ Core Capabilities

*   Repository scanner
*   Rule engine (documentation, config, structure categories)
*   Scoring engine (overall + per‑category)
*   Auto‑fix system (safe, non‑destructive fixes)
*   Machine-readable export (`hewd export`)
*   SVG badge generator
*   Configuration loader (`.hewd/config.yaml`)
*   Consistent command + flag architecture

### ✔ Diff Engine

*   Score delta detection
*   Category delta detection
*   New/resolved issue detection
*   Pretty diff output
*   JSON/YAML/Markdown diff output
*   Regression gating logic

### ✔ GitHub Integration

*   Composite GitHub Action
*   PR comment creation/updating
*   Markdown diff for PRs
*   CI-friendly exit codes
*   Diff-based regression checks

The project is now feature-complete for an initial public release.

***

# 🌱 Phase 1 — CLI UX & Documentation Polish *(Completed)*

We’ve completed a major consistency and polish pass across all user-facing commands.

### ✔ CLI UX Improvements

*   Unified command help text structure
*   Clearer short and long descriptions
*   Fully reworked examples (no HTML entities, consistent shell syntax)
*   Enhanced error messages
*   Improved flag conflict handling
*   Support for stdout output across commands
*   Better directory creation & validation in commands
*   More readable pretty outputs
*   Cleaner Markdown output (doctor + diff)
*   Standardized timestamp formatting

### ✔ Commands polished

*   `scan`
*   `doctor`
*   `export`
*   `diff`
*   `fix`
*   `init`
*   `badge`

### ✔ Documentation polish

*   Stable docs folder
*   Consistent language across all topics
*   Updated examples matching CLI UX improvements
*   Roadmap updated (this document)

This phase is officially **complete**.

***

# 🚀 Phase 2 — v0.1.0 Release Preparation *(current phase)*

We are now preparing the first official public release.

### 🔹 Release Engineering

*   Create `CHANGELOG.md`
*   Write release notes
*   Tag `v0.1.0` using release workflow
*   Build binaries for Linux, macOS, Windows
*   Validate installation experience (`go install` + binaries)
*   Optional: Homebrew tap

### 🔹 Example Repository

Create a small reference repo demonstrating:

*   before/after reports
*   GitHub Action PR comment
*   diff output examples (Markdown, JSON)
*   badge embedding
*   `.hewd/config.yaml` usage

### 🔹 Final Output Polish

*   Ensure pretty output alignment is consistent in all commands
*   Expand Markdown output (optional collapsible sections)
*   Provide final formatting for summary sections in doctor/diff

Once complete, hewd will be ready for broad adoption.

***

# 🧪 Phase 3 — Stability, Testing & Robustness

Focus: deepening trust and reliability.

### 🔹 Test Coverage

*   Unit tests for the diff engine
*   Unit tests for scoring logic
*   Unit tests for rule evaluation
*   Integration tests for entire workflows:
    *   `scan → doctor → export`
    *   fix application
    *   diff regression logic

### 🔹 Schema Stability

*   Finalize `MachineOutput` versioning strategy
*   Provide schema samples in docs
*   Document backward/forward compatibility policy

### 🔹 Robustness Improvements

*   Graceful fallback for missing or partial repositories
*   Additional validation for `.hewd/config.yaml`
*   More helpful end‑user error messages

***

# 📈 Phase 4 — Feature Expansion

Enhancements to broaden usefulness.

### 🔹 Extended Auto-Fixers

*   README scaffolding
*   SECURITY.md template
*   CODEOWNERS generator
*   ADR template scaffolding
*   Initial project layout generation

### 🔹 Additional Rule Types

*   Repo smell detection
*   Build/config inconsistency checks
*   Outdated documentation heuristics
*   Multi-language project structure rules

### 🔹 New Output Formats

*   HTML single-file report
*   Raw text output
*   Extended JSON schema with rule metadata

***

# 🌐 Phase 5 — Ecosystem & Integrations

Expanding beyond GitHub.

### 🔹 GitHub Action Enhancements

*   Auto-upload JSON/Markdown diff artifacts
*   Auto-run `export` before diff
*   Auto-detect base report for diffs
*   Optionally merge doctor + diff into unified PR comment

### 🔹 GitLab Integration

*   GitLab MR comments
*   JUnit-style output
*   Artifact uploads

### 🔹 VS Code Extension

*   Inline documentation completeness warnings
*   Side-panel project health dashboard
*   Quick-fix suggestions

***

# 🧭 Phase 6 — Long-Term Vision

Transform hewd into a full ecosystem for repository maturity.

### 🔮 hewd Web Dashboard

*   Upload/export reports
*   Trend visualization
*   Cross-project comparisons
*   Rule-level deep dives

### 🔮 Plugin System

*   Custom organization rules
*   Domain-specific rules
*   Pluggable scoring models

### 🔮 Project Templates

Example:

```bash
hewd init --template go-service
hewd init --template python-lib
```

Templates include:

*   README
*   docs/
*   CI workflows
*   CONTRIBUTING
*   CHANGELOG
*   starter project structure

***

# 🎉 Summary

hewd is now in **final release preparation**.  
Core features are complete, CLI UX is polished, and the project is ready to move into stability, testing, and packaging.

**Next steps:**

*   Produce CHANGELOG
*   Prepare v0.1.0 release
*   Build example repository
*   Final output styling polish

After that, hewd will be ready for its first public release and a path toward a mature ecosystem.

***

If you'd like, I can also create a **shorter roadmap**, a **more visual roadmap**, or convert this into **Markdown sections you can embed into README.md**.
