package api_shield_operation_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareAPIShieldOperation_Create(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// endpoint does not yet support the API tokens without an explicit scope.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceID := "cloudflare_api_shield_operation." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAPIShieldOperationDelete,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPIShieldOperation(rnd, zoneID, cloudflare.APIShieldBasicOperation{Method: "GET", Host: domain, Endpoint: "/example/path"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "method", "GET"),
					resource.TestCheckResourceAttr(resourceID, "host", domain),
					resource.TestCheckResourceAttr(resourceID, "endpoint", "/example/path"),
				),
			},
		},
	})
}

func TestAccCloudflareAPIShieldOperation_ForceNew(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// endpoint does not yet support the API tokens without an explicit scope.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceID := "cloudflare_api_shield_operation." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAPIShieldOperationDelete,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPIShieldOperation(rnd, zoneID, cloudflare.APIShieldBasicOperation{Method: "GET", Host: domain, Endpoint: "/example/path"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "method", "GET"),
					resource.TestCheckResourceAttr(resourceID, "host", domain),
					resource.TestCheckResourceAttr(resourceID, "endpoint", "/example/path"),
				),
			},
			{
				Config: testAccCloudflareAPIShieldOperation(rnd, zoneID, cloudflare.APIShieldBasicOperation{Method: "POST", Host: domain, Endpoint: "/example/path"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "method", "POST"), // check that we've 'updated' the value
					resource.TestCheckResourceAttr(resourceID, "host", domain),
					resource.TestCheckResourceAttr(resourceID, "endpoint", "/example/path"),
				),
			},
		},
	})
}

func testAccCheckAPIShieldOperationDelete(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
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
	return acctest.LoadTestCase("apishieldoperation.tf", resourceName, zone, op.Method, op.Host, op.Endpoint)
}
