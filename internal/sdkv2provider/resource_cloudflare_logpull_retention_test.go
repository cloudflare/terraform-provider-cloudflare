package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccLogpullRetentionSetStatus(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_logpull_retention." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testLogpullRetentionSetConfig(rnd, zoneID, "false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
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
