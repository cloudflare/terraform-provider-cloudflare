package cloudflare

import (
	"testing"

	"fmt"
	"os"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCloudflarePageRule_Import(t *testing.T) {
	t.Parallel()
	var pageRule cloudflare.PageRule
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	name := "cloudflare_page_rule.test"
	target := fmt.Sprintf("test-import.%s", zone)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigFullySpecified(zone, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(name, &pageRule),
				),
			},
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zone),
				ImportState:         true,
				ImportStateVerify:   true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(name, &pageRule),
				),
			},
		},
	})
}
