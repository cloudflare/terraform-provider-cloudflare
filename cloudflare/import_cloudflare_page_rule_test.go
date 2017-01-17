package cloudflare

import (
	"testing"

	"fmt"
	"os"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCloudFlarePageRule_Import(t *testing.T) {
	t.Parallel()
	var pageRule cloudflare.PageRule
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	name := "cloudflare_page_rule.test"
	target := fmt.Sprintf("test-import.%s", zone)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlarePageRuleConfigFullySpecified(zone, target),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlarePageRuleExists(name, &pageRule),
				),
			},
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zone),
				ImportState:         true,
				ImportStateVerify:   true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlarePageRuleExists(name, &pageRule),
				),
			},
		},
	})
}
