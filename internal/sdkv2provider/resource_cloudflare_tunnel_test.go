package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareTunnelCreate_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Argo Tunnel
	// endpoint does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared.%s", rnd)

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
	resource "cloudflare_zero_trust_tunnel_cloudflared" "%[2]s" {
		account_id = "%[1]s"
		name       = "%[2]s"
		secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
	}`, accID, name)
}

func TestAccCloudflareTunnelCreate_Managed(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Argo Tunnel
	// endpoint does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareTunnelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareTunnelManaged(accID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "secret", "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="),
					resource.TestMatchResourceAttr(name, "cname", regexp.MustCompile(".*\\.cfargotunnel\\.com")),
					resource.TestCheckResourceAttr(name, "config_src", "cloudflare"),
				),
			},
		},
	})
}

func testAccCheckCloudflareTunnelManaged(accID, name string) string {
	return fmt.Sprintf(`
	resource "cloudflare_zero_trust_tunnel_cloudflared" "%[2]s" {
		account_id = "%[1]s"
		name       = "%[2]s"
		secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
		config_src = "cloudflare"
	}`, accID, name)
}

func TestAccCloudflareTunnelCreate_Unmanaged(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Argo Tunnel
	// endpoint does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareTunnelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareTunnelUnmanaged(accID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "secret", "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="),
					resource.TestMatchResourceAttr(name, "cname", regexp.MustCompile(".*\\.cfargotunnel\\.com")),
					resource.TestCheckResourceAttr(name, "config_src", "local"),
				),
			},
		},
	})
}

func testAccCheckCloudflareTunnelUnmanaged(accID, name string) string {
	return fmt.Sprintf(`
	resource "cloudflare_zero_trust_tunnel_cloudflared" "%[2]s" {
		account_id = "%[1]s"
		name       = "%[2]s"
		secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
		config_src = "local"
	}`, accID, name)
}

func testAccCheckCloudflareTunnelDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_tunnel_cloudflared" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		tunnelID := rs.Primary.ID
		client := testAccProvider.Meta().(*cloudflare.API)
		tunnel, err := client.GetTunnel(context.Background(), cloudflare.AccountIdentifier(accountID), tunnelID)

		if err != nil {
			return err
		}

		if tunnel.DeletedAt == nil {
			return fmt.Errorf("argo tunnel with ID %s still exists", tunnel.ID)
		}
	}

	return nil
}
