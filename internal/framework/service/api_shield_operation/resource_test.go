package apishieldoperation_test

import (
	"context"
	"errors"
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

func init() {
	resource.AddTestSweepers("cloudflare_api_shield_operation", &resource.Sweeper{
		Name: "cloudflare_api_shield_operation",
		F:    testSweepCloudflareCloudAPIShieldOperations,
	})
}

func testSweepCloudflareCloudAPIShieldOperations(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	zone := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zone == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}

	operations, _, err := client.ListAPIShieldOperations(ctx, cloudflare.ZoneIdentifier(zone), cloudflare.ListAPIShieldOperationsParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list Cloudflare API Shield Operations: %s", err))
	}

	for _, operation := range operations {
		if err := client.DeleteAPIShieldOperation(
			ctx,
			cloudflare.ZoneIdentifier(zone),
			cloudflare.DeleteAPIShieldOperationParams{OperationID: operation.ID},
		); err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete Cloudflare API Shield Operation: %s", err))
		}
	}

	return nil
}

func TestAccCloudflareAPIShieldOperation_Create(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceID := "cloudflare_api_shield_operation." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		CheckDestroy:             testAccCheckAPIShieldOperationDelete,
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPIShieldOperation(rnd, zoneID, cloudflare.APIShieldBasicOperation{Method: "GET", Host: domain, Endpoint: "/example/path/foo/{fooId}"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "method", "GET"),
					resource.TestCheckResourceAttr(resourceID, "host", domain),
					resource.TestCheckResourceAttr(resourceID, "endpoint", "/example/path/foo/{fooId}"),
				),
			},
		},
	})
}

func TestAccCloudflareAPIShieldOperation_ForceNew(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceID := "cloudflare_api_shield_operation." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		CheckDestroy:             testAccCheckAPIShieldOperationDelete,
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPIShieldOperation(rnd, zoneID, cloudflare.APIShieldBasicOperation{Method: "GET", Host: domain, Endpoint: "/example/path/foo/{fooId}"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "method", "GET"),
					resource.TestCheckResourceAttr(resourceID, "host", domain),
					resource.TestCheckResourceAttr(resourceID, "endpoint", "/example/path/foo/{fooId}"),
				),
			},
			{
				Config: testAccCloudflareAPIShieldOperation(rnd, zoneID, cloudflare.APIShieldBasicOperation{Method: "POST", Host: domain, Endpoint: "/example/path/foo/{fooId}"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "method", "POST"), // check that we've 'updated' the value
					resource.TestCheckResourceAttr(resourceID, "host", domain),
					resource.TestCheckResourceAttr(resourceID, "endpoint", "/example/path/foo/{fooId}"),
				),
			},
		},
	})
}

func testAccCheckAPIShieldOperationDelete(s *terraform.State) error {
	client, err := acctest.SharedV1Client()
	if err != nil {
		tflog.Error(context.Background(), fmt.Sprintf("Failed to create Cloudflare client: %s", err))
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_api_shield_operation" {
			continue
		}

		_, err := client.GetAPIShieldOperation(
			context.Background(),
			cloudflare.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]),
			cloudflare.GetAPIShieldOperationParams{
				OperationID: rs.Primary.Attributes["id"],
			},
		)
		if err == nil {
			return fmt.Errorf("operation still exists")
		}

		var notFoundError *cloudflare.NotFoundError
		if !errors.As(err, &notFoundError) {
			return fmt.Errorf("expected not found error but got: %w", err)
		}
	}

	return nil
}

func testAccCloudflareAPIShieldOperation(resourceName, zone string, op cloudflare.APIShieldBasicOperation) string {
	return fmt.Sprintf(`
	resource "cloudflare_api_shield_operation" "%[1]s" {
		zone_id = "%[2]s"
		method = "%[3]s"
		host = "%[4]s"
		endpoint = "%[5]s"
	}
`, resourceName, zone, op.Method, op.Host, op.Endpoint)
}
