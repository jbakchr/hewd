import * as core from "@actions/core";
import * as github from "@actions/github";
import { exec } from "@actions/exec";
import fs from "fs";

async function postComment(markdown) {
  const token = core.getInput("github-token");
  const octokit = github.getOctokit(token);
  const context = github.context;

  const repo = context.repo;
  const prNumber = context.payload.pull_request.number;
  const marker = "<!-- hewd-report -->";

  // Find previous comment
  const { data: comments } = await octokit.rest.issues.listComments({
    owner: repo.owner,
    repo: repo.repo,
    issue_number: prNumber,
  });

  const existing = comments.find((c) => c.body && c.body.includes(marker));
  const body = `${marker}\n${markdown}`;

  if (existing) {
    await octokit.rest.issues.updateComment({
      owner: repo.owner,
      repo: repo.repo,
      comment_id: existing.id,
      body,
    });
  } else {
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

    await exec("hewd", args, options);

    // Write file
    fs.writeFileSync("hewd-report.md", output);

    // PR Commenting
    const context = github.context;
    if (
      core.getInput("pr-comment") === "true" &&
      context.payload.pull_request
    ) {
      await postComment(output);
    }

    core.setOutput("report", output);
  } catch (err) {
    core.setFailed(err.message);
  }
}

run();
