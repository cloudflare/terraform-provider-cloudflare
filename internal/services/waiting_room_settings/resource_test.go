package waiting_room_settings_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_waiting_room_settings", &resource.Sweeper{
		Name: "cloudflare_waiting_room_settings",
		F:    testSweepCloudflareWaitingRoomSettings,
	})
}

func testSweepCloudflareWaitingRoomSettings(r string) error {
	ctx := context.Background()
	// Waiting Room Settings is a zone-level configuration setting.
	// It's a singleton setting per zone, not something that accumulates.
	// No sweeping required.
	tflog.Info(ctx, "Waiting Room Settings doesn't require sweeping (zone setting)")
	return nil
}

func TestAccCloudflareWaitingRoomSettings_Create(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_waiting_room_settings.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareWaitingRoomSettings(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "search_engine_crawler_bypass", "true"),
				),
			},
		},
	})
}

func testAccCloudflareWaitingRoomSettings(resourceName, zoneID string) string {
	return acctest.LoadTestCase("waitingroomsettings.tf", resourceName, zoneID)
}
