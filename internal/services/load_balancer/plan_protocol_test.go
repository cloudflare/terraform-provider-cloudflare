package load_balancer_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"reflect"
	"sort"
	"strings"
	"sync"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	providerschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	frameworkresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

// Runs the real Terraform plan and apply protocol against a local API, so it
// needs neither Cloudflare credentials nor live resources.
type loadBalancerTestProvider struct{ baseURL string }

func (p *loadBalancerTestProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cloudflare"
}
func (p *loadBalancerTestProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = providerschema.Schema{}
}
func (p *loadBalancerTestProvider) Configure(_ context.Context, _ provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	resp.ResourceData = cloudflare.NewClient(option.WithBaseURL(p.baseURL), option.WithAPIToken("test-token"), option.WithMaxRetries(0))
}
func (p *loadBalancerTestProvider) Resources(context.Context) []func() frameworkresource.Resource {
	return []func() frameworkresource.Resource{load_balancer.NewResource}
}
func (p *loadBalancerTestProvider) DataSources(context.Context) []func() datasource.DataSource {
	return nil
}

// Writes replace the load balancer and fill omitted fields with server defaults.
// ticketDefaults selects the nullable defaults observed on the reported resource;
// the other fixture retains the simpler non-null defaults and geo map behavior.
type loadBalancerAPI struct {
	server *httptest.Server
	mu     sync.Mutex
	live   map[string]any
}

type loadBalancerDefaults int

const (
	standardLoadBalancerDefaults loadBalancerDefaults = iota
	ticketLoadBalancerDefaults
)

func fakeLoadBalancerAPI(t *testing.T, defaults loadBalancerDefaults) *loadBalancerAPI {
	api := &loadBalancerAPI{}
	writes := 0
	api.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.mu.Lock()
		defer api.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		const collection = "/zones/0da42c8d2132a9ddaf714f9e7c920711/load_balancers"
		switch {
		case r.Method == http.MethodPost && r.URL.Path == collection,
			r.Method == http.MethodPut && r.URL.Path == collection+"/lb-1":
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Errorf("decode request body: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			writes++
			affinityAttributes := map[string]any{
				"drain_duration": float64(0), "headers": nil, "require_all_headers": false,
				"samesite": "Auto", "secure": "Auto", "zero_downtime_failover": "none",
			}
			lb := map[string]any{
				"description":                 "",
				"ttl":                         float64(30),
				"enabled":                     true,
				"proxied":                     false,
				"session_affinity":            "none",
				"session_affinity_ttl":        nil,
				"steering_policy":             "",
				"adaptive_routing":            map[string]any{"failover_across_pools": false},
				"location_strategy":           map[string]any{"mode": "pop", "prefer_ecs": "proximity"},
				"random_steering":             map[string]any{"default_weight": float64(1), "pool_weights": nil},
				"country_pools":               map[string]any{},
				"pop_pools":                   map[string]any{},
				"region_pools":                map[string]any{},
				"networks":                    []any{"cloudflare"},
				"rules":                       []any{},
				"session_affinity_attributes": affinityAttributes,
			}
			if defaults == ticketLoadBalancerDefaults {
				lb["description"] = nil
				lb["ttl"] = nil
				lb["country_pools"] = nil
				affinityAttributes["require_all_headers"] = nil
			}
			for k, v := range body {
				if v == nil {
					continue
				}
				// A partial top-level nested object gets its own API defaults.
				if k == "adaptive_routing" || k == "location_strategy" || k == "random_steering" || k == "session_affinity_attributes" {
					defaults, ok := lb[k].(map[string]any)
					if !ok {
						t.Errorf("invalid default for %s: %#v", k, lb[k])
						w.WriteHeader(http.StatusInternalServerError)
						return
					}
					supplied, ok := v.(map[string]any)
					if !ok {
						t.Errorf("invalid request %s: expected object, got %#v", k, v)
						w.WriteHeader(http.StatusBadRequest)
						return
					}
					for field, value := range supplied {
						if value != nil {
							defaults[field] = value
						}
					}
					continue
				}
				lb[k] = v
			}
			if lb["session_affinity_ttl"] == nil {
				switch lb["session_affinity"] {
				case "cookie", "ip_cookie":
					lb["session_affinity_ttl"] = float64(82800)
				case "header":
					lb["session_affinity_ttl"] = float64(1800)
				}
			}
			if lb["session_affinity"] == "header" {
				ttl, ok := lb["session_affinity_ttl"].(float64)
				if !ok {
					t.Errorf("invalid session_affinity_ttl: expected number, got %#v", lb["session_affinity_ttl"])
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				if ttl < 30 || ttl > 3600 {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, `{"success":false,"errors":[{"code":1000,"message":"session_affinity_ttl %v out of range for header"}],"messages":[],"result":null}`, ttl)
					return
				}
			}
			// The API reports unspecified rule overrides as null, not inherited
			// load-balancer defaults or values from the previous rule at this index.
			rules, ok := lb["rules"].([]any)
			if !ok {
				t.Errorf("invalid request rules: expected array, got %#v", lb["rules"])
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			for _, raw := range rules {
				rule, ok := raw.(map[string]any)
				if !ok {
					t.Errorf("invalid request rule: expected object, got %#v", raw)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				if rule["overrides"] != nil {
					overrides, ok := rule["overrides"].(map[string]any)
					if !ok {
						t.Errorf("invalid request rule overrides: expected object, got %#v", rule["overrides"])
						w.WriteHeader(http.StatusBadRequest)
						return
					}
					for _, key := range []string{"adaptive_routing", "country_pools", "default_pools", "fallback_pool", "location_strategy", "pop_pools", "random_steering", "region_pools", "session_affinity", "session_affinity_attributes", "session_affinity_ttl", "steering_policy", "ttl"} {
						if _, present := overrides[key]; !present {
							overrides[key] = nil
						}
					}
				}
			}
			lb["id"] = "lb-1"
			lb["zone_name"] = "example.com"
			if api.live == nil {
				lb["created_on"] = "2026-11-03T10:00:00Z"
			} else {
				lb["created_on"] = api.live["created_on"]
			}
			lb["modified_on"] = fmt.Sprintf("2026-11-03T10:00:%02dZ", writes)
			api.live = lb
		case r.Method == http.MethodGet && r.URL.Path == collection+"/lb-1":
		case r.Method == http.MethodDelete && r.URL.Path == collection+"/lb-1":
			api.live = nil
			fmt.Fprint(w, `{"success":true,"errors":[],"messages":[],"result":{"id":"lb-1"}}`)
			return
		default:
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if api.live == nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, `{"success":false,"errors":[{"code":1000,"message":"not found"}],"messages":[],"result":null}`)
			return
		}
		if err := json.NewEncoder(w).Encode(map[string]any{"success": true, "errors": []any{}, "messages": []any{}, "result": api.live}); err != nil {
			t.Errorf("encode API response: %v", err)
		}
	}))
	return api
}

func loadBalancerConfig(fallbackPool, sessionAffinity string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer" "test" {
  zone_id          = "0da42c8d2132a9ddaf714f9e7c920711"
  name             = "lb.example.com"
  fallback_pool    = %[1]q
  default_pools    = [%[1]q]
  session_affinity = %[2]q
}
`, fallbackPool, sessionAffinity)
}

// expectOnlyUnknown fails unless the planned update leaves exactly the given
// attributes "(known after apply)".
type expectOnlyUnknown []string

func (want expectOnlyUnknown) CheckPlan(_ context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	for _, rc := range req.Plan.ResourceChanges {
		if rc.Address != "cloudflare_load_balancer.test" {
			continue
		}
		if !rc.Change.Actions.Update() {
			resp.Error = fmt.Errorf("want an in-place update, got %v", rc.Change.Actions)
			return
		}
		var got []string
		collectUnknown("", rc.Change.AfterUnknown, &got)
		sort.Strings(got)
		w := append([]string{}, want...)
		sort.Strings(w)
		if strings.Join(got, ",") != strings.Join(w, ",") {
			resp.Error = fmt.Errorf("planned as unknown: %v, want only %v", got, w)
		}
		return
	}
	resp.Error = fmt.Errorf("no planned change for cloudflare_load_balancer.test")
}

func collectUnknown(prefix string, v any, out *[]string) {
	switch v := v.(type) {
	case bool:
		if v {
			*out = append(*out, prefix)
		}
	case map[string]any:
		for k, inner := range v {
			p := k
			if prefix != "" {
				p = prefix + "." + k
			}
			collectUnknown(p, inner, out)
		}
	case []any:
		for i, inner := range v {
			collectUnknown(fmt.Sprintf("%s.%d", prefix, i), inner, out)
		}
	}
}

// Non-null API defaults remain covered; an unproxied DNS TTL cannot safely
// be reused because the API's default is not documented.
func TestLoadBalancerPoolChangePreservesSafeDefaults(t *testing.T) {
	terraformPath, err := exec.LookPath("terraform")
	if err != nil {
		t.Skip("Terraform CLI is required for the local protocol regression test")
	}
	t.Setenv("TF_ACC_TERRAFORM_PATH", terraformPath)
	t.Setenv("TF_ACC_PROVIDER_NAMESPACE", "cloudflare")

	api := fakeLoadBalancerAPI(t, standardLoadBalancerDefaults)
	defer api.server.Close()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"cloudflare": providerserver.NewProtocol6WithError(&loadBalancerTestProvider{baseURL: api.server.URL}),
		},
		Steps: []resource.TestStep{
			{
				Config: loadBalancerConfig("pool-a", "cookie"),
				Check:  resource.TestCheckResourceAttr("cloudflare_load_balancer.test", "session_affinity_ttl", "82800"),
			},
			{
				Config: loadBalancerConfig("pool-b", "cookie"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectOnlyUnknown{"modified_on", "ttl"}},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_load_balancer.test", "fallback_pool", "pool-b"),
					resource.TestCheckResourceAttr("cloudflare_load_balancer.test", "session_affinity_ttl", "82800"),
				),
			},
			{Config: loadBalancerConfig("pool-b", "cookie"), PlanOnly: true},
			{
				// The cookie TTL kept from state would be out of range for
				// header affinity, so the TTL and attributes stay unknown.
				Config: loadBalancerConfig("pool-b", "header"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectOnlyUnknown{
						"modified_on",
						"ttl",
						"session_affinity_ttl",
						"session_affinity_attributes",
					}},
				},
				Check: resource.TestCheckResourceAttr("cloudflare_load_balancer.test", "session_affinity_ttl", "1800"),
			},
			{Config: loadBalancerConfig("pool-b", "header"), PlanOnly: true},
		},
	})
}

// checkAPI compares the selected complete API values, including null versus
// empty collections. Rules are projected to their identity and full overrides.
func (api *loadBalancerAPI) checkAPI(want map[string]any) resource.TestCheckFunc {
	return func(_ *terraform.State) error {
		api.mu.Lock()
		defer api.mu.Unlock()
		for key, expected := range want {
			got, ok := api.live[key]
			if !ok {
				return fmt.Errorf("API response missing %q", key)
			}
			if key == "rules" {
				rules, ok := got.([]any)
				if !ok {
					return fmt.Errorf("API rules have unexpected type: %#v", got)
				}
				projected := make([]any, 0, len(rules))
				for _, raw := range rules {
					rule, ok := raw.(map[string]any)
					if !ok {
						return fmt.Errorf("API rule has unexpected type: %#v", raw)
					}
					projected = append(projected, map[string]any{"name": rule["name"], "overrides": rule["overrides"]})
				}
				got = projected
			}
			if !reflect.DeepEqual(got, expected) {
				return fmt.Errorf("API %s = %#v, want %#v", key, got, expected)
			}
		}
		return nil
	}
}

func protocolConfig(pool, attributes string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer" "test" {
  zone_id       = "0da42c8d2132a9ddaf714f9e7c920711"
  name          = "lb.example.com"
  fallback_pool = %[1]q
  default_pools = [%[1]q]
  %[2]s
}
`, pool, attributes)
}

func protocolRule(name, pool, extra string) string {
	return fmt.Sprintf(`{
      name = %[1]q
      condition = "dns.qry.type == 28"
      priority = 1
      overrides = {
        default_pools = [%[2]q]
        fallback_pool = %[2]q
        %[3]s
      }
    }`, name, pool, extra)
}

func protocolRules(rules ...string) string {
	return "rules = [" + strings.Join(rules, ",") + "]"
}

func expectedRule(name, pool string, ttl any, extra ...map[string]any) any {
	overrides := map[string]any{
		"adaptive_routing": nil, "country_pools": nil,
		"default_pools": []any{pool}, "fallback_pool": pool,
		"location_strategy": nil, "pop_pools": nil, "random_steering": nil,
		"region_pools": nil, "session_affinity": nil, "session_affinity_attributes": nil,
		"session_affinity_ttl": nil, "steering_policy": nil, "ttl": ttl,
	}
	for _, fields := range extra {
		for key, value := range fields {
			overrides[key] = value
		}
	}
	return map[string]any{"name": name, "overrides": overrides}
}

func checkRuleState(index int, name, pool string, ttl *string) resource.TestCheckFunc {
	const address = "cloudflare_load_balancer.test"
	prefix := fmt.Sprintf("rules.%d", index)
	checks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(address, prefix+".name", name),
		resource.TestCheckResourceAttr(address, prefix+".overrides.default_pools.#", "1"),
		resource.TestCheckResourceAttr(address, prefix+".overrides.default_pools.0", pool),
		resource.TestCheckResourceAttr(address, prefix+".overrides.fallback_pool", pool),
	}
	for _, key := range []string{"adaptive_routing", "country_pools", "location_strategy", "pop_pools", "random_steering", "region_pools", "session_affinity", "session_affinity_attributes", "session_affinity_ttl", "steering_policy"} {
		checks = append(checks, resource.TestCheckNoResourceAttr(address, prefix+".overrides."+key))
	}
	if ttl == nil {
		checks = append(checks, resource.TestCheckNoResourceAttr(address, prefix+".overrides.ttl"))
	} else {
		checks = append(checks, resource.TestCheckResourceAttr(address, prefix+".overrides.ttl", *ttl))
	}
	return resource.ComposeTestCheckFunc(checks...)
}

func TestLoadBalancerTicketDefaultsAndRuleReorder(t *testing.T) {
	terraformPath, err := exec.LookPath("terraform")
	if err != nil {
		t.Skip("Terraform CLI is required for the local protocol regression test")
	}
	t.Setenv("TF_ACC_TERRAFORM_PATH", terraformPath)
	t.Setenv("TF_ACC_PROVIDER_NAMESPACE", "cloudflare")
	api := fakeLoadBalancerAPI(t, ticketLoadBalancerDefaults)
	defer api.server.Close()

	green := protocolRule("Green", "green-pool", "")
	blue := protocolRule("Blue", "blue-pool", "")
	initial := protocolConfig("pool-a", "proxied = true\n"+protocolRules(green, blue))
	updated := protocolConfig("pool-b", "proxied = true\n"+protocolRules(green, blue))
	withTTL := protocolConfig("pool-b", "proxied = true\n"+protocolRules(protocolRule("Green", "green-pool", "ttl = 120"), blue))
	reordered := protocolConfig("pool-b", "proxied = true\n"+protocolRules(blue, protocolRule("Green", "green-new", "")))
	rules := []any{expectedRule("Green", "green-pool", nil), expectedRule("Blue", "blue-pool", nil)}
	const address = "cloudflare_load_balancer.test"
	ticketState := resource.ComposeTestCheckFunc(
		resource.TestCheckNoResourceAttr(address, "description"),
		resource.TestCheckNoResourceAttr(address, "ttl"),
		resource.TestCheckNoResourceAttr(address, "country_pools.%"),
		resource.TestCheckNoResourceAttr(address, "session_affinity_ttl"),
		resource.TestCheckResourceAttr(address, "pop_pools.%", "0"),
		resource.TestCheckResourceAttr(address, "region_pools.%", "0"),
		resource.TestCheckResourceAttr(address, "adaptive_routing.failover_across_pools", "false"),
		resource.TestCheckResourceAttr(address, "location_strategy.mode", "pop"),
		resource.TestCheckResourceAttr(address, "location_strategy.prefer_ecs", "proximity"),
		resource.TestCheckResourceAttr(address, "random_steering.default_weight", "1"),
		resource.TestCheckNoResourceAttr(address, "random_steering.pool_weights.%"),
		resource.TestCheckResourceAttr(address, "networks.#", "1"),
		resource.TestCheckResourceAttr(address, "networks.0", "cloudflare"),
		resource.TestCheckResourceAttr(address, "session_affinity_attributes.drain_duration", "0"),
		resource.TestCheckNoResourceAttr(address, "session_affinity_attributes.headers.#"),
		resource.TestCheckNoResourceAttr(address, "session_affinity_attributes.require_all_headers"),
		resource.TestCheckResourceAttr(address, "session_affinity_attributes.samesite", "Auto"),
		resource.TestCheckResourceAttr(address, "session_affinity_attributes.secure", "Auto"),
		resource.TestCheckResourceAttr(address, "session_affinity_attributes.zero_downtime_failover", "none"),
		resource.TestCheckResourceAttr(address, "rules.#", "2"),
	)
	apiDefaults := map[string]any{
		"description": nil, "ttl": nil, "country_pools": nil, "session_affinity_ttl": nil,
		"pop_pools": map[string]any{}, "region_pools": map[string]any{},
		"adaptive_routing":            map[string]any{"failover_across_pools": false},
		"location_strategy":           map[string]any{"mode": "pop", "prefer_ecs": "proximity"},
		"random_steering":             map[string]any{"default_weight": float64(1), "pool_weights": nil},
		"networks":                    []any{"cloudflare"},
		"session_affinity_attributes": map[string]any{"drain_duration": float64(0), "headers": nil, "require_all_headers": nil, "samesite": "Auto", "secure": "Auto", "zero_downtime_failover": "none"},
		"rules":                       rules,
		"fallback_pool":               "pool-b", "default_pools": []any{"pool-b"},
	}
	var nullState []statecheck.StateCheck
	for _, path := range []tfjsonpath.Path{
		tfjsonpath.New("description"), tfjsonpath.New("ttl"), tfjsonpath.New("country_pools"),
		tfjsonpath.New("session_affinity_ttl"), tfjsonpath.New("random_steering").AtMapKey("pool_weights"),
		tfjsonpath.New("session_affinity_attributes").AtMapKey("headers"),
		tfjsonpath.New("session_affinity_attributes").AtMapKey("require_all_headers"),
	} {
		nullState = append(nullState, statecheck.ExpectKnownValue(address, path, knownvalue.Null()))
	}
	for i := range 2 {
		for _, key := range []string{"adaptive_routing", "country_pools", "location_strategy", "pop_pools", "random_steering", "region_pools", "session_affinity", "session_affinity_attributes", "session_affinity_ttl", "steering_policy", "ttl"} {
			path := tfjsonpath.New("rules").AtSliceIndex(i).AtMapKey("overrides").AtMapKey(key)
			nullState = append(nullState, statecheck.ExpectKnownValue(address, path, knownvalue.Null()))
		}
	}
	ttl := "120"
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"cloudflare": providerserver.NewProtocol6WithError(&loadBalancerTestProvider{baseURL: api.server.URL}),
		},
		Steps: []resource.TestStep{
			{Config: initial, ConfigStateChecks: nullState, Check: resource.ComposeTestCheckFunc(ticketState, checkRuleState(0, "Green", "green-pool", nil), checkRuleState(1, "Blue", "blue-pool", nil))},
			{
				Config:            updated,
				ConfigStateChecks: nullState,
				ConfigPlanChecks:  resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{expectOnlyUnknown{"modified_on"}}},
				Check: resource.ComposeTestCheckFunc(ticketState, checkRuleState(0, "Green", "green-pool", nil), checkRuleState(1, "Blue", "blue-pool", nil),
					resource.TestCheckResourceAttr(address, "fallback_pool", "pool-b"), resource.TestCheckResourceAttr(address, "default_pools.0", "pool-b"), api.checkAPI(apiDefaults)),
			},
			{Config: updated, PlanOnly: true, ConfigPlanChecks: resource.ConfigPlanChecks{PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()}}},
			{Config: withTTL, Check: resource.ComposeTestCheckFunc(checkRuleState(0, "Green", "green-pool", &ttl), checkRuleState(1, "Blue", "blue-pool", nil), api.checkAPI(map[string]any{"rules": []any{expectedRule("Green", "green-pool", float64(120)), expectedRule("Blue", "blue-pool", nil)}}))},
			{Config: reordered, ConfigStateChecks: nullState, Check: resource.ComposeTestCheckFunc(ticketState, checkRuleState(0, "Blue", "blue-pool", nil), checkRuleState(1, "Green", "green-new", nil), api.checkAPI(map[string]any{"rules": []any{expectedRule("Blue", "blue-pool", nil), expectedRule("Green", "green-new", nil)}}))},
			{Config: reordered, PlanOnly: true, ConfigPlanChecks: resource.ConfigPlanChecks{PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()}}},
		},
	})
}

// Replacing a configuration must clear removed values rather than resending
// optional+computed values copied from the previous state.
func TestLoadBalancerRemovedConfigurationResetsAPIDefaults(t *testing.T) {
	terraformPath, err := exec.LookPath("terraform")
	if err != nil {
		t.Skip("Terraform CLI is required for the local protocol regression test")
	}
	t.Setenv("TF_ACC_TERRAFORM_PATH", terraformPath)
	t.Setenv("TF_ACC_PROVIDER_NAMESPACE", "cloudflare")
	api := fakeLoadBalancerAPI(t, standardLoadBalancerDefaults)
	defer api.server.Close()

	initial := protocolConfig("pool-a", `
  description = "configured"
  ttl = 120
  proxied = true
  steering_policy = "geo"
  pop_pools = { SJC = ["pop-pool"] }
  country_pools = { US = ["country-pool"] }
  region_pools = { WNAM = ["old-region"] }
  adaptive_routing = { failover_across_pools = true }
  location_strategy = { mode = "resolver_ip", prefer_ecs = "always" }
  random_steering = { default_weight = 0.5, pool_weights = { "pool-a" = 0.4 } }
  session_affinity = "header"
  session_affinity_ttl = 120
  session_affinity_attributes = {
    drain_duration = 30
    headers = ["x-session"]
    require_all_headers = true
    samesite = "Strict"
    secure = "Always"
    zero_downtime_failover = "temporary"
  }
  `+protocolRules(protocolRule("Green", "green-pool", "")))
	// Trigger an update so the framework marks omitted children as unknown.
	// Omission alone retains Optional+Computed state before modifiers run.
	childrenRemoved := strings.NewReplacer(
		`fallback_pool = "pool-a"`, `fallback_pool = "pool-b"`,
		`default_pools = ["pool-a"]`, `default_pools = ["pool-b"]`,
		"drain_duration = 30", "",
		"require_all_headers = true", "",
		`samesite = "Strict"`, "",
		`secure = "Always"`, "",
		`zero_downtime_failover = "temporary"`, "",
	).Replace(initial)
	regionOnly := protocolConfig("pool-a", `
  proxied = true
  steering_policy = "geo"
  region_pools = { ENAM = ["new-region"] }
  session_affinity = "header"
`)
	removed := protocolConfig("pool-a", `proxied = true
  steering_policy = "geo"`)
	const address = "cloudflare_load_balancer.test"
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"cloudflare": providerserver.NewProtocol6WithError(&loadBalancerTestProvider{baseURL: api.server.URL}),
		},
		Steps: []resource.TestStep{
			{
				Config: initial,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(address, "description", "configured"),
					resource.TestCheckResourceAttr(address, "ttl", "120"),
					resource.TestCheckResourceAttr(address, "pop_pools.SJC.0", "pop-pool"),
					resource.TestCheckResourceAttr(address, "country_pools.US.0", "country-pool"),
					resource.TestCheckResourceAttr(address, "region_pools.WNAM.0", "old-region"),
					resource.TestCheckResourceAttr(address, "adaptive_routing.failover_across_pools", "true"),
					resource.TestCheckResourceAttr(address, "location_strategy.mode", "resolver_ip"),
					resource.TestCheckResourceAttr(address, "random_steering.default_weight", "0.5"),
					resource.TestCheckResourceAttr(address, "random_steering.pool_weights.pool-a", "0.4"),
					resource.TestCheckResourceAttr(address, "session_affinity_ttl", "120"),
					resource.TestCheckResourceAttr(address, "session_affinity_attributes.headers.0", "x-session"),
					resource.TestCheckResourceAttr(address, "session_affinity_attributes.require_all_headers", "true"),
					checkRuleState(0, "Green", "green-pool", nil),
				),
			},
			{
				Config: childrenRemoved,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(address, "session_affinity_attributes.headers.0", "x-session"),
					resource.TestCheckResourceAttr(address, "session_affinity_attributes.drain_duration", "0"),
					resource.TestCheckResourceAttr(address, "session_affinity_attributes.require_all_headers", "false"),
					resource.TestCheckResourceAttr(address, "session_affinity_attributes.samesite", "Auto"),
					resource.TestCheckResourceAttr(address, "session_affinity_attributes.secure", "Auto"),
					resource.TestCheckResourceAttr(address, "session_affinity_attributes.zero_downtime_failover", "none"),
					api.checkAPI(map[string]any{
						"session_affinity_attributes": map[string]any{
							"headers": []any{"x-session"}, "drain_duration": float64(0),
							"require_all_headers": false, "samesite": "Auto", "secure": "Auto",
							"zero_downtime_failover": "none",
						},
					}),
				),
			},
			{Config: childrenRemoved, PlanOnly: true, ConfigPlanChecks: resource.ConfigPlanChecks{
				PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
			}},
			{
				Config: regionOnly,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(address, "description", ""),
					resource.TestCheckResourceAttr(address, "ttl", "30"),
					resource.TestCheckResourceAttr(address, "pop_pools.%", "0"),
					resource.TestCheckResourceAttr(address, "country_pools.%", "0"),
					resource.TestCheckResourceAttr(address, "region_pools.%", "1"),
					resource.TestCheckResourceAttr(address, "region_pools.ENAM.0", "new-region"),
					resource.TestCheckResourceAttr(address, "rules.#", "0"),
					resource.TestCheckResourceAttr(address, "adaptive_routing.failover_across_pools", "false"),
					resource.TestCheckResourceAttr(address, "location_strategy.mode", "pop"),
					resource.TestCheckResourceAttr(address, "location_strategy.prefer_ecs", "proximity"),
					resource.TestCheckResourceAttr(address, "random_steering.default_weight", "1"),
					resource.TestCheckNoResourceAttr(address, "random_steering.pool_weights.%"),
					resource.TestCheckResourceAttr(address, "session_affinity_ttl", "1800"),
					resource.TestCheckResourceAttr(address, "session_affinity_attributes.drain_duration", "0"),
					resource.TestCheckNoResourceAttr(address, "session_affinity_attributes.headers.#"),
					resource.TestCheckResourceAttr(address, "session_affinity_attributes.require_all_headers", "false"),
					resource.TestCheckResourceAttr(address, "session_affinity_attributes.samesite", "Auto"),
					resource.TestCheckResourceAttr(address, "session_affinity_attributes.secure", "Auto"),
					resource.TestCheckResourceAttr(address, "session_affinity_attributes.zero_downtime_failover", "none"),
					api.checkAPI(map[string]any{
						"description": "", "ttl": float64(30), "session_affinity": "header", "session_affinity_ttl": float64(1800),
						"pop_pools": map[string]any{}, "country_pools": map[string]any{}, "region_pools": map[string]any{"ENAM": []any{"new-region"}},
						"rules": []any{}, "adaptive_routing": map[string]any{"failover_across_pools": false},
						"location_strategy":           map[string]any{"mode": "pop", "prefer_ecs": "proximity"},
						"random_steering":             map[string]any{"default_weight": float64(1), "pool_weights": nil},
						"session_affinity_attributes": map[string]any{"drain_duration": float64(0), "headers": nil, "require_all_headers": false, "samesite": "Auto", "secure": "Auto", "zero_downtime_failover": "none"},
					}),
				),
			},
			{Config: regionOnly, PlanOnly: true, ConfigPlanChecks: resource.ConfigPlanChecks{PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()}}},
			{
				Config: removed,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(address, "region_pools.%", "0"),
					resource.TestCheckResourceAttr(address, "session_affinity", "none"),
					resource.TestCheckNoResourceAttr(address, "session_affinity_ttl"),
					api.checkAPI(map[string]any{"region_pools": map[string]any{}, "session_affinity": "none", "session_affinity_ttl": nil}),
				),
			},
			{Config: removed, PlanOnly: true, ConfigPlanChecks: resource.ConfigPlanChecks{PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()}}},
		},
	})
}

func TestLoadBalancerUnknownConfigAndDestroy(t *testing.T) {
	terraformPath, err := exec.LookPath("terraform")
	if err != nil {
		t.Skip("Terraform CLI is required for the local protocol regression test")
	}
	t.Setenv("TF_ACC_TERRAFORM_PATH", terraformPath)
	t.Setenv("TF_ACC_PROVIDER_NAMESPACE", "cloudflare")
	api := fakeLoadBalancerAPI(t, standardLoadBalancerDefaults)
	defer api.server.Close()

	initial := protocolConfig("pool-a", "")
	unknown := protocolConfig("pool-a", `description = terraform_data.description.output`) + `
resource "terraform_data" "description" {
  input = "from-data"
}
`
	const address = "cloudflare_load_balancer.test"
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"cloudflare": providerserver.NewProtocol6WithError(&loadBalancerTestProvider{baseURL: api.server.URL}),
		},
		Steps: []resource.TestStep{
			{Config: initial, Check: resource.TestCheckResourceAttr(address, "description", "")},
			{
				Config: unknown,
				ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{
					plancheck.ExpectUnknownValue(address, tfjsonpath.New("description")),
				}},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(address, "description", "from-data"),
					api.checkAPI(map[string]any{"description": "from-data"}),
				),
			},
			{Config: unknown, PlanOnly: true, ConfigPlanChecks: resource.ConfigPlanChecks{PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()}}},
			{
				Config:  unknown,
				Destroy: true,
				ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(address, plancheck.ResourceActionDestroy),
				}},
				Check: func(_ *terraform.State) error {
					api.mu.Lock()
					defer api.mu.Unlock()
					if api.live != nil {
						return fmt.Errorf("load balancer still exists after destroy: %#v", api.live)
					}
					return nil
				},
			},
		},
	})
}

func TestLoadBalancerRuleAffinityRemoval(t *testing.T) {
	terraformPath, err := exec.LookPath("terraform")
	if err != nil {
		t.Skip("Terraform CLI is required for the local protocol regression test")
	}
	t.Setenv("TF_ACC_TERRAFORM_PATH", terraformPath)
	t.Setenv("TF_ACC_PROVIDER_NAMESPACE", "cloudflare")
	api := fakeLoadBalancerAPI(t, ticketLoadBalancerDefaults)
	defer api.server.Close()

	blue := protocolRule("Blue", "blue-pool", "")
	greenHeader := protocolRule("Green", "green-pool", `
	  session_affinity = "header"
	  session_affinity_attributes = {
	    headers = ["x-rule"]
	    drain_duration = 30
	  }`)
	greenCookie := protocolRule("Green", "green-pool", `
	  session_affinity = "cookie"
	  session_affinity_attributes = { drain_duration = 30 }`)
	initial := protocolConfig("pool-a", "proxied = true\n"+protocolRules(greenHeader, blue))
	reordered := protocolConfig("pool-a", "proxied = true\n"+protocolRules(blue, greenCookie))
	updated := protocolConfig("pool-b", "proxied = true\n"+protocolRules(blue, greenCookie))
	const address = "cloudflare_load_balancer.test"
	const prefix = "rules.1.overrides.session_affinity_attributes"
	apiRules := []any{
		expectedRule("Blue", "blue-pool", nil),
		expectedRule("Green", "green-pool", nil, map[string]any{
			"session_affinity": "cookie",
			"session_affinity_attributes": map[string]any{
				"drain_duration": float64(30), "require_all_headers": false,
				"samesite": "Auto", "secure": "Auto", "zero_downtime_failover": "none",
			},
		}),
	}
	check := resource.ComposeTestCheckFunc(
		checkRuleState(0, "Blue", "blue-pool", nil),
		resource.TestCheckResourceAttr(address, "rules.1.name", "Green"),
		resource.TestCheckResourceAttr(address, "rules.1.overrides.session_affinity", "cookie"),
		resource.TestCheckNoResourceAttr(address, prefix+".headers.#"),
		resource.TestCheckResourceAttr(address, prefix+".drain_duration", "30"),
		api.checkAPI(map[string]any{"rules": apiRules}),
	)
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"cloudflare": providerserver.NewProtocol6WithError(&loadBalancerTestProvider{baseURL: api.server.URL}),
		},
		Steps: []resource.TestStep{
			{
				Config: initial,
				Check: resource.TestCheckResourceAttr(
					address, "rules.0.overrides.session_affinity_attributes.headers.0", "x-rule",
				),
			},
			{Config: reordered, Check: check},
			{
				Config: updated,
				ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{
					plancheck.ExpectKnownValue(address, tfjsonpath.New("rules").AtSliceIndex(1).
						AtMapKey("overrides").AtMapKey("session_affinity_attributes").
						AtMapKey("headers"), knownvalue.Null()),
				}},
				Check: check,
			},
			{Config: updated, PlanOnly: true, ConfigPlanChecks: resource.ConfigPlanChecks{
				PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
			}},
		},
	})
}
