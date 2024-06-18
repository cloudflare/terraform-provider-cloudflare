package total_tls_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stainless-sdks/cloudflare-terraform/internal/acctest"
	"github.com/stainless-sdks/cloudflare-terraform/internal/consts"
	"github.com/stainless-sdks/cloudflare-terraform/internal/utils"
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
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_total_tls." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
