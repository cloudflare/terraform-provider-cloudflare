package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	syncedLabel = "workflow/synced"
)

type serviceOwner struct {
	owner    string
	teamName string
}

type IssueName struct {
	Name string `json:"name"`
}

type IssueValue struct {
	Value string `json:"value"`
}

type IssueKey struct {
	Key string `json:"key"`
}

type IssueFields struct {
	Project     IssueKey     `json:"project"`
	Summary     string       `json:"summary"`
	Description string       `json:"description"`
	Teams       []IssueValue `json:"customfield_13100"`
	EngOwner    IssueName    `json:"customfield_16304"`
	MyTeam      IssueValue   `json:"customfield_14803"`
	SLA         IssueValue   `json:"customfield_15031"`
	IssueType   IssueName    `json:"issuetype"`
	Components  []IssueName  `json:"components"`
}

type InternalIssue struct {
	Fields IssueFields `json:"fields"`
}

type IssueCreationResponse struct {
	ID   string `json:"id"`
	Key  string `json:"key"`
	Self string `json:"self"`
}

var (
	// Label used to determine if the issue meets all the criteria and is ready
	// to be synced.
	acceptedLabel = "triage/accepted"

	// List of services that we permit syncing internally.
	allowedServiceLabels = []string{
		"provider/internals",

		"service/access",
		"service/cache",
		"service/iam",
		"service/load_balancing",
		"service/logs",
		"service/pages",
		"service/spectrum",
		"service/tls",
		"service/tunnel",
		"service/turnstile",
		"service/workers",
		"service/zones",
	}

	// Mapping of service label to owning internal team.
	serviceOwnership = map[string]serviceOwner{
		"provider/internals": {
			teamName: "API & Zones",
			owner:    "rupalim",
		},
		"service/zones": {
			teamName: "API & Zones",
			owner:    "rupalim",
		},
		"service/access": {
			teamName: "Access",
			owner:    "jroyal",
		},
		"service/logs": {
			teamName: "Logs",
			owner:    "duc",
		},
		"service/tls": {
			teamName: "SSL / TLS",
			owner:    "mihir",
		},
		"service/turnstile": {
			teamName: "Challenges and Turnstile",
			owner:    "opayne",
		},
		"service/workers": {
			teamName: "Workers Core Platform",
			owner:    "laszlo",
		},
		"service/tunnel": {
			teamName: "Tunnel/Teams Routing",
			owner:    "joliveirinha",
		},
		"service/load_balancing": {
			teamName: "Load Balancing",
			owner:    "laurence",
		},
		"service/cache": {
			teamName: "Content Delivery (Cache)",
			owner:    "charwood",
		},
		"service/iam": {
			teamName: "Identity and Access Management",
			owner:    "bnelson",
		},
		"service/spectrum": {
			teamName: "Spectrum",
			owner:    "njones",
		},
		"service/pages": {
			teamName: "Cloudflare Pages",
			owner:    "nrogers",
		},		
	}
)

func main() {
	ctx := context.Background()
	if len(os.Args) < 2 {
		log.Fatalf("Usage: sync-github-issue-to-jira <GitHub issue number>\n")
	}
	iss := os.Args[1]
	issueNumber, err := strconv.Atoi(iss)
	if err != nil {
		log.Fatalf("error parsing issue %q as a number: %s", iss, err)
	}

	githubRepositoryOwner := os.Getenv("GITHUB_OWNER")
	githubRepositoryName := os.Getenv("GITHUB_REPO")
	githubAccessToken := os.Getenv("GITHUB_TOKEN")
	jiraHostname := os.Getenv("JIRA_HOSTNAME")
	jiraAuthToken := os.Getenv("JIRA_AUTH_TOKEN")
	accessClientID := os.Getenv("CF_ACCESS_CLIENT_ID")
	accessClientSecret := os.Getenv("CF_ACCESS_CLIENT_SECRET")

	if githubRepositoryOwner == "" {
		log.Fatal("GITHUB_OWNER not set")
	}

	if githubRepositoryName == "" {
		log.Fatal("GITHUB_REPO not set")
	}

	if githubAccessToken == "" {
		log.Fatal("GITHUB_TOKEN not set")
	}

	if jiraHostname == "" {
		log.Fatal("JIRA_HOSTNAME not set")
	}

	if jiraAuthToken == "" {
		log.Fatal("JIRA_AUTH_TOKEN not set")
	}

	if accessClientID == "" {
		log.Fatal("CF_ACCESS_CLIENT_ID not set")
	}

	if accessClientSecret == "" {
		log.Fatal("CF_ACCESS_CLIENT_SECRET not set")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	issue, _, err := client.Issues.Get(ctx, githubRepositoryOwner, githubRepositoryName, issueNumber)
	if err != nil {
		log.Fatalf("error retrieving issue %s/%s#%d: %s", githubRepositoryOwner, githubRepositoryName, issueNumber, err)
	}

	if hasLabel(issue, syncedLabel) {
		log.Printf("issue is already marked as synced (%s), skipping", syncedLabel)
		os.Exit(0)
	}

	if !hasLabel(issue, acceptedLabel) {
		log.Printf("issue is not marked as ready for syncing using %s, skipping", acceptedLabel)
		os.Exit(0)
	}

	serviceLabel := getOwnershipLabel(issue)
	if serviceLabel == "" {
		fmt.Println("no service owner detected; exiting without creating a new JIRA issue")
		os.Exit(0)
	}

	service := serviceOwnership[serviceLabel]

	newIssue := InternalIssue{Fields: IssueFields{
		Project:     IssueKey{Key: "CUSTESC"},
		Summary:     *issue.Title,
		Description: jirafyBodyMarkdown(issue),
		Teams:       []IssueValue{{Value: service.teamName}},
		EngOwner:    IssueName{Name: service.owner},
		SLA:         IssueValue{Value: "Pro / Free"},
		MyTeam:      IssueValue{Value: "Other"},
		IssueType:   IssueName{Name: "Bug"},
		Components:  []IssueName{{Name: "SDK & Client API Libraries"}},
	}}

	res, err := json.Marshal(newIssue)
	if err != nil {
		fmt.Println(err)
	}

	url := fmt.Sprintf("https://%s/rest/api/latest/issue/", jiraHostname)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(res))
	if err != nil {
		log.Fatalf("failed to build HTTP request: %s", err)
	}

	req.Header.Set("authorization", "Basic "+jiraAuthToken)
	req.Header.Set("cf-access-client-id", accessClientID)
	req.Header.Set("cf-access-client-secret", accessClientSecret)
	req.Header.Set("content-type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %s", err)
	}

	var createdIssue IssueCreationResponse
	json.Unmarshal([]byte(body), &createdIssue)

	if resp.StatusCode != http.StatusCreated {
		fmt.Println(fmt.Sprintf("failed to create new JIRA issue"))
		os.Exit(1)
	}

	fmt.Println(fmt.Sprintf("successfully created internal JIRA issue: %s", createdIssue.Key))
	_, _, err = client.Issues.AddLabelsToIssue(ctx, githubRepositoryOwner, githubRepositoryName, issueNumber, []string{syncedLabel})
	if err != nil {
		log.Printf("error adding synced label for issue %s/%s#%d: %s", githubRepositoryOwner, githubRepositoryName, issueNumber, err)
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

func getOwnershipLabel(issue *github.Issue) string {
	for _, l := range issue.Labels {
		if strings.Contains(*l.Name, "service/") || strings.Contains(*l.Name, "provider/internals") {
			return *l.Name
		}
	}

	return ""
}

// jirafyBodyMarkdown takes GitHub markdown and makes it palatable for JIRA
// with reasonable formatting.
func jirafyBodyMarkdown(issue *github.Issue) string {
	output := "GitHub issue: " + *issue.HTMLURL + "\n\n---\n\n"

	output += *issue.Body
	output = strings.ReplaceAll(output, "- [X] ", "✅ ")
	output = strings.ReplaceAll(output, "###", "h3.")
	output = strings.ReplaceAll(output, "```hcl", "{code}")
	output = strings.ReplaceAll(output, "```", "{code}")

	return output
}
