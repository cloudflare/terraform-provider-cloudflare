package zone_lockdown_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
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
	resource.AddTestSweepers("cloudflare_zone_lockdown", &resource.Sweeper{
		Name: "cloudflare_zone_lockdown",
		F:    testSweepCloudflareZoneLockdowns,
	})
}

func testSweepCloudflareZoneLockdowns(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping zone lockdowns sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	lockdowns, _, err := client.ListZoneLockdowns(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.LockdownListParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch zone lockdowns: %s", err))
		return fmt.Errorf("failed to fetch zone lockdowns: %w", err)
	}

	if len(lockdowns) == 0 {
		tflog.Info(ctx, "No zone lockdowns to sweep")
		return nil
	}

	for _, lockdown := range lockdowns {
		// Use standard filtering helper on the description field
		if !utils.ShouldSweepResource(lockdown.Description) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting zone lockdown: %s (zone: %s)", lockdown.ID, zoneID))
		_, err := client.DeleteZoneLockdown(ctx, cloudflare.ZoneIdentifier(zoneID), lockdown.ID)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete zone lockdown %s: %s", lockdown.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted zone lockdown: %s", lockdown.ID))
	}

	return nil
}

func TestAccCloudflareZoneLockdown(t *testing.T) {
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone_lockdown." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneLockdownConfig(rnd, zoneID, "false", "1", "this is notes", rnd+"."+zoneName+"/*", "ip", "198.51.100.4"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "urls.#", "1"),
					resource.TestCheckResourceAttr(name, "configurations.#", "1"),
				),
			},
		},
	})
}

// test creating a config with only the required fields.
func TestAccCloudflareZoneLockdown_OnlyRequired(t *testing.T) {
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zone_lockdown." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneLockdownConfig(rnd, zoneID, "false", "1", "this is notes", rnd+"."+zoneName+"/*", "ip", "198.51.100.4"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "urls.#", "1"),
					resource.TestCheckResourceAttr(name, "configurations.#", "1"),
				),
			},
		},
	})
}

func TestAccCloudflareZoneLockdown_Import(t *testing.T) {
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	// name := "cloudflare_zone_lockdown." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareZoneLockdownConfig(rnd, zoneID, "false", "1", "this is notes", rnd+"."+zoneName+"/*", "ip", "198.51.100.4"),
			},
			// {
			// 	ResourceName:        name,
			// 	ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
			// 	ImportState:         true,
			// 	ImportStateVerify:   true,
			// },
		},
	})
}

func testCloudflareZoneLockdownConfig(resourceID, zoneID, paused, priority, description, url, target, value string) string {
	return acctest.LoadTestCase("cloudflarezonelockdownconfig.tf", resourceID, zoneID, paused, priority, description, url, target, value)
}
