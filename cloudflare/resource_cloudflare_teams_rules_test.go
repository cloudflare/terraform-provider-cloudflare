package cloudflare

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudflareTeamsRuleBasic(t *testing.T) {
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
	name := fmt.Sprintf("cloudflare_teams_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccessAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", "desc"),
					resource.TestCheckResourceAttr(name, "precedence", "12302"),
					resource.TestCheckResourceAttr(name, "action", "block"),
					resource.TestCheckResourceAttr(name, "filters.0", "dns"),
					resource.TestCheckResourceAttr(name, "traffic", "any(dns.domains[*] == \"example.com\")"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.block_page_enabled", "false"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.block_page_reason", "cuz"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.check_session.enforce", "true"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.check_session.duration", "200s"),
				),
			},
		},
	})
}

func testAccCloudflareTeamsRuleConfigBasic(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_teams_rule" "%[1]s" {
  name = "%[1]s"
  account_id = "%[2]s"
  description = "desc"
  precedence = 12302
  action = "block"
  filters = ["dns"]
  traffic = "any(dns.domains[*] == \"example.com\")"
  rule_settings {
    block_page_enabled = false
    block_page_reason = "cuz"
	check_session {
		enforce = true
		duration = "200s"
	}
  }
}
`, rnd, accountID)
}

func testAccCheckCloudflareTeamsRuleDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_teams_rule" {
			continue
		}

		_, err := client.TeamsRule(context.Background(), rs.Primary.Attributes["account_id"], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("teams rule still exists")
		}
	}

	return nil
}
