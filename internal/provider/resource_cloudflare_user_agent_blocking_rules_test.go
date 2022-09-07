package provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudflareUserAgentBlockingRules(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		defer func(apiToken string) {
			os.Setenv("CLOUDFLARE_API_TOKEN", apiToken)
		}(os.Getenv("CLOUDFLARE_API_TOKEN"))
		os.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_user_agent_blocking_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareUserAgentBlockingRules(rnd, zoneID, "js_challenge"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "mode", "js_challenge"),
					resource.TestCheckResourceAttr(name, "paused", "false"),
					resource.TestCheckResourceAttr(name, "description", "My description"),
					resource.TestCheckResourceAttr(name, "configuration.0.target", "ua"),
					resource.TestCheckResourceAttr(name, "configuration.0.value", "Mozilla"),
				),
			},
		},
		CheckDestroy: testAccCheckCloudflareUserAgentBlockingRulesDestroy,
	})
}

func testAccCloudflareUserAgentBlockingRules(rnd, zoneID, mode string) string {
	return fmt.Sprintf(`
resource "cloudflare_user_agent_blocking_rule" "%[1]s" {
	zone_id     = "%[2]s"
	mode        = "%[3]s"
	paused      = false
	description = "My description"
	configuration {
		target = "ua"
		value  = "Mozilla"
	}
}
`, rnd, zoneID, mode)
}

func testAccCheckCloudflareUserAgentBlockingRulesDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_device_posture_rule" {
			continue
		}

		_, err := client.UserAgentRule(context.Background(), rs.Primary.Attributes["zone_id"], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Device Posture Rule still exists")
		}
	}

	return nil
}
