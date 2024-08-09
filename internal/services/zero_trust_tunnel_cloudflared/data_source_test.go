package zero_trust_tunnel_cloudflared_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareTunnel_MatchName(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareTunnelMatchName(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudflare_zero_trust_tunnel_cloudflared."+rnd, "status", "inactive"),
				),
			},
		},
	})
}

func testCloudflareTunnelMatchName(name string) string {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	return acctest.LoadTestCase("cloudflaretunnelmatchname.tf", accountID, name)
}

func TestAccCloudflareTunnel_MatchIsDeleted(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareTunnelMatchIsDeleted_Default(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudflare_zero_trust_tunnel_cloudflared."+rnd, "status", "inactive"),
					resource.TestCheckResourceAttr("data.cloudflare_zero_trust_tunnel_cloudflared."+rnd, "is_deleted", "false"),
				),
			},
			{
				Config:  testCloudflareTunnelMatchIsDeleted_Default(rnd),
				Destroy: true,
			},
			{
				Config: testCloudflareTunnelMatchIsDeleted_DeletedTunnels(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudflare_zero_trust_tunnel_cloudflared."+rnd, "status", "inactive"),
					resource.TestCheckResourceAttr("data.cloudflare_zero_trust_tunnel_cloudflared."+rnd, "is_deleted", "true"),
				),
			},
			{
				Config:  testCloudflareTunnelMatchIsDeleted_DeletedTunnels(rnd),
				Destroy: true,
			},
			{
				Config: testCloudflareTunnelMatchIsDeleted_ActiveTunnels(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudflare_zero_trust_tunnel_cloudflared."+rnd, "status", "inactive"),
					resource.TestCheckResourceAttr("data.cloudflare_zero_trust_tunnel_cloudflared."+rnd, "is_deleted", "true"),
				),
				ExpectError: regexp.MustCompile(`Error: No tunnels with name: ` + rnd),
			},
		},
	})
}

func testCloudflareTunnelMatchIsDeleted_Default(name string) string {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	return acctest.LoadTestCase("default.tf", accountID, name)
}

func testCloudflareTunnelMatchIsDeleted_DeletedTunnels(name string) string {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	return fmt.Sprintf(`
data "cloudflare_zero_trust_tunnel_cloudflared" "%[2]s" {
	account_id = "%[1]s"
	name       = "%[2]s"
	is_deleted = true
}
`, accountID, name)
}

func testCloudflareTunnelMatchIsDeleted_ActiveTunnels(name string) string {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	return fmt.Sprintf(`
data "cloudflare_zero_trust_tunnel_cloudflared" "%[2]s" {
	account_id = "%[1]s"
	name       = "%[2]s"
	is_deleted = false
}
`, accountID, name)
}
