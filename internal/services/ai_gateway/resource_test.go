package ai_gateway_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/ai_gateway"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

const resourcePrefix = "tfacctest-"

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_ai_gateway", &resource.Sweeper{
		Name: "cloudflare_ai_gateway",
		F:    testSweepCloudflareAIGateways,
	})
}

func testSweepCloudflareAIGateways(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if accountID == "" {
		tflog.Info(ctx, "Skipping ai_gateway sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	list, err := client.AIGateway.List(ctx, ai_gateway.AIGatewayListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list AI gateways: %s", err))
		return fmt.Errorf("failed to list AI gateways: %w", err)
	}

	hasGateways := false
	for _, gw := range list.Result {
		hasGateways = true
		if !utils.ShouldSweepResource(gw.ID) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting AI gateway: %s (account: %s)", gw.ID, accountID))
		_, err := client.AIGateway.Delete(ctx, gw.ID, ai_gateway.AIGatewayDeleteParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete AI gateway %s: %s", gw.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted AI gateway: %s", gw.ID))
	}

	if !hasGateways {
		tflog.Info(ctx, "No AI gateways to sweep")
	}

	return nil
}

func TestAccCloudflareAIGateway_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := resourcePrefix + rnd
	name := "cloudflare_ai_gateway." + resourceName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAIGatewayConfig(resourceName, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.StringExact(resourceName)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("modified_at"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s", accountID, resourceName), nil
				},
			},
		},
	})
}

func testAccCloudflareAIGatewayConfig(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_ai_gateway" "%s" {
  account_id = "%s"
  id = "%s"
}
`, rnd, accountID, rnd)
}
