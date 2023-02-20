package sdkv2provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareZone_PreventZoneIdAndNameConflicts(t *testing.T) {
	t.Parallel()
	rnd := generateRandomResourceName()
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCloudflareZoneConfigConflictingFields(rnd),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("only one of `name,zone_id` can be specified")),
			},
		},
	})
}

func testAccCloudflareZoneConfigConflictingFields(rnd string) string {
	return fmt.Sprintf(`
data "cloudflare_zone" "%[1]s" {
  name    = "terraform.cfapi.net"
  zone_id = "abc123"
}
`, rnd)
}

func TestAccCloudflareZone_NameLookup(t *testing.T) {
	t.Parallel()
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_zone.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneConfigBasic(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareZonesDataSourceID(name),
					resource.TestCheckResourceAttr(name, "name", "terraform.cfapi.net"),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, testAccCloudflareZoneID),
					resource.TestCheckResourceAttr(name, "status", "active"),
				),
			},
		},
	})
}

func testAccCloudflareZoneConfigBasic(rnd string) string {
	return fmt.Sprintf(`
data "cloudflare_zone" "%[1]s" {
  name = "terraform.cfapi.net"
}
`, rnd)
}
