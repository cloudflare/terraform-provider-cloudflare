package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccCloudflareRecordDataSource(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_record.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	recordID := os.Getenv("CLOUDFLARE_RECORD_ID")
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRecordDataSourceConfig(rnd, zoneID, recordID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "hostname", "www.example.com"),
				),
			},
		},
	})
}

func testAccCloudflareRecordDataSourceConfig(rnd, zoneID, recordID string) string {
	return fmt.Sprintf(`
data "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  record_id = "%[3]s"
}`, rnd, zoneID, recordID)
}
