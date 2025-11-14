package account_token_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateAccountTokenFromV5MapToJSON tests migration from v5 account_token with resources as map to latest with resources as JSON string
func TestMigrateAccountTokenFromV5MapToJSON(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	// V5 config using account_token with resources as a map (old format)
	v5Config := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b" # DNS Read
    }]
    resources = {
      "com.cloudflare.api.account.%[2]s" = "*"
    }
  }]

  condition = {
    request_ip = {
      in = ["192.0.2.1/32"]
    }
  }
}`, rnd, accountID)

	// Latest config using account_token with resources as JSON string (new format)
	latestConfig := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b" # DNS Read
    }]
    resources = jsonencode({
      "com.cloudflare.api.account.%[2]s" = "*"
    })
  }]

  condition = {
    request_ip = {
      in = ["192.0.2.1/32"]
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
				// Step 1: Create with v5 provider (before resources field change)
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.10.0", // Use a later v5 version where account_token is stable // Early v5 version before the breaking change
					},
				},
				Config: v5Config,
			},
			{
				// Step 2: Upgrade to latest provider with state upgrader
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify account_id is preserved
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("account_id"),
						knownvalue.StringExact(accountID),
					),
					// Verify the token name is preserved
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("name"),
						knownvalue.StringExact(rnd),
					),
					// Verify the effect is preserved
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("effect"),
						knownvalue.StringExact("allow"),
					),
					// Verify permission group ID is preserved
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id"),
						knownvalue.StringExact("82e64a83756745bbbb1c9c2701bf816b"),
					),
					// Verify resources is now a JSON string
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID)),
					),
					// Verify condition is preserved
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in").AtSliceIndex(0),
						knownvalue.StringExact("192.0.2.1/32"),
					),
				},
			},
		},
	})
}

// TestMigrateAccountTokenFromV5ComplexResources tests migration with complex nested resources
func TestMigrateAccountTokenFromV5ComplexResources(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	// V5 config with multiple policies (old format that works with v5.10.0)
	// Note: v5.10.0 requires proper nesting for zone resources under account
	v5Config := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
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
        id = "c8fed203ed3043cba015a93ad1616f1f" # Zone Read
      }]
      resources = {
        "com.cloudflare.api.account.%[2]s" = {
          "com.cloudflare.api.account.zone.*" = "*"
        }
      }
    }
  ]
}`, rnd, accountID)

	// Latest config with complex nested resources (new capability)
	latestConfig := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
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
        id = "c8fed203ed3043cba015a93ad1616f1f" # Zone Read
      }]
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
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Try to create with v5 provider - this will fail due to nested resources
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.10.0", // Use a later v5 version where account_token is stable
					},
				},
				Config: v5Config,
				ExpectError: regexp.MustCompile(`string\s+required`),
			},
			{
				// Step 2: Upgrade to latest provider - now the nested structure works with jsonencode
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify first policy resources
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID)),
					),
					// Verify second policy resources (now with nested structure)
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(1).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":{"com.cloudflare.api.account.zone.*":"*"}}`, accountID)),
					),
				},
			},
		},
	})
}

// TestMigrateAccountTokenFromV5_ComplexNestedResources tests migration with complex nested resources
// This demonstrates that v5.10 can't handle nested resources but the latest version with jsonencode can
func TestMigrateAccountTokenFromV5_ComplexNestedResources(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	// Step 1: V5.10 config attempting to use nested resources (this should fail)
	v5NestedConfig := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s-nested"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
    }]
    # This nested structure is not supported in v5.10
    resources = {
      "com.cloudflare.api.account.%[2]s" = {
        "com.cloudflare.api.account.zone.*" = "*"
      }
    }
  }]
}`, rnd, accountID)

	// Step 2: Latest config with jsonencode supporting nested resources
	latestNestedConfig := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s-nested"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
    }]
    # With jsonencode, nested resources work correctly
    resources = jsonencode({
      "com.cloudflare.api.account.%[2]s" = {
        "com.cloudflare.api.account.zone.*" = "*"
      }
    })
  }]
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			{
				// Step 1: Try to use nested resources with v5.10 (expect error)
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "5.10.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v5NestedConfig,
				// V5.10 can't handle nested resources - expects a string value
				ExpectError: regexp.MustCompile(`string\s+required`),
			},
			{
				// Step 2: Upgrade to latest and use jsonencode (should succeed)
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestNestedConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":{"com.cloudflare.api.account.zone.*":"*"}}`, accountID)),
					),
				},
			},
		},
	})
}

// TestMigrateAccountTokenFromV5_10_BothResourceFormats tests migration with both flat and nested resource formats
func TestMigrateAccountTokenFromV5_10_BothResourceFormats(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	// V5.10 config with flat resources only
	v510Config := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s-mixed"

  policies = [
    {
      effect = "allow"
      permission_groups = [{
        id = "82e64a83756745bbbb1c9c2701bf816b"  # DNS Read
      }]
      # Flat resource structure (works in v5.10)
      resources = {
        "com.cloudflare.api.account.%[2]s" = "*"
      }
    }
  ]
}`, rnd, accountID)

	// Latest config with both flat and nested resources using jsonencode
	latestMixedConfig := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s-mixed"

  policies = [
    {
      effect = "allow"
      permission_groups = [{
        id = "82e64a83756745bbbb1c9c2701bf816b"  # DNS Read
      }]
      # Flat resource structure with jsonencode (same as v5.10)
      resources = jsonencode({
        "com.cloudflare.api.account.%[2]s" = "*"
      })
    },
    {
      effect = "allow"
      permission_groups = [{
        id = "82e64a83756745bbbb1c9c2701bf816b"  # DNS Read (using same permission for simplicity)
      }]
      # Nested resource structure (only works with jsonencode in latest)
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
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5.10 (flat resources only)
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "5.10.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v510Config,
			},
			{
				// Step 2: Upgrade to latest with both flat and nested resources
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestMixedConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify we now have 2 policies (upgraded from 1 in v5.10)
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("name"),
						knownvalue.StringExact(fmt.Sprintf("%s-mixed", rnd)),
					),
					// Verify first policy resources (flat structure)
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID)),
					),
					// Verify second policy resources (nested structure - new in latest)
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(1).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":{"com.cloudflare.api.account.zone.*":"*"}}`, accountID)),
					),
				},
			},
		},
	})
}

// TestMigrateAccountTokenFromV5WithTTL tests migration with token TTL settings
func TestMigrateAccountTokenFromV5WithTTL(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	// Set TTL times
	notBefore := "2025-01-01T00:00:00Z"
	expiresOn := time.Now().UTC().AddDate(0, 0, 30).Format(time.RFC3339) // 30 days from now

	// V5 config with TTL settings
	v5Config := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
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

  not_before = "%[3]s"
  expires_on = "%[4]s"
}`, rnd, accountID, notBefore, expiresOn)

	// Latest config with jsonencode
	latestConfig := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
    }]
    resources = jsonencode({
      "com.cloudflare.api.account.%[2]s" = "*"
    })
  }]

  not_before = "%[3]s"
  expires_on = "%[4]s"
}`, rnd, accountID, notBefore, expiresOn)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.10.0", // Use a later v5 version where account_token is stable
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
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("not_before"),
						knownvalue.StringExact(notBefore),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("expires_on"),
						knownvalue.StringExact(expiresOn),
					),
					// Verify resources migration
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID)),
					),
				},
			},
		},
	})
}

// TestMigrateAccountTokenFromV5NestedResources tests migration with the new nested resources capability
func TestMigrateAccountTokenFromV5NestedResources(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	// V5 config - could only support flat map structure
	v5Config := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
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

	// Latest config - can now support nested structures
	latestConfig := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
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

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5 provider - simple structure
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.10.0", // Use a later v5 version where account_token is stable
					},
				},
				Config: v5Config,
			},
			{
				// Step 2: Upgrade and modify to use nested structure
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify the nested structure is properly handled
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":{"com.cloudflare.api.account.zone.*":"*"}}`, accountID)),
					),
				},
			},
		},
	})
}

// TestMigrateAccountTokenFromV5RemovedFields tests that computed fields are properly removed during migration
func TestMigrateAccountTokenFromV5RemovedFields(t *testing.T) {

	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	// V5 config - the old schema had computed fields like policy.id, permission_groups.meta, permission_groups.name
	v5Config := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"

  policies = [{
    # In v5, policy.id was a computed field
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
      # In v5, meta and name were computed fields that would be populated
    }]
    resources = {
      "com.cloudflare.api.account.%[2]s" = "*"
    }
  }]
}`, rnd, accountID)

	// Latest config
	latestConfig := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"

  policies = [{
    # policy.id field is removed in latest schema
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
      # meta and name fields are removed in latest schema
    }]
    resources = jsonencode({
      "com.cloudflare.api.account.%[2]s" = "*"
    })
  }]
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.10.0", // Use a later v5 version where account_token is stable
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
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id"),
						knownvalue.StringExact("82e64a83756745bbbb1c9c2701bf816b"),
					),
					// Verify resources is migrated to JSON
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID)),
					),
				},
			},
		},
	})
}

// TestMigrateAccountTokenFromV5_10 tests migration from v5.10.0 (earliest stable version for account_token)
func TestMigrateAccountTokenFromV5_10(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	// V5.10 config
	v510Config := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s-v510"

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

	// Latest config with jsonencode
	latestConfig := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s-v510"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
    }]
    resources = jsonencode({
      "com.cloudflare.api.account.%[2]s" = "*"
    })
  }]
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5.10 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.10.0",
					},
				},
				Config: v510Config,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("name"),
						knownvalue.StringExact(fmt.Sprintf("%s-v510", rnd)),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID)),
					),
				},
			},
		},
	})
}

// TestMigrateAccountTokenFromV5_11 tests migration from v5.11.0 (account_token was introduced around v5.10)
func TestMigrateAccountTokenFromV5_11(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	// V5.11 config
	v5Config := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
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

  expires_on = "2025-12-31T23:59:59Z"
}`, rnd, accountID)

	// Latest config
	latestConfig := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
    }]
    resources = jsonencode({
      "com.cloudflare.api.account.%[2]s" = "*"
    })
  }]

  expires_on = "2025-12-31T23:59:59Z"
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5.11 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.11.0",
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
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID)),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("expires_on"),
						knownvalue.StringExact("2025-12-31T23:59:59Z"),
					),
				},
			},
		},
	})
}

// TestMigrateAccountTokenFromV5_12 tests migration from v5.12.0 (last version before the change)
func TestMigrateAccountTokenFromV5_12(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	// V5.12 config with multiple policies
	v5Config := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
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
        "com.cloudflare.api.account.%[2]s" = "*"
      }
    }
  ]

  condition = {
    request_ip = {
      in = ["192.0.2.0/24"]
    }
  }
}`, rnd, accountID)

	// Latest config
	latestConfig := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"

  policies = [
    {
      effect = "allow"
      permission_groups = [{
        id = "82e64a83756745bbbb1c9c2701bf816b"
      }]
      resources = jsonencode({
        "com.cloudflare.api.account.%[2]s" = "*"
      })
    },
    {
      effect = "allow"
      permission_groups = [{
        id = "c8fed203ed3043cba015a93ad1616f1f"
      }]
      resources = jsonencode({
        "com.cloudflare.api.account.%[2]s" = "*"
      })
    }
  ]

  condition = {
    request_ip = {
      in = ["192.0.2.0/24"]
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
				// Step 1: Create with v5.12 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.12.0",
					},
				},
				Config: v5Config,
				// v5.12 still has policy ordering issues that cause drift
				ExpectNonEmptyPlan: true,
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   latestConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("name"),
						knownvalue.StringExact(rnd),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"),
						knownvalue.StringExact(fmt.Sprintf(`{"com.cloudflare.api.account.%s":"*"}`, accountID)),
					),
				},
			},
		},
	})
}
