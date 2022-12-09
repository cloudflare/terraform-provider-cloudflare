package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareWAFRule_Import(t *testing.T) {
	skipV1WAFTestForNonConfiguredDefaultZone(t)

	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	ruleID := "100001"
	name := generateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareWAFRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWAFRuleConfig(zoneID, ruleID, "block", name),
			},
			{
				ResourceName:        "cloudflare_waf_rule." + name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}
