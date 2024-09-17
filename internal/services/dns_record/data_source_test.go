package dns_record_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareRecordDataSource(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_record.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRecordDataSourceConfig(rnd, zoneID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "hostname", rnd+"."+domain),
					resource.TestCheckResourceAttr(name, "type", "A"),
					resource.TestCheckResourceAttr(name, "value", "192.0.2.0"),
					resource.TestCheckResourceAttr(name, "proxied", "false"),
					resource.TestCheckResourceAttr(name, "ttl", "1"),
					resource.TestCheckResourceAttr(name, "proxiable", "true"),
					resource.TestCheckResourceAttr(name, "zone_name", domain),
				),
			},
		},
	})
}

func TestAccCloudflareRecordDataSourceTXT(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_record.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRecordDataSourceConfigTXT(rnd, zoneID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "hostname", rnd+"."+domain),
					resource.TestCheckResourceAttr(name, "type", "TXT"),
					resource.TestCheckResourceAttr(name, "value", "i am a text record"),
					resource.TestCheckResourceAttr(name, "proxied", "false"),
					resource.TestCheckResourceAttr(name, "ttl", "1"),
					resource.TestCheckResourceAttr(name, "proxiable", "false"),
					resource.TestCheckResourceAttr(name, "zone_name", domain),
				),
			},
		},
	})
}

func TestAccCloudflareRecordDataSourceMX(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_record.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRecordDataSourceConfigMX(rnd, zoneID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "hostname", rnd+"."+domain),
					resource.TestCheckResourceAttr(name, "type", "MX"),
					resource.TestCheckResourceAttr(name, "value", "mx1.example.com"),
					resource.TestCheckResourceAttr(name, "proxied", "false"),
					resource.TestCheckResourceAttr(name, "ttl", "1"),
					resource.TestCheckResourceAttr(name, "proxiable", "false"),
					resource.TestCheckResourceAttr(name, "zone_name", domain),
					resource.TestCheckResourceAttr(name, "priority", "10"),
				),
			},
		},
	})
}

func testAccCloudflareRecordDataSourceConfig(rnd, zoneID, domain string) string {
	return acctest.LoadTestCase("recorddatasourceconfig.tf", rnd, zoneID, domain)
}

func testAccCloudflareRecordDataSourceConfigTXT(rnd, zoneID, domain string) string {
	return acctest.LoadTestCase("recorddatasourceconfigtxt.tf", rnd, zoneID, domain)
}

func testAccCloudflareRecordDataSourceConfigMX(rnd, zoneID, domain string) string {
	return acctest.LoadTestCase("recorddatasourceconfigmx.tf", rnd, zoneID, domain)
}
