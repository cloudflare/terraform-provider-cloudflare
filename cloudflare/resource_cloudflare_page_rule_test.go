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

func TestAccCloudflarePageRule_UpdatingZoneForcesNewResource(t *testing.T) {
	var before, after cloudflare.PageRule
	oldZone := os.Getenv("CLOUDFLARE_DOMAIN")
	newZone := os.Getenv("CLOUDFLARE_ALT_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAltDomain(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigBasic(oldZone, fmt.Sprintf("test-updating-zone-value.%s", oldZone)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists("cloudflare_page_rule.test", &before),
					resource.TestCheckResourceAttr("cloudflare_page_rule.test", "zone", oldZone),
				),
			},
			{
				Config: testAccCheckCloudflarePageRuleConfigBasic(newZone, fmt.Sprintf("test-updating-zone-value.%s", newZone)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists("cloudflare_page_rule.test", &after),
					testAccCheckCloudflarePageRuleRecreated(&before, &after),
					resource.TestCheckResourceAttr("cloudflare_page_rule.test", "zone", newZone),
				),
			},
		},
	})
}

func TestAccCloudflarePageRuleMinifyAction(t *testing.T) {
	var pageRule cloudflare.PageRule
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	target := fmt.Sprintf("test-action-minify.%s", zone)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigMinify(zone, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists("cloudflare_page_rule.test", &pageRule),
					resource.TestCheckResourceAttr("cloudflare_page_rule.test", "actions.0.minify.0.css", "on"),
					resource.TestCheckResourceAttr("cloudflare_page_rule.test", "actions.0.minify.0.js", "off"),
					resource.TestCheckResourceAttr("cloudflare_page_rule.test", "actions.0.minify.0.html", "on"),
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

func TestAccCheckCloudflarePageRuleRecreated(before, after *cloudflare.PageRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if before.ID == after.ID {
			return fmt.Errorf("Expected change of PageRule Ids, but both were %v", before.ID)
		}
		return nil
	}
}

func TestAccCheckCloudflarePageRuleIDUnchanged(before, after *cloudflare.PageRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if before.ID != after.ID {
			return fmt.Errorf("ID should not change suring in place update, but got change %s -> %s", before.ID, after.ID)
		}
		return nil
	}
}

func TestAccCheckCloudflarePageRuleDestroy(s *terraform.State) error {
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

func TestAccCheckCloudflarePageRuleAttributesBasic(pageRule *cloudflare.PageRule) resource.TestCheckFunc {
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

		if val, ok := actionMap["ssl"]; ok {
			if _, ok := val.(string); !ok || val != "flexible" {
				return fmt.Errorf("'ssl' not specified correctly at api, found: %q", val)
			}
		} else {
			return fmt.Errorf("'ssl' not specified at api")
		}

		if len(pageRule.Actions) != 2 {
			return fmt.Errorf("api should only have attributes we set non-empty (%d) but got %d: %#v",
				2, len(pageRule.Actions), pageRule.Actions)
		}

		return nil
	}
}

func TestAccCheckCloudflarePageRuleAttributesUpdated(pageRule *cloudflare.PageRule) resource.TestCheckFunc {
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

func TestAccCheckCloudflarePageRuleExists(n string, pageRule *cloudflare.PageRule) resource.TestCheckFunc {
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

func TestAccManuallyDeletePageRule(name string, initialID *string) resource.TestCheckFunc {
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

func TestAccCheckCloudflarePageRuleConfigMinify(zone, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "test" {
	zone = "%s"
	target = "%s"
	actions {
		minify {
			js = "off"
			css = "on"
			html = "on"
		}
	}
}`, zone, target)
}

func TestAccCheckCloudflarePageRuleConfigBasic(zone, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "test" {
	zone = "%s"
	target = "%s"
	actions {
		always_online = "on"
		ssl = "flexible"
	}
}`, zone, target)
}

func TestAccCheckCloudflarePageRuleConfigNewValue(zone, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "test" {
	zone = "%s"
	target = "%s/updated"
	actions {
		always_online = "off"
		browser_check = "on"
		ssl = "strict"
		rocket_loader = "on"
	}
}`, zone, target)
}

func TestAccCheckCloudflarePageRuleConfigFullySpecified(zone, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "test" {
	zone = "%s"
	target = "%s"
	actions {
		always_online = "on"
		browser_check = "on"
		email_obfuscation = "on"
		ip_geolocation = "on"
		server_side_exclude = "on"
		disable_apps = true
		disable_performance = true
		disable_security = true
		cache_level = "bypass"
		security_level = "essentially_off"
		ssl = "flexible"
	}
}`, zone, target)
}

func TestAccCheckCloudflarePageRuleConfigForwardingOnly(zone, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "test" {
	zone = "%s"
	target = "%s"
	actions {
		// on/off options cannot even be set to off without causing error
		forwarding_url {
			url = "http://%[1]s/forward"
			status_code = 301
		}
	}
}`, zone, target)
}

func TestAccCheckCloudflarePageRuleConfigForwardingAndOthers(zone, target string) string {
	return fmt.Sprintf(`
resource "cloudflare_page_rule" "test" {
	zone = "%s"
	target = "%s"
	actions {
		disable_security = true
		forwarding_url {
			url = "http://%[1]s/forward"
			status_code = 301
		}
	}
}`, zone, target)
}
