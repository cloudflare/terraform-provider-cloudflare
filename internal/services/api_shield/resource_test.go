package api_shield_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/api_gateway"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_api_shield", &resource.Sweeper{
		Name: "cloudflare_api_shield",
		F:    testSweepCloudflareAPIShield,
	})
}

func testSweepCloudflareAPIShield(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping API Shield sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	tflog.Info(ctx, fmt.Sprintf("Clearing API Shield configuration for zone: %s", zoneID))
	_, err := client.APIGateway.Configurations.Update(
		ctx,
		api_gateway.ConfigurationUpdateParams{
			ZoneID: cloudflare.F(zoneID),
		},
		option.WithRequestBody("application/json", []byte(`{"auth_id_characteristics":[]}`)),
	)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to clear API Shield configuration: %s", err))
		return fmt.Errorf("failed to clear API Shield configuration: %w", err)
	}

	tflog.Info(ctx, "Cleared API Shield configuration")
	return nil
}

func TestAccCloudflareAPIShieldBasic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// endpoint does not yet support the API tokens without an explicit scope.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceID := "cloudflare_api_shield." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create the resource with a single header-based auth characteristic
			{
				Config: testAccCloudflareAPIShieldBasic(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.#", "1"),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.0.%", "2"),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.0.type", "header"),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.0.name", "my-example-header"),
				),
			},
			// Step 2: Update the resource to use a cookie-based auth characteristic instead
			// Tests that we can change both the type and name of characteristics
			{
				Config: testAccCloudflareAPIShieldUpdate(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.#", "1"),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.0.%", "2"),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.0.type", "cookie"),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.0.name", "my-example-cookie"),
				),
			},
			// Step 3: Remove all auth characteristics by just setting to an empty array
			// Tests that we can remove auth characteristics
			{
				Config: testAccCloudflareAPIShieldDeleteEmpty(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.#", "0"),
				),
			},
			// Step 4: Add multiple auth characteristics
			// Tests that we can have multiple auth characteristics of different types
			{
				Config: testAccCloudflareAPIShieldMultipleCharacteristics(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.#", "2"),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.0.type", "header"),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.0.name", "auth-header"),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.1.type", "cookie"),
					resource.TestCheckResourceAttr(resourceID, "auth_id_characteristics.1.name", "auth-cookie"),
				),
			},
			// Step 5: Import test
			// Verifies that the resource can be imported with the correct ID format
			{
				ResourceName:      resourceID,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					// this is a single resource per zone
					return zoneID, nil
				},
			},
		},
	})
}

func testAccCloudflareAPIShieldBasic(resourceName, zone string) string {
	return acctest.LoadTestCase("basic.tf", resourceName, zone)
}

func testAccCloudflareAPIShieldUpdate(resourceName, zone string) string {
	return acctest.LoadTestCase("update.tf", resourceName, zone)
}

func testAccCloudflareAPIShieldMultipleCharacteristics(resourceName, zone string) string {
	return acctest.LoadTestCase("multiple.tf", resourceName, zone)
}

func testAccCloudflareAPIShieldDeleteEmpty(resourceName, zone string) string {
	return acctest.LoadTestCase("delete_empty.tf", resourceName, zone)
}
