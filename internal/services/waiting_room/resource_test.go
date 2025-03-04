package waiting_room_test

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
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func init() {
	resource.AddTestSweepers("cloudflare_waiting_room", &resource.Sweeper{
		Name: "cloudflare_waiting_room",
		F:    testSweepCloudflareWaitingRoom,
	})
}

func testSweepCloudflareWaitingRoom(r string) error {
	ctx := context.Background()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	resp, err := client.ListWaitingRooms(ctx, zoneID)
	if err != nil {
		return err
	}

	for _, room := range resp {
		err := client.DeleteWaitingRoom(ctx, zoneID, room.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func TestAccCloudflareWaitingRoom_Create(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_waiting_room.%s", rnd)
	waitingRoomName := fmt.Sprintf("waiting_room_%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWaitingRoomDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareWaitingRoom(rnd, waitingRoomName, zoneID, domain, "/foobar"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", waitingRoomName),
					resource.TestCheckResourceAttr(name, "description", "my desc"),
					resource.TestCheckResourceAttr(name, "queueing_method", "fifo"),
					resource.TestCheckResourceAttr(name, "custom_page_html", "foobar"),
					resource.TestCheckResourceAttr(name, "default_template_language", "en-US"),
					resource.TestCheckResourceAttr(name, "disable_session_renewal", "true"),
					resource.TestCheckResourceAttr(name, "suspended", "true"),
					resource.TestCheckResourceAttr(name, "queue_all", "false"),
					resource.TestCheckResourceAttr(name, "new_users_per_minute", "400"),
					resource.TestCheckResourceAttr(name, "total_active_users", "405"),
					resource.TestCheckResourceAttr(name, "session_duration", "10"),
					resource.TestCheckResourceAttr(name, "json_response_enabled", "true"),
					resource.TestCheckResourceAttr(name, "cookie_suffix", "queue1"),
					resource.TestCheckResourceAttr(name, "additional_routes.#", "2"),
					resource.TestCheckResourceAttr(name, "additional_routes.0.host", "shop1."+domain),
					resource.TestCheckResourceAttr(name, "additional_routes.0.path", "/foobar"),
					resource.TestCheckResourceAttr(name, "additional_routes.1.host", "shop2."+domain),
					resource.TestCheckResourceAttr(name, "additional_routes.1.path", "/"),
					resource.TestCheckResourceAttr(name, "queueing_status_code", "200"),
				),
			},
		},
	})
}

func testAccCheckCloudflareWaitingRoomDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_waiting_room" {
			continue
		}

		_, err := client.WaitingRoom(context.Background(), rs.Primary.Attributes[consts.ZoneIDSchemaKey], rs.Primary.Attributes["name"])
		if err == nil {
			return fmt.Errorf("Waiting room still exists")
		}
	}

	return nil
}

func testAccCloudflareWaitingRoom(resourceName, waitingRoomName, zoneID, domain, path string) string {
	return acctest.LoadTestCase("waitingroom.tf", resourceName, waitingRoomName, zoneID, domain, path)
}
