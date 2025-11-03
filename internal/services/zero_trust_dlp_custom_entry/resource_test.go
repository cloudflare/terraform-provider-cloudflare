package zero_trust_dlp_custom_entry_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareZeroTrustDlpCustomEntry_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_dlp_custom_entry." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with minimal required attributes
				Config: testAccZeroTrustDlpCustomEntryConfig_Basic(rnd, accountID, "true", fmt.Sprintf("%s-test", rnd), "[0-9]{3}-[0-9]{2}-[0-9]{4}"),
				ConfigStateChecks: []statecheck.StateCheck{
					// Required attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s-test", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern").AtMapKey("regex"), knownvalue.StringExact("[0-9]{3}-[0-9]{2}-[0-9]{4}")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern").AtMapKey("validation"), knownvalue.Null()),
					// Computed attributes should be present
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_at"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("updated_at"), knownvalue.NotNull()),
				},
			},
			{
				// Step 2: Update enabled and name
				Config: testAccZeroTrustDlpCustomEntryConfig_Basic(rnd, accountID, "false", fmt.Sprintf("%s-updated", rnd), "[0-9]{3}-[0-9]{2}-[0-9]{4}"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s-updated", rnd))),
				},
			},
			{
				// Step 3: Import test
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"type"}, // type field has known issues with API response
				ImportStateIdFunc:       testAccZeroTrustDlpCustomEntryImportStateIdFunc(resourceName, accountID),
			},
		},
	})
}

func TestAccCloudflareZeroTrustDlpCustomEntry_WithValidation(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_dlp_custom_entry." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with validation field
				Config: testAccZeroTrustDlpCustomEntryConfig_WithValidation(rnd, accountID, fmt.Sprintf("%s-credit", rnd), "luhn"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern").AtMapKey("regex"), knownvalue.StringExact("[0-9]{13,16}")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern").AtMapKey("validation"), knownvalue.StringExact("luhn")),
				},
			},
			{
				// Step 2: Remove validation field (test optional attribute removal)
				Config: testAccZeroTrustDlpCustomEntryConfig_Basic(rnd, accountID, "true", fmt.Sprintf("%s-credit", rnd), "[0-9]{13,16}"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern").AtMapKey("validation"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestAccCloudflareZeroTrustDlpCustomEntry_UpdatePattern(t *testing.T) {
	// Test updating the pattern regex between different formats
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_dlp_custom_entry." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with SSN pattern
				Config: testAccZeroTrustDlpCustomEntryConfig_Basic(rnd, accountID, "true", fmt.Sprintf("%s-ssn", rnd), "[0-9]{3}-[0-9]{2}-[0-9]{4}"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s-ssn", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern").AtMapKey("regex"), knownvalue.StringExact("[0-9]{3}-[0-9]{2}-[0-9]{4}")),
				},
			},
			{
				// Step 2: Update to ZIP code pattern
				Config: testAccZeroTrustDlpCustomEntryConfig_Basic(rnd, accountID, "true", fmt.Sprintf("%s-zip", rnd), "[0-9]{5}(-[0-9]{4})?"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s-zip", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern").AtMapKey("regex"), knownvalue.StringExact("[0-9]{5}(-[0-9]{4})?")),
				},
			},
			{
				// Step 3: Import test
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"type"}, // type field has known issues with API response
				ImportStateIdFunc:       testAccZeroTrustDlpCustomEntryImportStateIdFunc(resourceName, accountID),
			},
		},
	})
}

func TestAccCloudflareZeroTrustDlpCustomEntry_CompleteWorkflow(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_dlp_custom_entry." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with all attributes including validation
				Config: testAccZeroTrustDlpCustomEntryConfig_Complete(rnd, accountID, "true", fmt.Sprintf("%s-complete", rnd), "4[0-9]{12}(?:[0-9]{3})?", "luhn"),
				ConfigStateChecks: []statecheck.StateCheck{
					// All configured attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s-complete", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern").AtMapKey("regex"), knownvalue.StringExact("4[0-9]{12}(?:[0-9]{3})?")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern").AtMapKey("validation"), knownvalue.StringExact("luhn")),
					// Verify profile_id is set
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("profile_id"), knownvalue.NotNull()),
				},
			},
			{
				// Step 2: Update all updatable fields
				Config: testAccZeroTrustDlpCustomEntryConfig_Complete(rnd, accountID, "false", fmt.Sprintf("%s-updated-complete", rnd), "5[0-9]{15}", ""),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s-updated-complete", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern").AtMapKey("regex"), knownvalue.StringExact("5[0-9]{15}")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("pattern").AtMapKey("validation"), knownvalue.Null()),
				},
			},
			{
				// Step 3: Import test
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"type"}, // type field has known issues with API response
				ImportStateIdFunc:       testAccZeroTrustDlpCustomEntryImportStateIdFunc(resourceName, accountID),
			},
		},
	})
}

func testAccZeroTrustDlpCustomEntryConfig_Basic(rnd, accountID, enabled, name, regex string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_dlp_custom_profile" "custom_profile" {
  name       = "%[1]s"
  account_id = "%[2]s"
}

resource "cloudflare_zero_trust_dlp_custom_entry" "%[1]s" {
  name       = "%[4]s"
  account_id = "%[2]s"
  profile_id = cloudflare_zero_trust_dlp_custom_profile.custom_profile.id
  enabled    = %[3]s

  pattern = {
    regex = "%[5]s"
  }
}
`, rnd, accountID, enabled, name, regex)
}

func testAccZeroTrustDlpCustomEntryConfig_WithValidation(rnd, accountID, name, validation string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_dlp_custom_profile" "custom_profile" {
  name       = "%[1]s"
  account_id = "%[2]s"
}

resource "cloudflare_zero_trust_dlp_custom_entry" "%[1]s" {
  name       = "%[3]s"
  account_id = "%[2]s"
  profile_id = cloudflare_zero_trust_dlp_custom_profile.custom_profile.id
  enabled    = true

  pattern = {
    regex      = "[0-9]{13,16}"
    validation = "%[4]s"
  }
}
`, rnd, accountID, name, validation)
}

func testAccZeroTrustDlpCustomEntryConfig_Complete(rnd, accountID, enabled, name, regex, validation string) string {
	validationAttr := ""
	if validation != "" {
		validationAttr = fmt.Sprintf(`
    validation = "%s"`, validation)
	}

	return fmt.Sprintf(`
resource "cloudflare_zero_trust_dlp_custom_profile" "custom_profile" {
  name       = "%[1]s"
  account_id = "%[2]s"
}

resource "cloudflare_zero_trust_dlp_custom_entry" "%[1]s" {
  name       = "%[4]s"
  account_id = "%[2]s"
  profile_id = cloudflare_zero_trust_dlp_custom_profile.custom_profile.id
  enabled    = %[3]s

  pattern = {
    regex = "%[5]s"%[6]s
  }
}
`, rnd, accountID, enabled, name, regex, validationAttr)
}

func testAccZeroTrustDlpCustomEntryImportStateIdFunc(resourceName, accountID string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}

		return fmt.Sprintf("%s/%s", accountID, rs.Primary.ID), nil
	}
}
