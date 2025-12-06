package zero_trust_dlp_predefined_profile_test

import (
	"context"
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
	resource.AddTestSweepers("cloudflare_zero_trust_dlp_predefined_profile", &resource.Sweeper{
		Name: "cloudflare_zero_trust_dlp_predefined_profile",
		F:    testSweepCloudflareZeroTrustDLPPredefinedProfile,
	})
}

func testSweepCloudflareZeroTrustDLPPredefinedProfile(r string) error {
	ctx := context.Background()
	// DLP Predefined Profile is an account-level predefined profile configuration.
	// It's a singleton setting per account, not something that accumulates.
	// No sweeping required.
	tflog.Info(ctx, "Zero Trust DLP Predefined Profile doesn't require sweeping (predefined configuration)")
	return nil
}

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
				// entries array will be empty on the response. enabled entries is preferred
				ExpectNonEmptyPlan: true,
			},
			{
				// Step 2: Update the resource, changing 'enabled' to false
				Config: testAccZeroTrustDlpPredefinedProfileConfig(rnd, accountID, "false"),
				// Check function to verify the resource attributes after update
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "ocr_enabled", "false"),
				),
				// entries array will be empty on the response. enabled entries is preferred
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccCloudflareZeroTrustDlpPredefinedProfileEnabledEntries_Basic(t *testing.T) {
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
				Config: testAccZeroTrustDlpPredefinedProfileConfigEnabledEntries(rnd, accountID, "true"),
				// Check function to verify the resource attributes after creation
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "ocr_enabled", "true"),
				),
			},
			{
				// Step 2: Update the resource, changing 'enabled' to false
				Config: testAccZeroTrustDlpPredefinedProfileConfigEnabledEntries(rnd, accountID, "false"),
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

func testAccZeroTrustDlpPredefinedProfileConfigEnabledEntries(rnd, accountID, enabled string) string {
	return acctest.LoadTestCase("enabled_entries.tf", rnd, accountID, enabled)
}
