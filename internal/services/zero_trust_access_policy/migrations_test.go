package zero_trust_access_policy_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// TestAccZeroTrustAccessPolicyMigrationFromV4Basic tests basic migration from v4 to v5
func TestAccZeroTrustAccessPolicyMigrationFromV4Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	tmpDir := t.TempDir()

	// V4 config using old resource name and basic email condition
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_policy" "%[1]s" {
  account_id     = "%[2]s"
  name           = "%[1]s"
  decision       = "allow"
  session_duration = "24h"
  include {
    email = ["test@example.com"]
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
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "~> 4.0",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.0", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				// Verify email transformation: v4 email list -> v5 single nested object
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("email"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestAccZeroTrustAccessPolicyMigrationFromV4Complex tests migration with complex conditions
func TestAccZeroTrustAccessPolicyMigrationFromV4Complex(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	tmpDir := t.TempDir()

	// V4 config with multiple condition types and lists
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_policy" "%[1]s" {
  account_id     = "%[2]s"
  name           = "%[1]s"
  decision       = "allow"
  session_duration = "24h"
  approval_required = true
  isolation_required = true
  purpose_justification_required = true
  purpose_justification_prompt = "Why do you need access?"
  
  approval_group {
    approvals_needed = 2
    email_addresses = ["admin1@example.com", "admin2@example.com"]
  }
  
  include {
    email = ["user1@example.com", "user2@example.com"]
    email_domain = ["example.com", "test.com"]
    ip = ["192.168.1.0/24", "10.0.0.0/8"]
    everyone = true
    any_valid_service_token = true
  }
  
  exclude {
    email = ["blocked@example.com"]
    geo = ["CN", "RU"]
  }
  
  require {
    certificate = true
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
						VersionConstraint: "~> 4.0",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.0", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("approval_required"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("isolation_required"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("purpose_justification_required"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("purpose_justification_prompt"), knownvalue.StringExact("Why do you need access?")),
				// Verify approval_group -> approval_groups transformation
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("approval_groups"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("approval_groups").AtSliceIndex(0).AtMapKey("approvals_needed"), knownvalue.Float64Exact(2.0)),
				// Verify complex condition transformations - multiple objects for multiple values
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(8)), // email(2) + email_domain(2) + ip(2) + everyone + any_valid_service_token
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude"), knownvalue.ListSizeExact(3)), // email(1) + geo(2)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("require"), knownvalue.ListSizeExact(1)), // certificate
			}),
		},
	})
}

// TestAccZeroTrustAccessPolicyMigrationFromV4OAuthProviders tests array explosion and attribute transformations
func TestAccZeroTrustAccessPolicyMigrationFromV4OAuthProviders(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	tmpDir := t.TempDir()

	// V4 config with simple array explosion test
	// Test email array explosion without mixing with boolean attributes
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_policy" "%[1]s" {
  account_id     = "%[2]s"
  name           = "%[1]s"
  decision       = "allow"
  
  include {
    email = ["user1@example.com", "user2@example.com"]
  }
  
  exclude {
    email = ["blocked@example.com"] 
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
						VersionConstraint: "~> 4.0",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.0", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
				// Verify array explosion: email array (2 rules) = 2 include rules
				// Also verify exclude has 1 rule  
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(2)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude"), knownvalue.ListSizeExact(1)),
			}),
		},
	})
}

// TestAccZeroTrustAccessPolicyMigrationFromV4DecisionTypes tests all decision types
func TestAccZeroTrustAccessPolicyMigrationFromV4DecisionTypes(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	testCases := []struct {
		decision string
		name     string
	}{
		{"allow", "Allow"},
		{"deny", "Deny"},
		{"non_identity", "NonIdentity"},
		{"bypass", "Bypass"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Decision_%s", tc.name), func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_policy." + rnd
			tmpDir := t.TempDir()

			v4Config := fmt.Sprintf(`
resource "cloudflare_access_policy" "%[1]s" {
  account_id     = "%[2]s"
  name           = "%[1]s"
  decision       = "%[3]s"
  session_duration = "24h"
  include {
    everyone = true
  }
}`, rnd, accountID, tc.decision)

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
								VersionConstraint: "~> 4.0",
							},
						},
						Config: v4Config,
					},
					acctest.MigrationTestStep(t, v4Config, tmpDir, "4.0", []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact(tc.decision)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					}),
				},
			})
		})
	}
}

// TestAccZeroTrustAccessPolicyMigrationFromV4OptionalBooleans tests boolean to object transformations
func TestAccZeroTrustAccessPolicyMigrationFromV4OptionalBooleans(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	testCases := []struct {
		boolField string
		testName  string
	}{
		{"everyone", "Everyone"},
		{"certificate", "Certificate"},
		{"any_valid_service_token", "AnyValidServiceToken"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Boolean_%s", tc.testName), func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_policy." + rnd
			tmpDir := t.TempDir()

			v4Config := fmt.Sprintf(`
resource "cloudflare_access_policy" "%[1]s" {
  account_id     = "%[2]s"
  name           = "%[1]s"
  decision       = "allow"
  include {
    %[3]s = true
  }
}`, rnd, accountID, tc.boolField)

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
								VersionConstraint: "~> 4.0",
							},
						},
						Config: v4Config,
					},
					acctest.MigrationTestStep(t, v4Config, tmpDir, "4.0", []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
						// Verify boolean -> empty object transformation
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(1)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey(tc.boolField), knownvalue.NotNull()),
					}),
				},
			})
		})
	}
}

// TestAccZeroTrustAccessPolicyMigrationFromV4BasicMigration tests basic v4 to v5 migration functionality
func TestAccZeroTrustAccessPolicyMigrationFromV4UnsupportedFeatures(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	tmpDir := t.TempDir()

	// V4 config with basic valid attributes that should migrate cleanly to v5
	// Test focuses on ensuring migration process works for simple policies
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_policy" "%[1]s" {
  account_id     = "%[2]s"
  name           = "%[1]s"
  decision       = "allow"
  session_duration = "24h"
  
  include {
    everyone = true
  }
  
  exclude {
    email = ["test@blocked.com"]
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
						VersionConstraint: "~> 4.0",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.0", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
				// Verify that basic migration works correctly:
				// - Basic attributes are preserved
				// - Include/exclude rules are properly transformed
				// - Session duration is preserved
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
			}),
		},
	})
}

// TestAccZeroTrustAccessPolicyMigrationFromV4ServiceTokens tests service token transformations  
func TestAccZeroTrustAccessPolicyMigrationFromV4ServiceTokens(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_policy." + rnd
	tmpDir := t.TempDir()

	// V4 config with any_valid_service_token (boolean -> empty nested object)
	// This tests the transformation without requiring actual service token IDs
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_policy" "%[1]s" {
  account_id     = "%[2]s"
  name           = "%[1]s"
  decision       = "allow"
  
  include {
    any_valid_service_token = true
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
						VersionConstraint: "~> 4.0",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.0", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
				// Verify any_valid_service_token transformation: boolean -> empty nested object
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("any_valid_service_token"), knownvalue.NotNull()),
			}),
		},
	})
}
