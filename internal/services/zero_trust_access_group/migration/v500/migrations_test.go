package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/compare"
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

//go:embed testdata/v4_list_explosion.tf
var v4ListExplosionConfig string

//go:embed testdata/v5_list_explosion.tf
var v5ListExplosionConfig string

//go:embed testdata/v4_github.tf
var v4GitHubConfig string

//go:embed testdata/v4_boolean.tf
var v4BooleanConfig string

//go:embed testdata/v5_boolean.tf
var v5BooleanConfig string

//go:embed testdata/v4_service_token_boolean.tf
var v4ServiceTokenBooleanConfig string

//go:embed testdata/v5_service_token_boolean.tf
var v5ServiceTokenBooleanConfig string

//go:embed testdata/v4_multi_list.tf
var v4MultiListConfig string

//go:embed testdata/v5_multi_list.tf
var v5MultiListConfig string

// TestMigrateZeroTrustAccessGroup_V4ToV5_Basic tests basic field migrations with DUAL test cases.
// This test validates:
// - Resource rename: cloudflare_access_group → cloudflare_zero_trust_access_group (v4 test only)
// - Basic field preservation (account_id, name)
// - Simple include block transformation
func TestMigrateZeroTrustAccessGroup_V4ToV5_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test setup
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := "tf-test-" + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with specific provider version
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
							// Verify account_id preserved
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_access_group."+rnd,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							// Verify name preserved
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_access_group."+rnd,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name),
							),
							// Verify include block exists
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_access_group."+rnd,
								tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("email").AtMapKey("email"),
								knownvalue.StringExact("user@example.com"),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustAccessGroup_V4ToV5_ListExplosion tests the critical "list explosion" transformation with DUAL test cases.
// This test validates:
// - v4: email = ["user1@example.com", "user2@example.com", "user3@example.com"]
// - v5: Three separate include blocks, each with one email object
func TestMigrateZeroTrustAccessGroup_V4ToV5_ListExplosion(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v4ListExplosionConfig, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v5ListExplosionConfig, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test setup
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := "tf-test-" + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with specific provider version
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
							// CRITICAL: Verify list explosion - 3 emails become 3 separate include blocks
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_access_group."+rnd,
								tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("email").AtMapKey("email"),
								knownvalue.StringExact("user1@example.com"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_access_group."+rnd,
								tfjsonpath.New("include").AtSliceIndex(1).AtMapKey("email").AtMapKey("email"),
								knownvalue.StringExact("user2@example.com"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_zero_trust_access_group."+rnd,
								tfjsonpath.New("include").AtSliceIndex(2).AtMapKey("email").AtMapKey("email"),
								knownvalue.StringExact("user3@example.com"),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustAccessGroup_V4ToV5_GitHub tests field rename and nested list explosion.
// This test validates:
// - Field rename: github → github_organization
// - Nested list explosion: teams = ["team1", "team2"] → two separate github_organization blocks
// - Field rename within nested object: teams → team (singular)
//
// TODO: This test is currently skipped because it requires cloudflare_access_identity_provider
// migration to be implemented first. Once that migration is complete, remove the Skip and enable this test.
func TestMigrateZeroTrustAccessGroup_V4ToV5_GitHub(t *testing.T) {
	t.Run("from_v4_latest", func(t *testing.T) {
		// Test setup
		accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
		rnd := utils.GenerateRandomResourceName()
		name := "tf-test-" + rnd
		tmpDir := t.TempDir()
		testConfig := fmt.Sprintf(v4GitHubConfig, rnd, accountID, name)
		version := acctest.GetLastV4Version()
		sourceVer, targetVer := acctest.InferMigrationVersions(version)

		resource.Test(t, resource.TestCase{
			WorkingDir: tmpDir,
			Steps: []resource.TestStep{
				{
					// Step 1: Create with v4 external provider
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: version,
						},
					},
					Config: testConfig,
				},
				// Step 2: Run migration and verify state
				acctest.MigrationV2TestStep(t, testConfig, tmpDir, version, sourceVer, targetVer,
					[]statecheck.StateCheck{
						// CRITICAL: Verify field rename github → github_organization
						statecheck.ExpectKnownValue(
							"cloudflare_zero_trust_access_group."+rnd,
							tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("github_organization").AtMapKey("name"),
							knownvalue.StringExact("my-org"),
						),
						// CRITICAL: Verify teams list explosion - 2 teams → 2 github_organization blocks
						statecheck.ExpectKnownValue(
							"cloudflare_zero_trust_access_group."+rnd,
							tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("github_organization").AtMapKey("team"),
							knownvalue.StringExact("team1"),
						),
						statecheck.ExpectKnownValue(
							"cloudflare_zero_trust_access_group."+rnd,
							tfjsonpath.New("include").AtSliceIndex(1).AtMapKey("github_organization").AtMapKey("team"),
							knownvalue.StringExact("team2"),
						),
						// Verify identity_provider_id preserved in both blocks (must match actual IDP resource)
						statecheck.CompareValuePairs(
							"cloudflare_zero_trust_access_group."+rnd,
							tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("github_organization").AtMapKey("identity_provider_id"),
							"cloudflare_zero_trust_access_identity_provider."+rnd+"_idp",
							tfjsonpath.New("id"),
							compare.ValuesSame(),
						),
					},
				),
			},
		})
	})
}

// TestMigrateZeroTrustAccessGroup_V4ToV5_Boolean tests boolean to empty object transformation.
// This test validates:
// - Boolean true → empty object {}
// - v4: everyone = true
// - v5: everyone {}
func TestMigrateZeroTrustAccessGroup_V4ToV5_Boolean(t *testing.T) {
	t.Run("from_v4_latest", func(t *testing.T) {
		// Test setup
		accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
		rnd := utils.GenerateRandomResourceName()
		name := "tf-test-" + rnd
		tmpDir := t.TempDir()
		testConfig := fmt.Sprintf(v4BooleanConfig, rnd, accountID, name)
		version := acctest.GetLastV4Version()
		sourceVer, targetVer := acctest.InferMigrationVersions(version)

		stateChecks := []statecheck.StateCheck{
			// CRITICAL: Verify boolean → empty object transformation
			// In v5, everyone is an empty object (no fields inside)
			// We verify the everyone key exists in the include block
			statecheck.ExpectKnownValue(
				"cloudflare_zero_trust_access_group."+rnd,
				tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("everyone"),
				knownvalue.MapExact(map[string]knownvalue.Check{}),
			),
		}

		// Use MigrationV2TestStepWithStateNormalization because the v5 provider's schema
		// returns nil selector fields (like common_name) that need to be normalized away
		migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, testConfig, tmpDir, version, sourceVer, targetVer, stateChecks)

		resource.Test(t, resource.TestCase{
			WorkingDir: tmpDir,
			Steps: append([]resource.TestStep{
				{
					// Step 1: Create with v4 external provider
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: version,
						},
					},
					Config: testConfig,
				},
			}, migrationSteps...),
		})
	})
}

// TestMigrateZeroTrustAccessGroup_V4ToV5_ServiceTokenBoolean tests the any_valid_service_token
// boolean-to-object transformation that was reported as broken in APIX-741.
//
// This is the exact scenario that fails without the UpgradeState fix:
//   - v4 state has any_valid_service_token = true (boolean) in include block
//   - v5 expects any_valid_service_token = {} (SingleNestedAttribute, empty object)
//   - Without the fix, UpgradeState fails with:
//     "invalid JSON, expected "{", got false"
func TestMigrateZeroTrustAccessGroup_V4ToV5_ServiceTokenBoolean(t *testing.T) {
	t.Run("from_v4_latest", func(t *testing.T) {
		// Test setup
		accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
		rnd := utils.GenerateRandomResourceName()
		name := "tf-test-" + rnd
		tmpDir := t.TempDir()
		testConfig := fmt.Sprintf(v4ServiceTokenBooleanConfig, rnd, accountID, name)
		version := acctest.GetLastV4Version()
		sourceVer, targetVer := acctest.InferMigrationVersions(version)

		stateChecks := []statecheck.StateCheck{
			// CRITICAL: Verify any_valid_service_token boolean → empty object transformation
			// This is the exact field that causes the APIX-741 error
			statecheck.ExpectKnownValue(
				"cloudflare_zero_trust_access_group."+rnd,
				tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("any_valid_service_token"),
				knownvalue.MapExact(map[string]knownvalue.Check{}),
			),
			// Verify exclude block email is preserved
			statecheck.ExpectKnownValue(
				"cloudflare_zero_trust_access_group."+rnd,
				tfjsonpath.New("exclude").AtSliceIndex(0).AtMapKey("email").AtMapKey("email"),
				knownvalue.StringExact("blocked@example.com"),
			),
		}

		// Use MigrationV2TestStepWithStateNormalization because the v5 provider's schema
		// returns nil selector fields that need to be normalized away
		migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, testConfig, tmpDir, version, sourceVer, targetVer, stateChecks)

		resource.Test(t, resource.TestCase{
			WorkingDir: tmpDir,
			Steps: append([]resource.TestStep{
				{
					// Step 1: Create with v4 external provider
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: version,
						},
					},
					Config: testConfig,
				},
			}, migrationSteps...),
		})
	})
}

// TestMigrateZeroTrustAccessGroup_V4ToV5_MultiList tests multiple list field explosions with DUAL test cases.
// This test validates:
// - Multiple list fields in include block (email, ip)
// - List fields in exclude block (geo)
// - Field rename: geo value → country_code
// - All lists explode correctly: v4 arrays → multiple v5 objects
func TestMigrateZeroTrustAccessGroup_V4ToV5_MultiList(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v4MultiListConfig, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v5MultiListConfig, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test setup
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := "tf-test-" + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			stateChecks := []statecheck.StateCheck{
				// Verify include block explosions - 4 total
				// Order matches transform.go processing order:
				// email (2) → ip (2)

				// Email 1
				statecheck.ExpectKnownValue(
					"cloudflare_zero_trust_access_group."+rnd,
					tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("email").AtMapKey("email"),
					knownvalue.StringExact("user1@example.com"),
				),
				// Email 2
				statecheck.ExpectKnownValue(
					"cloudflare_zero_trust_access_group."+rnd,
					tfjsonpath.New("include").AtSliceIndex(1).AtMapKey("email").AtMapKey("email"),
					knownvalue.StringExact("user2@example.com"),
				),
				// IP 1
				statecheck.ExpectKnownValue(
					"cloudflare_zero_trust_access_group."+rnd,
					tfjsonpath.New("include").AtSliceIndex(2).AtMapKey("ip").AtMapKey("ip"),
					knownvalue.StringExact("192.168.1.0/24"),
				),
				// IP 2
				statecheck.ExpectKnownValue(
					"cloudflare_zero_trust_access_group."+rnd,
					tfjsonpath.New("include").AtSliceIndex(3).AtMapKey("ip").AtMapKey("ip"),
					knownvalue.StringExact("10.0.0.0/8"),
				),
				// Verify exclude block explosions - 2 total (2 geos)
				// Geo 1 (verify field rename: geo → country_code)
				statecheck.ExpectKnownValue(
					"cloudflare_zero_trust_access_group."+rnd,
					tfjsonpath.New("exclude").AtSliceIndex(0).AtMapKey("geo").AtMapKey("country_code"),
					knownvalue.StringExact("US"),
				),
				// Geo 2
				statecheck.ExpectKnownValue(
					"cloudflare_zero_trust_access_group."+rnd,
					tfjsonpath.New("exclude").AtSliceIndex(1).AtMapKey("geo").AtMapKey("country_code"),
					knownvalue.StringExact("CA"),
				),
			}

			// Use MigrationV2TestStepWithStateNormalization because the v5 provider's schema
			// returns nil selector fields that need to be normalized away
			migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, stateChecks)

			resource.Test(t, resource.TestCase{
				WorkingDir: tmpDir,
				Steps: append([]resource.TestStep{
					{
						// Step 1: Create with specific provider version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
				}, migrationSteps...),
			})
		})
	}
}
