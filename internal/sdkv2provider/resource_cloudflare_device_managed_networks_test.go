package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareDeviceManagedNetworks(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_device_managed_networks.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDeviceManagedNetworks(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "tls"),
					resource.TestCheckResourceAttr(name, "config.0.tls_sockaddr", "foobar:1234"),
					resource.TestCheckResourceAttr(name, "config.0.sha256", "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"),
				),
			},
		},
	})
}

func testAccCloudflareDeviceManagedNetworks(accountID, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_device_managed_networks" "%[1]s" {
  account_id                = "%[2]s"
  name                      = "%[1]s"
  type                      = "tls"
  config {
	tls_sockaddr = "foobar:1234"
	sha256 = "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"
  }
}
`, rnd, accountID)
}
