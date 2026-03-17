# hewd

`hewd` is a command-line tool for initializing, scanning, and managing documentation assets for codebases.

It is written in Go and distributed as a standalone, globally‑runnable binary.

This README captures the initial state and intended behavior of the tool.

---

## 🚀 Overview

`hewd` currently provides three core commands:

- `hewd init`
- `hewd scan`
- `hewd doctor`

The **`scan`** and **`doctor`** commands support structured output through:

- `--json`
- `--yaml`
- `--pretty`

The **`init`** command does _not_ support these output flags.

---

## 🧩 Commands

### `hewd init`

Initializes basic project metadata or configuration.

Current behavior:

- Performs initial setup tasks needed by other `hewd` commands.
- May create configuration files or metadata directories (exact behavior evolving).

**Note:**  
This command **does not** support `--json`, `--yaml`, or `--pretty`.

Usage:

```bash
hewd init
```

### `hewd scan`

Analyzes the project directory and reports findings.

Initial planned capabilities:

- Detect languages present (Python, Node.js, Go, Rust, etc.)
- Identify configuration and documentation assets
- Gather metadata about the repository structure
- Output results in a machine- or human-friendly format

Supported output:

```
hewd scan --json
hewd scan --yaml
hewd scan --pretty
```

### hewd doctor

Runs diagnostic checks against the project.

Initial goals:

- Validate file structure and expected artifacts
- Identify missing or malformed configuration files
- Surface potential issues in documentation assets
- Provide actionable feedback

Supported output:

```
hewd doctor --json
hewd doctor --yaml
hewd doctor --pretty
```

## 🔧 Output Formats

### --json

Machine‑readable structured output suitable for CI pipelines or automated tooling.

### --yaml

Human‑friendly structured output.

### --pretty

Text-mode, human‑readable output.

Only scan and doctor use these.

## 📦 Installation (future)

Once releases are available:

```
curl -sSL https://<tbd>/install.sh | sh
```

Or download a binary directly from GitHub Releases.

## 🗺️ Roadmap (High-Level)

TBD.

## 📄 License

TBD.
