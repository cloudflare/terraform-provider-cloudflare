package waiting_room_event_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareWaitingRoomEvent_Create(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	waitingRoomID := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_waiting_room_event.%s", rnd)
	waitingRoomEventName := fmt.Sprintf("waiting_room_event_%s", rnd)
	waitingRoomName := fmt.Sprintf("waiting_room_%s", rnd)
	eventStartTime := time.Now().UTC()
	eventEndTime := eventStartTime.Add(5 * time.Minute).UTC()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWaitingRoomEventDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareWaitingRoomEvent(rnd, waitingRoomEventName, zoneID, waitingRoomID, eventStartTime, eventEndTime, domain, waitingRoomName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", waitingRoomEventName),
					resource.TestCheckResourceAttrSet(name, "waiting_room_id"),
					resource.TestCheckResourceAttr(name, "event_start_time", eventStartTime.Format(time.RFC3339)),
					resource.TestCheckResourceAttr(name, "event_end_time", eventEndTime.Format(time.RFC3339)),
					resource.TestCheckResourceAttr(name, "description", "my desc"),
					resource.TestCheckResourceAttr(name, "custom_page_html", "foobar"),
					resource.TestCheckResourceAttr(name, "disable_session_renewal", "true"),
					resource.TestCheckResourceAttr(name, "suspended", "true"),
					resource.TestCheckResourceAttr(name, "queueing_method", "fifo"),
					resource.TestCheckResourceAttr(name, "new_users_per_minute", "400"),
					resource.TestCheckResourceAttr(name, "total_active_users", "405"),
					resource.TestCheckResourceAttr(name, "session_duration", "10"),
					resource.TestCheckResourceAttr(name, "shuffle_at_event_start", "false"),
					resource.TestCheckNoResourceAttr(name, "prequeue_start_time"),
				),
			},
		},
	})
}

func testAccCheckCloudflareWaitingRoomEventDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_waiting_room_event" {
			continue
		}

		_, err := client.WaitingRoomEvent(context.Background(), rs.Primary.Attributes[consts.ZoneIDSchemaKey], rs.Primary.Attributes["waiting_room_id"], rs.Primary.Attributes["id"])
		if err == nil {
			return fmt.Errorf("waiting room event still exists")
		}
	}

	return nil
}

func testAccCloudflareWaitingRoomEvent(resourceName, waitingRoomEventName, zoneID, waitingRoomID string, startTime, endTime time.Time, domain, waitingRoomName string) string {
	return acctest.LoadTestCase("waitingroomevent.tf", resourceName, waitingRoomEventName, zoneID, waitingRoomID, startTime.Format(time.RFC3339), endTime.Format(time.RFC3339), domain, waitingRoomName)
}
