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

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_web_analytics_rule", &resource.Sweeper{
		Name: "cloudflare_web_analytics_rule",
		F:    testSweepCloudflareWebAnalyticsRules,
	})
}

func testSweepCloudflareWebAnalyticsRules(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		tflog.Info(ctx, "Skipping web analytics rules sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	// List all web analytics sites to find rulesets
	sites, _, err := client.ListWebAnalyticsSites(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListWebAnalyticsSitesParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch web analytics sites: %s", err))
		return fmt.Errorf("failed to fetch web analytics sites: %w", err)
	}

	if len(sites) == 0 {
		tflog.Info(ctx, "No web analytics sites to check for rules")
		return nil
	}

	for _, site := range sites {
		// List rules for each ruleset
		rulesResp, err := client.ListWebAnalyticsRules(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListWebAnalyticsRulesParams{
			RulesetID: site.Ruleset.ID,
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to fetch web analytics rules for ruleset %s: %s", site.Ruleset.ID, err))
			continue
		}

		if len(rulesResp.Rules) == 0 {
			continue
		}

		for _, rule := range rulesResp.Rules {
			// Use standard filtering helper on the host field
			if !utils.ShouldSweepResource(rule.Host) {
				continue
			}

			tflog.Info(ctx, fmt.Sprintf("Deleting web analytics rule: %s (host: %s, account: %s)", rule.ID, rule.Host, accountID))
			_, err := client.DeleteWebAnalyticsRule(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.DeleteWebAnalyticsRuleParams{
				RulesetID: site.Ruleset.ID,
				RuleID:    rule.ID,
			})
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete web analytics rule %s: %s", rule.ID, err))
				continue
			}
			tflog.Info(ctx, fmt.Sprintf("Deleted web analytics rule: %s", rule.ID))
		}
	}

	return nil
}

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
