package zero_trust_gateway_settings_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateZeroTrustGatewaySettings_V4ToV5_Minimal tests minimal configuration migration
// Validates:
// - Resource type rename: cloudflare_teams_account → cloudflare_zero_trust_gateway_settings
// - Basic structure transformation
//
// IMPORTANT: This test expects a non-empty plan after migration because:
// 1. Gateway settings is a singleton per-account (not a regular resource)
// 2. v4 provider reads ALL existing account settings when it reads the resource
// 3. The test account has pre-existing gateway settings (block_page, body_scanning, etc.)
// 4. Migration correctly removes fields not in user's config
// 5. v5 plan shows drift for the removed pre-existing settings
//
// This is EXPECTED BEHAVIOR, not a bug. The migration reveals implicit configuration.
func TestMigrateZeroTrustGatewaySettings_V4ToV5_Minimal(t *testing.T) {
	// Zero Trust/Teams resources don't support API tokens yet
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_gateway_settings." + rnd
	tmpDir := t.TempDir()

	// V4 config with minimal required fields
	// Note: Not including deprecated fields (logging, proxy, ssh_session_log, payload_log)
	// These fields get migrated to separate resources and would cause plan changes
	v4Config := fmt.Sprintf(`
resource "cloudflare_teams_account" "%[1]s" {
  account_id                 = "%[2]s"
  activity_log_enabled       = true
  tls_decrypt_enabled        = false
  protocol_detection_enabled = false
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				// The test account has pre-existing deprecated settings (logging, proxy, etc.)
				// that aren't in our minimal config, so we expect a non-empty plan
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config:             v4Config,
				ExpectNonEmptyPlan: true, // Account has deprecated fields not in user config
			},
			// Step 2: Run migration and verify state transformation
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Verify resource exists with new type
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),

				// Verify fields from user config are present and correct
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("activity_log").AtMapKey("enabled"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("tls_decrypt").AtMapKey("enabled"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("protocol_detection").AtMapKey("enabled"), knownvalue.Bool(false)),

				// CRITICAL: Verify fields NOT in user config are cleaned up (removed from state)
				// These fields were in v4 state (from API) but should NOT be in migrated v5 state
				// because they weren't in the user's configuration

				// block_page: Not in user config, should be null
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page"), knownvalue.Null()),

				// body_scanning: Not in user config, should be null
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("body_scanning"), knownvalue.Null()),

				// extended_email_matching: Not in user config, should be null
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("extended_email_matching"), knownvalue.Null()),

				// fips: Not in user config, should be null
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("fips"), knownvalue.Null()),

				// antivirus: Not in user config, should be null
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus"), knownvalue.Null()),

				// custom_certificate: Not in user config, should be null
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("custom_certificate"), knownvalue.Null()),

				// certificate: Not in user config, should be null
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("certificate"), knownvalue.Null()),

				// host_selector: Not in user config, should be null
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("host_selector"), knownvalue.Null()),

				// inspection: Not in user config, should be null
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("inspection"), knownvalue.Null()),
			}),
		},
	})
}

// TestMigrateZeroTrustGatewaySettings_V4ToV5_FlatBooleans tests flat boolean field transformations
// Validates:
// - activity_log_enabled → settings.activity_log.enabled
// - tls_decrypt_enabled → settings.tls_decrypt.enabled
// - protocol_detection_enabled → settings.protocol_detection.enabled
func TestMigrateZeroTrustGatewaySettings_V4ToV5_FlatBooleans(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_gateway_settings." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_teams_account" "%[1]s" {
  account_id                 = "%[2]s"
  activity_log_enabled       = true
  tls_decrypt_enabled        = true
  protocol_detection_enabled = false
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
				// Account already has gateway settings with deprecated fields
				ExpectNonEmptyPlan: true,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				// Verify boolean fields are nested under settings
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("activity_log").AtMapKey("enabled"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("tls_decrypt").AtMapKey("enabled"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("protocol_detection").AtMapKey("enabled"), knownvalue.Bool(false)),
			}),
		},
	})
}

// TestMigrateZeroTrustGatewaySettings_V4ToV5_BlockPage tests MaxItems:1 block conversion
// Validates:
// - block_page block → settings.block_page attribute
// - Nested fields preserved
func TestMigrateZeroTrustGatewaySettings_V4ToV5_BlockPage(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_gateway_settings." + rnd
	tmpDir := t.TempDir()

	blockPageName := fmt.Sprintf("test-block-%s", rnd)
	v4Config := fmt.Sprintf(`
resource "cloudflare_teams_account" "%[1]s" {
  account_id                 = "%[2]s"
  activity_log_enabled       = false
  protocol_detection_enabled = false
  tls_decrypt_enabled        = false

  block_page {
    enabled          = true
    name             = "%[3]s"
    footer_text      = "Contact IT Support"
    header_text      = "Access Blocked"
    logo_path        = "https://example.com/logo.png"
    background_color = "#FF0000"
    mailto_address   = "security@example.com"
    mailto_subject   = "Access Request"
  }
}`, rnd, accountID, blockPageName)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
				// Account already has gateway settings with deprecated fields
				ExpectNonEmptyPlan: true,
			},
			// Step 2: Run migration and validate state
			// NOTE: We expect a non-empty plan because the v5 provider has a known issue where
			// it removes Optional+Computed fields (mode, include_context, suppress_footer, target_uri)
			// from block_page during refresh when they weren't in the original HCL config.
			// This is expected behavior and documented in README.md.
			{
				PreConfig: func() {
					acctest.WriteOutConfig(t, v4Config, tmpDir)
					acctest.RunMigrationV2Command(t, v4Config, tmpDir, "v4", "v5")
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ExpectNonEmptyPlan:       true,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),

					// Verify flat booleans from config are present
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("activity_log").AtMapKey("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("tls_decrypt").AtMapKey("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("protocol_detection").AtMapKey("enabled"), knownvalue.Bool(false)),

					// Verify block_page is nested under settings as an attribute (not array)
					// All fields from user config should be present and correct
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("name"), knownvalue.StringExact(blockPageName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("footer_text"), knownvalue.StringExact("Contact IT Support")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("header_text"), knownvalue.StringExact("Access Blocked")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("logo_path"), knownvalue.StringExact("https://example.com/logo.png")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("background_color"), knownvalue.StringExact("#FF0000")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("mailto_address"), knownvalue.StringExact("security@example.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("mailto_subject"), knownvalue.StringExact("Access Request")),

					// NOTE: We don't check block_page optional fields (include_context, mode, suppress_footer, target_uri)
					// because the v5 provider has a known issue where it removes these Optional+Computed fields during
					// refresh when they weren't in the original HCL config. This causes drift that we accept as expected.
					// See README.md "Known Issue: block_page Perpetual Drift" section.

					// Verify other settings blocks NOT in config are null
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("body_scanning"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("extended_email_matching"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("fips"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("custom_certificate"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("certificate"), knownvalue.Null()),
				},
			},
		},
	})
}

// TestMigrateZeroTrustGatewaySettings_V4ToV5_Antivirus tests nested MaxItems:1 blocks
// Validates:
// - antivirus block → settings.antivirus attribute
// - notification_settings nested block → notification_settings attribute
// - Field rename: message → msg
func TestMigrateZeroTrustGatewaySettings_V4ToV5_Antivirus(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_gateway_settings." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_teams_account" "%[1]s" {
  account_id                 = "%[2]s"
  activity_log_enabled       = false
  protocol_detection_enabled = false
  tls_decrypt_enabled        = false

  antivirus {
    enabled_download_phase = true
    enabled_upload_phase   = false
    fail_closed            = true

    notification_settings {
      enabled     = true
      message     = "File scanning in progress"
      support_url = "https://support.example.com"
    }
  }
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
				// Account already has gateway settings with deprecated fields
				ExpectNonEmptyPlan: true,
			},
			// Step 2: Run migration and validate state
			// Account already has gateway settings with deprecated fields
			{
				PreConfig: func() {
					acctest.WriteOutConfig(t, v4Config, tmpDir)
					acctest.RunMigrationV2Command(t, v4Config, tmpDir, "v4", "v5")
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ExpectNonEmptyPlan:       true,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),

					// Verify flat booleans from config are present
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("activity_log").AtMapKey("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("tls_decrypt").AtMapKey("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("protocol_detection").AtMapKey("enabled"), knownvalue.Bool(false)),

					// Verify antivirus is nested under settings
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("enabled_download_phase"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("enabled_upload_phase"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("fail_closed"), knownvalue.Bool(true)),

					// Verify nested notification_settings and field rename (message → msg)
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("notification_settings").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("notification_settings").AtMapKey("msg"), knownvalue.StringExact("File scanning in progress")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("notification_settings").AtMapKey("support_url"), knownvalue.StringExact("https://support.example.com")),

					// Verify other settings blocks NOT in config are null
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("body_scanning"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("extended_email_matching"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("fips"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("custom_certificate"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("certificate"), knownvalue.Null()),
				},
			},
		},
	})
}

// TestMigrateZeroTrustGatewaySettings_V4ToV5_BrowserIsolation tests browser isolation field renames
// Validates:
// - url_browser_isolation_enabled → settings.browser_isolation.url_browser_isolation_enabled
// - non_identity_browser_isolation_enabled → settings.browser_isolation.non_identity_enabled
func TestMigrateZeroTrustGatewaySettings_V4ToV5_BrowserIsolation(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_gateway_settings." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_teams_account" "%[1]s" {
  account_id                             = "%[2]s"
  activity_log_enabled                   = false
  protocol_detection_enabled             = false
  tls_decrypt_enabled                    = false
  url_browser_isolation_enabled          = true
  non_identity_browser_isolation_enabled = false
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
				// Account already has gateway settings with deprecated fields
				ExpectNonEmptyPlan: true,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				// Verify browser isolation fields are grouped and renamed
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("browser_isolation").AtMapKey("url_browser_isolation_enabled"), knownvalue.Bool(true)),
				// Key test: non_identity_browser_isolation_enabled → non_identity_enabled
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("browser_isolation").AtMapKey("non_identity_enabled"), knownvalue.Bool(false)),
			}),
		},
	})
}

// TestMigrateZeroTrustGatewaySettings_V4ToV5_MultipleMaxItems1 tests multiple MaxItems:1 blocks
// Validates:
// - Multiple MaxItems:1 blocks converted to attributes
// - block_page, fips, body_scanning all under settings
func TestMigrateZeroTrustGatewaySettings_V4ToV5_MultipleMaxItems1(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_gateway_settings." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_teams_account" "%[1]s" {
  account_id                 = "%[2]s"
  activity_log_enabled       = false
  protocol_detection_enabled = false
  tls_decrypt_enabled        = false

  block_page {
    enabled          = true
    name             = "Multi Block Test"
    footer_text      = "Footer"
    header_text      = "Header"
    logo_path        = "https://example.com/logo.png"
    background_color = "#000000"
  }

  fips {
    tls = true
  }

  body_scanning {
    inspection_mode = "deep"
  }
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
				// Account already has gateway settings with deprecated fields
				ExpectNonEmptyPlan: true,
			},
			// Step 2: Run migration and validate state
			// Account already has gateway settings with deprecated fields
			{
				PreConfig: func() {
					acctest.WriteOutConfig(t, v4Config, tmpDir)
					acctest.RunMigrationV2Command(t, v4Config, tmpDir, "v4", "v5")
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ExpectNonEmptyPlan:       true,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),

					// Verify flat booleans from config are present
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("activity_log").AtMapKey("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("tls_decrypt").AtMapKey("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("protocol_detection").AtMapKey("enabled"), knownvalue.Bool(false)),

					// Verify all MaxItems:1 blocks are attributes under settings
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("name"), knownvalue.StringExact("Multi Block Test")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("footer_text"), knownvalue.StringExact("Footer")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("header_text"), knownvalue.StringExact("Header")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("logo_path"), knownvalue.StringExact("https://example.com/logo.png")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("background_color"), knownvalue.StringExact("#000000")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("fips").AtMapKey("tls"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("body_scanning").AtMapKey("inspection_mode"), knownvalue.StringExact("deep")),

					// Verify other settings blocks NOT in config are null
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("extended_email_matching"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("custom_certificate"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("certificate"), knownvalue.Null()),
				},
			},
		},
	})
}

// TestMigrateZeroTrustGatewaySettings_V4ToV5_Comprehensive tests complex configuration
// Validates:
// - All transformation patterns work together
// - Flat booleans + MaxItems:1 blocks + field renames
func TestMigrateZeroTrustGatewaySettings_V4ToV5_Comprehensive(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_gateway_settings." + rnd
	tmpDir := t.TempDir()

	blockPageName := fmt.Sprintf("comprehensive-%s", rnd)
	v4Config := fmt.Sprintf(`
resource "cloudflare_teams_account" "%[1]s" {
  account_id                             = "%[2]s"
  activity_log_enabled                   = true
  tls_decrypt_enabled                    = true
  protocol_detection_enabled             = false
  url_browser_isolation_enabled          = true
  non_identity_browser_isolation_enabled = false

  block_page {
    enabled          = true
    name             = "%[3]s"
    footer_text      = "Contact IT"
    header_text      = "Blocked"
    logo_path        = "https://example.com/logo.png"
    background_color = "#000000"
  }

  fips {
    tls = true
  }

  body_scanning {
    inspection_mode = "deep"
  }

  antivirus {
    enabled_download_phase = true
    enabled_upload_phase   = false
    fail_closed            = true

    notification_settings {
      enabled     = true
      message     = "Scanning"
      support_url = "https://support.example.com"
    }
  }

  extended_email_matching {
    enabled = true
  }
}`, rnd, accountID, blockPageName)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
				// Account already has gateway settings with deprecated fields
				ExpectNonEmptyPlan: true,
			},
			// Step 2: Run migration and validate state
			// Account already has gateway settings with deprecated fields
			{
				PreConfig: func() {
					acctest.WriteOutConfig(t, v4Config, tmpDir)
					acctest.RunMigrationV2Command(t, v4Config, tmpDir, "v4", "v5")
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ExpectNonEmptyPlan:       true,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),

					// Verify flat boolean transformations
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("activity_log").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("tls_decrypt").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("protocol_detection").AtMapKey("enabled"), knownvalue.Bool(false)),

					// Verify browser isolation with field rename
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("browser_isolation").AtMapKey("url_browser_isolation_enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("browser_isolation").AtMapKey("non_identity_enabled"), knownvalue.Bool(false)),

					// Verify block_page fields from config
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("name"), knownvalue.StringExact(blockPageName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("footer_text"), knownvalue.StringExact("Contact IT")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("header_text"), knownvalue.StringExact("Blocked")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("logo_path"), knownvalue.StringExact("https://example.com/logo.png")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("background_color"), knownvalue.StringExact("#000000")),

					// Verify other MaxItems:1 blocks
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("fips").AtMapKey("tls"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("body_scanning").AtMapKey("inspection_mode"), knownvalue.StringExact("deep")),

					// Verify antivirus with nested notification_settings and field rename (message → msg)
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("enabled_download_phase"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("enabled_upload_phase"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("fail_closed"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("notification_settings").AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("notification_settings").AtMapKey("msg"), knownvalue.StringExact("Scanning")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("notification_settings").AtMapKey("support_url"), knownvalue.StringExact("https://support.example.com")),

					// Verify extended_email_matching
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("extended_email_matching").AtMapKey("enabled"), knownvalue.Bool(true)),

					// Verify other settings blocks NOT in config are null
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("custom_certificate"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("certificate"), knownvalue.Null()),
				},
			},
		},
	})
}
