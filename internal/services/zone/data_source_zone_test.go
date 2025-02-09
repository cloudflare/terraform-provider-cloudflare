package zone_test

import (
	"fmt"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareZone_NameLookup(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_zone.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneConfigBasic(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", "terraform.cfapi.net"),
					resource.TestCheckResourceAttr(name, consts.IDSchemaKey, acctest.TestAccCloudflareZoneID),
					resource.TestCheckResourceAttr(name, "status", "active"),
				),
			},
		},
	})
}

func testAccCloudflareZoneConfigBasic(rnd string) string {
	return fmt.Sprintf(`
data "cloudflare_zone" "%[1]s" {
  filter = {
  	name = "terraform.cfapi.net"
  }
}
`, rnd)
}
