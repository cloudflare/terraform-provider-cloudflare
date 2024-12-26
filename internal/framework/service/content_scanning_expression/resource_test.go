package content_scanning_expression_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func init() {
	resource.AddTestSweepers("cloudflare_content_scanning_expression", &resource.Sweeper{
		Name: "cloudflare_content_scanning_expression",
		F:    testSweepCloudflareCSExpression,
	})
}

func testSweepCloudflareCSExpression(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Sweeper: failed to create Cloudflare client: %s", clientErr))
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}
	// fetch existing expressions from API
	expressions, err := client.ContentScanningListCustomExpressions(ctx, cfv1.ZoneIdentifier(zoneID), cfv1.ContentScanningListCustomExpressionsParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Sweeper: error listing customs scan expressions for Content Scanning: %s", err))
		return err
	}
	for _, exp := range expressions {
		deleteParam := cfv1.ContentScanningDeleteCustomExpressionsParams{ID: exp.ID}
		_, err := client.ContentScanningDeleteCustomExpression(ctx, cfv1.ZoneIdentifier(zoneID), deleteParam)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Sweeper: error deleting custom scan expression for Content Scanning: %s", err))
		}
	}

	return nil
}

func TestAccCloudflareContentScanningExpression_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_content_scanning_expression.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccBasicConfig(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name+"_first", "zone_id", zoneID),
					resource.TestCheckResourceAttr(name+"_first", "payload", "lookup_json_string(http.request.body.raw, \"file\")"),

					resource.TestCheckResourceAttr(name+"_second", "zone_id", zoneID),
					resource.TestCheckResourceAttr(name+"_second", "payload", "lookup_json_string(http.request.body.raw, \"document\")"),
				),
			},
			{
				Config: testAccBasicConfigChange(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name+"_second", "zone_id", zoneID),
					resource.TestCheckResourceAttr(name+"_second", "payload", "lookup_json_string(http.request.body.raw, \"txt\")"),
				),
			},
		},
	})
}

func testAccBasicConfig(name, zoneID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_content_scanning_expression" "%[1]s_first" {
    	zone_id = "%[2]s"
		payload = "lookup_json_string(http.request.body.raw, \"file\")"
  }

  resource "cloudflare_content_scanning_expression" "%[1]s_second" {
	zone_id = "%[2]s"
	payload = "lookup_json_string(http.request.body.raw, \"document\")"
  }`, name, zoneID)
}

func testAccBasicConfigChange(name, zoneID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_content_scanning_expression" "%[1]s_second" {
	zone_id = "%[2]s"
	payload = "lookup_json_string(http.request.body.raw, \"txt\")"
  }`, name, zoneID)
}
