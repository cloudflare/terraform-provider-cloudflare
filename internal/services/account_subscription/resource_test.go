package account_subscription_test

// NOTE: No sweeper is needed for zone_subscription as the resource cannot be deleted

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// https://github.com/cloudflare/terraform-provider-cloudflare/issues/5803
func TestAccCloudflareAccountSubscription_ImportNoChanges(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_account_subscription." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// First create a basic configuration
			{
				Config: testAccCloudflareAccountSubscriptionImportConfig(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(), // Should show no changes
					},
				},
			},
			// Then import the resource and verify it
			{
				ResourceName: resourceName,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					r, ok := s.RootModule().Resources[resourceName]
					if !ok {
						return "", fmt.Errorf("not found: %s", resourceName)
					}

					if r.Primary.ID == "" {
						return "", fmt.Errorf("no subscription ID is set")
					}
					subscriptionID := r.Primary.ID
					t.Logf("Importing with ID in import state id func: %s/%s", accountID, subscriptionID)

					return fmt.Sprintf("%s/%s", accountID, subscriptionID), nil
				},
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"current_period_end",
					"current_period_start",
				},
			},
		},
	})
}

func TestAccCloudflareAccountSubscription_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_account_subscription." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// First create a basic configuration
			{
				Config: testAccCloudflareAccountSubscriptionBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "rate_plan.id", "teams_free"),
				),
			},
		},
	})
}
func TestAccCloudflareAccountSubscription_BasicWithOptionalFields(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_account_subscription." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// First create a basic configuration
			{
				Config: testAccCloudflareAccountSubscriptionBasicWithOptionalFields(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "rate_plan.id", "teams_free"),
				),
			},
		},
	})
}

func testAccCloudflareAccountSubscriptionImportConfig(rnd, zoneID string) string {
	return acctest.LoadTestCase("import_config.tf", rnd, zoneID)
}

func testAccCloudflareAccountSubscriptionBasic(rnd, zoneID string) string {
	return acctest.LoadTestCase("basic.tf", rnd, zoneID)
}
func testAccCloudflareAccountSubscriptionBasicWithOptionalFields(rnd, zoneID string) string {
	return acctest.LoadTestCase("basic_with_optional_fields.tf", rnd, zoneID)
}
