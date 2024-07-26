package turnstile_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_turnstile_widget", &resource.Sweeper{
		Name: "cloudflare_turnstile_widget",
		F: func(region string) error {
			client, err := acctest.SharedV1Client()
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

			if err != nil {
				return fmt.Errorf("error establishing client: %w", err)
			}

			ctx := context.Background()
			widgets, _, err := client.ListTurnstileWidgets(ctx, cfv1.AccountIdentifier(accountID), cfv1.ListTurnstileWidgetParams{})
			if err != nil {
				return fmt.Errorf("failed to fetch turnstile widgets: %w", err)
			}

			for _, widget := range widgets {
				err := client.DeleteTurnstileWidget(ctx, cfv1.AccountIdentifier(accountID), widget.SiteKey)
				if err != nil {
					return fmt.Errorf("failed to delete turnstile widget %q: %w", widget.SiteKey, err)
				}
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
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
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
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
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
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
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

func testAccCheckCloudflareTurnstileWidgetBasic(rnd, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_turnstile_widget" "%[1]s" {
    account_id     = "%[2]s"
    name        = "%[1]s"
	bot_fight_mode = false
	domains = [ "example.com" ]
	mode = "invisible"
	region = "world"
  }`, rnd, accountID)
}

func testAccCheckCloudflareTurnstileWidgetMinimum(rnd, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_turnstile_widget" "%[1]s" {
    account_id     = "%[2]s"
    name        = "%[1]s"
	domains = [ "example.com" ]
	mode = "managed"
  }`, rnd, accountID)
}

func testAccCheckCloudflareTurnstileWidgetNoDomains(rnd, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_turnstile_widget" "%[1]s" {
    account_id     = "%[2]s"
    name        = "%[1]s"
	domains = [ ]
	mode = "managed"
  }`, rnd, accountID)
}
