# 🧭 **Phase 1 — Command Grouping (Step-by-Step Implementation Plan)**

Command grouping is simple to implement, but _highly impactful_ for overall UX.  
You’ll end up with a beautiful `hewd --help` output structured like:

    Analysis Commands:
      scan
      doctor
      diff

    Maintenance Commands:
      fix
      init

    Reporting Commands:
      export
      badge

Let’s break this into clean steps.

---

# ✅ **Step 1 — Decide on the command groups**

For hewd, the natural groups are:

### **1. Analysis Commands**

- scan
- doctor
- diff

### **2. Maintenance Commands**

- fix
- init

### **3. Reporting Commands**

- export
- badge

These match the conceptual model and how users think about your CLI.

We’ll define group IDs and titles:

| Group ID      | Title                |
| ------------- | -------------------- |
| `analysis`    | Analysis Commands    |
| `maintenance` | Maintenance Commands |
| `reporting`   | Reporting Commands   |

---

# ✅ **Step 2 — Add the groups in `root.go`**

Inside `NewRootCmd` (typically near the top of the function):

```go
rootCmd.AddGroup(&cobra.Group{
    ID:    "analysis",
    Title: "Analysis Commands",
})

rootCmd.AddGroup(&cobra.Group{
    ID:    "maintenance",
    Title: "Maintenance Commands",
})

rootCmd.AddGroup(&cobra.Group{
    ID:    "reporting",
    Title: "Reporting Commands",
})
```

This **declares the groups** to Cobra.

---

# ✅ **Step 3 — Assign each subcommand to its group**

For each `newXCmd()` call:

### Example for scan.go

In `newScanCmd()`:

```go
cmd.GroupID = "analysis"
```

### Example for doctor.go

```go
cmd.GroupID = "analysis"
```

### Example for diff.go

```go
cmd.GroupID = "analysis"
```

### Maintenance group

`fix.go`:

```go
cmd.GroupID = "maintenance"
```

`init.go`:

```go
cmd.GroupID = "maintenance"
```

### Reporting group

`export.go`:

```go
cmd.GroupID = "reporting"
```

`badge.go`:

```go
cmd.GroupID = "reporting"
```

---

# ✅ **Step 4 — Confirm order of adding commands**

In `root.go`, add commands _after_ declaring groups:

```go
rootCmd.AddCommand(newScanCmd())
rootCmd.AddCommand(newDoctorCmd())
rootCmd.AddCommand(newDiffCmd())

rootCmd.AddCommand(newFixCmd())
rootCmd.AddCommand(newInitCmd())

rootCmd.AddCommand(newExportCmd())
rootCmd.AddCommand(newBadgeCmd())
```

**Grouping does not rely on order**, but having them grouped visually in code helps clarity.

---

# ✅ **Step 5 — Run `hewd --help` and visually confirm**

You should now see:

    Analysis Commands:
      scan
      doctor
      diff

    Maintenance Commands:
      fix
      init

    Reporting Commands:
      export
      badge

If any command appears under the wrong section → check its `GroupID` assignment.

---

# ✅ **Step 6 — Decide on capitalization and phrasing**

Good options:

- “Analysis Commands” (recommended — consistent with docs)
- “Analysis”
- “Analysis Tools”

Same for other groups.

Consistency matters more than style — but your roadmap and docs already use similar group names.

---

# Optional Enhancements (small but impactful)

## ⭐ Add spacing between groups

Some custom help templates allow adding blank lines between sections for readability — we can do this later in Phase 6.

## ⭐ Add group descriptions

If you want, you can expand group titles with descriptions like:

    Analysis Commands:
      Scan and evaluate repository health.

    Maintenance Commands:
      Apply changes or initialize configuration.

    Reporting Commands:
      Generate outputs for CI, dashboards, or documentation.

This requires custom help template work, which we can do later.

---

# 🎉 **Phase 1 Summary**

| Step | Description                               |
| ---- | ----------------------------------------- |
| 1    | Define the command groups                 |
| 2    | Add group definitions in `root.go`        |
| 3    | Assign `cmd.GroupID` for each command     |
| 4    | Organize `AddCommand` calls (optional)    |
| 5    | Run `hewd --help` to verify               |
| 6    | (Optional) refine titles and descriptions |

This step is small, clean, and gives huge UX improvement.

---
