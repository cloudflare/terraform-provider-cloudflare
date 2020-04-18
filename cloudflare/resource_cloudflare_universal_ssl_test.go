package cloudflare

import (
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccCloudflareUniversalSSL(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_universal_ssl.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareUniversalSSLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareUniversalSSLConfig(zoneID, "on", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						name, "settings.0.status", "on"),
				),
			},
			{
				Config: testAccCheckCloudflareUniversalSSLConfig(zoneID, "off", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						name, "settings.0.status", "off"),
				),
			},
		},
	})
}

func testAccCheckCloudflareUniversalSSLDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_universal_ssl" {
			continue
		}

		ussl, err := client.UniversalSSLSettingDetails(rs.Primary.Attributes["zone_id"])
		if err != nil {
			return err
		}

		if ussl.Enabled != false {
			return fmt.Errorf("Expected USSL status to be reset to false, got: %t", ussl.Enabled)
		}
	}

	return nil
}

func testAccCheckCloudflareUniversalSSLConfig(zoneID, status, name string) string {
	return fmt.Sprintf(`
				resource "cloudflare_universal_ssl" "%[3]s" {
					zone_id = "%[1]s"
					settings {
						status = "%[2]s"
					}
				}`, zoneID, status, name)
}
