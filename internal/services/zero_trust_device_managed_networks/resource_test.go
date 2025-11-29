package zero_trust_device_managed_networks_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_device_managed_networks", &resource.Sweeper{
		Name: "cloudflare_zero_trust_device_managed_networks",
		F:    testSweepCloudflareZeroTrustDeviceManagedNetworks,
	})
}

func testSweepCloudflareZeroTrustDeviceManagedNetworks(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	networksResp, err := client.ZeroTrust.Devices.Networks.List(
		ctx,
		zero_trust.DeviceNetworkListParams{
			AccountID: cloudflare.F(accountID),
		},
	)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Zero Trust Device Managed Networks: %s", err))
		return err
	}

	for _, network := range networksResp.Result {
		if !utils.ShouldSweepResource(network.Name) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting Zero Trust Device Managed Network: %s", network.NetworkID))
		_, err := client.ZeroTrust.Devices.Networks.Delete(
			ctx,
			network.NetworkID,
			zero_trust.DeviceNetworkDeleteParams{
				AccountID: cloudflare.F(accountID),
			},
		)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete Zero Trust Device Managed Network %s: %s", network.NetworkID, err))
		}
	}

	return nil
}

func TestAccCloudflareDeviceManagedNetworks(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_managed_networks.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDeviceManagedNetworks(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "tls"),
					resource.TestCheckResourceAttr(name, "config.tls_sockaddr", "foobar:1234"),
					resource.TestCheckResourceAttr(name, "config.sha256", "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"),
				),
			},
		},
	})
}

func testAccCloudflareDeviceManagedNetworks(accountID, rnd string) string {
	return acctest.LoadTestCase("devicemanagednetworks.tf", rnd, accountID)
}
