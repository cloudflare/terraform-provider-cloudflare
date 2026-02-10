package byo_ip_prefix_test

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

// Config generators for v4 provider

func byoIPPrefixConfigV4Basic(rnd, accountID, prefixID string) string {
	return fmt.Sprintf(`
resource "cloudflare_byo_ip_prefix" "%[1]s" {
  account_id = "%[2]s"
  prefix_id  = "%[3]s"
  description = "Migration test prefix"
}`, rnd, accountID, prefixID)
}

func byoIPPrefixConfigV4WithAdvertisement(rnd, accountID, prefixID string) string {
	return fmt.Sprintf(`
resource "cloudflare_byo_ip_prefix" "%[1]s" {
  account_id    = "%[2]s"
  prefix_id     = "%[3]s"
  description   = "Migration test with advertisement"
  advertisement = "on"
}`, rnd, accountID, prefixID)
}

func byoIPPrefixConfigV4Minimal(rnd, accountID, prefixID string) string {
	return fmt.Sprintf(`
resource "cloudflare_byo_ip_prefix" "%[1]s" {
  account_id = "%[2]s"
  prefix_id  = "%[3]s"
}`, rnd, accountID, prefixID)
}

// TestMigrateBYOIPPrefix_V4ToV5_Basic tests basic migration with description
func TestMigrateBYOIPPrefix_V4ToV5_Basic(t *testing.T) {
	// BYO IP prefix tests require environment variables
	// These should point to a real prefix that already exists
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	prefixID := os.Getenv("CLOUDFLARE_BYO_IP_PREFIX_ID")

	if accountID == "" || prefixID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID and CLOUDFLARE_BYO_IP_PREFIX_ID must be set for BYO IP prefix migration tests")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_byo_ip_prefix." + rnd
	tmpDir := t.TempDir()

	v4Config := byoIPPrefixConfigV4Basic(rnd, accountID, prefixID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Import resource with v4 provider
				// Note: BYO IP prefixes cannot be created, they must be imported
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
				// Import the existing prefix
				ResourceName:       resourceName,
				ImportState:        true,
				ImportStateId:      fmt.Sprintf("%s/%s", accountID, prefixID),
				ImportStateVerify:  true,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Verify resource migrated correctly
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(prefixID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Migration test prefix")),
				// Verify prefix_id field was removed (it becomes 'id' in v5)
				// Verify advertisement field was removed (computed in v5)
				// Verify computed fields exist in v5 (populated by API)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("advertised"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("asn"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cidr"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateBYOIPPrefix_V4ToV5_WithAdvertisement tests migration with advertisement field
func TestMigrateBYOIPPrefix_V4ToV5_WithAdvertisement(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	prefixID := os.Getenv("CLOUDFLARE_BYO_IP_PREFIX_ID")

	if accountID == "" || prefixID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID and CLOUDFLARE_BYO_IP_PREFIX_ID must be set for BYO IP prefix migration tests")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_byo_ip_prefix." + rnd
	tmpDir := t.TempDir()

	v4Config := byoIPPrefixConfigV4WithAdvertisement(rnd, accountID, prefixID)

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
						VersionConstraint: "4.52.1",
					},
				},
				Config:             v4Config,
				ResourceName:       resourceName,
				ImportState:        true,
				ImportStateId:      fmt.Sprintf("%s/%s", accountID, prefixID),
				ImportStateVerify:  true,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(prefixID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Migration test with advertisement")),
				// Verify v4 'advertisement' field was removed and v5 'advertised' exists
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("advertised"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateBYOIPPrefix_V4ToV5_Minimal tests migration with only required fields
func TestMigrateBYOIPPrefix_V4ToV5_Minimal(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	prefixID := os.Getenv("CLOUDFLARE_BYO_IP_PREFIX_ID")

	if accountID == "" || prefixID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID and CLOUDFLARE_BYO_IP_PREFIX_ID must be set for BYO IP prefix migration tests")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_byo_ip_prefix." + rnd
	tmpDir := t.TempDir()

	v4Config := byoIPPrefixConfigV4Minimal(rnd, accountID, prefixID)

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
						VersionConstraint: "4.52.1",
					},
				},
				Config:             v4Config,
				ResourceName:       resourceName,
				ImportState:        true,
				ImportStateId:      fmt.Sprintf("%s/%s", accountID, prefixID),
				ImportStateVerify:  true,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Verify only required fields present after migration
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(prefixID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				// Verify v5 computed fields populated from API
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("asn"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cidr"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("advertised"), knownvalue.NotNull()),
			}),
		},
	})
}
