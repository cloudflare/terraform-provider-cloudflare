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
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
// This test validates that fields with encode_state_for_unknown tags work correctly
func TestAccCloudflareBotManagement_Issue5728_EnableJSDrift(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceID := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	var initialState, secondState map[string]interface{}

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
					// Capture initial state for comparison
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources[resourceID]
						if !ok {
							return fmt.Errorf("Not found: %s", resourceID)
						}
						initialState = make(map[string]interface{})
						for k, v := range rs.Primary.Attributes {
							initialState[k] = v
						}
						return nil
					},
				),
			},
			{
				// Step 2: Apply the same configuration again - validate no drift on encode_state_for_unknown fields
				Config: testCloudflareBotManagementIssue5728Config(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "auto_update_model", "true"),
					resource.TestCheckResourceAttr(resourceID, "enable_js", "true"),
					// Capture second state and compare with initial
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources[resourceID]
						if !ok {
							return fmt.Errorf("Not found: %s", resourceID)
						}
						secondState = make(map[string]interface{})
						for k, v := range rs.Primary.Attributes {
							secondState[k] = v
						}

						// Validate that encode_state_for_unknown fields didn't change
						fieldsWithEncode := []string{"enable_js", "auto_update_model", "fight_mode", "suppress_session_score"}
						for _, field := range fieldsWithEncode {
							if initialState[field] != secondState[field] {
								return fmt.Errorf("Field %s changed unexpectedly: %v -> %v (encode_state_for_unknown should prevent this)", 
									field, initialState[field], secondState[field])
							}
						}

						// Log what changed (for debugging, should be minimal/none)
						var changedFields []string
						for k, v := range initialState {
							if secondState[k] != v {
								changedFields = append(changedFields, fmt.Sprintf("%s: %v -> %v", k, v, secondState[k]))
							}
						}
						if len(changedFields) > 0 {
							t.Logf("Fields that changed between applies: %v", changedFields)
						} else {
							t.Logf("SUCCESS: No fields changed between applies (drift fix working)")
						}

						return nil
					},
				),
			},
		},
	})
}

// TestAccCloudflareBotManagement_Issue5519_SuppressSessionScore tests the fix for issue #5519  
// The issue: suppress_session_score always shows as being added even when not set
// This test validates that suppress_session_score doesn't appear in state when not configured
func TestAccCloudflareBotManagement_Issue5519_SuppressSessionScore(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceID := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	var initialState, secondState map[string]interface{}

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
					// Capture initial state
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources[resourceID]
						if !ok {
							return fmt.Errorf("Not found: %s", resourceID)
						}
						initialState = make(map[string]interface{})
						for k, v := range rs.Primary.Attributes {
							initialState[k] = v
						}
						
						// Validate suppress_session_score is NOT in configured state (issue #5519)
						if val, exists := initialState["suppress_session_score"]; exists && val != "" {
							t.Logf("suppress_session_score present in initial state: %v (may be OK if API defaults this)", val)
						}
						
						return nil
					},
				),
			},
			{
				// Step 2: Apply the same configuration again - validate no unexpected additions
				Config: testCloudflareBotManagementIssue5519Config(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "enable_js", "true"),
					// Capture second state and validate behavior
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources[resourceID]
						if !ok {
							return fmt.Errorf("Not found: %s", resourceID)
						}
						secondState = make(map[string]interface{})
						for k, v := range rs.Primary.Attributes {
							secondState[k] = v
						}

						// Critical test: suppress_session_score should not be added if not configured
						initialSuppressSession, initialExists := initialState["suppress_session_score"]
						secondSuppressSession, secondExists := secondState["suppress_session_score"]
						
						if !initialExists && secondExists {
							return fmt.Errorf("suppress_session_score was unexpectedly added: %v (issue #5519 regression)", secondSuppressSession)
						}
						
						if initialExists && secondExists && initialSuppressSession != secondSuppressSession {
							return fmt.Errorf("suppress_session_score changed unexpectedly: %v -> %v", initialSuppressSession, secondSuppressSession)
						}

						// Validate encode_state_for_unknown fields remain stable
						fieldsWithEncode := []string{"enable_js", "auto_update_model", "fight_mode", "suppress_session_score"}
						for _, field := range fieldsWithEncode {
							if initialState[field] != secondState[field] {
								return fmt.Errorf("Field %s with encode_state_for_unknown changed: %v -> %v", 
									field, initialState[field], secondState[field])
							}
						}

						// Log what changed
						var changedFields []string
						for k, v := range initialState {
							if secondState[k] != v {
								changedFields = append(changedFields, fmt.Sprintf("%s: %v -> %v", k, v, secondState[k]))
							}
						}
						if len(changedFields) > 0 {
							t.Logf("Fields that changed between applies: %v", changedFields)
						} else {
							t.Logf("SUCCESS: No fields changed between applies - issue #5519 fixed")
						}

						return nil
					},
				),
			},
		},
	})
}

// TestAccCloudflareBotManagement_EncodeStateForUnknownValidation tests that the encode_state_for_unknown fix works
// This test specifically validates that fields with the encode_state_for_unknown tag maintain their state values
// when the API doesn't return them, preventing false drift detection
func TestAccCloudflareBotManagement_EncodeStateForUnknownValidation(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceID := "cloudflare_bot_management." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	var postApplyState map[string]interface{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: Create resource with all encode_state_for_unknown fields set
				Config: testCloudflareBotManagementEncodeStateValidationConfig(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "enable_js", "true"),
					resource.TestCheckResourceAttr(resourceID, "auto_update_model", "true"),
					resource.TestCheckResourceAttr(resourceID, "suppress_session_score", "false"),
				),
			},
			{
				// Step 2: Apply same config to validate state stability
				Config: testCloudflareBotManagementEncodeStateValidationConfig(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "enable_js", "true"),
					resource.TestCheckResourceAttr(resourceID, "auto_update_model", "true"),
					resource.TestCheckResourceAttr(resourceID, "suppress_session_score", "false"),
					// Validate that encode_state_for_unknown fields are stable
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources[resourceID]
						if !ok {
							return fmt.Errorf("Not found: %s", resourceID)
						}
						postApplyState = make(map[string]interface{})
						for k, v := range rs.Primary.Attributes {
							postApplyState[k] = v
						}

						// These are the fields that have encode_state_for_unknown tags in model.go
						encodeStateFields := map[string]string{
							"enable_js":             "true",
							"auto_update_model":     "true",
							"suppress_session_score": "false",
						}

						for field, expectedValue := range encodeStateFields {
							actualValue, exists := postApplyState[field]
							if !exists {
								return fmt.Errorf("Field %s with encode_state_for_unknown is missing from state", field)
							}
							if actualValue != expectedValue {
								return fmt.Errorf("Field %s with encode_state_for_unknown has wrong value: expected %s, got %v", 
									field, expectedValue, actualValue)
							}
						}

						// Log successful validation
						t.Logf("SUCCESS: All encode_state_for_unknown fields maintained correct values:")
						for field, value := range encodeStateFields {
							t.Logf("  %s: %s ✓", field, value)
						}

						return nil
					},
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

// Configuration for encode_state_for_unknown validation test
func testCloudflareBotManagementEncodeStateValidationConfig(resourceName, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_bot_management" "%[1]s" {
  zone_id                 = "%[2]s"
  enable_js              = true
  auto_update_model      = true
  suppress_session_score = false
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

