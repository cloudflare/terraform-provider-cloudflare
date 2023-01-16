package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareRecord_ImportBasic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_record.%s", rnd)
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigBasic(zoneID, rnd, rnd),
			},
			{
				ResourceName:            resourceName,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", zoneID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"allow_overwrite"},
			},
		},
	})
}

func TestAccCloudflareRecord_ImportSRV(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_record.%s", rnd)
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRecordConfigSRV(zoneID, rnd, domain),
			},
			{
				ResourceName:            resourceName,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", zoneID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"allow_overwrite"},
			},
		},
	})
}
