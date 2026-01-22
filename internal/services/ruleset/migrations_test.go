package ruleset_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateCloudflareRulesetBasic tests migration of basic ruleset with only required attributes
func TestMigrateCloudflareRulesetBasic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	// V4 config with basic ruleset (no rules)
	v4Config := acctest.LoadTestCase("migrations/basic_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("My ruleset %s", rnd))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("phase"), knownvalue.StringExact("http_request_firewall_custom")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact("zone")),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			// Step 1: Create with v4 provider
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}

// TestMigrateCloudflareRulesetSimpleRules tests migration of rules block to list attribute
func TestMigrateCloudflareRulesetSimpleRules(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	// V4 config with simple rules using block syntax
	v4Config := acctest.LoadTestCase("migrations/simple_rules_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(2)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expression"), knownvalue.StringExact("ip.src eq 1.1.1.1")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("block")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("expression"), knownvalue.StringExact("ip.src eq 2.2.2.2")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action"), knownvalue.StringExact("log")),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}

// TestMigrateCloudflareRulesetRewriteRules tests migration with rewrite action and URI blocks
func TestMigrateCloudflareRulesetRewriteRules(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/rewrite_rules_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("rewrite")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("uri").AtMapKey("path").AtMapKey("value"), knownvalue.StringExact("/new-path")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("uri").AtMapKey("query").AtMapKey("value"), knownvalue.StringExact("foo=bar")),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}

// TestMigrateCloudflareRulesetDynamicRules tests migration with dynamic rules blocks
func TestMigrateCloudflareRulesetDynamicRules(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	// V4 config with dynamic rules blocks
	v4Config := acctest.LoadTestCase("migrations/dynamic_rules_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(3)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expression"), knownvalue.StringExact("ip.src eq 1.1.1.1")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("block")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("expression"), knownvalue.StringExact("http.request.uri.path contains \"/admin\"")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action"), knownvalue.StringExact("challenge")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("expression"), knownvalue.StringExact("ip.src eq 2.2.2.2")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action"), knownvalue.StringExact("log")),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}

// TestMigrateCloudflareRulesetRedirectFromValue tests migration with redirect from_value blocks
func TestMigrateCloudflareRulesetRedirectFromValue(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/redirect_from_value_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("phase"), knownvalue.StringExact("http_request_dynamic_redirect")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(2)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("redirect")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("from_value").AtMapKey("preserve_query_string"), knownvalue.Bool(true)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("from_value").AtMapKey("status_code"), knownvalue.Int64Exact(308)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("from_value").AtMapKey("target_url").AtMapKey("expression"), knownvalue.StringExact("concat(\"https://example.com\", http.request.uri.path)")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("enabled"), knownvalue.Bool(false)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("from_value").AtMapKey("target_url").AtMapKey("value"), knownvalue.StringExact("https://example.com/new-path")),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}

// TestMigrateCloudflareRulesetHeadersListToMap tests migration of headers from list to map format
func TestMigrateCloudflareRulesetHeadersListToMap(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/headers_list_to_map_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("phase"), knownvalue.StringExact("http_response_headers_transform")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(2)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("rewrite")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("headers"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action"), knownvalue.StringExact("rewrite")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("headers"), knownvalue.NotNull()),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}

// TestMigrateCloudflareRulesetCacheKeyQueryString tests migration of cache_key query_string include from list to object
func TestMigrateCloudflareRulesetCacheKeyQueryString(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/cache_key_query_string_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("phase"), knownvalue.StringExact("http_request_cache_settings")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(2)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("set_cache_settings")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("cache"), knownvalue.Bool(true)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("cache_key").AtMapKey("custom_key").AtMapKey("query_string").AtMapKey("include"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action"), knownvalue.StringExact("set_cache_settings")),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}
// TestMigrateCloudflareRulesetLogCustomFields tests migration of log custom fields (cookie_fields, request_fields, response_fields)
func TestMigrateCloudflareRulesetLogCustomFields(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/log_custom_fields_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("phase"), knownvalue.StringExact("http_log_custom_fields")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(3)),
		
		// Rule 0: cookie_fields transformation
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("log_custom_field")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("cookie_fields"), knownvalue.ListSizeExact(3)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("cookie_fields").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("session_id")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("cookie_fields").AtSliceIndex(1).AtMapKey("name"), knownvalue.StringExact("tracking_id")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("cookie_fields").AtSliceIndex(2).AtMapKey("name"), knownvalue.StringExact("user_token")),

		// Rule 1: request_fields transformation
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action"), knownvalue.StringExact("log_custom_field")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("request_fields"), knownvalue.ListSizeExact(2)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("request_fields").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("cf.bot_score")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("request_fields").AtSliceIndex(1).AtMapKey("name"), knownvalue.StringExact("http.user_agent")),

		// Rule 2: response_fields transformation
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action"), knownvalue.StringExact("log_custom_field")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action_parameters").AtMapKey("response_fields"), knownvalue.ListSizeExact(2)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action_parameters").AtMapKey("response_fields").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("cf.colo")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action_parameters").AtMapKey("response_fields").AtSliceIndex(1).AtMapKey("name"), knownvalue.StringExact("cf.ray_id")),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}

// TestMigrateCloudflareRulesetEdgeTTLStatusCode tests migration of edge_ttl status_code_ttl numeric fields
func TestMigrateCloudflareRulesetEdgeTTLStatusCode(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/edge_ttl_status_code_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("phase"), knownvalue.StringExact("http_request_cache_settings")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(3)),
		
		// Rule 0: Single status codes with numeric values
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("set_cache_settings")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("edge_ttl").AtMapKey("status_code_ttl"), knownvalue.ListSizeExact(2)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("edge_ttl").AtMapKey("status_code_ttl").AtSliceIndex(0).AtMapKey("status_code"), knownvalue.Int64Exact(200)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("edge_ttl").AtMapKey("status_code_ttl").AtSliceIndex(0).AtMapKey("value"), knownvalue.Int64Exact(86400)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("edge_ttl").AtMapKey("status_code_ttl").AtSliceIndex(1).AtMapKey("status_code"), knownvalue.Int64Exact(404)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("edge_ttl").AtMapKey("status_code_ttl").AtSliceIndex(1).AtMapKey("value"), knownvalue.Int64Exact(300)),
		
		// Rule 1: Status code range with numeric from/to
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action"), knownvalue.StringExact("set_cache_settings")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("edge_ttl").AtMapKey("status_code_ttl"), knownvalue.ListSizeExact(1)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("edge_ttl").AtMapKey("status_code_ttl").AtSliceIndex(0).AtMapKey("status_code_range").AtMapKey("from"), knownvalue.Int64Exact(200)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("edge_ttl").AtMapKey("status_code_ttl").AtSliceIndex(0).AtMapKey("status_code_range").AtMapKey("to"), knownvalue.Int64Exact(299)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("edge_ttl").AtMapKey("status_code_ttl").AtSliceIndex(0).AtMapKey("value"), knownvalue.Int64Exact(3600)),
		
		// Rule 2: Multiple status_code_ttl entries with mixed types
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action"), knownvalue.StringExact("set_cache_settings")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action_parameters").AtMapKey("edge_ttl").AtMapKey("status_code_ttl"), knownvalue.ListSizeExact(3)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action_parameters").AtMapKey("edge_ttl").AtMapKey("status_code_ttl").AtSliceIndex(0).AtMapKey("status_code"), knownvalue.Int64Exact(200)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action_parameters").AtMapKey("edge_ttl").AtMapKey("status_code_ttl").AtSliceIndex(1).AtMapKey("status_code_range").AtMapKey("from"), knownvalue.Int64Exact(400)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action_parameters").AtMapKey("edge_ttl").AtMapKey("status_code_ttl").AtSliceIndex(2).AtMapKey("status_code"), knownvalue.Int64Exact(500)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action_parameters").AtMapKey("edge_ttl").AtMapKey("status_code_ttl").AtSliceIndex(2).AtMapKey("value"), knownvalue.Int64Exact(0)),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}

// TestMigrateCloudflareRulesetWAFManagedOverrides tests migration of WAF overrides (categories and rules remain as arrays)
func TestMigrateCloudflareRulesetWAFManagedOverrides(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/waf_managed_overrides_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("phase"), knownvalue.StringExact("http_request_firewall_managed")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),

		// Rule 0: Categories overrides remain as array
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("execute")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("id"), knownvalue.StringExact("efb7b8c949ac4650a09736fc376e9aee")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("overrides").AtMapKey("categories"), knownvalue.ListSizeExact(2)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("overrides").AtMapKey("categories").AtSliceIndex(0).AtMapKey("category"), knownvalue.StringExact("wordpress")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("overrides").AtMapKey("categories").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("block")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("overrides").AtMapKey("categories").AtSliceIndex(1).AtMapKey("category"), knownvalue.StringExact("joomla")),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}

// TestMigrateCloudflareRulesetRateLimit tests migration of http_ratelimit phase
func TestMigrateCloudflareRulesetRateLimit(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/ratelimit_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("phase"), knownvalue.StringExact("http_ratelimit")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(3)),

		// Rule 0: Ratelimit with characteristics and mitigation_timeout
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("block")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("ratelimit").AtMapKey("characteristics"), knownvalue.ListSizeExact(2)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("ratelimit").AtMapKey("period"), knownvalue.Int64Exact(60)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("ratelimit").AtMapKey("requests_per_period"), knownvalue.Int64Exact(100)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("ratelimit").AtMapKey("mitigation_timeout"), knownvalue.Int64Exact(600)),

		// Rule 1: Ratelimit with counting_expression
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action"), knownvalue.StringExact("challenge")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("ratelimit").AtMapKey("characteristics"), knownvalue.ListSizeExact(2)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("ratelimit").AtMapKey("period"), knownvalue.Int64Exact(10)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("ratelimit").AtMapKey("requests_per_period"), knownvalue.Int64Exact(5)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("ratelimit").AtMapKey("counting_expression"), knownvalue.StringExact("http.request.method eq \"POST\"")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("ratelimit").AtMapKey("mitigation_timeout"), knownvalue.Int64Exact(0)),

		// Rule 2: Verify ratelimit block is object (not array)
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action"), knownvalue.StringExact("block")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("ratelimit").AtMapKey("characteristics"), knownvalue.ListSizeExact(2)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("ratelimit").AtMapKey("period"), knownvalue.Int64Exact(300)),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}

// TestMigrateCloudflareRulesetOriginRoute tests migration of origin routing rules
func TestMigrateCloudflareRulesetOriginRoute(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/origin_route_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("phase"), knownvalue.StringExact("http_request_origin")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(3)),

		// Rule 0: Origin routing with port override
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("route")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("origin").AtMapKey("port"), knownvalue.Int64Exact(8443)),

		// Rule 1: Origin routing with port override
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action"), knownvalue.StringExact("route")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("origin").AtMapKey("port"), knownvalue.Int64Exact(8080)),

		// Rule 2: Origin routing with port override
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action"), knownvalue.StringExact("route")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action_parameters").AtMapKey("origin").AtMapKey("port"), knownvalue.Int64Exact(443)),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}

// TestMigrateCloudflareRulesetLateTransform tests migration of late transform phase rules
func TestMigrateCloudflareRulesetLateTransform(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/late_transform_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("phase"), knownvalue.StringExact("http_request_late_transform")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(3)),

		// Rule 0: Rewrite headers - migrated to map format in v5
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("rewrite")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("headers"), knownvalue.NotNull()),

		// Rule 1: Rewrite headers for domain
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action"), knownvalue.StringExact("rewrite")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("headers"), knownvalue.NotNull()),

		// Rule 2: Add security headers
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action"), knownvalue.StringExact("rewrite")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action_parameters").AtMapKey("headers"), knownvalue.NotNull()),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}

// TestMigrateCloudflareRulesetRedirectFromList tests migration of redirect phase with from_list and from_value
func TestMigrateCloudflareRulesetRedirectFromList(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/redirect_from_list_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("phase"), knownvalue.StringExact("http_request_dynamic_redirect")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(3)),

		// Rule 0: Redirect blocked paths with from_value
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("redirect")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("from_value").AtMapKey("status_code"), knownvalue.Int64Exact(302)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("from_value").AtMapKey("target_url").AtMapKey("value"), knownvalue.StringExact("https://example.com/blocked")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("from_value").AtMapKey("preserve_query_string"), knownvalue.Bool(false)),

		// Rule 1: Dynamic redirect with expression
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action"), knownvalue.StringExact("redirect")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("from_value").AtMapKey("status_code"), knownvalue.Int64Exact(301)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("from_value").AtMapKey("target_url").AtMapKey("expression"), knownvalue.StringExact("concat(\"https://example.com/new\", http.request.uri.path)")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("from_value").AtMapKey("preserve_query_string"), knownvalue.Bool(true)),

		// Rule 2: Static redirect with from_value - verify from_value is object
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action"), knownvalue.StringExact("redirect")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action_parameters").AtMapKey("from_value").AtMapKey("status_code"), knownvalue.Int64Exact(301)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action_parameters").AtMapKey("from_value").AtMapKey("target_url").AtMapKey("value"), knownvalue.StringExact("https://example.com/new-location")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action_parameters").AtMapKey("from_value").AtMapKey("preserve_query_string"), knownvalue.Bool(true)),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}

// TestMigrateCloudflareRulesetSanitize tests migration of sanitize phase rules
func TestMigrateCloudflareRulesetSanitize(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/sanitize_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("phase"), knownvalue.StringExact("http_request_late_transform")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(3)),

		// Rule 0: Remove multiple headers - verify headers are converted to map
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("rewrite")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("headers").AtMapKey("X-Forwarded-For").AtMapKey("operation"), knownvalue.StringExact("remove")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("headers").AtMapKey("X-Real-IP").AtMapKey("operation"), knownvalue.StringExact("remove")),

		// Rule 1: Remove sensitive header
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action"), knownvalue.StringExact("rewrite")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("headers").AtMapKey("Authorization").AtMapKey("operation"), knownvalue.StringExact("remove")),

		// Rule 2: Set custom header
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action"), knownvalue.StringExact("rewrite")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action_parameters").AtMapKey("headers").AtMapKey("X-Custom-Header").AtMapKey("operation"), knownvalue.StringExact("set")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action_parameters").AtMapKey("headers").AtMapKey("X-Custom-Header").AtMapKey("value"), knownvalue.StringExact("custom-value")),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}

// TestMigrateCloudflareRulesetConfigSettings tests migration of config_settings phase rules
func TestMigrateCloudflareRulesetConfigSettings(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/config_settings_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("phase"), knownvalue.StringExact("http_config_settings")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(3)),

		// Rule 0: Basic config settings
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("set_config")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("automatic_https_rewrites"), knownvalue.Bool(true)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("bic"), knownvalue.Bool(true)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("disable_zaraz"), knownvalue.Bool(true)),

		// Rule 1: Security and performance settings
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action"), knownvalue.StringExact("set_config")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("email_obfuscation"), knownvalue.Bool(true)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("polish"), knownvalue.StringExact("lossless")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("security_level"), knownvalue.StringExact("medium")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("ssl"), knownvalue.StringExact("flexible")),

		// Rule 2: API-specific config
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action"), knownvalue.StringExact("set_config")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action_parameters").AtMapKey("automatic_https_rewrites"), knownvalue.Bool(false)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action_parameters").AtMapKey("sxg"), knownvalue.Bool(false)),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}

// TestMigrateCloudflareRulesetCustomErrors tests migration of custom_errors phase rules
func TestMigrateCloudflareRulesetCustomErrors(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/custom_errors_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("phase"), knownvalue.StringExact("http_custom_errors")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(3)),

		// Rule 0: Custom 404 error
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("serve_error")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("content"), knownvalue.StringExact("<html><body>Custom 404 Error</body></html>")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("content_type"), knownvalue.StringExact("text/html")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("status_code"), knownvalue.Int64Exact(404)),

		// Rule 1: Custom 500 error
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action"), knownvalue.StringExact("serve_error")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("status_code"), knownvalue.Int64Exact(500)),

		// Rule 2: Custom 403 error with JSON
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action"), knownvalue.StringExact("serve_error")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action_parameters").AtMapKey("content_type"), knownvalue.StringExact("application/json")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action_parameters").AtMapKey("status_code"), knownvalue.Int64Exact(403)),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}

// TestMigrateCloudflareRulesetResponseCompression tests migration of response_compression phase rules
func TestMigrateCloudflareRulesetResponseCompression(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/response_compression_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("phase"), knownvalue.StringExact("http_response_compression")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(3)),

		// Rule 0: Single compression algorithm - verify algorithms remain as array
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("compress_response")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("algorithms"), knownvalue.ListSizeExact(1)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("algorithms").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("gzip")),

		// Rule 1: Multiple algorithms
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action"), knownvalue.StringExact("compress_response")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("algorithms"), knownvalue.ListSizeExact(2)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("algorithms").AtSliceIndex(1).AtMapKey("name"), knownvalue.StringExact("brotli")),

		// Rule 2: Three algorithms
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action"), knownvalue.StringExact("compress_response")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action_parameters").AtMapKey("algorithms"), knownvalue.ListSizeExact(3)),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}

// TestMigrateCloudflareRulesetResponseFirewall tests migration of response_firewall_managed phase rules
func TestMigrateCloudflareRulesetResponseFirewall(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/response_firewall_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("phase"), knownvalue.StringExact("http_response_firewall_managed")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(2)),

		// Rule 0: Log 403 responses
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("log")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expression"), knownvalue.StringExact("http.response.code eq 403")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("description"), knownvalue.StringExact("Log 403 responses")),

		// Rule 1: Log responses with sensitive header
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action"), knownvalue.StringExact("log")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("expression"), knownvalue.StringExact("http.response.headers[\"x-custom-header\"][0] eq \"sensitive\"")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("description"), knownvalue.StringExact("Log responses with sensitive header")),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}

// TestMigrateCloudflareRulesetSBFM tests migration of SBFM (Super Bot Fight Mode) phase rules
func TestMigrateCloudflareRulesetSBFM(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/sbfm_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("phase"), knownvalue.StringExact("http_request_firewall_custom")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(3)),

		// Rule 0: Managed challenge for suspected bots
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("managed_challenge")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expression"), knownvalue.StringExact("cf.bot_management.score lt 30")),

		// Rule 1: Block known bad bots
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action"), knownvalue.StringExact("block")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("expression"), knownvalue.StringExact("(not cf.bot_management.verified_bot) and (cf.bot_management.score lt 10)")),

		// Rule 2: JS challenge for API endpoints
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action"), knownvalue.StringExact("js_challenge")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("expression"), knownvalue.StringExact("(http.request.uri.path contains \"/api\") and (cf.bot_management.score lt 50)")),
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{{
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		}}, migrationSteps...),
	})
}
