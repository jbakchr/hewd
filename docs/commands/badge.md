# 🏷️ `hewd badge` — Generate Project Health Badges

The `hewd badge` command generates an **SVG badge** that displays your repository’s overall project health score.  

The badge can be included in README files, dashboards, or documentation pages to quickly show the quality level of a project.

Badges are generated locally, do not depend on external services, and can be committed directly to your repository.

***

## 🚀 Usage

### Basic usage

```bash
hewd badge --output badge.svg
```

This will generate an SVG badge representing the current project score calculated by `hewd doctor`.

### Overwrite an existing badge

If you already have a badge named `badge.svg`:

```bash
hewd badge --output badge.svg
```

Running the same command again simply overwrites the file.

### Choose a custom output file

```bash
hewd badge --output docs/health-badge.svg
```

### Use in CI

Typical CI usage (after running `hewd doctor` or `hewd export`):

```bash
hewd badge --output badge.svg
```

You can then upload it as an artifact or commit it back to your repo (if using bots or special tokens).

***

## 🎨 Badge Appearance

The generated badge includes:

*   **Overall score (0–100)**
*   Automatic **color scaling** based on score
*   Clean, minimalistic SVG designed for README embedding
*   No external dependencies (self‑contained SVG)

While the exact styles may evolve, badges follow a scoring color scale similar to:

*   🟩 Green — Excellent
*   🟨 Yellow — Moderate
*   🟥 Red — Needs improvement

(Actual colors are determined programmatically in the badge generator.)

***

## 📦 Example Output

Example SVG preview (simplified):

```svg
<svg ...>
  <rect ... />
  <text>hewd score: 87</text>
  <rect fill="#4caf50" ... />
</svg>
```

Rendered inline in a README, it appears as a standard shields-style badge.

***

## 📄 Adding the Badge to a README

After generating:

```bash
hewd badge --output badge.svg
```

Add the badge to your README:

```md
badge.svg
```

Or if the badge lives in `docs/`:

```md
docs/health-badge.svg
```

If you want to host it from a GitHub branch:

```md
https://raw.githubusercontent.com/USERNAME/REPO/main/badge.svg
```

Replace `USERNAME` and `REPO` accordingly.

***

## 🔧 Integration Into CI Pipelines

You can generate fresh badges during CI builds:

```yaml
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Run hewd badge
        run: hewd badge --output badge.svg
      - name: Upload badge as artifact
        uses: actions/upload-artifact@v4
        with:
          name: hewd-badge
          path: badge.svg
```

To commit badges automatically, you would typically use a bot user or token with push permissions.

***

## 🌱 Future Enhancements

Planned or possible future extensions:

*   Category‑specific badges (docs score, config score, structure score)
*   Multi‑badge sets (one badge per category)
*   Version badges (`hewd version`)
*   CLI options for custom colors and formats

See `/docs/roadmap.md` for more.

