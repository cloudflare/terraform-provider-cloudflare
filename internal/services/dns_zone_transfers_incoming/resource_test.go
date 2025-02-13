package dns_zone_transfers_incoming_test

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
	resource.AddTestSweepers("cloudflare_dns_zone_transfers_incoming", &resource.Sweeper{
		Name: "cloudflare_dns_zone_transfers_incoming",
		F:    testSweepCloudflareDNSZoneTransfersIncoming,
	})
}

func testSweepCloudflareDNSZoneTransfersIncoming(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	// Clean up the connections between peers and zones
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}

	incomingZone, err := client.DNS.ZoneTransfers.Incoming.Get(context.Background(), dns.ZoneTransferIncomingGetParams{ZoneID: cloudflare.F(zoneID)})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare incoming secondary zones: %s", err))
	}

	if len(incomingZone.Peers) == 0 {
		log.Print("[DEBUG] No Cloudflare peers connected to zones records to sweep")
		return nil
	}

	_, err = client.DNS.ZoneTransfers.Incoming.Update(context.Background(), dns.ZoneTransferIncomingUpdateParams{ZoneID: cloudflare.F(zoneID), Peers: cloudflare.F([]string{})})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to clear peers on incoming zone: %s", err))
	}

	return nil
}

func TestAccCloudflareDNSZoneTransfersIncoming_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	name := "cloudflare_dns_zone_transfers_incoming." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	client := acctest.SharedClient()

	peer, err := client.DNS.ZoneTransfers.Peers.New(context.Background(), dns.ZoneTransferPeerNewParams{AccountID: cloudflare.F(accountID), Name: cloudflare.F("terraform-peer")})
	if err != nil {
		tflog.Error(context.Background(), fmt.Sprintf("Failed to bootstrap Cloudflare DNS peer: %s", err))
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDNSZoneTransfersIncomingConfig(rnd, zoneName, zoneID, 300, peer.ID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", zoneName),
					resource.TestCheckResourceAttr(name, "auto_refresh_seconds", "300"),
					resource.TestCheckResourceAttr(name, "peers.0", peer.ID),
				),
			},
		},
	})
	// Delete the original peer after we are done
	_, err = client.DNS.ZoneTransfers.Peers.Delete(context.Background(), peer.ID, dns.ZoneTransferPeerDeleteParams{AccountID: cloudflare.F(accountID)})
	if err != nil {
		tflog.Error(context.Background(), fmt.Sprintf("Failed to cleanup Cloudflare DNS peer in incoming test: %s", err))
	}
}

func TestAccCloudflareDNSZoneTransfersIncoming_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	name := "cloudflare_dns_zone_transfers_incoming." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	client := acctest.SharedClient()

	peer, err := client.DNS.ZoneTransfers.Peers.New(context.Background(), dns.ZoneTransferPeerNewParams{AccountID: cloudflare.F(accountID), Name: cloudflare.F("terraform-peer")})
	if err != nil {
		tflog.Error(context.Background(), fmt.Sprintf("Failed to bootstrap Cloudflare DNS peer: %s", err))
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDNSZoneTransfersIncomingConfig(rnd, zoneName, zoneID, 300, peer.ID),
			},
			{
				Config: testDNSZoneTransfersIncomingConfig(rnd, zoneName, zoneID, 500, peer.ID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", zoneName),
					resource.TestCheckResourceAttr(name, "auto_refresh_seconds", "500"),
					resource.TestCheckResourceAttr(name, "peers.0", peer.ID),
				),
			},
		},
	})
	// Delete the original peer after we are done
	_, err = client.DNS.ZoneTransfers.Peers.Delete(context.Background(), peer.ID, dns.ZoneTransferPeerDeleteParams{AccountID: cloudflare.F(accountID)})
	if err != nil {
		tflog.Error(context.Background(), fmt.Sprintf("Failed to cleanup Cloudflare DNS peer in incoming test: %s", err))
	}
}

func testDNSZoneTransfersIncomingConfig(resourceID, zoneName, zoneID string, autoRefreshSeconds int, peers string) string {
	return acctest.LoadTestCase("incoming.tf", resourceID, zoneID, autoRefreshSeconds, zoneName, peers)
}
