package dcv_delegation_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareDCVDelegationDataSource(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareDCVDelegationDataSourceConfig(zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.cloudflare_dcv_delegation.test", "id"),
					resource.TestCheckResourceAttrSet("data.cloudflare_dcv_delegation.test", "hostname"),
					resource.TestCheckResourceAttrSet("data.cloudflare_dcv_delegation.test", consts.ZoneIDSchemaKey),
				),
			},
		},
	})
}

func testAccCheckCloudflareDCVDelegationDataSourceConfig(zoneID string) string {
	return fmt.Sprintf(`
data "cloudflare_dcv_delegation" "test" {
    zone_id = "%s"
}
`, zoneID)
}
