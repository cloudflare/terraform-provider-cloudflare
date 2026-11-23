package sdkv2provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"strings"
	"sync"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestSuppressRedactedNotificationWebhookURLDiff(t *testing.T) {
	tests := map[string]struct {
		old, new string
		want     bool
	}{
		"slack path":          {"https://hooks.slack.com/...", "https://hooks.slack.com/services/T0/B0/XXXX", true},
		"query":               {"https://chat.googleapis.com/...", "https://chat.googleapis.com/v1/spaces/X/messages?key=k&token=t", true},
		"port kept":           {"https://example.com:8443/...", "https://example.com:8443/hook?api_key=secret", true},
		"userinfo and frag":   {"https://example.com/...", "https://user:pass@example.com/a/b#frag", true},
		"bare host":           {"https://example.com/...", "https://example.com", true},
		"http scheme":         {"http://example.com/...", "http://example.com/hook", true},
		"unchanged url":       {"https://example.com/hook", "https://example.com/hook", true},
		"host change":         {"https://example.com/...", "https://other.example.com/hook", false},
		"scheme change":       {"https://example.com/...", "http://example.com/hook", false},
		"port change":         {"https://example.com:8443/...", "https://example.com/hook", false},
		"full url differs":    {"https://example.com/a", "https://example.com/b", false},
		"unparseable config":  {"https://example.com/...", "not a url", false},
		"empty config":        {"https://example.com/...", "", false},
		"create from nothing": {"", "https://example.com/hook", false},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := suppressRedactedNotificationWebhookURLDiff("url", tc.old, tc.new, nil); got != tc.want {
				t.Fatalf("suppress(%q, %q) = %v, want %v", tc.old, tc.new, got, tc.want)
			}
		})
	}
}

type fakeWebhooksAPI struct {
	mu       sync.Mutex
	name     string
	url      string
	putURLs  []string
	creates  int
	deletes  int
	redacted bool
}

const fakeWebhookID = "11111111-2222-3333-4444-555555555555"

func (f *fakeWebhooksAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	var body struct{ Name, URL string }
	_ = json.NewDecoder(r.Body).Decode(&body)

	result := map[string]any{"id": fakeWebhookID}
	switch r.Method {
	case http.MethodPost:
		f.creates++
		f.name, f.url = body.Name, body.URL
	case http.MethodPut:
		f.name = body.Name
		f.putURLs = append(f.putURLs, body.URL)
	case http.MethodDelete:
		f.deletes++
	case http.MethodGet:
		shown := f.url
		if f.redacted {
			shown = "https://hooks.slack.com/..."
		}
		result = map[string]any{
			"id": fakeWebhookID, "name": f.name, "url": shown, "type": "slack",
			"created_at": "2026-01-01T00:00:00Z",
		}
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"success": true, "errors": []any{}, "messages": []any{}, "result": result})
}

func TestNotificationPolicyWebhooks_RedactedURLFromAPI(t *testing.T) {
	t.Setenv("TF_ACC_TERRAFORM_PATH", terraformPathForTest(t))

	fake := &fakeWebhooksAPI{redacted: true}
	srv := httptest.NewServer(fake)
	defer srv.Close()

	factories := map[string]func() (*schema.Provider, error){
		"cloudflare": func() (*schema.Provider, error) { //nolint:unparam // resource.TestCase.ProviderFactories requires this signature
			p := New("dev")()
			p.ConfigureContextFunc = func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
				api, err := cloudflare.NewWithAPIToken("test-token", cloudflare.BaseURL(srv.URL))
				if err != nil {
					return nil, diag.FromErr(err)
				}
				return api, nil
			}
			return p, nil
		},
	}

	config := func(name, url string) string {
		return fmt.Sprintf(`
resource "cloudflare_notification_policy_webhooks" "test" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = %q
  url        = %q
}`, name, url)
	}
	const slackURL = "https://hooks.slack.com/services/T0/B0/XXXX"

	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: factories,
		Steps: []resource.TestStep{
			{
				Config: config("hook", slackURL),
				Check:  resource.TestCheckResourceAttr("cloudflare_notification_policy_webhooks.test", "url", "https://hooks.slack.com/..."),
			},
			{
				Config:   config("hook", slackURL),
				PlanOnly: true,
			},
			{
				Config: config("renamed", slackURL),
				Check: func(*terraform.State) error {
					fake.mu.Lock()
					defer fake.mu.Unlock()
					if fake.creates != 1 || fake.deletes != 0 {
						return fmt.Errorf("rename replaced the webhook: creates=%d deletes=%d", fake.creates, fake.deletes)
					}
					if len(fake.putURLs) != 1 || fake.putURLs[0] != slackURL {
						return fmt.Errorf("update did not send the configured url")
					}
					return nil
				},
			},
			{
				Config:             config("renamed", "https://other.example.com/hook"),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func terraformPathForTest(t *testing.T) string {
	t.Helper()
	path, err := exec.LookPath("terraform")
	if err != nil {
		t.Skip("terraform binary not found")
	}
	return strings.TrimSpace(path)
}
