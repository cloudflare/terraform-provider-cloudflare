package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccLogpullRetentionSetStatus(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_logpull_retention." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testLogpullRetentionSetConfig(rnd, zoneID, "false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "enabled", "false"),
				),
			},
		},
	})
}

func testLogpullRetentionSetConfig(id, zoneID, enabled string) string {
	return fmt.Sprintf(`
  resource "cloudflare_logpull_retention" "%[1]s" {
    zone_id = "%[2]s"
	  enabled = "%[3]s"
  }`, id, zoneID, enabled)
}
