package dns_firewall_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareDNSFirewall_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_dns_firewall." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDNSFirewallConfig(rnd, accountID, rnd, "1.2.3.4", "1000"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "upstream_ips.#", "1"),
					resource.TestCheckResourceAttr(name, "upstream_ips.0", "1.2.3.4"),
					resource.TestCheckResourceAttr(name, "dns_firewall_ips.#", "4"),
					resource.TestCheckResourceAttr(name, "ratelimit", "1000"),
					resource.TestCheckResourceAttr(name, "attack_mitigation.enabled", "false"),
				),
			},
		},
	})
}

func TestAccCloudflareDNSFirewall_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_dns_firewall." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDNSFirewallConfig(rnd, accountID, rnd, "1.2.3.4", "1000"),
			},
			{
				Config: testDNSFirewallConfigWithAttackMitigation(rnd, accountID, rnd, "2.3.4.5", "2000"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "upstream_ips.#", "1"),
					resource.TestCheckResourceAttr(name, "upstream_ips.0", "2.3.4.5"),
					resource.TestCheckResourceAttr(name, "dns_firewall_ips.#", "4"),
					resource.TestCheckResourceAttr(name, "ratelimit", "2000"),
					resource.TestCheckResourceAttr(name, "attack_mitigation.enabled", "true"),
					resource.TestCheckResourceAttr(name, "attack_mitigation.only_when_upstream_unhealthy", "true"),
				),
			},
		},
	})
}

func testDNSFirewallConfig(resourceID, accountID, clusterName, upstreamIP, ratelimit string) string {
	return acctest.LoadTestCase("cluster.tf", resourceID, accountID, clusterName, upstreamIP, ratelimit)
}

func testDNSFirewallConfigWithAttackMitigation(resourceID, accountID, clusterName, upstreamIP, ratelimit string) string {
	return acctest.LoadTestCase("cluster_with_attack_mitigation.tf", resourceID, accountID, clusterName, upstreamIP, ratelimit)
}
