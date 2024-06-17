package sdkv2provider

import (
	"fmt"
	"os"
	"regexp"
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
				Config: testCloudflareTunnelMatchIsDeleted_Default(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudflare_tunnel."+rnd, "status", "inactive"),
					resource.TestCheckResourceAttr("data.cloudflare_tunnel."+rnd, "is_deleted", "false"),
				),
			},
			{
				Config:  testCloudflareTunnelMatchIsDeleted_Default(rnd),
				Destroy: true,
			},
			{
				Config: testCloudflareTunnelMatchIsDeleted_DeletedTunnels(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudflare_tunnel."+rnd, "status", "inactive"),
					resource.TestCheckResourceAttr("data.cloudflare_tunnel."+rnd, "is_deleted", "true"),
				),
			},
			{
				Config:  testCloudflareTunnelMatchIsDeleted_DeletedTunnels(rnd),
				Destroy: true,
			},
			{
				Config: testCloudflareTunnelMatchIsDeleted_ActiveTunnels(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudflare_tunnel."+rnd, "status", "inactive"),
					resource.TestCheckResourceAttr("data.cloudflare_tunnel."+rnd, "is_deleted", "true"),
				),
				ExpectError: regexp.MustCompile(`Error: No tunnels with name: ` + rnd),
			},
		},
	})
}

func testCloudflareTunnelMatchIsDeleted_Default(name string) string {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	return fmt.Sprintf(`
resource "cloudflare_tunnel" "%[2]s" {
	account_id = "%[1]s"
	name       = "%[2]s"
	secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}

data "cloudflare_tunnel" "%[2]s" {
	account_id = "%[1]s"
	name       = "%[2]s"
	is_deleted = false
	depends_on = [cloudflare_tunnel.%[2]s]
}
`, accountID, name)
}

func testCloudflareTunnelMatchIsDeleted_DeletedTunnels(name string) string {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	return fmt.Sprintf(`
data "cloudflare_tunnel" "%[2]s" {
	account_id = "%[1]s"
	name       = "%[2]s"
	is_deleted = true
}
`, accountID, name)
}

func testCloudflareTunnelMatchIsDeleted_ActiveTunnels(name string) string {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	return fmt.Sprintf(`
data "cloudflare_tunnel" "%[2]s" {
	account_id = "%[1]s"
	name       = "%[2]s"
	is_deleted = false
}
`, accountID, name)
}
