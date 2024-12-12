package content_scanning_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func init() {
	resource.AddTestSweepers("cloudflare_content_scanning", &resource.Sweeper{
		Name: "cloudflare_content_scanning",
		F:    testSweepCloudflareCS,
	})
}

func testSweepCloudflareCS(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}

	status, err := client.ContentScanningStatus(ctx, cfv1.ZoneIdentifier(zoneID), cfv1.ContentScanningStatusParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to GET Content Scanning status: %s", err))
	}

	if status.Result.Value == "disabled" {
		log.Print("[DEBUG] Content Scanning already disabled")
		return nil
	}
	_, err = client.ContentScanningDisable(ctx, cfv1.ZoneIdentifier(zoneID), cfv1.ContentScanningDisableParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to disable Content Scanning: %s", err))
		return err
	}

	return nil
}

func TestAccCloudflareContentScanning_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_content_scanning.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccContentScanningSimple(rnd, zoneID, "true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "true"),
				),
			},
			{
				Config: testAccContentScanningSimple(rnd, zoneID, "false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "false"),
				),
			},
		},
	})
}

func testAccContentScanningSimple(ID, zoneID, enabled string) string {
	return fmt.Sprintf(`
  resource "cloudflare_content_scanning" "%[1]s" {
    zone_id = "%[2]s"
    enabled = "%[3]s"
  }`, ID, zoneID, enabled)
}
