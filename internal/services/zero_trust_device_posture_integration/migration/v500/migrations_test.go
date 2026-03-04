package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Embed test configs
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_deprecated_name.tf
var v4DeprecatedNameConfig string

//go:embed testdata/v5_deprecated_name.tf
var v5DeprecatedNameConfig string

//go:embed testdata/v4_current_name.tf
var v4CurrentNameConfig string

//go:embed testdata/v5_current_name.tf
var v5CurrentNameConfig string

//go:embed testdata/v4_with_identifier.tf
var v4WithIdentifierConfig string

//go:embed testdata/v5_with_identifier.tf
var v5WithIdentifierConfig string

//go:embed testdata/v4_comprehensive.tf
var v4ComprehensiveConfig string

//go:embed testdata/v5_comprehensive.tf
var v5ComprehensiveConfig string

// TestMigrateZeroTrustDevicePostureIntegration_V4ToV5_Basic tests basic field migrations
// with DUAL test cases: from v4 legacy provider and from v5 current provider
//
// Tests transformation of:
// - Resource rename: cloudflare_device_posture_integration → cloudflare_zero_trust_device_posture_integration
// - Config structure: TypeList block → SingleNested attribute
// - interval field: Optional → Required (default "24h" if missing)
// - identifier field: Removed in v5
func TestMigrateZeroTrustDevicePostureIntegration_V4ToV5_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name, apiURL, clientID, clientSecret, customerID string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name, apiURL, clientID, clientSecret, customerID string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, accountID, name, apiURL, clientID, clientSecret, customerID)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name, apiURL, clientID, clientSecret, customerID string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, accountID, name, apiURL, clientID, clientSecret, customerID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
			originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
			if originalToken != "" {
				os.Unsetenv("CLOUDFLARE_API_TOKEN")
				defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
			}

			// Test setup - requires CrowdStrike credentials
			clientID := os.Getenv("CLOUDFLARE_CROWDSTRIKE_CLIENT_ID")
			clientSecret := os.Getenv("CLOUDFLARE_CROWDSTRIKE_CLIENT_SECRET")
			apiURL := os.Getenv("CLOUDFLARE_CROWDSTRIKE_API_URL")
			customerID := os.Getenv("CLOUDFLARE_CROWDSTRIKE_CUSTOMER_ID")
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

			rnd := utils.GenerateRandomResourceName()
			name := "tf-migrate-" + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name, apiURL, clientID, clientSecret, customerID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			// For v5 tests, use local provider; for v4 tests, use external provider
			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				// Use local v5 provider (has GetSchemaVersion, will create version=1 state)
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				// Use external v4 provider (will create version=0 state)
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
					acctest.TestAccPreCheck_CrowdStrike(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify critical identifier fields
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("type"),
								knownvalue.StringExact("crowdstrike_s2s"),
							),

							// Verify interval field (optional→required transformation)
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("interval"),
								knownvalue.StringExact("24h"),
							),

							// Verify config fields (array→pointer transformation)
							// Config should be converted from block [{...}] to attribute {...}
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("config").AtMapKey("api_url"),
								knownvalue.StringExact(apiURL),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("config").AtMapKey("client_id"),
								knownvalue.StringExact(clientID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("config").AtMapKey("customer_id"),
								knownvalue.StringExact(customerID),
							),

							// Verify ID exists (API-assigned)
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("id"),
								knownvalue.NotNull(),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustDevicePostureIntegration_V4ToV5_DeprecatedName tests migration
// from deprecated resource name with DUAL test cases.
//
// Tests:
// - Resource rename: cloudflare_device_posture_integration → cloudflare_zero_trust_device_posture_integration
// - Config block → attribute transformation
// - Empty string → null conversion for optional fields
func TestMigrateZeroTrustDevicePostureIntegration_V4ToV5_DeprecatedName(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name, apiURL, clientID, clientSecret, customerID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name, apiURL, clientID, clientSecret, customerID string) string {
				return fmt.Sprintf(v4DeprecatedNameConfig, rnd, accountID, name, apiURL, clientID, clientSecret, customerID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name, apiURL, clientID, clientSecret, customerID string) string {
				return fmt.Sprintf(v5DeprecatedNameConfig, rnd, accountID, name, apiURL, clientID, clientSecret, customerID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
			if originalToken != "" {
				os.Unsetenv("CLOUDFLARE_API_TOKEN")
				defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
			}

			clientID := os.Getenv("CLOUDFLARE_CROWDSTRIKE_CLIENT_ID")
			clientSecret := os.Getenv("CLOUDFLARE_CROWDSTRIKE_CLIENT_SECRET")
			apiURL := os.Getenv("CLOUDFLARE_CROWDSTRIKE_API_URL")
			customerID := os.Getenv("CLOUDFLARE_CROWDSTRIKE_CUSTOMER_ID")
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-migrate-deprecated-%s", rnd)
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name, apiURL, clientID, clientSecret, customerID)
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
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
					acctest.TestAccPreCheck_CrowdStrike(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify resource was renamed
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("type"),
								knownvalue.StringExact("crowdstrike_s2s"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("interval"),
								knownvalue.StringExact("24h"),
							),
							// Config should be converted from block to attribute
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("config").AtMapKey("api_url"),
								knownvalue.StringExact(apiURL),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("config").AtMapKey("client_id"),
								knownvalue.StringExact(clientID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("config").AtMapKey("customer_id"),
								knownvalue.StringExact(customerID),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustDevicePostureIntegration_V4ToV5_CurrentName tests migration
// from current (non-deprecated) resource name with DUAL test cases.
//
// Tests:
// - No resource rename (already using current name)
// - Config block → attribute transformation
func TestMigrateZeroTrustDevicePostureIntegration_V4ToV5_CurrentName(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name, apiURL, clientID, clientSecret, customerID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name, apiURL, clientID, clientSecret, customerID string) string {
				return fmt.Sprintf(v4CurrentNameConfig, rnd, accountID, name, apiURL, clientID, clientSecret, customerID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name, apiURL, clientID, clientSecret, customerID string) string {
				return fmt.Sprintf(v5CurrentNameConfig, rnd, accountID, name, apiURL, clientID, clientSecret, customerID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
			if originalToken != "" {
				os.Unsetenv("CLOUDFLARE_API_TOKEN")
				defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
			}

			clientID := os.Getenv("CLOUDFLARE_CROWDSTRIKE_CLIENT_ID")
			clientSecret := os.Getenv("CLOUDFLARE_CROWDSTRIKE_CLIENT_SECRET")
			apiURL := os.Getenv("CLOUDFLARE_CROWDSTRIKE_API_URL")
			customerID := os.Getenv("CLOUDFLARE_CROWDSTRIKE_CUSTOMER_ID")
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-migrate-current-%s", rnd)
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name, apiURL, clientID, clientSecret, customerID)
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
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
					acctest.TestAccPreCheck_CrowdStrike(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Resource name should stay the same (already using current name)
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("type"),
								knownvalue.StringExact("crowdstrike_s2s"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("interval"),
								knownvalue.StringExact("24h"),
							),
							// Config should be converted from block to attribute
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("config").AtMapKey("api_url"),
								knownvalue.StringExact(apiURL),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("config").AtMapKey("client_id"),
								knownvalue.StringExact(clientID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("config").AtMapKey("customer_id"),
								knownvalue.StringExact(customerID),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustDevicePostureIntegration_V4ToV5_WithIdentifier tests migration
// when identifier field exists with DUAL test cases.
//
// Tests:
// - Identifier field removal
// - Resource rename
// - Config block → attribute transformation
func TestMigrateZeroTrustDevicePostureIntegration_V4ToV5_WithIdentifier(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name, apiURL, clientID, clientSecret, customerID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name, apiURL, clientID, clientSecret, customerID string) string {
				return fmt.Sprintf(v4WithIdentifierConfig, rnd, accountID, name, apiURL, clientID, clientSecret, customerID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name, apiURL, clientID, clientSecret, customerID string) string {
				return fmt.Sprintf(v5WithIdentifierConfig, rnd, accountID, name, apiURL, clientID, clientSecret, customerID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
			if originalToken != "" {
				os.Unsetenv("CLOUDFLARE_API_TOKEN")
				defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
			}

			clientID := os.Getenv("CLOUDFLARE_CROWDSTRIKE_CLIENT_ID")
			clientSecret := os.Getenv("CLOUDFLARE_CROWDSTRIKE_CLIENT_SECRET")
			apiURL := os.Getenv("CLOUDFLARE_CROWDSTRIKE_API_URL")
			customerID := os.Getenv("CLOUDFLARE_CROWDSTRIKE_CUSTOMER_ID")
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-migrate-identifier-%s", rnd)
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name, apiURL, clientID, clientSecret, customerID)
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
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
					acctest.TestAccPreCheck_CrowdStrike(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify resource renamed and identifier removed
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("type"),
								knownvalue.StringExact("crowdstrike_s2s"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("interval"),
								knownvalue.StringExact("24h"),
							),
							// Config should be converted from block to attribute
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("config").AtMapKey("api_url"),
								knownvalue.StringExact(apiURL),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("config").AtMapKey("client_id"),
								knownvalue.StringExact(clientID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd,
								tfjsonpath.New("config").AtMapKey("customer_id"),
								knownvalue.StringExact(customerID),
							),
							// Note: identifier field should be removed from state (no check for it)
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustDevicePostureIntegration_V4ToV5_Comprehensive is a comprehensive test
// covering all migration scenarios with DUAL test cases.
//
// Tests:
// - Deprecated resource name rename
// - Current resource name (no rename)
// - Different interval values (1h, 6h, 24h)
// - Identifier field removal
// - Config block → attribute conversion
// - Empty string → null transformation
// - Multiple resources in single migration run
func TestMigrateZeroTrustDevicePostureIntegration_V4ToV5_Comprehensive(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd1, rnd2, rnd3, accountID, name1, name2, name3, apiURL, clientID, clientSecret, customerID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd1, rnd2, rnd3, accountID, name1, name2, name3, apiURL, clientID, clientSecret, customerID string) string {
				return fmt.Sprintf(v4ComprehensiveConfig, rnd1, rnd2, rnd3, accountID, name1, name2, name3, apiURL, clientID, clientSecret, customerID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd1, rnd2, rnd3, accountID, name1, name2, name3, apiURL, clientID, clientSecret, customerID string) string {
				return fmt.Sprintf(v5ComprehensiveConfig, rnd1, rnd2, rnd3, accountID, name1, name2, name3, apiURL, clientID, clientSecret, customerID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
			if originalToken != "" {
				os.Unsetenv("CLOUDFLARE_API_TOKEN")
				defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
			}

			clientID := os.Getenv("CLOUDFLARE_CROWDSTRIKE_CLIENT_ID")
			clientSecret := os.Getenv("CLOUDFLARE_CROWDSTRIKE_CLIENT_SECRET")
			apiURL := os.Getenv("CLOUDFLARE_CROWDSTRIKE_API_URL")
			customerID := os.Getenv("CLOUDFLARE_CROWDSTRIKE_CUSTOMER_ID")
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

			rnd1 := utils.GenerateRandomResourceName()
			rnd2 := utils.GenerateRandomResourceName()
			rnd3 := utils.GenerateRandomResourceName()
			name1 := fmt.Sprintf("tf-migrate-comprehensive-1-%s", rnd1)
			name2 := fmt.Sprintf("tf-migrate-comprehensive-2-%s", rnd2)
			name3 := fmt.Sprintf("tf-migrate-comprehensive-3-%s", rnd3)
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd1, rnd2, rnd3, accountID, name1, name2, name3, apiURL, clientID, clientSecret, customerID)
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
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
					acctest.TestAccPreCheck_CrowdStrike(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Resource 1: Verify deprecated name was renamed and identifier removed
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd1,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd1,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name1),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd1,
								tfjsonpath.New("type"),
								knownvalue.StringExact("crowdstrike_s2s"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd1,
								tfjsonpath.New("interval"),
								knownvalue.StringExact("1h"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd1,
								tfjsonpath.New("config").AtMapKey("api_url"),
								knownvalue.StringExact(apiURL),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd1,
								tfjsonpath.New("config").AtMapKey("client_id"),
								knownvalue.StringExact(clientID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd1,
								tfjsonpath.New("config").AtMapKey("customer_id"),
								knownvalue.StringExact(customerID),
							),

							// Resource 2: Verify current name preserved and unusual interval maintained
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd2,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd2,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name2),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd2,
								tfjsonpath.New("type"),
								knownvalue.StringExact("crowdstrike_s2s"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd2,
								tfjsonpath.New("interval"),
								knownvalue.StringExact("6h"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd2,
								tfjsonpath.New("config").AtMapKey("api_url"),
								knownvalue.StringExact(apiURL),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd2,
								tfjsonpath.New("config").AtMapKey("client_id"),
								knownvalue.StringExact(clientID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd2,
								tfjsonpath.New("config").AtMapKey("customer_id"),
								knownvalue.StringExact(customerID),
							),

							// Resource 3: Verify deprecated name renamed and standard interval preserved
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd3,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd3,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name3),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd3,
								tfjsonpath.New("type"),
								knownvalue.StringExact("crowdstrike_s2s"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd3,
								tfjsonpath.New("interval"),
								knownvalue.StringExact("24h"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd3,
								tfjsonpath.New("config").AtMapKey("api_url"),
								knownvalue.StringExact(apiURL),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd3,
								tfjsonpath.New("config").AtMapKey("client_id"),
								knownvalue.StringExact(clientID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_device_posture_integration."+rnd3,
								tfjsonpath.New("config").AtMapKey("customer_id"),
								knownvalue.StringExact(customerID),
							),
						},
					),
				},
			})
		})
	}
}
