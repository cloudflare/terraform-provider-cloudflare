package email_routing_catch_all_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

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
