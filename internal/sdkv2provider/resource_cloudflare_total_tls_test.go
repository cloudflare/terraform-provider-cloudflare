package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testTotalTLS(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_total_tls" "%[1]s" {
	zone_id = "%[2]s"
	enabled = true
	certificate_authority = "google"
}
`, rnd, zoneID)
}

func TestAccCloudflareTotalTLS(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_total_tls." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testTotalTLS(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, "certificate_authority", "google"),
				),
			},
			{
				ImportState:       true,
				ImportStateVerify: true,
				ResourceName:      name,
			},
		},
	})
}
