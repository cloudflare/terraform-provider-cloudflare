package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var maintainers = []string{"dependabot", "jacobbednarz", "patryk"}

func main() {
	ctx := context.Background()
	if len(os.Args) < 2 {
		log.Fatalf("Usage: maintainer-only-file-check PR#\n")
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

	// Don't worry about continuing if the PR was opened by a maintainer.
	if contains(maintainers, *pullRequest.User.Login) {
		os.Exit(0)
	}

	files, _, _ := client.PullRequests.ListFiles(ctx, owner, repo, prNo, &github.ListOptions{})
	if err != nil {
		log.Fatalf("error retrieving files on pull request %s/%s#%d: %s", owner, repo, prNo, err)
	}

	for _, file := range files {
		if file.GetFilename() == "go.mod" || file.GetFilename() == "go.sum" {
			body := `
			This project handles dependency version bumps (including upstream changes from cloudflare-go) independently of the standard PR process using automation. This allows the dependency upgrades to land without causing merge conflicts in multiple branches and handled in a consistent way. The exception to this is security related dependency upgrades but they should be co-ordinated with the maintainer team privately.

			Please remove the changes to the go.mod or go.sum files from this PR in order to proceed with review and merging.
			`

			_, _, _ = client.Issues.CreateComment(ctx, owner, repo, prNo, &github.IssueComment{
				Body: &body,
			})
			os.Exit(1)
		}

		if file.GetFilename() == "CHANGELOG.md" {
			body := `
			This pull request contains a CHANGELOG.md file which should only be modified by maintainers.

			If you are looking to include a CHANGELOG entry, you should use the process documented at https://github.com/cloudflare/terraform-provider-cloudflare/blob/master/docs/changelog-process.md instead.

			In order for this pull request to be merged, you need remove the modifications to CHANGELOG.md.
			`

			_, _, _ = client.Issues.CreateComment(ctx, owner, repo, prNo, &github.IssueComment{
				Body: &body,
			})
			os.Exit(1)
		}
	}

	os.Exit(0)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
