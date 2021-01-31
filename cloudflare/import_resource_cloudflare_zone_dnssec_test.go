package cloudflare

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccCloudflareZoneDNSSEC_Import(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zone_dnssec.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneDNSSECResourceConfig(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareZoneDNSSECDataSourceID(name),
					resource.TestCheckResourceAttrSet(name, "zone_id"),
					resource.TestMatchResourceAttr(name, "status", regexp.MustCompile("active|disabled|pending")),
					resource.TestCheckResourceAttrSet(name, "flags"),
					resource.TestCheckResourceAttrSet(name, "algorithm"),
					resource.TestCheckResourceAttrSet(name, "key_type"),
					resource.TestCheckResourceAttrSet(name, "digest_type"),
					resource.TestCheckResourceAttrSet(name, "digest_algorithm"),
					resource.TestCheckResourceAttrSet(name, "digest"),
					resource.TestCheckResourceAttrSet(name, "ds"),
					resource.TestCheckResourceAttrSet(name, "key_tag"),
					resource.TestCheckResourceAttrSet(name, "public_key"),
					resource.TestCheckResourceAttrSet(name, "modified_on"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareZoneDNSSECDataSourceID(name),
					resource.TestCheckResourceAttrSet(name, "zone_id"),
					resource.TestMatchResourceAttr(name, "status", regexp.MustCompile("active|disabled|pending")),
					resource.TestCheckResourceAttrSet(name, "flags"),
					resource.TestCheckResourceAttrSet(name, "algorithm"),
					resource.TestCheckResourceAttrSet(name, "key_type"),
					resource.TestCheckResourceAttrSet(name, "digest_type"),
					resource.TestCheckResourceAttrSet(name, "digest_algorithm"),
					resource.TestCheckResourceAttrSet(name, "digest"),
					resource.TestCheckResourceAttrSet(name, "ds"),
					resource.TestCheckResourceAttrSet(name, "key_tag"),
					resource.TestCheckResourceAttrSet(name, "public_key"),
					resource.TestCheckResourceAttrSet(name, "modified_on"),
				),
			},
		},
	})
}
