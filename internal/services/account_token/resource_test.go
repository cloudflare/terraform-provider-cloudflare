package account_token_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/accounts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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

	if accountID == "" {
		tflog.Info(ctx, "Skipping account tokens sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	// List all API tokens
	tokens, err := client.Accounts.Tokens.List(ctx, accounts.TokenListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch account tokens: %s", err))
		return fmt.Errorf("failed to fetch account tokens: %w", err)
	}

	if len(tokens.Result) == 0 {
		tflog.Info(ctx, "No account tokens to sweep")
		return nil
	}

	// Delete test tokens (those created by our tests)
	// Uses utils.ShouldSweepResource() to filter by standard test naming convention
	for _, token := range tokens.Result {
		if !utils.ShouldSweepResource(token.Name) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting account token: %s (%s) (account: %s)", token.Name, token.ID, accountID))
		_, err := client.Accounts.Tokens.Delete(ctx, token.ID, accounts.TokenDeleteParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete account token %s (%s): %s", token.Name, token.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted account token: %s (%s)", token.Name, token.ID))
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

// checkPermissionGroupIDsUnchanged verifies that the set of permission group
// IDs at policies[policyIdx] matches expectedIDs regardless of order.
// This replaces cross-step positional stability checks that assumed Set behavior.
func checkPermissionGroupIDsUnchanged(resourceName string, policyIdx int, expectedIDs *[]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %s not found in state", resourceName)
		}

		// Collect all permission_group IDs at this policy index
		var actualIDs []string
		for i := 0; ; i++ {
			key := fmt.Sprintf("policies.%d.permission_groups.%d.id", policyIdx, i)
			val, ok := rs.Primary.Attributes[key]
			if !ok {
				break
			}
			actualIDs = append(actualIDs, val)
		}

		// On first call, capture the IDs as the baseline
		if len(*expectedIDs) == 0 {
			*expectedIDs = make([]string, len(actualIDs))
			copy(*expectedIDs, actualIDs)
			return nil
		}

		// Compare as sets: same IDs regardless of order
		if len(actualIDs) != len(*expectedIDs) {
			return fmt.Errorf("permission_groups count changed: expected %d, got %d", len(*expectedIDs), len(actualIDs))
		}
		expected := make(map[string]bool, len(*expectedIDs))
		for _, id := range *expectedIDs {
			expected[id] = true
		}
		for _, id := range actualIDs {
			if !expected[id] {
				return fmt.Errorf("unexpected permission_group ID %s; expected one of %v", id, *expectedIDs)
			}
		}
		return nil
	}
}

func TestAccAccountToken_PermissionGroupOrder(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_account_token.test_account_token"

	// Track the set of permission group IDs across steps (order-insensitive)
	var permGroupIDs []string

	// Test that permission group order swap produces a benign update
	// but the same set of IDs is preserved
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
				},
				Check: checkPermissionGroupIDsUnchanged(resourceName, 0, &permGroupIDs),
			},
			{
				Config: acctest.LoadTestCase("account_token-permissiongroup-order2.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				Check: checkPermissionGroupIDsUnchanged(resourceName, 0, &permGroupIDs),
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
				},
				Check: checkPermissionGroupIDsUnchanged(resourceName, 0, &permGroupIDs),
			},
			{
				Config: acctest.LoadTestCase("account_token-permissiongroup-order1.tf", rnd+"-updated", accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				Check: checkPermissionGroupIDsUnchanged(resourceName, 0, &permGroupIDs),
			},
			// Import step — ignore policies ordering since import uses canonical
			// sort (no prior state) which may differ from the last config order
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"value", "policies"},
			},
		},
	})

	// Test the reverse order scenario
	var permGroupIDsReverse []string

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
				},
				Check: checkPermissionGroupIDsUnchanged(resourceName, 0, &permGroupIDsReverse),
			},
			{
				Config: acctest.LoadTestCase("account_token-permissiongroup-order1.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				Check: checkPermissionGroupIDsUnchanged(resourceName, 0, &permGroupIDsReverse),
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
				},
				Check: checkPermissionGroupIDsUnchanged(resourceName, 0, &permGroupIDsReverse),
			},
			{
				Config: acctest.LoadTestCase("account_token-permissiongroup-order2.tf", rnd+"-updated", accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups"), knownvalue.ListSizeExact(2)),
				},
				Check: checkPermissionGroupIDsUnchanged(resourceName, 0, &permGroupIDsReverse),
			},
			// Import step — ignore policies ordering since import uses canonical
			// sort (no prior state) which may differ from the last config order
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"value", "policies"},
			},
		},
	})
}

// checkAllPolicyPermGroupIDsUnchanged verifies that the combined set of
// permission group IDs across all policies is unchanged (order-insensitive).
func checkAllPolicyPermGroupIDsUnchanged(resourceName string, expectedIDs *[]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %s not found in state", resourceName)
		}

		// Collect all permission group IDs across all policies
		var actualIDs []string
		for pi := 0; ; pi++ {
			found := false
			for pgi := 0; ; pgi++ {
				key := fmt.Sprintf("policies.%d.permission_groups.%d.id", pi, pgi)
				val, ok := rs.Primary.Attributes[key]
				if !ok {
					break
				}
				found = true
				actualIDs = append(actualIDs, val)
			}
			if !found {
				break
			}
		}

		// On first call, capture the IDs as the baseline
		if len(*expectedIDs) == 0 {
			*expectedIDs = make([]string, len(actualIDs))
			copy(*expectedIDs, actualIDs)
			return nil
		}

		// Compare as sets: same IDs regardless of order or which policy they're in
		if len(actualIDs) != len(*expectedIDs) {
			return fmt.Errorf("total permission_group count changed: expected %d, got %d", len(*expectedIDs), len(actualIDs))
		}
		expected := make(map[string]int, len(*expectedIDs))
		for _, id := range *expectedIDs {
			expected[id]++
		}
		for _, id := range actualIDs {
			if expected[id] <= 0 {
				return fmt.Errorf("unexpected permission_group ID %s; expected IDs: %v", id, *expectedIDs)
			}
			expected[id]--
		}
		return nil
	}
}

// Test that policy group order doesn't affect plans. This test uses two
// corresponding .tf files that are identical except they swap the order of two
// policies.
func TestAccAccountToken_PolicyOrder(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_account_token.test_account_token"

	// Track the combined set of permission group IDs across all policies
	var allPermGroupIDs []string

	// Test that policy order swap produces a benign update
	// but the same set of permission group IDs is preserved
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
				},
				Check: checkAllPolicyPermGroupIDsUnchanged(resourceName, &allPermGroupIDs),
			},
			{
				Config: acctest.LoadTestCase("account_token-policy-order2.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				Check: checkAllPolicyPermGroupIDsUnchanged(resourceName, &allPermGroupIDs),
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
				},
				Check: checkAllPolicyPermGroupIDsUnchanged(resourceName, &allPermGroupIDs),
			},
			{
				Config: acctest.LoadTestCase("account_token-policy-order1.tf", rnd+"-updated", accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				Check: checkAllPolicyPermGroupIDsUnchanged(resourceName, &allPermGroupIDs),
			},
			// Import step — ignore policies ordering since import uses canonical
			// sort (no prior state) which may differ from the last config order
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"value", "policies"},
			},
		},
	})

	// Test the reverse order scenario
	var allPermGroupIDsReverse []string

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
				},
				Check: checkAllPolicyPermGroupIDsUnchanged(resourceName, &allPermGroupIDsReverse),
			},
			{
				Config: acctest.LoadTestCase("account_token-policy-order1.tf", rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				Check: checkAllPolicyPermGroupIDsUnchanged(resourceName, &allPermGroupIDsReverse),
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
				},
				Check: checkAllPolicyPermGroupIDsUnchanged(resourceName, &allPermGroupIDsReverse),
			},
			{
				Config: acctest.LoadTestCase("account_token-policy-order2.tf", rnd+"-updated", accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(2)),
				},
				Check: checkAllPolicyPermGroupIDsUnchanged(resourceName, &allPermGroupIDsReverse),
			},
			// Import step — ignore policies ordering since import uses canonical
			// sort (no prior state) which may differ from the last config order
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"value", "policies"},
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
			// Import step — ignore policies ordering since import uses canonical
			// sort (no prior state) which may differ from the last config order
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"value", "policies"},
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

func TestAccUpgradeAccountToken_FromPublishedV5(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	config := acctest.LoadTestCase("account_token-without-condition.tf", rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.16.0",
					},
				},
				Config: config,
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}
