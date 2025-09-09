package zero_trust_connectivity_settings_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareZeroTrustConnectivitySettings_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_connectivity_settings." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZeroTrustConnectivitySettingsBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "icmp_proxy_enabled", "true"),
					resource.TestCheckResourceAttr(name, "offramp_warp_enabled", "true"),
				),
			},
		},
	})
}

func testAccCloudflareZeroTrustConnectivitySettingsBasic(name, accountId string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_connectivity_settings" "%s" {
  account_id = "%s"
  icmp_proxy_enabled = true
  offramp_warp_enabled = true
}
`, name, accountId)
}

func TestAccCloudflareZeroTrustConnectivitySettings_Toggle(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_zero_trust_connectivity_settings." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZeroTrustConnectivitySettingsBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "icmp_proxy_enabled", "true"),
					resource.TestCheckResourceAttr(name, "offramp_warp_enabled", "true"),
				),
			},
			{
				Config: testAccCloudflareZeroTrustConnectivitySettingsToggle(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "icmp_proxy_enabled", "false"),
					resource.TestCheckResourceAttr(name, "offramp_warp_enabled", "false"),
				),
			},
		},
	})
}

func testAccCloudflareZeroTrustConnectivitySettingsToggle(name, accountId string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_connectivity_settings" "%s" {
  account_id = "%s"
  icmp_proxy_enabled = false
  offramp_warp_enabled = false
}
`, name, accountId)
}
