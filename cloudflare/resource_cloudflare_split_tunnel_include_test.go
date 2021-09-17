package cloudflare

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccCloudflareSplitTunnelInclude(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		defer func(apiToken string) {
			os.Setenv("CLOUDFLARE_API_TOKEN", apiToken)
		}(os.Getenv("CLOUDFLARE_API_TOKEN"))
		os.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_split_tunnel_include.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccessAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareSplitTunnelIncludeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareSplitTunnelInclude(rnd, accountID, "example domain", "*.example.com"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "description", "example domain",
					resource.TestCheckResourceAttr(name, "host", "*.example.com",

				),
			},
		},
	})
}

func testAccCloudflareSplitTunnelInclude(rnd, accountID string, description string, host string) string {
	return fmt.Sprintf(`
resource "cloudflare_split_tunnel_include" "%[1]s" {
	account_id                = "%[2]s"
	description               = "%[3]s"
	host                      = "%[4]s"
}
`, rnd, accountID, description, host)
}

func testAccCheckCloudflareSplitTunnelIncludeDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_split_tunnel_include" {
			continue
		}

		_, err := client.SplitTunnelInclude(context.Background(), rs.Primary.Attributes["account_id"])
		if err == nil {
			return fmt.Errorf("Split Tunnel Include still exists")
		}
	}

	return nil
}