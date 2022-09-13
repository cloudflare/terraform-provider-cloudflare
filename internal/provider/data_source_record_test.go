package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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

func testAccCloudflareRecordDataSourceConfig(rnd, zoneID, domain string) string {
	return fmt.Sprintf(`
data "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  record_id = cloudflare_record.%[1]s.id
}
resource "cloudflare_record" "%[1]s" {
	zone_id = "%[2]s"
	type = "A"
	name = "%[1]s.%[3]s"
	value = "192.0.2.0"
}`, rnd, zoneID, domain)
}
