package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCloudflareWAFRule_Import(t *testing.T) {
	t.Parallel()
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	ruleID := "100000"
	name := generateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareWAFRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWAFRuleConfig(zone, ruleID, "block", name),
			},
			{
				ResourceName:        "cloudflare_waf_rule." + name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zone),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}
