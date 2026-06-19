package turnstile_widget_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_turnstile_widget", &resource.Sweeper{
		Name: "cloudflare_turnstile_widget",
		F: func(region string) error {
			ctx := context.Background()
			client, err := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", err))
				return fmt.Errorf("error establishing client: %w", err)
			}

			if accountID == "" {
				tflog.Info(ctx, "Skipping turnstile widgets sweep: CLOUDFLARE_ACCOUNT_ID not set")
				return nil
			}

			widgets, _, err := client.ListTurnstileWidgets(ctx, cfv1.AccountIdentifier(accountID), cfv1.ListTurnstileWidgetParams{})
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to fetch turnstile widgets: %s", err))
				return fmt.Errorf("failed to fetch turnstile widgets: %w", err)
			}

			if len(widgets) == 0 {
				tflog.Info(ctx, "No turnstile widgets to sweep")
				return nil
			}

			for _, widget := range widgets {
				if !utils.ShouldSweepResource(widget.Name) {
					continue
				}
				tflog.Info(ctx, fmt.Sprintf("Deleting turnstile widget: %s (%s) (account: %s)", widget.Name, widget.SiteKey, accountID))
				err := client.DeleteTurnstileWidget(ctx, cfv1.AccountIdentifier(accountID), widget.SiteKey)
				if err != nil {
					tflog.Error(ctx, fmt.Sprintf("Failed to delete turnstile widget %s: %s", widget.SiteKey, err))
					continue
				}
				tflog.Info(ctx, fmt.Sprintf("Deleted turnstile widget: %s", widget.SiteKey))
			}

			return nil
		},
	})
}

func TestAccCloudflareTurnstileWidget_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_turnstile_widget." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareTurnstileWidgetBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "bot_fight_mode", "false"),
					resource.TestCheckResourceAttr(resourceName, "domains.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "domains.0", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "mode", "invisible"),
					resource.TestCheckResourceAttr(resourceName, "region", "world"),
					resource.TestCheckResourceAttr(resourceName, "offlabel", "false"),
				),
			},
			{
				ResourceName:        resourceName,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

// TestAccCloudflareTurnstileWidget_NoDomainReorderDrift verifies that listing
// the same domains in a different order is a no-op. The API returns domains
// sorted, so without ModifyPlan a reorder would plan a spurious update.
func TestAccCloudflareTurnstileWidget_NoDomainReorderDrift(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_turnstile_widget." + rnd

	noop := resource.ConfigPlanChecks{
		PreApply: []plancheck.PlanCheck{
			plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareTurnstileWidgetDomainOrder(rnd, accountID),
			},
			{
				// Same domains, same order: no-op.
				Config:           testAccCheckCloudflareTurnstileWidgetDomainOrder(rnd, accountID),
				ConfigPlanChecks: noop,
			},
			{
				// Same domains in a different order: no-op.
				Config:           testAccCheckCloudflareTurnstileWidgetDomainReorder(rnd, accountID),
				ConfigPlanChecks: noop,
			},
			{
				// A real domain change must still plan as an update.
				Config: testAccCheckCloudflareTurnstileWidgetDomainChanged(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
			},
		},
	})
}

func TestAccCloudflareTurnstileWidget_Minimum(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_turnstile_widget." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareTurnstileWidgetMinimum(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "bot_fight_mode", "false"),
					resource.TestCheckResourceAttr(resourceName, "domains.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "domains.0", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "mode", "managed"),
					resource.TestCheckResourceAttr(resourceName, "region", "world"),
					resource.TestCheckResourceAttr(resourceName, "offlabel", "false"),
				),
			},
			{
				ResourceName:        resourceName,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func TestAccCloudflareTurnstileWidget_NoDomains(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_turnstile_widget." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareTurnstileWidgetNoDomains(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "bot_fight_mode", "false"),
					resource.TestCheckResourceAttr(resourceName, "domains.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "mode", "managed"),
					resource.TestCheckResourceAttr(resourceName, "region", "world"),
					resource.TestCheckResourceAttr(resourceName, "offlabel", "false"),
				),
			},
			{
				ResourceName:        resourceName,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func TestAccCloudflareTurnstileWidget_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_turnstile_widget." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareTurnstileWidgetBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "mode", "invisible"),
					resource.TestCheckResourceAttr(resourceName, "domains.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "domains.0", "example.com"),
				),
			},
			{
				Config: testAccCheckCloudflareTurnstileWidgetUpdated(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd+"-updated"),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "mode", "invisible"),
					resource.TestCheckResourceAttr(resourceName, "domains.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "domains.0", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "domains.1", "test.example.com"),
				),
			},
			{
				ResourceName:        resourceName,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func TestAccCloudflareTurnstileWidget_NonInteractiveMode(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_turnstile_widget." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareTurnstileWidgetNonInteractive(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "mode", "non-interactive"),
					resource.TestCheckResourceAttr(resourceName, "domains.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "domains.0", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "region", "world"),
				),
			},
			{
				ResourceName:        resourceName,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func testAccCheckCloudflareTurnstileWidgetBasic(rnd, accountID string) string {
	return acctest.LoadTestCase("turnstilewidgetbasic.tf", rnd, accountID)
}

func testAccCheckCloudflareTurnstileWidgetDomainOrder(rnd, accountID string) string {
	return acctest.LoadTestCase("turnstilewidgetdomainorder.tf", rnd, accountID)
}

func testAccCheckCloudflareTurnstileWidgetDomainReorder(rnd, accountID string) string {
	return acctest.LoadTestCase("turnstilewidgetdomainreorder.tf", rnd, accountID)
}

func testAccCheckCloudflareTurnstileWidgetDomainChanged(rnd, accountID string) string {
	return acctest.LoadTestCase("turnstilewidgetdomainchanged.tf", rnd, accountID)
}

func testAccCheckCloudflareTurnstileWidgetMinimum(rnd, accountID string) string {
	return acctest.LoadTestCase("turnstilewidgetminimum.tf", rnd, accountID)
}

func testAccCheckCloudflareTurnstileWidgetNoDomains(rnd, accountID string) string {
	return acctest.LoadTestCase("turnstilewidgetnodomains.tf", rnd, accountID)
}

func testAccCheckCloudflareTurnstileWidgetUpdated(rnd, accountID string) string {
	return acctest.LoadTestCase("turnstilewidgetupdated.tf", rnd, accountID)
}

func testAccCheckCloudflareTurnstileWidgetNonInteractive(rnd, accountID string) string {
	return acctest.LoadTestCase("turnstilewidgetnoninteractive.tf", rnd, accountID)
}
