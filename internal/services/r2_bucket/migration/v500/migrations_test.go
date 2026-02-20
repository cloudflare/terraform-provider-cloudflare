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
// For r2_bucket:
// - Resource name unchanged (cloudflare_r2_bucket → cloudflare_r2_bucket)
// - New fields added: jurisdiction (default: "default"), storage_class (default: "Standard"), creation_date (computed)
// - All existing fields preserved: id, name, account_id, location

// Embed migration test configuration files
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_with_location.tf
var v4WithLocationConfig string

//go:embed testdata/v5_with_location.tf
var v5WithLocationConfig string

//go:embed testdata/v4_multiple.tf
var v4MultipleConfig string

// TestMigrateR2BucketBasic tests migration of a basic R2 bucket with minimal config
func TestMigrateR2BucketBasic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, bucketName string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, bucketName string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, accountID, bucketName)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, bucketName string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, accountID, bucketName)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			bucketName := fmt.Sprintf("tf-test-r2-bucket-%s", rnd)
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, bucketName)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			testSteps := []resource.TestStep{
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
			}

			// Step 2-4: Run migration with state normalization for v4→v5, or simple migration for v5→v5
			migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
				// Verify existing fields are preserved
				statecheck.ExpectKnownValue("cloudflare_r2_bucket."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_r2_bucket."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(bucketName)),
				// Note: jurisdiction, storage_class, and creation_date are normalized after migration
			})

			testSteps = append(testSteps, migrationSteps...)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps:      testSteps,
			})
		})
	}
}

// TestMigrateR2BucketWithLocation tests migration of R2 bucket with location specified
func TestMigrateR2BucketWithLocation(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, bucketName string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, bucketName string) string {
				return fmt.Sprintf(v4WithLocationConfig, rnd, accountID, bucketName)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, bucketName string) string {
				return fmt.Sprintf(v5WithLocationConfig, rnd, accountID, bucketName)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			bucketName := fmt.Sprintf("tf-test-r2-bucket-%s", rnd)
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, bucketName)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			testSteps := []resource.TestStep{
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
			}

			// Step 2-4: Run migration with state normalization
			migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_r2_bucket."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_r2_bucket."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(bucketName)),
				statecheck.ExpectKnownValue("cloudflare_r2_bucket."+rnd, tfjsonpath.New("location"), knownvalue.StringExact("WNAM")),
			})

			testSteps = append(testSteps, migrationSteps...)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps:      testSteps,
			})
		})
	}
}

// TestMigrateR2BucketMultiple tests migration of multiple R2 buckets in one config
func TestMigrateR2BucketMultiple(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	rnd1 := rnd + "1"
	rnd2 := rnd + "2"
	bucketName1 := fmt.Sprintf("tf-test-bucket-1-%s", rnd)
	bucketName2 := fmt.Sprintf("tf-test-bucket-2-%s", rnd)
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

	v4Config := fmt.Sprintf(v4MultipleConfig, rnd1, accountID, bucketName1, rnd2, accountID, bucketName2)

	testSteps := []resource.TestStep{
		{
			// Step 1: Create with v4 provider
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: version,
				},
			},
			Config: v4Config,
		},
	}

	// Step 2-4: Run migration with state normalization
	migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
		// Verify first bucket
		statecheck.ExpectKnownValue("cloudflare_r2_bucket."+rnd1, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
		statecheck.ExpectKnownValue("cloudflare_r2_bucket."+rnd1, tfjsonpath.New("name"), knownvalue.StringExact(bucketName1)),

		// Verify second bucket
		statecheck.ExpectKnownValue("cloudflare_r2_bucket."+rnd2, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
		statecheck.ExpectKnownValue("cloudflare_r2_bucket."+rnd2, tfjsonpath.New("name"), knownvalue.StringExact(bucketName2)),
		statecheck.ExpectKnownValue("cloudflare_r2_bucket."+rnd2, tfjsonpath.New("location"), knownvalue.StringExact("EEUR")),
	})

	testSteps = append(testSteps, migrationSteps...)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps:      testSteps,
	})
}

// TestMigrateR2BucketVariousLocations tests migration with different location values
func TestMigrateR2BucketVariousLocations(t *testing.T) {
	// Test multiple location variants to ensure location handling works correctly
	testCases := []struct {
		name     string
		location string
	}{
		{name: "WNAM", location: "WNAM"},
		{name: "ENAM", location: "ENAM"},
		{name: "WEUR", location: "WEUR"},
		{name: "EEUR", location: "EEUR"},
		{name: "APAC", location: "APAC"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			bucketName := fmt.Sprintf("tf-test-bucket-%s", rnd)
			tmpDir := t.TempDir()
			version := acctest.GetLastV4Version()
			sourceVer, targetVer := acctest.InferMigrationVersions(version)

			v4Config := fmt.Sprintf(`
resource "cloudflare_r2_bucket" "%s" {
  account_id = "%s"
  name       = "%s"
  location   = "%s"
}`, rnd, accountID, bucketName, tc.location)

			testSteps := []resource.TestStep{
				{
					// Step 1: Create with v4 provider
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: version,
						},
					},
					Config: v4Config,
				},
			}

			// Step 2-4: Run migration with state normalization
			migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_r2_bucket."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_r2_bucket."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(bucketName)),
				statecheck.ExpectKnownValue("cloudflare_r2_bucket."+rnd, tfjsonpath.New("location"), knownvalue.StringExact(tc.location)),
			})

			testSteps = append(testSteps, migrationSteps...)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps:      testSteps,
			})
		})
	}
}
