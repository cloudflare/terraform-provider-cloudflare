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

	"github.com/cloudflare/cloudflare-go"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	needsTriageLabel      = "needs-triage"
	needsInformationLabel = "triage/needs-information"
	debugLogDetected      = "triage/debug-log-attached"
	bugLabel              = "kind/bug"
	missingLogFileBody    = "Thank you for reporting this issue! For maintainers to dig into issues it is required that all issues include the entirety of `TF_LOG=DEBUG` output to be provided. The only parts that should be redacted are your user credentials in the `X-Auth-Key`, `X-Auth-Email` and `Authorization` HTTP headers. Details such as zone or account identifiers are not considered sensitive but can be redacted if you are very cautious. This log file provides additional context from Terraform, the provider and the Cloudflare API that helps in debugging issues. Without it, maintainers are very limited in what they can do and may hamper diagnosis efforts.\n\nThis issue has been marked with `triage/needs-information` and is unlikely to receive maintainer attention until the log file is provided making this a complete bug report."
	debugLogProvidedMsg   = "Terraform debug log detected :white_check_mark:"
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

	if !hasLabel(issue, bugLabel) {
		log.Printf("not a %s report, skipping", bugLabel)
		os.Exit(0)
	}

	comments, _, _ := client.Issues.ListComments(ctx, owner, repo, issueNumber, &github.IssueListCommentsOptions{})
	for _, comment := range comments {
		if strings.Contains(comment.GetBody(), "This issue has been marked with `triage/needs-information`") {
			if hasLabel(issue, debugLogDetected) {
				client.Issues.EditComment(ctx, owner, repo, *comment.ID, &github.IssueComment{
					Body: cloudflare.StringPtr(debugLogProvidedMsg),
				})
				os.Exit(0)
			}
		}
	}

	var re = regexp.MustCompile(`### Link to debug output\s+(?P<link>.*)\s+### Panic output`)
	matches := re.FindStringSubmatch(*issue.Body)

	if len(matches) == 0 {
		postMissingLogPayload(ctx, client, owner, repo, issueNumber)
		os.Exit(0)
	}

	link := strings.TrimSpace(matches[1])
	if strings.Contains(link, "gist.github.com") {
		link = link + "/raw"
	}

	res, err := http.Get(link)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		postMissingLogPayload(ctx, client, owner, repo, issueNumber)
		os.Exit(1)
	}

	if res.StatusCode != 200 {
		log.Printf("failed to fetch remote link: %s, status %d", link, res.StatusCode)
		postMissingLogPayload(ctx, client, owner, repo, issueNumber)
		os.Exit(1)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("could not ready response: %s", err)
		postMissingLogPayload(ctx, client, owner, repo, issueNumber)
		os.Exit(1)
	}
	body := string(resBody)

	if !strings.Contains(body, "Terraform version") &&
		!strings.Contains(body, "Go runtime version") &&
		!strings.Contains(body, "CLI args") &&
		!strings.Contains(body, "created provider logger") &&
		!strings.Contains(body, "CLI command args") &&
		!strings.Contains(body, "provider: plugin process exited") {

		postMissingLogPayload(ctx, client, owner, repo, issueNumber)
	} else {
		if len(comments) == 0 {
			client.Issues.CreateComment(ctx, owner, repo, issueNumber, &github.IssueComment{
				Body: cloudflare.StringPtr(debugLogProvidedMsg),
			})
		} else {
			for _, comment := range comments {
				if strings.Contains(comment.GetBody(), "This issue has been marked with `triage/needs-information`") {
					client.Issues.EditComment(ctx, owner, repo, *comment.ID, &github.IssueComment{
						Body: cloudflare.StringPtr(debugLogProvidedMsg),
					})
				}
			}
		}

		_, err = client.Issues.RemoveLabelForIssue(ctx, owner, repo, issueNumber, needsInformationLabel)
		if err != nil {
			log.Printf("error removing label for issue %s/%s#%d: %s", owner, repo, issueNumber, err)
		}

		_, _, err = client.Issues.AddLabelsToIssue(ctx, owner, repo, issueNumber, []string{debugLogDetected})
		if err != nil {
			log.Printf("error adding label for issue %s/%s#%d: %s", owner, repo, issueNumber, err)
		}
	}

	os.Exit(0)
}

func hasLabel(issue *github.Issue, label string) bool {
	for _, l := range issue.Labels {
		if *l.Name == label {
			return true
		}
	}

	return false
}

func postMissingLogPayload(ctx context.Context, client *github.Client, owner, repo string, issueNumber int) {
	comments, _, _ := client.Issues.ListComments(ctx, owner, repo, issueNumber, &github.IssueListCommentsOptions{})
	for _, comment := range comments {
		if strings.Contains(comment.GetBody(), "This issue has been marked with `triage/needs-information`") {
			log.Printf("request for debug log already exists, exiting")
			os.Exit(0)
		}
	}

	_, _, err := client.Issues.CreateComment(ctx, owner, repo, issueNumber, &github.IssueComment{
		Body: &missingLogFileBody,
	})
	if err != nil {
		log.Printf("failed to create comment for issue %s/%s#%d: %s", owner, repo, issueNumber, err)
	}

	_, _, err = client.Issues.AddLabelsToIssue(ctx, owner, repo, issueNumber, []string{needsInformationLabel})
	if err != nil {
		log.Printf("error adding label for issue %s/%s#%d: %s", owner, repo, issueNumber, err)
	}

	_, err = client.Issues.RemoveLabelForIssue(ctx, owner, repo, issueNumber, needsTriageLabel)
	if err != nil {
		log.Printf("error removing label for issue %s/%s#%d: %s", owner, repo, issueNumber, err)
	}
}
