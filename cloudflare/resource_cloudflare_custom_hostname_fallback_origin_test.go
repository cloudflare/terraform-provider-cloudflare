package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccCloudflareCustomHostnameFallbackOrigin(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_custom_hostname_fallback_origin." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomHostnameFallbackOrigin(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "zone_id", zoneID),
					resource.TestCheckResourceAttr(resourceName, "origin", fmt.Sprintf("fallback-origin.%s.terraform.cfapi.net", rnd)),
				),
			},
		},
	})
}

func testAccCheckCloudflareCustomHostnameFallbackOrigin(zoneID, rnd, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_custom_hostname_fallback_origin" "%[2]s" {
  zone_id = "%[1]s"
  origin = "fallback-origin.%[2]s.terraform.cfapi.net"
}

resource "cloudflare_record" "%[2]s" {
  zone_id = "%[1]s"
  name    = "fallback-origin.%[2]s.terraform.cfapi.net"
  value   = "192.168.0.1"
  type    = "A"
  ttl     = 3600
}`, zoneID, rnd, domain)
}
