package web3_hostname_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
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
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
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
	return acctest.LoadTestCase("buildweb3hostnameconfigethereum.tf", name, zoneID, domain)
}

func buildWeb3HostnameConfigIPFS(name, zoneID, domain string) string {
	return acctest.LoadTestCase("buildweb3hostnameconfigipfs.tf", name, zoneID, domain)
}

func TestAccCloudflareWeb3HostnameEthereum(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_web3_hostname.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_web3_hostname.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
