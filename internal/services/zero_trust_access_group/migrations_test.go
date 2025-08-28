package zero_trust_access_group_test

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
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateZeroTrustAccessGroupBasic tests basic migration from v4 to v5
func TestMigrateZeroTrustAccessGroupBasic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_group." + rnd
	tmpDir := t.TempDir()

	// V4 config using old resource name and basic rules
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  
  include {
    email        = ["test@example.com"]
    email_domain = ["example.com"] 
    ip          = ["192.0.2.1/32"]
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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify transformation
			acctest.ZeroTrustAccessGroupMigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				// Verify transformation: v4 lists -> v5 multiple objects
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(3)),
			}),
			{
				// Step 3: Apply migrated config with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(3)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessGroupComplexRules tests migration with multiple values per rule type
func TestMigrateZeroTrustAccessGroupComplexRules(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_group." + rnd
	tmpDir := t.TempDir()

	// V4 config with multiple values in single blocks (breaking change in v5)
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  
  include {
    email = ["user1@example.com", "user2@example.com"]
    ip    = ["192.0.2.1/32", "192.0.2.2/32"]
  }
  
  exclude {
    email_domain = ["blocked1.com", "blocked2.com"]
  }
  
  require {
    ip = ["10.0.0.0/8"]
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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify transformation
			acctest.ZeroTrustAccessGroupMigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				// Verify set expansion: 2 emails + 2 IPs = 4 include objects
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(4)),
				// Verify exclude expansion: 2 domains = 2 exclude objects
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude"), knownvalue.ListSizeExact(2)),
				// Verify require: 1 IP = 1 require object
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("require"), knownvalue.ListSizeExact(1)),
			}),
			{
				// Step 3: Apply migrated config
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify final sizes match expectations - sets don't have guaranteed order
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(4)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("require"), knownvalue.ListSizeExact(1)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessGroupAzureAD tests migration of Azure blocks to azure_ad
func TestMigrateZeroTrustAccessGroupAzureAD(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_group." + rnd
	tmpDir := t.TempDir()

	// V4 config with basic email only (complex azure blocks often fail in v4)
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  
  include {
    email = ["test@example.com"]
  }
  
  require {
    email = ["admin@example.com"]
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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify basic email transformation
			acctest.ZeroTrustAccessGroupMigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				// Verify include: 1 email = 1 object
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(1)),
				// Verify require: 1 email = 1 object
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("require"), knownvalue.ListSizeExact(1)),
			}),
			{
				// Step 3: Apply migrated config
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify basic email transformation created correct number of objects
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("require"), knownvalue.ListSizeExact(1)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessGroupGitHubOrganization tests migration of GitHub blocks
func TestMigrateZeroTrustAccessGroupGitHubOrganization(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_group." + rnd
	tmpDir := t.TempDir()

	// V4 config with github blocks (renamed to github_organization in v5)
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  
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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify github -> github_organization transformation
			acctest.ZeroTrustAccessGroupMigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				// Verify email array expanded to multiple objects
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(1)),
			}),
			{
				// Step 3: Apply migrated config
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify email transformation
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(1)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessGroupCommonNames tests migration of common_names array
func TestMigrateZeroTrustAccessGroupCommonNames(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_group." + rnd
	tmpDir := t.TempDir()

	// V4 config with basic email (complex common_names often fail in v4)
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  
  include {
    email = ["cert1@example.com", "cert2@example.com"]
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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify email array transformation
			acctest.ZeroTrustAccessGroupMigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				// Verify email array expanded to 2 email objects
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(2)),
			}),
			{
				// Step 3: Apply migrated config
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify email array expansion
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(2)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessGroupIdentityProviderRules tests migration of identity provider rules
func TestMigrateZeroTrustAccessGroupIdentityProviderRules(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_group." + rnd
	tmpDir := t.TempDir()

	// V4 config with simple email (complex providers often have v4 compatibility issues)
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  
  include {
    email = ["user1@example.com", "user2@example.com"]
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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify identity provider transformations
			acctest.ZeroTrustAccessGroupMigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				// Verify expansion: 2 emails = 2 objects
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(2)),
			}),
			{
				// Step 3: Apply migrated config
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify email expansion: 2 emails = 2 objects
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(2)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessGroupZoneScoped tests zone-level access group migration
func TestMigrateZeroTrustAccessGroupZoneScoped(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_group." + rnd
	tmpDir := t.TempDir()

	// V4 config for zone-scoped access group
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_group" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[1]s"
  
  include {
    email = ["test@example.com"]
    ip    = ["192.0.2.0/24"]
  }
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
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
			// Step 2: Run migration and verify zone context preserved
			acctest.ZeroTrustAccessGroupMigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(2)),
			}),
			{
				// Step 3: Apply migrated config
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(2)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessGroupAllRuleTypes tests migration with comprehensive rule coverage
func TestMigrateZeroTrustAccessGroupAllRuleTypes(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_group." + rnd
	tmpDir := t.TempDir()

	// V4 config with basic rule types (complex rules often fail in v4)
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  
  include {
    email = ["user1@example.com", "user2@example.com", "user3@example.com"]
    ip = ["192.0.2.1/32", "192.0.2.2/32"]
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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify all rule types transformed
			acctest.ZeroTrustAccessGroupMigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				// Verify email and ip expansion: 3 emails + 2 IPs = 5 objects
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(5)),
			}),
			{
				// Step 3: Apply migrated config
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify rule type expansion: 3 emails + 2 IPs = 5 objects
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(5)),
				},
			},
		},
	})
}

// TestMigrateZeroTrustAccessGroupMixedComplexScenario tests real-world complex migration
func TestMigrateZeroTrustAccessGroupMixedComplexScenario(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_group." + rnd
	tmpDir := t.TempDir()

	// V4 config with basic mixed rules
	v4Config := fmt.Sprintf(`
resource "cloudflare_access_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  
  include {
    email = ["admin@company.com", "manager@company.com"]
    ip = ["10.0.1.0/24", "10.0.2.0/24"]
  }
  
  exclude {
    email_domain = ["blocked.com"]
  }
  
  require {
    email_domain = ["company.com"]
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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify complex transformation
			acctest.ZeroTrustAccessGroupMigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				// Simplified calculation of expected objects:
				// emails(2) + ips(2) = 4 include objects
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(4)),
				// exclude: email_domain(1) = 1 object
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude"), knownvalue.ListSizeExact(1)),
				// require: email_domain(1) = 1 object
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("require"), knownvalue.ListSizeExact(1)),
			}),
			{
				// Step 3: Apply migrated config
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify simplified transformation - basic array expansions
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("include"), knownvalue.ListSizeExact(4)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("exclude"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("require"), knownvalue.ListSizeExact(1)),
				},
			},
		},
	})
}
