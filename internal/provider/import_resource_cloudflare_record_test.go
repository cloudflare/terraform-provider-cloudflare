package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareRecord_Import(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_record.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, name, rnd),
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
