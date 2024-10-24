package zero_trust_tunnel_cloudflared_test

import (
	"context"
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

func TestAccCloudflareTunnelCreate_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Argo Tunnel
	// endpoint does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTunnelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareTunnelBasic(accID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "tunnel_secret", "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="),
				),
			},
		},
	})
}

func testAccCheckCloudflareTunnelBasic(accID, name string) string {
	return acctest.LoadTestCase("tunnelbasic.tf", accID, name)
}

func TestAccCloudflareTunnelCreate_Managed(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Argo Tunnel
	// endpoint does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTunnelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareTunnelManaged(accID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "tunnel_secret", "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="),
					resource.TestCheckResourceAttr(name, "config_src", "cloudflare"),
				),
			},
		},
	})
}

func testAccCheckCloudflareTunnelManaged(accID, name string) string {
	return acctest.LoadTestCase("tunnelmanaged.tf", accID, name)
}

func TestAccCloudflareTunnelCreate_Unmanaged(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Argo Tunnel
	// endpoint does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTunnelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareTunnelUnmanaged(accID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "tunnel_secret", "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="),
					resource.TestCheckResourceAttr(name, "config_src", "local"),
				),
			},
		},
	})
}

func testAccCheckCloudflareTunnelUnmanaged(accID, name string) string {
	return acctest.LoadTestCase("tunnelunmanaged.tf", accID, name)
}

func testAccCheckCloudflareTunnelDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_tunnel_cloudflared" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		tunnelID := rs.Primary.ID
		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
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
