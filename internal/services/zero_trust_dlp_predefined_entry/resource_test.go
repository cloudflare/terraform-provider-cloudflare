package zero_trust_dlp_predefined_entry_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareZeroTrustDlpPredefinedEntry_Basic(t *testing.T) {
	// Generate a random resource name to avoid conflicts during testing
	rnd := utils.GenerateRandomResourceName()
	// Define the full resource name for checks
	resourceName := "cloudflare_zero_trust_dlp_predefined_entry." + rnd

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	// Define a predefined entry ID (e.g., VISA Card number)
	entryID := "5b1d5035-8c53-4bc9-a151-404eb32b34b4" // VISA Card number predefined entry ID

	resource.Test(t, resource.TestCase{
		// PreCheck ensures necessary environment variables are set
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		// ProtoV6ProviderFactories provides the provider instance for testing
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: Create the resource with 'enabled' set to true
				Config: testAccZeroTrustDlpPredefinedEntryConfig(rnd, accountID, entryID, true),
				// Check function to verify the resource attributes after creation
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "entry_id", entryID),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			{
				// Step 2: Update the resource, changing 'enabled' to false
				Config: testAccZeroTrustDlpPredefinedEntryConfig(rnd, accountID, entryID, false),
				// Check function to verify the resource attributes after update
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "entry_id", entryID),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				),
			},
		},
	})
}

// testAccZeroTrustDlpPredefinedEntryConfig generates the Terraform configuration string.
// It assumes acctest.LoadTestCase loads a template like:
//
//	resource "cloudflare_zero_trust_dlp_predefined_entry" "{{ .Rnd }}" {
//	  account_id = "{{ .AccountID }}"
//	  entry_id   = "{{ .EntryID }}"
//	  enabled    = {{ .Enabled }}
//	}
func testAccZeroTrustDlpPredefinedEntryConfig(rnd, accountID, entryID string, enabled bool) string {
	return acctest.LoadTestCase("basic.tf", rnd, accountID, entryID, enabled)
}
