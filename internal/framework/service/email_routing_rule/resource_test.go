package email_routing_rule_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_email_routing_rule", &resource.Sweeper{
		Name: "cloudflare_email_routing_rule",
		F: func(region string) error {
			client, err := acctest.SharedClient()
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

			if err != nil {
				return fmt.Errorf("error establishing client: %w", err)
			}

			ctx := context.Background()
			rules, _, err := client.ListEmailRoutingRules(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.ListEmailRoutingRulesParameters{})
			if err != nil {
				return fmt.Errorf("failed to fetch email routing rules: %w", err)
			}

			for _, rule := range rules {
				for _, matchers := range rule.Matchers {
					// you cannot delete a catch all rule
					if matchers.Type != "all" {
						_, err := client.DeleteEmailRoutingRule(ctx, cloudflare.ZoneIdentifier(zoneID), rule.Tag)
						if err != nil {
							return fmt.Errorf("failed to delete email routing rule %q: %w", rule.Name, err)
						}
					}
				}
			}

			return nil
		},
	})
}

func TestAccCloudflareEmailRoutingRule_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_email_routing_rule." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.ParallelTest(t, resource.TestCase{
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

					resource.TestCheckResourceAttr(name, "matcher.0.type", "literal"),
					resource.TestCheckResourceAttr(name, "matcher.0.field", "to"),
					resource.TestCheckResourceAttr(name, "matcher.0.value", "test@example.com"),

					resource.TestCheckResourceAttr(name, "action.0.type", "forward"),
					resource.TestCheckResourceAttr(name, "action.0.value.#", "1"),
					resource.TestCheckResourceAttr(name, "action.0.value.0", "destinationaddress@example.net"),
				),
			},
		},
	})
}

func TestAccCloudflareEmailRoutingRule_Drop(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_email_routing_rule." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.ParallelTest(t, resource.TestCase{
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

					resource.TestCheckResourceAttr(name, "matcher.0.type", "literal"),
					resource.TestCheckResourceAttr(name, "matcher.0.field", "to"),
					resource.TestCheckResourceAttr(name, "matcher.0.value", "test@example.com"),

					resource.TestCheckResourceAttr(name, "action.0.type", "drop"),
				),
			},
		},
	})
}

func testEmailRoutingRuleConfig(resourceID, zoneID string, enabled bool, priority int) string {
	return fmt.Sprintf(`
		resource "cloudflare_email_routing_rule" "%[1]s" {
		  zone_id = "%[2]s"
		  enabled = "%[3]t"
          priority = "%[4]d"
		  name = "terraform rule"
		  matcher {
			field  = "to"
			type = "literal"
			value = "test@example.com"
		  }

		  action {
			type = "forward"
			value = ["destinationaddress@example.net"]
		  }
	}
		`, resourceID, zoneID, enabled, priority)
}

func testEmailRoutingRuleConfigDrop(resourceID, zoneID string, enabled bool, priority int) string {
	return fmt.Sprintf(`
		resource "cloudflare_email_routing_rule" "%[1]s" {
		  zone_id = "%[2]s"
		  enabled = "%[3]t"
          priority = "%[4]d"
		  name = "%[1]s"
		  matcher {
			field  = "to"
			type = "literal"
			value = "test@example.com"
		  }

		  action {
			type = "drop"
		  }
	}
		`, resourceID, zoneID, enabled, priority)
}
