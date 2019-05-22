package cloudflare

import (
	"os"
	"testing"

	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCloudflareRateLimit_Import(t *testing.T) {
	t.Parallel()
	var rateLimit cloudflare.RateLimit
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := generateRandomResourceName()
	name := "cloudflare_rate_limit." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRateLimitConfigMatchingUrl(zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRateLimitExists(name, &rateLimit),
					testAccCheckCloudflareRateLimitIDIsValid(name, zone),
				),
			},
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zone),
				ImportState:         true,
				ImportStateVerify:   true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRateLimitExists(name, &rateLimit),
					testAccCheckCloudflareRateLimitIDIsValid(name, zone),
				),
			},
		},
	})
}
