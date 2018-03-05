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
					testAccCheckCloudFlarePageRuleAttributesBasic(&pageRule),
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
					testAccCheckCloudFlarePageRuleAttributesFullySpecified(&pageRule),
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
			{
				Config: testAccCheckCloudFlarePageRuleConfigForwardingAndOthers(zone, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlarePageRuleExists("cloudflare_page_rule.test", &pageRule),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "zone", zone),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "target", target),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "target", fmt.Sprintf("%s/", target)),
				),

				ExpectError: regexp.MustCompile("\"forwarding_url\" cannot be set with any other actions"),
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
				),
			},
			{
				Config: testAccCheckCloudFlarePageRuleConfigNewValue(zone, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlarePageRuleExists("cloudflare_page_rule.test", &after),
					testAccCheckCloudFlarePageRuleAttributesUpdated(&after),
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
						"cloudflare_page_rule.test", "actions.0.always_online", "off"),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "actions.0.browser_check", "on"),
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
		// this covers on/off attribute types and setting enum-type strings

		actionMap := pageRuleActionsToMap(pageRule.Actions)

		if val, ok := actionMap["always_online"]; ok {
			if _, ok := val.(string); !ok || val != "on" { // lots of booleans get mapped to on/off at api
				return fmt.Errorf("'always_online' not specified correctly at api, found: '%v'", val)
			}
		} else {
			return fmt.Errorf("'always_online' not specified at api")
		}

		if val, ok := actionMap["disable_apps"]; ok {
			if val != nil {
				return fmt.Errorf("'disable_apps' is a unitary value, expect nil value at api, but found: '%v'", val)
			}
		} else {
			return fmt.Errorf("'disable_apps' not specified at api")
		}

		if val, ok := actionMap["ssl"]; ok {
			if _, ok := val.(string); !ok || val != "flexible" {
				return fmt.Errorf("'ssl' not specified correctly at api, found: %q", val)
			}
		} else {
			return fmt.Errorf("'ssl' not specified at api")
		}

		if len(pageRule.Actions) != 3 {
			return fmt.Errorf("api should only have attributes we set non-empty (%d) but got %d: %#v",
				3, len(pageRule.Actions), pageRule.Actions)
		}

		return nil
	}
}

func testAccCheckCloudFlarePageRuleAttributesFullySpecified(pageRule *cloudflare.PageRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		// check boolean variables get set correctly
		actionMap := pageRuleActionsToMap(pageRule.Actions)

		if val, ok := actionMap["disable_apps"]; ok {
			if val != nil {
				return fmt.Errorf("'disable_apps' is a unitary value, expect nil value at api, but found: '%v'", val)
			}
		} else {
			return fmt.Errorf("'disable_apps' not specified at api")
		}

		if val, ok := actionMap["browser_cache_ttl"]; ok {
			if _, ok := val.(float64); !ok || val != 10000.000000 {
				return fmt.Errorf("'browser_cache_ttl' not specified correctly at api, found: '%f'", val.(float64))
			}
		} else {
			return fmt.Errorf("'browser_cache_ttl' not specified at api")
		}

		if len(pageRule.Actions) != 13 {
			return fmt.Errorf("api should return the attributes we set non-empty (count: %d) but got %d: %#v",
				13, len(pageRule.Actions), pageRule.Actions)
		}

		return nil
	}
}

func testAccCheckCloudFlarePageRuleAttributesUpdated(pageRule *cloudflare.PageRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		actionMap := pageRuleActionsToMap(pageRule.Actions)

		if _, ok := actionMap["disable_apps"]; ok {
			return fmt.Errorf("'disable_apps' found at api, but we should have removed it")
		}

		if val, ok := actionMap["always_online"]; ok {
			if _, ok := val.(string); !ok || val != "off" { // lots of booleans get mapped to on/off at api
				return fmt.Errorf("'always_online' not specified correctly at api, found: '%v'", val)
			}
		} else {
			return fmt.Errorf("'always_online' not specified at api")
		}

		if val, ok := actionMap["browser_check"]; ok {
			if _, ok := val.(string); !ok || val != "on" { // lots of booleans get mapped to on/off at api
				return fmt.Errorf("'browser_check' not specified correctly at api, found: '%v'", val)
			}
		} else {
			return fmt.Errorf("'browser_check' not specified at api")
		}

		if val, ok := actionMap["ssl"]; ok {
			if _, ok := val.(string); !ok || val != "strict" {
				return fmt.Errorf("'ssl' not specified correctly at api, found: %q", val)
			}
		} else {
			return fmt.Errorf("'ssl' not specified at api")
		}

		if val, ok := actionMap["rocket_loader"]; ok {
			if _, ok := val.(string); !ok || val != "automatic" {
				return fmt.Errorf("'rocket_loader' not specified correctly at api, found: %q", val)
			}
		} else {
			return fmt.Errorf("'rocket_loader' not specified at api")
		}

		if len(pageRule.Actions) != 4 {
			return fmt.Errorf("api should only have attributes we set non-empty (%d) but got %d: %#v",
				4, len(pageRule.Actions), pageRule.Actions)
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
		always_online = "on"
		ssl = "flexible"
 		disable_apps = true
	}
}`, zone, target)
}

func testAccCheckCloudFlarePageRuleConfigNewValue(zone, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "test" {
	zone = "%s"
	target = "%s/updated"
	actions = {
		always_online = "off"
		browser_check = "on"
		disable_apps = false
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
		always_online = "on"
		browser_check = "on"
		email_obfuscation = "on"
		ip_geolocation = "on"
		server_side_exclude = "on"
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
		// on/off options cannot even be set to off without causing error
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
