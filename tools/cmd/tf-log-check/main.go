package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
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
		log.Printf("not a %s report, skipping", bugLabel)
		os.Exit(0)
	}

	var re = regexp.MustCompile(`### Link to debug output\n\n(?P<link>.*)\n\n### Panic output`)
	matches := re.FindStringSubmatch(*issue.Body)
	link := matches[1]

	if strings.Contains(link, "gist.github.com") {
		link = link + "/raw"
	}

	res, err := http.Get(link)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	if res.StatusCode != 200 {
		log.Fatalf("failed to fetch remote link: %s, status %d", link, res.StatusCode)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("could not ready response: %s", err)
	}
	body := string(resBody)

	if !strings.Contains(body, "Terraform version") &&
		!strings.Contains(body, "Go runtime version") &&
		!strings.Contains(body, "CLI args") &&
		!strings.Contains(body, "created provider logger") &&
		!strings.Contains(body, "CLI command args") &&
		!strings.Contains(body, "provider: plugin process exited") {
		_, _, err := client.Issues.CreateComment(ctx, owner, repo, issueNumber, &github.IssueComment{
			Body: &missingLogFileBody,
		})
		if err != nil {
			log.Fatalf("failed to create comment for issue %s/%s#%d: %s", owner, repo, issueNumber, err)
		}

		_, err = client.Issues.RemoveLabelForIssue(ctx, owner, repo, issueNumber, needsTriageLabel)
		if err != nil {
			log.Fatalf("error removing label for issue %s/%s#%d: %s", owner, repo, issueNumber, err)
		}
		_, _, err = client.Issues.AddLabelsToIssue(ctx, owner, repo, issueNumber, []string{needsInformationLabel})
		if err != nil {
			log.Fatalf("error adding label for issue %s/%s#%d: %s", owner, repo, issueNumber, err)
		}
	}

	os.Exit(0)
}

func isBugReport(issue *github.Issue) bool {
	for _, label := range issue.Labels {
		if *label.Name == bugLabel {
			return true
		}
	}

	return false
}
