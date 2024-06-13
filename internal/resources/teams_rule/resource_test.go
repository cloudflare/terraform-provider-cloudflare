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

func TestAccCloudflareTeamsRule_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_teams_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", "desc"),
					resource.TestCheckResourceAttr(name, "precedence", "12302"),
					resource.TestCheckResourceAttr(name, "action", "block"),
					resource.TestCheckResourceAttr(name, "filters.0", "dns"),
					resource.TestCheckResourceAttr(name, "traffic", "any(dns.domains[*] == \"example.com\")"),
					resource.TestCheckResourceAttr(name, "rule_settings.#", "1"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.block_page_enabled", "true"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.block_page_reason", "cuz"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.insecure_disable_dnssec_validation", "false"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.egress.0.ipv4", "203.0.113.1"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.egress.0.ipv6", "2001:db8::/32"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.untrusted_cert.0.action", "error"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.payload_log.0.enabled", "true"),
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
    block_page_enabled = true
    block_page_reason = "cuz"
    insecure_disable_dnssec_validation = false
	egress {
		ipv4 = "203.0.113.1"
		ipv6 = "2001:db8::/32"
	}
	untrusted_cert {
		action = "error"
	}
	payload_log {
		enabled = true
	}
  }
}
`, rnd, accountID)
}

func TestAccCloudflareTeamsRule_NoSettings(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_teams_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigBasic(rnd, accountID),
			},
			{
				Config: testAccCloudflareTeamsRuleConfigNoSettings(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", "desc"),
					resource.TestCheckResourceAttr(name, "precedence", "12302"),
					resource.TestCheckResourceAttr(name, "action", "block"),
					resource.TestCheckResourceAttr(name, "filters.0", "dns"),
					resource.TestCheckResourceAttr(name, "traffic", "any(dns.domains[*] == \"example.com\")"),
					resource.TestCheckResourceAttr(name, "rule_settings.#", "0"),
				),
			},
		},
	})
}

func testAccCloudflareTeamsRuleConfigNoSettings(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_teams_rule" "%[1]s" {
  name = "%[1]s"
  account_id = "%[2]s"
  description = "desc"
  precedence = 12302
  action = "block"
  filters = ["dns"]
  traffic = "any(dns.domains[*] == \"example.com\")"
}
`, rnd, accountID)
}

func testAccCheckCloudflareTeamsRuleDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_teams_rule" {
			continue
		}

		_, err := client.TeamsRule(context.Background(), rs.Primary.Attributes[consts.AccountIDSchemaKey], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("teams rule still exists")
		}
	}

	return nil
}
