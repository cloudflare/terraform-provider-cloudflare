package argo_tunnel_test

import (
	"regexp"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccArgoTunnelResource_removed(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "cloudflare_argo_tunnel" "test" {
						account_id = "test-account-id"
						name = "test-tunnel"
						secret = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
					}
				`,
				ExpectError: regexp.MustCompile(`Resource Deprecated`),
			},
		},
	})
}
