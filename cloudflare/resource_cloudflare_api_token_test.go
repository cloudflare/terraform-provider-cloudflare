package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAPIToken(t *testing.T) {
	rnd := generateRandomResourceName()
	resourceID := "cloudflare_api_token." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAPITokenConfig(rnd, rnd, permissionID, zoneID, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, "name", rnd),
					resource.TestCheckResourceAttr(resourceID, "request_ip_in.0", "10.0.0.0/8"),
					resource.TestCheckResourceAttr(resourceID, "request_ip_not_in.0", "10.0.255.0/24"),
				),
			},
			{
				Config: testAPITokenConfig(rnd, rnd, permissionID, zoneID, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, "name", rnd),
					resource.TestCheckNoResourceAttr(resourceID, "request_ip_in.0"),
					resource.TestCheckNoResourceAttr(resourceID, "request_ip_not_in.0"),
				),
			},
		},
	})
}

func testAPITokenConfig(resourceID, name, permissionID, zoneID string, ips bool) string {
	var ipIn, ipNotIn string

	if ips {
		ipIn = `request_ip_in = ["10.0.0.0/8"]`
		ipNotIn = `request_ip_not_in = ["10.0.255.0/24"]`
	}

	return fmt.Sprintf(`
		resource "cloudflare_api_token" "%[1]s" {
		  name = "%[2]s"

          %[5]s
          %[6]s
		
		  policy {
			permission_groups = [
			  "%[3]s",
			]
			resources = [
			  "com.cloudflare.api.account.zone.%[4]s",
			]
		  }
		}
		`, resourceID, name, permissionID, zoneID, ipIn, ipNotIn)
}
