package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareRecordDataSource(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_record.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
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
					resource.TestCheckResourceAttr(name, "locked", "false"),
					resource.TestCheckResourceAttr(name, "zone_name", domain),
				),
			},
		},
	})
}

func TestAccCloudflareRecordDataSourceTXT(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_record.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
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
					resource.TestCheckResourceAttr(name, "locked", "false"),
					resource.TestCheckResourceAttr(name, "zone_name", domain),
				),
			},
		},
	})
}

func TestAccCloudflareRecordDataSourceMX(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_record.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
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
					resource.TestCheckResourceAttr(name, "locked", "false"),
					resource.TestCheckResourceAttr(name, "zone_name", domain),
					resource.TestCheckResourceAttr(name, "priority", "10"),
				),
			},
		},
	})
}

func testAccCloudflareRecordDataSourceConfig(rnd, zoneID, domain string) string {
	return fmt.Sprintf(`
data "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  hostname = cloudflare_record.%[1]s.hostname
}
resource "cloudflare_record" "%[1]s" {
	zone_id = "%[2]s"
	type = "A"
	name = "%[1]s.%[3]s"
	value = "192.0.2.0"
}`, rnd, zoneID, domain)
}

func testAccCloudflareRecordDataSourceConfigTXT(rnd, zoneID, domain string) string {
	return fmt.Sprintf(`
data "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  type = "TXT"
  hostname = cloudflare_record.%[1]s.hostname
}
resource "cloudflare_record" "%[1]s" {
	zone_id = "%[2]s"
	type = "TXT"
	name = "%[1]s.%[3]s"
	value = "i am a text record"
}`, rnd, zoneID, domain)
}

func testAccCloudflareRecordDataSourceConfigMX(rnd, zoneID, domain string) string {
	return fmt.Sprintf(`
data "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  type = "MX"
  priority = 10
  hostname = cloudflare_record.%[1]s.hostname
}
resource "cloudflare_record" "%[1]s" {
	zone_id = "%[2]s"
	type = "MX"
	name = "%[1]s.%[3]s"
	value = "mx1.example.com"
	priority = 10
}
resource "cloudflare_record" "%[1]s_2" {
	zone_id = "%[2]s"
	type = "MX"
	name = "%[1]s.%[3]s"
	value = "mx1.example.com"
	priority = 20
}
`, rnd, zoneID, domain)
}
