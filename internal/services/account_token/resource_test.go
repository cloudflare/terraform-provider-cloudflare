package account_token_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/accounts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_account_token", &resource.Sweeper{
		Name: "cloudflare_account_token",
		F:    testSweepCloudflareAccountToken,
	})
}

func testSweepCloudflareAccountToken(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	// List all API tokens
	tokens, err := client.Accounts.Tokens.List(ctx, accounts.TokenListParams{})
	if err != nil {
		return fmt.Errorf("failed to fetch account tokens: %w", err)
	}

	// Delete test tokens (those created by our tests)
	for _, token := range tokens.Result {
		if strings.Contains(token.Name, "terraform") {
			_, err := client.Accounts.Tokens.Delete(ctx, token.ID, accounts.TokenDeleteParams{
				AccountID: cloudflare.F(accountID),
			})
			if err != nil {
				return fmt.Errorf("failed to delete account token %s: %w", token.ID, err)
			}
		}
	}

	return nil
}

func testAccCheckCloudflareAccountTokenDestroy(s *terraform.State) error {
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_account_token" {
			continue
		}

		tokenID := rs.Primary.ID
		_, err := client.Accounts.Tokens.Get(context.Background(), tokenID, accounts.TokenGetParams{
			AccountID: cloudflare.F(accountID),
		})
		if err == nil {
			return fmt.Errorf("account token %s still exists", tokenID)
		}
	}

	return nil
}

func TestAccAccountToken_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_account_token.test_account_token"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccountTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("account_token-without-condition.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					// conditions by default should not be set
					resource.TestCheckNoResourceAttr(resourceName, "condition.request_ip.0.in"),
					resource.TestCheckNoResourceAttr(resourceName, "condition.request_ip.0.not_in"),
				),
			},
			{
				Config: acctest.LoadTestCase("account_token-without-condition.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			{
				Config: acctest.LoadTestCase("account_token-without-condition.tf", rnd+"-updated", accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd+"-updated"),
					// Verify conditions still not be set
					resource.TestCheckNoResourceAttr(resourceName, "condition.request_ip.0.in"),
					resource.TestCheckNoResourceAttr(resourceName, "condition.request_ip.0.not_in"),
				),
			},
			{
				Config: acctest.LoadTestCase("account_token-without-condition.tf", rnd+"-updated", accountID),
				// re-plan should not detect drift
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			// Import step
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"value"}, // API token value is not returned by the API
			},
		},
	})
}

func TestAccAccountToken_SetIndividualCondition(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_account_token.test_account_token"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccountTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("account_token-with-individual-condition.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in").AtSliceIndex(0), knownvalue.StringExact("192.0.2.1/32")),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "condition.request_ip.in.0", "192.0.2.1/32"),
					resource.TestCheckNoResourceAttr(resourceName, "condition.request_ip.not_in"),
				),
			},
			{
				Config: acctest.LoadTestCase("account_token-with-individual-condition.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			// Import step
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"value"}, // API token value is not returned by the API
			},
		},
	})
}

func TestAccAccountToken_SetAllCondition(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_account_token.test_account_token"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccountTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("account_token-with-all-condition.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					// Validate "in" condition
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in").AtSliceIndex(0), knownvalue.StringExact("192.0.2.1/32")),
					// Validate "not_in" condition
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("not_in"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("not_in").AtSliceIndex(0), knownvalue.StringExact("198.51.100.1/32")),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "condition.request_ip.in.0", "192.0.2.1/32"),
					resource.TestCheckResourceAttr(resourceName, "condition.request_ip.not_in.0", "198.51.100.1/32"),
				),
			},
			{
				Config: acctest.LoadTestCase("account_token-with-all-condition.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Import step
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"value"}, // API token value is not returned by the API
			},
		},
	})
}

func TestAccAccountToken_TokenTTL(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_account_token.test_account_token"

	oneDaysFromNow := time.Now().UTC().AddDate(0, 0, 1)
	expireTime := oneDaysFromNow.Format(time.RFC3339)
	twoDaysFromNow := time.Now().UTC().AddDate(0, 0, 2)
	updatedExpireTime := twoDaysFromNow.Format(time.RFC3339)
	notBefore := "2018-07-01T05:20:00Z"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccountTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("account_token-with-ttl.tf", rnd, accountID, expireTime),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("not_before"), knownvalue.StringExact(notBefore)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact(expireTime)),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "not_before", notBefore),
					resource.TestCheckResourceAttr(resourceName, "expires_on", expireTime),
				),
			},
			{
				Config: acctest.LoadTestCase("account_token-with-ttl.tf", rnd, accountID, expireTime),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			{
				Config: acctest.LoadTestCase("account_token-with-ttl.tf", rnd, accountID, updatedExpireTime),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact(updatedExpireTime)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact(updatedExpireTime)),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "not_before", notBefore),
					resource.TestCheckResourceAttr(resourceName, "expires_on", updatedExpireTime),
				),
			},
			// Import step
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"value"}, // API token value is not returned by the API
			},
		},
	})
}

func TestAccAccountToken_PermissionGroupOrder(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_account_token.test_account_token"

	permgroup0_SAME := statecheck.CompareValue(compare.ValuesSame())
	permgroup1_SAME := statecheck.CompareValue(compare.ValuesSame())

	// Test that permission group order doesn't affect plans
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccountTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("account_token-permissiongroup-order1.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups"), knownvalue.ListSizeExact(2)),
					permgroup0_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id")),
					permgroup1_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(1).AtMapKey("id")),
				},
			},
			{
				Config: acctest.LoadTestCase("account_token-permissiongroup-order2.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			{
				Config: acctest.LoadTestCase("account_token-permissiongroup-order2.tf", rnd+"-updated", accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups"), knownvalue.ListSizeExact(2)),
					permgroup0_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id")),
					permgroup1_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(1).AtMapKey("id")),
				},
			},
			{
				Config: acctest.LoadTestCase("account_token-permissiongroup-order1.tf", rnd+"-updated", accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			// Import step
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"value"}, // API token value is not returned by the API
			},
		},
	})

	// Test the reverse order scenario
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccountTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("account_token-permissiongroup-order2.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups"), knownvalue.ListSizeExact(2)),
					permgroup0_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id")),
					permgroup1_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(1).AtMapKey("id")),
				},
			},
			{
				Config: acctest.LoadTestCase("account_token-permissiongroup-order1.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			{
				Config: acctest.LoadTestCase("account_token-permissiongroup-order1.tf", rnd+"-updated", accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups"), knownvalue.ListSizeExact(2)),
					permgroup0_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id")),
					permgroup1_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(1).AtMapKey("id")),
				},
			},
			{
				Config: acctest.LoadTestCase("account_token-permissiongroup-order2.tf", rnd+"-updated", accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups"), knownvalue.ListSizeExact(2)),
					permgroup0_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id")),
					permgroup1_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(1).AtMapKey("id")),
				},
			},
			// Import step
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"value"}, // API token value is not returned by the API
			},
		},
	})
}

// Test that policy group order doesn't affect plans. This test uses two
// corresponding .tf files that are identical except they swap the order of two
// policies.
func TestAccAccountToken_PolicyOrder(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_account_token.test_account_token"

	policy1_permgroup_SAME := statecheck.CompareValue(compare.ValuesSame())
	policy2_permgroup_SAME := statecheck.CompareValue(compare.ValuesSame())

	// Test that policy order doesn't affect plans
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccountTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("account_token-policy-order1.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(2)),
					policy1_permgroup_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id")),
					policy2_permgroup_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(1).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id")),
				},
			},
			{
				Config: acctest.LoadTestCase("account_token-policy-order2.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			{
				Config: acctest.LoadTestCase("account_token-policy-order2.tf", rnd+"-updated", accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(2)),
					policy1_permgroup_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id")),
					policy2_permgroup_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(1).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id")),
				},
			},
			{
				Config: acctest.LoadTestCase("account_token-policy-order1.tf", rnd+"-updated", accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			// Import step
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"value"}, // API token value is not returned by the API
			},
		},
	})

	// Test the reverse order scenario
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccountTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("account_token-policy-order2.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(2)),
					policy1_permgroup_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id")),
					policy2_permgroup_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(1).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id")),
				},
			},
			{
				Config: acctest.LoadTestCase("account_token-policy-order1.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			{
				Config: acctest.LoadTestCase("account_token-policy-order1.tf", rnd+"-updated", accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(2)),
					policy1_permgroup_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id")),
					policy2_permgroup_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(1).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id")),
				},
			},
			{
				Config: acctest.LoadTestCase("account_token-policy-order2.tf", rnd+"-updated", accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(2)),
					policy1_permgroup_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id")),
					policy2_permgroup_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(1).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id")),
				},
			},
			// Import step
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"value"}, // API token value is not returned by the API
			},
		},
	})
}

func TestAccAccountToken_ResourcesFlexible(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_account_token.test_account_token"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccountTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("account_token-resources-flexible.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(2)),
				},
			},
			{
				Config: acctest.LoadTestCase("account_token-resources-flexible.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			{
				Config: acctest.LoadTestCase("account_token-resources-flexible.tf", rnd+"-updated", accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(2)),
				},
			},
			{
				Config: acctest.LoadTestCase("account_token-resources-flexible.tf", rnd+"-updated", accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			// Import step
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"value"}, // API token value is not returned by the API
			},
		},
	})
}

func TestAccAccountToken_CRUD(t *testing.T) {
	// Comprehensive test covering Create, Read, Update, Delete + Import
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_account_token.crud_test"

	initialName := rnd + "-initial"
	updatedName := rnd + "-updated"

	// TTL times
	oneDayFromNow := time.Now().UTC().AddDate(0, 0, 1).Format(time.RFC3339)
	twoDaysFromNow := time.Now().UTC().AddDate(0, 0, 2).Format(time.RFC3339)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccountTokenDestroy,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: acctest.LoadTestCase("account_token-crud.tf", initialName, accountID, oneDayFromNow, "192.0.2.1/32"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(initialName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact(oneDayFromNow)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in").AtSliceIndex(0), knownvalue.StringExact("192.0.2.1/32")),
				},
			},
			// Update name and TTL
			{
				Config: acctest.LoadTestCase("account_token-crud.tf", updatedName, accountID, twoDaysFromNow, "192.0.2.1/32"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(updatedName)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact(twoDaysFromNow)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(updatedName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact(twoDaysFromNow)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in").AtSliceIndex(0), knownvalue.StringExact("192.0.2.1/32")),
				},
			},
			// Update condition
			{
				Config: acctest.LoadTestCase("account_token-crud.tf", updatedName, accountID, twoDaysFromNow, "198.51.100.1/32"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in").AtSliceIndex(0), knownvalue.StringExact("198.51.100.1/32")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(updatedName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in").AtSliceIndex(0), knownvalue.StringExact("198.51.100.1/32")),
				},
			},
			// Import
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"value"}, // Token value is write-only
			},
		},
	})
}
