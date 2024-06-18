package web_analytics_site_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stainless-sdks/cloudflare-terraform/internal/acctest"
	"github.com/stainless-sdks/cloudflare-terraform/internal/consts"
	"github.com/stainless-sdks/cloudflare-terraform/internal/utils"
)

func TestAccCloudflareWebAnalyticsSite_Create(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_web_analytics_site.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWebAnalyticsSiteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareWebAnalyticsSite(rnd, accountID, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttrSet(name, "site_tag"),
					resource.TestCheckResourceAttr(name, "host", domain),
					resource.TestCheckResourceAttr(name, "auto_install", "false"),
					resource.TestCheckResourceAttrSet(name, "site_token"),
					resource.TestCheckResourceAttrSet(name, "snippet"),
				),
			},
		},
	})
}

func testAccCheckCloudflareWebAnalyticsSiteDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_web_analytics_site" {
			continue
		}

		_, err := client.GetWebAnalyticsSite(context.Background(), cloudflare.AccountIdentifier(rs.Primary.Attributes[consts.AccountIDSchemaKey]), cloudflare.GetWebAnalyticsSiteParams{
			SiteTag: rs.Primary.Attributes["site_tag"],
		})
		if err == nil {
			return fmt.Errorf("web analytics site still exists")
		}
	}

	return nil
}

func testAccCloudflareWebAnalyticsSite(resourceName, accountID, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_web_analytics_site" "%[1]s" {
  account_id    = "%[2]s"
  host          = "%[3]s"
  auto_install  = false
}
`, resourceName, accountID, domain)
}
