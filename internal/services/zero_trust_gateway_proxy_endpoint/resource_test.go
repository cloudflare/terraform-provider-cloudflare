package zero_trust_gateway_proxy_endpoint_test

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

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_gateway_proxy_endpoint", &resource.Sweeper{
		Name: "cloudflare_zero_trust_gateway_proxy_endpoint",
		F:    testSweepCloudflareZeroTrustGatewayProxyEndpoint,
	})
}

func testSweepCloudflareZeroTrustGatewayProxyEndpoint(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	endpoints, _, err := client.TeamsProxyEndpoints(ctx, accountID)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Zero Trust Gateway Proxy Endpoints: %s", err))
		return err
	}

	for _, endpoint := range endpoints {
		if !utils.ShouldSweepResource(endpoint.Name) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting Zero Trust Gateway Proxy Endpoint: %s", endpoint.ID))
		err := client.DeleteTeamsProxyEndpoint(ctx, accountID, endpoint.ID)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete Zero Trust Gateway Proxy Endpoint %s: %s", endpoint.ID, err))
		}
	}

	return nil
}

func TestAccCloudflareTeamsProxyEndpoint_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_gateway_proxy_endpoint.%s", rnd)
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
		if rs.Type != "cloudflare_zero_trust_gateway_proxy_endpoint" {
			continue
		}

		_, err := client.TeamsProxyEndpoint(context.Background(), rs.Primary.Attributes[consts.AccountIDSchemaKey], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("teams Proxy Endpoint still exists")
		}
	}

	return nil
}
