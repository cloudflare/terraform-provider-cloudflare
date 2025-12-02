package zero_trust_access_ai_controls_mcp_portal_test

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
	resource.AddTestSweepers("cloudflare_zero_trust_access_ai_controls_mcp_portal", &resource.Sweeper{
		Name: "cloudflare_zero_trust_access_ai_controls_mcp_portal",
		F:    testSweepCloudflareZeroTrustAccessAIControlsMcpPortal,
	})
}

func testSweepCloudflareZeroTrustAccessAIControlsMcpPortal(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	portalsResp, err := client.ZeroTrust.Access.AIControls.Mcp.Portals.List(
		ctx,
		zero_trust.AccessAIControlMcpPortalListParams{
			AccountID: cloudflare.F(accountID),
		},
	)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Zero Trust Access AI Controls MCP Portals: %s", err))
		return err
	}

	for _, portal := range portalsResp.Result {
		if !utils.ShouldSweepResource(portal.Name) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting Zero Trust Access AI Controls MCP Portal: %s", portal.ID))
		_, err := client.ZeroTrust.Access.AIControls.Mcp.Portals.Delete(
			ctx,
			portal.ID,
			zero_trust.AccessAIControlMcpPortalDeleteParams{
				AccountID: cloudflare.F(accountID),
			},
		)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete Zero Trust Access AI Controls MCP Portal %s: %s", portal.ID, err))
		}
	}

	return nil
}

func TestAccZeroTrustAccessAIControlsMcpPortal_basic(t *testing.T) {
	resourceName := "cloudflare_zero_trust_access_ai_controls_mcp_portal.tf-test"
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	name1 := "Test Portal"
	name2 := "Updated Test Portal"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessAIControlsMcpPortalDestroy,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: acctest.LoadTestCase("basic.tf", accountID, domain, name1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", "tf-test"),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "hostname", domain),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read testing
			{
				Config: acctest.LoadTestCase("basic.tf", accountID, domain, name2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", "tf-test"),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "hostname", domain),
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
			},
		},
	})
}

func testAccCheckCloudflareZeroTrustAccessAIControlsMcpPortalDestroy(s *terraform.State) error {
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_access_ai_controls_mcp_portal" {
			continue
		}

		_, err := client.ZeroTrust.Access.AIControls.Mcp.Portals.Read(
			context.Background(),
			rs.Primary.ID,
			zero_trust.AccessAIControlMcpPortalReadParams{
				AccountID: cloudflare.F(accountID),
			},
		)

		if err == nil {
			return fmt.Errorf("Zero Trust Access AI Controls Mcp Portal %s still exists", rs.Primary.ID)
		}
	}

	return nil
}
