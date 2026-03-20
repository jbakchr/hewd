# 🩺 `hewd doctor` — Project Health Diagnostics

`hewd doctor` is the core diagnostic command of the `hewd` CLI.

It scans your repository, runs all documentation/config/structure rules, and produces a detailed project health report with scoring, severity levels, and recommended fixes.

Use this command whenever you want to understand the overall quality and maturity of a repository — whether locally or inside CI.

---

## 📘 What `hewd doctor` Does

`hewd doctor` analyzes:

- **Documentation completeness**
  - README
  - LICENSE
  - CONTRIBUTING
  - CHANGELOG
  - SECURITY
  - docs/ directory
- **Configuration files**
  - .github/workflows/\*
  - language-specific config (e.g., go.mod, package.json)
  - Dockerfiles, OpenAPI files, etc.
- **Repository structure**
  - presence of certain directories
  - multi-language indicators
  - missing architectural documentation

Rules are grouped into categories:

- **documentation**
- **config**
- **structure**

Each rule produces:

- **info**
- **warn**
- **error**

And contributes to category and overall scoring.

---

## ⭐ CLI Usage

### Basic usage:

```bash
hewd doctor
```

### Markdown output (`--md`):

```bash
hewd doctor --md > health.md
```

### JSON output (`--json`):

```bash
hewd doctor --json
```

### YAML output (`--yaml`):

```bash
hewd doctor --yaml
```

### Only run specific categories (`--only`):

```bash
hewd doctor --only documentation
```

### Exclude specific categories (`--except`):

```bash
hewd doctor --except config
```

### Fail CI on a given severity (`--fail-on`):

```bash
hewd doctor --fail-on=warn
```

This is useful when gatekeeping PRs based on severity level.

---

## 📊 Output Explanation

### 🩺 Pretty Terminal Output

The default doctor output prints:

- Overall project score
- Category scores
- Detailed rule violations
- Grouped and sorted issues
- Severity indicators
- Recommendations

Example:

```
===== OVERALL SCORE =====
  82 / 100

===== CATEGORY SCORES =====
  Documentation:  70
  Config:         90
  Structure:      85

===== DOCUMENTATION ISSUES =====
  - DOC_LICENSE_MISSING (warn): LICENSE file is missing.
  - DOC_CONTRIBUTING_MISSING (info): CONTRIBUTING.md not found.
```

---

### 🧾 Markdown Output

Markdown output is ideal for:

- generating reports
- posting in PR comments
- sharing with teams

Example:

```bash
## Documentation Issues
- **DOC_LICENSE_MISSING** (warn): LICENSE file is missing.
- **DOC_README_STALE** (info): README appears outdated.
```

---

### 🧮 Machine‑Readable Output (JSON/YAML)

Use machine output for:

- CI pipelines
- dashboards
- programmatic analysis
- diff comparisons via hewd diff

JSON schema includes:

- score
- categoryScores
- results (rule evaluations)
- fixable items
- timestamps
- version info

Example:

```bash
hewd doctor --json > report.json
```

Then you can compare two reports:

```bash
hewd diff old.json new.json
```

---

## ⚙️ Flags

| **Flag**            | **Description**               |
| ------------------- | ----------------------------- |
| --md                | Output Markdown report        |
| --json              | Output JSON report            |
| --yaml              | Output YAML report            |
| --only <category>   | Only run specific categories  |
| --except <category> | Exclude categories            |
| --score             | Show only score (CI-friendly) |
| --category-score    | Show only category score      |
| --fail-on=info      | warn                          |

---

## 🔧 CI Usage

Common CI usage:

```bash
- name: Run hewd doctor
  run: hewd doctor --fail-on=warn
```

Or with the GitHub Action:

```bash
- uses: ./.github/actions/hewd-action
  with:
    github-token: ${{ secrets.GITHUB_TOKEN }}
    pr-comment: true
    md-report: true
```

This posts an updated PR comment with the doctor report.

---

## 🩹 Fixable Issues (`hewd fix`)

Some doctor rule failures are automatically fixable by `hewd fix`:

- LICENSE
- CONTRIBUTING.md
- CHANGELOG.md
- CI workflow skeleton
- docs/ folder

Try:

```bash
hewd fix --apply
```

---

## 🗂 Scores

### Overall Score (0–100)

Calculated from rule results and severity weights.

### Category Scores:

- Documentation Score
- Config Score
- Structure Score

Use these to track repository maturity over time.

---

## 🎯 When to Use `hewd doctor`

Use it when:

- starting work on a repo
- onboarding new contributors
- preparing a project for publication
- enforcing standards in CI
- maintaining consistent documentation across repos

It’s the core “health check” command of the `hewd` ecosystem.
