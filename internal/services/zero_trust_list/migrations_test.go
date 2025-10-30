package zero_trust_list_test

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

// TestMigrateZeroTrustList_V4ToV5_SimpleItems tests basic migration with simple items array
func TestMigrateZeroTrustList_V4ToV5_SimpleItems(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V4 config with simple items array
	v4Config := fmt.Sprintf(`
resource "cloudflare_teams_list" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-%[1]s"
  type       = "IP"
  items      = ["192.0.2.1", "192.0.2.2"]
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
						VersionConstraint: "4.48.0",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.48.0", "v4", "v5", []statecheck.StateCheck{
				// Resource should be renamed to cloudflare_zero_trust_list
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("tf-acc-test-"+rnd)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("IP")),
				// Items should be transformed from string array to object array
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("items"), knownvalue.SetSizeExact(2)),
			}),
		},
	})
}

// TestMigrateZeroTrustList_V4ToV5_ItemsWithDescription tests migration with items_with_description
func TestMigrateZeroTrustList_V4ToV5_ItemsWithDescription(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V4 config with items_with_description blocks
	v4Config := fmt.Sprintf(`
resource "cloudflare_teams_list" "%[1]s" {
  account_id  = "%[2]s"
  name        = "tf-acc-test-%[1]s"
  type        = "DOMAIN"
  description = "Test list with descriptions"
  
  items_with_description {
    value       = "example.com"
    description = "Main domain"
  }
  
  items_with_description {
    value       = "test.example.com"
    description = "Test subdomain"
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
						VersionConstraint: "4.48.0",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.48.0", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("tf-acc-test-"+rnd)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("DOMAIN")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("description"), knownvalue.StringExact("Test list with descriptions")),
				// Items should contain the merged items_with_description blocks
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("items"), knownvalue.SetSizeExact(2)),
			}),
		},
	})
}

// TestMigrateZeroTrustList_V4ToV5_MixedItems tests migration with both items and items_with_description
func TestMigrateZeroTrustList_V4ToV5_MixedItems(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V4 config with both items and items_with_description
	v4Config := fmt.Sprintf(`
resource "cloudflare_teams_list" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-%[1]s"
  type       = "IP"
  items      = ["192.0.2.1", "192.0.2.2"]
  
  items_with_description {
    value       = "192.0.2.3"
    description = "Special IP"
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
						VersionConstraint: "4.48.0",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.48.0", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("tf-acc-test-"+rnd)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("IP")),
				// Items should contain both regular items and items_with_description
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("items"), knownvalue.SetSizeExact(3)),
			}),
		},
	})
}

// TestMigrateZeroTrustList_V4ToV5_EmptyList tests migration with empty items list
func TestMigrateZeroTrustList_V4ToV5_EmptyList(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V4 config with empty items
	v4Config := fmt.Sprintf(`
resource "cloudflare_teams_list" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-%[1]s"
  type       = "SERIAL"
  items      = []
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
						VersionConstraint: "4.48.0",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.48.0", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("tf-acc-test-"+rnd)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("SERIAL")),
				// Items should be nil when empty (v4 stores empty as nil, v5 should match)
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("items"), knownvalue.Null()),
			}),
		},
	})
}

// TestMigrateZeroTrustList_V4ToV5_EmailList tests migration with EMAIL type list
func TestMigrateZeroTrustList_V4ToV5_EmailList(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V4 config with EMAIL type
	v4Config := fmt.Sprintf(`
resource "cloudflare_teams_list" "%[1]s" {
  account_id  = "%[2]s"
  name        = "tf-acc-test-%[1]s"
  type        = "EMAIL"
  description = "Email allowlist for testing"
  items       = ["user1@example.com", "user2@example.com", "admin@company.org"]
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
						VersionConstraint: "4.48.0",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.48.0", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("tf-acc-test-"+rnd)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("EMAIL")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("description"), knownvalue.StringExact("Email allowlist for testing")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("items"), knownvalue.SetSizeExact(3)),
			}),
		},
	})
}

// TestMigrateZeroTrustList_V4ToV5_URLList tests migration with URL type list
func TestMigrateZeroTrustList_V4ToV5_URLList(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V4 config with URL type and items_with_description
	// Using URLs with query parameters or fragments to avoid normalization
	v4Config := fmt.Sprintf(`
resource "cloudflare_teams_list" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-%[1]s"
  type       = "URL"
  
  items_with_description {
    value       = "https://example.com/admin/index.html"
    description = "Admin portal"
  }
  
  items_with_description {
    value       = "https://api.example.com/v1/users"
    description = "API endpoint"
  }
  
  items_with_description {
    value       = "https://test.example.com/app"
    description = "Test environment"
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
						VersionConstraint: "4.48.0",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.48.0", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("tf-acc-test-"+rnd)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("URL")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("items"), knownvalue.SetSizeExact(3)),
			}),
		},
	})
}

// TestMigrateZeroTrustList_V4ToV5_LargeList tests migration with many items
func TestMigrateZeroTrustList_V4ToV5_LargeList(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V4 config with many IP addresses
	v4Config := fmt.Sprintf(`
resource "cloudflare_teams_list" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-%[1]s"
  type       = "IP"
  items      = [
    "10.0.0.1", "10.0.0.2", "10.0.0.3", "10.0.0.4", "10.0.0.5",
    "10.0.0.6", "10.0.0.7", "10.0.0.8", "10.0.0.9", "10.0.0.10",
    "192.168.1.0/24", "192.168.2.0/24", "192.168.3.0/24",
    "172.16.0.0/16", "172.17.0.0/16"
  ]
  
  items_with_description {
    value       = "203.0.113.0/24"
    description = "Documentation range"
  }
  
  items_with_description {
    value       = "198.51.100.0/24"
    description = "Test network"
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
						VersionConstraint: "4.48.0",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.48.0", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("tf-acc-test-"+rnd)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("IP")),
				// Should have 15 regular items + 2 items_with_description = 17 total
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("items"), knownvalue.SetSizeExact(17)),
			}),
		},
	})
}

// TestMigrateZeroTrustList_V4ToV5_SpecialCharacters tests migration with special characters in values
func TestMigrateZeroTrustList_V4ToV5_SpecialCharacters(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V4 config with special characters in domains (but valid domain names)
	v4Config := fmt.Sprintf(`
resource "cloudflare_teams_list" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-%[1]s"
  type       = "DOMAIN"
  items      = [
    "sub-domain.example.com",
    "test-with-hyphens.com",
    "domain123.example.org"
  ]
  
  items_with_description {
    value       = "special-chars-123.example.org"
    description = "Domain with numbers and hyphens"
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
						VersionConstraint: "4.48.0",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.48.0", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("tf-acc-test-"+rnd)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("DOMAIN")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("items"), knownvalue.SetSizeExact(4)),
			}),
		},
	})
}