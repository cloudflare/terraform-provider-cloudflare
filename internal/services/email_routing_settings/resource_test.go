package email_routing_settings_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/email_routing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_email_routing_settings", &resource.Sweeper{
		Name: "cloudflare_email_routing_settings",
		F: func(region string) error {
			client := acctest.SharedClient()
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			ctx := context.Background()

			// Check if email routing is enabled first
			settings, err := client.EmailRouting.Get(ctx, email_routing.EmailRoutingGetParams{
				ZoneID: cloudflare.F(zoneID),
			})
			if err != nil {
				return fmt.Errorf("failed to get email routing settings: %w", err)
			}

			// Delete email routing DNS configuration (removes all DNS records and subdomains)
			deletedRecords, err := client.EmailRouting.DNS.Delete(ctx, email_routing.DNSDeleteParams{
				ZoneID: cloudflare.F(zoneID),
			})
			if err != nil {
				fmt.Printf("Note: DNS delete returned error (might be expected): %v\n", err)
			} else if deletedRecords != nil && deletedRecords.Result != nil {
				fmt.Printf("Deleted %d email routing DNS records\n", len(deletedRecords.Result))
			}

			// Also disable email routing if it's still enabled
			if settings.Enabled == email_routing.SettingsEnabledTrue {
				_, err = client.EmailRouting.Disable(ctx, email_routing.EmailRoutingDisableParams{
					ZoneID: cloudflare.F(zoneID),
					Body:   map[string]interface{}{"name": settings.Name},
				})
				if err != nil {
					return fmt.Errorf("failed to disable email routing settings: %w", err)
				}
				fmt.Printf("Disabled email routing settings\n")
			} else {
				fmt.Printf("Email routing settings are already disabled, DNS cleanup completed\n")
			}

			return nil
		},
	})
}

func testEmailRoutingSettingsConfig(resourceID, zoneID string, enabled bool) string {
	return acctest.LoadTestCase("emailroutingsettingsconfig.tf", resourceID, zoneID, enabled)
}

func TestAccTestEmailRoutingSettings(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_email_routing_settings." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testEmailRoutingSettingsConfig(rnd, zoneID, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
				),
			},
		},
	})
}
