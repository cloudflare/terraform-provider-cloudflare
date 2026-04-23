package custom_pages_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

/* Migration tests don't include every possible permutation, but do cover:
 * - At least one migration from v4 and one from v5 for each identifier
 * - One migration from v4 provider without a "type" attribute
 * - Zone-level and account-level custom pages migrated from both v4 and v5
 * - Migrations of "default" and "customized" pages from v5
 * - Migrations of "customized" pages from v5 (note: couldn't actually create "default" page in v4 provider)
 * - Import verification in a few cases
 */

func TestMigrateCustomPagesMigrationFromV4(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	// V4 config using old "type" attribute
	v4Config := fmt.Sprintf(`
resource "cloudflare_custom_pages" "%[1]s" {
  account_id = "%[2]s"
  type       = "500_errors"
  state      = "customized"
  url        = "https://custom-pages-basic.terraform-provider-acceptance-testing.workers.dev/"
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
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("500_errors")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-basic.terraform-provider-acceptance-testing.workers.dev/")),
			}),
		},
	})
}

// Because we didn't set a schema version when we moved to V5,
// the state upgrade will also run on custom_pages created with
// earlier V5 providers;
// this test ensures the upgrade is a no-op for these resources
func TestMigrateCustomPagesMigrationFromV5(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			// Step 1: Create a custom_pages resource with 5.2.0 version constraint
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "= 5.2.0", // custom_pages didn't exist in 5.2.0 and 5.1.0
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesAccountConfig(rnd, accountID, "500_errors", "customized", "https://custom-pages-basic.terraform-provider-acceptance-testing.workers.dev/"),
			},
			// Step 2: Upgrade to current and make sure the plan is empty
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testAccCustomPagesAccountConfig(rnd, accountID, "500_errors", "customized", "https://custom-pages-basic.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("500_errors")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-basic.terraform-provider-acceptance-testing.workers.dev/")),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ResourceName:             resourceName,
				ImportStateIdPrefix:      fmt.Sprintf("accounts/%s/", accountID),
				ImportState:              true,
				ImportStateVerify:        true,
				// TODO these can't be imported because the response schema isn't defined in OpenAPI
				ImportStateVerifyIgnore: []string{"state", "url"},
			},
		},
	})
}

func TestMigrateCustomPagesMigrationFromV5Default(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			// Step 1: Create a custom_pages resource with 5.2.0 version constraint
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "= 5.2.0", // custom_pages didn't exist in 5.2.0 and 5.1.0
						Source:            "cloudflare/cloudflare",
					},
				},
				// Work around known drift in v5.2.0 provider
				ExpectNonEmptyPlan: true,
				Config:             testAccCustomPagesAccountConfig(rnd, accountID, "500_errors", "default", ""),
			},
			// Step 2: Upgrade to current and make sure the plan is empty
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testAccCustomPagesAccountConfig(rnd, accountID, "500_errors", "default", ""),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("500_errors")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("")),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ResourceName:             resourceName,
				ImportStateIdPrefix:      fmt.Sprintf("accounts/%s/", accountID),
				ImportState:              true,
				ImportStateVerify:        true,
				// TODO these can't be imported because the response schema isn't defined in OpenAPI
				ImportStateVerifyIgnore: []string{"state", "url"},
			},
		},
	})
}

// Account-level basic_challenge tests

func TestMigrateCustomPagesMigrationFromV4BasicChallengeCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	// V4 config using old "type" attribute
	v4Config := fmt.Sprintf(`
resource "cloudflare_custom_pages" "%[1]s" {
  account_id = "%[2]s"
  type       = "basic_challenge"
  state      = "customized"
  url        = "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"
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
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("basic_challenge")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
			}),
		},
	})
}

func TestMigrateCustomPagesMigrationFromV5BasicChallengeCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "= 5.2.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesAccountConfig(rnd, accountID, "basic_challenge", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testAccCustomPagesAccountConfig(rnd, accountID, "basic_challenge", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("basic_challenge")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestMigrateCustomPagesMigrationFromV5BasicChallengeDefault(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "= 5.2.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				// Work around known drift in v5.2.0 provider
				ExpectNonEmptyPlan: true,
				Config:             testAccCustomPagesAccountConfig(rnd, accountID, "basic_challenge", "default", ""),
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testAccCustomPagesAccountConfig(rnd, accountID, "basic_challenge", "default", ""),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("basic_challenge")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("")),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

// Account-level country_challenge tests
func TestMigrateCustomPagesMigrationFromV4AccountCountryChallengeCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	// V4 config using old "type" attribute
	v4Config := fmt.Sprintf(`
resource "cloudflare_custom_pages" "%[1]s" {
  account_id = "%[2]s"
  type       = "country_challenge"
  state      = "customized"
  url        = "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"
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
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("country_challenge")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
			}),
		},
	})
}

func TestMigrateCustomPagesMigrationFromV5AccountCountryChallengeCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "= 5.2.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesAccountConfig(rnd, accountID, "country_challenge", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testAccCustomPagesAccountConfig(rnd, accountID, "country_challenge", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("country_challenge")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

// Account-level waf_block tests

func TestMigrateCustomPagesMigrationFromV4WafBlockCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	// V4 config using old "type" attribute (account-level)
	v4Config := fmt.Sprintf(`
resource "cloudflare_custom_pages" "%[1]s" {
  account_id = "%[2]s"
  type       = "waf_block"
  state      = "customized"
  url        = "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"
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
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("waf_block")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
			}),
		},
	})
}

func TestMigrateCustomPagesMigrationFromV5WafBlockCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "= 5.2.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesAccountConfig(rnd, accountID, "waf_block", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testAccCustomPagesAccountConfig(rnd, accountID, "waf_block", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("waf_block")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

// Zone-level ip_block tests

func TestMigrateCustomPagesMigrationFromV4IpBlockCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()

	// V4 config using old "type" attribute (zone-level)
	v4Config := fmt.Sprintf(`
resource "cloudflare_custom_pages" "%[1]s" {
  zone_id = "%[2]s"
  type    = "ip_block"
  state   = "customized"
  url     = "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
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
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("ip_block")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
			}),
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ResourceName:             resourceName,
				ImportStateIdPrefix:      fmt.Sprintf("zones/%s/", zoneID),
				ImportState:              true,
				ImportStateVerify:        true,
				ImportStateVerifyIgnore:  []string{"state", "url"},
			},
		},
	})
}

func TestMigrateCustomPagesMigrationFromV5IpBlockCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "= 5.2.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "ip_block", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testAccCustomPagesZoneConfig(rnd, zoneID, "ip_block", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("ip_block")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ResourceName:             resourceName,
				ImportStateIdPrefix:      fmt.Sprintf("zones/%s/", zoneID),
				ImportState:              true,
				ImportStateVerify:        true,
				ImportStateVerifyIgnore:  []string{"state", "url"},
			},
		},
	})
}

// Zone-level country_challenge tests

func TestMigrateCustomPagesMigrationFromV4CountryChallengeCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()

	// V4 config using old "type" attribute (zone-level)
	v4Config := fmt.Sprintf(`
resource "cloudflare_custom_pages" "%[1]s" {
  zone_id = "%[2]s"
  type    = "country_challenge"
  state   = "customized"
  url     = "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
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
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("country_challenge")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
			}),
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ResourceName:             resourceName,
				ImportStateIdPrefix:      fmt.Sprintf("zones/%s/", zoneID),
				ImportState:              true,
				ImportStateVerify:        true,
				ImportStateVerifyIgnore:  []string{"state", "url"},
			},
		},
	})
}

func TestMigrateCustomPagesMigrationFromV5CountryChallengeCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "= 5.2.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "country_challenge", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testAccCustomPagesZoneConfig(rnd, zoneID, "country_challenge", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("country_challenge")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestMigrateCustomPagesMigrationFromV5CountryChallengeDefault(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "= 5.2.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				// Work around known drift in v5.2.0 provider
				ExpectNonEmptyPlan: true,
				Config:             testAccCustomPagesZoneConfig(rnd, zoneID, "country_challenge", "default", ""),
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testAccCustomPagesZoneConfig(rnd, zoneID, "country_challenge", "default", ""),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("country_challenge")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("")),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

// Zone-level 1000_errors tests

func TestMigrateCustomPagesMigrationFromV41000ErrorsCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()

	// V4 config using old "type" attribute (zone-level)
	v4Config := fmt.Sprintf(`
resource "cloudflare_custom_pages" "%[1]s" {
  zone_id = "%[2]s"
  type    = "1000_errors"
  state   = "customized"
  url     = "https://custom-pages-1000-errors.terraform-provider-acceptance-testing.workers.dev/"
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
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
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("1000_errors")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-1000-errors.terraform-provider-acceptance-testing.workers.dev/")),
			}),
		},
	})
}

func TestMigrateCustomPagesMigrationFromV51000ErrorsCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "= 5.2.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "1000_errors", "customized", "https://custom-pages-1000-errors.terraform-provider-acceptance-testing.workers.dev/"),
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testAccCustomPagesZoneConfig(rnd, zoneID, "1000_errors", "customized", "https://custom-pages-1000-errors.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("1000_errors")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-1000-errors.terraform-provider-acceptance-testing.workers.dev/")),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestMigrateCustomPagesMigrationFromV51000ErrorsDefault(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "= 5.2.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				// Work around known drift in v5.2.0 provider
				ExpectNonEmptyPlan: true,
				Config:             testAccCustomPagesZoneConfig(rnd, zoneID, "1000_errors", "default", ""),
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testAccCustomPagesZoneConfig(rnd, zoneID, "1000_errors", "default", ""),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("1000_errors")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("")),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}
func TestMigrateCustomPagesMigrationFromV4ManagedChallengeCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()

	// V4 config using old "type" attribute (zone-level)
	v4Config := fmt.Sprintf(`
resource "cloudflare_custom_pages" "%[1]s" {
  zone_id = "%[2]s"
  type    = "managed_challenge"
  state   = "customized"
  url     = "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
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
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("managed_challenge")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
			}),
		},
	})
}

func TestMigrateCustomPagesMigrationFromV57ManagedChallengeCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "= 5.7.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "managed_challenge", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testAccCustomPagesZoneConfig(rnd, zoneID, "managed_challenge", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("managed_challenge")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ResourceName:             resourceName,
				ImportStateIdPrefix:      fmt.Sprintf("zones/%s/", zoneID),
				ImportState:              true,
				ImportStateVerify:        true,
				// TODO these can't be imported because the response schema isn't defined in OpenAPI
				ImportStateVerifyIgnore: []string{"state", "url"},
			},
		},
	})
}

func TestMigrateCustomPagesMigrationFromV4RatelimitBlockCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()

	// V4 config using old "type" attribute (zone-level)
	v4Config := fmt.Sprintf(`
resource "cloudflare_custom_pages" "%[1]s" {
  zone_id = "%[2]s"
  type    = "ratelimit_block"
  state   = "customized"
  url     = "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
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
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("ratelimit_block")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
			}),
		},
	})
}

func TestMigrateCustomPagesMigrationFromV5RatelimitBlockCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "= 5.2.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "ratelimit_block", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testAccCustomPagesZoneConfig(rnd, zoneID, "ratelimit_block", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("ratelimit_block")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestMigrateCustomPagesMigrationFromV4UnderAttackCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()

	// V4 config using old "type" attribute (zone-level)
	v4Config := fmt.Sprintf(`
resource "cloudflare_custom_pages" "%[1]s" {
  zone_id = "%[2]s"
  type    = "under_attack"
  state   = "customized"
  url     = "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
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
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("under_attack")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
			}),
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testAccCustomPagesZoneConfig(rnd, zoneID, "under_attack", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("under_attack")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestMigrateCustomPagesMigrationFromV5UnderAttackCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "= 5.2.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "under_attack", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testAccCustomPagesZoneConfig(rnd, zoneID, "under_attack", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("under_attack")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

// TestMigrateCustomPagesMigrationFromV4WafChallenge tests the critical missing identifier: waf_challenge
func TestMigrateCustomPagesMigrationFromV4WafChallenge(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	// V4 config using "type" for waf_challenge
	v4Config := fmt.Sprintf(`
resource "cloudflare_custom_pages" "%[1]s" {
  account_id = "%[2]s"
  type       = "waf_challenge"
  state      = "customized"
  url        = "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"
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
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("waf_challenge")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
			}),
		},
	})
}

// TestMigrateCustomPagesMigrationFromV5WafChallengeZone tests waf_challenge at zone level
func TestMigrateCustomPagesMigrationFromV5WafChallengeZone(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "= 5.2.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "waf_challenge", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testAccCustomPagesZoneConfig(rnd, zoneID, "waf_challenge", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("waf_challenge")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

// TestMigrateCustomPagesMigrationFromV4IpBlockAccount tests ip_block at account level (was only zone-level)
func TestMigrateCustomPagesMigrationFromV4IpBlockAccount(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_custom_pages" "%[1]s" {
  account_id = "%[2]s"
  type       = "ip_block"
  state      = "customized"
  url        = "https://custom-pages-ip-block.terraform-provider-acceptance-testing.workers.dev/"
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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("ip_block")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-ip-block.terraform-provider-acceptance-testing.workers.dev/")),
			}),
		},
	})
}

// TestMigrateCustomPagesMigrationFromV5BasicChallengeZone tests basic_challenge at zone level (was only account-level)
func TestMigrateCustomPagesMigrationFromV5BasicChallengeZone(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "= 5.2.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "basic_challenge", "customized", "https://custom-pages-basic-challenge.terraform-provider-acceptance-testing.workers.dev/"),
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testAccCustomPagesZoneConfig(rnd, zoneID, "basic_challenge", "customized", "https://custom-pages-basic-challenge.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("basic_challenge")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-basic-challenge.terraform-provider-acceptance-testing.workers.dev/")),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

// TestMigrateCustomPagesMigrationFromV4ManagedChallengeAccount tests managed_challenge at account level (was only zone-level)
func TestMigrateCustomPagesMigrationFromV4ManagedChallengeAccount(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_pages." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_custom_pages" "%[1]s" {
  account_id = "%[2]s"
  type       = "managed_challenge"
  state      = "customized"
  url        = "https://custom-pages-managed-challenge.terraform-provider-acceptance-testing.workers.dev/"
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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("managed_challenge")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-managed-challenge.terraform-provider-acceptance-testing.workers.dev/")),
			}),
		},
	})
}

func testAccCustomPagesOldConfig(rnd, accountID, identifier, state, url string) string {
	return acctest.LoadTestCase("migrations_old.tf", rnd, accountID, identifier, state, url)
}

func testAccCustomPagesOldZoneConfig(rnd, zoneID, identifier, state, url string) string {
	return acctest.LoadTestCase("migrations_old_zone.tf", rnd, zoneID, identifier, state, url)
}
