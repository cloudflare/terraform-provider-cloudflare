package api_token_test

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
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateCloudflareAPIToken_Migration_Basic_MultiVersion tests the api_token
// migration from v4 to v5. This test ensures that:
// 1. policy block is converted to policies array attribute
// 2. permission_groups is converted from list of strings to list of objects with id field
// 3. resources is wrapped with jsonencode()
// 4. condition block is converted to condition object attribute
// 5. The migration tool successfully transforms both configuration and state files
// 6. Resources remain functional after migration without requiring manual intervention
func TestMigrateCloudflareAPIToken_Migration_Basic_MultiVersion(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd string) string
	}{
		{
			name:     "from_v4_52_1", // Last v4 release
			version:  "4.52.1",
			configFn: testAccCloudflareAPITokenMigrationConfigV4Basic,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_api_token." + rnd
			testConfig := tc.configFn(rnd)
			tmpDir := t.TempDir()

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create api_token with v4 provider
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								VersionConstraint: tc.version,
								Source:            "cloudflare/cloudflare",
							},
						},
						Config: testConfig,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policy"), knownvalue.NotNull()),
						},
					},
					// Step 2: Migrate to v5 provider
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, "v4", "v5", []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.NotNull()),
					}),
					{
						// Step 3: Apply the migrated configuration with v5 provider
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("effect"), knownvalue.StringExact("allow")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("82e64a83756745bbbb1c9c2701bf816b")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"), knownvalue.StringExact(`{"com.cloudflare.api.account.zone.*":"*"}`)),
						},
					},
				},
			})
		})
	}
}

// TestMigrateCloudflareAPIToken_Migration_WithCondition tests migration of api_token
// with the optional condition field to ensure all fields are properly migrated.
func TestMigrateCloudflareAPIToken_Migration_WithCondition(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_token." + rnd
	v4Config := testAccCloudflareAPITokenMigrationConfigV4WithCondition(rnd)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create api_token with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition"), knownvalue.NotNull()),
				},
			},
			// Step 2: Migrate to v5 provider
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition"), knownvalue.NotNull()),
			}),
			{
				// Step 3: Apply migrated config with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in").AtSliceIndex(0), knownvalue.StringExact("192.0.2.1/32")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"), knownvalue.StringExact(`{"com.cloudflare.api.account.zone.*":"*"}`)),
				},
			},
		},
	})
}

// TestMigrateCloudflareAPIToken_Migration_WithMultiplePolicies tests migration with multiple policies
func TestMigrateCloudflareAPIToken_Migration_WithMultiplePolicies(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_token." + rnd
	v4Config := testAccCloudflareAPITokenMigrationConfigV4MultiplePolicies(rnd, accountID)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create api_token with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policy"), knownvalue.NotNull()),
				},
			},
			// Step 2: Migrate to v5 provider
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.NotNull()),
			}),
			{
				// Step 3: Apply migrated config with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(2)),
					// First policy
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"), knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID))),
					// Second policy
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(1).AtMapKey("resources"), knownvalue.StringExact(`{"com.cloudflare.api.account.zone.*":"*"}`)),
				},
			},
		},
	})
}

// TestMigrateCloudflareAPIToken_Migration_WithTTL tests migration with TTL fields
func TestMigrateCloudflareAPIToken_Migration_WithTTL(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_token." + rnd
	v4Config := testAccCloudflareAPITokenMigrationConfigV4WithTTL(rnd)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create api_token with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("not_before"), knownvalue.StringExact("2025-01-01T00:00:00Z")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact("2025-12-31T23:59:59Z")),
				},
			},
			// Step 2: Migrate to v5 provider
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("not_before"), knownvalue.StringExact("2025-01-01T00:00:00Z")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact("2025-12-31T23:59:59Z")),
			}),
			{
				// Step 3: Apply migrated config with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("not_before"), knownvalue.StringExact("2025-01-01T00:00:00Z")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact("2025-12-31T23:59:59Z")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"), knownvalue.StringExact(`{"com.cloudflare.api.account.zone.*":"*"}`)),
				},
			},
		},
	})
}

// V4 Configuration Functions

func testAccCloudflareAPITokenMigrationConfigV4Basic(rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policy {
    permission_groups = [
      "82e64a83756745bbbb1c9c2701bf816b" # DNS Read
    ]
    resources = {
      "com.cloudflare.api.account.zone.*" = "*"
    }
  }
}
`, rnd)
}

func testAccCloudflareAPITokenMigrationConfigV4WithCondition(rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policy {
    permission_groups = [
      "82e64a83756745bbbb1c9c2701bf816b" # DNS Read
    ]
    resources = {
      "com.cloudflare.api.account.zone.*" = "*"
    }
  }

  condition {
    request_ip {
      in = ["192.0.2.1/32"]
    }
  }
}
`, rnd)
}

func testAccCloudflareAPITokenMigrationConfigV4MultiplePolicies(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policy {
    permission_groups = [
      "82e64a83756745bbbb1c9c2701bf816b" # DNS Read
    ]
    resources = {
      "com.cloudflare.api.account.%[2]s" = "*"
    }
  }

  policy {
    permission_groups = [
      "e199d584e69344eba202452019deafe3"
    ]
    resources = {
      "com.cloudflare.api.account.zone.*" = "*"
    }
  }
}
`, rnd, accountID)
}

func testAccCloudflareAPITokenMigrationConfigV4WithTTL(rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policy {
    permission_groups = [
      "82e64a83756745bbbb1c9c2701bf816b"
    ]
    resources = {
      "com.cloudflare.api.account.zone.*" = "*"
    }
  }

  not_before = "2025-01-01T00:00:00Z"
  expires_on = "2025-12-31T23:59:59Z"
}
`, rnd)
}
