package zero_trust_tunnel_cloudflared_virtual_network_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_tunnel_cloudflared_virtual_network", &resource.Sweeper{
		Name: "cloudflare_zero_trust_tunnel_cloudflared_virtual_network",
		F:    testSweepCloudflareTunnelVirtualNetwork,
	})
}

func testSweepCloudflareTunnelVirtualNetwork(r string) error {
	ctx := context.Background()

	client := acctest.SharedClient() // TODO(terraform): replace with SharedV2Clent

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return errors.New("CLOUDFLARE_ACCOUNT_ID must be set")
	}

	tunnelVirtualNetworks, err := client.ZeroTrust.Networks.VirtualNetworks.List(
		context.Background(), zero_trust.NetworkVirtualNetworkListParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Tunnel Virtual Networks: %s", err))
	}

	if len(tunnelVirtualNetworks.Result) == 0 {
		log.Print("[DEBUG] No Cloudflare Tunnel Virtual Networks to sweep")
		return nil
	}

	for _, vnet := range tunnelVirtualNetworks.Result {
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Tunnel Virtual Network %s", vnet.ID))
		//nolint:errcheck
		client.ZeroTrust.Networks.VirtualNetworks.Delete(
			context.Background(), vnet.ID, zero_trust.NetworkVirtualNetworkDeleteParams{})
	}

	return nil
}

func TestAccCloudflareTunnelVirtualNetwork_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared_virtual_network.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

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
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "comment", rnd),
					resource.TestCheckResourceAttr(name, "is_default_network", "false"),
				),
			},
			// Update
			{
				Config: testAccCloudflareTunnelVirtualNetworkSimple(rnd, rnd+"-updated", accountID, rnd, false),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(
							name,
							tfjsonpath.New("comment"),
							knownvalue.StringExact(rnd+"-updated"),
						),
						plancheck.ExpectKnownValue(
							name,
							tfjsonpath.New("name"),
							knownvalue.StringExact(rnd),
						),
						plancheck.ExpectKnownValue(
							name,
							tfjsonpath.New("is_default_network"),
							knownvalue.Bool(false),
						),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "comment", rnd+"-updated"),
					resource.TestCheckResourceAttr(name, "is_default_network", "false"),
				),
			},
			// Re-applying same change does not produce drift
			{
				Config: testAccCloudflareTunnelVirtualNetworkSimple(rnd, rnd+"-updated", accountID, rnd, false),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Import
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func testAccCloudflareTunnelVirtualNetworkSimple(ID, comment, accountID, name string, isDefault bool) string {
	return acctest.LoadTestCase("tunnelvirtualnetworksimple.tf", ID, comment, accountID, name, isDefault)
}
