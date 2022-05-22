package provider

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareSplitTunnel_Include(t *testing.T) {
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
	name := fmt.Sprintf("cloudflare_split_tunnel.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccessAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareSplitTunnelInclude(rnd, accountID, "example domain", "*.example.com", "include"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "mode", "include"),
					resource.TestCheckResourceAttr(name, "tunnels.0.description", "example domain"),
					resource.TestCheckResourceAttr(name, "tunnels.0.host", "*.example.com"),
				),
			},
			{
				Config: testAccCloudflareSplitTunnelInclude(rnd, accountID, "example domain", "test.example.com", "include"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "mode", "include"),
					resource.TestCheckResourceAttr(name, "tunnels.0.description", "example domain"),
					resource.TestCheckResourceAttr(name, "tunnels.0.host", "test.example.com"),
				),
			},
		},
	})
}

func TestAccCloudflareSplitTunnel_ConflictingTunnelProperties(t *testing.T) {
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

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccessAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCloudflareSplitTunnelConflictingTunnelProperties(rnd, accountID, "example domain", "include"),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("address and host are mutually exclusive and cannot be applied together in the same block")),
			},
		},
	})
}

func testAccCloudflareSplitTunnelInclude(rnd, accountID string, description string, host string, mode string) string {
	return fmt.Sprintf(`
resource "cloudflare_split_tunnel" "%[1]s" {
  account_id = "%[2]s"
  mode = "%[5]s"
  tunnels {
    description = "%[3]s"
    host = "%[4]s"
  }
}
`, rnd, accountID, description, host, mode)
}

func testAccCloudflareSplitTunnelConflictingTunnelProperties(rnd, accountID string, description string, mode string) string {
	return fmt.Sprintf(`
resource "cloudflare_split_tunnel" "%[1]s" {
  account_id = "%[2]s"
  mode = "%[4]s"
  tunnels {
    description = "%[3]s"
    address = "192.0.2.0/24"
    host = "example.com"
  }
}
`, rnd, accountID, description, mode)
}

func TestAccCloudflareSplitTunnel_Exclude(t *testing.T) {
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
	name := fmt.Sprintf("cloudflare_split_tunnel.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccessAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareSplitTunnelExclude(rnd, accountID, "example domain", "*.example.com", "exclude"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "mode", "exclude"),
					resource.TestCheckResourceAttr(name, "tunnels.0.description", "example domain"),
					resource.TestCheckResourceAttr(name, "tunnels.0.host", "*.example.com"),
				),
			},
		},
	})
}

func testAccCloudflareSplitTunnelExclude(rnd, accountID string, description string, host string, mode string) string {
	return fmt.Sprintf(`
resource "cloudflare_split_tunnel" "%[1]s" {
  account_id = "%[2]s"
  mode = "%[5]s"
  tunnels {
    description= "%[3]s"
    host = "%[4]s"
  }
}
`, rnd, accountID, description, host, mode)
}
