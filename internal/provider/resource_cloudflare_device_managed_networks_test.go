package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
				Config: testAccCloudflareDeviceManagedNetworks(accountID, rnd, "custom profile"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "custom"),
					resource.TestCheckResourceAttr(name, "entry.0.name", fmt.Sprintf("%s_entry1", rnd)),
					resource.TestCheckResourceAttr(name, "entry.0.enabled", "true"),
					resource.TestCheckResourceAttr(name, "entry.0.pattern.0.regex", "^4[0-9]"),
					resource.TestCheckResourceAttr(name, "entry.0.pattern.0.validation", "luhn"),
				),
			},
		},
	})
}

func testAccCloudflareDeviceManagedNetworks(accountID, rnd, description string) string {
	return fmt.Sprintf(`
resource "cloudflare_device_managed_networks" "%[1]s" {
  account_id                = "%[3]s"
  name                      = "%[1]s"
  type                      = "tls"
  config {
	tls_sockaddr = "foobar:1234"
	sha256 = b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c
  }
}
`, rnd, description, accountID)
}
