# 🤖 **hewd GitHub Action**

The **hewd GitHub Action** allows you to automatically run `hewd doctor`, `hewd diff`, regression gating, and Markdown reporting on every pull request.  
It is implemented as a **composite action** using your repository's own code, and includes:

- `doctor` mode (default)
- `diff` mode
- Markdown PR comments (create/update)
- Regression gating (fail PRs on regressions)
- JSON diff artifacts (optional)
- Lightweight, dependency‑free execution

This document explains how to use the action in your workflows.

---

## 📦 Action Location

The action lives inside your repository:

    .github/actions/hewd-action/

You reference it from workflows using:

```yaml
uses: ./.github/actions/hewd-action
```

---

## 🚀 Quick Start (Doctor Mode)

Add this workflow to `.github/workflows/hewd.yml`:

```yaml
name: hewd

on:
  pull_request:

jobs:
  hewd:
    runs-on: ubuntu-latest

    permissions:
      contents: read
      pull-requests: write

    steps:
      - uses: actions/checkout@v4

      - name: Run hewd action
        uses: ./.github/actions/hewd-action
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          pr-comment: true
          md-report: true
```

This will:

- Run `hewd doctor`
- Generate a Markdown report
- Post or update a single PR comment

---

## 🔁 Diff Mode (Recommended for PRs)

Diff mode compares **old vs new** hewd reports and is ideal for PR regression detection.

### Step 1 — Generate old.json (base branch)

Example:

```yaml
- uses: actions/checkout@v4
- run: hewd export --output old.json
```

### Step 2 — Generate new\.json (PR branch)

```yaml
- run: hewd export --output new.json
```

### Step 3 — Run diff mode

```yaml
- uses: ./.github/actions/hewd-action
  with:
    github-token: ${{ secrets.GITHUB_TOKEN }}
    diff: true
    diff-old: old.json
    diff-new: new.json
    diff-pr-comment: true
    diff-md-report: true
```

This will:

- Run `hewd diff old.json new.json`
- Generate JSON and Markdown diff outputs
- Post or update a PR comment containing the formatted diff
- Run regression gating (fail PR if regressions are detected)

---

## 🚨 Regression Gating

The action uses this command internally during diff mode:

```bash
hewd diff old.json new.json --fail-on-any-regression
```

A PR will fail (exit code `1`) if:

- the overall score decreases
- any category score decreases
- any new issues appear

This enforces quality standards in PRs automatically.

---

## 📝 PR Comment Behavior

The action creates **one** PR comment and **updates it** on every run.

It identifies its own comment using the header:

    📊 Hewd Diff Report

This prevents duplicate comments and keeps PRs clean.

### In diff mode, the comment contains:

- overall score delta table
- category score deltas
- new issues (grouped, sorted)
- resolved issues
- trend arrows & color indicators

### In doctor mode, the comment contains:

- Markdown doctor report

---

## 🧱 Input Reference

These are the inputs defined in `action.yml`.

### General Inputs

| Input          | Description                      | Default               |
| -------------- | -------------------------------- | --------------------- |
| `github-token` | Required for posting PR comments | `${{ github.token }}` |
| `pr-comment`   | Post/update doctor report in PR  | `true`                |
| `md-report`    | Use Markdown output for doctor   | `true`                |

### Diff Mode Inputs

| Input             | Description                    | Default |
| ----------------- | ------------------------------ | ------- |
| `diff`            | Enable diff mode               | `false` |
| `diff-old`        | Path to old hewd JSON report   | `""`    |
| `diff-new`        | Path to new hewd JSON report   | `""`    |
| `diff-md-report`  | Generate Markdown diff         | `true`  |
| `diff-pr-comment` | Post/update diff comment in PR | `true`  |

---

## 🧠 Doctor Mode Flow

When `diff=false`, the action:

1.  Builds the hewd binary
2.  Runs `hewd doctor`
3.  Generates Markdown (`--md`)
4.  Posts/updates PR comment (if enabled)

Used for general diagnostics.

---

## 🔁 Diff Mode Flow

When `diff=true`, the action:

1.  Builds hewd
2.  Runs JSON diff
3.  Runs Markdown diff
4.  Writes `diff.json` and `diff.md`
5.  Runs regression gating
6.  Posts/updates PR comment
7.  Fails CI if regressions are detected

Used for pull request quality enforcement.

---

## 🎨 Example PR Comment

The action will produce a Markdown diff similar to:

```md
# 📊 Hewd Diff Report

## 📈 Score Summary

| Metric        | Old | New | Δ   | Trend |
| ------------- | --- | --- | --- | ----- |
| Overall Score | 78  | 87  | +9  | 🟩⬆️  |
| Documentation | 65  | 75  | +10 | 🟩⬆️  |
| Config        | 80  | 84  | +4  | 🟩⬆️  |
| Structure     | 90  | 92  | +2  | 🟩⬆️  |

## 🆕 New Issues

_No new issues! 🎉_

## ✅ Resolved Issues

### documentation

- **DOC_LICENSE_MISSING** (warn)
```

---

## 🧪 Example Full PR Workflow (Diff Mode)

```yaml
name: hewd-ci

on:
  pull_request:

jobs:
  hewd-diff:
    runs-on: ubuntu-latest
    permissions:
      contents: read
      pull-requests: write

    steps:
      - uses: actions/checkout@v4

      # Export the baseline report from main
      - name: Export base report
        run: |
          git fetch origin main
          git checkout origin/main
          hewd export --output old.json

      # Export the PR report
      - name: Export PR report
        run: |
          git checkout ${{ github.sha }}
          hewd export --output new.json

      # Run hewd diff
      - name: hewd diff
        uses: ./.github/actions/hewd-action
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          diff: true
          diff-old: old.json
          diff-new: new.json
          diff-pr-comment: true
```

---

## 📂 Implementation Details

Internally, the action:

- builds the hewd binary using Go
- installs it into `/usr/local/bin/hewd`
- runs a Bash `entrypoint.sh`
- authenticates via the GitHub token
- uses the `gh api` command to create/update PR comments
- handles missing inputs safely
- exits with meaningful exit codes
- supports diff and doctor seamlessly

See `.github/actions/hewd-action/entrypoint.sh` for full logic.

---

## 🧭 Troubleshooting

### ❌ “hewd: command not found”

Ensure the build step in the action is present (this is included by default).

### ❌ Action doesn’t post comments

Check that:

- the workflow includes `permissions: pull-requests: write`
- `github-token` is passed
- the PR is not from a fork without elevated permissions

### ❌ Regression gating fails unexpectedly

Confirm your reports (`old.json`, `new.json`) use the same schema version.
