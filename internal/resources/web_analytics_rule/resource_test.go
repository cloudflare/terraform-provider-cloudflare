package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareWebAnalyticsRule_Create(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	ruleName := fmt.Sprintf("cloudflare_web_analytics_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareWebAnalyticsRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareWebAnalyticsRule(rnd, accountID, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(ruleName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttrSet(ruleName, "ruleset_id"),
					resource.TestCheckResourceAttr(ruleName, "host", zoneID),
				),
			},
		},
	})
}

func testAccCheckCloudflareWebAnalyticsRuleDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_web_analytics_rule" {
			continue
		}

		rules, err := client.ListWebAnalyticsRules(context.Background(), cloudflare.AccountIdentifier(rs.Primary.Attributes[consts.AccountIDSchemaKey]), cloudflare.ListWebAnalyticsRulesParams{
			RulesetID: rs.Primary.Attributes["ruleset_id"],
		})

		if err == nil {
			for _, rule := range rules.Rules {
				if rule.ID == rs.Primary.Attributes["id"] {
					return fmt.Errorf("web analytics rule still exists")
				}
			}
		}
	}

	return nil
}

func testAccCloudflareWebAnalyticsRule(resourceName, accountID, zoneId string) string {
	return fmt.Sprintf(`
resource "cloudflare_web_analytics_site" "%[1]s" {
  account_id   = "%[2]s"
  zone_tag     = "%[3]s"
  auto_install = true
}

resource "cloudflare_web_analytics_rule" "%[1]s" {
  depends_on = [cloudflare_web_analytics_site.%[1]s]
  account_id = "%[2]s"
  ruleset_id = cloudflare_web_analytics_site.%[1]s.ruleset_id
  host       = "%[3]s"
  paths      = ["/excluded"]
  inclusive  = false
  is_paused  = false
}
`, resourceName, accountID, zoneId)
}
