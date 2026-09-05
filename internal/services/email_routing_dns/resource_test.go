package email_routing_dns_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/email_routing"
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
	resource.AddTestSweepers("cloudflare_email_routing_dns", &resource.Sweeper{
		Name: "cloudflare_email_routing_dns",
		F: func(region string) error {
			ctx := context.Background()
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

			if zoneID == "" {
				tflog.Info(ctx, "Skipping email_routing_dns sweep: CLOUDFLARE_ZONE_ID not set")
				return nil
			}

			client := acctest.SharedClient()

			_, err := client.EmailRouting.DNS.Delete(ctx, email_routing.DNSDeleteParams{
				ZoneID: cloudflare.F(zoneID),
			})
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete email routing DNS for zone %s: %s", zoneID, err))
				return fmt.Errorf("failed to delete email routing DNS: %w", err)
			}

			tflog.Info(ctx, fmt.Sprintf("Deleted email routing DNS for zone: %s", zoneID))
			return nil
		},
	})
}

// noRefreshFields lists fields that are set at create time but not returned by
// the GET endpoint and therefore cannot round-trip through import.
var noRefreshFields = []string{
	"created",
	"enabled",
	"modified",
	"name",
	"skip_wizard",
	"status",
	"support_subaddress",
	"tag",
}

func TestAccCloudflareEmailRoutingDNS_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_email_routing_dns." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareEmailRoutingDNSBasic(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: noRefreshFields,
			},
		},
	})
}

func testAccCloudflareEmailRoutingDNSBasic(rnd, zoneID string) string {
	return acctest.LoadTestCase("emailroutingdnsbasic.tf", rnd, zoneID)
}
