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

// TODO parallel tests run into rate limiting, update after client limiting is merged

func TestAccCloudFlarePageRule_Basic(t *testing.T) {
	var pageRule cloudflare.PageRule
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	target := fmt.Sprintf("test-basic.%s", zone)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlarePageRuleConfigBasic(zone, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlarePageRuleExists("cloudflare_page_rule.test", &pageRule),
					//testAccCheckCloudFlarePageRuleAttributes(&pageRule),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "zone", zone),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "target", fmt.Sprintf("%s/", target)),
				),
			},
		},
	})
}

func TestAccCloudFlarePageRule_FullySpecified(t *testing.T) {
	var pageRule cloudflare.PageRule
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	target := fmt.Sprintf("test-fully-specified.%s", zone)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlarePageRuleConfigFullySpecified(zone, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlarePageRuleExists("cloudflare_page_rule.test", &pageRule),
					//testAccCheckCloudFlarePageRuleAttributes(&pageRule),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "zone", zone),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "target", fmt.Sprintf("%s/", target)),
				),
			},
		},
	})
}

func TestAccCloudFlarePageRule_ForwardingOnly(t *testing.T) {
	var pageRule cloudflare.PageRule
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	target := fmt.Sprintf("test-fwd-only.%s", zone)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlarePageRuleConfigForwardingOnly(zone, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlarePageRuleExists("cloudflare_page_rule.test", &pageRule),
					//testAccCheckCloudFlarePageRuleAttributes(&pageRule),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "zone", zone),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "target", fmt.Sprintf("%s/", target)),
				),
			},
		},
	})
}

func TestAccCloudFlarePageRule_ForwardingAndOthers(t *testing.T) {
	var pageRule cloudflare.PageRule
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	target := fmt.Sprintf("test-fwd-others.%s", zone)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlarePageRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckCloudFlarePageRuleConfigForwardingAndOthers(zone, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlarePageRuleExists("cloudflare_page_rule.test", &pageRule),
					//testAccCheckCloudFlarePageRuleAttributes(&pageRule),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "zone", zone),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "target", target),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "target", fmt.Sprintf("%s/", target)),
				),

				ExpectError: regexp.MustCompile("HTTP status 400"),
			},
		},
	})
}

func TestAccCloudFlarePageRule_Updated(t *testing.T) {
	var before, after cloudflare.PageRule
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	target := fmt.Sprintf("test-updated.%s", zone)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlarePageRuleConfigBasic(zone, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlarePageRuleExists("cloudflare_page_rule.test", &before),
					//testAccCheckCloudFlarePageRuleAttributes(&pageRule), // TODO check attributes
				),
			},
			{
				Config: testAccCheckCloudFlarePageRuleConfigNewValue(zone, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlarePageRuleExists("cloudflare_page_rule.test", &after),
					//testAccCheckCloudFlarePageRuleAttributesUpdated(&pageRule),
					testAccCheckCloudFlarePageRuleIDUnchanged(&before, &after),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "target", fmt.Sprintf("%s/updated", target)),
				),
			},
		},
	})
}

func TestAccCloudFlarePageRule_CreateAfterManualDestroy(t *testing.T) {
	var before, after cloudflare.PageRule
	var initialID string
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	target := fmt.Sprintf("test-updated.%s", zone)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlarePageRuleConfigBasic(zone, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlarePageRuleExists("cloudflare_page_rule.test", &before),
					testAccManuallyDeletePageRule("cloudflare_page_rule.test", &initialID),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCloudFlarePageRuleConfigNewValue(zone, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlarePageRuleExists("cloudflare_page_rule.test", &after),
					testAccCheckCloudFlarePageRuleRecreated(&before, &after),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "target", fmt.Sprintf("%s/updated", target)),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "actions.0.always_online", "false"),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "actions.0.automatic_https_rewrites", "true"),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "actions.0.rocket_loader", "automatic"),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "actions.0.ssl", "strict"),
				),
			},
		},
	})
}

func testAccCheckCloudFlarePageRuleRecreated(before, after *cloudflare.PageRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if before.ID == after.ID {
			return fmt.Errorf("Expected change of PageRule Ids, but both were %v", before.ID)
		}
		return nil
	}
}

func testAccCheckCloudFlarePageRuleIDUnchanged(before, after *cloudflare.PageRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if before.ID != after.ID {
			return fmt.Errorf("ID should not change suring in place update, but got change %s -> %s", before.ID, after.ID)
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

func testAccCheckCloudFlarePageRuleAttributesBasic(pageRule *cloudflare.PageRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		// check the api only has attributes we set non-empty values for

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

func testAccCheckCloudFlarePageRuleAttributesFullySpecified(pageRule *cloudflare.PageRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		// TODO check each of the different types of (non-empty) values are set at API level

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

		// TODO check some updated attributes get set

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

func testAccManuallyDeletePageRule(name string, initialID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		*initialID = rs.Primary.ID
		err := client.DeletePageRule(rs.Primary.Attributes["zone_id"], rs.Primary.ID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckCloudFlarePageRuleConfigBasic(zone, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "test" {
	zone = "%s"
	target = "%s"
	actions = {
		always_online = true
		ssl = "flexible"
	}
}`, zone, target)
}

func testAccCheckCloudFlarePageRuleConfigNewValue(zone, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "test" {
	zone = "%s"
	target = "%s/updated"
	actions = {
		always_online = false
		automatic_https_rewrites = true
		ssl = "strict"
		rocket_loader = "automatic"
	}
}`, zone, target)
}

func testAccCheckCloudFlarePageRuleConfigFullySpecified(zone, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "test" {
	zone = "%s"
	target = "%s"
	actions = {
		always_online = true
		automatic_https_rewrites = true
		browser_check = true
		email_obfuscation = true
		ip_geolocation = true
		opportunistic_encryption = true
		server_side_exclude = true
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
}`, zone, target)
}

func testAccCheckCloudFlarePageRuleConfigForwardingOnly(zone, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "test" {
	zone = "%s"
	target = "%s"
	actions = {
		always_online = false
		automatic_https_rewrites = false
		browser_check = false
		email_obfuscation = false
		ip_geolocation = false
		opportunistic_encryption = false
		server_side_exclude = false
        always_use_https = false
        disable_apps = false
        disable_performance = false
        disable_security = false
		forwarding_url {
        	url = "http://%[1]s/forward"
			status_code = 301
		}
	}
}`, zone, target)
}

func testAccCheckCloudFlarePageRuleConfigForwardingAndOthers(zone, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "test" {
	zone = "%s"
	target = "%s"
	actions = {
        disable_security = true
		forwarding_url {
        	url = "http://%[1]s/forward"
			status_code = 301
		}
	}
}`, zone, target)
}
