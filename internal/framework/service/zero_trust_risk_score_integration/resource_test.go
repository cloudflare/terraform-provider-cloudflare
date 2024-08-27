package zero_trust_risk_score_integration_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_risk_score_integration", &resource.Sweeper{
		Name: "cloudflare_zero_trust_risk_score_integration",
		F: func(region string) error {
			client, err := acctest.SharedV1Client()
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

			if err != nil {
				return fmt.Errorf("error establishing client: %w", err)
			}

			ctx := context.Background()

			integrations, err := client.ListRiskScoreIntegrations(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListRiskScoreIntegrationParams{})
			if err != nil {
				return fmt.Errorf("failed to get risk score integrations: %w", err)
			}

			// Clean up old integrations
			for _, integration := range integrations {
				err := client.DeleteRiskScoreIntegration(ctx, cloudflare.AccountIdentifier(accountID), integration.ID)
				if err != nil {
					return fmt.Errorf("failed to delete risk score integration: %w", err)
				}
			}

			return nil
		},
	})
}

func TestAccCloudflareRiskScoreIntegration_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_risk_score_integration." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource cloudflare_zero_trust_risk_score_integration %s {
					account_id = "%s"
					integration_type = "Okta"
					tenant_url = "https://test-tenant.okta.com"
				}`, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "integration_type", "Okta"),
					resource.TestCheckResourceAttr(name, "tenant_url", "https://test-tenant.okta.com"),
					resource.TestCheckResourceAttr(name, "active", "true"), // Test function uses the stringified version for comparison
				),
			},
		},
	})
}

func TestAccCloudflareRiskScoreIntegration_Reference_ID(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_risk_score_integration." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource cloudflare_zero_trust_risk_score_integration %s {
					account_id = "%s"
					integration_type = "Okta"
					tenant_url = "https://test-tenant.okta.com"
					reference_id = "58ee8f00-f28a-4b09-b955-93f7c557bd43"
				}`, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "reference_id", "58ee8f00-f28a-4b09-b955-93f7c557bd43"),
				),
			},
		},
	})
}
