# 📘 Getting Started with `hewd`

Welcome to `hewd` — a fast, dependency‑free tool for evaluating and improving the health of modern software repositories.

This guide will help you install `hewd`, run your first commands, and generate useful reports in just a few minutes.

---

## 🚀 What You Can Do With `hewd`

`hewd` helps you:

- scan a repository for documentation, structure, and config files
- run diagnostics and compute a project health score
- generate Markdown, JSON, and YAML reports
- automatically fix missing assets (LICENSE, CONTRIBUTING, etc.)
- compare two hewd reports using hewd diff
- enforce quality in CI using regression gating
- post automated PR comments using the GitHub Action

## 📦 Installation

Build from source (recommended for now):

```bash
git clone https://github.com/jbakchr/hewd
cd hewd
go build -o hewd ./cmd/hewd
```

This will produce a hewd binary in the project directory.

You can move it into your PATH if desired:

```bash
sudo mv hewd /usr/local/bin/hewd
sudo chmod +x /usr/local/bin/hewd
```

---

## 🧪 Verify Installation

Check that the CLI is available:

```bash
hewd --help
```

You should see a list of available commands:

```
hewd scan
hewd doctor
hewd diff
hewd fix
hewd export
hewd badge
```

---

## 🔍 Step 1: Scan Your Repository (`hewd scan`)

Start by scanning the current directory:

```bash
hewd scan --pretty
```

This will detect:

- programming languages
- documentation files
- configuration files
- docs/ directory
- CI workflows
- key project structure indicators

---

## 🩺 Step 2: Run Diagnostics (`hewd doctor`)

Run the full project health check:

```bash
hewd doctor
```

or a Markdown-formatted version:

```bash
hewd doctor --md > health.md
```

You’ll get:

- a list of issues grouped by category
- overall score (0–100)
- documentation/config/structure scores

This is the command most users run regularly.

---

## 🧾 Step 3: Export Machine‑Readable Output (`hewd export`)

To create a JSON report suitable for diffing or dashboards:

```bash
hewd export --output hewd.json
```

The exported file contains:

- project metadata
- rule results
- fixable issues
- category scores
- overall score
- version info
- timestamps

This is ideal for CI pipelines and trend tracking.

---

## 🔁 Step 4: Compare Two Reports (`hewd diff`)

The diff command shows how the project changed over time:

```bash
hewd diff old.json new.json
```

### Pretty diff:

- score delta
- category score changes
- new issues
- resolved issues
- sorted + grouped output

### Markdown diff (great for PR comments):

```bash
hewd diff old.json new.json --md
```

### JSON/YAML diff:

```bash
hewd diff old.json new.json --json
hewd diff old.json new.json --yaml
```

### Regression gating (for CI):

```bash
hewd diff old.json new.json --fail-on-any-regression
```

This exits with non‑zero status if quality decreases.

---

## 🛠 Step 5: Automatically Fix Issues (hewd fix)

Dry‑run:

```bash
hewd fix
```

Apply fixes:

```bash
hewd fix --apply
```

`hewd` can generate:

- LICENSE
- CONTRIBUTING.md
- CHANGELOG.md
- .github/workflows/ci.yml
- docs/ directory

_More fixers will be added over time._

---

## 🏷 Step 6: Generate a Score Badge

```bash
hewd badge --output badge.svg
```

You can include the badge in your README.

---

## 🤖 Step 7: Use `hewd` in GitHub Actions

`hewd` ships with a powerful GitHub Action.

It can:

- run hewd doctor
- run hewd diff
- detect regressions
- post or update PR comments
- output Markdown/JSON diff artifacts

Add a minimal workflow:

```bash
jobs:
  hewd:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: ./.github/actions/hewd-action
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          pr-comment: true
          md-report: true
```

See /docs/github-action.md for full details.

---

## 🎉 Next Steps

Explore additional documentation:

- ./commands/
- ./github-action.md
- ./commands/diff.md
- ./configuration.md
- ./machine-output.md

You now have everything you need to start using `hewd` to improve the health and maintainability of your project.