package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"regexp"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccCloudFlarePageRule_Basic(t *testing.T) {
	var pageRule cloudflare.PageRule
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	target := fmt.Sprintf("test.%s", domain)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlarePageRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckCloudFlarePageRuleConfigBasic(domain, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlarePageRuleExists("cloudflare_page_rule.foobar", &pageRule),
					//testAccCheckCloudFlarePageRuleAttributes(&pageRule),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.foobar", "domain", domain),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.foobar", "target", fmt.Sprintf("test.%s", domain)),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.foobar", "effective_target", fmt.Sprintf("test.%s/", domain)),
				),
			},
		},
	})
}

func TestAccCloudFlarePageRule_FullySpecified(t *testing.T) {
	var pageRule cloudflare.PageRule
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	target := fmt.Sprintf("test.%s", domain)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlarePageRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckCloudFlarePageRuleConfigFullySpecified(domain, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlarePageRuleExists("cloudflare_page_rule.foobar", &pageRule),
					//testAccCheckCloudFlarePageRuleAttributes(&pageRule),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.foobar", "domain", domain),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.foobar", "target", fmt.Sprintf("test.%s", domain)),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.foobar", "effective_target", fmt.Sprintf("test.%s/", domain)),
				),
			},
		},
	})
}

func TestAccCloudFlarePageRule_ForwardingOnly(t *testing.T) {
	var pageRule cloudflare.PageRule
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	target := fmt.Sprintf("test.%s", domain)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlarePageRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckCloudFlarePageRuleConfigForwardingOnly(domain, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlarePageRuleExists("cloudflare_page_rule.foobar", &pageRule),
					//testAccCheckCloudFlarePageRuleAttributes(&pageRule),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.foobar", "domain", domain),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.foobar", "target", fmt.Sprintf("test.%s", domain)),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.foobar", "effective_target", fmt.Sprintf("test.%s/", domain)),
				),

				ExpectError: regexp.MustCompile("\\\"forwarding_url\\\" may not be used with \\\"any setting\\\""),
			},
		},
	})
}

func TestAccCloudFlarePageRule_Updated(t *testing.T) {
	var pageRule cloudflare.PageRule
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	target := fmt.Sprintf("test.%s", domain)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlarePageRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckCloudFlarePageRuleConfigBasic(domain, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlarePageRuleExists("cloudflare_page_rule.foobar", &pageRule),
					//testAccCheckCloudFlarePageRuleAttributes(&pageRule),
				),
			},
			resource.TestStep{
				Config: testAccCheckCloudFlarePageRuleConfigNewValue(domain, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlarePageRuleExists("cloudflare_page_rule.foobar", &pageRule),
					//testAccCheckCloudFlarePageRuleAttributesUpdated(&pageRule),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.foobar", "target", fmt.Sprintf("test.%s", domain)),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.foobar", "actions.0.always_online", "on"),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.foobar", "actions.0.ssl", "strict"),
				),
			},
		},
	})
}

func testAccCheckCloudFlarePageRuleRecreated(t *testing.T,
	before, after *cloudflare.PageRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if before.ID == after.ID {
			t.Fatalf("Expected change of PageRule Ids, but both were %v", before.ID)
		}
		return nil
	}
}

func testAccCheckCloudFlarePageRuleDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_page_rule" {
			continue
		}

		_, err := client.PageRule(rs.Primary.Attributes["zone_id"], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("PageRule still exists")
		}
	}

	return nil
}

func testAccCheckCloudFlarePageRuleAttributes(pageRule *cloudflare.PageRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if pageRule.Actions[0].ID != "always_online" {
			return fmt.Errorf("Bad type for actions[0]: %s", pageRule.Actions[0].ID)
		}

		if pageRule.Actions[0].Value != "on" {
			return fmt.Errorf("Bad value for actions.always_online: %s", pageRule.Actions[0].Value)
		}

		if pageRule.Actions[1].ID != "ssl" {
			return fmt.Errorf("Bad type for actions[0]: %s", pageRule.Actions[0].ID)
		}

		if pageRule.Actions[1].Value != "flexible" {
			return fmt.Errorf("Bad value for actions.ssl: %s", pageRule.Actions[0].Value)
		}

		return nil
	}
}

func testAccCheckCloudFlarePageRuleAttributesUpdated(pageRule *cloudflare.PageRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if pageRule.Actions[0].ID != "always_online" {
			return fmt.Errorf("Bad type for actions[0]: %s", pageRule.Actions[0].ID)
		}

		if pageRule.Actions[0].Value != "on" {
			return fmt.Errorf("Bad value for actions.always_online: %s", pageRule.Actions[0].Value)
		}

		if pageRule.Actions[0].ID != "ssl" {
			return fmt.Errorf("Bad type for actions[0]: %s", pageRule.Actions[0].ID)
		}

		if pageRule.Actions[0].Value != "strict" {
			return fmt.Errorf("Bad value for actions.ssl: %s", pageRule.Actions[0].Value)
		}

		return nil
	}
}

func testAccCheckCloudFlarePageRuleExists(n string, pageRule *cloudflare.PageRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No PageRule ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundPageRule, err := client.PageRule(rs.Primary.Attributes["zone_id"], rs.Primary.ID)
		if err != nil {
			return err
		}

		if foundPageRule.ID != rs.Primary.ID {
			return fmt.Errorf("PageRule not found")
		}

		*pageRule = foundPageRule

		return nil
	}
}

func testAccCheckCloudFlarePageRuleConfigBasic(domain, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "foobar" {
	domain = "%s"
	target = "%s"
	actions = {
		always_online = "on",
		ssl = "flexible",
	}
}`, domain, target)
}

func testAccCheckCloudFlarePageRuleConfigNewValue(domain, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "foobar" {
	domain = "%s"
	target = "%s"
	actions = {
		always_online = "on",
		ssl = "strict",
	}
}`, domain, target)
}

func testAccCheckCloudFlarePageRuleConfigFullySpecified(domain, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "foobar" {
	domain = "%s"
	target = "%s"
	actions = {
		always_online = "on"
		automatic_https_rewrites = "on"
		browser_check = "on"
		email_obfuscation = "on"
		ip_geolocation = "on"
		opportunistic_encryption = "on"
		server_side_exclude = "on"
        always_use_https = true
        disable_apps = true
        disable_performance = true
        disable_security = true
        browser_cache_ttl = 10000
        edge_cache_ttl = 10000
        cache_level = "bypass"
		security_level = "essentially_off"
		ssl = "flexible"
	}
}`, domain, target)
}

func testAccCheckCloudFlarePageRuleConfigForwardingOnly(domain, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "foobar" {
	domain = "%s"
	target = "%s"
	actions = {
		always_online = "off"
		automatic_https_rewrites = "off"
		browser_check = "off"
		email_obfuscation = "off"
		ip_geolocation = "off"
		opportunistic_encryption = "off"
		server_side_exclude = "off"
        always_use_https = false
        disable_apps = false
        disable_performance = false
        disable_security = false
		forwarding_url {
        	url = "http://%[1]s/forward"
			status_code = 301
		}
	}
}`, domain, target)
}
