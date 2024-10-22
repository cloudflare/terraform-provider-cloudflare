package page_shield_policy_test

import (
	"context"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflarePageShieldPolicy_Create(t *testing.T) {
	ctx := context.Background()
	randomID := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_page_shield_policy." + randomID

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Error(ctx, "CLOUDFLARE_ZONE_ID must be set for cloudflare_page_shield_policy")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
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
		},
	})
}

func TestAccCloudflarePageShieldPolicy_Update(t *testing.T) {
	ctx := context.Background()
	randomID := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_page_shield_policy." + randomID

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Error(ctx, "CLOUDFLARE_ZONE_ID must be set for cloudflare_page_shield_policy")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testPageShieldPolicyConfig(randomID, zoneID, "description", "log", "true", "false", "script-src 'none'"),
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
