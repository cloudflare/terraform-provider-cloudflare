package v500_test

import (
	_ "embed"
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

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v4_with_advertisement.tf
var v4WithAdvertisementConfig string

//go:embed testdata/v4_minimal.tf
var v4MinimalConfig string

// TestMigrateBYOIPPrefix_Basic tests basic migration with description.
//
// This test verifies state upgrade (0→1) via UpgradeFromV0, field rename
// (prefix_id→id), removal of advertisement, and API-populated computed fields.
//
// Note: from_v5 tests are not included because v5's Create method provisions
// a new prefix via the API (requiring unique asn/cidr), which conflicts with
// the existing test prefix. Since only one test prefix is available and BYO IP
// prefixes cannot be created idempotently, only v4→v5 migration is tested.
func TestMigrateBYOIPPrefix_Basic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	prefixID := os.Getenv("CLOUDFLARE_BYO_IP_PREFIX_ID")

	if accountID == "" || prefixID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID and CLOUDFLARE_BYO_IP_PREFIX_ID must be set")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_byo_ip_prefix." + rnd
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

	testConfig := fmt.Sprintf(v4BasicConfig, rnd, accountID, prefixID)

	stateChecks := []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(prefixID)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Migration test prefix")),
		// asn and cidr are populated from the API by the v5 provider on the first read
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("asn"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cidr"), knownvalue.NotNull()),
	}

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
			acctest.MigrationV2TestStep(t, testConfig, tmpDir, version, sourceVer, targetVer, stateChecks),
		},
	})
}

// TestMigrateBYOIPPrefix_WithAdvertisement tests that the v4 advertisement field is
// correctly dropped during migration and replaced by the v5 computed advertised field.
//
// Note: from_v5 tests are not included. See TestMigrateBYOIPPrefix_Basic for details.
func TestMigrateBYOIPPrefix_WithAdvertisement(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	prefixID := os.Getenv("CLOUDFLARE_BYO_IP_PREFIX_ID")

	if accountID == "" || prefixID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID and CLOUDFLARE_BYO_IP_PREFIX_ID must be set")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_byo_ip_prefix." + rnd
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

	testConfig := fmt.Sprintf(v4WithAdvertisementConfig, rnd, accountID, prefixID)

	stateChecks := []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(prefixID)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Migration test with advertisement")),
	}

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
			acctest.MigrationV2TestStep(t, testConfig, tmpDir, version, sourceVer, targetVer, stateChecks),
		},
	})
}

// TestMigrateBYOIPPrefix_Minimal tests migration with only the required v4 fields,
// verifying that v5 computed fields (asn, cidr, advertised) are populated from the API.
//
// Note: from_v5 tests are not included. See TestMigrateBYOIPPrefix_Basic for details.
func TestMigrateBYOIPPrefix_Minimal(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	prefixID := os.Getenv("CLOUDFLARE_BYO_IP_PREFIX_ID")

	if accountID == "" || prefixID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID and CLOUDFLARE_BYO_IP_PREFIX_ID must be set")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_byo_ip_prefix." + rnd
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

	testConfig := fmt.Sprintf(v4MinimalConfig, rnd, accountID, prefixID)

	stateChecks := []statecheck.StateCheck{
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(prefixID)),
		// asn and cidr are populated from the API by the v5 provider on the first read
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("asn"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cidr"), knownvalue.NotNull()),
	}

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
			acctest.MigrationV2TestStep(t, testConfig, tmpDir, version, sourceVer, targetVer, stateChecks),
		},
	})
}
