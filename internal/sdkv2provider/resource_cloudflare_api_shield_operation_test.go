package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"os"
	"testing"
)

func TestAccAPIShieldOperation_Create(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// endpoint does not yet support the API tokens without an explicit scope.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	resourceID := "cloudflare_api_shield_operation." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAPIShieldOperationDelete,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPIShieldOperation(rnd, zoneID, cloudflare.APIShieldBasicOperation{Method: "GET", Host: "api.cloudflare.com", Endpoint: "/client/v4/zones"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "method", "GET"),
					resource.TestCheckResourceAttr(resourceID, "host", "api.cloudflare.com"),
					resource.TestCheckResourceAttr(resourceID, "endpoint", "/client/v4/zones"),
				),
			},
		},
	})
}

func TestAccAPIShieldOperation_ForceNew(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// endpoint does not yet support the API tokens without an explicit scope.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	resourceID := "cloudflare_api_shield_operation." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAPIShieldOperationDelete,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPIShieldOperation(rnd, zoneID, cloudflare.APIShieldBasicOperation{Method: "GET", Host: "api.cloudflare.com", Endpoint: "/client/v4/zones"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "method", "GET"),
					resource.TestCheckResourceAttr(resourceID, "host", "api.cloudflare.com"),
					resource.TestCheckResourceAttr(resourceID, "endpoint", "/client/v4/zones"),
				),
			},
			{
				Config: testAccCloudflareAPIShieldOperation(rnd, zoneID, cloudflare.APIShieldBasicOperation{Method: "POST", Host: "api.cloudflare.com", Endpoint: "/client/v4/zones"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "method", "POST"), // check that we've 'updated' the value
					resource.TestCheckResourceAttr(resourceID, "host", "api.cloudflare.com"),
					resource.TestCheckResourceAttr(resourceID, "endpoint", "/client/v4/zones"),
				),
			},
		},
	})
}

func testAccCheckAPIShieldOperationDelete(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

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
