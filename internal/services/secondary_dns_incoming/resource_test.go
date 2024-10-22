package secondary_dns_incoming_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/secondary_dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func init() {
	resource.AddTestSweepers("cloudflare_secondary_dns_incoming", &resource.Sweeper{
		Name: "cloudflare_secondary_dns_incoming",
		F:    testSweepCloudflareSecondaryDNSIncoming,
	})
}

func testSweepCloudflareSecondaryDNSIncoming(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	// Clean up the connections between peers and zones
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}

	incomingZone, err := client.SecondaryDNS.Incoming.Get(context.Background(), secondary_dns.IncomingGetParams{ZoneID: cloudflare.F(zoneID)})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare incoming secondary zones: %s", err))
	}

	if len(incomingZone.Peers) == 0 {
		log.Print("[DEBUG] No Cloudflare peers connected to zones records to sweep")
		return nil
	}

	_, err = client.SecondaryDNS.Incoming.Update(context.Background(), secondary_dns.IncomingUpdateParams{ZoneID: cloudflare.F(zoneID), Peers: cloudflare.F([]interface{}{})})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to clear peers on incoming zone: %s", err))
	}

	return nil
}

func TestAccCloudflareSecondaryDNSIncoming_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	name := "cloudflare_secondary_dns_incoming." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	client := acctest.SharedClient()

	// Make a new peer that we can connect our zone to
	peerObj := secondary_dns.PeerParam{
		Port: cloudflare.F(float64(53)),
		Name: cloudflare.F("terraform-peer"),
		IP:   cloudflare.F("1.2.3.4"),
	}
	peer, err := client.SecondaryDNS.Peers.New(context.Background(), secondary_dns.PeerNewParams{AccountID: cloudflare.F(accountID), Body: peerObj})
	if err != nil {
		tflog.Error(context.Background(), fmt.Sprintf("Failed to bootstrap Cloudflare DNS peer: %s", err))
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testSecondaryDNSIncomingConfig(rnd, zoneName, zoneID, 300, peer.ID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", zoneName),
					resource.TestCheckResourceAttr(name, "auto_refresh_seconds", "300"),
					resource.TestCheckResourceAttr(name, "peers.0", peer.ID),
				),
			},
		},
	})
	// Delete the original peer after we are done
	_, err = client.SecondaryDNS.Peers.Delete(context.Background(), peer.ID, secondary_dns.PeerDeleteParams{AccountID: cloudflare.F(accountID)})
	if err != nil {
		tflog.Error(context.Background(), fmt.Sprintf("Failed to cleanup Cloudflare DNS peer in incoming test: %s", err))
	}
}

func TestAccCloudflareSecondaryDNSIncoming_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	name := "cloudflare_secondary_dns_incoming." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	client := acctest.SharedClient()

	// Make a new peer that we can connect our zone to
	peerObj := secondary_dns.PeerParam{
		Port: cloudflare.F(float64(53)),
		Name: cloudflare.F("terraform-peer"),
		IP:   cloudflare.F("1.2.3.4"),
	}
	peer, err := client.SecondaryDNS.Peers.New(context.Background(), secondary_dns.PeerNewParams{AccountID: cloudflare.F(accountID), Body: peerObj})
	if err != nil {
		tflog.Error(context.Background(), fmt.Sprintf("Failed to bootstrap Cloudflare DNS peer: %s", err))
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testSecondaryDNSIncomingConfig(rnd, zoneName, zoneID, 300, peer.ID),
			},
			{
				Config: testSecondaryDNSIncomingConfig(rnd, zoneName, zoneID, 500, peer.ID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", zoneName),
					resource.TestCheckResourceAttr(name, "auto_refresh_seconds", "500"),
					resource.TestCheckResourceAttr(name, "peers.0", peer.ID),
				),
			},
		},
	})
	// Delete the original peer after we are done
	_, err = client.SecondaryDNS.Peers.Delete(context.Background(), peer.ID, secondary_dns.PeerDeleteParams{AccountID: cloudflare.F(accountID)})
	if err != nil {
		tflog.Error(context.Background(), fmt.Sprintf("Failed to cleanup Cloudflare DNS peer in incoming test: %s", err))
	}
}

func testSecondaryDNSIncomingConfig(resourceID, zoneName, zoneID string, autoRefreshSeconds int, peers string) string {
	return acctest.LoadTestCase("incoming.tf", resourceID, zoneID, autoRefreshSeconds, zoneName, peers)
}
