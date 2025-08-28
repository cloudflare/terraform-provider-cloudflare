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

// Config generators for different provider versions
func accessPolicyConfigV4Basic(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_policy" "%[1]s" {
  account_id       = "%[2]s"
  name            = "%[1]s"
  decision        = "allow"
  session_duration = "24h"
  
  include {
    email = ["test@example.com"]
  }
}`, rnd, accountID)
}

func accessPolicyConfigV4Complex(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_policy" "%[1]s" {
  account_id       = "%[2]s"
  name            = "%[1]s"
  decision        = "allow"
  session_duration = "24h"
  
  include {
    email = ["user1@example.com", "user2@example.com"]
    ip    = ["192.168.1.0/24", "10.0.0.0/8"]
  }
  
  exclude {
    email = ["blocked@example.com"]
  }
  
  require {
    email_domain = ["example.com"]
  }
}`, rnd, accountID)
}

func accessPolicyConfigV4WithApproval(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_policy" "%[1]s" {
  account_id                    = "%[2]s"
  name                         = "%[1]s"
  decision                     = "allow"
  session_duration             = "24h"
  approval_required            = true
  purpose_justification_required = true
  purpose_justification_prompt = "Why do you need access?"
  
  approval_group {
    approvals_needed = 2
    email_addresses  = ["admin1@example.com", "admin2@example.com"]
  }
  
  include {
    email = ["test@example.com"]
  }
}`, rnd, accountID)
}

func accessPolicyConfigV5Basic(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_access_policy" "%[1]s" {
  account_id       = "%[2]s"
  name            = "%[1]s"
  decision        = "allow"
  session_duration = "24h"
  
  include = [
    {
      email = {
        email = "test@example.com"
      }
    }
  ]
}`, rnd, accountID)
}

func accessPolicyConfigV5Complex(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_access_policy" "%[1]s" {
  account_id       = "%[2]s"
  name            = "%[1]s"
  decision        = "allow"
  session_duration = "24h"
  
  include = [
    {
      email = {
        email = "user1@example.com"
      }
    },
    {
      email = {
        email = "user2@example.com"
      }
    },
    {
      ip = {
        ip = "192.168.1.0/24"
      }
    },
    {
      ip = {
        ip = "10.0.0.0/8"
      }
    }
  ]
  
  exclude = [
    {
      email = {
        email = "blocked@example.com"
      }
    }
  ]
  
  require = [
    {
      email_domain = {
        domain = "example.com"
      }
    }
  ]
}`, rnd, accountID)
}

func accessPolicyConfigV5WithApproval(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_access_policy" "%[1]s" {
  account_id                    = "%[2]s"
  name                         = "%[1]s"
  decision                     = "allow"
  session_duration             = "24h"
  approval_required            = true
  purpose_justification_required = true
  purpose_justification_prompt = "Why do you need access?"
  
  approval_groups = [
    {
      approvals_needed = 2
      email_addresses  = ["admin1@example.com", "admin2@example.com"]
    }
  ]
  
  include = [
    {
      email = {
        email = "test@example.com"
      }
    }
  ]
}`, rnd, accountID)
}

// TestMigrateZeroTrustAccessPolicyMultiVersion_Basic tests basic migration from multiple versions
func TestMigrateZeroTrustAccessPolicyMultiVersion_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	testCases := []struct {
		name       string
		version    string
		configFunc func(rnd, accountID string) string
	}{
		{
			name:       "from_v4_52_1",
			version:    "4.52.1",
			configFunc: accessPolicyConfigV4Basic,
		},
		{
			name:       "from_v5_0_0",
			version:    "5.0.0",
			configFunc: accessPolicyConfigV5Basic,
		},
		{
			name:       "from_v5_8_4", // Current stable release
			version:    "5.8.4",
			configFunc: accessPolicyConfigV5Basic,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_policy." + rnd
			tmpDir := t.TempDir()

			config := tc.configFunc(rnd, accountID)

			// Build test steps
			steps := []resource.TestStep{}

			// Step 1: Create with specific provider version
			steps = append(steps, resource.TestStep{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: tc.version,
					},
				},
				Config: config,
			})

			// Step 2: Run migration (for v4) or just upgrade provider (for v5)
			steps = append(steps,
				acctest.MigrationTestStep(t, config, tmpDir, tc.version, []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					// Verify email transformation: v4 email list -> v5 single nested object
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(1)),
				}),
			)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps:      steps,
			})
		})
	}
}

// TestMigrateZeroTrustAccessPolicyMultiVersion_Complex tests migration with complex conditions
func TestMigrateZeroTrustAccessPolicyMultiVersion_Complex(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	testCases := []struct {
		name       string
		version    string
		configFunc func(rnd, accountID string) string
	}{
		{
			name:       "from_v4_52_1",
			version:    "4.52.1",
			configFunc: accessPolicyConfigV4Complex,
		},
		{
			name:       "from_v5_0_0",
			version:    "5.0.0",
			configFunc: accessPolicyConfigV5Complex,
		},
		{
			name:       "from_v5_8_4", // Current stable release
			version:    "5.8.4",
			configFunc: accessPolicyConfigV5Complex,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_policy." + rnd
			tmpDir := t.TempDir()

			config := tc.configFunc(rnd, accountID)

			// Build test steps
			steps := []resource.TestStep{}

			// Step 1: Create with specific provider version
			steps = append(steps, resource.TestStep{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: tc.version,
					},
				},
				Config: config,
			})

			// Step 2: Run migration and verify transformation
			steps = append(steps,
				acctest.MigrationTestStep(t, config, tmpDir, tc.version, []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
					// Verify complex condition transformations - multiple objects for multiple values
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(4)), // email(2) + ip(2)
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude"), knownvalue.ListSizeExact(1)), // email(1)
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("require"), knownvalue.ListSizeExact(1)), // email_domain(1)
				}),
			)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps:      steps,
			})
		})
	}
}

// TestMigrateZeroTrustAccessPolicyMultiVersion_WithApproval tests migration with approval groups
func TestMigrateZeroTrustAccessPolicyMultiVersion_WithApproval(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	testCases := []struct {
		name       string
		version    string
		configFunc func(rnd, accountID string) string
	}{
		{
			name:       "from_v4_52_1",
			version:    "4.52.1",
			configFunc: accessPolicyConfigV4WithApproval,
		},
		{
			name:       "from_v5_0_0",
			version:    "5.0.0",
			configFunc: accessPolicyConfigV5WithApproval,
		},
		{
			name:       "from_v5_8_4", // Current stable release
			version:    "5.8.4",
			configFunc: accessPolicyConfigV5WithApproval,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_policy." + rnd
			tmpDir := t.TempDir()

			config := tc.configFunc(rnd, accountID)

			// Build test steps
			steps := []resource.TestStep{}

			// Step 1: Create with specific provider version
			steps = append(steps, resource.TestStep{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: tc.version,
					},
				},
				Config: config,
			})

			// Step 2: Run migration and verify transformation
			steps = append(steps,
				acctest.MigrationTestStep(t, config, tmpDir, tc.version, []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("approval_required"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("purpose_justification_required"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("purpose_justification_prompt"), knownvalue.StringExact("Why do you need access?")),
					// Verify approval_group -> approval_groups transformation
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("approval_groups"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(1)),
				}),
			)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps:      steps,
			})
		})
	}
}

// TestMigrateZeroTrustAccessPolicy_EdgeCases tests various edge cases in migration
func TestMigrateZeroTrustAccessPolicy_EdgeCases(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	// Test edge cases with different rule combinations
	edgeCases := []struct {
		name           string
		version        string
		configFunc     func(rnd, accountID string) string
		expectedChecks func(resourceName, accountID, rnd string) []statecheck.StateCheck
	}{
		{
			name:       "v4_single_email",
			version:    "4.52.1",
			configFunc: accessPolicyConfigV4Basic,
			expectedChecks: func(resourceName, accountID, rnd string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("decision"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(1)),
				}
			},
		},
		{
			name:    "v4_multiple_values_per_type",
			version: "4.52.1",
			configFunc: func(rnd, accountID string) string {
				return fmt.Sprintf(`
resource "cloudflare_access_policy" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  decision   = "allow"
  
  include {
    email        = ["email1@example.com", "email2@example.com", "email3@example.com"]
    email_domain = ["example.com", "test.com"]
  }
}`, rnd, accountID)
			},
			expectedChecks: func(resourceName, accountID, rnd string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					// 3 emails + 2 domains = 5 include objects
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(5)),
				}
			},
		},
		{
			name:       "v5_basic_remains_unchanged",
			version:    "5.0.0",
			configFunc: accessPolicyConfigV5Basic,
			expectedChecks: func(resourceName, accountID, rnd string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(1)),
				}
			},
		},
		{
			name:       "v5_complex_remains_unchanged",
			version:    "5.8.4",
			configFunc: accessPolicyConfigV5Complex,
			expectedChecks: func(resourceName, accountID, rnd string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(4)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("require"), knownvalue.ListSizeExact(1)),
				}
			},
		},
		{
			name:    "v4_with_boolean_fields",
			version: "4.52.1",
			configFunc: func(rnd, accountID string) string {
				return fmt.Sprintf(`
resource "cloudflare_access_policy" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  decision   = "allow"
  
  include {
    everyone                = true
    any_valid_service_token = true
  }
  
  require {
    certificate = true
  }
}`, rnd, accountID)
			},
			expectedChecks: func(resourceName, accountID, rnd string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					// Boolean fields convert to single objects
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(2)), // everyone + any_valid_service_token
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("require"), knownvalue.ListSizeExact(1)), // certificate
				}
			},
		},
		{
			name:    "v4_with_geo",
			version: "4.52.1",
			configFunc: func(rnd, accountID string) string {
				return fmt.Sprintf(`
resource "cloudflare_access_policy" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  decision   = "allow"
  
  include {
    geo = ["US", "GB", "DE"]
  }
}`, rnd, accountID)
			},
			expectedChecks: func(resourceName, accountID, rnd string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					// 3 geo locations = 3 include objects
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(3)),
				}
			},
		},
	}

	for _, tc := range edgeCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_policy." + rnd
			tmpDir := t.TempDir()

			config := tc.configFunc(rnd, accountID)
			expectedChecks := tc.expectedChecks(resourceName, accountID, rnd)

			// Build test steps
			steps := []resource.TestStep{}

			// Step 1: Create with specific provider version
			steps = append(steps, resource.TestStep{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: tc.version,
					},
				},
				Config: config,
			})

			// Step 2: Run migration (for v4) or just upgrade provider (for v5)
			steps = append(steps,
				acctest.MigrationTestStep(t, config, tmpDir, tc.version, expectedChecks),
			)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps:      steps,
			})
		})
	}
}
