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
)

const aiGatewayResourcePrefix = "tfacctest-ai-gateway-"

func init() {
	resource.AddTestSweepers("cloudflare_ai_gateway", &resource.Sweeper{
		Name: "cloudflare_ai_gateway",
		F:    testSweepCloudflareAIGateway,
	})
}

func testSweepCloudflareAIGateway(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if accountID == "" {
		tflog.Info(ctx, "Skipping AI Gateway sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	list, err := client.AIGateway.List(ctx, ai_gateway.AIGatewayListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list AI Gateways: %s", err))
		return fmt.Errorf("failed to list AI Gateways: %w", err)
	}

	if len(list.Result) == 0 {
		tflog.Info(ctx, "No AI Gateways to sweep")
		return nil
	}

	for _, gw := range list.Result {
		if !utils.ShouldSweepResource(gw.ID) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting AI Gateway: %s (account: %s)", gw.ID, accountID))
		_, err := client.AIGateway.Delete(ctx, gw.ID, ai_gateway.AIGatewayDeleteParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete AI Gateway %s: %s", gw.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted AI Gateway: %s", gw.ID))
	}

	return nil
}

func TestAccCloudflareAIGateway_Basic(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	resourceName := aiGatewayResourcePrefix + rnd
	name := "cloudflare_ai_gateway." + resourceName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAIGatewayConfigBasic(resourceName, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "id", resourceName),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttrSet(name, "created_at"),
				),
			},
			{
				ResourceName:            name,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"created_at", "modified_at"},
			},
		},
	})
}

func TestAccCloudflareAIGateway_WithRateLimiting(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	resourceName := aiGatewayResourcePrefix + rnd
	name := "cloudflare_ai_gateway." + resourceName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAIGatewayConfigWithRateLimiting(resourceName, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "id", resourceName),
					resource.TestCheckResourceAttr(name, "rate_limiting_limit", "100"),
					resource.TestCheckResourceAttr(name, "rate_limiting_interval", "60"),
					resource.TestCheckResourceAttr(name, "rate_limiting_technique", "fixed"),
				),
			},
			{
				Config: testAccCloudflareAIGatewayConfigWithUpdatedRateLimiting(resourceName, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "rate_limiting_limit", "200"),
					resource.TestCheckResourceAttr(name, "rate_limiting_interval", "120"),
					resource.TestCheckResourceAttr(name, "rate_limiting_technique", "sliding"),
				),
			},
		},
	})
}

func TestAccCloudflareAIGateway_WithDLP(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	resourceName := aiGatewayResourcePrefix + rnd
	name := "cloudflare_ai_gateway." + resourceName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAIGatewayConfigWithDLP(resourceName, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "dlp.0.action", "BLOCK"),
					resource.TestCheckResourceAttr(name, "dlp.0.enabled", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareAIGateway_WithOpenTelemetry(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	resourceName := aiGatewayResourcePrefix + rnd
	name := "cloudflare_ai_gateway." + resourceName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAIGatewayConfigWithOpenTelemetry(resourceName, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "otel.0.url", "https://otel.example.com"),
					resource.TestCheckResourceAttr(name, "otel.0.content_type", "json"),
				),
			},
		},
	})
}

func TestAccCloudflareAIGateway_WithWorkersAI(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	resourceName := aiGatewayResourcePrefix + rnd
	name := "cloudflare_ai_gateway." + resourceName
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAIGatewayConfigWithWorkersAI(resourceName, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "workers_ai_billing_mode", "prepaid"),
				),
			},
		},
	})
}

func testAccCloudflareAIGatewayConfigBasic(resourceName, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_ai_gateway" "%s" {
  account_id = "%s"
  id        = "%s"
}
`, resourceName, accountID, resourceName)
}

func testAccCloudflareAIGatewayConfigWithRateLimiting(resourceName, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_ai_gateway" "%s" {
  account_id               = "%s"
  id                       = "%s"
  rate_limiting_limit      = 100
  rate_limiting_interval   = 60
  rate_limiting_technique  = "fixed"
}
`, resourceName, accountID, resourceName)
}

func testAccCloudflareAIGatewayConfigWithUpdatedRateLimiting(resourceName, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_ai_gateway" "%s" {
  account_id               = "%s"
  id                       = "%s"
  rate_limiting_limit      = 200
  rate_limiting_interval   = 120
  rate_limiting_technique  = "sliding"
}
`, resourceName, accountID, resourceName)
}

func testAccCloudflareAIGatewayConfigWithDLP(resourceName, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_ai_gateway" "%s" {
  account_id = "%s"
  id         = "%s"

  dlp {
    action   = "BLOCK"
    enabled  = true
    profiles = ["profile1", "profile2"]
  }
}
`, resourceName, accountID, resourceName)
}

func testAccCloudflareAIGatewayConfigWithOpenTelemetry(resourceName, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_ai_gateway" "%s" {
  account_id = "%s"
  id         = "%s"

  otel {
    url          = "https://otel.example.com"
    content_type = "json"
    headers = {
      "X-Custom-Header" = "value"
    }
  }
}
`, resourceName, accountID, resourceName)
}

func testAccCloudflareAIGatewayConfigWithWorkersAI(resourceName, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_ai_gateway" "%s" {
  account_id             = "%s"
  id                     = "%s"
  workers_ai_billing_mode = "prepaid"
}
`, resourceName, accountID, resourceName)
}
