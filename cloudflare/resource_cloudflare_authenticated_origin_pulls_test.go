package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccCloudflareAuthenticatedOriginPullsGlobalConfig(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_authenticated_origin_pulls.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAuthenticatedOriginPullsFullGlobalConfig(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "status", "on"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAuthenticatedOriginPullsFullGlobalConfig(zoneID, name string) string {
	return fmt.Sprintf(`
  resource "cloudflare_authenticated_origin_pulls" "%[2]s" {
	  zone_id        = "%[1]s"
    status = "on"
  }`, zoneID, name)
}
