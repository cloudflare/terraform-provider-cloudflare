package cloudflare

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccCloudflareArgoTunnelCreate(t *testing.T) {
	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_argo_tunnel.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareArgoTunnelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareArgoTunnelBasic(accID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "secret", "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="),
				),
			},
		},
	})
}

func testAccCheckCloudflareArgoTunnelBasic(accID, name string) string {
	return fmt.Sprintf(`
	resource "cloudflare_argo_tunnel" "%[2]s" {
		account_id = "%[1]s"
		name       = "%[2]s"
		secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
	}`, accID, name)
}

func testAccCheckCloudflareArgoTunnelDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_argo_tunnel" {
			continue
		}

		accountID := rs.Primary.Attributes["account_id"]
		tunnelID := rs.Primary.ID
		client := testAccProvider.Meta().(*cloudflare.API)
		tunnel, err := client.ArgoTunnel(context.Background(), accountID, tunnelID)

		if err != nil {
			return err
		}

		if tunnel.DeletedAt == nil {
			return fmt.Errorf("argo tunnel with ID %s still exists", tunnel.ID)
		}

	}

	return nil
}
