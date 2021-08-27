package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccCloudflareRecord_Import(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	name := "cloudflare_record.foobar"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, name),
			},
			{
				ResourceName:            name,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", zoneID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"allow_overwrite"},
			},
		},
	})
}
