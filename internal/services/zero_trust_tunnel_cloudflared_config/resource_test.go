package zero_trust_tunnel_cloudflared_config_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testTunnelConfig(resourceID, accountID, tunnelSecret, domain string) string {
	return acctest.LoadTestCase("tunnelconfig.tf", resourceID, accountID, tunnelSecret, domain)
}

func testTunnelConfigShort(resourceID, accountID, tunnelSecret string) string {
	return acctest.LoadTestCase("tunnelconfigshort.tf", resourceID, accountID, tunnelSecret)
}

func testTunnelConfigNilPointer(resourceID, accountID, tunnelSecret string) string {
	return acctest.LoadTestCase("tunnelconfignilpointer.tf", resourceID, accountID, tunnelSecret)
}

func TestAccCloudflareTunnelConfig_Full(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_tunnel_cloudflared_config." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	tunnelSecret := utils.RandStringFromCharSet(32, utils.CharSetAlpha)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testTunnelConfig(rnd, zoneID, tunnelSecret, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "config.warp_routing.enabled", "true"),
					resource.TestCheckResourceAttr(name, "config.origin_request.connect_timeout", "60"),
					resource.TestCheckResourceAttr(name, "config.origin_request.tls_timeout", "60"),
					resource.TestCheckResourceAttr(name, "config.origin_request.tcp_keep_alive", "60"),
					resource.TestCheckResourceAttr(name, "config.origin_request.no_happy_eyeballs", "false"),
					resource.TestCheckResourceAttr(name, "config.origin_request.keep_alive_connections", "1024"),
					resource.TestCheckResourceAttr(name, "config.origin_request.keep_alive_timeout", "60"),
					resource.TestCheckResourceAttr(name, "config.origin_request.http_host_header", "baz"),
					resource.TestCheckResourceAttr(name, "config.origin_request.origin_server_name", "foobar"),
					resource.TestCheckResourceAttr(name, "config.origin_request.ca_pool", "/path/to/unsigned/ca/pool"),
					resource.TestCheckResourceAttr(name, "config.origin_request.no_tls_verify", "false"),
					resource.TestCheckResourceAttr(name, "config.origin_request.disable_chunked_encoding", "false"),
					resource.TestCheckResourceAttr(name, "config.origin_request.proxy_type", "socks"),
					resource.TestCheckResourceAttr(name, "config.origin_request.http2_origin", "true"),

					resource.TestCheckResourceAttr(name, "config.ingress.0.hostname", "foo"),
					resource.TestCheckResourceAttr(name, "config.ingress.0.path", "/bar"),
					resource.TestCheckResourceAttr(name, "config.ingress.0.service", "http://10.0.0.2:8080"),
					resource.TestCheckResourceAttr(name, "config.ingress.0.origin_request.connect_timeout", "15"),

					resource.TestCheckResourceAttr(name, "config.ingress.1.hostname", "bar"),
					resource.TestCheckResourceAttr(name, "config.ingress.1.path", "/foo"),
					resource.TestCheckResourceAttr(name, "config.ingress.1.service", "http://10.0.0.3:8081"),
					resource.TestCheckResourceAttr(name, "config.ingress.1.origin_request.access.required", "true"),
					resource.TestCheckResourceAttr(name, "config.ingress.1.origin_request.access.team_name", "terraform"),
					resource.TestCheckResourceAttr(name, "config.ingress.1.origin_request.access.aud_tag.#", "1"),
					resource.TestCheckResourceAttr(name, "config.ingress.1.origin_request.access.aud_tag.0", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),

					resource.TestCheckResourceAttr(name, "config.ingress.2.service", "https://10.0.0.4:8082"),
				),
			},
			// {
			// 	ResourceName:        name,
			// 	ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
			// 	ImportState:         true,
			// 	ImportStateVerify:   true,
			// },
		},
	})
}

func TestAccCloudflareTunnelConfig_Short(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_tunnel_cloudflared_config." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tunnelSecret := utils.RandStringFromCharSet(32, utils.CharSetAlpha)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testTunnelConfigShort(rnd, zoneID, tunnelSecret),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "config.ingress.0.service", "https://10.0.0.1:8081"),
				),
			},
			// {
			// 	ResourceName:        name,
			// 	ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
			// 	ImportState:         true,
			// 	ImportStateVerify:   true,
			// },
		},
	})
}

func TestAccCloudflareTunnelConfig_NilPointer(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_tunnel_cloudflared_config." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tunnelSecret := utils.RandStringFromCharSet(32, utils.CharSetAlpha)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testTunnelConfigNilPointer(rnd, zoneID, tunnelSecret),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "config.warp_routing.enabled", "true"),
					resource.TestCheckResourceAttr(name, "config.origin_request.no_tls_verify", "true"),
					resource.TestCheckResourceAttr(name, "config.ingress.0.hostname", "foo"),
					resource.TestCheckResourceAttr(name, "config.ingress.0.service", "https://10.0.0.1:8006"),
					resource.TestCheckResourceAttr(name, "config.ingress.1.service", "http_status:501"),
				),
			},
			// {
			// 	ResourceName:        name,
			// 	ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
			// 	ImportState:         true,
			// 	ImportStateVerify:   true,
			// },
		},
	})
}
