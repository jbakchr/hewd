const core = require("@actions/core");
const exec = require("@actions/exec");
const github = require("@actions/github");
const fs = require("fs");

async function postComment(markdown) {
  const token = core.getInput("github-token");
  const octokit = github.getOctokit(token);
  const context = github.context;

  const repo = context.repo;
  const prNumber = context.payload.pull_request.number;
  const marker = "<!-- hewd-report -->";

  // Find existing comment
  const { data: comments } = await octokit.rest.issues.listComments({
    owner: repo.owner,
    repo: repo.repo,
    issue_number: prNumber,
  });

  const existing = comments.find((c) => c.body && c.body.includes(marker));
  const body = `${marker}\n${markdown}`;

  if (existing) {
    // Update
    await octokit.rest.issues.updateComment({
      owner: repo.owner,
      repo: repo.repo,
      comment_id: existing.id,
      body,
    });
  } else {
    // Create
    await octokit.rest.issues.createComment({
      owner: repo.owner,
      repo: repo.repo,
      issue_number: prNumber,
      body,
    });
  }
}

async function run() {
  try {
    const failOn = core.getInput("fail-on");
    const only = core.getInput("only");
    const except = core.getInput("except");
    const mdReport = core.getInput("md-report") === "true";

    let args = ["doctor"];

    if (failOn) args.push(`--fail-on=${failOn}`);
    if (only) args.push(`--only=${only}`);
    if (except) args.push(`--except=${except}`);
    if (mdReport) args.push("--md");

    let output = "";
    const options = {
      listeners: {
        stdout: (data) => (output += data.toString()),
      },
    };

    // Run hewd
    await exec.exec("hewd", args, options);

    // Save Markdown output for artifact
    fs.writeFileSync("hewd-report.md", output);

    // ===============================================
    // 🔥 PR COMMENT LOGIC GOES RIGHT HERE
    // ===============================================
    const context = github.context;
    if (
      core.getInput("pr-comment") === "true" &&
      context.payload.pull_request
    ) {
      await postComment(output);
    }
    // ===============================================

    core.setOutput("report", output);
  } catch (err) {
    core.setFailed(err.message);
  }
}

run();
