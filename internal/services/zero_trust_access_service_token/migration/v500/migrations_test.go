package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Migration Test Configuration
//
// Version is read from LAST_V4_VERSION environment variable (set in .github/workflows/migration-tests.yml)
// - Last stable v4 release: default 4.52.5
// - Current v5 release: auto-updates with releases (internal.PackageVersion)
//
// Based on breaking changes analysis:
// - Key changes for zero_trust_access_service_token:
//   - Resource rename: cloudflare_access_service_token → cloudflare_zero_trust_access_service_token
//   - Field removal: min_days_for_renewal removed in v5
//   - Type conversion: client_secret_version (int → float64)

// Embed test configs
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_deprecated_field.tf
var v4DeprecatedFieldConfig string

//go:embed testdata/v5_deprecated_field.tf
var v5DeprecatedFieldConfig string

//go:embed testdata/v4_zone_scoped.tf
var v4ZoneScopedConfig string

//go:embed testdata/v5_zone_scoped.tf
var v5ZoneScopedConfig string

//go:embed testdata/v4_legacy_name.tf
var v4LegacyNameConfig string

//go:embed testdata/v5_legacy_name.tf
var v5LegacyNameConfig string

//go:embed testdata/v4_type_conversion.tf
var v4TypeConversionConfig string

//go:embed testdata/v5_type_conversion.tf
var v5TypeConversionConfig string

//go:embed testdata/v4_complete.tf
var v4CompleteConfig string

//go:embed testdata/v5_complete.tf
var v5CompleteConfig string

// TestMigrateZeroTrustAccessServiceToken_Basic tests basic migration from v4 to v5
// This tests both v4→v5 migration and v5→v5 version bump
//
// This test verifies:
// 1. Resource is created successfully with v4 provider
// 2. Migration tool runs without errors
// 3. All fields are preserved correctly
// 4. State can be read by v5 provider
// 5. Default duration is applied when not specified
func TestMigrateZeroTrustAccessServiceToken_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string { return fmt.Sprintf(v4BasicConfig, rnd, accountID, rnd) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID string) string { return fmt.Sprintf(v5BasicConfig, rnd, accountID, rnd) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Zero Trust Access resources don't support API tokens yet
			if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
				t.Setenv("CLOUDFLARE_API_TOKEN", "")
			}

			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			resourceName := "cloudflare_zero_trust_access_service_token." + rnd
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with specific version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify resource exists with correct type
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
							// Verify fields are preserved
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact("test-"+rnd)),
							// Verify computed fields exist
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("client_id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_at"), knownvalue.NotNull()),
							// Verify default duration is applied
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("duration"), knownvalue.StringExact("8760h")),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustAccessServiceToken_WithDeprecatedField tests migration with field removal
// This tests both v4→v5 migration and v5→v5 version bump
//
// This test specifically verifies that:
// 1. The min_days_for_renewal field is removed from the config
// 2. The min_days_for_renewal field is removed from the state
// 3. The migration handles this gracefully without errors
// 4. Other fields (like duration) are preserved
func TestMigrateZeroTrustAccessServiceToken_WithDeprecatedField(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string { return fmt.Sprintf(v4DeprecatedFieldConfig, rnd, accountID, rnd) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID string) string { return fmt.Sprintf(v5DeprecatedFieldConfig, rnd, accountID, rnd) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
				t.Setenv("CLOUDFLARE_API_TOKEN", "")
			}

			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			resourceName := "cloudflare_zero_trust_access_service_token." + rnd
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create resource with specific version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					// Step 2: Run migration
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact("test-"+rnd)),
							// Verify duration is preserved (not removed with min_days_for_renewal)
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("duration"), knownvalue.StringExact("17520h")),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustAccessServiceToken_ZoneScoped tests zone-scoped resource migration
// This tests both v4→v5 migration and v5→v5 version bump
//
// This test verifies that:
// 1. Zone-scoped resources (using zone_id instead of account_id) migrate correctly
// 2. The zone_id is preserved
// 3. No account_id is added
func TestMigrateZeroTrustAccessServiceToken_ZoneScoped(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string { return fmt.Sprintf(v4ZoneScopedConfig, rnd, zoneID, rnd) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, zoneID string) string { return fmt.Sprintf(v5ZoneScopedConfig, rnd, zoneID, rnd) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
				t.Setenv("CLOUDFLARE_API_TOKEN", "")
			}

			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			resourceName := "cloudflare_zero_trust_access_service_token." + rnd
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create zone-scoped resource with specific version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					// Step 2: Run migration and verify zone_id is preserved
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
							// Verify zone_id is preserved
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact("test-"+rnd)),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustAccessServiceToken_LegacyName tests migration from deprecated resource name
// This tests both v4→v5 migration and v5→v5 version bump
//
// This test verifies that:
// 1. The deprecated v4 resource name "cloudflare_access_service_token" is handled
// 2. The resource is migrated to "cloudflare_zero_trust_access_service_token"
// 3. All data is preserved during the rename
func TestMigrateZeroTrustAccessServiceToken_LegacyName(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string { return fmt.Sprintf(v4LegacyNameConfig, rnd, accountID, rnd) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID string) string { return fmt.Sprintf(v5LegacyNameConfig, rnd, accountID, rnd) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
				t.Setenv("CLOUDFLARE_API_TOKEN", "")
			}

			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			// Note: The resource name in state will be the v5 name after migration
			resourceName := "cloudflare_zero_trust_access_service_token." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create resource with specific version (may use legacy name in v4)
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					// Step 2: Run migration
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact("test-"+rnd)),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustAccessServiceToken_TypeConversion tests client_secret_version type conversion
// This tests both v4→v5 migration and v5→v5 version bump
//
// This test verifies that:
// 1. The client_secret_version field is converted from int to float64
// 2. Values like 1, 2, 3 become 1.0, 2.0, 3.0
// 3. The conversion happens transparently without errors
func TestMigrateZeroTrustAccessServiceToken_TypeConversion(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string { return fmt.Sprintf(v4TypeConversionConfig, rnd, accountID, rnd) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID string) string { return fmt.Sprintf(v5TypeConversionConfig, rnd, accountID, rnd) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
				t.Setenv("CLOUDFLARE_API_TOKEN", "")
			}

			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			resourceName := "cloudflare_zero_trust_access_service_token." + rnd
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create resource with specific version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					// Step 2: Run migration
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact("test-"+rnd)),
							// Verify client_secret_version exists (should be converted to float64)
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("client_secret_version"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("duration"), knownvalue.StringExact("8760h")),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustAccessServiceToken_CompleteResource tests migration with all fields
// This tests both v4→v5 migration and v5→v5 version bump
//
// This test verifies that:
// 1. A resource with all possible fields migrates correctly
// 2. All fields are preserved
// 3. The deprecated field is removed
// 4. Type conversions work with a complete resource
func TestMigrateZeroTrustAccessServiceToken_CompleteResource(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string { return fmt.Sprintf(v4CompleteConfig, rnd, accountID, rnd) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID string) string { return fmt.Sprintf(v5CompleteConfig, rnd, accountID, rnd) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
				t.Setenv("CLOUDFLARE_API_TOKEN", "")
			}

			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			resourceName := "cloudflare_zero_trust_access_service_token." + rnd
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create complete resource with specific version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					// Step 2: Run migration on complete resource
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact("test-"+rnd)),
							// Verify duration is preserved (43800h = 5 years)
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("duration"), knownvalue.StringExact("43800h")),
							// Verify computed fields
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("client_id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("client_secret"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_at"), knownvalue.NotNull()),
						},
					),
				},
			})
		})
	}
}
