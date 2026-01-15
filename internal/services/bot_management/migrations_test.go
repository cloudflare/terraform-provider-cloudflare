package bot_management_test

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

// TestMigrateBotManagement_V4ToV5_Basic tests migration with basic bot management fields
func TestMigrateBotManagement_V4ToV5_Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	tmpDir := t.TempDir()

	// V4 config with basic fields
	v4Config := fmt.Sprintf(`
resource "cloudflare_bot_management" "%[1]s" {
  zone_id           = "%[2]s"
  enable_js         = true
  auto_update_model = true
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
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
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Verify resource exists with correct zone_id
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				// Verify basic fields are preserved
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_update_model"), knownvalue.Bool(true)),
				// Verify ID is present
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateBotManagement_V4ToV5_AIBotsProtection tests migration with AI bots protection
func TestMigrateBotManagement_V4ToV5_AIBotsProtection(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	tmpDir := t.TempDir()

	// V4 config with AI bots protection
	v4Config := fmt.Sprintf(`
resource "cloudflare_bot_management" "%[1]s" {
  zone_id               = "%[2]s"
  enable_js             = true
  auto_update_model     = true
  ai_bots_protection    = "block"
  suppress_session_score = false
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
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
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_bots_protection"), knownvalue.StringExact("block")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_update_model"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("suppress_session_score"), knownvalue.Bool(false)),
			}),
		},
	})
}

// TestMigrateBotManagement_V4ToV5_AIBotsProtectionDisabled tests migration with AI bots protection disabled
func TestMigrateBotManagement_V4ToV5_AIBotsProtectionDisabled(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	tmpDir := t.TempDir()

	// V4 config with AI bots protection disabled
	v4Config := fmt.Sprintf(`
resource "cloudflare_bot_management" "%[1]s" {
  zone_id            = "%[2]s"
  enable_js          = true
  auto_update_model  = true
  ai_bots_protection = "disabled"
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
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
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_bots_protection"), knownvalue.StringExact("disabled")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(true)),
			}),
		},
	})
}

// TestMigrateBotManagement_V4ToV5_SessionScoreSuppress tests migration with session score suppression
func TestMigrateBotManagement_V4ToV5_SessionScoreSuppress(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	tmpDir := t.TempDir()

	// V4 config with session score suppression
	v4Config := fmt.Sprintf(`
resource "cloudflare_bot_management" "%[1]s" {
  zone_id                = "%[2]s"
  enable_js              = true
  auto_update_model      = true
  suppress_session_score = true
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
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
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_update_model"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("suppress_session_score"), knownvalue.Bool(true)),
			}),
		},
	})
}

// TestMigrateBotManagement_V4ToV5_MultipleFields tests migration with multiple fields
func TestMigrateBotManagement_V4ToV5_MultipleFields(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	tmpDir := t.TempDir()

	// V4 config with multiple available fields
	v4Config := fmt.Sprintf(`
resource "cloudflare_bot_management" "%[1]s" {
  zone_id                = "%[2]s"
  enable_js              = true
  auto_update_model      = true
  suppress_session_score = false
  ai_bots_protection     = "block"
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
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
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_update_model"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("suppress_session_score"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_bots_protection"), knownvalue.StringExact("block")),
			}),
		},
	})
}

// TestMigrateBotManagement_V4ToV5_EnableJSFalse tests migration with enable_js=false
func TestMigrateBotManagement_V4ToV5_EnableJSFalse(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	tmpDir := t.TempDir()

	// V4 config with enable_js=false
	v4Config := fmt.Sprintf(`
resource "cloudflare_bot_management" "%[1]s" {
  zone_id           = "%[2]s"
  enable_js         = false
  auto_update_model = true
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
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
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_update_model"), knownvalue.Bool(true)),
			}),
		},
	})
}

// TestMigrateBotManagement_V4ToV5_AutoUpdateModelFalse tests migration with auto_update_model=false
func TestMigrateBotManagement_V4ToV5_AutoUpdateModelFalse(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_bot_management." + rnd
	tmpDir := t.TempDir()

	// V4 config with auto_update_model=false
	v4Config := fmt.Sprintf(`
resource "cloudflare_bot_management" "%[1]s" {
  zone_id           = "%[2]s"
  enable_js         = true
  auto_update_model = false
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
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
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_update_model"), knownvalue.Bool(false)),
			}),
		},
	})
}
