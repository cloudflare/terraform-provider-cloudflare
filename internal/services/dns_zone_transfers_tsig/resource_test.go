package dns_zone_transfers_tsig_test

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
	resource.AddTestSweepers("cloudflare_dns_zone_transfers_tsig", &resource.Sweeper{
		Name: "cloudflare_dns_zone_transfers_tsig",
		F:    testSweepCloudflareSecondaryDNSTSIG,
	})
}

func testSweepCloudflareSecondaryDNSTSIG(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	// Clean up the account level tsigs
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return errors.New("CLOUDFLARE_ACCOUNT_ID must be set")
	}

	tsigs, err := client.DNS.ZoneTransfers.TSIGs.List(context.Background(), dns.ZoneTransferTSIGListParams{AccountID: cloudflare.F(accountID)})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare tsigs: %s", err))
	}

	if len(tsigs.Result) == 0 {
		log.Print("[DEBUG] No Cloudflare tsigs records to sweep")
		return nil
	}

	for _, tsig := range tsigs.Result {
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare tsig ID: %s", tsig.ID))
		//nolint:errcheck
		client.DNS.ZoneTransfers.TSIGs.Delete(context.TODO(), tsig.ID, dns.ZoneTransferTSIGDeleteParams{AccountID: cloudflare.F(accountID)})
	}

	return nil
}

func TestAccCloudflareSecondaryDNSTSIG_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_dns_zone_transfers_tsig." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testSecondaryDNSTSIGConfig(rnd, accountID, rnd, "hmac-sha512.", "caf79a7804b04337c9c66ccd7bef9190a1e1679b5dd03d8aa10f7ad45e1a9dab92b417896c15d4d007c7c14194538d2a5d0feffdecc5a7f0e1c570cfa700837c"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd+"."),
					resource.TestCheckResourceAttr(name, "algo", "hmac-sha512."),
					resource.TestCheckResourceAttr(name, "secret", "caf79a7804b04337c9c66ccd7bef9190a1e1679b5dd03d8aa10f7ad45e1a9dab92b417896c15d4d007c7c14194538d2a5d0feffdecc5a7f0e1c570cfa700837c"),
				),
			},
		},
	})
}

func TestAccCloudflareSecondaryDNSTSIG_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_dns_zone_transfers_tsig." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testSecondaryDNSTSIGConfig(rnd, accountID, rnd, "hmac-sha512.", "baf79a7804b04337c9c66ccd7bef9190a1e1679b5dd03d8aa10f7ad45e1a9dab92b417896c15d4d007c7c14194538d2a5d0feffdecc5a7f0e1c570cfa700837c"),
			},
			{
				Config: testSecondaryDNSTSIGConfig(rnd, accountID, rnd, "hmac-sha512.", "caf79a7804b04337c9c66ccd7bef9190a1e1679b5dd03d8aa10f7ad45e1a9dab92b417896c15d4d007c7c14194538d2a5d0feffdecc5a7f0e1c570cfa700837c"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd+"."),
					resource.TestCheckResourceAttr(name, "algo", "hmac-sha512."),
					resource.TestCheckResourceAttr(name, "secret", "caf79a7804b04337c9c66ccd7bef9190a1e1679b5dd03d8aa10f7ad45e1a9dab92b417896c15d4d007c7c14194538d2a5d0feffdecc5a7f0e1c570cfa700837c"),
				),
			},
		},
	})
}

func testSecondaryDNSTSIGConfig(resourceID, accountID, name, algo, secret string) string {
	return acctest.LoadTestCase("tsig.tf", resourceID, accountID, name+".", algo, secret)
}
