package r2_bucket_test

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

// TestMigrateR2BucketBasic tests basic migration from v4 to v5 with minimal config
func TestMigrateR2BucketBasic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_r2_bucket." + rnd
	tmpDir := t.TempDir()
	bucketName := fmt.Sprintf("tf-test-bucket-%s", rnd)

	// V4 config - simple pass-through migration (no transformations needed)
	v4Config := fmt.Sprintf(`
resource "cloudflare_r2_bucket" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[3]s"
}`, rnd, accountID, bucketName)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
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
				// Verify resource exists with same type (no rename)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(bucketName)),
				// Verify new v5 computed fields are present
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("storage_class"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("creation_date"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateR2BucketWithLocation tests migration with location specified
func TestMigrateR2BucketWithLocation(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_r2_bucket." + rnd
	tmpDir := t.TempDir()
	bucketName := fmt.Sprintf("tf-test-bucket-%s", rnd)

	// V4 config with uppercase location (v4 style)
	v4Config := fmt.Sprintf(`
resource "cloudflare_r2_bucket" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[3]s"
  location   = "WNAM"
}`, rnd, accountID, bucketName)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(bucketName)),
				// Location should be preserved (v5 handles case normalization)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("location"), knownvalue.NotNull()),
				// Verify new v5 fields
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("storage_class"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("creation_date"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateR2BucketMultiple tests migration of multiple R2 buckets in one config
func TestMigrateR2BucketMultiple(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	rnd1 := rnd + "1"
	rnd2 := rnd + "2"
	resourceName1 := "cloudflare_r2_bucket." + rnd1
	resourceName2 := "cloudflare_r2_bucket." + rnd2
	tmpDir := t.TempDir()
	bucketName1 := fmt.Sprintf("tf-test-bucket-1-%s", rnd)
	bucketName2 := fmt.Sprintf("tf-test-bucket-2-%s", rnd)

	v4Config := fmt.Sprintf(`
resource "cloudflare_r2_bucket" "%[1]s" {
  account_id = "%[3]s"
  name       = "%[4]s"
}

resource "cloudflare_r2_bucket" "%[2]s" {
  account_id = "%[3]s"
  name       = "%[5]s"
  location   = "EEUR"
}`, rnd1, rnd2, accountID, bucketName1, bucketName2)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
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
				// Verify first bucket
				statecheck.ExpectKnownValue(resourceName1, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName1, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName1, tfjsonpath.New("name"), knownvalue.StringExact(bucketName1)),
				statecheck.ExpectKnownValue(resourceName1, tfjsonpath.New("jurisdiction"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName1, tfjsonpath.New("storage_class"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName1, tfjsonpath.New("creation_date"), knownvalue.NotNull()),
				// Verify second bucket
				statecheck.ExpectKnownValue(resourceName2, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName2, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName2, tfjsonpath.New("name"), knownvalue.StringExact(bucketName2)),
				statecheck.ExpectKnownValue(resourceName2, tfjsonpath.New("location"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName2, tfjsonpath.New("jurisdiction"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName2, tfjsonpath.New("storage_class"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName2, tfjsonpath.New("creation_date"), knownvalue.NotNull()),
			}),
		},
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
		{name: "OC", location: "OC"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_r2_bucket." + rnd
			tmpDir := t.TempDir()
			bucketName := fmt.Sprintf("tf-test-bucket-%s", rnd)

			v4Config := fmt.Sprintf(`
resource "cloudflare_r2_bucket" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[3]s"
  location   = "%[4]s"
}`, rnd, accountID, bucketName, tc.location)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
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
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(bucketName)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("location"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("storage_class"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("creation_date"), knownvalue.NotNull()),
					}),
				},
			})
		})
	}
}
