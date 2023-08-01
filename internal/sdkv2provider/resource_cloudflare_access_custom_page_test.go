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
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessCustomPage_IdentityDenied(rnd),
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
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessCustomPage_Forbidden(rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "type", "forbidden"),
					resource.TestCheckResourceAttr(resourceName, "custom_html", "<html><body><h1>Forbidden</h1></body></html>"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessCustomPage_IdentityDenied(name string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_custom_page" "%[1]s" {
	name = "%[1]s"
	type = "identity_denied"
	custom_html = "<html><body><h1>Access Denied</h1></body></html>"
}
	`, name)
}

func testAccCheckCloudflareAccessCustomPage_Forbidden(name string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_custom_page" "%[1]s" {
	name = "%[1]s"
	type = "forbidden"
	custom_html = "<html><body><h1>Access Denied</h1></body></html>"
}
	`, name)
}
