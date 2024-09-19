package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareTarget_MatchHostname(t *testing.T) {
	rnd1 := generateRandomResourceName()
	rnd2 := generateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareInfrastructureTargetsMatchNoIpv6(rnd1),
				Check: resource.ComposeTestCheckFunc(
					// We should expect this data source to have 1 resource
					resource.TestCheckResourceAttr("data.cloudflare_zero_trust_infrastructure_targets."+rnd1, "resources.#", "1"),
					// Check that there is no ipv6 object in this resource
					resource.TestCheckNoResourceAttr("data.cloudflare_zero_trust_infrastructure_targets."+rnd1, "ip.ipv6"),
					// Check the existing attributes of this resource
					resource.TestCheckTypeSetElemNestedAttrs("data.cloudflare_zero_trust_infrastructure_targets."+rnd1, "resources.*", map[string]string{
						"hostname":                   rnd1,
						"ip.ipv4.ip_addr":            "187.26.29.249",
						"ip.ipv4.virtual_network_id": "c77b744e-acc8-428f-9257-6878c046ed55",
					}),
				),
			},
			{
				Config: testCloudflareInfrastructureTargetsMatchAll(rnd1, rnd2),
				Check: resource.ComposeTestCheckFunc(
					// Expect this data source to have 2 resources
					resource.TestCheckResourceAttr("data.cloudflare_zero_trust_infrastructure_targets.all", "resources.#", "2"),
					// Check the attributes of the first resource
					resource.TestCheckTypeSetElemNestedAttrs("data.cloudflare_zero_trust_infrastructure_targets.all", "resources.*", map[string]string{
						"hostname":                   rnd1,
						"ip.ipv4.ip_addr":            "187.26.29.249",
						"ip.ipv4.virtual_network_id": "c77b744e-acc8-428f-9257-6878c046ed55",
					}),
					// Check the attributes of the second resource
					resource.TestCheckTypeSetElemNestedAttrs("data.cloudflare_zero_trust_infrastructure_targets.all", "resources.*", map[string]string{
						"hostname":                   rnd2,
						"ip.ipv4.ip_addr":            "250.26.29.250",
						"ip.ipv6.ip_addr":            "64c0:64e8:f0b4:8dbf:7104:72b0:ec8f:f5e0",
						"ip.ipv6.virtual_network_id": "01920a8c-dc14-7bb2-b67b-14c858494a54",
					}),
				),
			},
		},
	})
}

func testCloudflareInfrastructureTargetsMatchNoIpv6(hostname string) string {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_infrastructure_target" "%[2]s" {
	account_id = "%[1]s"
	hostname   = "%[2]s"
	ip = {
		ipv4 = {
           ip_addr = "187.26.29.249",
           virtual_network_id = "c77b744e-acc8-428f-9257-6878c046ed55"
        }
	}
}

data "cloudflare_zero_trust_infrastructure_targets" "%[2]s" {
	depends_on =  [cloudflare_zero_trust_infrastructure_target.%[2]s]
}
`, accountID, hostname)
}

func testCloudflareInfrastructureTargetsMatchAll(hostname1 string, hostname2 string) string {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_infrastructure_target" "%[2]s" {
	account_id = "%[1]s"
	hostname   = "%[2]s"
	ip = {
		ipv4 = {
           ip_addr = "187.26.29.249",
           virtual_network_id = "c77b744e-acc8-428f-9257-6878c046ed55"
        }
	}
}

resource "cloudflare_zero_trust_infrastructure_target" "%[3]s" {
	account_id = "%[1]s"
	hostname   = "%[3]s"
	ip = {
		ipv4 = {
           ip_addr = "250.26.29.250",
           virtual_network_id = "01920a8c-dc14-7bb2-b67b-14c858494a54"
        },
		ipv6 = {
           ip_addr = "64c0:64e8:f0b4:8dbf:7104:72b0:ec8f:f5e0",
           virtual_network_id = "01920a8c-dc14-7bb2-b67b-14c858494a54"
        }
	}
}

data "cloudflare_zero_trust_infrastructure_targets" "all" {
	depends_on =  [
		cloudflare_zero_trust_infrastructure_target.%[2]s,
		cloudflare_zero_trust_infrastructure_target.%[3]s
	]
}
`, accountID, hostname1, hostname2)
}
