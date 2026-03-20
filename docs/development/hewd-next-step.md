# 🧭 **Final CLI Polish Pass — Development Plan**

We break this into **6 phases**, each small enough to complete independently and in sequence.

This gives us clarity, direction, and checkpoints.

---

# **Phase 1 — Command Grouping (High Impact, Small Effort)**

Grouping commands improves discoverability and makes `hewd --help` look professional.

## 🎯 Goals

- Create logical command groups:
  - **Analysis Commands:** scan, doctor, diff
  - **Maintenance Commands:** fix, init
  - **Reporting Commands:** export, badge
- Update `root.go` to define these groups
- Assign each subcommand to its group

## 🛠 Steps

1.  Add groups to `NewRootCmd` using `AddGroup(&cobra.Group{...})`
2.  Add each command to the correct group via `cmd.GroupID = "analysis"` etc.
3.  Verify help output (run `hewd --help`)
4.  Refine wording of group titles if needed

## 📌 Outcome

Users get a **clean, structured**, and intuitive help output.

---

# **Phase 2 — Global Flag Audit & Consistency Pass**

We ensure consistency across commands for flags like:

- `--json`
- `--yaml`
- `--md`
- `--pretty`
- `--debug`
- `--no-color`
- `--quiet` (optional future)

## 🎯 Goals

- Make flags consistent across all commands
- Ensure all commands use the same descriptions
- Validate conflict rules (e.g. JSON vs YAML vs Markdown)
- Make sure every command follows the same flag ordering pattern

## 🛠 Steps

1.  Review every command’s `Flags()` section
2.  Identify inconsistent wording or missing flags
3.  Align descriptions using a shared style
4.  Implement consistent conflict logic (already mostly done)
5.  Document naming rules in `CONTRIBUTING.md`

## 📌 Outcome

Commands feel consistent and “designed”, not accidental.

---

# **Phase 3 — Command Help Structure Validation**

Each command’s helptext should feel:

✔ parallel in structure  
✔ equal in tone  
✔ equal in length  
✔ clear  
✔ legible

## 🎯 Goals

- Ensure all helptext files in `internal/helptext` have consistent formatting
- Check for broken indentation
- Ensure examples are realistic and helpful
- Prevent editorial drift between commands

## 🛠 Steps

1.  Review each helptext file side-by-side
2.  Normalize structure:
    - 2–3 paragraphs max
    - Description → capabilities → purpose
    - Bullet lists use same format
3.  Normalize examples:
    - Same style, spacing, indentation
4.  Remove unnecessary or overly verbose lines
5.  Make sure no HTML escapes remain

## 📌 Outcome

The help system feels unified and professional.

---

# **Phase 4 — Pretty Output Polishing (Optional but High Value)**

This involves improving:

- alignment
- spacing
- indentation
- severity icons
- section headers
- horizontal separators

This affects:

- `scan_pretty.go`
- `doctor_pretty.go`
- `diff_pretty.go`

## 🎯 Goals

- Improve whitespace and readability
- Add consistent color usage (if coloring is desired)
- Standardize headers like `===== OVERALL SCORE =====`
- Review severity indicators (info/warn/error)

## 🛠 Steps

1.  Decide on a unified format
2.  Apply the description to each file
3.  Ensure consistency across commands
4.  Validate on repos of varying sizes

## 📌 Outcome

hewd feels more polished, more readable, and more modern.

---

# **Phase 5 — Error Message Standardization**

Error messages should be:

- short
- precise
- actionable
- consistent
- non-technical (for user-facing errors)

## 🎯 Goals

- Ensure all errors follow similar phrasing
- Improve messages for:
  - missing flags
  - invalid flag combinations
  - missing files
  - invalid config
  - schema mismatch
- Add context where needed (e.g. include filenames)

## 🛠 Steps

1.  Search the codebase for `fmt.Errorf`
2.  Extract all error strings
3.  Normalize style (e.g. always lowercase, no trailing period)
4.  Improve clarity where needed

## 📌 Outcome

Clear, predictable, and helpful error messages.

---

# **Phase 6 — Optional: Custom Help Template (Advanced Polish)**

This is optional but provides a **huge visual bump**.

Cobra allows custom templates for:

- command summaries
- flag listings
- examples
- group sections
- usage formatting

## 🎯 Goals

- Create a template that matches hewd’s brand & tone
- Replace the default Cobra help template
- Make `--help` a joy to read

## 🛠 Steps

1.  Design a template in a separate file (e.g. `help_template.go`)
2.  Use `cobra.AddTemplateFunc` + `cmd.SetUsageTemplate`
3.  Test across commands

## 📌 Outcome

hewd feels like a fully polished product.

---

# 🧩 Summary of the Final CLI Polish Plan

| Phase | Description                     | Goal                              |
| ----- | ------------------------------- | --------------------------------- |
| **1** | Command grouping                | Better structure in `hewd --help` |
| **2** | Global flag consistency         | Predictable UX                    |
| **3** | Helptext polish                 | Clean, unified documentation      |
| **4** | Pretty output improvements      | High-quality terminal UX          |
| **5** | Error message polish            | Clear, actionable errors          |
| **6** | Custom help template (optional) | Professional CLI experience       |
