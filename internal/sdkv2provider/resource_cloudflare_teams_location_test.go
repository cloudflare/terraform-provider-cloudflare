package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareTeamsLocationBasic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_dns_location.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareTeamsLocationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsLocationConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "client_default", "false"),
					resource.TestCheckResourceAttr(name, "ecs_support", "false"),
					resource.TestCheckResourceAttr(name, "networks.#", "1"),
					resource.TestCheckResourceAttr(name, "networks.0.network", "2.5.6.200/32"),
					resource.TestCheckResourceAttr(name, "endpoints.#", "1"),
					resource.TestCheckResourceAttr(name, "endpoints.0.ipv4.0.enabled", "true"),
					resource.TestCheckResourceAttr(name, "endpoints.0.ipv4.0.authentication_enabled", "true"),

					resource.TestCheckResourceAttr(name, "endpoints.0.ipv6.0.enabled", "true"),
					resource.TestCheckResourceAttr(name, "endpoints.0.ipv6.0.authentication_enabled", "true"),
					resource.TestCheckResourceAttr(name, "endpoints.0.ipv6.0.networks.0.network", "2a09:bac5:50c3:400::6b:57/128"),

					resource.TestCheckResourceAttr(name, "endpoints.0.doh.0.enabled", "true"),
					resource.TestCheckResourceAttr(name, "endpoints.0.doh.0.authentication_enabled", "true"),
					resource.TestCheckResourceAttr(name, "endpoints.0.doh.0.networks.0.network", "2.5.6.202/32"),
					resource.TestCheckResourceAttr(name, "endpoints.0.doh.0.networks.1.network", "3.5.6.203/32"),

					resource.TestCheckResourceAttr(name, "endpoints.0.dot.0.enabled", "true"),
					resource.TestCheckResourceAttr(name, "endpoints.0.dot.0.authentication_enabled", "true"),
					resource.TestCheckResourceAttr(name, "endpoints.0.dot.0.networks.0.network", "2.5.6.201/32"),
				),
			},
		},
	})
}

func testAccCloudflareTeamsLocationConfigBasic(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_dns_location" "%[1]s" {
  name        = "%[1]s"
  account_id  = "%[2]s"
  networks = [{ network = "2.5.6.200/32" }]

  dns_destination_ips_id = "0e4a32c6-6fb8-4858-9296-98f51631e8e6"
  
  endpoints   {
		ipv4   { 
			enabled = true
		}
		ipv6   { 
			enabled = true 
			networks  = [ { network = "2a09:bac5:50c3:400::6b:57/128" } ]
		}
		dot   {
			enabled = true
			networks  = [ { network = "2.5.6.201/32" } ]
		}
		doh    {
			enabled = true 
			networks  = [ { network = "2.5.6.202/32" }, { network = "3.5.6.203/32" } ]
		}
	}
}
`, rnd, accountID)
}

func testAccCheckCloudflareTeamsLocationDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_dns_location" {
			continue
		}

		_, err := client.TeamsLocation(context.Background(), rs.Primary.Attributes[consts.AccountIDSchemaKey], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("teams Location still exists")
		}
	}

	return nil
}
