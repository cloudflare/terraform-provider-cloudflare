package ai_gateway_dynamic_routing_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/ai_gateway_dynamic_routing"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	providerschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	frameworkresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

type dynamicRoutingTestProvider struct{}

type dynamicRoutingTestProviderModel struct {
	BaseURL types.String `tfsdk:"base_url"`
}

func (dynamicRoutingTestProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cloudflare"
	resp.Version = "test"
}

func (dynamicRoutingTestProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = providerschema.Schema{
		Attributes: map[string]providerschema.Attribute{
			"base_url": providerschema.StringAttribute{Optional: true},
		},
	}
}

func (dynamicRoutingTestProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data dynamicRoutingTestProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := cloudflare.NewClient(option.WithBaseURL(data.BaseURL.ValueString()))
	resp.ResourceData = client
	resp.DataSourceData = client
}

func (dynamicRoutingTestProvider) Resources(context.Context) []func() frameworkresource.Resource {
	return []func() frameworkresource.Resource{ai_gateway_dynamic_routing.NewResource}
}

func (dynamicRoutingTestProvider) DataSources(context.Context) []func() datasource.DataSource {
	return nil
}

type dynamicRoutingFakeServer struct {
	mu      sync.RWMutex
	graph   string
	exists  bool
	deletes int
}

func (s *dynamicRoutingFakeServer) setGraph(graph string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.graph = graph
}

func (s *dynamicRoutingFakeServer) graphJSON() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.graph
}

func (s *dynamicRoutingFakeServer) restoreGraph(graph string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.graph = graph
	s.exists = true
}

func (s *dynamicRoutingFakeServer) wasDeleted() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.deletes > 0 && !s.exists
}

func (s *dynamicRoutingFakeServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	routePath := "/accounts/account-1/ai-gateway/gateways/gateway-1/routes"
	itemPath := routePath + "/route-1"
	if r.URL.Path != routePath && r.URL.Path != itemPath {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodPost:
		if r.URL.Path != routePath {
			http.NotFound(w, r)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var request map[string]json.RawMessage
		if err := json.Unmarshal(body, &request); err != nil || request["elements"] == nil {
			http.Error(w, "invalid create request", http.StatusBadRequest)
			return
		}
		var elements []json.RawMessage
		if err := json.Unmarshal(request["elements"], &elements); err != nil || len(elements) != 2 {
			http.Error(w, "unexpected graph", http.StatusBadRequest)
			return
		}
		s.mu.Lock()
		s.graph = string(request["elements"])
		s.exists = true
		s.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, `{"result":{"id":"route-1","name":"route","gateway_id":"gateway-1","elements":%s,"version":{"active":"true","data":"documented","version_id":"version-1"}},"success":true}`, s.graphJSON())
	case http.MethodGet:
		if r.URL.Path != itemPath {
			http.NotFound(w, r)
			return
		}
		s.mu.RLock()
		exists := s.exists
		s.mu.RUnlock()
		if !exists {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, `{"result":{"id":"route-1","name":"route","gateway_id":"gateway-1","version":{"active":"true","data":%s,"version_id":"version-1"}},"success":true}`, s.graphJSON())
	case http.MethodDelete:
		if r.URL.Path != itemPath {
			http.NotFound(w, r)
			return
		}
		s.mu.Lock()
		s.exists = false
		s.deletes++
		s.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"result":{"id":"route-1"},"success":true}`))
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func TestAIGatewayDynamicRoutingProtocol(t *testing.T) {
	graphA := `[{"id":"start","outputs":{"next":{"elementId":"end"}},"type":"start"},{"id":"end","outputs":{},"type":"end"}]`
	graphB := `[{"id":"changed","outputs":{"next":{"elementId":"other"}},"type":"start"},{"id":"other","outputs":{},"type":"end"}]`

	fake := &dynamicRoutingFakeServer{graph: graphA}
	server := httptest.NewServer(fake)
	defer server.Close()

	config := func(startID, nextID, endID string) string {
		return fmt.Sprintf(`provider "cloudflare" {
  base_url = %q
}

resource "cloudflare_ai_gateway_dynamic_routing" "test" {
  account_id = "account-1"
  gateway_id = "gateway-1"
  name       = "route"

  elements = [
    {
      id   = %q
      type = "start"
      outputs = {
        next = { element_id = %q }
      }
    },
    {
      id      = %q
      type    = "end"
      outputs = {}
    },
  ]
}
`, server.URL, startID, nextID, endID)
	}

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"cloudflare": func() (tfprotov6.ProviderServer, error) {
				return providerserver.NewProtocol6(dynamicRoutingTestProvider{})(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: config("start", "end", "end"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				PreConfig: func() {
					fake.setGraph(graphB)
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
				RefreshPlanChecks: resource.RefreshPlanChecks{
					PostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("cloudflare_ai_gateway_dynamic_routing.test", plancheck.ResourceActionReplace),
					},
				},
			},
			{
				Config:  config("changed", "other", "other"),
				Destroy: true,
			},
			{
				PreConfig: func() {
					fake.restoreGraph(graphB)
				},
				Config:             config("changed", "other", "other"),
				ResourceName:       "cloudflare_ai_gateway_dynamic_routing.test",
				ImportState:        true,
				ImportStateId:      "account-1/gateway-1/route-1",
				ImportStatePersist: true,
				ImportStateCheck: func(states []*terraform.InstanceState) error {
					if len(states) != 1 {
						return fmt.Errorf("expected one imported resource, got %d", len(states))
					}
					attributes := states[0].Attributes
					for path, expected := range map[string]string{
						"id":                                 "route-1",
						"elements.#":                         "2",
						"elements.0.id":                      "changed",
						"elements.0.outputs.next.element_id": "other",
						"elements.1.id":                      "other",
						"version.data":                       graphB,
					} {
						if got := attributes[path]; got != expected {
							return fmt.Errorf("%s = %q, want %q", path, got, expected)
						}
					}
					return nil
				},
			},
			{
				Config:             config("changed", "other", "other"),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})

	if !fake.wasDeleted() {
		t.Error("expected the test lifecycle to delete the fake route")
	}
}
