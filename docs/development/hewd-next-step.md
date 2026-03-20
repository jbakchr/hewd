# ✅ **Overall Plan for Polishing All CLI Help Text**

We will do this in 4 small, manageable phases:

---

## **Phase 1 — Define a polished help‑text style guide**

This gives us:

- consistent tone
- consistent formatting
- consistent example style
- rules for when to use Short vs Long descriptions
- consistent command structure

This prevents rewriting help text multiple times later.

---

## **Phase 2 — Polish help text for each top‑level command**

We do one command at a time, in this order:

1.  `hewd scan`
2.  `hewd doctor`
3.  `hewd fix`
4.  `hewd export`
5.  `hewd badge`
6.  `hewd diff`
7.  `hewd init` (optional but recommended)

For each command I will generate:

- `Short` description
- `Long` description
- Usage examples
- Flag explanations
- Consistent formatting

You copy/paste directly into your Cobra command files.

---

## **Phase 3 — Polish the root help (`hewd --help`)**

We’ll produce:

- a polished banner
- short command descriptions
- example usage block
- summary of what hewd does
- link to docs

---

## **Phase 4 — Optional UX Polish**

- Add emojis to severity/help
- Add "Examples:" section in every command
- Add color to terminal output (if you want)
- Consider adding command aliases

---

# ⭐ Let's begin with **Phase 1: Define the Help‑Text Style Guide**

This is the foundation that ensures everything feels professional and cohesive.

---

# 🔧 **Phase 1 — hewd Help Text Style Guide**

Here is the exact style we’ll follow (you can approve or adjust anything):

---

## **1. Short Descriptions (one-line)**

- Always written in **imperative tone**
- Should answer: “What does this command do?”
- Avoid details — keep it simple and fast to read

**Examples:**

- “Scan the repository and detect documentation, config, and structure indicators.”
- “Run full project diagnostics and compute health scores.”

---

## **2. Long Descriptions**

- Multi-line
- Paragraph-based
- Contains:
  - What the command does
  - Why it’s useful
  - What it outputs
  - Common scenarios

**Tone:**  
Clear, confident, friendly, but professional.  
No rambling — 4–8 sentences max.

---

## **3. Examples Section**

Every command should have **3–5 examples** formatted like:

```bash
hewd command [flags]
```

Examples should be short, real, and directly useful.

---

## **4. Flags Section**

Flags should:

- Always be documented
- Use consistent formatting
- Group related flags together
- Follow this style:

<!---->

      --json        Output JSON format
      --yaml        Output YAML format
      --pretty      Pretty human-friendly output

---

## **5. Markdown-like formatting encouraged in Long text**

We use:

- backticks for commands
- minimal emoji (optional)
- short sections for readability

Example:

    `hewd scan` is a fast repository scanner...

---

## **6. Avoid:**

- Overly long explanations
- Academic tone
- Excessive emojis
- Repeating info from README
- Multi-paragraph essays
- Huge sections of output inside help text

---

## **7. Target user:**

- Developers
- Maintainers
- CI users
- Open‑source contributors

The help text should help them run the command _immediately_.

---

# 🎉 **Phase 1 is now complete.**

If you're happy with this style guide, we can proceed to:

# 👉 **Phase 2, Step 1: Polish `hewd scan` help text**

Just say:

> **“Polish help text for hewd scan”**

And we'll rewrite the Short, Long, Examples, and Flags for that command.
