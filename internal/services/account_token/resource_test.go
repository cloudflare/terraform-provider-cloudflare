package account_token_test

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

func TestAccAccountToken_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var policyId string

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("account_token-without-condition.tf", rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "name", rnd),
					resource.TestCheckResourceAttrSet("cloudflare_account_token.test_account_token", "policies.0.id"),
					resource.TestCheckResourceAttrWith("cloudflare_account_token.test_account_token", "policies.0.id", func(value string) error {
						policyId = value
						return nil
					}),
					// conditions by default should not be set
					resource.TestCheckNoResourceAttr("cloudflare_account_token.test_account_token", "condition.request_ip.0.in"),
					resource.TestCheckNoResourceAttr("cloudflare_account_token.test_account_token", "condition.request_ip.0.not_in"),
				),
			},
			{
				Config: acctest.LoadTestCase("account_token-without-condition.tf", rnd, accountID),
				// re-plan should not detect drift
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				Config: acctest.LoadTestCase("account_token-without-condition.tf", rnd+"-updated", accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "name", rnd+"-updated"),
					resource.TestCheckResourceAttrSet("cloudflare_account_token.test_account_token", "policies.0.id"),
					resource.TestCheckResourceAttrWith("cloudflare_account_token.test_account_token", "policies.0.id", func(value string) error {
						if value != policyId {
							return fmt.Errorf("policy ID changed from %s to %s", policyId, value)
						}
						return nil
					}),
					// conditions still not be set
					resource.TestCheckNoResourceAttr("cloudflare_account_token.test_account_token", "condition.request_ip.0.in"),
					resource.TestCheckNoResourceAttr("cloudflare_account_token.test_account_token", "condition.request_ip.0.not_in"),
				),
			},
			{
				Config: acctest.LoadTestCase("account_token-without-condition.tf", rnd+"-updated", accountID),
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

func TestAccAccountToken_SetIndividualCondition(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("account_token-with-individual-condition.tf", rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "name", rnd),
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "condition.request_ip.in.0", "192.0.2.1/32"),
					resource.TestCheckNoResourceAttr("cloudflare_account_token.test_account_token", "condition.request_ip.not_in"),
				),
			},
			{
				Config: acctest.LoadTestCase("account_token-with-individual-condition.tf", rnd, accountID),
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

func TestAccAccountToken_SetAllCondition(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("account_token-with-all-condition.tf", rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "name", rnd),
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "condition.request_ip.in.0", "192.0.2.1/32"),
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "condition.request_ip.not_in.0", "198.51.100.1/32"),
				),
			},
			{
				Config: acctest.LoadTestCase("account_token-with-all-condition.tf", rnd, accountID),
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

func TestAccAccountToken_TokenTTL(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	oneDaysFromNow := time.Now().UTC().AddDate(0, 0, 1)
	expireTime := oneDaysFromNow.Format(time.RFC3339)
	twoDaysFromNow := time.Now().UTC().AddDate(0, 0, 2)
	updatedExpireTime := twoDaysFromNow.Format(time.RFC3339)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("account_token-with-ttl.tf", rnd, accountID, expireTime),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "name", rnd),
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "not_before", "2018-07-01T05:20:00Z"),
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "expires_on", expireTime),
				),
			},
			{
				Config: acctest.LoadTestCase("account_token-with-ttl.tf", rnd, accountID, expireTime),
				// re-plan should not detect drift
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				Config: acctest.LoadTestCase("account_token-with-ttl.tf", rnd, accountID, updatedExpireTime),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "name", rnd),
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "not_before", "2018-07-01T05:20:00Z"),
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "expires_on", updatedExpireTime),
				),
			},
		},
	})
}

func TestAccAccountToken_PermissionGroupOrder(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	permissionID0 := ""
	permissionID1 := ""

	var policyId string

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// not seting permission IDs first, retrieving them from API by name
				Config: acctest.LoadTestCase("account_token-permissiongroup-order.tf", rnd, accountID, "", ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "name", rnd),
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "policies.#", "1"),
					resource.TestCheckResourceAttrSet("cloudflare_account_token.test_account_token", "policies.0.id"),
					resource.TestCheckResourceAttrWith("cloudflare_account_token.test_account_token", "policies.0.id", func(value string) error {
						policyId = value
						return nil
					}),
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "policies.0.permission_groups.#", "2"),
					resource.TestCheckResourceAttrWith("cloudflare_account_token.test_account_token", "policies.0.permission_groups.0.id", func(value string) error {
						permissionID0 = value
						return nil
					}),
					resource.TestCheckResourceAttrWith("cloudflare_account_token.test_account_token", "policies.0.permission_groups.1.id", func(value string) error {
						permissionID1 = value
						return nil
					}),
				),
			},
			// below we try changing the order of the permission group IDs and
			// verify there are no plan changes
			{
				Config: acctest.LoadTestCase("account_token-permissiongroup-order.tf", rnd, accountID, permissionID0, permissionID1),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				Config: acctest.LoadTestCase("account_token-permissiongroup-order.tf", rnd, accountID, permissionID1, permissionID0),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// try updating the token and ensure policy information hasn't
			// changed
			{
				Config: acctest.LoadTestCase("account_token-permissiongroup-order.tf", rnd+"updated", accountID, permissionID1, permissionID0),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "name", rnd+"updated"),
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "policies.#", "1"),
					resource.TestCheckResourceAttrWith("cloudflare_account_token.test_account_token", "policies.0.id", func(value string) error {
						if value != policyId {
							return fmt.Errorf("policy ID changed from %s to %s", policyId, value)
						}
						return nil
					}),
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "policies.0.permission_groups.#", "2"),
					resource.TestCheckResourceAttrWith("cloudflare_account_token.test_account_token", "policies.0.permission_groups.0.id", func(value string) error {
						if value != permissionID0 {
							return fmt.Errorf("permission ID 0 changed from %s to %s", permissionID0, value)
						}
						return nil
					}),
					resource.TestCheckResourceAttrWith("cloudflare_account_token.test_account_token", "policies.0.permission_groups.1.id", func(value string) error {
						if value != permissionID1 {
							return fmt.Errorf("permission ID 1 changed from %s to %s", permissionID1, value)
						}
						return nil
					}),
				),
			},
			{
				Config: acctest.LoadTestCase("account_token-permissiongroup-order.tf", rnd+"updated2", accountID, permissionID0, permissionID1),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "name", rnd+"updated2"),
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "policies.#", "1"),
					resource.TestCheckResourceAttrWith("cloudflare_account_token.test_account_token", "policies.0.id", func(value string) error {
						if value != policyId {
							return fmt.Errorf("policy ID changed from %s to %s", policyId, value)
						}
						return nil
					}),
					resource.TestCheckResourceAttr("cloudflare_account_token.test_account_token", "policies.0.permission_groups.#", "2"),
					resource.TestCheckResourceAttrWith("cloudflare_account_token.test_account_token", "policies.0.permission_groups.0.id", func(value string) error {
						if value != permissionID0 {
							return fmt.Errorf("permission ID 0 changed from %s to %s", permissionID0, value)
						}
						return nil
					}),
					resource.TestCheckResourceAttrWith("cloudflare_account_token.test_account_token", "policies.0.permission_groups.1.id", func(value string) error {
						if value != permissionID1 {
							return fmt.Errorf("permission ID 1 changed from %s to %s", permissionID1, value)
						}
						return nil
					}),
				),
			},
		},
	})
}
