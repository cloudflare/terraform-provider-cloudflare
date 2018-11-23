package cloudflare

import (
	"os"
	"testing"

	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCloudflareSpectrumApplication_Import(t *testing.T) {
	t.Parallel()
	var application cloudflare.SpectrumApplication
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
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
				ImportStateIdPrefix: fmt.Sprintf("%s/", zone),
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
