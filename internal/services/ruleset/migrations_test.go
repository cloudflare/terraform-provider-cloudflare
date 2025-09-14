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
	v4Config := acctest.LoadTestCase( "migrations/simple_rules_v4.tf", zoneID, rnd)

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

	v4Config := acctest.LoadTestCase( "migrations/rewrite_rules_v4.tf", zoneID, rnd)

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