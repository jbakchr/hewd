# 🔍 `hewd scan` — Repository Structure & Documentation Scanner

`hewd scan` inspects your repository and produces a high‑level summary of its structure, documentation files, configuration files, languages, and metadata.

It is fast, dependency‑free, and forms the foundation for all other hewd commands — including `hewd doctor`, `hewd export`, and `hewd diff`.

This command is ideal when you want a quick overview of the state of a repository.

---

## 🚀 Usage

### Basic:

```bash
hewd scan
```

### Pretty / human‑friendly output:

```bash
hewd scan --pretty
```

### JSON output:

```bash
hewd scan --json
```

### YAML output:

```bash
hewd scan --yaml
```

### Output to a file:

```bash
hewd scan --json > scan.json
```

---

## 📘 What `hewd scan` Does

hewd scan performs a lightweight analysis of your repository and detects:

### 📄 Documentation Files

- `README` / `README.md`
- `LICENSE`
- `CONTRIBUTING.md`
- `CHANGELOG.md`
- `SECURITY.md`
- Other markdown documentation

### 🧰 Configuration Files

- `.github/workflows/*`
- `go.mod`, `package.json`
- Dockerfiles
- OpenAPI specs
- Misc project metadata

### 📁 Structure Indicators

- Presence of `docs/` directory
- CI directories
- Multi‑language project markers
- Root-level structure consistency

### 📊 Repository Stats

- Total file count
- Total directory count
- Detected programming languages
- Key paths and patterns

The output is intentionally informative but not overly verbose — use `hewd doctor` for full rule evaluation.

---

## 🌈 Output Modes

### ▶️ Pretty Output

Pretty output provides a human‑friendly summary:

    ===== SCAN SUMMARY =====
    Languages: go
    Documentation: README.md, LICENSE
    Config files: go.mod, .github/workflows/ci.yml
    Docs directory: missing
    Total files: 53
    Total directories: 18

This is ideal for quick local inspection.

---

### 🧾 JSON Output

JSON output is stable, machine‑readable, and used internally by:

- `hewd doctor`
- `hewd export`
- CI pipelines
- dashboards
- testing frameworks

Example:

```bash
hewd scan --json > scan.json
```

JSON includes:

- files discovered
- languages
- documentation status
- config files
- metadata
- directory structure indicators

---

### 📄 YAML Output

For environments where YAML is preferred:

```bash
hewd scan --yaml
```

---

## 📦 Use Cases

Use `hewd scan` when you want to:

- quickly understand a repository’s structure
- prepare a repo for diagnostics
- inspect missing documentation
- gather metadata before exporting
- feed scan output into custom tooling
- power local scripts or dashboard automations

It’s the fastest way to get a “big picture” view of a repository.

---

## 🔧 Options

| Flag       | Description                   |
| ---------- | ----------------------------- |
| `--pretty` | Human‑friendly summary output |
| `--json`   | Output JSON                   |
| `--yaml`   | Output YAML                   |

---

## 🛠 Tips

- Use `hewd scan --pretty` before running `hewd doctor`
- For automation, prefer `--json` or `--yaml`
- You can pipe JSON into tools like `jq` for filtering:

```bash
hewd scan --json | jq '.languages'
```

---

## 🎯 When to Use `hewd scan`

Run this command when you want fast, broad insight into:

- What documentation your repo has
- What config files are present
- Whether CI workflows exist
- What languages the project uses
- What issues doctor mode might report

It’s the “quick overview” tool of the hewd command suite.
