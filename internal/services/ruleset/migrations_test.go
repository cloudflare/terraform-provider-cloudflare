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
	t.Skip("TODO: investigate refresh plan non-empty")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/cache_key_query_string_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
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
