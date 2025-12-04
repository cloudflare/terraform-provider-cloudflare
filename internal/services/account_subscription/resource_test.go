package account_subscription_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

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
				Config: testAccCloudflareAccountSubscriptionWithPlan(rnd, accountID, "teams_free"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "rate_plan.id", "teams_free"),
				),
			},
			{
				ResourceName: resourceName,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return importStateIdFuncHelper(s, resourceName, accountID)
				},
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Verify no plan change
			{
				Config: testAccCloudflareAccountSubscriptionWithPlan(rnd, accountID, "teams_free"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

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
					resource.TestCheckResourceAttr(resourceName, "rate_plan.id", "teams_free"),
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
					return importStateIdFuncHelper(s, resourceName, accountID)
				},
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// Test that import fails with invalid format
func TestAccCloudflareAccountSubscription_ImportInvalidFormat(t *testing.T) {
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
			{
				Config: testAccCloudflareAccountSubscriptionWithPlan(rnd, accountID, "teams_free"),
			},
			{
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: accountID, // Wrong format - missing subscription_id
				ExpectError:   regexp.MustCompile("invalid ID"),
			},
		},
	})
}

// Test that import fails with non-existent subscription
func TestAccCloudflareAccountSubscription_ImportNonExistent(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_account_subscription." + rnd
	fakeSubscriptionID := "00000000000000000000000000000000"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:        testAccCloudflareAccountSubscriptionWithPlan(rnd, accountID, "teams_free"),
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s/%s", accountID, fakeSubscriptionID),
				ExpectError:   regexp.MustCompile("Subscription not found"),
			},
		},
	})
}

// Test computed fields are populated correctly after import
func TestAccCloudflareAccountSubscription_ComputedFieldsAfterImport(t *testing.T) {
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
			{
				Config: testAccCloudflareAccountSubscriptionWithPlan(rnd, accountID, "teams_free"),
			},
			{
				ResourceName: resourceName,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return importStateIdFuncHelper(s, resourceName, accountID)
				},
				ImportState: true,
				Check: resource.ComposeTestCheckFunc(
					// Verify computed fields are populated
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "currency"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "price"),
					resource.TestCheckResourceAttrSet(resourceName, "rate_plan.currency"),
					resource.TestCheckResourceAttrSet(resourceName, "rate_plan.public_name"),
					resource.TestCheckResourceAttrSet(resourceName, "rate_plan.externally_managed"),
					resource.TestCheckResourceAttrSet(resourceName, "rate_plan.is_contract"),
				),
			},
		},
	})
}

func testAccCloudflareAccountSubscriptionImportConfig(rnd, zoneID string) string {
	return acctest.LoadTestCase("import_config.tf", rnd, zoneID)
}

func testAccCloudflareAccountSubscriptionWithPlan(rnd, zoneID string, planID string) string {
	return acctest.LoadTestCase("with_plan.tf", rnd, zoneID, planID)
}

// func testAccCloudflareAccountSubscriptionBasicWithOptionalFields(rnd, zoneID string) string {
// 	return acctest.LoadTestCase("basic_with_optional_fields.tf", rnd, zoneID)
// }

func importStateIdFuncHelper(s *terraform.State, resourceName string, accountID string) (string, error) {
	r, ok := s.RootModule().Resources[resourceName]
	if !ok {
		return "", fmt.Errorf("not found: %s", resourceName)
	}

	if r.Primary.ID == "" {
		return "", fmt.Errorf("no subscription ID is set")
	}
	subscriptionID := r.Primary.ID

	return fmt.Sprintf("%s/%s", accountID, subscriptionID), nil
}
