# 🧭 Phase 1E — Error Message Standardization

## **Detailed Implementation Plan**

This phase improves the _clarity, consistency, predictability, and professionalism_ of all error messages across the hewd CLI.

Error messages are one of the strongest determinants of perceived quality, especially in CLI tools used in CI/CD. Well‑structured errors:

- Reduce confusion
- Make CI failures readable
- Help users diagnose their repos
- Make docs clearer
- Improve trust in the tool

Let’s break this phase into **five clear steps**, then I'll ask you a couple of clarifying questions so I can apply the style consistently across all commands.

---

# ⭐ Step 1 — Define hewd’s global error message style guide

A consistent error message style helps avoid drift.

### **1. Error messages should be lowercase unless proper nouns appear**

Examples:

    cannot combine --json and --yaml
    failed to parse config file: <path>
    invalid value for --fail-on: <value>

### **2. No trailing periods**

CLI errors typically avoid trailing punctuation.

### **3. No capitalization unless:**

- referencing file paths
- referring to proper nouns
- starting with a function/cause annotation (rare)

### **4. No leading newline**

Errors should not begin with `\n`.

### **5. Use simple, direct phrasing**

Preferred verbs:

- cannot
- failed to
- invalid
- missing
- expected

Avoid:

- “Oops!”
- overly verbose sentences
- emotional tone

### **6. Provide actionable information**

Examples:

    failed to read file: foo.json
    cannot open directory: .hewd/ (missing permissions)
    invalid config: missing "rules" block

### **7. Avoid ambiguity**

Bad:

    something went wrong

Good:

    failed to parse JSON: unexpected token at offset 52

### **8. Use standard error wrapping**

Every error caused by another error should wrap:

```go
return fmt.Errorf("failed to read file %s: %w", p, err)
```

### **9. Use consistent ordering**

Structure:

    action: context

Not:

    context: action failed

Correct order:

    failed to write output file: foo/bar.json

---

# ⭐ Step 2 — Identify all major error categories in hewd

Errors in hewd fall into these buckets:

### **Category 1 — Input / flag misuse**

Examples:

- invalid flag combination
- missing required argument
- missing file paths
- using flags incompatible with output formats

### **Category 2 — File system errors**

- directory does not exist
- cannot write file
- permissions
- cannot create directory

### **Category 3 — Repo-level errors**

- failed to scan directory
- missing documentation files
- incomplete metadata

### **Category 4 — Output / formatting errors**

- markdown generation failure
- JSON marshal failure
- YAML marshal failure

### **Category 5 — Diff / schema errors**

- schema mismatch
- invalid diff input
- unreadable old/new report files

### **Category 6 — Config errors**

- malformed YAML
- missing config path
- bad values

---

# ⭐ Step 3 — Create standardized error phrasing patterns for each category

These are the templates we will apply.

### **Flag errors**

    cannot combine --json and --yaml
    invalid value for --fail-on: %s
    --output is required (example: --output file.json)

### **Filesystem errors**

    failed to read file %s: %w
    failed to write file %s: %w
    failed to create directory %s: %w

### **Scan errors**

    failed to scan directory: %w

### **Config errors**

    failed to load config: %w
    invalid config: missing <field>

### **Diff errors**

    schema version mismatch: old=%d new=%d
    failed to parse diff input file %s: %w

### **Doctor errors**

    failed to run diagnostics: %w

### **Markdown/JSON/YAML output errors**

    failed to generate markdown: %w
    failed to marshal JSON: %w
    failed to marshal YAML: %w

---

# ⭐ Step 4 — Audit all existing hewd code for every error message

This includes:

- `scan.go`
- `doctor.go`
- `diff.go`
- `export.go`
- `badge.go`
- `fix.go`
- `init.go`

And also:

- `scan_pretty.go`
- `doctor_pretty.go`
- `diff_pretty.go`
- validation helpers
- config loader
- diff parser
- fix detectors

This step ensures _no outdated phrasing remains_.

I can generate a **complete standardized error-message rewrite** for all commands, but before I do, I need your preferences (see below).

---

# ⭐ Step 5 — Apply consistent error translations across all commands

This means:

- replacing old phrases
- adding `%w` error wrapping
- ensuring all error messages appear in one global style
- making sure CI-facing commands give clear messages
- ensuring pretty outputs don’t print errors (errors always go to stderr)

After this step, hewd’s errors will feel like those in top-tier CLI tools.

---

# ❓ Before I rewrite all error messages, I need your preferences:

Please answer these 3 questions so I apply the exact style you want.

---

### **Q1 — Do you want errors to include the command name?**

Options:

A) **Yes**, like:

    hewd doctor: cannot combine --json and --yaml

B) **No** (cleaner):

    cannot combine --json and --yaml

C) **Only for top-level failures** (recommended):

- command-specific errors do **not** include the name
- root command failures **do** include it

---

### **Q2 — Should error messages use bold/colored text in the terminal?**

Options:

A) **No**, keeping errors in plain red:

    ERROR: cannot read file foo.json

B) **Yes**, use subtle color:

- red for errors (`cliutils.RedBold`)
- no bold text
- keep simple punctuation

C) **No ANSI colors in error messages** (strict CI friendliness)

❗ Note: GH Actions supports ANSI colors today, so option B is modern and safe.

---

### **Q3 — Should error messages include suggestions?**

Example:

    cannot combine --json and --yaml
    hint: use only one machine-readable format at a time

Options:

A) **Yes, include hints** (recommended)  
B) **No, keep it terse**  
C) **Include hints only for common mistakes**

---
