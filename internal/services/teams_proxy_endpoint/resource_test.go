package teams_proxy_endpoint_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_teams_proxy_endpoint.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsProxyEndpointDestroy,
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
	return acctest.LoadTestCase("teamsproxyendpointconfigbasic.tf", rnd, accountID)
}

func testAccCheckCloudflareTeamsProxyEndpointDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

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
