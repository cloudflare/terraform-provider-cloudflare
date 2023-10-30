package sdkv2provider

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareZoneHold_Full(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zone_hold.%s", rnd)
	currentTime := time.Now()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneHoldOnResourceConfig(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareZoneDNSSECDataSourceID(name),
					resource.TestCheckResourceAttrSet(name, consts.ZoneIDSchemaKey),
					resource.TestCheckResourceAttr(name, "hold", "true"),
				),
			},
			{
				Config: testAccCloudflareZoneHoldOffWithTimeAfterResourceConfig(zoneID, rnd, currentTime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareZoneDNSSECDataSourceID(name),
					resource.TestCheckResourceAttrSet(name, consts.ZoneIDSchemaKey),
					resource.TestCheckResourceAttr(name, "hold", "false"),
					resource.TestCheckResourceAttr(name, "hold_after", currentTime.Add(time.Duration(1*time.Hour)).UTC().Format(time.RFC3339)),
				),
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"hold_after"},
			},
		},
	})
}

func testAccCloudflareZoneHoldOnResourceConfig(zoneID string, name string) string {
	return fmt.Sprintf(`
resource "cloudflare_zone_hold" "%s" {
	zone_id = "%s"
	hold = true
}`, name, zoneID)
}

func testAccCloudflareZoneHoldOffWithTimeAfterResourceConfig(zoneID string, name string, t time.Time) string {
	hold := t.Add(time.Duration(1 * time.Hour)).UTC().Format(time.RFC3339)
	return fmt.Sprintf(`
resource "cloudflare_zone_hold" "%s" {
	zone_id = "%s"
	hold = false
	hold_after = "%s"
}`, name, zoneID, hold)
}
