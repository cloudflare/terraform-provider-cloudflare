package access_rule_test

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestMigrateAccessRule_AccountLevel(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()
	resourceName := "cloudflare_access_rule." + rnd

	v4Config := acctest.LoadTestCase("accessrulemigrationaccount.tf", accountID, rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: testAccCheckCloudflareAccessRuleDestroy,
		WorkingDir:   tmpDir,
		Steps: append([]resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
		}, // Step 2: Run migration and verify state
			acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mode"), knownvalue.StringExact("block")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("notes"), knownvalue.StringExact("Block malicious IP")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("configuration").AtMapKey("target"), knownvalue.StringExact("ip")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("configuration").AtMapKey("value"), knownvalue.StringExact("198.51.100.50")),
			})...),
	})
}

func TestMigrateAccessRule_ZoneLevel(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()
	resourceName := "cloudflare_access_rule." + rnd

	v4Config := acctest.LoadTestCase("accessrulemigrationzone.tf", zoneID, rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		CheckDestroy: testAccCheckCloudflareAccessRuleDestroy,
		WorkingDir:   tmpDir,
		Steps: append([]resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
		}, // Step 2: Run migration and verify state
			acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mode"), knownvalue.StringExact("challenge")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("notes"), knownvalue.StringExact("Challenge suspicious country")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("configuration").AtMapKey("target"), knownvalue.StringExact("country")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("configuration").AtMapKey("value"), knownvalue.StringExact("TV")),
			})...),
	})
}

func TestMigrateAccessRule_WithoutNotes(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()
	resourceName := "cloudflare_access_rule." + rnd

	v4Config := acctest.LoadTestCase("accessrulemigrationnonotes.tf", accountID, rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: testAccCheckCloudflareAccessRuleDestroy,
		WorkingDir:   tmpDir,
		Steps: append([]resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
		}, // Step 2: Run migration and verify state
			acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mode"), knownvalue.StringExact("whitelist")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("configuration").AtMapKey("target"), knownvalue.StringExact("ip_range")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("configuration").AtMapKey("value"), knownvalue.StringExact("198.51.100.0/24")),
			})...),
	})
}
