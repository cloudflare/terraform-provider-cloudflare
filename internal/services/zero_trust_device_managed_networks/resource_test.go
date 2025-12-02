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
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
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
	resourceName := fmt.Sprintf("cloudflare_zero_trust_device_managed_networks.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDeviceManagedNetworksDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDeviceManagedNetworks(accountID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("tls")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("tls_sockaddr"), knownvalue.StringExact("foobar:1234")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("config").AtMapKey("sha256"), knownvalue.StringExact("b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c")),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

func testAccCloudflareDeviceManagedNetworks(accountID, rnd string) string {
	return acctest.LoadTestCase("devicemanagednetworks.tf", rnd, accountID)
}

func testAccCheckCloudflareDeviceManagedNetworksDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_device_managed_networks" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		networkID := rs.Primary.ID
		_, err := client.ZeroTrust.Devices.Networks.Get(context.Background(), networkID, zero_trust.DeviceNetworkGetParams{
			AccountID: cloudflare.F(accountID),
		})
		if err == nil {
			return fmt.Errorf("device managed network still exists: %s", networkID)
		}
	}

	return nil
}
