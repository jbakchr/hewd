# hewd — Root Help Output (v0.1.0‑rc.1 Reference)

This file contains the canonical root help output for **hewd** as of the  
**v0.1.0‑rc.1** release candidate.  
It is intended for:

- UX reference
- CLI regression testing
- Documentation consistency
- Contributor onboarding

This snapshot represents the expected top-level help output a user sees when
running:

```

hewd --help

```

---

## Root Help Output (Rendered)

> **Note:** ANSI colors (bold white, bold cyan) are shown here in plain text
> for readability. In the actual CLI, headers appear bold‑white and the word
> “hewd” appears bold‑cyan in descriptive text.

---

**hewd — repository health diagnostics, scoring, and automated fixes**

**hewd** analyzes repository health by evaluating documentation,
configuration, and structural conventions. It provides fast scanning,
actionable feedback, health scores, diff reports, and automated fixes.

## Features

- Fast, dependency-free repository scanner
- Curated rules for documentation, structure, and configuration
- Automated fixes for common issues
- JSON, YAML, Markdown, and pretty outputs
- Regression gating for CI pipelines
- GitHub Action for PR comments
- SVG badge generation

Use **hewd** to maintain consistent documentation, detect regressions,
enforce standards, and track repository maturity over time.

---

## Typical Workflow

1. **Scan the repository**

```

hewd scan

```

2. **Run full diagnostics to identify issues and compute scores**

```

hewd doctor

```

3. **Export a snapshot of the current health state (before changes)**

```

hewd export --json --output old.json

```

4. **Apply an improvement discovered from running `hewd doctor`**  
   Example improvements:

- Create a LICENSE file
- Add a CONTRIBUTING.md
- Improve README content
- Fix configuration or structure issues

5. **Export a new snapshot after making improvements**

```

hewd export --json --output new\.json

```

6. **Compare reports using the diff engine**

```

hewd diff old.json new\.json --md > diff.md

```

---

## Usage

```

hewd \[command] \[flags]

```

---

## Examples

_(See “Typical Workflow” above for how these commands fit together — all commands shown here are safe to run.)_

### Scan a repository

```

hewd scan

```

### Run full diagnostics

```

hewd doctor

```

### Export machine-readable project health

```

hewd export --json --output hewd.json

```

### Compare reports using the diff engine

```

hewd diff old.json new\.json --md > diff.md

```

### Apply automated fixes

```

hewd fix --apply

```

### Generate an SVG badge

```

hewd badge --output badge.svg

```

### Initialize a hewd configuration file

```

hewd init

```

---

## Analysis Commands

- **diff**  
  Compare two hewd reports and show score, category, and issue differences.

- **doctor**  
  Run full diagnostics and compute documentation, config, and structure scores.

- **scan**  
  Scan the repository and detect documentation, config, and structure indicators.

---

## Maintenance Commands

- **fix**  
  Automatically generate missing documentation, structure, and CI files.

- **init**  
  Initialize a new hewd configuration in the current repository.

---

## Reporting Commands

- **badge**  
  Generate an SVG badge displaying the project's health score.

- **export**  
  Export a complete machine-readable project health report.

---

## Additional Commands

- **completion**  
  Generate the autocompletion script for the specified shell.

- **help**  
  Help about any command.

---

## Flags

```

-h, --help      help for hewd
-v, --version   version for hewd

```

---

**Use `hewd [command] --help` for more information about a command.**
