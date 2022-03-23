package cloudflare

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudflareTeamsProxyEndpoint_Basic(t *testing.T) {
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
	name := fmt.Sprintf("cloudflare_teams_proxy_endpoint.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccessAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareTeamsProxyEndpointDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsProxyEndpointConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "ips.0", "104.16.132.229/32"),
					resource.TestMatchResourceAttr(name, "subdomain", regexp.MustCompile("^[a-zA-Z0-9]+$")),
				),
			},
		},
	})
}

func testAccCloudflareTeamsProxyEndpointConfigBasic(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_teams_proxy_endpoint" "%[1]s" {
  name        = "%[1]s"
  account_id  = "%[2]s"
  ips  = ["104.16.132.229/32"]
}
`, rnd, accountID)
}

func testAccCheckCloudflareTeamsProxyEndpointDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_teams_proxy_endpoint" {
			continue
		}

		_, err := client.TeamsProxyEndpoint(context.Background(), rs.Primary.Attributes["account_id"], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("teams Proxy Endpoint still exists")
		}
	}

	return nil
}
