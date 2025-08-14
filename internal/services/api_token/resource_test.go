package api_token_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestAccAPIToken_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()

	var policyId string

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("api_token-without-condition.tf", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "name", rnd),
					resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "policies.#", "1"),
					resource.TestCheckResourceAttrSet("cloudflare_api_token.test_account_token", "policies.0.id"),
					resource.TestCheckResourceAttrWith("cloudflare_api_token.test_account_token", "policies.0.id", func(value string) error {
						policyId = value
						return nil
					}),
					// conditions by default should not be set
					resource.TestCheckNoResourceAttr("cloudflare_api_token.test_account_token", "condition.request_ip.0.in"),
					resource.TestCheckNoResourceAttr("cloudflare_api_token.test_account_token", "condition.request_ip.0.not_in"),
				),
			},
			{
				Config: acctest.LoadTestCase("api_token-without-condition.tf", rnd),
				// re-plan should not detect drift
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				Config: acctest.LoadTestCase("api_token-without-condition.tf", rnd+"-updated"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "name", rnd+"-updated"),
					resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "policies.#", "1"),
					resource.TestCheckResourceAttrSet("cloudflare_api_token.test_account_token", "policies.0.id"),
					resource.TestCheckResourceAttrWith("cloudflare_api_token.test_account_token", "policies.0.id", func(value string) error {
						if value != policyId {
							return fmt.Errorf("policy ID changed from %s to %s", policyId, value)
						}
						return nil
					}),
					// conditions still not set
					resource.TestCheckNoResourceAttr("cloudflare_api_token.test_account_token", "condition.request_ip.0.in"),
					resource.TestCheckNoResourceAttr("cloudflare_api_token.test_account_token", "condition.request_ip.0.not_in"),
				),
			},
			{
				Config: acctest.LoadTestCase("api_token-without-condition.tf", rnd+"-updated"),
				// re-plan should not detect drift
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestAccAPIToken_SetIndividualCondition(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("api_token-with-individual-condition.tf", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "name", rnd),
					resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "condition.request_ip.in.0", "192.0.2.1/32"),
					resource.TestCheckNoResourceAttr("cloudflare_api_token.test_account_token", "condition.request_ip.not_in"),
				),
			},
			{
				Config: acctest.LoadTestCase("api_token-with-individual-condition.tf", rnd),
				// re-plan should not detect drift
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestAccAPIToken_SetAllCondition(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("api_token-with-all-condition.tf", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "name", rnd),
					resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "condition.request_ip.in.0", "192.0.2.1/32"),
					resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "condition.request_ip.not_in.0", "198.51.100.1/32"),
				),
			},
		},
	})
}

func TestAccAPIToken_TokenTTL(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()

	oneDaysFromNow := time.Now().UTC().AddDate(0, 0, 1)
	expireTime := oneDaysFromNow.Format(time.RFC3339)
	twoDaysFromNow := time.Now().UTC().AddDate(0, 0, 2)
	updatedExpireTime := twoDaysFromNow.Format(time.RFC3339)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("api_token-with-ttl.tf", rnd, expireTime),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "name", rnd),
					resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "not_before", "2018-07-01T05:20:00Z"),
					resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "expires_on", expireTime),
				),
			},
			{
				Config: acctest.LoadTestCase("api_token-with-ttl.tf", rnd, expireTime),
				// re-plan should not detect drift
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				Config: acctest.LoadTestCase("api_token-with-ttl.tf", rnd, updatedExpireTime),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "name", rnd),
					resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "not_before", "2018-07-01T05:20:00Z"),
					resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "expires_on", updatedExpireTime),
				),
			},
		},
	})
}

func TestAccAPIToken_PermissionGroupOrder(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_api_token." + rnd
	permissionID1 := "82e64a83756745bbbb1c9c2701bf816b" // DNS read
	permissionID2 := "e199d584e69344eba202452019deafe3" // Disable ESC read

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("api_token-permissiongroup-order.tf", rnd, permissionID1, permissionID2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "policies.0.permission_groups.0.id", permissionID1),
					resource.TestCheckResourceAttr(name, "policies.0.permission_groups.1.id", permissionID2),
				),
			},
			{
				Config: acctest.LoadTestCase("api_token-permissiongroup-order.tf", rnd, permissionID2, permissionID1),
				// changing the order of permission groups should not affect plan
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("api_token-permissiongroup-order.tf", rnd, permissionID2, permissionID1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "policies.0.permission_groups.0.id", permissionID1),
					resource.TestCheckResourceAttr(name, "policies.0.permission_groups.1.id", permissionID2),
				),
			},
			{
				Config: acctest.LoadTestCase("api_token-permissiongroup-order.tf", rnd, permissionID2, permissionID1),
				// re-applying same change does not produce drift
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				Config: acctest.LoadTestCase("api_token-permissiongroup-order.tf", rnd, permissionID1, permissionID2),
				// changing the order of permission groups should not affect plan
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})

	// rnd := utils.GenerateRandomResourceName()
	// permissionID0 := ""
	// permissionID1 := ""

	// var policyId string

	// resource.Test(t, resource.TestCase{
	// 	PreCheck:                 func() { acctest.TestAccPreCheck(t) },
	// 	ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
	// 	Steps: []resource.TestStep{
	// 		{
	// 			// not seting permission IDs first, retrieving them from API by name
	// 			Config: acctest.LoadTestCase("api_token-permissiongroup-order.tf", rnd, "", ""),
	// 			Check: resource.ComposeTestCheckFunc(
	// 				resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "name", rnd),
	// 				resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "policies.#", "1"),
	// 				resource.TestCheckResourceAttrSet("cloudflare_api_token.test_account_token", "policies.0.id"),
	// 				resource.TestCheckResourceAttrWith("cloudflare_api_token.test_account_token", "policies.0.id", func(value string) error {
	// 					policyId = value
	// 					return nil
	// 				}),
	// 				resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "policies.0.permission_groups.#", "2"),
	// 				resource.TestCheckResourceAttrWith("cloudflare_api_token.test_account_token", "policies.0.permission_groups.0.id", func(value string) error {
	// 					permissionID0 = value
	// 					return nil
	// 				}),
	// 				resource.TestCheckResourceAttrWith("cloudflare_api_token.test_account_token", "policies.0.permission_groups.1.id", func(value string) error {
	// 					permissionID1 = value
	// 					return nil
	// 				}),
	// 			),
	// 		},
	// 		// below we try changing the order of the permission group IDs and
	// 		// verify there are no plan changes
	// 		{
	// 			Config: acctest.LoadTestCase("api_token-permissiongroup-order.tf", rnd, permissionID0, permissionID1),
	// 			ConfigPlanChecks: resource.ConfigPlanChecks{
	// 				PreApply: []plancheck.PlanCheck{
	// 					plancheck.ExpectEmptyPlan(),
	// 				},
	// 			},
	// 		},
	// 		{
	// 			Config: acctest.LoadTestCase("api_token-permissiongroup-order.tf", rnd, permissionID1, permissionID0),
	// 			ConfigPlanChecks: resource.ConfigPlanChecks{
	// 				PreApply: []plancheck.PlanCheck{
	// 					plancheck.ExpectEmptyPlan(),
	// 				},
	// 			},
	// 		},
	// 		// try updating the token and ensure policy information hasn't
	// 		// changed
	// 		{
	// 			Config: acctest.LoadTestCase("api_token-permissiongroup-order.tf", rnd+"updated", permissionID1, permissionID0),
	// 			ConfigPlanChecks: resource.ConfigPlanChecks{
	// 				PreApply: []plancheck.PlanCheck{
	// 					plancheck.ExpectNonEmptyPlan(),
	// 				},
	// 			},
	// 			Check: resource.ComposeTestCheckFunc(
	// 				resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "name", rnd+"updated"),
	// 				resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "policies.#", "1"),
	// 				resource.TestCheckResourceAttrWith("cloudflare_api_token.test_account_token", "policies.0.id", func(value string) error {
	// 					if value != policyId {
	// 						return fmt.Errorf("policy ID changed from %s to %s", policyId, value)
	// 					}
	// 					return nil
	// 				}),
	// 				resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "policies.0.permission_groups.#", "2"),
	// 				resource.TestCheckResourceAttrWith("cloudflare_api_token.test_account_token", "policies.0.permission_groups.0.id", func(value string) error {
	// 					if value != permissionID0 {
	// 						return fmt.Errorf("permission ID 0 changed from %s to %s", permissionID0, value)
	// 					}
	// 					return nil
	// 				}),
	// 				resource.TestCheckResourceAttrWith("cloudflare_api_token.test_account_token", "policies.0.permission_groups.1.id", func(value string) error {
	// 					if value != permissionID1 {
	// 						return fmt.Errorf("permission ID 1 changed from %s to %s", permissionID1, value)
	// 					}
	// 					return nil
	// 				}),
	// 			),
	// 		},
	// 		{
	// 			Config: acctest.LoadTestCase("api_token-permissiongroup-order.tf", rnd+"updated2", permissionID0, permissionID1),
	// 			ConfigPlanChecks: resource.ConfigPlanChecks{
	// 				PreApply: []plancheck.PlanCheck{
	// 					plancheck.ExpectNonEmptyPlan(),
	// 				},
	// 			},
	// 			Check: resource.ComposeTestCheckFunc(
	// 				resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "name", rnd+"updated2"),
	// 				resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "policies.#", "1"),
	// 				resource.TestCheckResourceAttrWith("cloudflare_api_token.test_account_token", "policies.0.id", func(value string) error {
	// 					if value != policyId {
	// 						return fmt.Errorf("policy ID changed from %s to %s", policyId, value)
	// 					}
	// 					return nil
	// 				}),
	// 				resource.TestCheckResourceAttr("cloudflare_api_token.test_account_token", "policies.0.permission_groups.#", "2"),
	// 				resource.TestCheckResourceAttrWith("cloudflare_api_token.test_account_token", "policies.0.permission_groups.0.id", func(value string) error {
	// 					if value != permissionID0 {
	// 						return fmt.Errorf("permission ID 0 changed from %s to %s", permissionID0, value)
	// 					}
	// 					return nil
	// 				}),
	// 				resource.TestCheckResourceAttrWith("cloudflare_api_token.test_account_token", "policies.0.permission_groups.1.id", func(value string) error {
	// 					if value != permissionID1 {
	// 						return fmt.Errorf("permission ID 1 changed from %s to %s", permissionID1, value)
	// 					}
	// 					return nil
	// 				}),
	// 			),
	// 		},
	// 	},
	// })
}

func TestAccAPIToken_Resources_SimpleToNested_NoDrift(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_api_token." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_APIToken(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Simple: wildcard string mapping
				Config: acctest.LoadTestCase("api_token-resources-simple.tf", rnd, rnd, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "policies.0.permission_groups.0.id", permissionID),
				),
			},
			{
				// Re-apply should produce empty plan (no drift)
				Config: acctest.LoadTestCase("api_token-resources-simple.tf", rnd, rnd, permissionID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				// Nested: account -> { zone.* = "*" }
				Config: acctest.LoadTestCase("api_token-resources-nested.tf", rnd, rnd, permissionID, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "policies.0.permission_groups.0.id", permissionID),
				),
			},
			{
				// Re-apply nested should produce empty plan (no drift)
				Config: acctest.LoadTestCase("api_token-resources-nested.tf", rnd, rnd, permissionID, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}
