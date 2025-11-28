package zero_trust_access_custom_page_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_access_custom_page", &resource.Sweeper{
		Name: "cloudflare_zero_trust_access_custom_page",
		F:    testSweepCloudflareZeroTrustAccessCustomPage,
	})
}

func testSweepCloudflareZeroTrustAccessCustomPage(r string) error {
	ctx := context.Background()
	// Access Custom Pages are account-level custom HTML page configurations.
	// They are managed as configuration and don't accumulate.
	// No sweeping required.
	tflog.Info(ctx, "Zero Trust Access Custom Page doesn't require sweeping (account configuration)")
	return nil
}

func TestAccCloudflareAccessCustomPage_IdentityDenied(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_access_custom_page.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessCustomPage_CustomHTML(rnd, accountID, "identity_denied", "<html><body><h1>Access Denied</h1></body></html>"),
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

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_access_custom_page.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessCustomPage_CustomHTML(rnd, accountID, "forbidden", "<html><body><h1>Forbidden</h1></body></html>"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "type", "forbidden"),
					resource.TestCheckResourceAttr(resourceName, "custom_html", "<html><body><h1>Forbidden</h1></body></html>"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessCustomPage_CustomHTML(rnd, accountID, pageType, markup string) string {
	return acctest.LoadTestCase("customhtml.tf", rnd, accountID, pageType, markup)
}
