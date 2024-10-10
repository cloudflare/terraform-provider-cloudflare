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
	name := fmt.Sprintf("cloudflare_zero_trust_gateway_policy.%s", rnd)

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
resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
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
	name := fmt.Sprintf("cloudflare_zero_trust_gateway_policy.%s", rnd)

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
					resource.TestCheckResourceAttr(name, "rule_settings.#", "1"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.block_page_enabled", "true"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.block_page_reason", "cuz"),
				),
			},
		},
	})
}

func testAccCloudflareTeamsRuleConfigNoSettings(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
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
		if rs.Type != "cloudflare_zero_trust_gateway_policy" {
			continue
		}

		_, err := client.TeamsRule(context.Background(), rs.Primary.Attributes[consts.AccountIDSchemaKey], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("teams rule still exists")
		}
	}

	return nil
}

func TestAccCloudflareTeamsRule_CustomResolver(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_gateway_policy.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigCustomResolver(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "description", "desc"),
					resource.TestCheckResourceAttr(name, "precedence", "12302"),
					resource.TestCheckResourceAttr(name, "action", "resolve"),
					resource.TestCheckResourceAttr(name, "filters.0", "dns_resolver"),
					resource.TestCheckResourceAttr(name, "traffic", "any(dns.domains[*] == \"example.com\")"),
					resource.TestCheckResourceAttr(name, "rule_settings.#", "1"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.block_page_enabled", "false"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.dns_resolvers.0.ipv4.0.ip", "1.34.4.4"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.dns_resolvers.0.ipv4.1.ip", "1.34.4.5"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.dns_resolvers.0.ipv4.2.ip", "1.34.4.6"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.dns_resolvers.0.ipv6.0.ip", "2a09:bac1:76c0:1378::b:248"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.dns_resolvers.0.ipv6.1.ip", "2a09:bac1:76c0:1378::b:148"),

					resource.TestCheckResourceAttr(name, "rule_settings.0.dns_resolvers.0.ipv4.0.port", "53"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.dns_resolvers.0.ipv4.1.port", "30000"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.dns_resolvers.0.ipv4.2.port", "200"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.dns_resolvers.0.ipv6.0.port", "53"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.dns_resolvers.0.ipv6.1.port", "53"),

					resource.TestCheckResourceAttr(name, "rule_settings.0.dns_resolvers.0.ipv4.0.vnet_id", "5fb4dfe1-4fe7-4ff4-980e-72f89dd9af9e"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.dns_resolvers.0.ipv4.1.vnet_id", ""),
					resource.TestCheckResourceAttr(name, "rule_settings.0.dns_resolvers.0.ipv4.2.vnet_id", ""),
					resource.TestCheckResourceAttr(name, "rule_settings.0.dns_resolvers.0.ipv6.0.vnet_id", ""),
					resource.TestCheckResourceAttr(name, "rule_settings.0.dns_resolvers.0.ipv6.1.vnet_id", ""),

					resource.TestCheckResourceAttr(name, "rule_settings.0.dns_resolvers.0.ipv4.0.route_through_private_network", "true"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.dns_resolvers.0.ipv4.1.route_through_private_network", "false"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.dns_resolvers.0.ipv4.2.route_through_private_network", "false"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.dns_resolvers.0.ipv6.0.route_through_private_network", "false"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.dns_resolvers.0.ipv6.1.route_through_private_network", "false"),

					resource.TestCheckResourceAttr(name, "rule_settings.0.untrusted_cert.0.action", "error"),
					resource.TestCheckResourceAttr(name, "rule_settings.0.payload_log.0.enabled", "true"),
				),
			},
		},
	})
}

func testAccCloudflareTeamsRuleConfigCustomResolver(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  name = "%[1]s"
  account_id = "%[2]s"
  description = "desc"
  precedence = 12302
  action = "resolve"
  filters = ["dns_resolver"]
  traffic = "any(dns.domains[*] == \"example.com\")"
  rule_settings {
    insecure_disable_dnssec_validation = false
	untrusted_cert {
		action = "error"
	}
	payload_log {
		enabled = true
	}
	dns_resolvers {
		ipv4 {
				ip = "1.34.4.4" 
				port = 53
				vnet_id = "5fb4dfe1-4fe7-4ff4-980e-72f89dd9af9e"
				route_through_private_network = true
		}
		ipv4 { 
				ip = "1.34.4.5"
				port = 30000
		}
		ipv4 { 
				ip = "1.34.4.6"
				port = 200
		}


		ipv6 { 
				ip = "2a09:bac1:76c0:1378::b:248"
				port = 53
		}
		ipv6 {
			ip = "2a09:bac1:76c0:1378::b:148"
			port = 53
		}
		

	}
  }
}
`, rnd, accountID)
}

func TestAccCloudflareTeamsRule_WithClipboardRedirection(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_gateway_policy.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareTeamsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsRuleConfigWithClipboardRedirection(rnd, accountID),
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
					resource.TestCheckResourceAttr(name, "rule_settings.0.biso_admin_controls.0.disable_clipboard_redirection", "true"), // Check new parameter
				),
			},
		},
	})
}

func testAccCloudflareTeamsRuleConfigWithClipboardRedirection(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
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
    biso_admin_controls {
      disable_printing = false
      disable_copy_paste = false
      disable_download = false
      disable_upload = false
      disable_keyboard = false
      disable_clipboard_redirection = true // Set new parameter
    }
  }
}
`, rnd, accountID)
}
