package page_shield_policy_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflarePageShieldPolicy_Basic(t *testing.T) {
	randomID := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_page_shield_policy." + randomID
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testPageShieldPolicyConfig(randomID, zoneID, "description", "log", "true", "false", "script-src 'none'"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "action", "log"),
					resource.TestCheckResourceAttr(resourceName, "expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "value", "script-src 'none'"),
				),
			},
			{
				Config: testPageShieldPolicyConfig(
					randomID,
					zoneID,
					"description2",
					"allow",
					"(http.request.uri.path contains \\\"checkout\\\")",
					"true",
					"script-src 'none' connect-src 'none'",
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "action", "allow"),
					resource.TestCheckResourceAttr(resourceName, "expression", "(http.request.uri.path contains \"checkout\")"),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "value", "script-src 'none' connect-src 'none'"),
				),
			},
		},
	})
}

func testPageShieldPolicyConfig(resourceName, zoneID, description, action, expression, enabled, value string) string {
	return acctest.LoadTestCase("page_shield_policy.tf", resourceName, zoneID, description, action, expression, enabled, value)
}
