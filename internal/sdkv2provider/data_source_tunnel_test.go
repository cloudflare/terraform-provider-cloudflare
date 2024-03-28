package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareTunnel_MatchName(t *testing.T) {
	rnd := generateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareTunnelMatchName(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudflare_tunnel."+rnd, "status", "inactive"),
				),
			},
		},
	})
}

func testCloudflareTunnelMatchName(name string) string {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	return fmt.Sprintf(`
resource "cloudflare_tunnel" "%[2]s" {
	account_id = "%[1]s"
	name       = "%[2]s"
	secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}

data "cloudflare_tunnel" "%[2]s" {
	account_id = cloudflare_tunnel.%[2]s.account_id
	name       = cloudflare_tunnel.%[2]s.name
	depends_on = [cloudflare_tunnel.%[2]s]
}
`, accountID, name)
}

func TestAccCloudflareTunnel_MatchIsDeleted(t *testing.T) {
	rnd := generateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareTunnelMatchIsDeletedStep1(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudflare_tunnel."+rnd, "status", "inactive"),
				),
			},
			{
				Config: "// delete tunnel resource",
			},
			{
				Config: testCloudflareTunnelMatchIsDeletedStep3(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudflare_tunnel."+rnd+"_deleted", "status", "inactive"),
					resource.TestCheckResourceAttr("data.cloudflare_tunnel."+rnd+"_default", "status", "inactive"),
				),
			},
		},
	})
}

func testCloudflareTunnelMatchIsDeletedStep1(name string) string {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	return fmt.Sprintf(`
resource "cloudflare_tunnel" "%[2]s" {
	account_id = "%[1]s"
	name       = "%[2]s"
	secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}

data "cloudflare_tunnel" "%[2]s" {
	account_id = cloudflare_tunnel.%[2]s.account_id
	name       = cloudflare_tunnel.%[2]s.name
    is_deleted = false
	depends_on = [cloudflare_tunnel.%[2]s]
}
`, accountID, name)
}

func testCloudflareTunnelMatchIsDeletedStep3(name string) string {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	return fmt.Sprintf(`
data "cloudflare_tunnel" "%[2]s_deleted" {
	account_id = "%[1]s"
	name       = "%[2]s"
    is_deleted = true
}

data "cloudflare_tunnel" "%[2]s_default" {
	account_id = "%[1]s"
	name       = "%[2]s"
}
`, accountID, name)
}
