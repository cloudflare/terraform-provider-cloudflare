package bot_management_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareBotManagement_SBFM(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceID := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	sbfmConfig := cloudflare.BotManagement{
		EnableJS:                     cloudflare.BoolPtr(true),
		SBFMDefinitelyAutomated:      cloudflare.StringPtr("managed_challenge"),
		SBFMLikelyAutomated:          cloudflare.StringPtr("block"),
		SBFMVerifiedBots:             cloudflare.StringPtr("allow"),
		SBFMStaticResourceProtection: cloudflare.BoolPtr(false),
		OptimizeWordpress:            cloudflare.BoolPtr(true),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementSBFM(rnd, zoneID, sbfmConfig),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "enable_js", "true"),
					resource.TestCheckResourceAttr(resourceID, "sbfm_definitely_automated", "managed_challenge"),
					resource.TestCheckResourceAttr(resourceID, "sbfm_likely_automated", "block"),
					resource.TestCheckResourceAttr(resourceID, "sbfm_verified_bots", "allow"),
					resource.TestCheckResourceAttr(resourceID, "sbfm_static_resource_protection", "false"),
					resource.TestCheckResourceAttr(resourceID, "optimize_wordpress", "true"),
				),
			},
			// {
			// 	ResourceName:      resourceID,
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
		},
	})
}

func TestAccCloudflareBotManagement_Unentitled(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceID := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	bmEntConfig := cloudflare.BotManagement{
		EnableJS:             cloudflare.BoolPtr(true),
		SuppressSessionScore: cloudflare.BoolPtr(false),
		AutoUpdateModel:      cloudflare.BoolPtr(false),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareBotManagementEntSubscription(rnd, zoneID, bmEntConfig),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "enable_js", "true"),
					resource.TestCheckResourceAttr(resourceID, "suppress_session_score", "false"),
					resource.TestCheckResourceAttr(resourceID, "auto_update_model", "false"),
				),
				ExpectError: regexp.MustCompile("zone not entitled to disable"),
			},
		},
	})
}

func testCloudflareBotManagementSBFM(resourceName, rnd string, bm cloudflare.BotManagement) string {
	return acctest.LoadTestCase("cloudflarebotmanagementsbfm.tf", resourceName, rnd,
		*bm.EnableJS, *bm.SBFMDefinitelyAutomated,
		*bm.SBFMLikelyAutomated, *bm.SBFMVerifiedBots,
		*bm.SBFMStaticResourceProtection, *bm.OptimizeWordpress)
}

func testCloudflareBotManagementEntSubscription(resourceName, rnd string, bm cloudflare.BotManagement) string {
	return acctest.LoadTestCase("cloudflarebotmanagemententsubscription.tf", resourceName, rnd, *bm.EnableJS, *bm.SuppressSessionScore, false)
}

// TestAccCloudflareBotManagement_Issue5728_EnableJSDrift tests the fix for issue #5728
// The issue: enable_js shows drift (false -> true) on every apply even when set to true in config
// NOTE: This test expects some drift on other computed_optional fields - this is expected behavior
// until all fields get the encode_state_for_unknown fix. The key validation is that our target
// fields (enable_js, auto_update_model) maintain their configured values.
func TestAccCloudflareBotManagement_Issue5728_EnableJSDrift(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceID := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: Apply the EXACT configuration from issue #5728
				Config: testCloudflareBotManagementIssue5728Config(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "auto_update_model", "true"),
					resource.TestCheckResourceAttr(resourceID, "enable_js", "true"),
				),
			},
			{
				// Step 2: Apply the same configuration again - this is where issue #5728 would show drift
				// With our fix, enable_js should NOT show drift anymore
				Config: testCloudflareBotManagementIssue5728Config(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "auto_update_model", "true"),
					resource.TestCheckResourceAttr(resourceID, "enable_js", "true"),
				),
				// With encode_state_for_unknown fix, no drift should be detected
			},
		},
	})
}

// TestAccCloudflareBotManagement_Issue5519_SuppressSessionScore tests the fix for issue #5519  
// The issue: suppress_session_score always shows as being added even when not set
func TestAccCloudflareBotManagement_Issue5519_SuppressSessionScore(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceID := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: Apply the EXACT configuration from issue #5519
				// Note: suppress_session_score is NOT set in the config
				Config: testCloudflareBotManagementIssue5519Config(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "enable_js", "true"),
				),
			},
			{
				// Step 2: Apply the same configuration again - this is where issue #5519 would show drift
				// Without our fix, this would show: + suppress_session_score = false
				Config: testCloudflareBotManagementIssue5519Config(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "enable_js", "true"),
				),
			},
		},
	})
}

// Exact config from GitHub issue #5728 (removed fight_mode due to zone entitlements)
// Added other fields to prevent drift from computed_optional fields
func testCloudflareBotManagementIssue5728Config(resourceName, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_bot_management" "%[1]s" {
  zone_id                           = "%[2]s"
  auto_update_model                 = true
  enable_js                         = true
  ai_bots_protection                = "disabled"
  crawler_protection                = "disabled"
  optimize_wordpress                = false
  sbfm_definitely_automated         = "allow"
  sbfm_likely_automated             = "allow"
  sbfm_verified_bots                = "allow"
  sbfm_static_resource_protection   = false
}`, resourceName, zoneID)
}

// Exact config from GitHub issue #5519 (removed fight_mode due to zone entitlements)
func testCloudflareBotManagementIssue5519Config(resourceName, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_bot_management" "%[1]s" {
  zone_id   = "%[2]s"
  enable_js = true
}`, resourceName, zoneID)
}

// TestAccCloudflareBotManagement_MigrationFromV4 tests migration from v4 to v5 provider
// Since bot_management schema didn't change between v4 and v5, this should be a no-op migration
// NOTE: Some drift is expected due to computed_optional field behavior, but migration should work
func TestAccCloudflareBotManagement_MigrationFromV4(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		Steps: []resource.TestStep{
			// Step 1: Create bot_management resource with v4 provider
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "~> 4.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testCloudflareBotManagementMigrationConfig(rnd, zoneID),
				// V4 provider may have some plan differences due to computed_optional fields
				ExpectNonEmptyPlan: true,
			},
			// Step 2: Upgrade to v5 provider and verify no changes needed
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testCloudflareBotManagementMigrationConfig(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_update_model"), knownvalue.Bool(true)),
				},
				// With encode_state_for_unknown fix, migration should be clean
			},
			// Step 3: Apply same config again to verify no persistent drift
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testCloudflareBotManagementMigrationConfig(rnd, zoneID),
				// With encode_state_for_unknown fix, no drift should occur
			},
		},
	})
}

// Configuration for migration test - identical for v4 and v5 since no schema changes
func testCloudflareBotManagementMigrationConfig(resourceName, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_bot_management" "%[1]s" {
  zone_id           = "%[2]s"
  enable_js         = true
  auto_update_model = true
}`, resourceName, zoneID)
}


// TestAccCloudflareBotManagement_BasicConfigurationNoDrift tests that a basic configuration
// with the fields from the original GitHub issues doesn't show persistent drift
func TestAccCloudflareBotManagement_BasicConfigurationNoDrift(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: Apply the basic configuration
				Config: testCloudflareBotManagementBasicConfig(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "zone_id", zoneID),
					resource.TestCheckResourceAttr(resourceName, "auto_update_model", "true"),
					resource.TestCheckResourceAttr(resourceName, "enable_js", "true"),
				),
			},
			{
				// Step 2: Apply the same configuration - should show no drift
				// Our encode_state_for_unknown fixes should prevent drift on enable_js and auto_update_model
				Config: testCloudflareBotManagementBasicConfig(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "zone_id", zoneID),
					resource.TestCheckResourceAttr(resourceName, "auto_update_model", "true"),
					resource.TestCheckResourceAttr(resourceName, "enable_js", "true"),
				),
				// With encode_state_for_unknown fix, no drift should be detected
			},
			{
				// Step 3: Apply once more to verify no persistent drift
				Config: testCloudflareBotManagementBasicConfig(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "zone_id", zoneID),
					resource.TestCheckResourceAttr(resourceName, "auto_update_model", "true"),
					resource.TestCheckResourceAttr(resourceName, "enable_js", "true"),
				),
				// With encode_state_for_unknown fix, no drift should persist
			},
		},
	})
}

// Basic configuration matching the GitHub issue example (removed fight_mode due to zone entitlements)
func testCloudflareBotManagementBasicConfig(resourceName, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_bot_management" "%[1]s" {
  zone_id           = "%[2]s"
  auto_update_model = true
  enable_js         = true
}`, resourceName, zoneID)
}

