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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_email_routing_settings", &resource.Sweeper{
		Name: "cloudflare_email_routing_settings",
		F: func(region string) error {
			ctx := context.Background()
			client := acctest.SharedClient()
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

			if zoneID == "" {
				tflog.Info(ctx, "Skipping email routing settings sweep: CLOUDFLARE_ZONE_ID not set")
				return nil
			}

			// Check if email routing is enabled first
			settings, err := client.EmailRouting.Get(ctx, email_routing.EmailRoutingGetParams{
				ZoneID: cloudflare.F(zoneID),
			})
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to get email routing settings: %s", err))
				return fmt.Errorf("failed to get email routing settings: %w", err)
			}

			// Delete email routing DNS configuration (removes all DNS records and subdomains)
			tflog.Info(ctx, fmt.Sprintf("Deleting email routing DNS configuration (zone: %s)", zoneID))
			deletedRecords, err := client.EmailRouting.DNS.Delete(ctx, email_routing.DNSDeleteParams{
				ZoneID: cloudflare.F(zoneID),
			})
			if err != nil {
				tflog.Info(ctx, fmt.Sprintf("Note: DNS delete returned error (might be expected): %v", err))
			} else if deletedRecords != nil && deletedRecords.Result != nil {
				tflog.Info(ctx, fmt.Sprintf("Deleted %d email routing DNS records", len(deletedRecords.Result)))
			}

			// Also disable email routing if it's still enabled
			if settings.Enabled == email_routing.SettingsEnabledTrue {
				tflog.Info(ctx, fmt.Sprintf("Disabling email routing settings (zone: %s)", zoneID))
				_, err = client.EmailRouting.Disable(ctx, email_routing.EmailRoutingDisableParams{
					ZoneID: cloudflare.F(zoneID),
					Body:   map[string]interface{}{"name": settings.Name},
				})
				if err != nil {
					tflog.Error(ctx, fmt.Sprintf("Failed to disable email routing settings: %s", err))
					return fmt.Errorf("failed to disable email routing settings: %w", err)
				}
				tflog.Info(ctx, "Disabled email routing settings")
			} else {
				tflog.Info(ctx, "Email routing settings are already disabled, DNS cleanup completed")
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
