package secondary_dns_outgoing_test

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
	resource.AddTestSweepers("cloudflare_secondary_dns_outgoing", &resource.Sweeper{
		Name: "cloudflare_secondary_dns_outgoing",
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

	outgoingZone, err := client.SecondaryDNS.Outgoing.Get(context.Background(), secondary_dns.OutgoingGetParams{ZoneID: cloudflare.F(zoneID)})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare outgoing secondary zones: %s", err))
	}

	if len(outgoingZone.Peers) == 0 {
		log.Print("[DEBUG] No Cloudflare peers connected to zones records to sweep")
		return nil
	}

	_, err = client.SecondaryDNS.Outgoing.Update(context.Background(), secondary_dns.OutgoingUpdateParams{ZoneID: cloudflare.F(zoneID), Peers: cloudflare.F([]interface{}{})})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to clear peers on outgoing zone: %s", err))
	}

	return nil
}

func TestAccCloudflareSecondaryDNSOutgoing_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	name := "cloudflare_secondary_dns_outgoing." + rnd
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
				Config: testSecondaryDNSOutgoingConfig(rnd, zoneName, zoneID, peer.ID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", zoneName),
					resource.TestCheckResourceAttr(name, "peers.0", peer.ID),
				),
			},
		},
	})
	// Delete the original peer after we are done
	_, err = client.SecondaryDNS.Peers.Delete(context.Background(), peer.ID, secondary_dns.PeerDeleteParams{AccountID: cloudflare.F(accountID)})
	if err != nil {
		tflog.Error(context.Background(), fmt.Sprintf("Failed to cleanup Cloudflare DNS peer in outgoing test: %s", err))
	}
}

func TestAccCloudflareSecondaryDNSOutgoing_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	name := "cloudflare_secondary_dns_outgoing." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	client := acctest.SharedClient()

	// Make a new peer that we can connect our zone to
	peerObj := secondary_dns.PeerParam{
		Port: cloudflare.F(float64(53)),
		Name: cloudflare.F("terraform-peer"),
		IP:   cloudflare.F("1.2.3.4"),
	}
	peer1, err := client.SecondaryDNS.Peers.New(context.Background(), secondary_dns.PeerNewParams{AccountID: cloudflare.F(accountID), Body: peerObj})
	if err != nil {
		tflog.Error(context.Background(), fmt.Sprintf("Failed to bootstrap Cloudflare DNS peer: %s", err))
	}
	peerObj.IP = cloudflare.F("1.2.3.6")
	peer2, err := client.SecondaryDNS.Peers.New(context.Background(), secondary_dns.PeerNewParams{AccountID: cloudflare.F(accountID), Body: peerObj})
	if err != nil {
		tflog.Error(context.Background(), fmt.Sprintf("Failed to bootstrap Cloudflare DNS peer: %s", err))
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testSecondaryDNSOutgoingConfig(rnd, zoneName, zoneID, peer1.ID),
			},
			{
				Config: testSecondaryDNSOutgoingConfig(rnd, zoneName, zoneID, peer2.ID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", zoneName),
					resource.TestCheckResourceAttr(name, "peers.0", peer2.ID),
				),
			},
		},
	})
	// Delete the original peers after we are done
	for _, id := range []string{peer1.ID, peer2.ID} {
		_, err = client.SecondaryDNS.Peers.Delete(context.Background(), id, secondary_dns.PeerDeleteParams{AccountID: cloudflare.F(accountID)})
		if err != nil {
			tflog.Error(context.Background(), fmt.Sprintf("Failed to cleanup Cloudflare DNS peer in outgoing test: %s", err))
		}
	}

}

func testSecondaryDNSOutgoingConfig(resourceID, zoneName, zoneID string, peers string) string {
	return acctest.LoadTestCase("outgoing.tf", resourceID, zoneID, zoneName, peers)
}
