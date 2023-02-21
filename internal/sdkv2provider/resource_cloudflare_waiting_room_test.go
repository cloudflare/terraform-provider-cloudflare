package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudflareWaitingRoom_Create(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_waiting_room.%s", rnd)
	waitingRoomName := fmt.Sprintf("waiting_room_%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareWaitingRoomDestroy,
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
				),
			},
		},
	})
}

func testAccCheckCloudflareWaitingRoomDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

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
	return fmt.Sprintf(`
resource "cloudflare_waiting_room" "%[1]s" {
  name                      = "%[2]s"
  zone_id                   = "%[3]s"
  host                      = "www.%[4]s"
  new_users_per_minute      = 400
  total_active_users        = 405
  path                      = "%[5]s"
  session_duration          = 10
  queueing_method           = "fifo"
  custom_page_html          = "foobar"
  default_template_language = "en-US"
  description               = "my desc"
  disable_session_renewal   = true
  suspended                 = true
  queue_all                 = false
  json_response_enabled     = true
}
`, resourceName, waitingRoomName, zoneID, domain, path)
}
