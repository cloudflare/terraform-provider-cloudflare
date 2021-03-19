package cloudflare

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func init() {
	resource.AddTestSweepers("cloudflare_firewall_rule", &resource.Sweeper{
		Name: "cloudflare_firewall_rule",
		F:    testSweepCloudflareFirewallRuleSweeper,
	})
}

func TestAccCloudflareFirewallRulesMatchPaused(t *testing.T) {
	t.Parallel()
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_firewall_rules.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareFirewallRulesConfigMatchPaused(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "filter.0.paused", "true"),
					resource.TestCheckResourceAttr(name, "firewall_rules.#", "1"),
				),
			},
		},
	})
}

func TestAccCloudflareFirewallRulesMatchPriorityFilter(t *testing.T) {
	t.Parallel()
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_zones.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareFirewallRulesConfigMatchPriority(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareFirewallRulesDataSourceID(name),
					resource.TestCheckResourceAttr(name, "filter.0.priority", "2"),
					resource.TestCheckResourceAttr(name, "filter.0.match_type", "gte"),
					resource.TestCheckResourceAttr(name, "firewall_rules.#", "2"),
				),
			},
		},
	})
}

func testAccCheckCloudflareFirewallRulesDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		all := s.RootModule().Resources
		rs, ok := all[n]
		if !ok {
			return fmt.Errorf("Can't find firewall rule data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Snapshot firewall rule source ID not set")
		}
		return nil
	}
}

func testAccCloudflareFirewallRulesConfigMatchPaused(rnd string) string {
	return fmt.Sprintf(`
data "cloudflare_firewall_rules" "%[2]s" {
  filter {
    paused = "${cloudflare_firewall_rule.baa_com.paused}"
  }
}

%[1]s
`, testRules, rnd)
}

func testAccCloudflareFirewallRulesConfigMatchPriority(rnd string) string {
	return fmt.Sprintf(`
data "cloudflare_firewall_rules" "%[2]s" {
  filter {
    priority   = 2
	match_type = "gte"
    // This is an ordering fix to ensure that the test suite doesn't assert
    // state before all the resources are available.
    paused = "${cloudflare_firewall_rule.baa_net.paused}"
  }
}

%[1]s
`, testRules, rnd)
}

const testRules = `resource "cloudflare_firewall_rule" "baa_com" {
  zone_id  = "0da42c8d2132a9ddaf714f9e7c920711"
  paused   = true
  action   = "block"
  priority = 1
}

resource "cloudflare_firewall_rule" "baa_org" {
  zone_id  = "0da42c8d2132a9ddaf714f9e7c920711"
  paused   = false
  action   = "block"
  priority = 2
}

resource "cloudflare_firewall_rule" "baa_net" {
  zone_id  = "0da42c8d2132a9ddaf714f9e7c920711"
  paused   = false
  action   = "block"
  priority = 3
}

resource "cloudflare_firewall_rule" "foo_net" {
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
  paused     = false
  action     = "block"
  depends_on = ["cloudflare_firewall_rule.baa_net", "cloudflare_firewall_rule.baa_org", "cloudflare_firewall_rule.baa_com"]
}`
