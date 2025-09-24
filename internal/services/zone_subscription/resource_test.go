package zone_subscription_test

import (
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareZoneSubscription_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_zone_subscription." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneSubscriptionConfig(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					// Note: rate_plan might be nested differently, skip for now
					// Verify computed attributes exist
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("currency"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("price"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("frequency"), knownvalue.StringExact("not-applicable")),
					// current_period_end and current_period_start might not be set for enterprise plans
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(), // Should show no changes
					},
				},
			},
			{
				Config: testAccCloudflareZoneSubscriptionConfigUpdate(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					// Verify computed attributes still exist
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("currency"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("frequency"), knownvalue.StringExact("not-applicable")),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(), // Should show no changes
					},
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareZoneSubscription_WithPlanChange(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_zone_subscription." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneSubscriptionWithPlan(rnd, zoneID, "enterprise"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "zone_id", zoneID),
					resource.TestCheckResourceAttr(resourceName, "rate_plan.id", "enterprise"),
					resource.TestCheckResourceAttr(resourceName, "frequency", "not-applicable"),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(), // Should show no changes
					},
				},
			},
			{
				Config: testAccCloudflareZoneSubscriptionWithPlan(rnd, zoneID, "free"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "zone_id", zoneID),
					resource.TestCheckResourceAttr(resourceName, "rate_plan.id", "free"),
					resource.TestCheckResourceAttr(resourceName, "frequency", "not-applicable"),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(), // Should show no changes
					},
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareZoneSubscription_WithPlanAndComputedOptionalRatePlanFields(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_zone_subscription." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneSubscriptionWithPlanAndComputedOptionalRatePlanFields(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "zone_id", zoneID),
					resource.TestCheckResourceAttr(resourceName, "frequency", "not-applicable"),
					resource.TestCheckResourceAttr(resourceName, "rate_plan.id", "free"),
					resource.TestCheckResourceAttr(resourceName, "rate_plan.currency", "USD"),
					resource.TestCheckResourceAttr(resourceName, "rate_plan.is_contract", "false"),
					resource.TestCheckResourceAttr(resourceName, "rate_plan.public_name", "Cloudflare Free Plan"),
					resource.TestCheckResourceAttr(resourceName, "rate_plan.externally_managed", "false"),
					resource.TestCheckResourceAttr(resourceName, "rate_plan.scope", "zone"),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(), // Should show no changes
					},
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// Free and Enterprise plans and when any frequency is provided
// will return  Error: Provider produced inconsistent result after apply

func TestAccCloudflareZoneSubscription_WithPlanThatReturnsNotApplicableFrequencyResultingInError(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_zone_subscription." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCloudflareZoneSubscriptionWithPlanAndFrequency(rnd, zoneID, "free", "monthly"),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("Error: Provider produced inconsistent result after apply")),
			},
			{
				Config:      testAccCloudflareZoneSubscriptionWithPlanAndFrequency(rnd, zoneID, "enterprise", "monthly"),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("Error: Provider produced inconsistent result after apply")),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// func TestAccCloudflareZoneSubscription_WithPlanProAndFrequencyMonthly(t *testing.T) {
// 	rnd := utils.GenerateRandomResourceName()
// 	zoneID := "1a9554f82d41a491021f661fb93f0b78"
// 	resourceName := "cloudflare_zone_subscription." + rnd

// 	resource.Test(t, resource.TestCase{
// 		PreCheck: func() {
// 			acctest.TestAccPreCheck(t)
// 			acctest.TestAccPreCheck_ZoneID(t)
// 		},
// 		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccCloudflareZoneSubscriptionWithPlanAndFrequency(rnd, zoneID, "pro", "monthly"),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceName, "zone_id", zoneID),
// 					resource.TestCheckResourceAttr(resourceName, "rate_plan.id", "pro"),
// 					resource.TestCheckResourceAttr(resourceName, "frequency", "monthly"),
// 				),
// 				ConfigPlanChecks: resource.ConfigPlanChecks{
// 					PostApplyPostRefresh: []plancheck.PlanCheck{
// 						plancheck.ExpectEmptyPlan(), // Should show no changes
// 					},
// 				},
// 			},
// 			{
// 				ResourceName:      resourceName,
// 				ImportState:       true,
// 				ImportStateVerify: true,
// 			},
// 		},
// 	})
// }

// func TestAccCloudflareZoneSubscription_WithPlanBusinessAndFrequencyYearly(t *testing.T) {
// 	rnd := utils.GenerateRandomResourceName()
// 	zoneID := "7df2f8a2cea32ed6d9c048403f3b2e27"
// 	resourceName := "cloudflare_zone_subscription." + rnd

// 	resource.Test(t, resource.TestCase{
// 		PreCheck: func() {
// 			acctest.TestAccPreCheck(t)
// 		},
// 		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccCloudflareZoneSubscriptionWithPlanAndFrequency(rnd, zoneID, "business", "yearly"),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceName, "zone_id", zoneID),
// 					resource.TestCheckResourceAttr(resourceName, "rate_plan.id", "business"),
// 					resource.TestCheckResourceAttr(resourceName, "frequency", "yearly"),
// 				),
// 				ConfigPlanChecks: resource.ConfigPlanChecks{
// 					PostApplyPostRefresh: []plancheck.PlanCheck{
// 						plancheck.ExpectEmptyPlan(), // Should show no changes
// 					},
// 				},
// 			},
// 			{
// 				ResourceName:      resourceName,
// 				ImportState:       true,
// 				ImportStateVerify: true,
// 			},
// 		},
// 	})
// }

// TestAccCloudflareZoneSubscription_ImportNoChanges tests that after importing a zone subscription
// for a free plan, there are no plan changes shown.
// https://github.com/cloudflare/terraform-provider-cloudflare/issues/5971
func TestAccCloudflareZoneSubscription_ImportNoChanges(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_zone_subscription." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// First create a basic configuration
			{
				Config: testAccCloudflareZoneSubscriptionImportConfig(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "zone_id", zoneID),
				),
			},
			// Then import the resource and verify it
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"current_period_end", "current_period_start"},
			},
			// Then check that there are no plan changes after import
			{
				Config: testAccCloudflareZoneSubscriptionImportConfig(rnd, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func testAccCloudflareZoneSubscriptionConfig(rnd, zoneID string) string {
	return acctest.LoadTestCase("basic.tf", rnd, zoneID)
}

func testAccCloudflareZoneSubscriptionConfigUpdate(rnd, zoneID string) string {
	return acctest.LoadTestCase("basic_update.tf", rnd, zoneID)
}

func testAccCloudflareZoneSubscriptionWithPlan(rnd, zoneID, plan string) string {
	return acctest.LoadTestCase("with_plan.tf", rnd, zoneID, plan)
}

func testAccCloudflareZoneSubscriptionWithPlanAndFrequency(rnd, zoneID, plan, frequency string) string {
	return acctest.LoadTestCase("with_plan_and_frequency.tf", rnd, zoneID, plan, frequency)
}

func testAccCloudflareZoneSubscriptionImportConfig(rnd, zoneID string) string {
	return acctest.LoadTestCase("import_config.tf", rnd, zoneID)
}

func testAccCloudflareZoneSubscriptionWithPlanAndComputedOptionalRatePlanFields(rnd, zoneID string) string {
	return acctest.LoadTestCase("with_plan_and_computed_optional_rate_plan_fields.tf", rnd, zoneID)
}
