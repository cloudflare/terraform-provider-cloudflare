package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

var (
	currentProviderVersion = internal.PackageVersion
)

// Embed test configs
//
//go:embed testdata/v4_flat_booleans.tf
var v4FlatBooleansConfig string

//go:embed testdata/v5_flat_booleans.tf
var v5FlatBooleansConfig string

//go:embed testdata/v4_antivirus.tf
var v4AntivirusConfig string

//go:embed testdata/v5_antivirus.tf
var v5AntivirusConfig string

//go:embed testdata/v4_comprehensive.tf
var v4ComprehensiveConfig string

//go:embed testdata/v5_comprehensive.tf
var v5ComprehensiveConfig string

//go:embed testdata/v4_browser_isolation.tf
var v4BrowserIsolationConfig string

//go:embed testdata/v5_browser_isolation.tf
var v5BrowserIsolationConfig string

//go:embed testdata/v4_block_page.tf
var v4BlockPageConfig string

//go:embed testdata/v5_block_page.tf
var v5BlockPageConfig string

//go:embed testdata/v4_maxitems1_blocks.tf
var v4MaxItems1BlocksConfig string

//go:embed testdata/v5_maxitems1_blocks.tf
var v5MaxItems1BlocksConfig string

// skipIfAPIToken skips the test when using API token auth.
// Zero Trust resources don't support API token authentication.
func skipIfAPIToken(t *testing.T) {
	t.Helper()
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Skip("Skipping: zero_trust_gateway_settings does not support API token authentication")
	}
}

// TestMigrateZeroTrustGatewaySettings_V4ToV5_FlatBooleans tests flat boolean field
// transformations using the dual test case pattern (from_v4_latest and from_v5).
//
// Validates:
//   - activity_log_enabled → settings.activity_log.enabled
//   - tls_decrypt_enabled → settings.tls_decrypt.enabled
//   - protocol_detection_enabled → settings.protocol_detection.enabled
//   - Resource rename: cloudflare_teams_account → cloudflare_zero_trust_gateway_settings
func TestMigrateZeroTrustGatewaySettings_V4ToV5_FlatBooleans(t *testing.T) {
	skipIfAPIToken(t)

	certID := os.Getenv("CLOUDFLARE_GATEWAY_CERTIFICATE_ID")

	testCases := []struct {
		name              string
		version           string
		configFn          func(rnd, accountID string) string
		tlsDecryptEnabled bool
		needsCert         bool
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4FlatBooleansConfig, rnd, accountID)
			},
			tlsDecryptEnabled: false,
		},
		{
			name:      "from_v5",
			version:   currentProviderVersion,
			needsCert: true,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5FlatBooleansConfig, rnd, accountID, certID)
			},
			tlsDecryptEnabled: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if certID == "" {
				t.Skip("CLOUDFLARE_GATEWAY_CERTIFICATE_ID must be set for this sub-test.")
			}
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_gateway_settings." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
					// Gateway settings is a singleton; account may have pre-existing settings
					ExpectNonEmptyPlan: true,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
							// Verify flat booleans were moved to nested settings
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("activity_log").AtMapKey("enabled"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("tls_decrypt").AtMapKey("enabled"), knownvalue.Bool(tc.tlsDecryptEnabled)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("protocol_detection").AtMapKey("enabled"), knownvalue.Bool(false)),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustGatewaySettings_V4ToV5_Antivirus tests nested block transformation
// with field rename using the dual test case pattern.
//
// Validates:
//   - antivirus block → settings.antivirus attribute
//   - notification_settings nested block → notification_settings attribute
//   - Field rename: message → msg (v4 → v5)
//   - notification_settings[0] array element → object
func TestMigrateZeroTrustGatewaySettings_V4ToV5_Antivirus(t *testing.T) {
	skipIfAPIToken(t)

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4AntivirusConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5AntivirusConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_gateway_settings." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config:             testConfig,
					ExpectNonEmptyPlan: true,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
							// Verify antivirus fields
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("enabled_download_phase"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("enabled_upload_phase"), knownvalue.Bool(false)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("fail_closed"), knownvalue.Bool(true)),
							// Verify notification_settings with field rename (message → msg)
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("notification_settings").AtMapKey("enabled"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("notification_settings").AtMapKey("msg"), knownvalue.StringExact("File scanning in progress")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("notification_settings").AtMapKey("support_url"), knownvalue.StringExact("https://support.example.com/")),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustGatewaySettings_V4ToV5_Comprehensive tests multiple transformations
// together using the dual test case pattern.
//
// Validates:
//   - All flat boolean fields → nested under settings.*
//   - Browser isolation field rename: non_identity_browser_isolation_enabled → non_identity_enabled
//   - fips MaxItems:1 block → settings.fips attribute
//   - body_scanning MaxItems:1 block → settings.body_scanning attribute
//   - antivirus with nested notification_settings + message → msg rename
//
// Note: block_page and extended_email_matching are excluded due to pre-existing v5 provider
// drift on API-computed fields (mode, version, suppress_footer, etc.) that are not in config.
func TestMigrateZeroTrustGatewaySettings_V4ToV5_Comprehensive(t *testing.T) {
	skipIfAPIToken(t)

	certID := os.Getenv("CLOUDFLARE_GATEWAY_CERTIFICATE_ID")

	testCases := []struct {
		name              string
		version           string
		configFn          func(rnd, accountID string) string
		tlsDecryptEnabled bool
		needsCert         bool
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4ComprehensiveConfig, rnd, accountID)
			},
			tlsDecryptEnabled: false,
		},
		{
			name:      "from_v5",
			version:   currentProviderVersion,
			needsCert: true,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5ComprehensiveConfig, rnd, accountID, certID)
			},
			tlsDecryptEnabled: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if certID == "" {
				t.Skip("CLOUDFLARE_GATEWAY_CERTIFICATE_ID must be set for this sub-test.")
			}
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_gateway_settings." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config:             testConfig,
					ExpectNonEmptyPlan: true,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),

							// Flat boolean → nested
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("activity_log").AtMapKey("enabled"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("tls_decrypt").AtMapKey("enabled"), knownvalue.Bool(tc.tlsDecryptEnabled)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("protocol_detection").AtMapKey("enabled"), knownvalue.Bool(false)),

							// Browser isolation with field rename
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("browser_isolation").AtMapKey("url_browser_isolation_enabled"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("browser_isolation").AtMapKey("non_identity_enabled"), knownvalue.Bool(false)),

							// fips
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("fips").AtMapKey("tls"), knownvalue.Bool(true)),

							// body_scanning
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("body_scanning").AtMapKey("inspection_mode"), knownvalue.StringExact("deep")),

							// antivirus with field rename (message → msg)
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("enabled_download_phase"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("enabled_upload_phase"), knownvalue.Bool(false)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("fail_closed"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("notification_settings").AtMapKey("enabled"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("notification_settings").AtMapKey("msg"), knownvalue.StringExact("Scanning")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus").AtMapKey("notification_settings").AtMapKey("support_url"), knownvalue.StringExact("https://support.example.com/")),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustGatewaySettings_V4ToV5_Minimal tests that fields absent from the
// user's config are null in migrated state (null-cleanup behavior).
//
// Validates:
//   - Fields in config are correctly set after migration
//   - Fields NOT in config (block_page, fips, antivirus, etc.) are null — the migration
//     must not carry over v4 API-returned values for blocks the user never configured
func TestMigrateZeroTrustGatewaySettings_V4ToV5_Minimal(t *testing.T) {
	skipIfAPIToken(t)

	testCases := []struct {
		name              string
		version           string
		configFn          func(rnd, accountID string) string
		tlsDecryptEnabled bool
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4FlatBooleansConfig, rnd, accountID)
			},
			tlsDecryptEnabled: false,
		},
		{
			// from_v5 uses a truly minimal config (no cert, tls_decrypt=false) so that
			// the null-cleanup assertions below hold — this test's purpose is to verify
			// that fields absent from the user's config are null after migration.
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(`
resource "cloudflare_zero_trust_gateway_settings" "%[1]s" {
  account_id = "%[2]s"
  settings = {
    activity_log = {
      enabled = true
    }
    tls_decrypt = {
      enabled = false
    }
    protocol_detection = {
      enabled = false
    }
  }
}`, rnd, accountID)
			},
			tlsDecryptEnabled: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_gateway_settings." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config:             testConfig,
					ExpectNonEmptyPlan: true,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
							// Flat booleans in config are present and correct
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("activity_log").AtMapKey("enabled"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("tls_decrypt").AtMapKey("enabled"), knownvalue.Bool(tc.tlsDecryptEnabled)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("protocol_detection").AtMapKey("enabled"), knownvalue.Bool(false)),
							// Fields NOT in user config must be null — migration must not carry
							// over v4 API-returned values for blocks the user never configured
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page"), knownvalue.Null()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("body_scanning"), knownvalue.Null()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("extended_email_matching"), knownvalue.Null()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("fips"), knownvalue.Null()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus"), knownvalue.Null()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("custom_certificate"), knownvalue.Null()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("certificate"), knownvalue.Null()),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustGatewaySettings_V4ToV5_BrowserIsolation tests browser isolation
// field rename using the dual test case pattern.
//
// Validates:
//   - url_browser_isolation_enabled → settings.browser_isolation.url_browser_isolation_enabled
//   - non_identity_browser_isolation_enabled → settings.browser_isolation.non_identity_enabled
func TestMigrateZeroTrustGatewaySettings_V4ToV5_BrowserIsolation(t *testing.T) {
	skipIfAPIToken(t)

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4BrowserIsolationConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5BrowserIsolationConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_gateway_settings." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config:             testConfig,
					ExpectNonEmptyPlan: true,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
							// Key test: non_identity_browser_isolation_enabled → non_identity_enabled
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("browser_isolation").AtMapKey("url_browser_isolation_enabled"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("browser_isolation").AtMapKey("non_identity_enabled"), knownvalue.Bool(false)),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustGatewaySettings_V4ToV5_BlockPage tests block_page MaxItems:1 block
// conversion using the dual test case pattern.
//
// Validates:
//   - block_page block → settings.block_page attribute
//   - All user-configurable fields preserved after migration
//
// Note: This test uses ExpectNonEmptyPlan on the migration step because the v5 provider
// has a known drift issue on API-computed block_page fields (mode, version, suppress_footer,
// target_uri, include_context) that are returned by the API but not present in user config.
// This is a pre-existing v5 provider issue, not a migration bug.
func TestMigrateZeroTrustGatewaySettings_V4ToV5_BlockPage(t *testing.T) {
	skipIfAPIToken(t)

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4BlockPageConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5BlockPageConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_gateway_settings." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
					// block_page computed fields (mode, version, etc.) cause drift even in pure v5
					ExpectNonEmptyPlan: true,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config:             testConfig,
					ExpectNonEmptyPlan: true,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					// Use manual migration step (not MigrationV2TestStep) to allow ExpectNonEmptyPlan
					// due to the pre-existing block_page computed-field drift in the v5 provider.
					{
						PreConfig: func() {
							acctest.WriteOutConfig(t, testConfig, tmpDir)
							acctest.RunMigrationV2Command(t, testConfig, tmpDir, sourceVer, targetVer)
						},
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ExpectNonEmptyPlan:       true,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
							// Verify user-configurable block_page fields were migrated correctly
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("enabled"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("name"), knownvalue.StringExact(rnd)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("footer_text"), knownvalue.StringExact("Contact IT")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("header_text"), knownvalue.StringExact("Blocked")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("logo_path"), knownvalue.StringExact("https://example.com/logo.png")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("background_color"), knownvalue.StringExact("#FF0000")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("mailto_address"), knownvalue.StringExact("security@example.com")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("block_page").AtMapKey("mailto_subject"), knownvalue.StringExact("Access Request")),
							// Other blocks not in config are null
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("fips"), knownvalue.Null()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus"), knownvalue.Null()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("body_scanning"), knownvalue.Null()),
						},
					},
				},
			})
		})
	}
}

// TestMigrateZeroTrustGatewaySettings_V4ToV5_MultipleMaxItems1 tests that multiple
// MaxItems:1 blocks are all correctly converted to nested attributes using the dual
// test case pattern.
//
// Validates:
//   - fips MaxItems:1 block → settings.fips attribute
//   - body_scanning MaxItems:1 block → settings.body_scanning attribute
func TestMigrateZeroTrustGatewaySettings_V4ToV5_MultipleMaxItems1(t *testing.T) {
	skipIfAPIToken(t)

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4MaxItems1BlocksConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5MaxItems1BlocksConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_gateway_settings." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config:             testConfig,
					ExpectNonEmptyPlan: true,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
							// Verify all MaxItems:1 blocks are attributes under settings
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("fips").AtMapKey("tls"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("body_scanning").AtMapKey("inspection_mode"), knownvalue.StringExact("deep")),
							// Other blocks not in config are null
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("antivirus"), knownvalue.Null()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("extended_email_matching"), knownvalue.Null()),
						},
					),
				},
			})
		})
	}
}
