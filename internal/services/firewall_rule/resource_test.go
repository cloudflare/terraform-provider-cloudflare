package firewall_rule_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}


func init() {
	resource.AddTestSweepers("cloudflare_firewall_rule", &resource.Sweeper{
		Name: "cloudflare_firewall_rule",
		F:    testSweepCloudflareFirewallRuleSweeper,
	})
}

func testSweepCloudflareFirewallRuleSweeper(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s",clientErr))
		return clientErr
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping firewall rules sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}
	rules, _, rulesErr := client.FirewallRules(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.FirewallRuleListParams{})

	if rulesErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare firewall rules: %s",rulesErr))
		return rulesErr
	}
	if len(rules) == 0 {
		tflog.Info(ctx, "No Cloudflare firewall rules to sweep")
		return nil
	}

	for _, rule := range rules {
		// Use standard filtering helper on the description field
		if !utils.ShouldSweepResource(rule.Description) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting firewall rule: %s (zone: %s)", rule.ID, zoneID))
		err := client.DeleteFirewallRule(ctx, cloudflare.ZoneIdentifier(zoneID), rule.ID)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete firewall rule %s: %s", rule.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted firewall rule: %s", rule.ID))
	}

	return nil
}

func TestAccFirewallRuleSimple(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_firewall_rule." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	filterQuoted := `(http.request.uri.path ~ \".*wp-login-` + rnd + `.php\" or http.request.uri.path ~ \".*xmlrpc.php\") and ip.src ne 192.0.2.1`

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testFirewallRuleConfig(rnd, zoneID, "true", "this is notes", filterQuoted, "allow", "1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "description", "this is notes"),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "action", "allow"),
					resource.TestCheckResourceAttr(name, "priority", "1"),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
				),
			},
		},
	})
}

func testFirewallRuleConfig(resourceID, zoneID, paused, description, expression, action, priority string) string {
	return acctest.LoadTestCase("firewallruleconfig.tf", resourceID, zoneID, paused, description, expression, action, priority)
}

func TestAccFirewallRuleBypass(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_firewall_rule." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	filterQuoted := `(http.host eq \"` + domain + `\")`

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testFirewallRuleBypassConfig(rnd, zoneID, "false", "this is notes", filterQuoted),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "description", "this is notes"),
					resource.TestCheckResourceAttr(name, "paused", "false"),
					resource.TestCheckResourceAttr(name, "action", "bypass"),
					resource.TestCheckResourceAttr(name, "priority", "2"),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "products.#", "2"),
				),
			},
		},
	})
}

func testFirewallRuleBypassConfig(resourceID, zoneID, paused, description, expression string) string {
	return acctest.LoadTestCase("firewallrulebypassconfig.tf", resourceID, zoneID, paused, description, expression)
}

func TestAccFirewallRuleWithUnicodeAndHTMLEntity(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_firewall_rule." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	expression := `(http.host eq \"` + domain + `\")`

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallRuleWithUnicodeAndHTMLEntityConfig(rnd, zoneID, "true", "this is a 'test'", expression, "allow", "1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "description", "this is a 'test'"),
					resource.TestCheckResourceAttr(name, "paused", "true"),
					resource.TestCheckResourceAttr(name, "action", "allow"),
					resource.TestCheckResourceAttr(name, "priority", "1"),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func testAccFirewallRuleWithUnicodeAndHTMLEntityConfig(resourceID, zoneID, paused, description, expression, action, priority string) string {
	return acctest.LoadTestCase("accfirewallrulewithunicodeandhtmlentityconfig.tf", resourceID, zoneID, paused, description, expression, action, priority)
}
