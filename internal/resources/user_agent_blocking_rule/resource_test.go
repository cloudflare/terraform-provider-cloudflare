package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareUserAgentBlockingRule_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the UA
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_user_agent_blocking_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareUserAgentBlockingRule(rnd, zoneID, "js_challenge"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "mode", "js_challenge"),
					resource.TestCheckResourceAttr(name, "paused", "false"),
					resource.TestCheckResourceAttr(name, "description", "My description"),
					resource.TestCheckResourceAttr(name, "configuration.0.target", "ua"),
					resource.TestCheckResourceAttr(name, "configuration.0.value", "Mozilla"),
				),
			},
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
		CheckDestroy: testAccCheckCloudflareUserAgentBlockingRulesDestroy,
	})
}

func testAccCloudflareUserAgentBlockingRule(rnd, zoneID, mode string) string {
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
		if rs.Type != "cloudflare_user_agent_blocking_rule" {
			continue
		}

		_, err := client.UserAgentRule(context.Background(), rs.Primary.Attributes[consts.ZoneIDSchemaKey], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("User Agent Blocking Rule still exists")
		}
	}

	return nil
}
