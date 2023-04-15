package challenge_widget_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareChallengeWidgetBasic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_challenge_widget." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareChallengeWidgetBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "bot_fight_mode", "false"),
					resource.TestCheckResourceAttr(resourceName, "domains.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "domains.0", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "mode", "invisible"),
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

func testAccCheckCloudflareChallengeWidgetBasic(rnd, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_challenge_widget" "%[1]s" {
    account_id     = "%[2]s"
    name        = "%[1]s"
	bot_fight_mode = false
	domains = [ "example.com" ]
	mode = "invisible"
	region = "world"
  }`, rnd, accountID)
}
