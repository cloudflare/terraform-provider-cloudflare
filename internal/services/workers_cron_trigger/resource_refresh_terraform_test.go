package workers_cron_trigger_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_cron_trigger"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	providerschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	frameworkresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestWorkersCronTriggerTerraformDetectsDrift(t *testing.T) {
	var mu sync.Mutex
	schedules := []map[string]string{}
	createdOn := ""
	updates := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		if req.URL.Path != "/accounts/account-123/workers/scripts/test-worker/schedules" {
			t.Errorf("unexpected path: %s", req.URL.Path)
		}
		switch req.Method {
		case http.MethodPut:
			if err := json.NewDecoder(req.Body).Decode(&schedules); err != nil {
				t.Error(err)
			}
			updates++
			createdOn = fmt.Sprintf("2026-11-01T00:00:%02dZ", updates)
		case http.MethodGet:
		default:
			t.Errorf("unexpected method: %s", req.Method)
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		responseSchedules := []map[string]string{}
		for _, schedule := range schedules {
			responseSchedules = append(responseSchedules, map[string]string{"cron": schedule["cron"], "created_on": createdOn, "modified_on": "2026-11-02T00:00:00Z"})
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]any{"success": true, "errors": []any{}, "messages": []any{}, "result": map[string]any{"schedules": responseSchedules}}); err != nil {
			t.Error(err)
		}
	}))
	defer server.Close()

	const name = "cloudflare_workers_cron_trigger.test"
	const config = `resource "cloudflare_workers_cron_trigger" "test" {
  account_id = "account-123"
  script_name = "test-worker"
  schedules = [{ cron = "0 0 * * *" }]
 }`
	steps := []resource.TestStep{
		{Config: config, Check: resource.TestCheckResourceAttr(name, "schedules.0.cron", "0 0 * * *")},
		{Config: config, PlanOnly: true, ConfigPlanChecks: resource.ConfigPlanChecks{PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()}}},
	}
	for _, drift := range []string{"0 12 * * *", ""} {
		drift := drift
		steps = append(steps,
			resource.TestStep{
				Config: config, PlanOnly: true, ExpectNonEmptyPlan: true,
				PreConfig: func() {
					mu.Lock()
					defer mu.Unlock()
					schedules = []map[string]string{}
					if drift != "" {
						schedules = append(schedules, map[string]string{"cron": drift})
					}
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate)}},
			},
			resource.TestStep{Config: config, Check: resource.TestCheckResourceAttr(name, "schedules.0.cron", "0 0 * * *")},
			resource.TestStep{Config: config, PlanOnly: true, ConfigPlanChecks: resource.ConfigPlanChecks{PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()}}},
		)
	}
	steps = append(steps, resource.TestStep{ResourceName: name, ImportState: true, ImportStateId: "account-123/test-worker", ImportStateVerify: true})
	// Updating configuration must still preserve the planned cron and populate timestamps.
	steps = append(steps, resource.TestStep{Config: strings.ReplaceAll(config, "0 0 * * *", "0 6 * * *"), Check: resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr(name, "schedules.0.cron", "0 6 * * *"),
		resource.TestCheckResourceAttrSet(name, "schedules.0.created_on"),
	)})
	resource.UnitTest(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"cloudflare": providerserver.NewProtocol6WithError(&cronTestProvider{baseURL: server.URL})}, Steps: steps})
}

// Only the resource under test is served. Local tests use a local HTTP server;
// the opt-in acceptance test supplies its live API client explicitly.
type cronTestProvider struct {
	baseURL string
	client  *cloudflare.Client
}

func (p *cronTestProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cloudflare"
}
func (p *cronTestProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = providerschema.Schema{}
}
func (p *cronTestProvider) Configure(_ context.Context, _ provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	if p.client != nil {
		resp.ResourceData = p.client
		return
	}
	resp.ResourceData = cloudflare.NewClient(option.WithBaseURL(p.baseURL), option.WithAPIToken("test-token"), option.WithMaxRetries(0))
}
func (p *cronTestProvider) Resources(context.Context) []func() frameworkresource.Resource {
	return []func() frameworkresource.Resource{workers_cron_trigger.NewResource}
}
func (p *cronTestProvider) DataSources(context.Context) []func() datasource.DataSource { return nil }
