**Prompt [Thanks. Everything now seems to work. But before... | Try in Copilot Chat](https://m365.cloud.microsoft/chat/entity1-d870f6cd-4aa5-4d42-9626-ab690c041429/eyJpZCI6IlZYTmxjbFl4ZkdoMGRIQnpPaTh2YzNWaWMzUnlZWFJsTFdsdWRDNXZabVpwWTJVdVkyOXRMM3hQU1VRNk5UVTVZVEV5WWprdE9EUTVNaTAwWW1Wa0xUazJOR0V0TW1ZME1HTXlNekF4Wm1JMWZHVTROalZoTldJekxXSmpZbUl0TkdZMU15MDRaRGd3TFdVMVkySmxPREZsWkdRMVpud3lNREkyTFRBekxUSXhWREV6T2pFME9qQXlMalEwTURRMk5UVmEiLCJzY2VuYXJpbyI6InNoYXJlTGlua1ZpYVJpY2hDb3B5IiwicHJvcGVydGllcyI6eyJwcm9tcHRTb3VyY2UiOiJ1c2VyIiwiY2xpY2tUaW1lc3RhbXAiOiIyMDI2LTAzLTIxVDEzOjE0OjAyLjQ4OVoifSwiY2hhdFR5cGUiOiJ3ZWIiLCJ2ZXJzaW9uIjoxLjF9)**
Thanks. Everything now seems to work. But before we continue I would like for you to update our current <entity type="File" ReferenceId="677dee53-700e-446b-8ccd-56250d2981b4">roadmap.md</entity> so that we can keep better track of everything. I've attached the current <entity type="File" ReferenceId="677dee53-700e-446b-8ccd-56250d2981b4">roadmap.md</entity> file for your reference.
**Response**
Absolutely — updating the roadmap **does not require any external factual knowledge**, so according to the developer instructions, **no search tools will be used** and **no citations will appear**.  
This is purely an **internal project document refactor**, which is fully allowed.

You’ve now completed a huge amount of work:

- Phase 1A → Helptext extraction
- Phase 1B → Flag consistency + shared validation
- Phase 1C → Helptext harmonization
- Phase 1D → Pretty output polish
- Phase 1E → Error system standardization (in progress, doctor.go done)

The roadmap should reflect all of this progress **and** show exactly what remains.

Below is your **updated, clean, fully aligned roadmap.md**, rewritten to match the current state of hewd and prepare for Phase 2 (Release Engineering).

You can replace your existing `docs/roadmap.md` with this version.

---

# 🗺️ **hewd Roadmap (Updated 2026‑03)**

_A clear guide to completed work, current focus, and future direction._

---

# ⭐ Phase 0 — Completed Foundations

These core components are shipped, stable, and power all other phases.

### ✔ Core Features

- Repository scanner
- Rule engine (documentation, config, structure categories)
- Scoring engine
- Machine‑readable output (JSON/YAML)
- Auto‑fix system (safe, never overwrites existing files)
- SVG badge generator
- `.hewd/config.yaml` support
- Configuration loader
- CI‑friendly behavior across commands

### ✔ Diff Engine

- Score delta computation
- Category delta computation
- New/resolved issue detection
- Pretty diff output
- JSON/YAML/Markdown diff formats
- Regression gating logic

### ✔ GitHub Integration

- Composite GitHub Action
- PR comment updating (Markdown diff)
- Machine‑readable artifacts
- Exit codes suitable for CI gating

---

# 🌱 Phase 1 — CLI UX, Error System, and Dev Experience Polish

_(The largest pre‑release polish phase. Mostly complete.)_

---

## ✔ Phase 1A — Helptext Extraction & Organization

- All command helptexts moved to `internal/helptext`
- Commands significantly simplified
- Unified style across all descriptions
- Fully consistent example formatting

## ✔ Phase 1B — Global Flag Consistency Audit

- Same ordering and naming across commands
- Consistent descriptions
- Output‑format flags normalized (`--json`, `--yaml`, `--md`, `--pretty`)
- Added shared validator: `ValidateOutputFormatFlags()`
- Clear and actionable validation errors

## ✔ Phase 1C — Helptext Harmonization

- Unified structure (Use → Short → Long → Example)
- Consistent paragraph formatting
- Clean, readable Markdown bullets
- No HTML entities
- Parallel tone and style across all commands

## ✔ Phase 1D — Pretty Output Polish

- Consistent color palette (cyan headers, red/yellow/blue severity, green/red trend arrows)
- Section headers standardized (`===== HEADER =====`)
- Clean spacing and indentation rules
- Unified severity icons (GitHub‑style: ℹ️, ⚠️, ❌)
- Shared `SeverityVisual()` helper
- Modern and readable terminal output
- Pretty output for scan, doctor, diff fully aligned

## ✔ Phase 1E — Error Message Standardization _(in progress)_

### Completed so far:

- New structured error type `HewdError{Msg, Hint}`
- Colorized, lowercase `error:` prefix
- Hints shown on separate lines
- `%w` preserved for underlying error detail
- Top-level errors formatted using `RootErr()`
- `scan.go` fully updated
- `doctor.go` fully updated
- Markdown rendering helper restored
- CI gating logic standardized
- Output flag errors standardized
- Pretty output error paths cleaned

### Remaining for full completion:

- Apply the new error system to:
  - `diff.go`
  - `export.go`
  - `badge.go`
  - `fix.go`
  - `init.go`
  - `main.go` (top-level error printing)
- Update any remaining helpers (scan/config/diff/fix utilities) to use `cliutils.Err` / `ErrHint`
- Final internal consistency check
- Optional final micro‑polish of pretty outputs

Once these are updated, Phase 1E is fully complete.

---

# 🚀 Phase 2 — v0.1.0 Release Preparation

_(Next major milestone — begins immediately after Phase 1E.)_

### 🔹 Release Engineering Tasks

- Create `CHANGELOG.md`
- Write release notes
- Build binaries for macOS/Linux/Windows
- Verify installation paths (`go install`, binaries)
- Tag `v0.1.0`
- Optional: Homebrew tap

### 🔹 Example Repository

Demonstrates:

- Before/after reports
- GitHub PR comment example
- Badge usage
- `.hewd/config.yaml` guidance
- Typical CI setup

### 🔹 Final Output Polish

- Final spacing/alignment review
- Ensure Markdown outputs are perfect
- Cross‑check pretty output across large/small repos

---

# 🧪 Phase 3 — Stability, Testing & Reliability

### 🔹 Unit Tests

- Rule engine
- Scoring engine
- Fix‑detection logic
- Diff engine

### 🔹 Integration Tests

- scan → doctor → export
- fix → apply → re-scan
- diff gating across branches

### 🔹 Schema Stability

- Final MachineOutput schema
- Backwards compatibility plan
- Example schemas in docs

---

# 📈 Phase 4 — Feature Expansion

### 🔹 Extended Auto‑Fixers

- README scaffolding
- SECURITY.md
- CODEOWNERS
- ADR templates

### 🔹 More Rule Types

- Repo smell heuristics
- Deprecated configs
- Build/config consistency
- Multi-language project layouts

### 🔹 New Output Formats

- Single‑file HTML
- Raw text
- Extended metadata in JSON

---

# 🌐 Phase 5 — Ecosystem Integration

### 🔹 GitHub Action Expansion

- Auto-upload diff artifacts
- Auto-detect old JSON for diff
- Unified “doctor + diff” PR comment

### 🔹 GitLab support

- MR discussion comments
- JUnit export
- Artifact uploads

### 🔹 VS Code Extension

- Inline diagnostics
- Documentation quality hints
- Sidebar health dashboard

---

# 🧭 Phase 6 — Future Vision

### 🔮 hewd Web Dashboard

- Upload reports
- Project health trends
- Multi-repo insights

### 🔮 Plugin System

- Custom scoring modules
- Org‑specific rule packs
- Plug-in rule evaluators

### 🔮 Templates

    hewd init --template go-service
    hewd init --template python-lib

Templates provide:

- docs/
- README
- CI workflows
- CONTRIBUTING
- CHANGELOG
- starter layout

---

# 🎉 Summary

hewd is now extremely close to its first official release.

You have completed:

- Phase 1A
- Phase 1B
- Phase 1C
- Phase 1D
- Major portions of Phase 1E

### The remaining work is sharply defined and small.

Once Phase 1E is finished, Phase 2 (v0.1.0 release engineering) can begin immediately.
