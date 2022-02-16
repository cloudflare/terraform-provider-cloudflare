package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	needsTriageLabel      = "needs-triage"
	needsInformationLabel = "triage/needs-information"
	bugLabel              = "kind/bug"
	missingLogFileBody    = "Thank you for reporting this issue! For maintainers to dig into issues it is required that all issues include the entirety of `TF_LOG=DEBUG` output to be provided. The only parts that should be redacted are your user credentials in the `X-Auth-Key`, `X-Auth-Email` and `Authorization` HTTP headers. Details such as zone or account identifiers are not considered sensitive but can be redacted if you are very cautious. This log file provides additional context from Terraform, the provider and the Cloudflare API that helps in debugging issues. Without it, maintainers are very limited in what they can do and may hamper diagnosis efforts.\n\nThis issue has been marked with `triage/needs-information` and is unlikely to receive maintainer attention until the log file is provided making this a complete bug report."
)

func main() {
	ctx := context.Background()
	if len(os.Args) < 2 {
		log.Fatalf("Usage: tf-log-check issue#\n")
	}
	issueParam := os.Args[1]
	issueNumber, err := strconv.Atoi(issueParam)
	if err != nil {
		log.Fatalf("error parsing issue %q as a number: %s", issueParam, err)
	}

	owner := os.Getenv("GITHUB_OWNER")
	repo := os.Getenv("GITHUB_REPO")
	token := os.Getenv("GITHUB_TOKEN")

	if owner == "" {
		log.Fatalf("GITHUB_OWNER not set")
	}

	if repo == "" {
		log.Fatalf("GITHUB_REPO not set")
	}

	if token == "" {
		log.Fatalf("GITHUB_TOKEN not set")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	issue, _, err := client.Issues.Get(ctx, owner, repo, issueNumber)
	if err != nil {
		log.Fatalf("error retrieving issue %s/%s#%d: %s", owner, repo, issueNumber, err)
	}

	if !isBugReport(issue) {
		os.Exit(0)
	}

	if !strings.Contains(*issue.Body, "Terraform version") &&
		!strings.Contains(*issue.Body, "Go runtime version") &&
		!strings.Contains(*issue.Body, "CLI args") &&
		!strings.Contains(*issue.Body, "created provider logger") &&
		!strings.Contains(*issue.Body, "CLI command args") &&
		!strings.Contains(*issue.Body, "provider: plugin process exited") {

		client.Issues.CreateComment(ctx, owner, repo, issueNumber, &github.IssueComment{
			Body: &missingLogFileBody,
		})

		client.Issues.RemoveLabelForIssue(ctx, owner, repo, issueNumber, needsTriageLabel)
		client.Issues.AddLabelsToIssue(ctx, owner, repo, issueNumber, []string{needsInformationLabel})
	}

	os.Exit(0)
}

func isBugReport(issue *github.Issue) bool {
	for _, label := range issue.Labels {
		if label.Name == &bugLabel {
			return true
		}
	}

	return false
}
