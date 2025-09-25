package zone_dnssec_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareZoneDNSSEC(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_zone_dnssec.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneDNSSECConfig(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.StringRegexp(regexp.MustCompile("active|disabled|pending"))),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("flags"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("algorithm"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("key_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("digest_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("digest_algorithm"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("digest"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("ds"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("key_tag"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("public_key"), knownvalue.NotNull()),
				},
				ExpectNonEmptyPlan: true, // Data source depends on DNSSEC resource with changing computed values
			},
		},
	})
}


func testAccCloudflareZoneDNSSECConfig(zoneID string, name string) string {
	return fmt.Sprintf(`
data "cloudflare_zone_dnssec" "%s" {
	zone_id = cloudflare_zone_dnssec.%s.zone_id
}

%s
`, name, name, testAccCloudflareZoneDNSSECResourceConfig(zoneID, name))
}

func testAccCloudflareZoneDNSSECResourceConfig(zoneID string, name string) string {
	return acctest.LoadTestCase("zonednssecresourceconfig.tf", name, zoneID)
}
