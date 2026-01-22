package api_token_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateAPITokenFromV5MapToJSON tests migration from v5 api_token with resources as map to latest with resources as JSON string
func TestMigrateAPITokenFromV5MapToJSON(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V5 config using api_token with resources as a map (old format)
	// This represents the configuration that users would have with v5.x.x
	v5Config := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b" # DNS Read
    }]
    resources = {
      "com.cloudflare.api.account.zone.*" = "*"
    }
  }]

  condition = {
    request_ip = {
      in = ["192.0.2.1/32"]
    }
  }
}`, rnd)

	// Latest config using api_token with resources as JSON string (new format)
	// This is what the configuration should look like after migration
	latestConfig := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b" # DNS Read
    }]
    resources = jsonencode({
      "com.cloudflare.api.account.zone.*" = "*"
    })
  }]

  condition = {
    request_ip = {
      in = ["192.0.2.1/32"]
    }
  }
}`, rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5 provider (before resources field change)
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.0.0", // Early v5 version before the breaking change
					},
				},
				Config: v5Config,
			},
			{
				// Step 2: Upgrade to latest provider with state upgrader
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify the token name is preserved
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_api_token.%s", rnd),
						tfjsonpath.New("name"),
						knownvalue.StringExact(rnd),
					),
					// Verify the effect is preserved
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_api_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("effect"),
						knownvalue.StringExact("allow"),
					),
					// Verify permission group ID is preserved
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_api_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id"),
						knownvalue.StringExact("82e64a83756745bbbb1c9c2701bf816b"),
					),
					// Verify resources is now a JSON string
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_api_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(`{"com.cloudflare.api.account.zone.*":"*"}`),
					),
					// Verify condition is preserved
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_api_token.%s", rnd),
						tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in").AtSliceIndex(0),
						knownvalue.StringExact("192.0.2.1/32"),
					),
				},
			},
		},
	})
}

// TestMigrateAPITokenFromV5ComplexResources tests migration with complex nested resources
func TestMigrateAPITokenFromV5ComplexResources(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	// V5 config with simple map resources (old format limitation)
	v5Config := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policies = [
    {
      effect = "allow"
      permission_groups = [{
        id = "82e64a83756745bbbb1c9c2701bf816b" # DNS Read
      }]
      resources = {
        "com.cloudflare.api.account.%[2]s" = "*"
      }
    },
    {
      effect = "allow"
      permission_groups = [{
        id = "e199d584e69344eba202452019deafe3" # Another permission
      }]
      resources = {
        "com.cloudflare.api.account.zone.*" = "*"
      }
    }
  ]
}`, rnd, accountID)

	// Latest config with complex nested resources (new capability)
	latestConfig := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policies = [
    {
      effect = "allow"
      permission_groups = [{
        id = "82e64a83756745bbbb1c9c2701bf816b" # DNS Read
      }]
      resources = jsonencode({
        "com.cloudflare.api.account.%[2]s" = "*"
      })
    },
    {
      effect = "allow"
      permission_groups = [{
        id = "e199d584e69344eba202452019deafe3" # Another permission
      }]
      resources = jsonencode({
        "com.cloudflare.api.account.zone.*" = "*"
      })
    }
  ]
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.0.0",
					},
				},
				Config: v5Config,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify first policy resources
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_api_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID)),
					),
					// Verify second policy resources
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_api_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(1).AtMapKey("resources"),
						knownvalue.StringExact(`{"com.cloudflare.api.account.zone.*":"*"}`),
					),
				},
			},
		},
	})
}

// TestMigrateAPITokenFromV5WithTTL tests migration with token TTL settings
func TestMigrateAPITokenFromV5WithTTL(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V5 config with TTL settings
	v5Config := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
    }]
    resources = {
      "com.cloudflare.api.account.zone.*" = "*"
    }
  }]

  not_before = "2027-01-01T00:00:00Z"
  expires_on = "2027-12-31T23:59:59Z"
}`, rnd)

	// Latest config with jsonencode
	latestConfig := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
    }]
    resources = jsonencode({
      "com.cloudflare.api.account.zone.*" = "*"
    })
  }]

  not_before = "2027-01-01T00:00:00Z"
  expires_on = "2027-12-31T23:59:59Z"
}`, rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.0.0",
					},
				},
				Config: v5Config,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify TTL settings are preserved
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_api_token.%s", rnd),
						tfjsonpath.New("not_before"),
						knownvalue.StringExact("2027-01-01T00:00:00Z"),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_api_token.%s", rnd),
						tfjsonpath.New("expires_on"),
						knownvalue.StringExact("2027-12-31T23:59:59Z"),
					),
					// Verify resources migration
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_api_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(`{"com.cloudflare.api.account.zone.*":"*"}`),
					),
				},
			},
		},
	})
}

// TestMigrateAPITokenFromV5RemovedFields tests that computed fields are properly removed during migration
func TestMigrateAPITokenFromV5RemovedFields(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V5 config - the old schema had computed fields like policy.id, permission_groups.meta, permission_groups.name
	v5Config := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
      # In v5, meta and name were computed fields that would be populated
    }]
    resources = {
      "com.cloudflare.api.account.zone.*" = "*"
    }
  }]
}`, rnd)

	// Latest config
	latestConfig := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
      # meta and name fields are removed in latest schema
    }]
    resources = jsonencode({
      "com.cloudflare.api.account.zone.*" = "*"
    })
  }]
}`, rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.0.0",
					},
				},
				Config: v5Config,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify permission group only has id field
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_api_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id"),
						knownvalue.StringExact("82e64a83756745bbbb1c9c2701bf816b"),
					),
					// Verify resources is migrated to JSON
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_api_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(`{"com.cloudflare.api.account.zone.*":"*"}`),
					),
				},
			},
		},
	})
}

// TestMigrateAPITokenFromV5_4 tests migration from v5.4.0
func TestMigrateAPITokenFromV5_4(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V5.4 config
	v5Config := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
    }]
    resources = {
      "com.cloudflare.api.account.zone.*" = "*"
    }
  }]
}`, rnd)

	// Latest config
	latestConfig := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
    }]
    resources = jsonencode({
      "com.cloudflare.api.account.zone.*" = "*"
    })
  }]
}`, rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5.4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.4.0",
					},
				},
				Config: v5Config,
				// v5.4.0 has known computed field differences
				ExpectNonEmptyPlan: true,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_api_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(`{"com.cloudflare.api.account.zone.*":"*"}`),
					),
				},
			},
		},
	})
}

// TestMigrateAPITokenFromV5_7 tests migration from v5.7.0
func TestMigrateAPITokenFromV5_7(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V5.7 config
	v5Config := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
    }]
    resources = {
      "com.cloudflare.api.account.zone.*" = "*"
    }
  }]

  expires_on = "2027-12-31T23:59:59Z"
}`, rnd)

	// Latest config
	latestConfig := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
    }]
    resources = jsonencode({
      "com.cloudflare.api.account.zone.*" = "*"
    })
  }]

  expires_on = "2027-12-31T23:59:59Z"
}`, rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5.7 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.7.0",
					},
				},
				Config: v5Config,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_api_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(`{"com.cloudflare.api.account.zone.*":"*"}`),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_api_token.%s", rnd),
						tfjsonpath.New("expires_on"),
						knownvalue.StringExact("2027-12-31T23:59:59Z"),
					),
				},
			},
		},
	})
}

// TestMigrateAPITokenFromV5_ComplexNestedResources tests migration with nested resources that v5 can't handle
func TestMigrateAPITokenFromV5_ComplexNestedResources(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	// V5 config - can only use flat resources, not nested
	v5Config := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
    }]
    resources = {
      "com.cloudflare.api.account.%[2]s" = "*"
    }
  }]
}`, rnd, accountID)

	// Latest config - now supports nested resources via jsonencode
	latestConfigNested := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
    }]
    resources = jsonencode({
      "com.cloudflare.api.account.%[2]s" = {
        "com.cloudflare.api.account.zone.*" = "*"
      }
    })
  }]
}`, rnd, accountID)

	// Config that tries to use nested structure directly in v5 (will fail)
	v5NestedConfig := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
    }]
    resources = {
      "com.cloudflare.api.account.%[2]s" = {
        "com.cloudflare.api.account.zone.*" = "*"
      }
    }
  }]
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Try nested structure with v5.7 - should fail
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.7.0",
					},
				},
				Config:      v5NestedConfig,
				ExpectError: regexp.MustCompile(`Incorrect attribute value type`),
			},
			{
				// Step 2: Create with v5.7 using flat structure
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.7.0",
					},
				},
				Config: v5Config,
			},
			{
				// Step 3: Upgrade to latest provider with nested structure
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfigNested,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_api_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":{"com.cloudflare.api.account.zone.*":"*"}}`, accountID)),
					),
				},
			},
		},
	})
}

// TestMigrateAPITokenFromV5_4_BothResourceFormats tests both flat and nested resource formats
func TestMigrateAPITokenFromV5_4_BothResourceFormats(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	// V5.4 config with multiple policies - only flat resources work
	v5Config := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policies = [
    {
      effect = "allow"
      permission_groups = [{
        id = "82e64a83756745bbbb1c9c2701bf816b"
      }]
      resources = {
        "com.cloudflare.api.account.%[2]s" = "*"
      }
    },
    {
      effect = "allow"
      permission_groups = [{
        id = "c8fed203ed3043cba015a93ad1616f1f"
      }]
      resources = {
        "com.cloudflare.api.account.zone.*" = "*"
      }
    }
  ]
}`, rnd, accountID)

	// Latest config - mix of flat and nested resources via jsonencode
	latestConfig := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policies = [
    {
      effect = "allow"
      permission_groups = [{
        id = "82e64a83756745bbbb1c9c2701bf816b"
      }]
      # Flat structure still works with jsonencode
      resources = jsonencode({
        "com.cloudflare.api.account.%[2]s" = "*"
      })
    },
    {
      effect = "allow"
      permission_groups = [{
        id = "c8fed203ed3043cba015a93ad1616f1f"
      }]
      # Nested structure now possible with jsonencode
      resources = jsonencode({
        "com.cloudflare.api.account.%[2]s" = {
          "com.cloudflare.api.account.zone.*" = "*"
        }
      })
    }
  ]
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5.4 provider - flat only
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.4.0",
					},
				},
				Config: v5Config,
				// v5.4.0 has known computed field differences
				ExpectNonEmptyPlan: true,
			},
			{
				// Step 2: Upgrade to latest with mixed flat and nested
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					// First policy - flat structure
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_api_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID)),
					),
					// Second policy - nested structure
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_api_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(1).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":{"com.cloudflare.api.account.zone.*":"*"}}`, accountID)),
					),
				},
			},
		},
	})
}

// ==================== V4 to V5 Migration Tests ====================

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
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("not_before"), knownvalue.StringExact("2027-01-01T00:00:00Z")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact("2027-12-31T23:59:59Z")),
				},
			},
			// Step 2: Migrate to v5 provider
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("not_before"), knownvalue.StringExact("2027-01-01T00:00:00Z")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact("2027-12-31T23:59:59Z")),
			}),
			{
				// Step 3: Apply migrated config with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("not_before"), knownvalue.StringExact("2027-01-01T00:00:00Z")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact("2027-12-31T23:59:59Z")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"), knownvalue.StringExact(`{"com.cloudflare.api.account.zone.*":"*"}`)),
				},
			},
		},
	})
}

// TestMigrateCloudflareAPIToken_Migration_DenyEffect tests migration with "deny" effect
// policy to ensure deny policies are properly migrated.
func TestMigrateCloudflareAPIToken_Migration_DenyEffect(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_token." + rnd
	v4Config := testAccCloudflareAPITokenMigrationConfigV4DenyEffect(rnd)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create api_token with v4 provider using deny effect
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("effect"), knownvalue.StringExact("deny")),
			}),
			{
				// Step 3: Apply the migrated configuration with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("effect"), knownvalue.StringExact("deny")),
				},
			},
		},
	})
}

// TestMigrateCloudflareAPIToken_Migration_ConditionNotIn tests migration with condition
// using only not_in (without in) to ensure partial condition blocks are handled correctly.
func TestMigrateCloudflareAPIToken_Migration_ConditionNotIn(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_token." + rnd
	v4Config := testAccCloudflareAPITokenMigrationConfigV4ConditionNotIn(rnd)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create api_token with v4 provider using only not_in
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("not_in"), knownvalue.ListSizeExact(1)),
			}),
			{
				// Step 3: Apply the migrated configuration with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("not_in"), knownvalue.ListSizeExact(1)),
				},
			},
		},
	})
}

// TestMigrateCloudflareAPIToken_Migration_MultipleResources tests migration
// with multiple resource keys to ensure proper JSON encoding with multiple entries.
func TestMigrateCloudflareAPIToken_Migration_MultipleResources(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for this acceptance test")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_token." + rnd
	v4Config := testAccCloudflareAPITokenMigrationConfigV4MultipleResources(rnd, accountID)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create api_token with v4 provider with multiple resources
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
				// Verify JSON encoding preserves multiple resources
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"), knownvalue.NotNull()),
			}),
			{
				// Step 3: Apply the migrated configuration with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
				},
			},
		},
	})
}

// TestMigrateCloudflareAPIToken_Migration_MixedEffects tests migration with multiple
// policies having different effects (allow and deny) to ensure proper handling.
func TestMigrateCloudflareAPIToken_Migration_MixedEffects(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for this acceptance test")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_token." + rnd
	v4Config := testAccCloudflareAPITokenMigrationConfigV4MixedEffects(rnd, accountID)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create api_token with v4 provider with mixed effects
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policy"), knownvalue.SetSizeExact(2)),
				},
			},
			// Step 2: Migrate to v5 provider
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(2)),
				// Verify first policy has allow effect
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("effect"), knownvalue.StringExact("allow")),
				// Verify second policy has deny effect
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(1).AtMapKey("effect"), knownvalue.StringExact("deny")),
			}),
			{
				// Step 3: Apply the migrated configuration with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("effect"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(1).AtMapKey("effect"), knownvalue.StringExact("deny")),
				},
			},
		},
	})
}

// ==================== V4 Configuration Helper Functions ====================

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

  not_before = "2027-01-01T00:00:00Z"
  expires_on = "2027-12-31T23:59:59Z"
}
`, rnd)
}

func testAccCloudflareAPITokenMigrationConfigV4DenyEffect(rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policy {
    effect = "deny"
    permission_groups = [
      "82e64a83756745bbbb1c9c2701bf816b"
    ]
    resources = {
      "com.cloudflare.api.account.zone.*" = "*"
    }
  }
}
`, rnd)
}

func testAccCloudflareAPITokenMigrationConfigV4ConditionNotIn(rnd string) string {
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

  condition {
    request_ip {
      not_in = ["192.0.2.0/24"]
    }
  }
}
`, rnd)
}

func testAccCloudflareAPITokenMigrationConfigV4MultipleResources(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policy {
    permission_groups = [
      "82e64a83756745bbbb1c9c2701bf816b"
    ]
    resources = {
      "com.cloudflare.api.account.%[2]s" = "*"
      "com.cloudflare.api.account.zone.*" = "*"
    }
  }
}
`, rnd, accountID)
}

func testAccCloudflareAPITokenMigrationConfigV4MixedEffects(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policy {
    effect = "allow"
    permission_groups = [
      "82e64a83756745bbbb1c9c2701bf816b"
    ]
    resources = {
      "com.cloudflare.api.account.%[2]s" = "*"
    }
  }

  policy {
    effect = "deny"
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
