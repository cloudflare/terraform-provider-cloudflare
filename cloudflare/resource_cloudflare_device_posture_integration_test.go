package cloudflare

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudflareDevicePostureIntegrationCreate(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		defer func(apiToken string) {
			os.Setenv("CLOUDFLARE_API_TOKEN", apiToken)
		}(os.Getenv("CLOUDFLARE_API_TOKEN"))
		os.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	token := os.Getenv("WORKSPACE_ONE_API_TOKEN")
	if token == "" {
		t.Fatal("Missing workspace one api token")
	}

	id := os.Getenv("WORKSPACE_ONE_API_ID")
	if id == "" {
		t.Fatal("Missing wokspace one api id")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_device_posture_integration.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccessAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareDevicePostureIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureIntegration(rnd, accountID, id, token),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "workspace_one"),
					resource.TestCheckResourceAttr(name, "interval", "24h"),
					resource.TestCheckResourceAttr(name, "config.0.auth_url", "https://test.uemauth.vmwservices.com/connect/token"),
					resource.TestCheckResourceAttr(name, "config.0.api_url", "https://example.com/api-url"),
					resource.TestCheckResourceAttr(name, "config.0.client_id", "client-id"),
				),
			},
		},
	})
}

func testAccCloudflareDevicePostureIntegration(rnd, accountID, ws1ID, ws1Token string) string {
	return fmt.Sprintf(`
resource "cloudflare_device_posture_integration" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "workspace_one"
	interval                  = "24h"
	config {
		api_url       =  "https://example.com/api-url"
		auth_url      =  "https://test.uemauth.vmwservices.com/connect/token"
		client_id     =  "%[3]s"
		client_secret =  "%[4]s"
	}
}
`, rnd, accountID, ws1ID, ws1Token)
}

func testAccCheckCloudflareDevicePostureIntegrationDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_device_posture_integration" {
			continue
		}

		_, err := client.DevicePostureIntegration(context.Background(), rs.Primary.Attributes["account_id"], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Device Posture Integration still exists")
		}
	}

	return nil
}
