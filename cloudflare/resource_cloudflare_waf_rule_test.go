package cloudflare

import (
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccCloudflareWAFRule_CreateThenUpdate(t *testing.T) {
	t.Parallel()
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	ruleID := "100000"
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_waf_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareWAFRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWAFRuleConfig(zone, ruleID, "simulate", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "rule_id", ruleID),
					resource.TestCheckResourceAttr(name, "zone", zone),
					resource.TestCheckResourceAttrSet(name, "zone_id"),
					resource.TestCheckResourceAttrSet(name, "package_id"),
					resource.TestCheckResourceAttr(name, "mode", "simulate"),
				),
			},
			{
				Config: testAccCheckCloudflareWAFRuleConfig(zone, ruleID, "challenge", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "rule_id", ruleID),
					resource.TestCheckResourceAttr(name, "zone", zone),
					resource.TestCheckResourceAttrSet(name, "zone_id"),
					resource.TestCheckResourceAttrSet(name, "package_id"),
					resource.TestCheckResourceAttr(name, "mode", "challenge"),
				),
			},
		},
	})
}

func testAccCheckCloudflareWAFRuleDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_waf_rule" {
			continue
		}

		rule, err := client.WAFRule(rs.Primary.Attributes["zone_id"], rs.Primary.Attributes["package_id"], rs.Primary.ID)
		if err != nil {
			return err
		}

		if rule.Mode != "default" {
			return fmt.Errorf("Expected mode to be reset to default, got: %s", rule.Mode)
		}
	}

	return nil
}

func testAccCheckCloudflareWAFRuleConfig(zone, ruleID, mode, name string) string {
	return fmt.Sprintf(`
				resource "cloudflare_waf_rule" "%[4]s" {
					rule_id = %[2]s
					zone = "%[1]s"
					mode = "%[3]s"
				}`, zone, ruleID, mode, name)
}
