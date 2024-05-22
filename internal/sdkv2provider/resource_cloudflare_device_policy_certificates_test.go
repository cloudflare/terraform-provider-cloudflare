package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareDevicePolicyCertificatesCreate(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_device_policy_certificates.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareDevicePostureIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareDevicePolicyCertificates(rnd, zoneID, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
				),
			},
		},
	})
}

func testCloudflareDevicePolicyCertificates(rnd, zoneID string, enable bool) string {
	return fmt.Sprintf(`
resource "cloudflare_device_policy_certificates" "%[1]s" {
	zone_id = "%[2]s"
	enabled = "%[3]t"
}
`, rnd, zoneID, enable)
}
