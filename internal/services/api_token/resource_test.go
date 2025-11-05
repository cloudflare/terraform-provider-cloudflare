package api_token_test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/user"
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
	resource.AddTestSweepers("cloudflare_api_token", &resource.Sweeper{
		Name: "cloudflare_api_token",
		F:    testSweepCloudflareAPIToken,
	})
}

func testSweepCloudflareAPIToken(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	// List all API tokens
	tokens, err := client.User.Tokens.List(ctx, user.TokenListParams{})
	if err != nil {
		return fmt.Errorf("failed to fetch API tokens: %w", err)
	}

	// Delete test tokens (those created by our tests)
	for _, token := range tokens.Result {
		if strings.Contains(token.Name, "terraform") {
			_, err := client.User.Tokens.Delete(ctx, token.ID)
			if err != nil {
				return fmt.Errorf("failed to delete API token %s: %w", token.ID, err)
			}
		}
	}

	return nil
}

func testAccCheckCloudflareAPITokenDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_api_token" {
			continue
		}

		tokenID := rs.Primary.ID
		_, err := client.User.Tokens.Get(context.Background(), tokenID)
		if err == nil {
			return fmt.Errorf("API token %s still exists", tokenID)
		}
	}

	return nil
}

func TestAccAPIToken_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_token.test_account_token"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAPITokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("api_token-without-condition.tf", rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
				},
				Check: resource.ComposeTestCheckFunc(
					// Verify conditions are not set (ConfigStateChecks can't easily check for absence)
					resource.TestCheckNoResourceAttr(resourceName, "condition.request_ip.0.in"),
					resource.TestCheckNoResourceAttr(resourceName, "condition.request_ip.0.not_in"),
				),
			},
			{
				Config: acctest.LoadTestCase("api_token-without-condition.tf", rnd),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			{
				Config: acctest.LoadTestCase("api_token-without-condition.tf", rnd+"-updated"),
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
					// Verify conditions still not set
					resource.TestCheckNoResourceAttr(resourceName, "condition.request_ip.0.in"),
					resource.TestCheckNoResourceAttr(resourceName, "condition.request_ip.0.not_in"),
				),
			},
			{
				Config: acctest.LoadTestCase("api_token-without-condition.tf", rnd+"-updated"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
		},
	})
}

func TestAccAPIToken_SetIndividualCondition(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_token.test_account_token"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAPITokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("api_token-with-individual-condition.tf", rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in").AtSliceIndex(0), knownvalue.StringExact("192.0.2.1/32")),
				},
				Check: resource.ComposeTestCheckFunc(
					// Verify not_in condition is not set
					resource.TestCheckNoResourceAttr(resourceName, "condition.request_ip.not_in"),
				),
			},
			{
				Config: acctest.LoadTestCase("api_token-with-individual-condition.tf", rnd),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
		},
	})
}

func TestAccAPIToken_SetAllCondition(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_token.test_account_token"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAPITokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("api_token-with-all-condition.tf", rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					// Validate "in" condition
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in").AtSliceIndex(0), knownvalue.StringExact("192.0.2.1/32")),
					// Validate "not_in" condition
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("not_in"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("not_in").AtSliceIndex(0), knownvalue.StringExact("198.51.100.1/32")),
				},
			},
		},
	})
}

func TestAccAPIToken_TokenTTL(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_token.test_account_token"

	oneDaysFromNow := time.Now().UTC().AddDate(0, 0, 1)
	expireTime := oneDaysFromNow.Format(time.RFC3339)
	twoDaysFromNow := time.Now().UTC().AddDate(0, 0, 2)
	updatedExpireTime := twoDaysFromNow.Format(time.RFC3339)
	notBefore := "2018-07-01T05:20:00Z"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAPITokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("api_token-with-ttl.tf", rnd, expireTime),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("not_before"), knownvalue.StringExact(notBefore)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact(expireTime)),
				},
			},
			{
				Config: acctest.LoadTestCase("api_token-with-ttl.tf", rnd, expireTime),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			{
				Config: acctest.LoadTestCase("api_token-with-ttl.tf", rnd, updatedExpireTime),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact(updatedExpireTime)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("not_before"), knownvalue.StringExact(notBefore)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact(updatedExpireTime)),
				},
			},
		},
	})
}

// Test that permission group order doesn't affect plans. This test uses two
// corresponding .tf files that are identical except they swap the order of two
// permission groups on the same policy. We use statecheck.CompareValue to
// assert that permission group 0 never changes and permission group 1 never
// changes despite changing the order of the permission groups and updating the
// token name.
func TestAccAPIToken_PermissionGroupOrder(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_token.test_account_token"

	permgroup0_SAME := statecheck.CompareValue(compare.ValuesSame())
	permgroup1_SAME := statecheck.CompareValue(compare.ValuesSame())

	// Test that permission group order doesn't affect plans
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAPITokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("api_token-permissiongroup-order1.tf", rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups"), knownvalue.ListSizeExact(2)),
					permgroup0_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id")),
					permgroup1_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(1).AtMapKey("id")),
				},
			},
			{
				Config: acctest.LoadTestCase("api_token-permissiongroup-order2.tf", rnd),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			{
				Config: acctest.LoadTestCase("api_token-permissiongroup-order2.tf", rnd+"-updated"),
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
				Config: acctest.LoadTestCase("api_token-permissiongroup-order1.tf", rnd+"-updated"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
		},
	})

	// Test the reverse order scenario
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAPITokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("api_token-permissiongroup-order2.tf", rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups"), knownvalue.ListSizeExact(2)),
					permgroup0_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id")),
					permgroup1_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(1).AtMapKey("id")),
				},
			},
			{
				Config: acctest.LoadTestCase("api_token-permissiongroup-order1.tf", rnd),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			{
				Config: acctest.LoadTestCase("api_token-permissiongroup-order1.tf", rnd+"-updated"),
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
				Config: acctest.LoadTestCase("api_token-permissiongroup-order2.tf", rnd+"-updated"),
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
		},
	})
}

// Test that policy group order doesn't affect plans. This test uses two
// corresponding .tf files that are identical except they swap the order of two
// policies.
func TestAccAPIToken_PolicyOrder(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_token.test_account_token"

	policy1_permgroup_SAME := statecheck.CompareValue(compare.ValuesSame())
	policy2_permgroup_SAME := statecheck.CompareValue(compare.ValuesSame())

	// Test that policy order doesn't affect plans
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAPITokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("api_token-policy-order1.tf", rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(2)),
					policy1_permgroup_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id")),
					policy2_permgroup_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(1).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id")),
				},
			},
			{
				Config: acctest.LoadTestCase("api_token-policy-order2.tf", rnd),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			{
				Config: acctest.LoadTestCase("api_token-policy-order2.tf", rnd+"-updated"),
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
				Config: acctest.LoadTestCase("api_token-policy-order1.tf", rnd+"-updated"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
		},
	})

	// Test the reverse order scenario
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAPITokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("api_token-policy-order2.tf", rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(2)),
					policy1_permgroup_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id")),
					policy2_permgroup_SAME.AddStateValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(1).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id")),
				},
			},
			{
				Config: acctest.LoadTestCase("api_token-policy-order1.tf", rnd),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			{
				Config: acctest.LoadTestCase("api_token-policy-order1.tf", rnd+"-updated"),
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
				Config: acctest.LoadTestCase("api_token-policy-order2.tf", rnd+"-updated"),
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
		},
	})
}

func TestAccAPIToken_CRUD(t *testing.T) {
	// Comprehensive test covering Create, Read, Update, Delete + Import
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_token.crud_test"

	initialName := rnd + "-initial"
	updatedName := rnd + "-updated"

	// TTL times
	oneDayFromNow := time.Now().UTC().AddDate(0, 0, 1).Format(time.RFC3339)
	twoDaysFromNow := time.Now().UTC().AddDate(0, 0, 2).Format(time.RFC3339)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAPITokenDestroy,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAPITokenCRUDConfig(initialName, oneDayFromNow, "192.0.2.1/32"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(initialName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact(oneDayFromNow)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in").AtSliceIndex(0), knownvalue.StringExact("192.0.2.1/32")),
				},
			},
			// Update name and TTL
			{
				Config: testAccAPITokenCRUDConfig(updatedName, twoDaysFromNow, "192.0.2.1/32"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(updatedName)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact(twoDaysFromNow)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(updatedName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact(twoDaysFromNow)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in").AtSliceIndex(0), knownvalue.StringExact("192.0.2.1/32")),
				},
			},
			// Update condition
			{
				Config: testAccAPITokenCRUDConfig(updatedName, twoDaysFromNow, "198.51.100.1/32"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in").AtSliceIndex(0), knownvalue.StringExact("198.51.100.1/32")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(updatedName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in").AtSliceIndex(0), knownvalue.StringExact("198.51.100.1/32")),
				},
			},
			// Import
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"value"}, // Token value is write-only
			},
		},
	})
}

func testAccAPITokenCRUDConfig(name, expiresOn, ipCondition string) string {
	return fmt.Sprintf(`
data "cloudflare_api_token_permission_groups_list" "dns_read" {
  name  = "DNS Read"
  scope = "com.cloudflare.api.account.zone"
}

resource "cloudflare_api_token" "crud_test" {
  name       = "%[1]s"
  expires_on = "%[2]s"

  policies = [
    {
      effect = "allow"
      permission_groups = [
        { id = data.cloudflare_api_token_permission_groups_list.dns_read.result[0].id }
      ]
      resources = {
        "com.cloudflare.api.account.zone.*" = "*"
      }
    }
  ]

  condition = {
    request_ip = {
      in = ["%[3]s"]
    }
  }
}
`, name, expiresOn, ipCondition)
}
