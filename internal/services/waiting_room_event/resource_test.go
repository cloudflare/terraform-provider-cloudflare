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

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_waiting_room_event", &resource.Sweeper{
		Name: "cloudflare_waiting_room_event",
		F:    testSweepCloudflareWaitingRoomEvents,
	})
}

func testSweepCloudflareWaitingRoomEvents(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping waiting room events sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	// First, list all waiting rooms
	waitingRooms, err := client.ListWaitingRooms(ctx, zoneID)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch waiting rooms: %s", err))
		return fmt.Errorf("failed to fetch waiting rooms: %w", err)
	}

	if len(waitingRooms) == 0 {
		tflog.Info(ctx, "No waiting rooms found, skipping waiting room events sweep")
		return nil
	}

	// For each waiting room, list and delete its events
	for _, waitingRoom := range waitingRooms {
		events, err := client.ListWaitingRoomEvents(ctx, zoneID, waitingRoom.ID)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to fetch waiting room events for room %s: %s", waitingRoom.ID, err))
			continue
		}

		for _, event := range events {
			// Use standard filtering helper on the event name
			if !utils.ShouldSweepResource(event.Name) {
				continue
			}

			tflog.Info(ctx, fmt.Sprintf("Deleting waiting room event: %s (waiting room: %s, zone: %s)", event.ID, waitingRoom.ID, zoneID))
			err := client.DeleteWaitingRoomEvent(ctx, zoneID, waitingRoom.ID, event.ID)
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete waiting room event %s: %s", event.ID, err))
				continue
			}
			tflog.Info(ctx, fmt.Sprintf("Deleted waiting room event: %s", event.ID))
		}
	}

	return nil
}

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
