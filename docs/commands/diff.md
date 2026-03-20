# 🔁 `hewd diff` — Report Comparison & Regression Detection

`hewd diff` compares two machine‑readable hewd reports (hewd.json) and shows how your project’s health has changed over time.

This command is ideal for:

- tracking repository maturity
- enforcing documentation & structure standards
- detecting regressions in pull requests
- generating Markdown reports for CI/PR comments
- machine‑readable diff output for pipelines

## 📘 Overview

`hewd diff` accepts **two JSON reports**, typically produced using:

```bash
hewd export --output hewd.json
```

The diff engine then compares:

- overall score
- documentation / config / structure scores
- new issues
- resolved issues

Output formats:

- **Pretty terminal output**
- **Markdown** (--md)
- **JSON** (--json)
- **YAML** (--yaml)

---

## 🚀 Usage

### Basic:

```bash
hewd diff old.json new.json
```

### Markdown output (for PR comments):

```bash
hewd diff old.json new.json --md
```

### JSON output:

```bash
hewd diff old.json new.json --json
```

### YAML output:

```bash
hewd diff old.json new.json --yaml
```

---

## 📊 What the diff shows

### ✔ Score comparison

Shows old score, new score, and the delta:

```bash
Old: 78
New: 87
Change: +9 ↑
```

### ✔ Category score deltas

```bash
Documentation:  65 → 75   (+10)
Config:         80 → 84   (+4)
Structure:      90 → 92   (+2)
```

### ✔ New issues

Issues present in the new report but NOT in the old report.

Sorted by severity → category → ID.

### ✔ Resolved issues

Issues present in the old report but NOT in the new report.

Also sorted and grouped.

### ✔ Grouped output

Both new and resolved issues are grouped by category:

```
### documentation
- DOC_LICENSE_MISSING (warn): LICENSE file is missing.
```

---

## 🧮 Pretty Terminal Output

Example:

```bash
===== OVERALL SCORE =====
Old report score: 78
New report score: 87
Score change: +9

===== CATEGORY SCORES =====
  Documentation:   65 → 75   (+10)
  Config:          80 → 84   (+4)
  Structure:       90 → 92   (+2)

===== NEW ISSUES =====
(none)

===== RESOLVED ISSUES =====
documentation
  - DOC_LICENSE_MISSING (warn)
```

Optimized for readability within CLI environments.

---

## 📝 Markdown Output

Ideal for GitHub pull requests.

```bash
hewd diff old.json new.json --md > diff.md
```

Produces:

```
# 📊 Hewd Diff Report

## 📈 Score Summary

| Metric         | Old | New | Δ     | Trend  |
|----------------|-----|-----|-------|--------|
| Overall Score  | 78  | 87  | +9    | 🟩⬆️    |
| Documentation  | 65  | 75  | +10   | 🟩⬆️    |
| Config         | 80  | 84  | +4    | 🟩⬆️    |
| Structure      | 90  | 92  | +2    | 🟩⬆️    |

## 🆕 New Issues
_No new issues! 🎉_

## ✅ Resolved Issues
### documentation
- **DOC_LICENSE_MISSING** (warn)
```

Used by the GitHub Action to post or update PR comments.

---

## 🧾 JSON / YAML Output

Machine‑readable diff format for CI or automation.

### JSON:

```bash
hewd diff old.json new.json --json > diff.json
```

Example structure:

```bash
{
  "schemaVersion": 1,
  "hewdVersion": "0.1.0",
  "generatedAt": "...",
  "scoreDelta": 9,
  "categoryDeltas": {
    "documentation": 10,
    "config": 4,
    "structure": 2
  },
  "newIssues": [],
  "resolvedIssues": [...]
}
```

### YAML:

```bash
hewd diff old.json new.json --yaml > diff.yaml
```

---

## 🚨 Regression Gating

The diff engine can fail CI builds if regressions are found.

### Fail if score drops by N:

```bash
hewd diff old.json new.json --fail-on-score-drop=5
```

### Fail if new error‑level issues appear:

```bash
hewd diff old.json new.json --fail-on-new-errors
```

### Fail on _any_ regression (recommended in CI):

```bash
hewd diff old.json new.json --fail-on-any-regression
```

These exit with **code 1**, making them ideal for automated checks.

---

## 🤖 GitHub Action Integration

The hewd GitHub Action can automatically:

- run hewd diff
- generate Markdown diff
- post/update PR comments
- fail PRs based on regression gating
- generate JSON artifacts

Example usage:

```bash
- uses: ./.github/actions/hewd-action
  with:
    github-token: ${{ secrets.GITHUB_TOKEN }}
    diff: true
    diff-old: old.json
    diff-new: new.json
    diff-pr-comment: true
```

More details → /docs/github-action.md

---

## 📦 When to Use hewd diff

Use it when you want to:

- compare project health over time
- detect regressions in pull requests
- track repository maturity
- measure documentation improvement
- generate PR-ready Markdown summaries
- enforce documentation standards via CI
- integrate into GitHub Actions workflows

It is the “before vs after” engine for the entire `hewd` ecosystem.
