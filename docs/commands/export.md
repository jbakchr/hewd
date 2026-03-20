# 🧾 `hewd export` — Machine‑Readable Project Health Output

`hewd export` generates a complete, machine‑readable JSON (or YAML) representation of a repository’s documentation, configuration, and structure health.

This exported file is the foundation for:

*   diff comparisons (`hewd diff`)
*   dashboards
*   CI pipelines
*   historical tracking
*   regression detection
*   external integrations

It captures **everything** about the repository at a given point in time.

***

## 🚀 Usage

### Export a JSON report:

```bash
hewd export --output hewd.json
```

### Export YAML:

```bash
hewd export --yaml --output hewd.yaml
```

### Pretty JSON (human‑friendly):

```bash
hewd export --json --pretty
```

> If you omit `--output`, the report prints to stdout.

***

## 📦 What the Export Contains

The exported file is based on the stable `MachineOutput` schema used throughout the hewd ecosystem.

It includes:

### ✔ Schema version

Ensures compatibility across versions.

### ✔ Hewd version

The CLI version that produced the output.

### ✔ Timestamp

`generatedAt` with full time metadata.

### ✔ Overall score

A 0–100 project health score.

### ✔ Category scores

*   Documentation
*   Config
*   Structure

### ✔ Rule results

A complete list of all evaluated rules, each with:

*   rule ID
*   severity (info/warn/error)
*   message
*   category
*   file context (if any)

### ✔ Fixable items

Items that `hewd fix` can correct automatically.

### ✔ All metadata

Project stats, file counts, languages detected, etc.

***

## 📄 Example JSON Output

This is an example structure (simplified):

```json
{
  "schemaVersion": 1,
  "hewdVersion": "0.1.0",
  "generatedAt": "2026-03-18T12:45:32Z",
  "score": 84,
  "categoryScores": {
    "documentation": 70,
    "config": 90,
    "structure": 85
  },
  "results": [
    {
      "id": "DOC_LICENSE_MISSING",
      "level": "warn",
      "category": "documentation",
      "message": "LICENSE file is missing."
    },
    {
      "id": "STR_DOCS_DIR_MISSING",
      "level": "warn",
      "category": "structure",
      "message": "docs/ directory is missing."
    }
  ],
  "fixable": [
    {
      "ruleId": "DOC_LICENSE_MISSING",
      "message": "Create a LICENSE file",
      "filePath": "LICENSE"
    }
  ]
}
```

***

## 🔁 Using Exported Reports with `hewd diff`

`hewd diff` requires two exported reports:

```bash
hewd export --output old.json
# make changes…
hewd export --output new.json

hewd diff old.json new.json
```

This allows hewd to compute:

*   score deltas
*   category deltas
*   new issues
*   resolved issues

Markdown, JSON, and YAML diff formats are fully supported.

Learn more → `docs/commands/diff.md`

***

## 🧪 CI Usage Example

Store the export as an artifact for later comparison:

```yaml
- name: Export health report
  run: hewd export --output hewd.json

- uses: actions/upload-artifact@v4
  with:
    name: hewd-report
    path: hewd.json
```

Later jobs or workflows can download and diff:

```yaml
- uses: actions/download-artifact@v4
  with:
    name: hewd-report
    path: old/

- name: Diff reports
  run: hewd diff old/hewd.json new/hewd.json --md
```

***

## 📖 Why Use `hewd export`?

### ✔ For CI pipelines

Generate reports to enforce quality standards.

### ✔ For diffing

Compare any two snapshots of a project’s health.

### ✔ For dashboards

Build internal tools or visualizations.

### ✔ For automation

MachineOutput is stable and self-contained.

### ✔ For historical tracking

Track maturity changes over releases.

***

## 🎯 When to Use This Command

Use `hewd export` when you want to:

*   generate input for `hewd diff`
*   create a baseline before making changes
*   record project health at specific commits
*   feed data into GitHub Actions
*   integrate hewd into build pipelines
*   analyze repo state programmatically

This command is an essential part of the `hewd` workflow.
