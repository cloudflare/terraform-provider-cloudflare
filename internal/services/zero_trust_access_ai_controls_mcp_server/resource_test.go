package zero_trust_access_ai_controls_mcp_server_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_access_ai_controls_mcp_server", &resource.Sweeper{
		Name: "cloudflare_zero_trust_access_ai_controls_mcp_server",
		F:    testSweepCloudflareZeroTrustAccessAIControlsMcpServer,
	})
}

func testSweepCloudflareZeroTrustAccessAIControlsMcpServer(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	serversResp, err := client.ZeroTrust.Access.AIControls.Mcp.Servers.List(
		ctx,
		zero_trust.AccessAIControlMcpServerListParams{
			AccountID: cloudflare.F(accountID),
		},
	)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Zero Trust Access AI Controls MCP Servers: %s", err))
		return err
	}

	for _, server := range serversResp.Result {
		if !utils.ShouldSweepResource(server.Name) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting Zero Trust Access AI Controls MCP Server: %s", server.ID))
		_, err := client.ZeroTrust.Access.AIControls.Mcp.Servers.Delete(
			ctx,
			server.ID,
			zero_trust.AccessAIControlMcpServerDeleteParams{
				AccountID: cloudflare.F(accountID),
			},
		)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete Zero Trust Access AI Controls MCP Server %s: %s", server.ID, err))
		}
	}

	return nil
}

func TestAccZeroTrustAccessAIControlsMcpServer_basic(t *testing.T) {
	resourceName := "cloudflare_zero_trust_access_ai_controls_mcp_server.tf-test"
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	mcpUrl := "https://docs.mcp.cloudflare.com/mcp"
	name1 := "Test Server"
	name2 := "Updated Test Server"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessAIControlsMcpServerDestroy,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: acctest.LoadTestCase("basic.tf", accountID, mcpUrl, name1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", "tf-test"),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "hostname", mcpUrl),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read testing
			{
				Config: acctest.LoadTestCase("basic.tf", accountID, mcpUrl, name2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", "tf-test"),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "hostname", mcpUrl),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// ImportState testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources[resourceName]
					if !ok {
						return "", fmt.Errorf("not found: %s", resourceName)
					}
					return fmt.Sprintf("%s/%s", rs.Primary.Attributes["account_id"], rs.Primary.ID), nil
				},
				ImportStateVerifyIgnore: []string{"last_synced"},
			},
		},
	})
}

func testAccCheckCloudflareZeroTrustAccessAIControlsMcpServerDestroy(s *terraform.State) error {
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_access_ai_controls_mcp_server" {
			continue
		}

		_, err := client.ZeroTrust.Access.AIControls.Mcp.Servers.Read(
			context.Background(),
			rs.Primary.ID,
			zero_trust.AccessAIControlMcpServerReadParams{
				AccountID: cloudflare.F(accountID),
			},
		)

		if err == nil {
			return fmt.Errorf("Zero Trust Access AI Controls Mcp Server %s still exists", rs.Primary.ID)
		}
	}

	return nil
}
