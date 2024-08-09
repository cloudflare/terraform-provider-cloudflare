package risk_behavior_test

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
	resource.AddTestSweepers("cloudflare_zero_trust_risk_behavior", &resource.Sweeper{
		Name: "cloudflare_zero_trust_risk_behavior",
		F: func(region string) error {
			client, err := acctest.SharedV1Client()
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

			if err != nil {
				return fmt.Errorf("error establishing client: %w", err)
			}

			ctx := context.Background()

			behaviors, err := client.Behaviors(ctx, accountID)
			if err != nil {
				return fmt.Errorf("failed to get risk behaviors: %w", err)
			}

			// set all risk behavior values to false/low before running update
			for _, behavior := range behaviors.Behaviors {
				behavior.Enabled = cloudflare.BoolPtr(false)
				behavior.RiskLevel = cloudflare.Low
			}

			_, err = client.UpdateBehaviors(ctx, accountID, behaviors)
			if err != nil {
				return fmt.Errorf("failed to reset risk behaviors: %w", err)
			}

			return nil
		},
	})
}

func TestAccCloudflareRiskBehavior_Partial(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_risk_behavior." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRiskBehaviorsPartial(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "behavior.#", "1"),
				),
			},
		},
	})
}

func testAccCloudflareRiskBehaviorsPartial(name, accountId string) string {
	return fmt.Sprintf(`
	resource cloudflare_zero_trust_risk_behavior %s {
		account_id = "%s"
		behavior {
			name = "imp_travel"
			enabled = true
			risk_level = "high"
		}
	}`, name, accountId)
}

func TestAccCloudflareRiskBehavior_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_risk_behavior." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareRiskBehaviors(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "behavior.#", "2"),
				),
			},
		},
	})
}

func testAccCloudflareRiskBehaviors(name, accountId string) string {
	return fmt.Sprintf(`
	resource cloudflare_zero_trust_risk_behavior %s {
		account_id = "%s"
		behavior {
			name = "imp_travel"
			enabled = true
			risk_level = "high"
		}
		behavior {
			name = "high_dlp"
			enabled = true
			risk_level = "medium"
		}
	}`, name, accountId)
}
