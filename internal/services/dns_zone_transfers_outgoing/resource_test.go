package dns_zone_transfers_outgoing_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func init() {
	resource.AddTestSweepers("cloudflare_dns_zone_transfers_outgoing", &resource.Sweeper{
		Name: "cloudflare_dns_zone_transfers_outgoing",
		F:    testSweepCloudflareSecondaryDNSOutgoing,
	})
}

func testSweepCloudflareSecondaryDNSOutgoing(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	// Clean up the connections between peers and zones
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}

	outgoingZone, err := client.DNS.ZoneTransfers.Outgoing.Get(context.Background(), dns.ZoneTransferOutgoingGetParams{ZoneID: cloudflare.F(zoneID)})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare outgoing secondary zones: %s", err))
	}

	if len(outgoingZone.Peers) == 0 {
		log.Print("[DEBUG] No Cloudflare peers connected to zones records to sweep")
		return nil
	}

	_, err = client.DNS.ZoneTransfers.Outgoing.Update(context.Background(), dns.ZoneTransferOutgoingUpdateParams{ZoneID: cloudflare.F(zoneID), Peers: cloudflare.F([]string{})})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to clear peers on outgoing zone: %s", err))
	}

	return nil
}

func TestAccCloudflareSecondaryDNSOutgoing_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	name := "cloudflare_dns_zone_transfers_outgoing." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testSecondaryDNSOutgoingConfig(rnd, zoneName, zoneID, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", zoneName),
				),
			},
		},
	})
}

func TestAccCloudflareSecondaryDNSOutgoing_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	name := "cloudflare_dns_zone_transfers_outgoing." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testSecondaryDNSOutgoingConfig(rnd, zoneName, zoneID, accountID),
			},
			{
				Config: testSecondaryDNSOutgoingConfig(rnd, zoneName, zoneID, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", zoneName),
				),
			},
		},
	})
}

func testSecondaryDNSOutgoingConfig(resourceID, zoneName, zoneID string, accountID string) string {
	return acctest.LoadTestCase("outgoing.tf", resourceID, zoneID, zoneName, accountID)
}
