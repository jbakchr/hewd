# 🧾 **Machine Output Schema (`hewd export`)**

`hewd export` produces a stable, machine‑readable JSON (or YAML) summary of a repository’s health.  
This document explains the structure of the output and how it can be consumed by tooling, CI pipelines, dashboards, and the `hewd diff` engine.

This schema is referred to throughout the codebase as **`MachineOutput`**.

***

## 📦 Example Usage

```bash
hewd export --output hewd.json
```

Or:

```bash
hewd doctor --json > report.json
```

Generated files may be used directly with:

```bash
hewd diff old.json new.json
```

***

## 🧬 Schema Overview

The `MachineOutput` structure includes:

*   metadata about the tool and schema
*   timestamps
*   overall and category scores
*   all rule evaluation results
*   fixable items
*   project metadata

It is designed to be:

*   **deterministic**
*   **stable across versions**
*   **friendly for both machines and humans**
*   **ideal for PR comparisons, CI, and analytics**

***

## 🧱 **Top‑Level Structure**

Below is the complete top‑level structure of `MachineOutput`, exactly as used in your codebase:

```json
{
  "schemaVersion": 1,
  "hewdVersion": "string",
  "generatedAt": "RFC3339 timestamp",
  "score": 0,
  "categoryScores": {
    "documentation": 0,
    "config": 0,
    "structure": 0
  },
  "results": [
    {
      "id": "RULE_ID",
      "level": "info|warn|error",
      "category": "documentation|config|structure",
      "message": "string",
      "file": "optional file context"
    }
  ],
  "fixable": [
    {
      "ruleId": "RULE_ID",
      "message": "string",
      "filePath": "path/to/file"
    }
  ]
}
```

This structure does not include any fields that your repo does not generate — ensuring stability and clarity.

***

## 🔢 Field‑By‑Field Documentation

Below is a detailed explanation of each field.

***

### **`schemaVersion`** (int)

Indicates the version of the machine‑readable output format.

*   Incremented only when the schema changes in **breaking** ways.
*   Used by `hewd diff` to ensure compatibility.

***

### **`hewdVersion`** (string)

The version of the **hewd CLI** that produced the report.

Useful for:

*   debugging
*   tracking improvements across tool versions
*   CI reproducibility

***

### **`generatedAt`** (RFC3339 timestamp)

Timestamp of when the report was created.

Example:

    2026-03-18T12:45:32Z

***

### **`score`** (int, 0–100)

The overall project health score computed by the scoring engine.

Represents a weighted evaluation of:

*   documentation completeness
*   config completeness
*   structure soundness

***

### **`categoryScores`** (object)

Contains detailed per‑category scoring:

```json
{
  "documentation": 70,
  "config": 90,
  "structure": 85
}
```

Each category has its own scoring rules and deductions.

***

### **`results`** (array of rule evaluations)

Each entry corresponds to a rule (from the documentation, config, or structure rule sets).

Fields include:

*   **id** — unique rule identifier (`DOC_LICENSE_MISSING`)
*   **level** — `info`, `warn`, or `error`
*   **category** — rule category
*   **message** — human‑readable explanation
*   **file** — (optional) file associated with the rule

Example:

```json
{
  "id": "DOC_LICENSE_MISSING",
  "level": "warn",
  "category": "documentation",
  "message": "LICENSE file is missing."
}
```

***

### **`fixable`** (array of fix suggestions)

Lists problems that the `hewd fix` engine is capable of resolving automatically.

Each entry contains:

*   `ruleId` — rule this fix corresponds to
*   `message` — description of the fix
*   `filePath` — the file that will be created or modified

Example:

```json
{
  "ruleId": "DOC_LICENSE_MISSING",
  "message": "Create a LICENSE file",
  "filePath": "LICENSE"
}
```

***

## 📊 How MachineOutput Is Used

The `MachineOutput` schema powers several major features of hewd.

***

## 🔁 `hewd diff`

The diff engine compares **two `MachineOutput` JSON files**:

*   score delta
*   category deltas
*   new issues
*   resolved issues
*   regression detection
*   JSON/YAML/Markdown diff output

MachineOutput schemas must have matching `schemaVersion` for diffing.

***

## 🤖 CI Pipelines

MachineOutput is ideal for CI because it is:

*   deterministic
*   versioned
*   easily diffed
*   JSON/YAML friendly

Use it to power:

*   regression gating
*   dashboards
*   artifacts

***

## 🧮 Dashboards & Automation

You can build:

*   documentation health dashboards
*   trend lines over time
*   custom quality metrics
*   bots and assistants

Simply parse the `results`, `score`, and `categoryScores` fields.

***

## 📂 Example

Running:

```bash
hewd export --output report.json
```

Produces a file compatible with:

```bash
hewd diff old.json new.json
```

And:

```bash
jq '.score' report.json
jq '.results[] | select(.level=="error")' report.json
```

***

## 🔮 Future Schema Enhancements

Potential improvements (see `/docs/roadmap.md`):

*   HTML output
*   Extended metadata (contributors, commit count)
*   Per‑rule scoring weights included
*   Additional categories
*   Versioned rule sets
*   Compatibility layers for future diff versions
