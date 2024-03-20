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
				Config: testCloudflareTunnelMatchIsDeletedStep1(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudflare_tunnel."+rnd+"_default", "status", "inactive"),
					resource.TestCheckResourceAttr("data.cloudflare_tunnel."+rnd+"_default", "is_deleted", "false"),
					resource.TestCheckResourceAttr("data.cloudflare_tunnel."+rnd+"_not_deleted", "status", "inactive"),
					resource.TestCheckResourceAttr("data.cloudflare_tunnel."+rnd+"_not_deleted", "is_deleted", "false"),
				),
			},
			{
				Config:  testCloudflareTunnelMatchIsDeletedStep1(rnd),
				Destroy: true,
			},
			{
				Config: testCloudflareTunnelMatchIsDeletedStep2(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudflare_tunnel."+rnd+"_default", "status", "inactive"),
					resource.TestCheckResourceAttr("data.cloudflare_tunnel."+rnd+"_default", "is_deleted", "true"),
					resource.TestCheckResourceAttr("data.cloudflare_tunnel."+rnd+"_deleted", "status", "inactive"),
					resource.TestCheckResourceAttr("data.cloudflare_tunnel."+rnd+"_deleted", "is_deleted", "true"),
				),
			},
			{
				Config: testCloudflareTunnelMatchIsDeletedStep3(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudflare_tunnel."+rnd+"_not_deleted", "status", "inactive"),
				),
				ExpectError: regexp.MustCompile(`Error: No tunnels with name: ` + rnd),
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

data "cloudflare_tunnel" "%[2]s_default" {
	account_id = cloudflare_tunnel.%[2]s.account_id
	name       = cloudflare_tunnel.%[2]s.name
	depends_on = [cloudflare_tunnel.%[2]s]
}

data "cloudflare_tunnel" "%[2]s_not_deleted" {
	account_id = "%[1]s"
	name       = "%[2]s"
	is_deleted = false
	depends_on = [cloudflare_tunnel.%[2]s]
}
`, accountID, name)
}

func testCloudflareTunnelMatchIsDeletedStep2(name string) string {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	return fmt.Sprintf(`
data "cloudflare_tunnel" "%[2]s_default" {
	account_id = "%[1]s"
	name       = "%[2]s"
}

data "cloudflare_tunnel" "%[2]s_deleted" {
	account_id = "%[1]s"
	name       = "%[2]s"
	is_deleted = true
}
`, accountID, name)
}

func testCloudflareTunnelMatchIsDeletedStep3(name string) string {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	return fmt.Sprintf(`
data "cloudflare_tunnel" "%[2]s_not_deleted" {
	account_id = "%[1]s"
	name       = "%[2]s"
	is_deleted = false
}
`, accountID, name)
}
