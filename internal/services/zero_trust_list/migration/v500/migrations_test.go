package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

var (
	currentProviderVersion = internal.PackageVersion
)

//go:embed testdata/v4_simple_items.tf
var v4SimpleItemsConfig string

//go:embed testdata/v5_simple_items.tf
var v5SimpleItemsConfig string

//go:embed testdata/v4_empty.tf
var v4EmptyConfig string

//go:embed testdata/v5_empty.tf
var v5EmptyConfig string

// unsetAPIToken is needed because the Access service does not yet support API tokens
func unsetAPIToken(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}
}

// ============================================================================
// Ported from migrations_test.go (service root)
// ============================================================================

func TestMigrateZeroTrustList_FromV4_SimpleItems(t *testing.T) {
	unsetAPIToken(t)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	testConfig := fmt.Sprintf(v4SimpleItemsConfig, rnd, accountID)
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

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
						VersionConstraint: version,
					},
				},
				Config: testConfig,
			},
			acctest.MigrationV2TestStep(t, testConfig, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("tf-acc-test-"+rnd)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("IP")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("items"), knownvalue.SetSizeExact(2)),
			}),
		},
	})
}

func TestMigrateZeroTrustList_FromV4_ItemsWithDescription(t *testing.T) {
	unsetAPIToken(t)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()

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
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

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
						VersionConstraint: version,
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("DOMAIN")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("description"), knownvalue.StringExact("Test list with descriptions")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("items"), knownvalue.SetSizeExact(2)),
			}),
		},
	})
}

func TestMigrateZeroTrustList_FromV4_MixedItems(t *testing.T) {
	unsetAPIToken(t)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()

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
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

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
						VersionConstraint: version,
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("IP")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("items"), knownvalue.SetSizeExact(3)),
			}),
		},
	})
}

func TestMigrateZeroTrustList_FromV4_EmptyList(t *testing.T) {
	unsetAPIToken(t)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	testConfig := fmt.Sprintf(v4EmptyConfig, rnd, accountID)
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

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
						VersionConstraint: version,
					},
				},
				Config: testConfig,
			},
			acctest.MigrationV2TestStep(t, testConfig, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("SERIAL")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("items"), knownvalue.Null()),
			}),
		},
	})
}

func TestMigrateZeroTrustList_FromV4_EmailList(t *testing.T) {
	unsetAPIToken(t)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()

	v4Config := fmt.Sprintf(`
resource "cloudflare_teams_list" "%[1]s" {
  account_id  = "%[2]s"
  name        = "tf-acc-test-%[1]s"
  type        = "EMAIL"
  description = "Email allowlist for testing"
  items       = ["user1@example.com", "user2@example.com", "admin@company.org"]
}`, rnd, accountID)
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

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
						VersionConstraint: version,
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("EMAIL")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("description"), knownvalue.StringExact("Email allowlist for testing")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("items"), knownvalue.SetSizeExact(3)),
			}),
		},
	})
}

func TestMigrateZeroTrustList_FromV4_URLList(t *testing.T) {
	unsetAPIToken(t)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()

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
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

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
						VersionConstraint: version,
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("URL")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("items"), knownvalue.SetSizeExact(3)),
			}),
		},
	})
}

func TestMigrateZeroTrustList_FromV4_LargeList(t *testing.T) {
	unsetAPIToken(t)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()

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
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

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
						VersionConstraint: version,
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("IP")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("items"), knownvalue.SetSizeExact(17)),
			}),
		},
	})
}

func TestMigrateZeroTrustList_FromV4_SpecialCharacters(t *testing.T) {
	unsetAPIToken(t)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()

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
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

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
						VersionConstraint: version,
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("DOMAIN")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_list."+rnd, tfjsonpath.New("items"), knownvalue.SetSizeExact(4)),
			}),
		},
	})
}

// ============================================================================
// New dual test cases (from_v4_latest / from_v5)
// ============================================================================

func TestMigrateZeroTrustList_SimpleItems(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string { return fmt.Sprintf(v4SimpleItemsConfig, rnd, accountID) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID string) string { return fmt.Sprintf(v5SimpleItemsConfig, rnd, accountID) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			unsetAPIToken(t)
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_list." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact("tf-acc-test-"+rnd)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("IP")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.SetSizeExact(2)),
					}),
				},
			})
		})
	}
}

func TestMigrateZeroTrustList_EmptyList(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string { return fmt.Sprintf(v4EmptyConfig, rnd, accountID) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID string) string { return fmt.Sprintf(v5EmptyConfig, rnd, accountID) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			unsetAPIToken(t)
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_list." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("SERIAL")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.Null()),
					}),
				},
			})
		})
	}
}
