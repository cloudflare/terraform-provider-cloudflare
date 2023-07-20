package sdkv2provider

import (
	"testing"

	"fmt"
	"os"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflarePageRule_Import(t *testing.T) {
	t.Parallel()
	var pageRule cloudflare.PageRule
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	target := fmt.Sprintf("%s.%s", rnd, domain)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflarePageRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflarePageRuleConfigFullySpecified(zoneID, target, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
				),
			},
			{
				ResourceName:        resourceName,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
				ImportState:         true,
				ImportStateVerify:   true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
				),
			},
		},
	})
}

func TestAccCloudflarePageRule_ImportWithBrowserCacheTTL30(t *testing.T) {
	var pageRule cloudflare.PageRule
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	target := fmt.Sprintf("%s.%s", rnd, domain)

	testAccRunResourceTestSteps(t, []resource.TestStep{
		{
			Config: buildPageRuleConfig(rnd, zoneID, "browser_cache_ttl = 30", target),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
			),
		},
		{
			ResourceName:        resourceName,
			ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
			ImportState:         true,
			ImportStateVerify:   true,
			Check: resource.ComposeTestCheckFunc(
				testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
			),
		},
	})
}

func TestAccCloudflarePageRule_ImportWithoutBrowserCacheTTL(t *testing.T) {
	var pageRule cloudflare.PageRule
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := generateRandomResourceName()
	resourceName := "cloudflare_page_rule." + rnd
	target := fmt.Sprintf("%s.%s", rnd, domain)

	testAccRunResourceTestSteps(t, []resource.TestStep{
		{
			Config: buildPageRuleConfig(rnd, zoneID, `browser_check = "on"`, target),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
			),
		},
		{
			ResourceName:        resourceName,
			ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
			ImportState:         true,
			ImportStateVerify:   true,
			Check: resource.ComposeTestCheckFunc(
				testAccCheckCloudflarePageRuleExists(resourceName, &pageRule),
			),
		},
	})
}
