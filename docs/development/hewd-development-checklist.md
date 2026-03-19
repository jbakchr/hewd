# Hewd Development Checklist

_A reusable checklist for developing new features, improvements, or refactoring within the hewd project._

This checklist helps ensure consistent architecture, clear goals, and smooth development.

---

## 1. Feature Overview

- [ ] Describe the feature or improvement.
- [ ] Define the purpose and problem it solves.
- [ ] Explain how it fits into hewd’s architecture.
- [ ] Provide expected user experience (UX).

---

## 2. Current State

- [ ] What is already implemented?
- [ ] What is partially implemented?
- [ ] What is missing?
- [ ] Known issues or blockers.
- [ ] Relevant TODOs or design notes.

---

## 3. Repository Structure Snapshot

Include only relevant directories/files, for example:

```
internal/
diff/
diff.go
printer.go
models/
report.go
cmd/
hewd/
diff.go
```

---

## 4. Relevant Code Files

Attach or paste the necessary parts:

- [ ] Struct definitions (e.g., Report, Issue, CategoryScore)
- [ ] Functions related to the feature
- [ ] Internal helpers to reuse
- [ ] Any new files needed
- [ ] Interfaces or contracts to maintain

---

## 5. Expected Input/Output Formats

Provide examples for:

- [ ] JSON output
- [ ] YAML output
- [ ] Markdown output (for PR comments)
- [ ] Pretty terminal output
- [ ] Schema for machine-readable diff output

---

## 6. CLI Requirements

- [ ] Command syntax (e.g., `hewd diff old.json new.json`)
- [ ] Flags and their behavior
- [ ] Required/optional arguments
- [ ] Error cases and messages
- [ ] Exit code policies
- [ ] Help text expectations

---

## 7. Constraints & Preferences

- [ ] Style preferences (functions, methods, error handling)
- [ ] Dependency policy (stdlib-only, allow small libs, etc.)
- [ ] Naming consistency rules
- [ ] Performance expectations
- [ ] What patterns to avoid (e.g., generics, reflection)

---

## 8. Testing Expectations

- [ ] Table-driven tests
- [ ] Golden files (for output stability)
- [ ] Snapshot tests
- [ ] Unit vs integration testing
- [ ] Specific example test cases
- [ ] Required test coverage levels

---

## 9. Compatibility & Versioning

- [ ] Backward compatibility requirements
- [ ] Versioning impact (semantic versioning, schema versions)
- [ ] CI changes or new artifacts
- [ ] Any public API changes

---

## 10. Non-Goals

Clarify what should _**NOT**_ be done:

- [ ] Breaking changes
- [ ] Adding new directories
- [ ] Modifying existing output formats
- [ ] Introducing new dependencies
- [ ] Changing struct names or APIs unnecessarily

---

## 11. Optional Attachments

If useful:

- [ ] Example hewd report (`.json`)
- [ ] GitHub Action config or logs
- [ ] Sample project directories for testing
- [ ] Design documents or diagrams

---

## Usage

Before requesting new feature development or code generation, fill out this checklist to ensure clarity, consistency, and architectural alignment.
