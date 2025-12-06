package dns_firewall_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/dns_firewall"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_dns_firewall", &resource.Sweeper{
		Name: "cloudflare_dns_firewall",
		F: func(region string) error {
			ctx := context.Background()
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			client := acctest.SharedClient()

			if accountID == "" {
				tflog.Info(ctx, "Skipping DNS Firewall clusters sweep: CLOUDFLARE_ACCOUNT_ID not set")
				return nil
			}

			clusters, err := client.DNSFirewall.List(ctx, dns_firewall.DNSFirewallListParams{
				AccountID: cloudflare.F(accountID),
			})
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to fetch DNS Firewall clusters: %s", err))
				return fmt.Errorf("failed to fetch DNS Firewall clusters: %w", err)
			}

			if len(clusters.Result) == 0 {
				tflog.Info(ctx, "No DNS Firewall clusters to sweep")
				return nil
			}

			for _, cluster := range clusters.Result {
				if !utils.ShouldSweepResource(cluster.Name) {
					continue
				}
				tflog.Info(ctx, fmt.Sprintf("Deleting DNS Firewall cluster: %s (%s) (account: %s)", cluster.Name, cluster.ID, accountID))
				_, err := client.DNSFirewall.Delete(ctx, cluster.ID, dns_firewall.DNSFirewallDeleteParams{
					AccountID: cloudflare.F(accountID),
				})
				if err != nil {
					tflog.Error(ctx, fmt.Sprintf("Failed to delete DNS Firewall cluster %s: %s", cluster.Name, err))
					continue
				}
				tflog.Info(ctx, fmt.Sprintf("Deleted DNS Firewall cluster: %s", cluster.ID))
			}

			return nil
		},
	})
}

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
