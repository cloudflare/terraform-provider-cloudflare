package account_token_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestAccAccountToken_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceID := "cloudflare_account_token." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccountTokenWithoutCondition(rnd, accountID, rnd, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, "name", rnd),
					resource.TestCheckResourceAttr(resourceID, "policies.0.permission_groups.0.id", permissionID),
				),
			},
			{
				Config: testAccCloudflareAccountTokenWithoutCondition(rnd, accountID, rnd+"-updated", permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, "name", rnd+"-updated"),
					resource.TestCheckResourceAttr(resourceID, "policies.0.permission_groups.0.id", permissionID),
				),
			},
		},
	})
}

func TestAccAccountToken_DoesNotSetConditions(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_account_token." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccountTokenWithoutCondition(rnd, accountID, rnd, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckNoResourceAttr(name, "condition.request_ip.0.in"),
					resource.TestCheckNoResourceAttr(name, "condition.request_ip.0.not_in"),
				),
			},
		},
	})
}

func testAccCloudflareAccountTokenWithoutCondition(resourceName, accountId, rnd, permissionID string) string {
	return acctest.LoadTestCase("account_token-without-condition.tf", resourceName, accountId, rnd, permissionID)
}

func TestAccAccountToken_SetIndividualCondition(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_account_token." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccountTokenWithIndividualCondition(rnd, accountID, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "condition.request_ip.in.0", "192.0.2.1/32"),
					resource.TestCheckNoResourceAttr(name, "condition.request_ip.not_in"),
				),
			},
		},
	})
}

func testAccCloudflareAccountTokenWithIndividualCondition(rnd, accountID, permissionID string) string {
	return acctest.LoadTestCase("account_token-with-individual-condition.tf", rnd, accountID, permissionID)
}

func TestAccAccountToken_SetAllCondition(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_account_token." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccountTokenWithAllCondition(rnd, accountID, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "condition.request_ip.in.0", "192.0.2.1/32"),
					resource.TestCheckResourceAttr(name, "condition.request_ip.not_in.0", "198.51.100.1/32"),
				),
			},
		},
	})
}

func testAccCloudflareAccountTokenWithAllCondition(rnd, accountID, permissionID string) string {
	return acctest.LoadTestCase("account_token-with-all-condition.tf", rnd, accountID, permissionID)
}

func TestAccAccountToken_TokenTTL(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_account_token." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccountTokenWithTTL(rnd, accountID, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "not_before", "2018-07-01T05:20:00Z"),
					resource.TestCheckResourceAttr(name, "expires_on", "2032-01-01T00:00:00Z"),
				),
			},
		},
	})
}

func testAccCloudflareAccountTokenWithTTL(rnd, accountID, permissionID string) string {
	return acctest.LoadTestCase("account_token-with-ttl.tf", rnd, accountID, permissionID)
}

func TestAccAccountToken_PermissionGroupOrder(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_account_token." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	permissionID1 := "82e64a83756745bbbb1c9c2701bf816b" // DNS read
	permissionID2 := "e199d584e69344eba202452019deafe3" // Disable ESC read

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("account_token-permissiongroup-order.tf", rnd, accountID, permissionID1, permissionID2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "policies.0.permission_groups.0.id", permissionID1),
					resource.TestCheckResourceAttr(name, "policies.0.permission_groups.1.id", permissionID2),
				),
			},
			{
				Config: acctest.LoadTestCase("account_token-permissiongroup-order.tf", rnd, accountID, permissionID2, permissionID1),
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
				Config: acctest.LoadTestCase("account_token-permissiongroup-order.tf", rnd, accountID, permissionID2, permissionID1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "policies.0.permission_groups.0.id", permissionID1),
					resource.TestCheckResourceAttr(name, "policies.0.permission_groups.1.id", permissionID2),
				),
			},
			{
				Config: acctest.LoadTestCase("account_token-permissiongroup-order.tf", rnd, accountID, permissionID2, permissionID1),
				// re-applying same change does not produce drift
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				Config: acctest.LoadTestCase("account_token-permissiongroup-order.tf", rnd, accountID, permissionID1, permissionID2),
				// changing the order of permission groups should not affect plan
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}
