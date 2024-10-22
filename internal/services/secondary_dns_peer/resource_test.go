package secondary_dns_peer_test

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
	resource.AddTestSweepers("cloudflare_secondary_dns_peer", &resource.Sweeper{
		Name: "cloudflare_secondary_dns_peer",
		F:    testSweepCloudflareSecondaryDNSPeer,
	})
}

func testSweepCloudflareSecondaryDNSPeer(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	// Clean up the account level peers
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return errors.New("CLOUDFLARE_ACCOUNT_ID must be set")
	}

	peers, err := client.SecondaryDNS.Peers.List(context.Background(), secondary_dns.PeerListParams{AccountID: cloudflare.F(accountID)})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare peers: %s", err))
	}

	if len(peers.Result) == 0 {
		log.Print("[DEBUG] No Cloudflare peers records to sweep")
		return nil
	}

	for _, peer := range peers.Result {
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare peer ID: %s", peer.ID))
		//nolint:errcheck
		client.SecondaryDNS.Peers.Delete(context.TODO(), peer.ID, secondary_dns.PeerDeleteParams{AccountID: cloudflare.F(accountID)})
	}

	return nil
}

func TestAccCloudflareSecondaryDNSPeer_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_secondary_dns_peer." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	client := acctest.SharedClient()

	// Make a new TSIG that we can connect our peer to
	tsigObj := secondary_dns.TSIGParam{
		Algo: cloudflare.F("hmac-sha512."),
		Name: cloudflare.F("terraform-tsig"),
	}
	tsig, err := client.SecondaryDNS.TSIGs.New(context.Background(), secondary_dns.TSIGNewParams{AccountID: cloudflare.F(accountID), TSIG: tsigObj})
	if err != nil {
		tflog.Error(context.Background(), fmt.Sprintf("Failed to bootstrap Cloudflare DNS tsig: %s", err))
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testSecondaryDNSPeerConfig(rnd, accountID, rnd, "1.2.3.4", tsig.ID, 53, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "ip", "1.2.3.4"),
					resource.TestCheckResourceAttr(name, "ixfr_enable", "true"),
					resource.TestCheckResourceAttr(name, "port", "53"),
					resource.TestCheckResourceAttr(name, "tsig_id", tsig.ID),
				),
			},
		},
	})

	// Delete the original tsig after we are done
	_, err = client.SecondaryDNS.TSIGs.Delete(context.Background(), tsig.ID, secondary_dns.TSIGDeleteParams{AccountID: cloudflare.F(accountID)})
	if err != nil {
		tflog.Error(context.Background(), fmt.Sprintf("Failed to cleanup Cloudflare DNS tsig in peer test: %s", err))
	}
}

func TestAccCloudflareSecondaryDNSPeer_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_secondary_dns_peer." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testSecondaryDNSPeerNoTsigConfig(rnd, accountID, rnd, "1.2.3.4", 53),
			},
			{
				Config: testSecondaryDNSPeerNoTsigConfig(rnd, accountID, rnd, "1.2.3.5", 53),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "ip", "1.2.3.5"),
					resource.TestCheckResourceAttr(name, "port", "53"),
				),
			},
		},
	})
}

func testSecondaryDNSPeerConfig(resourceID, accountID, name, ip, tsigID string, port int, ixfrEnable bool) string {
	return acctest.LoadTestCase("peer.tf", resourceID, accountID, name, ip, ixfrEnable, port, tsigID)
}

func testSecondaryDNSPeerNoTsigConfig(resourceID, accountID, name, ip string, port int) string {
	return acctest.LoadTestCase("peer_no_tsig.tf", resourceID, accountID, name, ip, port)
}
