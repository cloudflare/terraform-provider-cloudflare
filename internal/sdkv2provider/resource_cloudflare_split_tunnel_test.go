package sdkv2provider

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareSplitTunnel_Include(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_split_tunnel.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDefaultSplitTunnelInclude(rnd, accountID, "example domain", "*.example.com", "include"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "mode", "include"),
					resource.TestCheckResourceAttr(name, "tunnels.0.description", "example domain"),
					resource.TestCheckResourceAttr(name, "tunnels.0.host", "*.example.com"),
					resource.TestCheckNoResourceAttr(name, "policy_id"),
				),
			},
			{
				Config: testAccCloudflareDefaultSplitTunnelInclude(rnd, accountID, "example domain", "test.example.com", "include"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "mode", "include"),
					resource.TestCheckResourceAttr(name, "tunnels.0.description", "example domain"),
					resource.TestCheckResourceAttr(name, "tunnels.0.host", "test.example.com"),
					resource.TestCheckNoResourceAttr(name, "policy_id"),
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
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
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

func testAccCloudflareDefaultSplitTunnelInclude(rnd, accountID string, description string, host string, mode string) string {
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

func testAccCloudflareSplitTunnelInclude(rnd, accountID string, description string, host string, mode string) string {
	return fmt.Sprintf(`
resource "cloudflare_device_settings_policy" "%[1]s" {
	account_id                = "%[2]s"
	allow_mode_switch         = true
	allow_updates             = true
	allowed_to_leave          = true
	auto_connect              = 0
	captive_portal            = 5
	disable_auto_fallback     = true
	enabled                   = true
	match                     = "identity.email == \"foo@example.com\""
	name                      = "%[1]s"
	precedence                = 10
	support_url               = "https://cloudflare.com"
	switch_locked             = true
	exclude_office_ips        = false
}

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
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_split_tunnel.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareSplitTunnelExclude(rnd, accountID, "example domain", "*.example.com", "exclude"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "mode", "exclude"),
					resource.TestCheckResourceAttr(name, "tunnels.0.description", "example domain"),
					resource.TestCheckResourceAttr(name, "tunnels.0.host", "*.example.com"),
				),
			},
		},
	})
}

func TestAccCloudflareSplitTunnel_IncludeTunnelOrdering(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDefaultSplitTunnelIncludeMultiplesOrdered(rnd, accountID),
			},
			{
				Config: testAccCloudflareDefaultSplitTunnelIncludeMultiplesFlippedOrder(rnd, accountID),
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

func testAccCloudflareDefaultSplitTunnelIncludeMultiplesOrdered(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_split_tunnel" "%[1]s" {
  account_id = "%[2]s"
  mode = "include"
  tunnels {
    description = "example 1"
    host = "*.example.com"
  }

  tunnels {
    description = "example 2"
    host = "*.example.net"
  }
}
`, rnd, accountID)
}

func testAccCloudflareDefaultSplitTunnelIncludeMultiplesFlippedOrder(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_split_tunnel" "%[1]s" {
  account_id = "%[2]s"
  mode = "include"
  tunnels {
    description = "example 1"
    host = "*.example.com"
  }

  tunnels {
    description = "example 3"
    host = "*.example.org"
  }

  tunnels {
    description = "example 2"
    host = "*.example.net"
  }
}
`, rnd, accountID)
}
