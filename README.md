# hewd
_A fast, zero‑dependency CLI tool for initializing, scanning, and diagnosing documentation and configuration assets in software projects._

---

## 📑 Table of Contents

- [Features](#-features)
- [Installation](#-installation)
- [Quick Start](#-quick-start)
- [Commands](#-commands)
  - [hewd init](#hewd-init)
  - [hewd scan](#hewd-scan)
  - [hewd](#hewd-doctor)
- [Output Formats](#-output-formats)
- [Configuration (WIP)](#-configuration-wip)
- [Roadmap](#-roadmap)
- [License](#-license)
- [Project Links](#-project-links)

---

## ✨ Features

- 🔍 Scan project structure to detect languages, documentation, config files, and metadata
- 🩺 Run diagnostics to identify missing, malformed, or inconsistent documentation assets
- 🧰 Initialize config for project‑wide heuristics (planned/experimental)
- 📤 Structured output formats (--json, --yaml, --pretty)
- 🚀 Single static binary, no external runtime required
- 🛠️ Designed for CI pipelines, documentation bots, and developer tooling

---

## 📦 Installation

> **Note**: Installation methods below are placeholders until releases are available.

### Download binary (future)

```bash
curl -sSL https://github.com/<org>/<repo>/releases/latest/download/hewd \
  -o /usr/local/bin/hewd && chmod +x /usr/local/bin/hewd
```

### Go install (future)

```bash
go install github.com/<org>/<repo>/cmd/hewd@latest
```

### Homebrew (planned)

```bash
brew install <org>/tap/hewd
```

---

## 🚀 Quick Start

After installing:

```bash
hewd scan --pretty
```

Example output:

```bash
Project Summary:
  Files:       139
  Directories: 100

Languages detected:
  Go (6 files)
  Markdown (1 files)
  YAML (1 files)

Documentation:
  README.md: present
  LICENSE: missing
  CONTRIBUTING.md: missing

Scan complete.
```

---

## 📚 Commands

### `hewd init`

Initialize project‑level configuration or metadata.

```bash
hewd init
```

#### Current behavior (WIP):

- Creates basic project configuration
- Prepares directory metadata used by other commands

> `init` does not support --json, --yaml, or --pretty.

---

### `hewd scan`

Analyze the repository and surface structured information.

```bash
hewd scan
hewd scan --json
hewd scan --json --pretty
hewd scan --yaml
```

Detects (planned & evolving):

- Programming languages present
- Documentation assets (README.md, docs/, ADRs, etc.)
- Project metadata (LICENSE, CHANGELOG, module files, etc.)
- Configuration surface (Dockerfiles, CI configs, manifests)

Example **YAML** output:

```bash
files: 139
directories: 100
languages:
    Go: 6
    Markdown: 1
    YAML: 1
documentation:
    CONTRIBUTING.md: false
    LICENSE: false
    README.md: true
```

---

### hewd doctor

Run diagnostics on the project structure.

```bash
hewd doctor
hewd doctor --json
hewd doctor --json --pretty
hewd doctor --yaml
```

Checks (current + roadmap):

- Missing or malformed documentation
- Inconsistent file structure
- Undeclared or unused files
- Gaps in config or discovery metadata
- Areas that would benefit from documentation templates

Example **JSON** (_pretty_) output:

```bash
{
  "Files": 139,
  "Directories": 100,
  "Languages": {
    "Go": 6,
    "Markdown": 1,
    "YAML": 1
  },
  "Documentation": {
    "CONTRIBUTING.md": false,
    "LICENSE": false,
    "README.md": true
  }
}
```

---

## 🧩 Output Formats

### --json

Structured machine‑readable output for pipelines.

### --yaml

Human‑friendly structured output.

### --pretty

Rich text output for terminal use.

> Output flags apply to `scan` and `doctor`, **not** `init`.

## ⚙️ Configuration (WIP)

`hewd init` _may_ generate:

```
.hewd/
  config.yaml
```

A future configuration example:

```
project_name: my-service
scan:
  include:
    - src
    - cmd
  exclude:
    - vendor
doctor:
  enable_checks:
    - docs
    - config
    - metadata
```

---

## 🗺️ Roadmap

- [ ] Language + framework‑aware documentation heuristics
- [ ] Project maturity scoring
- [ ] ADR discovery & validation
- [ ] GitHub Action / CI integration
- [ ] Documentation template generation
- [ ] Repo structure “smell” detection
- [ ] LLM‑assisted documentation gap analysis
- [ ] Auto‑linking docs across directories

## 📄 License

To be determined.

## 🔗 Project Links

- **Source Code**: https://github.com/jbakchr/hewd
- **Issues**: https://github.com/jbakchr/hewd/issues
- **Releases**: https://github.com/jbakchr/hewd/releases
