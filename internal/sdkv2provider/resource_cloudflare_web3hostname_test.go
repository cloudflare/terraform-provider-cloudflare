package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

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
