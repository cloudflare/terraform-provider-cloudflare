package email_routing_dns_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testEmailRoutingDNSConfig(resourceID, zoneID string, subDomain string) string {
	return acctest.LoadTestCase("emailroutingdnsconfig.tf", resourceID, zoneID, subDomain)
}

func TestAccTestEmailRoutingDNS(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_email_routing_dns." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	subDomain := fmt.Sprintf("%s.%s", rnd, domain)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testEmailRoutingDNSConfig(rnd, zoneID, subDomain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
				),
			},
		},
	})
}
