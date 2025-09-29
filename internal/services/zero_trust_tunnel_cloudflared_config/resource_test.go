package zero_trust_tunnel_cloudflared_config_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func testTunnelConfig(resourceID, accountID, tunnelSecret, domain string) string {
	return acctest.LoadTestCase("tunnelconfig.tf", resourceID, accountID, tunnelSecret, domain)
}

func testTunnelConfigUpdate(resourceID, accountID, tunnelSecret, domain string) string {
	return acctest.LoadTestCase("tunnelconfigupdate.tf", resourceID, accountID, tunnelSecret, domain)
}

func testTunnelConfigShort(resourceID, accountID, tunnelSecret, service string) string {
	return acctest.LoadTestCase("tunnelconfigshort.tf", resourceID, accountID, tunnelSecret, service)
}

func testTunnelConfigNilPointer(resourceID, accountID, tunnelSecret string) string {
	return acctest.LoadTestCase("tunnelconfignilpointer.tf", resourceID, accountID, tunnelSecret)
}

func testTunnelConfigNilPointerUpdate(resourceID, accountID, tunnelSecret string) string {
	return acctest.LoadTestCase("tunnelconfignilpointerupdate.tf", resourceID, accountID, tunnelSecret)
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
					resource.TestCheckResourceAttr(name, "config.warp_routing.enabled", "false"),
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
			{
				Config: testTunnelConfigUpdate(rnd, zoneID, tunnelSecret, domain),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "config.warp_routing.enabled", "false"),
					resource.TestCheckResourceAttr(name, "config.origin_request.connect_timeout", "61"),
					resource.TestCheckResourceAttr(name, "config.origin_request.tls_timeout", "61"),
					resource.TestCheckResourceAttr(name, "config.origin_request.tcp_keep_alive", "61"),
					resource.TestCheckResourceAttr(name, "config.origin_request.no_happy_eyeballs", "true"),
					resource.TestCheckResourceAttr(name, "config.origin_request.keep_alive_connections", "1028"),
					resource.TestCheckResourceAttr(name, "config.origin_request.keep_alive_timeout", "61"),
					resource.TestCheckResourceAttr(name, "config.origin_request.http_host_header", "bez"),
					resource.TestCheckResourceAttr(name, "config.origin_request.origin_server_name", "fuuber"),
					resource.TestCheckResourceAttr(name, "config.origin_request.ca_pool", "/path/to/unsigned/ca/pool"),
					resource.TestCheckResourceAttr(name, "config.origin_request.no_tls_verify", "false"),
					resource.TestCheckResourceAttr(name, "config.origin_request.disable_chunked_encoding", "false"),
					resource.TestCheckResourceAttr(name, "config.origin_request.proxy_type", "socks"),
					resource.TestCheckResourceAttr(name, "config.origin_request.http2_origin", "true"),

					resource.TestCheckResourceAttr(name, "config.ingress.0.hostname", "fuu"),
					resource.TestCheckResourceAttr(name, "config.ingress.0.path", "/ber"),
					resource.TestCheckResourceAttr(name, "config.ingress.0.service", "http://10.0.0.3:8080"),
					resource.TestCheckResourceAttr(name, "config.ingress.0.origin_request.connect_timeout", "20"),

					resource.TestCheckResourceAttr(name, "config.ingress.1.hostname", "ber"),
					resource.TestCheckResourceAttr(name, "config.ingress.1.path", "/fuu"),
					resource.TestCheckResourceAttr(name, "config.ingress.1.service", "http://10.0.0.5:8081"),
					resource.TestCheckResourceAttr(name, "config.ingress.1.origin_request.access.required", "false"),
					resource.TestCheckResourceAttr(name, "config.ingress.1.origin_request.access.team_name", "terraform2"),
					resource.TestCheckResourceAttr(name, "config.ingress.1.origin_request.access.aud_tag.#", "1"),
					resource.TestCheckResourceAttr(name, "config.ingress.1.origin_request.access.aud_tag.0", "BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"),

					resource.TestCheckResourceAttr(name, "config.ingress.2.service", "https://10.0.0.5:8082"),
				),
			},
			{
				Config: testTunnelConfigUpdate(rnd, zoneID, tunnelSecret, domain),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
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
				Config: testTunnelConfigShort(rnd, zoneID, tunnelSecret, "https://10.0.0.1:8081"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "config.ingress.0.service", "https://10.0.0.1:8081"),
				),
			},
			{
				Config: testTunnelConfigShort(rnd, zoneID, tunnelSecret, "https://10.0.0.10:8081"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "config.ingress.0.service", "https://10.0.0.10:8081"),
				),
			},
			{
				Config: testTunnelConfigShort(rnd, zoneID, tunnelSecret, "https://10.0.0.10:8081"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
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
					resource.TestCheckResourceAttr(name, "config.warp_routing.enabled", "false"),
					resource.TestCheckResourceAttr(name, "config.origin_request.no_tls_verify", "true"),
					resource.TestCheckResourceAttr(name, "config.ingress.0.hostname", "foo"),
					resource.TestCheckResourceAttr(name, "config.ingress.0.service", "https://10.0.0.1:8006"),
					resource.TestCheckResourceAttr(name, "config.ingress.1.service", "http_status:501"),
				),
			},
			{
				Config: testTunnelConfigNilPointerUpdate(rnd, zoneID, tunnelSecret),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "config.warp_routing.enabled", "false"),
					resource.TestCheckResourceAttr(name, "config.origin_request.no_tls_verify", "false"),
					resource.TestCheckResourceAttr(name, "config.ingress.0.hostname", "bar"),
					resource.TestCheckResourceAttr(name, "config.ingress.0.service", "https://10.0.0.10:8006"),
					resource.TestCheckResourceAttr(name, "config.ingress.1.service", "http_status:502"),
				),
			},
			{
				Config: testTunnelConfigNilPointerUpdate(rnd, zoneID, tunnelSecret),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}
