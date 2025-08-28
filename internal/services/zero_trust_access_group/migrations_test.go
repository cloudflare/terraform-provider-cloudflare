package zero_trust_access_group_test

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

// Config generators for different provider versions
func accessGroupConfigV4Basic(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  
  include {
    email = ["test@example.com"]
  }
}`, rnd, accountID)
}

func accessGroupConfigV4MultiValue(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  
  include {
    email = ["user1@example.com", "user2@example.com"]
    ip    = ["192.0.2.1/32", "192.0.2.2/32"]
  }
  
  exclude {
    email_domain = ["blocked.com"]
  }
  
  require {
    ip = ["10.0.0.0/8"]
  }
}`, rnd, accountID)
}

func accessGroupConfigV4ZoneScoped(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_group" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[1]s"
  
  include {
    email = ["test@example.com"]
    ip    = ["192.0.2.0/24"]
  }
}`, rnd, zoneID)
}

func accessGroupConfigV5Basic(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_access_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  
  include = [
    {
      email = {
        email = "test@example.com"
      }
    }
  ]
}`, rnd, accountID)
}

func accessGroupConfigV5MultiValue(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_access_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  
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
        ip = "192.0.2.1/32"
      }
    },
    {
      ip = {
        ip = "192.0.2.2/32"
      }
    }
  ]
  
  exclude = [
    {
      email_domain = {
        domain = "blocked.com"
      }
    }
  ]
  
  require = [
    {
      ip = {
        ip = "10.0.0.0/8"
      }
    }
  ]
}`, rnd, accountID)
}

func accessGroupConfigV5ZoneScoped(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_access_group" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[1]s"
  
  include = [
    {
      email = {
        email = "test@example.com"
      }
    },
    {
      ip = {
        ip = "192.0.2.0/24"
      }
    }
  ]
}`, rnd, zoneID)
}

// TestMigrateZeroTrustAccessGroupMultiVersion_Basic tests basic migration from multiple versions
func TestMigrateZeroTrustAccessGroupMultiVersion_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
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
			configFunc: accessGroupConfigV4Basic,
		},
		{
			name:       "from_v5_0_0",
			version:    "5.0.0",
			configFunc: accessGroupConfigV5Basic,
		},
		{
			name:       "from_v5_8_4", // Current stable release
			version:    "5.8.4",
			configFunc: accessGroupConfigV5Basic,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_group." + rnd
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
			// Use special migration function for access group
			steps = append(steps,
				acctest.ZeroTrustAccessGroupMigrationTestStep(t, config, tmpDir, tc.version, []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
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

// TestMigrateZeroTrustAccessGroupMultiVersion_ComplexRules tests migration with multiple values
func TestMigrateZeroTrustAccessGroupMultiVersion_ComplexRules(t *testing.T) {
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
			configFunc: accessGroupConfigV4MultiValue,
		},
		{
			name:       "from_v5_0_0",
			version:    "5.0.0",
			configFunc: accessGroupConfigV5MultiValue,
		},
		{
			name:       "from_v5_8_4", // Current stable release
			version:    "5.8.4",
			configFunc: accessGroupConfigV5MultiValue,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_group." + rnd
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
				acctest.ZeroTrustAccessGroupMigrationTestStep(t, config, tmpDir, tc.version, []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					// Verify expansion: 2 emails + 2 IPs = 4 include objects
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(4)),
					// Verify exclude: 1 domain = 1 exclude object
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude"), knownvalue.ListSizeExact(1)),
					// Verify require: 1 IP = 1 require object
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("require"), knownvalue.ListSizeExact(1)),
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

// TestMigrateZeroTrustAccessGroupMultiVersion_ZoneScoped tests zone-level migration
func TestMigrateZeroTrustAccessGroupMultiVersion_ZoneScoped(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	testCases := []struct {
		name       string
		version    string
		configFunc func(rnd, zoneID string) string
	}{
		{
			name:       "from_v4_52_1",
			version:    "4.52.1",
			configFunc: accessGroupConfigV4ZoneScoped,
		},
		{
			name:       "from_v5_0_0",
			version:    "5.0.0",
			configFunc: accessGroupConfigV5ZoneScoped,
		},
		{
			name:       "from_v5_8_4", // Current stable release
			version:    "5.8.4",
			configFunc: accessGroupConfigV5ZoneScoped,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_group." + rnd
			tmpDir := t.TempDir()

			config := tc.configFunc(rnd, zoneID)

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

			// Step 2: Run migration and verify zone context preserved
			steps = append(steps,
				acctest.ZeroTrustAccessGroupMigrationTestStep(t, config, tmpDir, tc.version, []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(2)),
				}),
			)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps:      steps,
			})
		})
	}
}

// TestMigrateZeroTrustAccessGroup_EdgeCases tests various edge cases in migration
func TestMigrateZeroTrustAccessGroup_EdgeCases(t *testing.T) {
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
			configFunc: accessGroupConfigV4Basic,
			expectedChecks: func(resourceName, accountID, rnd string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(1)),
				}
			},
		},
		{
			name:    "v4_multiple_emails_and_ips",
			version: "4.52.1",
			configFunc: func(rnd, accountID string) string {
				return fmt.Sprintf(`
resource "cloudflare_access_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  
  include {
    email        = ["email1@example.com", "email2@example.com", "email3@example.com"]
    email_domain = ["example.com", "test.com"]
    ip          = ["192.0.2.1/32"]
  }
}`, rnd, accountID)
			},
			expectedChecks: func(resourceName, accountID, rnd string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					// 3 emails + 2 domains + 1 IP = 6 include objects
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(6)),
				}
			},
		},
		{
			name:       "v5_basic_remains_unchanged",
			version:    "5.0.0",
			configFunc: accessGroupConfigV5Basic,
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
			configFunc: accessGroupConfigV5MultiValue,
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
			name:    "v4_all_rule_types",
			version: "4.52.1",
			configFunc: func(rnd, accountID string) string {
				return fmt.Sprintf(`
resource "cloudflare_access_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  
  include {
    email = ["admin@company.com", "manager@company.com"]
    ip    = ["10.0.1.0/24", "10.0.2.0/24"]
  }
  
  exclude {
    email_domain = ["blocked.com"]
  }
  
  require {
    email_domain = ["company.com"]
  }
}`, rnd, accountID)
			},
			expectedChecks: func(resourceName, accountID, rnd string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					// Include: 2 emails + 2 IPs = 4 objects
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(4)),
					// Exclude: 1 domain = 1 object
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude"), knownvalue.ListSizeExact(1)),
					// Require: 1 domain = 1 object
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("require"), knownvalue.ListSizeExact(1)),
				}
			},
		},
	}

	for _, tc := range edgeCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_access_group." + rnd
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
				acctest.ZeroTrustAccessGroupMigrationTestStep(t, config, tmpDir, tc.version, expectedChecks),
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
