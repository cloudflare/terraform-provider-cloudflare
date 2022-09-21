package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareAccount_Basic(t *testing.T) {
	t.Parallel()

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_account.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccountName(fmt.Sprintf("%s_old", rnd)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						name, "name", fmt.Sprintf("%s_old", rnd)),
				),
			},
			{
				Config: testAccCheckCloudflareAccountName(fmt.Sprintf("%s_new", rnd)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						name, "name", fmt.Sprintf("%s_new", rnd)),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccountName(name string) string {
	return fmt.Sprintf(`
  resource "cloudflare_account" "%[1]s" {
	  name = "%[1]s"
  }`, name)
}
