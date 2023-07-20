package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareDevices(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_devices.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicesConfig(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCloudflareDevicesDataSourceId(name),
				),
			},
		},
	})
}

func testAccCloudflareDevicesDataSourceId(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		all := s.RootModule().Resources
		rs, ok := all[n]

		if !ok {
			return fmt.Errorf("can't find Devices data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Snapshot Devices source ID not set")
		}
		return nil
	}
}

func testAccCloudflareDevicesConfig(name string, accountID string) string {
	return fmt.Sprintf(`
	data "cloudflare_devices" "%[1]s" {
		account_id = "%[2]s"
	}`, name, accountID)
}
