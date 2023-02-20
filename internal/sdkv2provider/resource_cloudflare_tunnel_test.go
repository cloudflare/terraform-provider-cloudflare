package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudflareTunnelCreate_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Argo Tunnel
	// endpoint does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_tunnel.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareTunnelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareTunnelBasic(accID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "secret", "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="),
					resource.TestMatchResourceAttr(name, "cname", regexp.MustCompile(".*\\.cfargotunnel\\.com")),
				),
			},
		},
	})
}

func testAccCheckCloudflareTunnelBasic(accID, name string) string {
	return fmt.Sprintf(`
	resource "cloudflare_tunnel" "%[2]s" {
		account_id = "%[1]s"
		name       = "%[2]s"
		secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
	}`, accID, name)
}

func testAccCheckCloudflareTunnelDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_tunnel" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		tunnelID := rs.Primary.ID
		client := testAccProvider.Meta().(*cloudflare.API)
		tunnel, err := client.Tunnel(context.Background(), cloudflare.AccountIdentifier(accountID), tunnelID)

		if err != nil {
			return err
		}

		if tunnel.DeletedAt == nil {
			return fmt.Errorf("argo tunnel with ID %s still exists", tunnel.ID)
		}
	}

	return nil
}
