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

// TestMigrateZeroTrustAccessGroupMultiVersion_Basic tests basic migration from v4 to v5
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
						VersionConstraint: tc.version,
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: config,
			})

			// Step 2: Run migration (for v4) or just upgrade provider (for v5)
			// Use helper that handles state normalization for resources with nil field removal
			migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(1)),
			})
			steps = append(steps, migrationSteps...)

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

// TestMigrateZeroTrustAccessGroupMultiVersion_ComplexRules tests migration with multiple values from v4 to v5
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
						VersionConstraint: tc.version,
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: config,
			})

			// Step 2: Run migration and verify transformation
			migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				// Verify expansion: 2 emails + 2 IPs = 4 include objects
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(4)),
				// Verify exclude: 1 domain = 1 exclude object
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude"), knownvalue.ListSizeExact(1)),
				// Verify require: 1 IP = 1 require object
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("require"), knownvalue.ListSizeExact(1)),
			})
			steps = append(steps, migrationSteps...)

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

// TestMigrateZeroTrustAccessGroupMultiVersion_ZoneScoped tests zone-level migration from v4 to v5
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
						VersionConstraint: tc.version,
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: config,
			})

			// Step 2: Run migration and verify zone context preserved
			migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(2)),
			})
			steps = append(steps, migrationSteps...)

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

	// Test edge cases with different rule combinations from v4 to v5
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
						VersionConstraint: tc.version,
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: config,
			})

			// Step 2: Run migration (for v4) or just upgrade provider (for v5)
			migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, config, tmpDir, "4.52.1", "v4", "v5", expectedChecks)
			steps = append(steps, migrationSteps...)

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
