package account_dns_settings_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/require"
)

const foundationDNSTestAccountID = "0123456789abcdef0123456789abcdef"

// fakeAccountDNSSettingsAPI mimics the account DNS settings API after
// zone_defaults.foundation_dns was removed: it is never returned, and a PATCH
// that includes it is rejected.
type fakeAccountDNSSettingsAPI struct {
	mu             sync.Mutex
	nameserverType string
	nsTTL          float64
	patches        []map[string]any
}

func newFakeAccountDNSSettingsAPI(t *testing.T, nameserverType string) *fakeAccountDNSSettingsAPI {
	t.Helper()
	api := &fakeAccountDNSSettingsAPI{nameserverType: nameserverType, nsTTL: 86400}
	server := httptest.NewServer(http.HandlerFunc(api.serveHTTP))
	t.Cleanup(server.Close)
	t.Setenv("CLOUDFLARE_BASE_URL", server.URL)
	t.Setenv("CLOUDFLARE_API_TOKEN", strings.Repeat("x", 40))
	t.Setenv("CLOUDFLARE_API_KEY", "")
	t.Setenv("CLOUDFLARE_EMAIL", "")
	return api
}

func (a *fakeAccountDNSSettingsAPI) patchBodies() []map[string]any {
	a.mu.Lock()
	defer a.mu.Unlock()
	return append([]map[string]any(nil), a.patches...)
}

func (a *fakeAccountDNSSettingsAPI) serveHTTP(w http.ResponseWriter, r *http.Request) {
	a.mu.Lock()
	defer a.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")

	if r.URL.Path != "/accounts/"+foundationDNSTestAccountID+"/dns_settings" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"success":false,"result":null,"errors":[{"code":7003,"message":"No route for the URI"}],"messages":[]}`)
		return
	}

	if r.Method == http.MethodPatch {
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `{"success":false,"result":null,"errors":[{"code":1000,"message":"invalid body"}],"messages":[]}`)
			return
		}
		a.patches = append(a.patches, body)
		zoneDefaults, _ := body["zone_defaults"].(map[string]any)
		if _, ok := zoneDefaults["foundation_dns"]; ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `{"success":false,"result":null,"errors":[{"code":1034,"message":"foundation_dns is no longer supported. Use nameservers.type to configure Advanced Nameservers."}],"messages":[]}`)
			return
		}
		if nameservers, ok := zoneDefaults["nameservers"].(map[string]any); ok {
			if nameserverType, ok := nameservers["type"].(string); ok {
				a.nameserverType = nameserverType
			}
		}
		if nsTTL, ok := zoneDefaults["ns_ttl"].(float64); ok {
			a.nsTTL = nsTTL
		}
	}

	result := map[string]any{
		"enforce_dns_only": false,
		"zone_defaults": map[string]any{
			"flatten_all_cnames":  false,
			"multi_provider":      false,
			"ns_ttl":              a.nsTTL,
			"secondary_overrides": false,
			"zone_mode":           "standard",
			"internal_dns":        map[string]any{"reference_zone_id": nil},
			"nameservers":         map[string]any{"type": a.nameserverType},
			"soa": map[string]any{
				"mname":   nil,
				"rname":   "dns.cloudflare.com",
				"refresh": 10000,
				"retry":   2400,
				"expire":  604800,
				"min_ttl": 1800,
				"ttl":     3600,
			},
		},
	}
	_ = json.NewEncoder(w).Encode(map[string]any{"success": true, "errors": []any{}, "messages": []any{}, "result": result})
}

func configuredProviderServer(t *testing.T) tfprotov6.ProviderServer {
	t.Helper()
	ctx := context.Background()
	server, err := acctest.TestAccProtoV6ProviderFactories["cloudflare"]()
	require.NoError(t, err)

	schemaResp, err := server.GetProviderSchema(ctx, &tfprotov6.GetProviderSchemaRequest{})
	require.NoError(t, err)
	require.Empty(t, schemaResp.Diagnostics)
	providerType := schemaResp.Provider.ValueType().(tftypes.Object)
	nulls := make(map[string]tftypes.Value, len(providerType.AttributeTypes))
	for name, attrType := range providerType.AttributeTypes {
		nulls[name] = tftypes.NewValue(attrType, nil)
	}
	config, err := tfprotov6.NewDynamicValue(providerType, tftypes.NewValue(providerType, nulls))
	require.NoError(t, err)

	configureResp, err := server.ConfigureProvider(ctx, &tfprotov6.ConfigureProviderRequest{TerraformVersion: "1.9.0", Config: &config})
	require.NoError(t, err)
	require.Empty(t, configureResp.Diagnostics)
	return server
}

func zoneDefaultsAttributes(t *testing.T, server tfprotov6.ProviderServer, state *tfprotov6.DynamicValue) map[string]tftypes.Value {
	t.Helper()
	providerSchema, err := server.GetProviderSchema(context.Background(), &tfprotov6.GetProviderSchemaRequest{})
	require.NoError(t, err)
	value, err := state.Unmarshal(providerSchema.ResourceSchemas["cloudflare_account_dns_settings"].ValueType())
	require.NoError(t, err)
	var attributes map[string]tftypes.Value
	require.NoError(t, value.As(&attributes))
	var zoneDefaults map[string]tftypes.Value
	require.NoError(t, attributes["zone_defaults"].As(&zoneDefaults))
	return zoneDefaults
}

// State written while zone_defaults.foundation_dns was in the schema must
// upgrade and refresh without error, and the attribute must be gone afterwards.
func TestAccountDNSSettingsStateWithFoundationDNSUpgradesAndRefreshes(t *testing.T) {
	newFakeAccountDNSSettingsAPI(t, "cloudflare.standard")
	server := configuredProviderServer(t)
	ctx := context.Background()

	priorState, err := json.Marshal(map[string]any{
		"account_id":       foundationDNSTestAccountID,
		"enforce_dns_only": false,
		"zone_defaults": map[string]any{
			"flatten_all_cnames":  false,
			"foundation_dns":      false,
			"multi_provider":      false,
			"ns_ttl":              86400,
			"secondary_overrides": false,
			"zone_mode":           "standard",
			"internal_dns":        map[string]any{"reference_zone_id": nil},
			"nameservers":         map[string]any{"type": "cloudflare.standard"},
			"soa": map[string]any{
				"expire":  604800,
				"min_ttl": 1800,
				"mname":   nil,
				"refresh": 10000,
				"retry":   2400,
				"rname":   "dns.cloudflare.com",
				"ttl":     3600,
			},
		},
	})
	require.NoError(t, err)

	for _, version := range []int64{0, 500} {
		t.Run(fmt.Sprintf("version_%d", version), func(t *testing.T) {
			upgradeResp, err := server.UpgradeResourceState(ctx, &tfprotov6.UpgradeResourceStateRequest{
				TypeName: "cloudflare_account_dns_settings",
				Version:  version,
				RawState: &tfprotov6.RawState{JSON: priorState},
			})
			require.NoError(t, err)
			require.Empty(t, upgradeResp.Diagnostics)
			upgraded := zoneDefaultsAttributes(t, server, upgradeResp.UpgradedState)
			_, ok := upgraded["foundation_dns"]
			require.False(t, ok, "upgraded state still has foundation_dns")
			require.True(t, upgraded["ns_ttl"].Equal(tftypes.NewValue(tftypes.Number, 86400)))

			readResp, err := server.ReadResource(ctx, &tfprotov6.ReadResourceRequest{
				TypeName:     "cloudflare_account_dns_settings",
				CurrentState: upgradeResp.UpgradedState,
			})
			require.NoError(t, err)
			require.Empty(t, readResp.Diagnostics)
			refreshed := zoneDefaultsAttributes(t, server, readResp.NewState)
			_, ok = refreshed["foundation_dns"]
			require.False(t, ok, "refreshed state still has foundation_dns")
			var nameservers map[string]tftypes.Value
			require.NoError(t, refreshed["nameservers"].As(&nameservers))
			require.True(t, nameservers["type"].Equal(tftypes.NewValue(tftypes.String, "cloudflare.standard")))
		})
	}
}

func accountDNSSettingsConfig(nsTTL int, extra string) string {
	return fmt.Sprintf(`
resource "cloudflare_account_dns_settings" "test" {
  account_id       = %q
  enforce_dns_only = false
  zone_defaults = {
    flatten_all_cnames  = false
    multi_provider      = false
    secondary_overrides = false
    ns_ttl              = %d
    zone_mode           = "standard"
    nameservers = {
      type = "cloudflare.advanced"
    }
%s
    soa = {
      rname   = "dns.cloudflare.com"
      refresh = 10000
      retry   = 2400
      expire  = 604800
      min_ttl = 1800
      ttl     = 3600
    }
  }
}
`, foundationDNSTestAccountID, nsTTL, extra)
}

func requireNoFoundationDNSInPatches(api *fakeAccountDNSSettingsAPI) resource.TestCheckFunc {
	return func(_ *terraform.State) error {
		bodies := api.patchBodies()
		if len(bodies) == 0 {
			return fmt.Errorf("no PATCH was sent")
		}
		for _, body := range bodies {
			if _, ok := body["foundation_dns"]; ok {
				return fmt.Errorf("PATCH body includes foundation_dns: %v", body)
			}
			if zoneDefaults, ok := body["zone_defaults"].(map[string]any); ok {
				if _, ok := zoneDefaults["foundation_dns"]; ok {
					return fmt.Errorf("PATCH body includes zone_defaults.foundation_dns: %v", body)
				}
			}
		}
		return nil
	}
}

// Terraform drops an undeclared key from an object assigned to a nested
// attribute, so a configuration that still sets zone_defaults.foundation_dns
// keeps working and the key is never sent. zone_defaults.nameservers.type
// accepts cloudflare.advanced in its place.
func TestAccountDNSSettingsFoundationDNSRemoved(t *testing.T) {
	api := newFakeAccountDNSSettingsAPI(t, "cloudflare.standard")

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: accountDNSSettingsConfig(86400, "    foundation_dns = true\n"),
				Check: resource.ComposeAggregateTestCheckFunc(
					requireNoFoundationDNSInPatches(api),
					resource.TestCheckResourceAttr("cloudflare_account_dns_settings.test", "zone_defaults.nameservers.type", "cloudflare.advanced"),
					resource.TestCheckNoResourceAttr("cloudflare_account_dns_settings.test", "zone_defaults.foundation_dns"),
				),
			},
			{
				Config:             accountDNSSettingsConfig(86400, ""),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			{
				Config: accountDNSSettingsConfig(3600, ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					requireNoFoundationDNSInPatches(api),
					resource.TestCheckResourceAttr("cloudflare_account_dns_settings.test", "zone_defaults.ns_ttl", "3600"),
				),
			},
			{
				Config:             accountDNSSettingsConfig(3600, ""),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}
