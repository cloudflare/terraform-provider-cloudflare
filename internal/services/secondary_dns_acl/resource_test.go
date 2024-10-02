package secondary_dns_acl_test

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
	resource.AddTestSweepers("cloudflare_secondary_dns_acl", &resource.Sweeper{
		Name: "cloudflare_secondary_dns_acl",
		F:    testSweepCloudflareSecondaryDNSACL,
	})
}

func testSweepCloudflareSecondaryDNSACL(r string) error {
	ctx := context.Background()
	client := acctest.SharedV2Client()

	// Clean up the account level acls
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return errors.New("CLOUDFLARE_ACCOUNT_ID must be set")
	}

	acls, err := client.SecondaryDNS.ACLs.List(context.Background(), secondary_dns.ACLListParams{AccountID: cloudflare.F(accountID)})
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
		client.SecondaryDNS.ACLs.Delete(context.TODO(), acl.ID, secondary_dns.ACLDeleteParams{AccountID: cloudflare.F(accountID)})
	}

	return nil
}

func TestAccCloudflareSecondaryDNSACL_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_secondary_dns_acl." + rnd
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
	name := "cloudflare_secondary_dns_acl." + rnd
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
