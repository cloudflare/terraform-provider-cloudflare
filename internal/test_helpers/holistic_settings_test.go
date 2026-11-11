package test_helpers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_gateway_settings"
	gatewaymigration "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_gateway_settings/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_organization"
	organizationmigration "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_organization/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	providerschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	frameworkresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// Use the real Terraform planning/apply protocol with only these three resources
// and a local API. This needs neither Cloudflare credentials nor live resources.
type holisticSettingsProvider struct{ baseURL string }

func (p *holisticSettingsProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cloudflare"
}
func (p *holisticSettingsProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = providerschema.Schema{}
}
func (p *holisticSettingsProvider) Configure(_ context.Context, _ provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	resp.ResourceData = cloudflare.NewClient(option.WithBaseURL(p.baseURL), option.WithAPIToken("test-token"), option.WithMaxRetries(0))
}
func (p *holisticSettingsProvider) Resources(context.Context) []func() frameworkresource.Resource {
	return []func() frameworkresource.Resource{
		zero_trust_organization.NewResource,
		zero_trust_gateway_settings.NewResource,
		zero_trust_device_settings.NewResource,
	}
}
func (p *holisticSettingsProvider) DataSources(context.Context) []func() datasource.DataSource {
	return nil
}

// BUGS-2161: settings unavailable in v4 are restored into explicit v5 HCL
// from a pre-migration API snapshot. No provider omission behavior is changed.
func TestHolisticSettingsMigrationWithExplicitConfiguration(t *testing.T) {
	terraformPath, err := exec.LookPath("terraform")
	if err != nil {
		t.Skip("Terraform CLI is required for the local protocol regression test")
	}
	t.Setenv("TF_ACC_TERRAFORM_PATH", terraformPath)
	t.Setenv("TF_ACC_PROVIDER_NAMESPACE", "cloudflare")

	var mu sync.Mutex
	live := map[string]map[string]any{
		"/access/organizations": {
			"name": "Live organization", "auth_domain": "example.cloudflareaccess.com",
			"strict_service_token_auth":                   false,
			"deny_unmatched_requests":                     true,
			"deny_unmatched_requests_exempted_zone_names": []any{"public.example.com"},
			"mfa_required_for_all_apps":                   true, "mfa_configuration_allowed": true,
			"mfa_config":               map[string]any{"allowed_authenticators": []any{"totp"}, "session_duration": "24h"},
			"mfa_piv_key_requirements": map[string]any{"pin_policy": "always", "touch_policy": "always", "require_fips_device": true},
		},
		"/gateway/configuration": {
			"created_at": "2026-10-01T00:00:00Z", "updated_at": "2026-10-01T00:00:00Z",
			"settings": map[string]any{
				"tls_decrypt":   map[string]any{"enabled": true},
				"host_selector": map[string]any{"enabled": true},
				"sandbox":       map[string]any{"enabled": true, "fallback_action": "block"},
			},
		},
		"/devices/settings": {
			"external_emergency_signal_enabled":     true,
			"external_emergency_signal_url":         "https://192.0.2.1/signal",
			"external_emergency_signal_interval":    "7m",
			"external_emergency_signal_fingerprint": strings.Repeat("a", 64),
			"gateway_proxy_enabled":                 false,
		},
	}
	deviceWrites := 0
	initialOrg := make(map[string]any)
	for name, value := range live["/access/organizations"] {
		initialOrg[name] = value
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		var key string
		for candidate := range live {
			if strings.HasSuffix(r.URL.Path, candidate) {
				key = candidate
				break
			}
		}
		if key == "" {
			http.Error(w, "unexpected endpoint", http.StatusNotFound)
			return
		}
		if r.Method == http.MethodPut || r.Method == http.MethodPatch {
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Errorf("invalid request body: %v", err)
				http.Error(w, "invalid JSON", http.StatusBadRequest)
				return
			}
			if key == "/devices/settings" {
				deviceWrites++
			}
			// Organization and Gateway PUT replace their writable configuration.
			// Device updates merge supplied fields, matching the service code.
			if key == "/access/organizations" {
				live[key] = map[string]any{}
			}
			for name, value := range body {
				live[key][name] = value
			}
		}
		result := live[key]
		if exemptions, ok := result["deny_unmatched_requests_exempted_zone_names"].([]any); key == "/access/organizations" && ok && len(exemptions) == 0 {
			// Match the API's omission of cleared optional lists without changing
			// the mock's stored object, which the checks inspect separately.
			result = make(map[string]any, len(live[key]))
			for name, value := range live[key] {
				result[name] = value
			}
			delete(result, "deny_unmatched_requests_exempted_zone_names")
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"success": true, "errors": []any{}, "messages": []any{}, "result": result})
	}))
	defer server.Close()

	check := func(enabled bool) resource.TestCheckFunc {
		return func(state *terraform.State) error {
			mu.Lock()
			defer mu.Unlock()
			org := live["/access/organizations"]
			for _, name := range []string{"deny_unmatched_requests", "mfa_required_for_all_apps"} {
				if org[name] != enabled {
					return fmt.Errorf("%s = %v, want %v", name, org[name], enabled)
				}
			}
			for _, name := range []string{"mfa_config", "mfa_piv_key_requirements"} {
				if !reflect.DeepEqual(org[name], initialOrg[name]) {
					return fmt.Errorf("org security setting %s changed: got %v, want %v", name, org[name], initialOrg[name])
				}
			}
			wantExemptions := []any{}
			if enabled {
				wantExemptions = initialOrg["deny_unmatched_requests_exempted_zone_names"].([]any)
			}
			if !reflect.DeepEqual(org["deny_unmatched_requests_exempted_zone_names"], wantExemptions) {
				return fmt.Errorf("organization exemptions = %v, want %v", org["deny_unmatched_requests_exempted_zone_names"], wantExemptions)
			}
			settings := live["/gateway/configuration"]["settings"].(map[string]any)
			for _, name := range []string{"host_selector", "sandbox"} {
				value, ok := settings[name].(map[string]any)
				if !ok || value["enabled"] != enabled {
					return fmt.Errorf("Gateway %s = %v, want enabled=%v", name, settings[name], enabled)
				}
			}
			device := live["/devices/settings"]
			for name, want := range map[string]any{
				"gateway_proxy_enabled": true, "gateway_udp_proxy_enabled": true,
				"root_certificate_installation_enabled": true, "use_zt_virtual_ip": false, "disable_for_time": float64(0),
			} {
				if device[name] != want {
					return fmt.Errorf("migrated device setting %s = %v, want %v", name, device[name], want)
				}
			}
			if device["external_emergency_signal_enabled"] != enabled ||
				device["external_emergency_signal_url"] != "https://192.0.2.1/signal" ||
				device["external_emergency_signal_interval"] != "7m" ||
				device["external_emergency_signal_fingerprint"] != strings.Repeat("a", 64) {
				return fmt.Errorf("live emergency disconnect settings were changed: %v", device)
			}
			if deviceWrites == 0 {
				return fmt.Errorf("device-settings create path was not exercised")
			}
			if !enabled && deviceWrites < 2 {
				return fmt.Errorf("device-settings update path was not exercised")
			}
			for _, address := range []string{"cloudflare_zero_trust_organization.migrated", "cloudflare_zero_trust_gateway_settings.migrated", "cloudflare_zero_trust_device_settings.migrated_proxy"} {
				if state.RootModule().Resources[address] == nil {
					return fmt.Errorf("missing migrated resource %s", address)
				}
			}
			return nil
		}
	}

	workingDir := t.TempDir()
	resource.UnitTest(t, resource.TestCase{
		WorkingDir: workingDir,
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"cloudflare": providerserver.NewProtocol6WithError(&holisticSettingsProvider{baseURL: server.URL}),
		},
		Steps: []resource.TestStep{
			{
				PreConfig: func() { seedHolisticV4State(t, workingDir) },
				Config:    holisticMigrationConfig(false), Check: check(true),
			},
			{Config: holisticMigrationConfig(true), Check: check(false)},
			{Config: holisticMigrationConfig(true), PlanOnly: true},
			{
				Config: holisticMigrationWithoutRestoredSettings(),
				Check: func(_ *terraform.State) error {
					mu.Lock()
					defer mu.Unlock()
					org := live["/access/organizations"]
					if org["mfa_config"] != nil || org["mfa_piv_key_requirements"] != nil {
						return fmt.Errorf("removing explicit organization settings did not clear them: %v", org)
					}
					settings := live["/gateway/configuration"]["settings"].(map[string]any)
					if settings["host_selector"] != nil || settings["sandbox"] != nil {
						return fmt.Errorf("removing explicit Gateway settings did not clear them: %v", settings)
					}
					return nil
				},
			},
		},
	})
}

func holisticMigrationWithoutRestoredSettings() string {
	return strings.NewReplacer(
		`  mfa_config = { allowed_authenticators = ["totp"], session_duration = "24h" }`, "",
		`  mfa_piv_key_requirements = { pin_policy = "always", touch_policy = "always", require_fips_device = true }`, "",
		`host_selector = { enabled = false }`, "",
		`sandbox = { enabled = false, fallback_action = "allow" }`, "",
	).Replace(holisticMigrationConfig(true))
}

func holisticMigrationConfig(disable bool) string {
	// These explicit values are the output of the migration-only snapshot step.
	orgSettings := `deny_unmatched_requests = true
  deny_unmatched_requests_exempted_zone_names = ["public.example.com"]
  mfa_required_for_all_apps = true`
	gatewaySettings := `host_selector = { enabled = true }
  sandbox = { enabled = true, fallback_action = "block" }`
	deviceSettings := "external_emergency_signal_enabled = true"
	if disable {
		orgSettings = `deny_unmatched_requests = false
 mfa_required_for_all_apps = false
  deny_unmatched_requests_exempted_zone_names = []`
		gatewaySettings = "host_selector = { enabled = false }\n sandbox = { enabled = false, fallback_action = \"allow\" }"
		deviceSettings = "external_emergency_signal_enabled = false"
	}
	return fmt.Sprintf(`
terraform {
  required_providers {
    cloudflare = { source = "cloudflare/cloudflare" }
  }
}
provider "cloudflare" {}
moved {
  from = cloudflare_zero_trust_access_organization.migrated
  to = cloudflare_zero_trust_organization.migrated
}
resource "cloudflare_zero_trust_organization" "migrated" {
  account_id = "0123456789abcdef0123456789abcdef"
  name = "Migrated organization"
  auth_domain = "example.cloudflareaccess.com"
  mfa_config = { allowed_authenticators = ["totp"], session_duration = "24h" }
  mfa_piv_key_requirements = { pin_policy = "always", touch_policy = "always", require_fips_device = true }
  %s
}
moved {
  from = cloudflare_teams_account.migrated
  to = cloudflare_zero_trust_gateway_settings.migrated
}
resource "cloudflare_zero_trust_gateway_settings" "migrated" {
  account_id = "0123456789abcdef0123456789abcdef"
  settings = {
    tls_decrypt = { enabled = false }
    %s
  }
}
resource "cloudflare_zero_trust_device_settings" "migrated_proxy" {
  account_id = "0123456789abcdef0123456789abcdef"
  gateway_proxy_enabled = true
  gateway_udp_proxy_enabled = true
  root_certificate_installation_enabled = true
  use_zt_virtual_ip = false
  disable_for_time = 0
  external_emergency_signal_url = "https://192.0.2.1/signal"
  external_emergency_signal_interval = "7m"
  external_emergency_signal_fingerprint = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  %s
}
`, orgSettings, gatewaySettings, deviceSettings)
}

// Seed representative SDKv2 state, including the legacy singleton-list proxy
// block. Terraform itself invokes the real MoveState handlers before refresh,
// planning and the first apply. No live API or v5 import is used to build state.
func seedHolisticV4State(t *testing.T, workingDir string) {
	t.Helper()
	entries, err := os.ReadDir(workingDir)
	if err != nil {
		t.Fatal(err)
	}
	var directories []string
	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), "work") {
			directories = append(directories, filepath.Join(workingDir, entry.Name()))
		}
	}
	if len(directories) != 1 {
		t.Fatalf("expected one Terraform working directory, got %v", directories)
	}
	legacyResource := func(resourceType string, legacySchema schema.Schema, attributes map[string]any) map[string]any {
		for name := range legacySchema.Attributes {
			if _, configured := attributes[name]; !configured {
				attributes[name] = nil
			}
		}
		for name := range legacySchema.Blocks {
			attributes[name] = []any{}
		}
		return map[string]any{
			"mode": "managed", "type": resourceType, "name": "migrated",
			"provider":  `provider["registry.terraform.io/cloudflare/cloudflare"]`,
			"instances": []any{map[string]any{"schema_version": 0, "attributes": attributes}},
		}
	}
	accountID := "0123456789abcdef0123456789abcdef"
	state := map[string]any{
		"version": 4, "terraform_version": "1.14.5", "serial": 1,
		"lineage": "975bec6b-a878-4731-82b5-ec3fd42cd55d", "outputs": map[string]any{},
		"resources": []any{
			legacyResource("cloudflare_zero_trust_access_organization", organizationmigration.SourceCloudflareAccessOrganizationSchema(), map[string]any{
				"id": accountID, "account_id": accountID, "name": "Live organization", "auth_domain": "example.cloudflareaccess.com",
				"allow_authenticate_via_warp": false, "auto_redirect_to_identity": false, "is_ui_read_only": false,
			}),
			legacyResource("cloudflare_teams_account", gatewaymigration.SourceV4ZeroTrustGatewaySettingsSchema(), map[string]any{
				"account_id": accountID, "tls_decrypt_enabled": true,
				"proxy": []any{map[string]any{"tcp": true, "udp": true, "root_ca": true, "virtual_ip": false, "disable_for_time": 0}},
			}),
		},
	}
	data, err := json.Marshal(state)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(directories[0], "terraform.tfstate"), data, 0600); err != nil {
		t.Fatal(err)
	}
}
