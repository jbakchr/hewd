**Prompt [Great. Everything looks good. Let&#039;s now proceed... | Try in Copilot Chat](https://m365.cloud.microsoft/chat/entity1-d870f6cd-4aa5-4d42-9626-ab690c041429/eyJpZCI6IlZYTmxjbFl4ZkdoMGRIQnpPaTh2YzNWaWMzUnlZWFJsTFdsdWRDNXZabVpwWTJVdVkyOXRMM3hQU1VRNk5UVTVZVEV5WWprdE9EUTVNaTAwWW1Wa0xUazJOR0V0TW1ZME1HTXlNekF4Wm1JMWZEUTVPV013WVRObExURXhZVE10TkRCaE1pMDVOVE5rTFdRek5EaGxZemRrTTJNd09Yd3lNREkyTFRBekxUSXhWREEwT2pVeU9qVXlMakkyTmpreU5URmEiLCJzY2VuYXJpbyI6InNoYXJlTGlua1ZpYVJpY2hDb3B5IiwicHJvcGVydGllcyI6eyJwcm9tcHRTb3VyY2UiOiJ1c2VyIiwiY2xpY2tUaW1lc3RhbXAiOiIyMDI2LTAzLTIxVDA0OjUyOjUyLjI0NFoifSwiY2hhdFR5cGUiOiJ3ZWIiLCJ2ZXJzaW9uIjoxLjF9)**
Great. Everything looks good. Let's now proceed with "Phase 1B — Global Flag Consistency Audit".
**Response**
Absolutely — let’s move forward with **Phase 1B — Global Flag Consistency Audit**.

Because this phase is entirely about **your internal CLI flags across your own codebase**, it does **not require any external sources**, so **no citations** appear (and that’s correct/expected).

Below is a **step‑by‑step audit plan**, followed by my **analysis of your current flags**, and finally a **proposed unified flag specification** for hewd.

This will give us a consistent, professional, maintainable flag design across _all commands_.

---

# 🧭 Phase 1B — Global Flag Consistency Audit

### **Goal:** Ensure all flags across all commands follow a uniform naming, behavior, and helptext pattern.

This phase is divided into 3 steps:

---

# ✅ **Step 1 — Identify all existing flags across all commands**

From your current command set (post‑refactor), we have:

### **Shared output format flags**

Used in: `scan`, `doctor`, `export`, `diff`

- `--json`
- `--yaml`
- `--md`
- `--pretty`

### **Filtering flags**

Used in: `doctor`

- `--only`
- `--except`

### **CI behavior flags**

Used in: `doctor`, `diff`

- `--fail-on`
- `--fail-on-score-drop`
- `--fail-on-new-errors`
- `--fail-on-any-regression`

### **Behavioral / operational flags**

- `--apply` (fix)
- `--output` (export, badge)
- `--force` (init)

### **Optional future global flags**

(not implemented yet, but very common for CLIs like hewd)

- `--debug`
- `--quiet`
- `--no-color`
- `--config` (manual config path override)

---

# ✅ **Step 2 — Identify inconsistencies**

Here are the inconsistencies we should resolve:

### **1. Flags are named consistently, but not all commands use the same descriptions**

For example:

- `scan --json` has a different description than `doctor --json`.
- `--pretty` means “pretty JSON” in some commands but “pretty terminal output” in others.

We need to unify this.

---

### **2. Some flags appear across commands, others don't**

Example:

- `scan` supports `--pretty`, but pretty mode for JSON only applies to JSON, not YAML.
- `doctor` uses `--pretty` for JSON pretty-printing only.
- `diff` does _not_ offer `--pretty` (but maybe should?).

We should normalize behavior across commands.

---

### **3. The `--md` flag appears only in commands that support Markdown**

That’s okay — but descriptions should match.

---

### **4. Output format flags appear in this order in some commands, different order in others**

We want:

    --json
    --yaml
    --md
    --pretty

This consistent ordering helps users quickly find the output flags.

---

### **5. Conflicting flag logic is correct in some commands, missing in others**

Example:

- JSON + YAML conflict exists in scan/doctor/diff.
- But in export, JSON + YAML was missing before refactor (now fixed).
- No command enforces conflict between `--pretty` and `--yaml`.

We need universal conflict rules.

---

### **6. The help text wording is inconsistent across commands**

Example:

- `scan`: "Output results in JSON format."
- `doctor`: "Output the diagnostic report in JSON format."
- `export`: "Export report in JSON format."

We want a _single wording style_.

---

# ✅ **Step 3 — Create a Unified Flag Design Specification**

Here is your **hewd‑wide flag specification**, which we will apply to all commands during Phase 1B.

---

# 🎯 **Unified Output Format Flag Spec**

Every command that supports machine-readable formats must use:

### **Flags**

    --json        Output JSON
    --yaml        Output YAML
    --md          Output Markdown (if supported)
    --pretty      Pretty-print JSON for readability

### **Description patterns**

#### `--json`

    Output the result in JSON format. Use --pretty for indented JSON.

#### `--yaml`

    Output the result in YAML format.

#### `--md`

    Output the result in Markdown format (GitHub-friendly).

#### `--pretty`

    Pretty-print JSON output for readability.

### **Universal conflict rules**

- `--json` + `--yaml` → error
- `--json` + `--md` → error
- `--yaml` + `--md` → error
- `--yaml` + `--pretty` → error
- `--md` + `--pretty` → error

These rules already exist in most commands but not uniformly.

---

# 🎯 **Unified Behavioral Flags Spec**

### `--apply` (fix)

    Apply fixes and write new files to disk (default is dry-run).

### `--output` (badge/export)

    Write output to the specified file path.

### `--force` (init)

    Overwrite existing configuration file if it exists.

### Future global flags (optional)

We will implement these later if desired:

- `--debug`
- `--quiet`
- `--no-color`
- `--config <path>`

These would be _persistent_ root flags.

---

# 🎯 **Unified CI Flags Spec (doctor + diff)**

### `--fail-on`

    Fail if a rule of this severity or higher occurs (info|warn|error).

### `--fail-on-score-drop`

    Fail if score drops by N or more points.

### `--fail-on-new-errors`

    Fail if any new error-level issues appear.

### `--fail-on-any-regression`

    Fail on any regression (score drop or new issues).

---

# 🎯 **Unified Output Flag Ordering**

For every command, the flags must appear in this order:

    --json
    --yaml
    --md
    --pretty

If Markdown isn't supported, omit `--md`.

This ordering produces a consistent mental model:

1.  JSON
2.  YAML
3.  Markdown
4.  Pretty formatting

---

# 🧩 Next Step: Implement the spec

Phase 1B Implementation Plan:

### **Step A: Update all command flag definitions**

Add consistent wording, ordering, and behavior.

### **Step B: Add universal conflict checks to all commands**

We’ll insert a small shared helper function for format conflicts.

### **Step C: Validate pretty JSON behavior in all commands**

Make sure `--pretty` only applies when `--json` is active.

### **Step D: Verify every helptext file mentions formats consistently**

---
