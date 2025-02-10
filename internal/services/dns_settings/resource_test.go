package dns_settings_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDNSSettings_UpdateAccount(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_dns_settings." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDNSSettingsConfigAccount(rnd, accountID, false, false, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "zone_defaults.flatten_all_cnames", "false"),
					resource.TestCheckResourceAttr(name, "zone_defaults.foundation_dns", "false"),
					resource.TestCheckResourceAttr(name, "zone_defaults.multi_provider", "false"),
					resource.TestCheckResourceAttr(name, "zone_defaults.nameservers.type", "cloudflare.standard"),
					resource.TestCheckResourceAttr(name, "zone_defaults.ns_ttl", "86400"),
					resource.TestCheckResourceAttr(name, "zone_defaults.secondary_overrides", "false"),
					resource.TestCheckResourceAttr(name, "zone_defaults.soa.expire", "604800"),
					resource.TestCheckResourceAttr(name, "zone_defaults.soa.min_ttl", "1800"),
					resource.TestCheckResourceAttr(name, "zone_defaults.soa.mname", "kristina.ns.cloudflare.com"),
					resource.TestCheckResourceAttr(name, "zone_defaults.soa.refresh", "10000"),
					resource.TestCheckResourceAttr(name, "zone_defaults.soa.retry", "2400"),
					resource.TestCheckResourceAttr(name, "zone_defaults.soa.rname", "admin.example.com"),
					resource.TestCheckResourceAttr(name, "zone_defaults.soa.ttl", "3600"),
					resource.TestCheckResourceAttr(name, "zone_defaults.zone_mode", "standard"),
				),
			},
		},
	})
}

func TestAccDNSSettings_UpdateZone(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_dns_settings." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDNSSettingsConfigZone(rnd, zoneID, false, false, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttr(name, "zone_defaults.flatten_all_cnames", "false"),
					resource.TestCheckResourceAttr(name, "zone_defaults.foundation_dns", "false"),
					resource.TestCheckResourceAttr(name, "zone_defaults.multi_provider", "false"),
					resource.TestCheckResourceAttr(name, "zone_defaults.nameservers.type", "cloudflare.standard"),
					resource.TestCheckResourceAttr(name, "zone_defaults.ns_ttl", "86400"),
					resource.TestCheckResourceAttr(name, "zone_defaults.secondary_overrides", "false"),
					resource.TestCheckResourceAttr(name, "zone_defaults.soa.expire", "604800"),
					resource.TestCheckResourceAttr(name, "zone_defaults.soa.min_ttl", "1800"),
					resource.TestCheckResourceAttr(name, "zone_defaults.soa.mname", "kristina.ns.cloudflare.com"),
					resource.TestCheckResourceAttr(name, "zone_defaults.soa.refresh", "10000"),
					resource.TestCheckResourceAttr(name, "zone_defaults.soa.retry", "2400"),
					resource.TestCheckResourceAttr(name, "zone_defaults.soa.rname", "admin.example.com"),
					resource.TestCheckResourceAttr(name, "zone_defaults.soa.ttl", "3600"),
					resource.TestCheckResourceAttr(name, "zone_defaults.zone_mode", "standard"),
				),
			},
		},
	})
}

func testDNSSettingsConfigAccount(resourceID, ID string, flattenCnames, foundationDNS, multiProvider bool) string {
	return acctest.LoadTestCase("account_setting.tf", resourceID, ID, flattenCnames, foundationDNS, multiProvider)
}

func testDNSSettingsConfigZone(resourceID, ID string, flattenCnames, foundationDNS, multiProvider bool) string {
	return acctest.LoadTestCase("zone_setting.tf", resourceID, ID, flattenCnames, foundationDNS, multiProvider)
}
