package zone_subscription_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zone_subscription", &resource.Sweeper{
		Name: "cloudflare_zone_subscription",
		F:    testSweepCloudflareZoneSubscription,
	})
}

func testSweepCloudflareZoneSubscription(r string) error {
	ctx := context.Background()
	// Zone Subscription is a zone-level billing subscription setting.
	// It cannot be deleted, only updated.
	// No sweeping required.
	tflog.Info(ctx, "Zone Subscription doesn't require sweeping (zone billing setting)")
	return nil
}

func TestAccCloudflareZoneSubscription_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ALT_ZONE_ID")
	resourceName := "cloudflare_zone_subscription." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		// No CheckDestroy as zone_subscription cannot be deleted
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneSubscriptionConfig(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("currency"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("price"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccCloudflareZoneSubscriptionConfig(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					// Verify computed attributes still exist
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("currency"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("price"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Verify no plan change
			{
				Config: testAccCloudflareZoneSubscriptionConfig(rnd, zoneID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestAccCloudflareZoneSubscriptionResource_WithPlanChange(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ALT_ZONE_ID")
	resourceName := "cloudflare_zone_subscription." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneSubscriptionWithPlan(rnd, zoneID, "enterprise"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "zone_id", zoneID),
					resource.TestCheckResourceAttr(resourceName, "rate_plan.id", "enterprise"),
				),
			},
			{
				Config: testAccCloudflareZoneSubscriptionWithPlan(rnd, zoneID, "free"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "zone_id", zoneID),
					resource.TestCheckResourceAttr(resourceName, "rate_plan.id", "free"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Verify no plan change
			{
				Config: testAccCloudflareZoneSubscriptionWithPlan(rnd, zoneID, "free"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

// Tests that creating a zone with subscription doesn't cause drift on computed fields
// The bug was that computed fields (currency, price, state, rate_plan nested fields, frequency)
// were causing drift on subsequent applies because they were being set to null in config
func TestAccCloudflareZoneSubscriptionResource_CreateZoneWithPlan_CUSTESC_57375(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneName := fmt.Sprintf("%s.net", rnd)
	zoneResourceName := fmt.Sprintf("zone_%s", rnd)
	subscriptionResourceName := "cloudflare_zone_subscription." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneSubscriptionCreateZoneWithPlan(rnd, "enterprise", zoneResourceName, accountID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(subscriptionResourceName, "rate_plan.id", "enterprise"),
					resource.TestCheckResourceAttrSet(subscriptionResourceName, "zone_id"),
					resource.TestCheckResourceAttr(subscriptionResourceName, "frequency", "not-applicable"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(subscriptionResourceName, tfjsonpath.New("frequency"), knownvalue.StringExact("not-applicable")),
				},
			},
			{
				ResourceName:      subscriptionResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// This test is to verify that applying the same configuration does not result in a plan change
			// This is the critical test for CUSTESC-57375 - the bug manifested on the 2nd apply
			{
				Config: testAccCloudflareZoneSubscriptionCreateZoneWithPlan(rnd, "enterprise", zoneResourceName, accountID, zoneName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Third apply to ensure stability (regression test)
			{
				Config: testAccCloudflareZoneSubscriptionCreateZoneWithPlan(rnd, "enterprise", zoneResourceName, accountID, zoneName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

// https://github.com/cloudflare/terraform-provider-cloudflare/issues/5971
// https://github.com/cloudflare/terraform-provider-cloudflare/issues/6485
// Tests that importing a zone subscription with frequency="not-applicable" doesn't cause drift
func TestAccCloudflareZoneSubscriptionResource_ImportNoChanges_BILLSUB_247(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ALT_ZONE_ID")
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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rate_plan").AtMapKey("id"), knownvalue.StringExact("free")),
					// Explicitly check frequency is computed on free plan returns "not-applicable"
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("frequency"), knownvalue.StringExact("not-applicable")),
				},
			},
			// Then import the resource and verify it
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Verify no drift after import
			{
				Config: testAccCloudflareZoneSubscriptionWithPlan(rnd, zoneID, "free"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

// Test that setting frequency on enterprise plans that don't support it causes drift
// This documents the expected behavior for zones that don't support frequency configuration
// The API accepts the value during create/update but returns "not-applicable" on read, causing drift
func TestAccCloudflareZoneSubscriptionResource_FrequencyNotSupported(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ALT_ZONE_ID")
	resourceName := "cloudflare_zone_subscription." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneSubscriptionWithFrequency(rnd, zoneID, "enterprise", "monthly"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rate_plan.id", "enterprise"),
					// After apply, state contains the configured value
					resource.TestCheckResourceAttr(resourceName, "frequency", "monthly"),
				),
				// Expect non-empty plan on refresh because API returns "not-applicable" instead of "monthly"
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// https://github.com/cloudflare/terraform-provider-cloudflare/issues/6374
func TestAccCloudflareZoneSubscriptionResource_PartnersEnt(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneName := fmt.Sprintf("%s.net", rnd)
	zoneResourceName := fmt.Sprintf("zone_%s", rnd)
	subscriptionResourceName := "cloudflare_zone_subscription." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneSubscriptionCreateZoneWithPlan(rnd, "partners_ent", zoneResourceName, accountID, zoneName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(subscriptionResourceName, "rate_plan.id", "partners_ent"),
				),
				// Expect non-empty plan on refresh because API returns "enterprise" instead of "partners_ent"
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCloudflareZoneSubscriptionConfig(rnd, zoneID string) string {
	return acctest.LoadTestCase("basic.tf", rnd, zoneID)
}

func testAccCloudflareZoneSubscriptionWithPlan(rnd, zoneID, plan string) string {
	return acctest.LoadTestCase("with_plan.tf", rnd, zoneID, plan)
}

func testAccCloudflareZoneSubscriptionImportConfig(rnd, zoneID string) string {
	return acctest.LoadTestCase("import_config.tf", rnd, zoneID)
}

func testAccCloudflareZoneSubscriptionCreateZoneWithPlan(rnd, plan, zoneResourceName, accountID, zoneName string) string {
	return acctest.LoadTestCase("create_zone_with_plan.tf", rnd, plan, zoneResourceName, accountID, zoneName)
}

func testAccCloudflareZoneSubscriptionWithFrequency(rnd, zoneID, plan, frequency string) string {
	return acctest.LoadTestCase("with_plan_and_frequency.tf", rnd, zoneID, plan, frequency)
}
