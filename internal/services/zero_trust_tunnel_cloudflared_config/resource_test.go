package zero_trust_tunnel_cloudflared_config_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("connect_timeout"), knownvalue.Int64Exact(60)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("tls_timeout"), knownvalue.Int64Exact(60)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("tcp_keep_alive"), knownvalue.Int64Exact(60)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("no_happy_eyeballs"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("keep_alive_connections"), knownvalue.Int64Exact(1024)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("keep_alive_timeout"), knownvalue.Int64Exact(60)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("http_host_header"), knownvalue.StringExact("baz")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("origin_server_name"), knownvalue.StringExact("foobar")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("ca_pool"), knownvalue.StringExact("/path/to/unsigned/ca/pool")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("no_tls_verify"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("disable_chunked_encoding"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("proxy_type"), knownvalue.StringExact("socks")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("http2_origin"), knownvalue.Bool(true)),

					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.StringExact("foo")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(0).AtMapKey("path"), knownvalue.StringExact("/bar")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(0).AtMapKey("service"), knownvalue.StringExact("http://10.0.0.2:8080")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(0).AtMapKey("origin_request").AtMapKey("connect_timeout"), knownvalue.Int64Exact(15)),

					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(1).AtMapKey("hostname"), knownvalue.StringExact("bar")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(1).AtMapKey("path"), knownvalue.StringExact("/foo")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(1).AtMapKey("service"), knownvalue.StringExact("http://10.0.0.3:8081")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(1).AtMapKey("origin_request").AtMapKey("access").AtMapKey("required"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(1).AtMapKey("origin_request").AtMapKey("access").AtMapKey("team_name"), knownvalue.StringExact("terraform")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(1).AtMapKey("origin_request").AtMapKey("access").AtMapKey("aud_tag"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(1).AtMapKey("origin_request").AtMapKey("access").AtMapKey("aud_tag").AtSliceIndex(0), knownvalue.StringExact("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")),

					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(2).AtMapKey("service"), knownvalue.StringExact("https://10.0.0.4:8082")),
				},
			},
			{
				Config: testTunnelConfigUpdate(rnd, zoneID, tunnelSecret, domain),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("connect_timeout"), knownvalue.Int64Exact(61)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("tls_timeout"), knownvalue.Int64Exact(61)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("tcp_keep_alive"), knownvalue.Int64Exact(61)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("no_happy_eyeballs"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("keep_alive_connections"), knownvalue.Int64Exact(1028)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("keep_alive_timeout"), knownvalue.Int64Exact(61)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("http_host_header"), knownvalue.StringExact("bez")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("origin_server_name"), knownvalue.StringExact("fuuber")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("ca_pool"), knownvalue.StringExact("/path/to/unsigned/ca/pool")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("no_tls_verify"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("disable_chunked_encoding"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("proxy_type"), knownvalue.StringExact("socks")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("http2_origin"), knownvalue.Bool(true)),

					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.StringExact("fuu")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(0).AtMapKey("path"), knownvalue.StringExact("/ber")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(0).AtMapKey("service"), knownvalue.StringExact("http://10.0.0.3:8080")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(0).AtMapKey("origin_request").AtMapKey("connect_timeout"), knownvalue.Int64Exact(20)),

					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(1).AtMapKey("hostname"), knownvalue.StringExact("ber")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(1).AtMapKey("path"), knownvalue.StringExact("/fuu")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(1).AtMapKey("service"), knownvalue.StringExact("http://10.0.0.5:8081")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(1).AtMapKey("origin_request").AtMapKey("access").AtMapKey("required"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(1).AtMapKey("origin_request").AtMapKey("access").AtMapKey("team_name"), knownvalue.StringExact("terraform2")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(1).AtMapKey("origin_request").AtMapKey("access").AtMapKey("aud_tag"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(1).AtMapKey("origin_request").AtMapKey("access").AtMapKey("aud_tag").AtSliceIndex(0), knownvalue.StringExact("BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")),

					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(2).AtMapKey("service"), knownvalue.StringExact("https://10.0.0.5:8082")),
				},
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
				ImportStateVerifyIgnore: []string{},
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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(0).AtMapKey("service"), knownvalue.StringExact("https://10.0.0.1:8081")),
				},
			},
			{
				Config: testTunnelConfigShort(rnd, zoneID, tunnelSecret, "https://10.0.0.10:8081"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(name, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(0).AtMapKey("service"), knownvalue.StringExact("https://10.0.0.10:8081")),
				},
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
				ImportStateVerifyIgnore: []string{},
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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("no_tls_verify"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.StringExact("foo")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(0).AtMapKey("service"), knownvalue.StringExact("https://10.0.0.1:8006")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(1).AtMapKey("service"), knownvalue.StringExact("http_status:501")),
				},
			},
			{
				Config: testTunnelConfigNilPointerUpdate(rnd, zoneID, tunnelSecret),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("origin_request").AtMapKey("no_tls_verify"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.StringExact("bar")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(0).AtMapKey("service"), knownvalue.StringExact("https://10.0.0.10:8006")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("ingress").AtSliceIndex(1).AtMapKey("service"), knownvalue.StringExact("http_status:502")),
				},
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
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}
