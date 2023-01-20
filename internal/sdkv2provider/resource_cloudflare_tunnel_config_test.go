package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testTunnelConfig(resourceID, accountID, tunnelSecret string) string {
	return fmt.Sprintf(`
		resource "cloudflare_tunnel" "%[1]s" {
		  account_id = "%[2]s"
		  name       = "%[1]s"
		  secret     = "%[3]s"
		}

		resource "cloudflare_tunnel_config" "%[1]s" {
		  account_id         = "%[2]s"
		  tunnel_id          = cloudflare_tunnel.%[1]s.id

		  config {
			warp_routing {
			  enabled = true
			}
			origin_request {
			  connect_timeout = "1m0s"
			  tls_timeout = "1m0s"
			  tcp_keep_alive = "1m0s"
			  no_happy_eyeballs = false
			  keep_alive_connections = 1024
			  keep_alive_timeout = "1m0s"
			  http_host_header = "baz"
			  origin_server_name = "foobar"
			  ca_pool = "/path/to/unsigned/ca/pool"
			  no_tls_verify = false
			  disable_chunked_encoding = false
			  bastion_mode = false
			  proxy_address = "10.0.0.1"
			  proxy_port = "8123"
			  proxy_type = "socks"
			  ip_rules {
				prefix = "/web"
				ports = [80, 443]
				allow = false
			  }
			}
			ingress_rule {
			  hostname = "foo"
			  path = "/bar"
			  service = "http://10.0.0.2:8080"
			}
			ingress_rule {
				service = "https://10.0.0.3:8081"
			  }
		  }
		}
		`, resourceID, accountID, tunnelSecret)
}

func testTunnelConfigShort(resourceID, accountID, tunnelSecret string) string {
	return fmt.Sprintf(`
		resource "cloudflare_tunnel" "%[1]s" {
		  account_id = "%[2]s"
		  name       = "%[1]s"
		  secret     = "%[3]s"
		}

		resource "cloudflare_tunnel_config" "%[1]s" {
		  account_id         = "%[2]s"
		  tunnel_id          = cloudflare_tunnel.%[1]s.id

		  config {
			ingress_rule {
				service = "https://10.0.0.1:8081"
			  }
		  }
		}
		`, resourceID, accountID, tunnelSecret)
}

func TestAccCloudflareTunnelConfig_Full(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_tunnel_config." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tunnelSecret := acctest.RandStringFromCharSet(32, acctest.CharSetAlpha)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testTunnelConfig(rnd, zoneID, tunnelSecret),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "config.0.warp_routing.0.enabled", "true"),
					resource.TestCheckResourceAttr(name, "config.0.origin_request.0.connect_timeout", "1m0s"),
					resource.TestCheckResourceAttr(name, "config.0.origin_request.0.tls_timeout", "1m0s"),
					resource.TestCheckResourceAttr(name, "config.0.origin_request.0.tcp_keep_alive", "1m0s"),
					resource.TestCheckResourceAttr(name, "config.0.origin_request.0.no_happy_eyeballs", "false"),
					resource.TestCheckResourceAttr(name, "config.0.origin_request.0.keep_alive_connections", "1024"),
					resource.TestCheckResourceAttr(name, "config.0.origin_request.0.keep_alive_timeout", "1m0s"),
					resource.TestCheckResourceAttr(name, "config.0.origin_request.0.http_host_header", "baz"),
					resource.TestCheckResourceAttr(name, "config.0.origin_request.0.origin_server_name", "foobar"),
					resource.TestCheckResourceAttr(name, "config.0.origin_request.0.ca_pool", "/path/to/unsigned/ca/pool"),
					resource.TestCheckResourceAttr(name, "config.0.origin_request.0.no_tls_verify", "false"),
					resource.TestCheckResourceAttr(name, "config.0.origin_request.0.disable_chunked_encoding", "false"),
					resource.TestCheckResourceAttr(name, "config.0.origin_request.0.bastion_mode", "false"),
					resource.TestCheckResourceAttr(name, "config.0.origin_request.0.proxy_address", "10.0.0.1"),
					resource.TestCheckResourceAttr(name, "config.0.origin_request.0.proxy_port", "8123"),
					resource.TestCheckResourceAttr(name, "config.0.origin_request.0.proxy_type", "socks"),

					resource.TestCheckResourceAttr(name, "config.0.origin_request.0.ip_rules.0.prefix", "/web"),
					resource.TestCheckResourceAttr(name, "config.0.origin_request.0.ip_rules.0.ports.#", "2"),
					resource.TestCheckResourceAttr(name, "config.0.origin_request.0.ip_rules.0.ports.0", "80"),
					resource.TestCheckResourceAttr(name, "config.0.origin_request.0.ip_rules.0.ports.1", "443"),
					resource.TestCheckResourceAttr(name, "config.0.origin_request.0.ip_rules.0.allow", "false"),

					resource.TestCheckResourceAttr(name, "config.0.ingress_rule.#", "2"),
					resource.TestCheckResourceAttr(name, "config.0.ingress_rule.0.hostname", "foo"),
					resource.TestCheckResourceAttr(name, "config.0.ingress_rule.0.path", "/bar"),
					resource.TestCheckResourceAttr(name, "config.0.ingress_rule.0.service", "http://10.0.0.2:8080"),
					resource.TestCheckResourceAttr(name, "config.0.ingress_rule.1.hostname", ""),
					resource.TestCheckResourceAttr(name, "config.0.ingress_rule.1.path", ""),
					resource.TestCheckResourceAttr(name, "config.0.ingress_rule.1.service", "https://10.0.0.3:8081"),
				),
			},
		},
	})
}

func TestAccCloudflareTunnelConfig_Short(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_tunnel_config." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tunnelSecret := acctest.RandStringFromCharSet(32, acctest.CharSetAlpha)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testTunnelConfigShort(rnd, zoneID, tunnelSecret),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "config.0.ingress_rule.#", "1"),
					resource.TestCheckResourceAttr(name, "config.0.ingress_rule.0.service", "https://10.0.0.1:8081"),
				),
			},
		},
	})
}
