package zero_trust_tunnel_cloudflared_virtual_network_test

import (
	"context"
	"errors"
	"fmt"
	"log"
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

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_tunnel_cloudflared_virtual_network", &resource.Sweeper{
		Name: "cloudflare_zero_trust_tunnel_cloudflared_virtual_network",
		F:    testSweepCloudflareTunnelVirtualNetwork,
	})
}

func testSweepCloudflareTunnelVirtualNetwork(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return errors.New("CLOUDFLARE_ACCOUNT_ID must be set")
	}

	tunnelVirtualNetworks, err := client.ListTunnelVirtualNetworks(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.TunnelVirtualNetworksListParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Tunnel Virtual Networks: %s", err))
	}

	if len(tunnelVirtualNetworks) == 0 {
		log.Print("[DEBUG] No Cloudflare Tunnel Virtual Networks to sweep")
		return nil
	}

	for _, vnet := range tunnelVirtualNetworks {
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Tunnel Virtual Network %s", vnet.ID))
		//nolint:errcheck
		client.DeleteTunnelVirtualNetwork(context.Background(), cloudflare.AccountIdentifier(accountID), vnet.ID)
	}

	return nil
}

func TestAccCloudflareTunnelVirtualNetwork_Exists(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared_virtual_network.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var TunnelVirtualNetwork cloudflare.TunnelVirtualNetwork

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTunnelVirtualNetworkSimple(rnd, rnd, accountID, rnd, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareTunnelVirtualNetworkExists(name, &TunnelVirtualNetwork),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "comment", rnd),
					resource.TestCheckResourceAttr(name, "is_default_network", "false"),
				),
			},
		},
	})
}

func testAccCheckCloudflareTunnelVirtualNetworkExists(name string, virtualNetwork *cloudflare.TunnelVirtualNetwork) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Tunnel Virtual Network is set")
		}

		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		foundTunnelVirtualNetworks, err := client.ListTunnelVirtualNetworks(context.Background(), cloudflare.AccountIdentifier(rs.Primary.Attributes[consts.AccountIDSchemaKey]), cloudflare.TunnelVirtualNetworksListParams{
			IsDeleted: cloudflare.BoolPtr(false),
			ID:        rs.Primary.ID,
		})

		if err != nil {
			return err
		}

		*virtualNetwork = foundTunnelVirtualNetworks[0]

		return nil
	}
}

func TestAccCloudflareTunnelVirtualNetwork_UpdateComment(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared_virtual_network.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var TunnelVirtualNetwork cloudflare.TunnelVirtualNetwork

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTunnelVirtualNetworkSimple(rnd, rnd, accountID, rnd, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareTunnelVirtualNetworkExists(name, &TunnelVirtualNetwork),
					resource.TestCheckResourceAttr(name, "comment", rnd),
				),
			},
			{
				Config: testAccCloudflareTunnelVirtualNetworkSimple(rnd, rnd+"-updated", accountID, rnd, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareTunnelVirtualNetworkExists(name, &TunnelVirtualNetwork),
					resource.TestCheckResourceAttr(name, "comment", rnd+"-updated"),
				),
			},
		},
	})
}

func testAccCloudflareTunnelVirtualNetworkSimple(ID, comment, accountID, name string, isDefault bool) string {
	return acctest.LoadTestCase("tunnelvirtualnetworksimple.tf", ID, comment, accountID, name, isDefault)
}
