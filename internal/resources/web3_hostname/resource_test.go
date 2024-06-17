package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func init() {
	resource.AddTestSweepers("cloudflare_web3_hostname", &resource.Sweeper{
		Name: "cloudflare_web3_hostname",
		F:    testSweepCloudflareWeb3Hostname,
	})
}

func testSweepCloudflareWeb3Hostname(r string) error {
	ctx := context.Background()
	client, clientErr := sharedClient()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}

	hostnames, err := client.ListWeb3Hostnames(context.Background(), cloudflare.Web3HostnameListParameters{ZoneID: zoneID})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Web3 hostnames: %s", err))
	}

	if len(hostnames) == 0 {
		log.Print("[DEBUG] No Cloudflare Web3 hostnames to sweep")
		return nil
	}

	for _, hostname := range hostnames {
		tflog.Info(ctx, fmt.Sprintf("DeletingCloudflare Web3 hostname ID: %s", hostname.ID))
		//nolint:errcheck
		client.DeleteWeb3Hostname(context.Background(), cloudflare.Web3HostnameDetailsParameters{ZoneID: zoneID, Identifier: hostname.ID})
	}

	return nil
}

func buildWeb3HostnameConfigEthereum(name, zoneID, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_web3_hostname" "%[1]s" {
	zone_id = "%[2]s"
	name = "%[1]s.%[3]s"
	target = "ethereum"
	description = "test"
}
`, name, zoneID, domain)
}

func buildWeb3HostnameConfigIPFS(name, zoneID, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_web3_hostname" "%[1]s" {
	zone_id = "%[2]s"
	name = "%[1]s.%[3]s"
	target = "ipfs"
	description = "test"
	dnslink = "/ipns/onboarding.ipfs.cloudflare.com"
}
`, name, zoneID, domain)
}

func TestAccCloudflareWeb3HostnameEthereum(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_web3_hostname.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: buildWeb3HostnameConfigEthereum(rnd, zoneID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd+"."+domain),
					resource.TestCheckResourceAttr(name, "target", "ethereum"),
					resource.TestCheckResourceAttr(name, "description", "test"),
				),
			},
		},
	})
}

func TestAccCloudflareWeb3Hostname(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_web3_hostname.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: buildWeb3HostnameConfigIPFS(rnd, zoneID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd+"."+domain),
					resource.TestCheckResourceAttr(name, "target", "ipfs"),
					resource.TestCheckResourceAttr(name, "description", "test"),
					resource.TestCheckResourceAttr(name, "dnslink", "/ipns/onboarding.ipfs.cloudflare.com"),
				),
			},
		},
	})
}
