# 🧰 `hewd fix` — Auto‑Fix Missing Documentation & Structure

`hewd fix` automatically creates missing files and directories recommended by the hewd diagnostic engine (hewd doctor).

It helps you quickly bring a repository up to a healthy baseline with consistent documentation and structure.

Auto‑fix supports:

- generating missing documentation files
- adding default CI workflows
- scaffolding documentation directories
- preparing repos for open‑source readiness

Use this command _after_ running:

```bash
hewd doctor
```

to apply fixes for issues marked as **fixable**.

---

## 🚀 Usage

### Dry‑run (default)

```bash
hewd fix
```

Shows which fixes would be applied without modifying the repository.

### Apply fixes

```bash
hewd fix --apply
```

This writes changes to disk.

---

## 🛠 What hewd fix Can Generate Today

The auto‑fix engine currently supports creating:

### ✔ LICENSE

A default open-source license placeholder if no license file exists.

Generated file:

```bash
LICENSE
```

### ✔ CONTRIBUTING.md

Explains contribution guidelines for collaborators.

Generated file:

```bash
CONTRIBUTING.md
```

### ✔ CHANGELOG.md

A simple starter changelog following recommended conventions.

Generated file:

```bash
CHANGELOG.md
```

### ✔ Default CI Workflow

Creates a minimal GitHub Actions CI workflow:

Generated file:

```bash
.github/workflows/ci.yml
```

### ✔ docs/ directory

Creates an empty **docs/** folder if missing.

Generated path:

```bash
docs/
```

---

## 📦 Example Output

Running:

```bash
hewd fix
```

might produce:

```bash
[fix] Would create LICENSE
[fix] Would create CONTRIBUTING.md
[fix] Would create CHANGELOG.md
[fix] Would create .github/workflows/ci.yml
[fix] Would create docs/ directory
```

Running with **--apply**:

```bash
hewd fix --apply
```

would produce:

```bash
[fix] Created LICENSE
[fix] Created CONTRIBUTING.md
[fix] Created CHANGELOG.md
[fix] Created .github/workflows/ci.yml
[fix] Created docs/ directory
```

---

## 🔧 When to Use `hewd fix`

Use this command when:

- starting a new repository
- preparing a project for open‑source publication
- onboarding collaborators
- cleaning up older repositories
- restoring missing documentation
- enforcing minimum quality standards

It provides a fast way to bring repositories to a consistent baseline.

---

## 🧹 Notes & Future Expansion

The auto‑fix engine is intentionally conservative and will never overwrite existing files.

Future versions may add support for:

- README scaffolding
- SECURITY.md
- CODEOWNERS
- Architectural document templates
- ADR boilerplates

Follow the project roadmap for more.
