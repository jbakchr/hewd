#!/usr/bin/env bash
set -euo pipefail

echo "== Hewd GitHub Action =="

echo "Working directory: $(pwd)"
echo ""

# ------------------------------------------------------------
# Read Inputs (GitHub exposes inputs as environment variables)
# ------------------------------------------------------------
FAIL_ON="${INPUT_FAIL_ON}"
ONLY="${INPUT_ONLY}"
EXCEPT="${INPUT_EXCEPT}"
MD_REPORT="${INPUT_MD_REPORT}"
TOKEN="${INPUT_GITHUB_TOKEN}"
PR_COMMENT="${INPUT_PR_COMMENT}"

DIFF="${INPUT_DIFF}"
DIFF_OLD="${INPUT_DIFF_OLD}"
DIFF_NEW="${INPUT_DIFF_NEW}"
DIFF_MD="${INPUT_DIFF_MD_REPORT}"
DIFF_PR_COMMENT="${INPUT_DIFF_PR_COMMENT}"

echo "Inputs:"
echo "  diff: $DIFF"
echo "  diff-old: $DIFF_OLD"
echo "  diff-new: $DIFF_NEW"
echo "  diff-md-report: $DIFF_MD"
echo "  diff-pr-comment: $DIFF_PR_COMMENT"
echo ""

# ------------------------------------------------------------
# Authenticate GH CLI (required for PR comments)
# ------------------------------------------------------------
if [[ "${PR_COMMENT}" == "true" || "${DIFF_PR_COMMENT}" == "true" ]]; then
    if [[ -z "$TOKEN" ]]; then
        echo "❌ GitHub token not provided. PR comments cannot be posted."
        exit 1
    fi

    echo "$TOKEN" | gh auth login --with-token
fi

# ------------------------------------------------------------
# Helper: find existing PR comment matching our unique marker
# ------------------------------------------------------------
find_existing_comment() {
    local pr_number="$1"

    gh api repos/"${GITHUB_REPOSITORY}"/issues/"${pr_number}"/comments \
      --jq '.[] | select(.body | contains("📊 Hewd Diff Report")) | .id' \
      2>/dev/null || true
}

# ------------------------------------------------------------
# Helper: create OR update PR comment
# ------------------------------------------------------------
post_or_update_pr_comment() {
    local pr_number="$1"
    local body_file="$2"

    local existing_id
    existing_id=$(find_existing_comment "$pr_number")

    if [[ -n "$existing_id" ]]; then
        echo "Updating existing PR comment #$existing_id..."
        gh api \
          repos/"${GITHUB_REPOSITORY}"/issues/comments/"${existing_id}" \
          -X PATCH \
          -F body="$(cat "$body_file")"
    else
        echo "Creating new PR comment..."
        gh api \
          repos/"${GITHUB_REPOSITORY}"/issues/"${pr_number}"/comments \
          -f body="$(cat "$body_file")"
    fi
}

# ------------------------------------------------------------
# DIFF MODE
# ------------------------------------------------------------
if [[ "$DIFF" == "true" ]]; then
    echo "== DIFF MODE =="
    echo ""

    # Validate required paths
    if [[ -z "$DIFF_OLD" || -z "$DIFF_NEW" ]]; then
        echo "❌ diff-old and diff-new must both be provided when diff=true"
        exit 1
    fi

    # Create machine-readable diff
    echo "Generating diff.json..."
    hewd diff "$DIFF_OLD" "$DIFF_NEW" --json > diff.json

    # Markdown diff (for PR comments)
    if [[ "$DIFF_MD" == "true" ]]; then
        echo "Generating diff.md..."
        hewd diff "$DIFF_OLD" "$DIFF_NEW" --md > diff.md
    fi

    # Regression gating (strict mode)
    echo "Running regression gating..."
    set +e
    hewd diff "$DIFF_OLD" "$DIFF_NEW" --fail-on-any-regression
    DIFF_EXIT=$?
    set -e

    # PR comment
    if [[ "$DIFF_PR_COMMENT" == "true" && -f diff.md ]]; then
        PR_NUMBER=$(jq -r .pull_request.number "$GITHUB_EVENT_PATH" 2>/dev/null || true)

        if [[ -n "$PR_NUMBER" ]]; then
            echo "Posting (or updating) diff PR comment for PR #$PR_NUMBER..."
            post_or_update_pr_comment "$PR_NUMBER" diff.md
        else
            echo "No PR number found — skipping PR comment."
        fi
    fi

    # Fail CI on regression
    if [[ $DIFF_EXIT -ne 0 ]]; then
        echo "❌ Regression detected. Failing Action."
        exit 1
    fi

    echo "Diff mode completed successfully."
    exit 0
fi

# ------------------------------------------------------------
# DEFAULT MODE: HEWD DOCTOR
# ------------------------------------------------------------
echo "== DOCTOR MODE =="
echo ""

if [[ "$MD_REPORT" == "true" ]]; then
    echo "Generating Markdown doctor report..."
    hewd doctor --md > report.md
else
    echo "Generating plain doctor report..."
    hewd doctor > report.md
fi

# PR comment (doctor mode)
if [[ "$PR_COMMENT" == "true" ]]; then
    PR_NUMBER=$(jq -r .pull_request.number "$GITHUB_EVENT_PATH" 2>/dev/null || true)

    if [[ -n "$PR_NUMBER" ]]; then
        echo "Posting (or updating) doctor PR comment for PR #$PR_NUMBER..."
        post_or_update_pr_comment "$PR_NUMBER" report.md
    else
        echo "No PR number found — skipping PR comment."
    fi
fi

echo "Doctor mode completed."
exit 0