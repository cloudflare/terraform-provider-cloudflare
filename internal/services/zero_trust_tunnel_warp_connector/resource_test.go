package zero_trust_tunnel_warp_connector_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cloudflare6 "github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_tunnel_warp_connector", &resource.Sweeper{
		Name: "cloudflare_zero_trust_tunnel_warp_connector",
		F:    testSweepCloudflareZeroTrustTunnelWARPConnector,
	})
}

func testSweepCloudflareZeroTrustTunnelWARPConnector(region string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if accountID == "" {
		tflog.Info(ctx, "Skipping WARP Connector tunnels sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	page, err := client.ZeroTrust.Tunnels.WARPConnector.List(
		ctx,
		zero_trust.TunnelWARPConnectorListParams{
			AccountID: cloudflare6.F(accountID),
		},
	)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list WARP Connector tunnels: %s", err))
		return fmt.Errorf("error listing WARP Connector tunnels for sweep: %w", err)
	}

	tunnelCount := 0
	for page != nil && len(page.Result) > 0 {
		for _, tunnel := range page.Result {
			if tunnel.ID == "" {
				tflog.Debug(ctx, fmt.Sprintf("Skipping WARP Connector tunnel with empty ID: %s", tunnel.Name))
				continue
			}

			if !tunnel.DeletedAt.IsZero() {
				tflog.Debug(ctx, fmt.Sprintf("Skipping already deleted WARP Connector tunnel: %s (%s)", tunnel.Name, tunnel.ID))
				continue
			}

			if !utils.ShouldSweepResource(tunnel.Name) {
				continue
			}

			tunnelCount++
			tflog.Info(ctx, fmt.Sprintf("Deleting WARP Connector tunnel: %s (%s) (account: %s)", tunnel.Name, tunnel.ID, accountID))

			_, err := client.ZeroTrust.Tunnels.WARPConnector.Delete(
				ctx,
				tunnel.ID,
				zero_trust.TunnelWARPConnectorDeleteParams{
					AccountID: cloudflare6.F(accountID),
				},
				option.WithQuery("cascade", "true"),
			)
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete WARP Connector tunnel %s (%s): %s", tunnel.Name, tunnel.ID, err))
			}
		}

		nextPage, err := page.GetNextPage()
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Error getting next page: %s", err))
			break
		}
		page = nextPage
	}

	tflog.Info(ctx, fmt.Sprintf("Successfully swept %d WARP Connector tunnel(s)", tunnelCount))
	return nil
}

func TestAccWARPConnectorCreateBasic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_tunnel_warp_connector.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWARPConnectorBasic(accID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "tun_type", "warp_connector"),
				),
			},
		},
	})
}

func TestAccWARPConnectorUpdateName(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_tunnel_warp_connector.%s", rnd)

	name1 := fmt.Sprintf("%s_1", rnd)
	name2 := fmt.Sprintf("%s_2", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWARPConnectorUpdateName(accID, rnd, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
					resource.TestCheckResourceAttr(resourceName, "tun_type", "warp_connector"),
				),
			},
			{
				Config: testAccCheckWARPConnectorUpdateName(accID, rnd, name2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
					resource.TestCheckResourceAttr(resourceName, "tun_type", "warp_connector"),
				),
			},
		},
	})
}

func testAccCheckWARPConnectorBasic(accID, name string) string {
	return acctest.LoadTestCase("warp_connector_basic.tf", accID, name)
}

func testAccCheckWARPConnectorUpdateName(accID, resourceName, name string) string {
	return acctest.LoadTestCase("warp_connector_update_name.tf", accID, resourceName, name)
}
