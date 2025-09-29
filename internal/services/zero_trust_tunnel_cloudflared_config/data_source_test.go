package zero_trust_tunnel_cloudflared_config_test

import (
	"math/big"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func createPath(key string) tfjsonpath.Path {
	paths := strings.Split(key, ".")

	var path tfjsonpath.Path
	for _, p := range paths {
		n, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			path = path.AtMapKey(p)
		} else {
			path = path.AtSliceIndex(int(n))
		}
	}
	return path
}

func TestAccCloudflareTunnelCloudflaredConfigDatasource(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_tunnel_cloudflared_config." + rnd
	dataSourceName := "data.cloudflare_zero_trust_tunnel_cloudflared_config." + rnd
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	tunnelSecret := utils.RandStringFromCharSet(32, utils.CharSetAlpha)
	acctest.TestAccPreCheck_AccountID(t)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareTunnelConfigDataSource(rnd, accountID, tunnelSecret, domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, createPath("config.warp_routing.enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, createPath("config.origin_request.connect_timeout"), knownvalue.NumberExact(big.NewFloat(60))),
					statecheck.ExpectKnownValue(resourceName, createPath("config.origin_request.tls_timeout"), knownvalue.NumberExact(big.NewFloat(60))),
					statecheck.ExpectKnownValue(resourceName, createPath("config.origin_request.tcp_keep_alive"), knownvalue.NumberExact(big.NewFloat(60))),
					statecheck.ExpectKnownValue(resourceName, createPath("config.origin_request.no_happy_eyeballs"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, createPath("config.origin_request.keep_alive_connections"), knownvalue.NumberExact(big.NewFloat(1024))),
					statecheck.ExpectKnownValue(resourceName, createPath("config.origin_request.keep_alive_timeout"), knownvalue.NumberExact(big.NewFloat(60))),
					statecheck.ExpectKnownValue(resourceName, createPath("config.origin_request.http_host_header"), knownvalue.StringExact("baz")),
					statecheck.ExpectKnownValue(resourceName, createPath("config.origin_request.origin_server_name"), knownvalue.StringExact("foobar")),
					statecheck.ExpectKnownValue(resourceName, createPath("config.origin_request.ca_pool"), knownvalue.StringExact("/path/to/unsigned/ca/pool")),
					statecheck.ExpectKnownValue(resourceName, createPath("config.origin_request.no_tls_verify"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, createPath("config.origin_request.disable_chunked_encoding"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, createPath("config.origin_request.proxy_type"), knownvalue.StringExact("socks")),
					statecheck.ExpectKnownValue(resourceName, createPath("config.origin_request.http2_origin"), knownvalue.Bool(true)),

					statecheck.ExpectKnownValue(resourceName, createPath("config.ingress.0.hostname"), knownvalue.StringExact("foo")),
					statecheck.ExpectKnownValue(resourceName, createPath("config.ingress.0.path"), knownvalue.StringExact("/bar")),
					statecheck.ExpectKnownValue(resourceName, createPath("config.ingress.0.service"), knownvalue.StringExact("http://10.0.0.2:8080")),
					statecheck.ExpectKnownValue(resourceName, createPath("config.ingress.0.origin_request.connect_timeout"), knownvalue.NumberExact(big.NewFloat(15))),

					statecheck.ExpectKnownValue(resourceName, createPath("config.ingress.1.hostname"), knownvalue.StringExact("bar")),
					statecheck.ExpectKnownValue(resourceName, createPath("config.ingress.1.path"), knownvalue.StringExact("/foo")),
					statecheck.ExpectKnownValue(resourceName, createPath("config.ingress.1.service"), knownvalue.StringExact("http://10.0.0.3:8081")),
					statecheck.ExpectKnownValue(resourceName, createPath("config.ingress.1.origin_request.access.required"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, createPath("config.ingress.1.origin_request.access.team_name"), knownvalue.StringExact("terraform")),
					statecheck.ExpectKnownValue(resourceName, createPath("config.ingress.1.origin_request.access.aud_tag.0"), knownvalue.StringExact("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")),

					statecheck.ExpectKnownValue(resourceName, createPath("config.ingress.2.service"), knownvalue.StringExact("https://10.0.0.4:8082")),

					// Check data source attributes match resource attributes
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.warp_routing.enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.origin_request.connect_timeout"), knownvalue.NumberExact(big.NewFloat(60))),
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.origin_request.tls_timeout"), knownvalue.NumberExact(big.NewFloat(60))),
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.origin_request.tcp_keep_alive"), knownvalue.NumberExact(big.NewFloat(60))),
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.origin_request.no_happy_eyeballs"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.origin_request.keep_alive_connections"), knownvalue.NumberExact(big.NewFloat(1024))),
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.origin_request.keep_alive_timeout"), knownvalue.NumberExact(big.NewFloat(60))),
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.origin_request.http_host_header"), knownvalue.StringExact("baz")),
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.origin_request.origin_server_name"), knownvalue.StringExact("foobar")),
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.origin_request.ca_pool"), knownvalue.StringExact("/path/to/unsigned/ca/pool")),
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.origin_request.no_tls_verify"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.origin_request.disable_chunked_encoding"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.origin_request.proxy_type"), knownvalue.StringExact("socks")),
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.origin_request.http2_origin"), knownvalue.Bool(true)),

					statecheck.ExpectKnownValue(dataSourceName, createPath("config.ingress.0.hostname"), knownvalue.StringExact("foo")),
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.ingress.0.path"), knownvalue.StringExact("/bar")),
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.ingress.0.service"), knownvalue.StringExact("http://10.0.0.2:8080")),
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.ingress.0.origin_request.connect_timeout"), knownvalue.NumberExact(big.NewFloat(15))),

					statecheck.ExpectKnownValue(dataSourceName, createPath("config.ingress.1.hostname"), knownvalue.StringExact("bar")),
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.ingress.1.path"), knownvalue.StringExact("/foo")),
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.ingress.1.service"), knownvalue.StringExact("http://10.0.0.3:8081")),
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.ingress.1.origin_request.access.required"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.ingress.1.origin_request.access.team_name"), knownvalue.StringExact("terraform")),
					statecheck.ExpectKnownValue(dataSourceName, createPath("config.ingress.1.origin_request.access.aud_tag.0"), knownvalue.StringExact("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")),

					statecheck.ExpectKnownValue(dataSourceName, createPath("config.ingress.2.service"), knownvalue.StringExact("https://10.0.0.4:8082")),
				},
			},
		},
	})
}

func testCloudflareTunnelConfigDataSource(resourceID, accountID, tunnelSecret, domain string) string {
	return acctest.LoadTestCase("tunnelconfigdata.tf", resourceID, accountID, tunnelSecret, domain)
}
