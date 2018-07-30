package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"reflect"
	"regexp"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

// TODO parallel tests run into rate limiting, update after client limiting is merged

func TestAccCloudflarePageRule_Basic(t *testing.T) {
	var pageRule cloudflare.PageRule
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	target := fmt.Sprintf("test-basic.%s", zone)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigBasic(zone, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists("cloudflare_page_rule.test", &pageRule),
					testAccCheckCloudflarePageRuleAttributesBasic(&pageRule),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "zone", zone),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "target", fmt.Sprintf("%s/", target)),
				),
			},
		},
	})
}

func TestAccCloudflarePageRule_FullySpecified(t *testing.T) {
	var pageRule cloudflare.PageRule
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	target := fmt.Sprintf("test-fully-specified.%s", zone)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigFullySpecified(zone, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists("cloudflare_page_rule.test", &pageRule),
					testAccCheckCloudflarePageRuleAttributesFullySpecified(&pageRule),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "zone", zone),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "target", fmt.Sprintf("%s/", target)),
				),
			},
		},
	})
}

func TestAccCloudflarePageRule_ForwardingOnly(t *testing.T) {
	var pageRule cloudflare.PageRule
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	target := fmt.Sprintf("test-fwd-only.%s", zone)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigForwardingOnly(zone, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists("cloudflare_page_rule.test", &pageRule),
					//testAccCheckCloudflarePageRuleAttributes(&pageRule),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "zone", zone),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "target", fmt.Sprintf("%s/", target)),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test",
						"actions.0.forwarding_url.0.url",
						fmt.Sprintf("http://%s/forward", zone),
					),
				),
			},
		},
	})
}

func TestAccCloudflarePageRule_ForwardingAndOthers(t *testing.T) {
	var pageRule cloudflare.PageRule
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	target := fmt.Sprintf("test-fwd-others.%s", zone)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigForwardingAndOthers(zone, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists("cloudflare_page_rule.test", &pageRule),
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

func TestAccCloudflarePageRule_Updated(t *testing.T) {
	var before, after cloudflare.PageRule
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	target := fmt.Sprintf("test-updated.%s", zone)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigBasic(zone, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists("cloudflare_page_rule.test", &before),
				),
			},
			{
				Config: testAccCheckCloudflarePageRuleConfigNewValue(zone, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists("cloudflare_page_rule.test", &after),
					testAccCheckCloudflarePageRuleAttributesUpdated(&after),
					testAccCheckCloudflarePageRuleIDUnchanged(&before, &after),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "target", fmt.Sprintf("%s/updated", target)),
				),
			},
		},
	})
}

func TestAccCloudflarePageRule_CreateAfterManualDestroy(t *testing.T) {
	var before, after cloudflare.PageRule
	var initialID string
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	target := fmt.Sprintf("test-updated.%s", zone)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigBasic(zone, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists("cloudflare_page_rule.test", &before),
					testAccManuallyDeletePageRule("cloudflare_page_rule.test", &initialID),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCloudflarePageRuleConfigNewValue(zone, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists("cloudflare_page_rule.test", &after),
					testAccCheckCloudflarePageRuleRecreated(&before, &after),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "target", fmt.Sprintf("%s/updated", target)),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "actions.0.always_online", "off"),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "actions.0.browser_check", "on"),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "actions.0.rocket_loader", "on"),
					resource.TestCheckResourceAttr(
						"cloudflare_page_rule.test", "actions.0.ssl", "strict"),
				),
			},
		},
	})
}

func TestTranformForwardingURL(t *testing.T) {
	key, val, err := transformFromCloudflarePageRuleAction(&cloudflare.PageRuleAction{
		ID: "forwarding_url",
		Value: map[string]interface{}{
			"url":         "http://test.com/forward",
			"status_code": 302,
		},
	})
	if err != nil {
		t.Fatalf("Unexpected error transforming page rule action: %s", err)
	}

	if key != "forwarding_url" {
		t.Fatalf("Unexpected key transforming page rule action. Expected \"forwarding_url\", got \"%s\"", key)
	}

	// the transformed value for a forwarding_url should be [{url: "", "status_code": 302}] (single item slice where the
	// element in the slice is a map)
	if sl, isSlice := val.([]interface{}); !isSlice {
		t.Fatalf("Unexpected value type from transforming page rule action. Expected slice, got %s", reflect.TypeOf(val).Kind())
	} else if len(sl) != 1 {
		t.Fatalf("Unexpected slice length after transforming page rule action. Expected 1, got %d", len(sl))
	} else if _, isMap := sl[0].(map[string]interface{}); !isMap {
		t.Fatalf("Unexpected type in slice after tranforming page rule action. Expected map[string]interface{}, got %s", reflect.TypeOf(sl[0]).Kind())
	}
}

func testAccCheckCloudflarePageRuleRecreated(before, after *cloudflare.PageRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if before.ID == after.ID {
			return fmt.Errorf("Expected change of PageRule Ids, but both were %v", before.ID)
		}
		return nil
	}
}

func testAccCheckCloudflarePageRuleIDUnchanged(before, after *cloudflare.PageRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if before.ID != after.ID {
			return fmt.Errorf("ID should not change suring in place update, but got change %s -> %s", before.ID, after.ID)
		}
		return nil
	}
}

func testAccCheckCloudflarePageRuleDestroy(s *terraform.State) error {
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

func testAccCheckCloudflarePageRuleAttributesBasic(pageRule *cloudflare.PageRule) resource.TestCheckFunc {
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

func testAccCheckCloudflarePageRuleAttributesFullySpecified(pageRule *cloudflare.PageRule) resource.TestCheckFunc {
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

func testAccCheckCloudflarePageRuleAttributesUpdated(pageRule *cloudflare.PageRule) resource.TestCheckFunc {
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
			if _, ok := val.(string); !ok || val != "on" {
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

func testAccCheckCloudflarePageRuleExists(n string, pageRule *cloudflare.PageRule) resource.TestCheckFunc {
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

func testAccCheckCloudflarePageRuleConfigBasic(zone, target string) string {
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

func testAccCheckCloudflarePageRuleConfigNewValue(zone, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "test" {
	zone = "%s"
	target = "%s/updated"
	actions = {
		always_online = "off"
		browser_check = "on"
		disable_apps = false
		ssl = "strict"
		rocket_loader = "on"
	}
}`, zone, target)
}

func testAccCheckCloudflarePageRuleConfigFullySpecified(zone, target string) string {
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

func testAccCheckCloudflarePageRuleConfigForwardingOnly(zone, target string) string {
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

func testAccCheckCloudflarePageRuleConfigForwardingAndOthers(zone, target string) string {
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
