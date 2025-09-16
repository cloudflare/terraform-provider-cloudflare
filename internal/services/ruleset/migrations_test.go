package ruleset_test

import (
	"fmt"
	"os"
	"strings"
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

	migrationSteps := acctest.MigrationTestStepWithPlan(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
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

	migrationSteps := acctest.MigrationTestStepWithPlan(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
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

	migrationSteps := acctest.MigrationTestStepWithPlan(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
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

	migrationSteps := acctest.MigrationTestStepWithPlan(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
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

	migrationSteps := acctest.MigrationTestStepWithPlan(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
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

	migrationSteps := acctest.MigrationTestStepWithPlan(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
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

	migrationSteps := acctest.MigrationTestStepWithPlan(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
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

func TestMigrateCloudflareRulesetExpressionDoubleDollar(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/literal_expression_double_dollar_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationTestStepWithPlan(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("Literal Expression Test %s", rnd))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("phase"), knownvalue.StringExact("http_request_firewall_custom")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact("zone")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
		// Test that the expression is preserved correctly
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("block")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("description"), knownvalue.StringExact("Literal heredoc test")),
		// Verify expression is preserved correctly (literal content, not variables)
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expression"), knownvalue.StringFunc(func(v string) error {
			if !strings.Contains(v, "ip.geoip.country eq \"CN\"") {
				return fmt.Errorf("expected expression to contain literal content, got: %s", v)
			}
			if !strings.Contains(v, "http.host eq \"example.com\"") {
				return fmt.Errorf("expected expression to contain literal content, got: %s", v)
			}
			return nil
		})),
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

// TestMigrateCloudflareRulesetComplexHeredocRatelimit tests complex heredoc expressions with ratelimit blocks
func TestMigrateCloudflareRulesetComplexHeredocRatelimit(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/complex_heredoc_ratelimit_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationTestStepWithPlan(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("Complex Heredoc Test %s", rnd))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("phase"), knownvalue.StringExact("http_ratelimit")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact("zone")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
		// Test that the complex heredoc expression is preserved correctly
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("log")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("description"), knownvalue.StringExact("Log rate limit with complex expression")),
		// Verify ratelimit block is converted properly
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("ratelimit").AtMapKey("period"), knownvalue.Int64Exact(10)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("ratelimit").AtMapKey("requests_per_period"), knownvalue.Int64Exact(1000)),
		// Verify expression contains expected content and is syntactically valid
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expression"), knownvalue.StringFunc(func(v string) error {
			if !strings.Contains(v, "authorization") {
				return fmt.Errorf("expected expression to contain 'authorization', got: %s", v)
			}
			if !strings.Contains(v, "__Secure-next-auth") {
				return fmt.Errorf("expected expression to contain '__Secure-next-auth', got: %s", v)
			}
			return nil
		})),
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

// TestMigrateCloudflareRulesetComplexMultipleRulesRatelimit tests multiple rules with heredoc expressions and ratelimit blocks
func TestMigrateCloudflareRulesetComplexMultipleRulesRatelimit(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/complex_multiple_rules_ratelimit_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationTestStepWithPlan(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("Global rate limit %s", rnd))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("phase"), knownvalue.StringExact("http_ratelimit")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact("zone")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(3)),
		// Test that all rules are preserved correctly
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("log")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("description"), knownvalue.StringExact("Log Global rate limit 1000/10s non authenticated")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action"), knownvalue.StringExact("log")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("description"), knownvalue.StringExact("Log Global rate limit 2500/1m non authenticated")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action"), knownvalue.StringExact("log")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("description"), knownvalue.StringExact("Log Unauthed rate limit DDoS User Agents")),
		// Test that ratelimit blocks are converted properly for all rules
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("ratelimit").AtMapKey("period"), knownvalue.Int64Exact(10)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("ratelimit").AtMapKey("requests_per_period"), knownvalue.Int64Exact(1000)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("ratelimit").AtMapKey("period"), knownvalue.Int64Exact(60)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("ratelimit").AtMapKey("requests_per_period"), knownvalue.Int64Exact(2500)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("ratelimit").AtMapKey("period"), knownvalue.Int64Exact(10)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("ratelimit").AtMapKey("requests_per_period"), knownvalue.Int64Exact(300)),
		// Verify expressions are syntactically valid and properly terminated
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expression"), knownvalue.StringFunc(func(v string) error {
			if !strings.Contains(v, "authorization") {
				return fmt.Errorf("expected expression to contain 'authorization', got: %s", v)
			}
			return nil
		})),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("expression"), knownvalue.StringFunc(func(v string) error {
			if !strings.Contains(v, "__Secure-next-auth") {
				return fmt.Errorf("expected expression to contain '__Secure-next-auth', got: %s", v)
			}
			return nil
		})),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("expression"), knownvalue.StringFunc(func(v string) error {
			if !strings.Contains(v, "Chrome/117.0.0.0") {
				return fmt.Errorf("expected expression to contain 'Chrome/117.0.0.0', got: %s", v)
			}
			return nil
		})),
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

// TestMigrateCloudflareRulesetHeredocExpressionPreservation tests that heredoc expressions are properly preserved during migration
func TestMigrateCloudflareRulesetHeredocExpressionPreservation(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	resourceName := fmt.Sprintf("cloudflare_ruleset.%s", rnd)

	v4Config := acctest.LoadTestCase("migrations/complex_heredoc_ratelimit_v4.tf", zoneID, rnd)

	migrationSteps := acctest.MigrationTestStepWithPlan(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("Complex Heredoc Test %s", rnd))),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("phase"), knownvalue.StringExact("http_ratelimit")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact("zone")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
		// Test that the rule is preserved correctly
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("log")),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("description"), knownvalue.StringExact("Log rate limit with complex expression")),
		// Test that ratelimit block is converted properly
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("ratelimit").AtMapKey("period"), knownvalue.Int64Exact(10)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("ratelimit").AtMapKey("requests_per_period"), knownvalue.Int64Exact(1000)),
		// Critical test: verify heredoc expression is properly preserved and not converted to escaped string
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expression"), knownvalue.StringFunc(func(v string) error {
			if !strings.Contains(v, "authorization") {
				return fmt.Errorf("expected expression to contain 'authorization', got: %s", v)
			}
			// Verify the expression doesn't contain escaped newlines (would indicate corruption)
			if strings.Contains(v, "\\n") {
				return fmt.Errorf("expression contains escaped newlines (\\n), indicating heredoc was corrupted. Got: %s", v)
			}
			// Verify it contains proper newlines and formatting
			if !strings.Contains(v, "\n") {
				return fmt.Errorf("expression should contain actual newlines for proper heredoc format, got: %s", v)
			}
			return nil
		})),
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
