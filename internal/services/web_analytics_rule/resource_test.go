package web_analytics_rule_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareWebAnalyticsRule_Create(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	ruleName := fmt.Sprintf("cloudflare_web_analytics_rule.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWebAnalyticsRuleDestroy,
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
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

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
	return acctest.LoadTestCase("webanalyticsrule.tf", resourceName, accountID, zoneId)
}
