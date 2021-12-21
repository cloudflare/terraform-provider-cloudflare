package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	changelogEntryFileFormat      = ".changelog/%d.txt"
	skipLabel                     = "workflow/skip-changelog-entry"
	changelogProcessDocumentation = "https://github.com/cloudflare/terraform-provider-cloudflare/blob/master/docs/changelog-process.md"
)

var (
	changelogEntryPresent = false
)

func main() {
	ctx := context.Background()
	if len(os.Args) < 2 {
		log.Fatalf("Usage: changelog-check PR#\n")
	}
	pr := os.Args[1]
	prNo, err := strconv.Atoi(pr)
	if err != nil {
		log.Fatalf("error parsing PR %q as a number: %s", pr, err)
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

	pullRequest, _, err := client.PullRequests.Get(ctx, owner, repo, prNo)
	if err != nil {
		log.Fatalf("error retrieving pull request %s/%s#%d: %s", owner, repo, prNo, err)
	}

	for _, label := range pullRequest.Labels {
		if label.GetName() == skipLabel {
			log.Printf("%s label found, exiting as changelog is not required\n", skipLabel)
			os.Exit(0)
		}
	}

	files, _, _ := client.PullRequests.ListFiles(ctx, owner, repo, prNo, &github.ListOptions{})
	if err != nil {
		log.Fatalf("error retrieving files on pull request %s/%s#%d: %s", owner, repo, prNo, err)
	}

	for _, file := range files {
		if file.GetFilename() == fmt.Sprintf(changelogEntryFileFormat, prNo) {
			changelogEntryPresent = true
		}
	}

	if changelogEntryPresent {
		log.Printf("changelog found for %d, skipping remainder of checks\n", prNo)
		os.Exit(0)
	}

	body := "Oops! It looks like no changelog entry is attached to" +
		" this PR. Please include a release note as described in " +
		changelogProcessDocumentation + ".\n\nExample: " +
		"\n\n~~~\n```release-note:TYPE\nRelease note" +
		"\n```\n~~~\n\n" +
		"If you do not require a release note to be included, please add the `workflow/skip-changelog-entry` label."

	_, _, err = client.Issues.CreateComment(ctx, owner, repo, prNo, &github.IssueComment{
		Body: &body,
	})

	os.Exit(1)
}
