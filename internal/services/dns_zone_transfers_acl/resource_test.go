package dns_zone_transfers_acl_test

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
	resource.AddTestSweepers("cloudflare_dns_zone_transfers_acl", &resource.Sweeper{
		Name: "cloudflare_dns_zone_transfers_acl",
		F:    testSweepCloudflareSecondaryDNSACL,
	})
}

func testSweepCloudflareSecondaryDNSACL(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	// Clean up the account level acls
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return errors.New("CLOUDFLARE_ACCOUNT_ID must be set")
	}

	acls, err := client.DNS.ZoneTransfers.ACLs.List(context.Background(), dns.ZoneTransferACLListParams{AccountID: cloudflare.F(accountID)})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare secondary DNS ACLs: %s", err))
	}

	if len(acls.Result) == 0 {
		log.Print("[DEBUG] No Cloudflare ACLs to sweep")
		return nil
	}

	for _, acl := range acls.Result {
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare ACL ID: %s", acl.ID))
		//nolint:errcheck
		client.DNS.ZoneTransfers.ACLs.Delete(context.TODO(), acl.ID, dns.ZoneTransferACLDeleteParams{AccountID: cloudflare.F(accountID)})
	}

	return nil
}

func TestAccCloudflareSecondaryDNSACL_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_dns_zone_transfers_acl." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testSecondaryDNSACLConfig(rnd, accountID, rnd, "1.2.3.4/32"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "ip_range", "1.2.3.4/32"),
				),
			},
		},
	})
}

func TestAccCloudflareSecondaryDNSACL_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_dns_zone_transfers_acl." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testSecondaryDNSACLConfig(rnd, accountID, rnd, "1.2.3.4/32"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "ip_range", "1.2.3.4/32"),
				),
			},
			{
				Config: testSecondaryDNSACLConfig(rnd, accountID, rnd, "1.2.3.5/32"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "ip_range", "1.2.3.5/32"),
				),
			},
		},
	})
}

func testSecondaryDNSACLConfig(resourceID, accountID, name, ipRange string) string {
	return acctest.LoadTestCase("acl.tf", resourceID, accountID, name, ipRange)
}
