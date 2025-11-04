package zero_trust_access_ai_controls_mcp_portal_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccZeroTrustAccessAIControlsMcpPortal_basic(t *testing.T) {
	resourceName := "cloudflare_zero_trust_access_ai_controls_mcp_portal.tf-test"
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	hostname := "tf-test.playground.that-test.site"
	name1 := "Test Portal"
	name2 := "Updated Test Portal"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZeroTrustAccessAIControlsMcpPortalDestroy,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: acctest.LoadTestCase("basic.tf", accountID, hostname, name1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", "tf-test"),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "hostname", hostname),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read testing
			{
				Config: acctest.LoadTestCase("basic.tf", accountID, hostname, name2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", "tf-test"),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "hostname", hostname),
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
