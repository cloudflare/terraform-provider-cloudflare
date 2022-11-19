package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func testChallengeWidget(resourceID, accountID, widgetType string) string {
	return fmt.Sprintf(`
		resource "cloudflare_challenge_widget" "%[1]s" {
		  account_id = "%[2]s"
		  name = "%[1]s"
		  type = "%[3]s"
		  domains = ["example.com"]
		}
		`, resourceID, accountID, widgetType)
}

func testChallengeWidgetImport(resourceID, accountID, widgetType, siteKey, Secret string) string {
	return fmt.Sprintf(`
		resource "cloudflare_challenge_widget" "%[1]s" {
		  account_id = "%[2]s"
		  site_key = "%[4]s"
          secret = "%[5]s"
		  name = "%[1]s"
		  type = "%[3]s"
		  domains = ["example.com"]
		}
		`, resourceID, accountID, widgetType, siteKey, Secret)
}

func TestAccCloudflareChallengeWidget(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "cloudflare_challenge_widget." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testChallengeWidget(rnd, accountID, "invisible"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "invisible"),
					resource.TestCheckResourceAttr(name, "domains.#", "1"),
					resource.TestCheckResourceAttr(name, "domains.0", "example.com"),
					resource.TestCheckResourceAttrSet(name, "site_key"),
					resource.TestCheckResourceAttrSet(name, "secret"),
				),
			},
		},
	})
}

func TestAccCloudflareChallengeWidget_Import(t *testing.T) {

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_challenge_widget.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testChallengeWidgetImport(rnd, accountID, "invisible", "0x4AAF00AAAABn0R22HWm-YUc", "0x4AAF00AAAABn0R22HWm098HVBjhdsYUc"),
			},
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}
