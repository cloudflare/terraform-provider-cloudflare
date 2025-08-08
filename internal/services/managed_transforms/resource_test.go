package managed_transforms_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v5/managed_transforms"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"

	cloudflare "github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/pkg/errors"
)

func init() {
	resource.AddTestSweepers("cloudflare_managed_headers", &resource.Sweeper{
		Name: "cloudflare_managed_headers",
		F:    testSweepCloudflareManagedTransforms,
	})
}

func testSweepCloudflareManagedTransforms(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	if client == nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client"))
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}

	err := client.ManagedTransforms.Delete(
		ctx,
		managed_transforms.ManagedTransformDeleteParams{
			ZoneID: cloudflare.F(zoneID),
		},
	)

	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to delete Cloudflare managed transforms: %s", err))
	}

	return nil
}

func TestAccCloudflareManagedHeaders(t *testing.T) {
	t.Parallel()


	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_managed_transforms." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareManagedTransforms(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.0.id", "add_true_client_ip_headers"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.1.id", "add_visitor_location_headers"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.1.enabled", "true"),

					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.0.id", "add_security_headers"),
					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.0.enabled", "true"),
				),
			},
			{
				Config: testAccCheckCloudflareManagedTransformsRemovedHeader(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.0.id", "add_true_client_ip_headers"),
					resource.TestCheckResourceAttr(resourceName, "managed_request_headers.0.enabled", "true"),

					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.0.id", "add_security_headers"),
					resource.TestCheckResourceAttr(resourceName, "managed_response_headers.0.enabled", "true"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCloudflareManagedTransforms(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransforms.tf", rnd, zoneID)
}

func testAccCheckCloudflareManagedTransformsRemovedHeader(rnd, zoneID string) string {
	return acctest.LoadTestCase("managedtransformsremovedheader.tf", rnd, zoneID)
}
