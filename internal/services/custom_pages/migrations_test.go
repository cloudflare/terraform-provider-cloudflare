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

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			// Step 1: Create a custom_pages resource with v4 version constraint
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: " ~> 4.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesOldConfig(rnd, accountID, "500_errors", "customized", "https://custom-pages-basic.terraform-provider-acceptance-testing.workers.dev/"),
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

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: " ~> 4.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesOldConfig(rnd, accountID, "basic_challenge", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
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

func TestMigrateCustomPagesMigrationFromV57BasicChallengeDefault(t *testing.T) {
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
						VersionConstraint: "= 5.7.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesAccountConfig(rnd, accountID, "basic_challenge", "default", ""),
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

// Account-level waf_challenge tests

func TestMigrateCustomPagesMigrationFromV4WafChallengeCustomized(t *testing.T) {
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
						VersionConstraint: " ~> 4.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesOldConfig(rnd, accountID, "waf_challenge", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testAccCustomPagesAccountConfig(rnd, accountID, "waf_challenge", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("waf_challenge")),
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

func TestMigrateCustomPagesMigrationFromV5WafChallengeCustomized(t *testing.T) {
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
				Config: testAccCustomPagesAccountConfig(rnd, accountID, "waf_challenge", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testAccCustomPagesAccountConfig(rnd, accountID, "waf_challenge", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("waf_challenge")),
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

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: " ~> 4.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesOldConfig(rnd, accountID, "waf_block", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
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

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: " ~> 4.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesOldZoneConfig(rnd, zoneID, "ip_block", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
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

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: " ~> 4.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesOldZoneConfig(rnd, zoneID, "country_challenge", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
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

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: " ~> 4.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesOldZoneConfig(rnd, zoneID, "1000_errors", "customized", "https://custom-pages-1000-errors.terraform-provider-acceptance-testing.workers.dev/"),
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

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "~> 4.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesOldZoneConfig(rnd, zoneID, "managed_challenge", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
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
func TestMigrateCustomPagesMigrationFromV5ManagedChallengeCustomized(t *testing.T) {
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
		},
	})
}

func TestMigrateCustomPagesMigrationFromV4RatelimitBlockCustomized(t *testing.T) {
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
						VersionConstraint: "~> 4.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesOldZoneConfig(rnd, zoneID, "ratelimit_block", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
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

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "~> 4.0",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testAccCustomPagesOldZoneConfig(rnd, zoneID, "under_attack", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
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

func testAccCustomPagesOldConfig(rnd, accountID, identifier, state, url string) string {
	return acctest.LoadTestCase("migrations_old.tf", rnd, accountID, identifier, state, url)
}

func testAccCustomPagesOldZoneConfig(rnd, zoneID, identifier, state, url string) string {
	return acctest.LoadTestCase("migrations_old_zone.tf", rnd, zoneID, identifier, state, url)
}
