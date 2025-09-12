package zero_trust_dlp_predefined_profile_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareZeroTrustDlpPredefinedProfile_Basic(t *testing.T) {
	// Generate a random resource name to avoid conflicts during testing
	rnd := utils.GenerateRandomResourceName()
	// Define the full resource name for checks
	resourceName := "cloudflare_zero_trust_dlp_predefined_profile." + rnd

	// Retrieve Cloudflare Account ID from environment variables
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		// PreCheck ensures necessary environment variables are set
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		// ProtoV6ProviderFactories provides the provider instance for testing
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: Create the resource with 'enabled' set to true
				Config: testAccZeroTrustDlpPredefinedProfileConfig(rnd, accountID, "true"),
				// Check function to verify the resource attributes after creation
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "ocr_enabled", "true"),
				),
			},
			{
				// Step 2: Update the resource, changing 'enabled' to false
				Config: testAccZeroTrustDlpPredefinedProfileConfig(rnd, accountID, "false"),
				// Check function to verify the resource attributes after update
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "ocr_enabled", "false"),
				),
			},
		},
	})
}

func testAccZeroTrustDlpPredefinedProfileConfig(rnd, accountID, enabled string) string {
	return acctest.LoadTestCase("basic.tf", rnd, accountID, enabled)
}
