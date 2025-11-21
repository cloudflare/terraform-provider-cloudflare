package email_routing_catch_all_test

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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_email_routing_catch_all", &resource.Sweeper{
		Name: "cloudflare_email_routing_catch_all",
		F: func(region string) error {
			client := acctest.SharedClient()
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			ctx := context.Background()

			// Get the catch-all rule
			catchAll, err := client.EmailRouting.Rules.CatchAlls.Get(ctx, email_routing.RuleCatchAllGetParams{
				ZoneID: cloudflare.F(zoneID),
			})
			if err != nil {
				return fmt.Errorf("failed to fetch email routing catch-all: %w", err)
			}

			// Disable the catch-all rule if it's enabled
			if catchAll.Enabled {
				actionParams := make([]email_routing.CatchAllActionParam, 0)
				for _, action := range catchAll.Actions {
					actionParams = append(actionParams, email_routing.CatchAllActionParam{
						Type:  cloudflare.F(action.Type),
						Value: cloudflare.F(action.Value),
					})
				}
				matcherParams := make([]email_routing.CatchAllMatcherParam, 0)
				for _, matcher := range catchAll.Matchers {
					matcherParams = append(matcherParams, email_routing.CatchAllMatcherParam{
						Type: cloudflare.F(matcher.Type),
					})
				}
				_, err := client.EmailRouting.Rules.CatchAlls.Update(ctx, email_routing.RuleCatchAllUpdateParams{
					ZoneID:   cloudflare.F(zoneID),
					Actions:  cloudflare.F(actionParams),
					Matchers: cloudflare.F(matcherParams),
					Enabled:  cloudflare.F(email_routing.RuleCatchAllUpdateParamsEnabledFalse),
				})
				if err != nil {
					return fmt.Errorf("failed to disable email routing catch-all: %w", err)
				}
				fmt.Printf("Disabled email routing catch-all rule\n")
			} else {
				fmt.Printf("Email routing catch-all rule is already disabled\n")
			}

			return nil
		},
	})
}

func testEmailRoutingRuleCatchAllConfig(resourceID, zoneID string, enabled bool) string {
	return acctest.LoadTestCase("emailroutingrulecatchallconfig.tf", resourceID, zoneID, enabled)
}

func TestAccCloudflareEmailRoutingCatchAll(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_email_routing_catch_all." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testEmailRoutingRuleCatchAllConfig(rnd, zoneID, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", "terraform rule catch all"),

					resource.TestCheckResourceAttr(name, "matchers.0.type", "all"),

					resource.TestCheckResourceAttr(name, "actions.0.type", "forward"),
					resource.TestCheckResourceAttr(name, "actions.0.value.#", "1"),
					resource.TestCheckResourceAttr(name, "actions.0.value.0", "destinationaddress@example.net"),
				),
			},
		},
	})
}
