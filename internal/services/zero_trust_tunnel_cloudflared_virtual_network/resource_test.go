package zero_trust_tunnel_cloudflared_virtual_network_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
		context.Background(), zero_trust.NetworkVirtualNetworkListParams{
			AccountID: cloudflare.F(accountID),
		})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Tunnel Virtual Networks: %s", err))
		return err
	}

	if len(tunnelVirtualNetworks.Result) == 0 {
		log.Print("[DEBUG] No Cloudflare Tunnel Virtual Networks to sweep")
		return nil
	}

	for _, vnet := range tunnelVirtualNetworks.Result {
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Tunnel Virtual Network %s", vnet.ID))
		//nolint:errcheck
		client.ZeroTrust.Networks.VirtualNetworks.Delete(
			context.Background(), vnet.ID, zero_trust.NetworkVirtualNetworkDeleteParams{
				AccountID: cloudflare.F(accountID),
			})
	}

	return nil
}

func testAccCheckCloudflareTunnelVirtualNetworkDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_tunnel_cloudflared_virtual_network" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		_, err := client.ZeroTrust.Networks.VirtualNetworks.Get(context.Background(), rs.Primary.ID, zero_trust.NetworkVirtualNetworkGetParams{
			AccountID: cloudflare.F(accountID),
		})
		if err == nil {
			return fmt.Errorf("tunnel virtual network still exists")
		}
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
		CheckDestroy:             testAccCheckCloudflareTunnelVirtualNetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTunnelVirtualNetworkSimple(rnd, rnd, accountID, rnd, false),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("comment"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("is_default_network"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("comment"), knownvalue.StringExact(rnd+"-updated")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("is_default_network"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
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

func TestAccCloudflareTunnelVirtualNetwork_Minimal(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_tunnel_cloudflared_virtual_network.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTunnelVirtualNetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTunnelVirtualNetworkMinimal(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("comment"), knownvalue.StringExact("")), // Default value
					statecheck.ExpectKnownValue(name, tfjsonpath.New("is_default_network"), knownvalue.Bool(false)), // Default value
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("deleted_at"), knownvalue.Null()),
				},
			},
			{
				ResourceName:        name,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}


func testAccCloudflareTunnelVirtualNetworkSimple(ID, comment, accountID, name string, isDefault bool) string {
	return acctest.LoadTestCase("tunnelvirtualnetworksimple.tf", ID, comment, accountID, name, isDefault)
}

func testAccCloudflareTunnelVirtualNetworkMinimal(name, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_tunnel_cloudflared_virtual_network" "%s" {
	account_id = "%s"
	name       = "%s"
}`, name, accountID, name)
}


