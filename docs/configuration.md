# ⚙️ **hewd Configuration**

hewd supports a simple, YAML‑based configuration system that lets you customize:

- which rules run
- severity levels
- scan include/exclude paths
- how strict CI gating should be
- rule weighting
- per‑category or per‑rule adjustments

Configuration is optional — hewd works with zero configuration out of the box — but adding a configuration file enables fine‑tuning for specific repositories or team workflows.

---

## 📁 Configuration File Location

When running:

```bash
hewd init
```

hewd creates:

    .hewd/
      config.yaml

This file is loaded automatically by all hewd commands:

- `hewd doctor`
- `hewd scan`
- `hewd fix`
- `hewd export`
- `hewd diff`

You do **not** need to pass any flags to load it.

---

## 📄 Example Configuration File

Below is a typical `.hewd/config.yaml`:

```yaml
rules:
  DOC_README_MISSING: true
  DOC_LICENSE_MISSING: true
  DOC_CONTRIBUTING_MISSING: true
  STR_DOCS_DIR_MISSING: true
  CFG_CI_WORKFLOW_MISSING: true

weights:
  DOC_LICENSE_MISSING: 3
  DOC_README_MISSING: 2

scan:
  include: []
  exclude:
    - node_modules
    - vendor

failOn:
  severity: error
```

This is only an example — customize freely.

---

## 🧩 Configuration Sections Explained

### 1. **rules**

Enable or disable specific rules by their ID.

```yaml
rules:
  RULE_ID: true|false
```

Example:

```yaml
rules:
  DOC_LICENSE_MISSING: false
  CFG_CI_WORKFLOW_MISSING: true
```

This allows you to:

- turn off rules that don’t apply to your project
- selectively enable/disable rule categories
- enforce organization‑specific standards

If a rule is missing from the file, hewd uses its default behavior.

---

### 2. **weights**

Customize the _impact_ of specific rule violations on scoring.

```yaml
weights:
  RULE_ID: number
```

Example:

```yaml
weights:
  DOC_LICENSE_MISSING: 3
  DOC_SECURITY_MISSING: 5
```

Rule weights are numerical — higher means a larger deduction in score.

Useful when you want:

- missing LICENSE to impact score more than missing README
- structural issues to matter more than docs issues

---

### 3. **scan**

Controls which files/directories the scanner should include or ignore.

```yaml
scan:
  include:
    - src
  exclude:
    - node_modules
    - vendor
    - build
```

#### Include

If provided, hewd scans **only** these paths.

#### Exclude

Exclusion always works, even if `include` is empty.

This is useful for:

- monorepos
- large vendor directories
- nested subprojects

---

### 4. **failOn**

Controls CI behavior for doctor mode.

```yaml
failOn:
  severity: info|warn|error
```

Example:

```yaml
failOn:
  severity: warn
```

This tells `hewd doctor` to exit with non‑zero status if a violation of the given severity (or higher) is found.

`failOn` applies only to **doctor mode**.  
For diff gating, use diff flags:

- `--fail-on-score-drop`
- `--fail-on-new-errors`
- `--fail-on-any-regression`

---

## 🔧 How Configuration Interacts With Commands

- **doctor** uses:
  - `rules`
  - `weights`
  - `scan.include` / `scan.exclude`
  - `failOn.severity`

- **diff** uses:
  - nothing directly (only schema)
  - but reports reflect config-driven scoring

- **scan** uses:
  - `scan.include`
  - `scan.exclude`

- **fix** uses:
  - only rule enable/disable (does not use weights)

- **export**:
  - includes all configured behavior inside its output

- **badge**:
  - uses scoring influenced by `weights` and `rules`

---

## 🧪 Validating Configuration

To check whether your configuration is being applied correctly:

```bash
hewd doctor --json
```

Look at:

- rule results
- scoring breakdown
- category scores

To confirm scan paths:

```bash
hewd scan --json | jq '.files'
hewd scan --pretty
```

---

## 🧭 Recommended Config for Most Repositories

A good baseline for most teams:

```yaml
rules:
  DOC_README_MISSING: true
  DOC_LICENSE_MISSING: true
  DOC_CONTRIBUTING_MISSING: true
  STR_DOCS_DIR_MISSING: true
  CFG_CI_WORKFLOW_MISSING: true

weights:
  DOC_LICENSE_MISSING: 3
  STR_DOCS_DIR_MISSING: 2
  CFG_CI_WORKFLOW_MISSING: 3

scan:
  exclude:
    - node_modules
    - vendor
    - dist
    - .git

failOn:
  severity: warn
```

---

## 🔮 Future Enhancements

Configuration may expand to include:

- per-project rule overrides
- custom rule definitions
- custom fix templates
- automatic detection of language‑specific best practices
- separate severity adjustment per rule

See `/docs/roadmap.md` for projections.
