package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCloudflareSpectrumApplication_Import(t *testing.T) {
	t.Parallel()
	var application cloudflare.SpectrumApplication
	zone := os.Getenv("CLOUDFLARE_DOMAIN")

	// This declaration can be removed once CLOUDFLARE_ZONE_ID is set in
	// TeamCity for the acceptance tests. In the meantime, this environment
	// variable needs to be set in order to match the CLOUDFLARE_DOMAIN the
	// tests will run against. This is the zone ID for the
	// hashicorptest.com zone.
	zoneID := "25afd8e9b39af234c001b657a2eb2c5c"
	if id := os.Getenv("CLOUDFLARE_ZONE_ID"); id != "" {
		zoneID = id
	}

	rnd := acctest.RandString(10)
	name := "cloudflare_spectrum_application." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareSpectrumApplicationConfigBasic(zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareSpectrumApplicationExists(name, &application),
					testAccCheckCloudflareSpectrumApplicationIDIsValid(name),
				),
			},
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
				ImportState:         true,
				ImportStateVerify:   true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareSpectrumApplicationExists(name, &application),
					testAccCheckCloudflareSpectrumApplicationIDIsValid(name),
				),
			},
		},
	})
}
