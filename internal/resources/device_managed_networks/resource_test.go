package device_managed_networks_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stainless-sdks/cloudflare-terraform/internal/acctest"
	"github.com/stainless-sdks/cloudflare-terraform/internal/consts"
	"github.com/stainless-sdks/cloudflare-terraform/internal/utils"
)

func TestAccCloudflareDeviceManagedNetworks(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_device_managed_networks.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
