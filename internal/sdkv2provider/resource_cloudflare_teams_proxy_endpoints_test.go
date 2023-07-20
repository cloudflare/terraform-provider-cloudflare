package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareTeamsProxyEndpoint_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_teams_proxy_endpoint.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareTeamsProxyEndpointDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsProxyEndpointConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
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

		_, err := client.TeamsProxyEndpoint(context.Background(), rs.Primary.Attributes[consts.AccountIDSchemaKey], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("teams Proxy Endpoint still exists")
		}
	}

	return nil
}
