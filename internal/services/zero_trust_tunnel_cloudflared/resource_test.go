package zero_trust_tunnel_cloudflared_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	cloudflare6 "github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
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
	resource.AddTestSweepers("cloudflare_zero_trust_tunnel_cloudflared", &resource.Sweeper{
		Name: "cloudflare_zero_trust_tunnel_cloudflared",
		F:    testSweepCloudflareZeroTrustTunnelCloudflared,
	})
}

func testSweepCloudflareZeroTrustTunnelCloudflared(region string) error {
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if accountID == "" {
		log.Print("[DEBUG] CLOUDFLARE_ACCOUNT_ID not set, skipping sweep")
		return nil
	}

	ctx := context.Background()

	// List all cloudflared tunnels using v6 SDK
	page, err := client.ZeroTrust.Tunnels.Cloudflared.List(
		ctx,
		zero_trust.TunnelCloudflaredListParams{
			AccountID: cloudflare6.F(accountID),
		},
	)
	if err != nil {
		return fmt.Errorf("error listing cloudflared tunnels for sweep: %w", err)
	}

	tunnelCount := 0
	for page != nil && len(page.Result) > 0 {
		tunnelCount += len(page.Result)

		for _, tunnel := range page.Result {
			if tunnel.ID == "" {
				log.Printf("[DEBUG] Skipping cloudflared tunnel with empty ID: %s", tunnel.Name)
				continue
			}

			// Skip tunnels that are already deleted
			if !tunnel.DeletedAt.IsZero() {
				log.Printf("[DEBUG] Skipping already deleted cloudflared tunnel: %s (%s)", tunnel.Name, tunnel.ID)
				continue
			}

			log.Printf("[INFO] Deleting cloudflared tunnel: %s (%s)", tunnel.Name, tunnel.ID)

			_, err := client.ZeroTrust.Tunnels.Cloudflared.Delete(
				ctx,
				tunnel.ID,
				zero_trust.TunnelCloudflaredDeleteParams{
					AccountID: cloudflare6.F(accountID),
				},
				option.WithQuery("cascade", "true"),
			)
			if err != nil {
				log.Printf("[ERROR] Failed to delete cloudflared tunnel %s: %v", tunnel.ID, err)
				continue
			}

			log.Printf("[DEBUG] Successfully deleted cloudflared tunnel: %s (%s) with cascade", tunnel.Name, tunnel.ID)
		}

		// Get next page
		page, err = page.GetNextPage()
		if err != nil {
			log.Printf("[ERROR] Failed to get next page of tunnels: %v", err)
			break
		}
	}

	if tunnelCount == 0 {
		log.Print("[DEBUG] No cloudflared tunnels found to sweep")
	} else {
		log.Printf("[DEBUG] Found %d cloudflared tunnels to sweep", tunnelCount)
	}

	return nil
}

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

func TestAccCloudflareTunnelRotateSecret(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Argo Tunnel
	// endpoint does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared.%s", rnd)

	secretV1 := "dGhpc19pc18xX3NlY3JldF9mb3JfdGhlX2ZpcnN0"
	secretV2 := "Ml9zZWNyZXRfMl90dW5uZWxfdXBkYXRl"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTunnelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareTunnelRotateSecret(accID, rnd, secretV1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "tunnel_secret", secretV1),
				),
			},
			{
				Config: testAccCheckCloudflareTunnelRotateSecret(accID, rnd, secretV2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "tunnel_secret", secretV2),
				),
			},
		},
	})
}

func testAccCheckCloudflareTunnelRotateSecret(accID, name, secret string) string {
	return acctest.LoadTestCase("tunnelsecret.tf", accID, name, secret)
}
