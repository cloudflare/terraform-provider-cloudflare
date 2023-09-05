package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareAccessCustomPage_IdentityDenied(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_access_custom_page.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessCustomPage_CustomHTML(rnd, zoneID, "identity_denied", "<html><body><h1>Access Denied</h1></body></html>"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "type", "identity_denied"),
					resource.TestCheckResourceAttr(resourceName, "custom_html", "<html><body><h1>Access Denied</h1></body></html>"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessCustomPage_Forbidden(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_access_custom_page.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessCustomPage_CustomHTML(rnd, zoneID, "forbidden", "<html><body><h1>Forbidden</h1></body></html>"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "type", "forbidden"),
					resource.TestCheckResourceAttr(resourceName, "custom_html", "<html><body><h1>Forbidden</h1></body></html>"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessCustomPage_CustomHTML(rnd, zoneID, pageType, markup string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_custom_page" "%[1]s" {
	zone_id = "%[2]s"
	name = "%[1]s"
	type = "%[3]s"
	custom_html = "%[4]s"
}
	`, rnd, zoneID, pageType, markup)
}
