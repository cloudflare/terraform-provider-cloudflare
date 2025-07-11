package api_token_test

import (
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAPIToken_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceID := "cloudflare_api_token." + rnd
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPITokenWithoutCondition(rnd, rnd, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, "name", rnd),
				),
			},
			{
				Config: testAccCloudflareAPITokenWithoutCondition(rnd, rnd+"-updated", permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, "name", rnd+"-updated"),
				),
			},
		},
	})
}

func TestAccAPIToken_DoesNotSetConditions(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_api_token." + rnd
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPITokenWithoutCondition(rnd, rnd, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckNoResourceAttr(name, "condition.request_ip.0.in"),
					resource.TestCheckNoResourceAttr(name, "condition.request_ip.0.not_in"),
				),
			},
		},
	})
}

func testAccCloudflareAPITokenWithoutCondition(resourceName, rnd, permissionID string) string {
	return acctest.LoadTestCase("apitokenwithoutcondition.tf", resourceName, rnd, permissionID)
}

func TestAccAPIToken_SetIndividualCondition(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_api_token." + rnd
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPITokenWithIndividualCondition(rnd, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "condition.request_ip.in.0", "192.0.2.1/32"),
					resource.TestCheckNoResourceAttr(name, "condition.request_ip.not_in"),
				),
			},
		},
	})
}

func testAccCloudflareAPITokenWithIndividualCondition(rnd string, permissionID string) string {
	return acctest.LoadTestCase("apitokenwithindividualcondition.tf", rnd, permissionID)
}

func TestAccAPIToken_SetAllCondition(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_api_token." + rnd
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPITokenWithAllCondition(rnd, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "condition.request_ip.in.0", "192.0.2.1/32"),
					resource.TestCheckResourceAttr(name, "condition.request_ip.not_in.0", "198.51.100.1/32"),
				),
			},
		},
	})
}

func testAccCloudflareAPITokenWithAllCondition(rnd string, permissionID string) string {
	return acctest.LoadTestCase("apitokenwithallcondition.tf", rnd, permissionID)
}

func TestAccAPIToken_TokenTTL(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_api_token." + rnd
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPITokenWithTTL(rnd, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "not_before", "2018-07-01T05:20:00Z"),
					resource.TestCheckResourceAttr(name, "expires_on", "2032-01-01T00:00:00Z"),
				),
			},
		},
	})
}

func testAccCloudflareAPITokenWithTTL(rnd string, permissionID string) string {
	return acctest.LoadTestCase("apitokenwithttl.tf", rnd, permissionID)
}
