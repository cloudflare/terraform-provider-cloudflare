package cloudflare

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccCloudflareAccountMemberBasic(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_account_mamber." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckAccount(t)
			testAccPreCheckEmail(t)
			testAccPreCheckApiKey(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareAccountMemberBasicConfig(rnd, fmt.Sprintf("%s@domain.com", rnd)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "email_address", fmt.Sprintf("%s@domain.com", rnd)),
					resource.TestCheckResourceAttr(name, "role_ids.#", "1"),
					resource.TestCheckResourceAttr(name, "role_ids.0", "05784afa30c1afe1440e79d9351c7430"),
				),
			},
		},
	})
}

func testCloudflareAccountMemberBasicConfig(resourceID, emailAddress string) string {
	return fmt.Sprintf(`
				resource "cloudflare_account_member" "%[1]s" {
					email_address = "%[2]s"
					role_ids = [ "05784afa30c1afe1440e79d9351c7430" ]
				}`, resourceID, emailAddress)
}
