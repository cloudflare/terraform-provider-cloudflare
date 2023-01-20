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

func TestAccCloudflareWaitingRoomRules_Create(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := generateRandomResourceName()
	waitingRoomName := fmt.Sprintf("waiting_room_%s", rnd)
	name := fmt.Sprintf("cloudflare_waiting_room_rules.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareWaitingRoomRulesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareWaitingRoomRules(rnd, zoneID, domain, waitingRoomName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttrSet(name, "waiting_room_id"),

					resource.TestCheckResourceAttr(name, "rules.0.description", "ip bypass"),
					resource.TestCheckResourceAttr(name, "rules.0.expression", "ip.src in {192.0.2.1}"),
					resource.TestCheckResourceAttr(name, "rules.0.action", "bypass_waiting_room"),
					resource.TestCheckResourceAttr(name, "rules.0.status", "enabled"),
					resource.TestCheckResourceAttr(name, "rules.0.version", "1"),

					resource.TestCheckResourceAttr(name, "rules.1.description", "query string bypass"),
					resource.TestCheckResourceAttr(name, "rules.1.expression", "http.request.uri.query contains \"bypass=true\""),
					resource.TestCheckResourceAttr(name, "rules.1.action", "bypass_waiting_room"),
					resource.TestCheckResourceAttr(name, "rules.1.status", "disabled"),
					resource.TestCheckResourceAttr(name, "rules.1.version", "1"),
				),
			},
		},
	})
}

func testAccCheckCloudflareWaitingRoomRulesDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_waiting_room_rules" {
			continue
		}

		waitingRoomRules, err := client.ListWaitingRoomRules(context.Background(), cloudflare.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]), cloudflare.ListWaitingRoomRuleParams{
			WaitingRoomID: rs.Primary.Attributes["waiting_room_id"],
		})
		if err == nil {
			return fmt.Errorf("error reading waiting room rules")
		}
		// when no rules exist, an empty list is returned
		if len(waitingRoomRules) != 0 {
			return fmt.Errorf("waiting room rules still exists")
		}
	}

	return nil
}

func testAccCloudflareWaitingRoomRules(resourceName, zoneID, domain, waitingRoomName string) string {
	return fmt.Sprintf(`
resource "cloudflare_waiting_room" "%[1]s" {
  name                      = "%[4]s"
  zone_id                   = "%[2]s"
  host                      = "www.%[3]s"
  new_users_per_minute      = 400
  total_active_users        = 405
  path                      = "/foobar"
  session_duration          = 10
  custom_page_html          = "foobar"
  description               = "my desc"
  disable_session_renewal   = true
  suspended                 = true
  queue_all                 = false
  json_response_enabled     = true
}

resource "cloudflare_waiting_room_rules" "%[1]s" {
  zone_id		  = "%[2]s"
  waiting_room_id = cloudflare_waiting_room.%[1]s.id

  rules {
    action      = "bypass_waiting_room"
    expression  = "ip.src in {192.0.2.1}"
    description = "ip bypass"
    status 	    = "enabled"
  }

  rules {
    action      = "bypass_waiting_room"
    expression 	= "http.request.uri.query contains \"bypass=true\""
    description = "query string bypass"
    status 	    = "disabled"
  }
}
`, resourceName, zoneID, domain, waitingRoomName)
}
