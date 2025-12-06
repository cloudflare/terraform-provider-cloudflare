package sso_connector_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/iam"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_sso_connector", &resource.Sweeper{
		Name: "cloudflare_sso_connector",
		F:    testSweepCloudflareSSOConnectors,
	})
}

func testSweepCloudflareSSOConnectors(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		tflog.Info(ctx, "Skipping SSO connectors sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	connectors, err := client.IAM.SSO.List(ctx, iam.SSOListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch SSO connectors: %s", err))
		return fmt.Errorf("failed to fetch SSO connectors: %w", err)
	}

	if len(connectors.Result) == 0 {
		tflog.Info(ctx, "No SSO connectors to sweep")
		return nil
	}

	for _, connector := range connectors.Result {
		// Use standard filtering helper on the email domain field
		if !utils.ShouldSweepResource(connector.EmailDomain) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting SSO connector: %s (email_domain: %s, account: %s)", connector.ID, connector.EmailDomain, accountID))
		_, err := client.IAM.SSO.Delete(ctx, connector.ID, iam.SSODeleteParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete SSO connector %s: %s", connector.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted SSO connector: %s", connector.ID))
	}

	return nil
}

func TestAccCloudflareSsoConnector_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_sso_connector." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSsoConnectorConfig(rnd, accountID, false),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email_domain"), knownvalue.StringExact(fmt.Sprintf("%s.example.com", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("updated_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("use_fedramp_language"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("verification"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccSsoConnectorConfig(rnd, accountID, true),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email_domain"), knownvalue.StringExact(fmt.Sprintf("%s.example.com", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("use_fedramp_language"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:              resourceName,
				ImportStateIdPrefix:       fmt.Sprintf("%s/", accountID),
				ImportState:               true,
				ImportStateVerify:         true,
				ImportStateVerifyIgnore:   []string{"begin_verification"},
			},
		},
	})
}

func testAccSsoConnectorConfig(rnd, accountID string, useFedramp bool) string {
	if useFedramp {
		return acctest.LoadTestCase("with_fedramp_language.tf", rnd, accountID)
	}
	return acctest.LoadTestCase("basic.tf", rnd, accountID)
}
