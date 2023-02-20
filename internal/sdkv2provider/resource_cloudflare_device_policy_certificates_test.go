package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareDevicePolicyCertificatesCreate(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

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
