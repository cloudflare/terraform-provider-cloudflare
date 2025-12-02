package email_routing_rule_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/email_routing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_email_routing_rule", &resource.Sweeper{
		Name: "cloudflare_email_routing_rule",
		F: func(region string) error {
			ctx := context.Background()
			client := acctest.SharedClient()
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

			if zoneID == "" {
				tflog.Info(ctx, "Skipping email routing rules sweep: CLOUDFLARE_ZONE_ID not set")
				return nil
			}

			// List all email routing rules
			rules, err := client.EmailRouting.Rules.List(ctx, email_routing.RuleListParams{
				ZoneID: cloudflare.F(zoneID),
			})
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to fetch email routing rules: %s", err))
				return fmt.Errorf("failed to fetch email routing rules: %w", err)
			}

			ruleList := rules.Result
			if len(ruleList) == 0 {
				tflog.Info(ctx, "No email routing rules to sweep")
				return nil
			}

			tflog.Info(ctx, fmt.Sprintf("Found %d email routing rules", len(ruleList)))
			deletedCount := 0
			skippedCount := 0

			for _, rule := range ruleList {
				isCatchAll := false
				for _, matcher := range rule.Matchers {
					// you cannot delete a catch all rule
					if matcher.Type == "all" {
						isCatchAll = true
						break
					}
				}

				if isCatchAll {
					skippedCount++
					continue
				}

				if !utils.ShouldSweepResource(rule.Name) {
					continue
				}

				tflog.Info(ctx, fmt.Sprintf("Deleting email routing rule: %s (%s) (zone: %s)", rule.Name, rule.Tag, zoneID))
				_, err := client.EmailRouting.Rules.Delete(ctx, rule.Tag, email_routing.RuleDeleteParams{
					ZoneID: cloudflare.F(zoneID),
				})
				if err != nil {
					tflog.Error(ctx, fmt.Sprintf("Failed to delete email routing rule %s: %s", rule.Name, err))
					continue
				}
				tflog.Info(ctx, fmt.Sprintf("Deleted email routing rule: %s", rule.Tag))
				deletedCount++
			}

			tflog.Info(ctx, fmt.Sprintf("Deleted %d email routing rules, skipped %d catch-all rules", deletedCount, skippedCount))
			return nil
		},
	})
}

func TestAccCloudflareEmailRoutingRule_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_email_routing_rule." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testEmailRoutingRuleConfig(rnd, zoneID, true, 10),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "priority", "10"),
					resource.TestCheckResourceAttr(name, "name", "terraform rule"),

					resource.TestCheckResourceAttr(name, "matchers.0.type", "literal"),
					resource.TestCheckResourceAttr(name, "matchers.0.field", "to"),
					resource.TestCheckResourceAttr(name, "matchers.0.value", "test@example.com"),

					resource.TestCheckResourceAttr(name, "actions.0.type", "forward"),
					resource.TestCheckResourceAttr(name, "actions.0.value.#", "1"),
					resource.TestCheckResourceAttr(name, "actions.0.value.0", "destinationaddress@example.net"),
				),
			},
		},
	})
}

func TestAccCloudflareEmailRoutingRule_Drop(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_email_routing_rule." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testEmailRoutingRuleConfigDrop(rnd, zoneID, true, 10),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "priority", "10"),
					resource.TestCheckResourceAttr(name, "name", rnd),

					resource.TestCheckResourceAttr(name, "matchers.0.type", "literal"),
					resource.TestCheckResourceAttr(name, "matchers.0.field", "to"),
					resource.TestCheckResourceAttr(name, "matchers.0.value", "test@example.com"),

					resource.TestCheckResourceAttr(name, "actions.0.type", "drop"),
				),
			},
		},
	})
}

func testEmailRoutingRuleConfig(resourceID, zoneID string, enabled bool, priority int) string {
	return acctest.LoadTestCase("emailroutingruleconfig.tf", resourceID, zoneID, enabled, priority)
}

func testEmailRoutingRuleConfigDrop(resourceID, zoneID string, enabled bool, priority int) string {
	return acctest.LoadTestCase("emailroutingruleconfigdrop.tf", resourceID, zoneID, enabled, priority)
}
