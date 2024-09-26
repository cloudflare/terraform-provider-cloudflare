package infrastructure_access_target_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareInfrastructureAccessTarget_DataSource(t *testing.T) {
	rnd1 := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareInfrastructureTargetsMatchNoIpv6(rnd1),
				Check: resource.ComposeTestCheckFunc(
					// We should expect this data source to have 1 resource.
					resource.TestCheckResourceAttr("data.cloudflare_infrastructure_access_targets."+rnd1, "targets.#", "1"),
					// Check that there is no ipv6 object in this resource.
					resource.TestCheckNoResourceAttr("data.cloudflare_infrastructure_access_targets."+rnd1, "ip.ipv6"),
					// Check the existing attributes of this resource
					resource.TestCheckTypeSetElemNestedAttrs("cloudflare_infrastructure_access_targets."+rnd1, "resources.*", map[string]string{
						"hostname":                   rnd1,
						"ip.ipv4.ip_addr":            "250.26.29.250",
						"ip.ipv4.virtual_network_id": "b9c90134-52de-4903-81e8-004a3a06b435",
					}),
				),
			},
		},
	})
}

func testCloudflareInfrastructureTargetsMatchNoIpv6(hostname string) string {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	return fmt.Sprintf(`
resource "cloudflare_infrastructure_access_target" "%[2]s" {
	account_id = "%[1]s"
	hostname   = "%[2]s"
	ip = {
		ipv4 = {
           ip_addr = "250.26.29.250",
           virtual_network_id = "b9c90134-52de-4903-81e8-004a3a06b435"
        }
	}
}

data "cloudflare_infrastructure_access_targets" "%[2]s" {
	depends_on =  [cloudflare_infrastructure_access_target.%[2]s]
	account_id = "%[1]s"
}
`, accountID, hostname)
}
