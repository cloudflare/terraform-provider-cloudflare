package cloudflare

import (
	"os"
	"testing"

	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCloudFlareRateLimit_Import(t *testing.T) {
	t.Parallel()
	var rateLimit cloudflare.RateLimit
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := acctest.RandString(10)
	name := "cloudflare_rate_limit." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareRateLimitConfigMatchingUrl(zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRateLimitExists(name, &rateLimit),
					testAccCheckCloudFlareRateLimitIDIsValid(name, zone),
				),
			},
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zone),
				ImportState:         true,
				ImportStateVerify:   true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareRateLimitExists(name, &rateLimit),
					testAccCheckCloudFlareRateLimitIDIsValid(name, zone),
				),
			},
		},
	})
}
