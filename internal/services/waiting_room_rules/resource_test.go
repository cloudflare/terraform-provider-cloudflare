package waiting_room_rules_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareWaitingRoomRules_Create(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	waitingRoomName := fmt.Sprintf("waiting_room_%s", rnd)
	name := fmt.Sprintf("cloudflare_waiting_room_rules.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWaitingRoomRulesDestroy,
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
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

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
	return acctest.LoadTestCase("waitingroomrules.tf", resourceName, zoneID, domain, waitingRoomName)
}
