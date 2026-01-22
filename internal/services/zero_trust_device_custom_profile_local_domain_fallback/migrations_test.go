package zero_trust_device_custom_profile_local_domain_fallback_test

import (
	"fmt"
	"hash/fnv"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// generateUniquePrecedence generates a unique precedence value based on the random resource name
// to avoid conflicts with existing profiles in the account
func generateUniquePrecedence(rnd string) int {
	h := fnv.New32a()
	h.Write([]byte(rnd))
	// Generate a precedence in range 101-999 to avoid conflicts
	// API requires precedence between 1-999 for custom profiles
	return 101 + int(h.Sum32()%899)
}

// TestMigrateZeroTrustDeviceCustomProfileLocalDomainFallback_Basic tests migration of a custom profile fallback domain
// Tests: resource rename, policy_id preservation, Set → List transformation
// This test creates both a custom device profile and a fallback domain referencing it
func TestMigrateZeroTrustDeviceCustomProfileLocalDomainFallback_Basic(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	precedence := generateUniquePrecedence(rnd)
	tmpDir := t.TempDir()

	// v4 config with custom device profile + fallback domain
	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_device_profiles" "%[1]s_profile" {
  account_id  = "%[2]s"
  name        = "Test Custom Profile"
  description = "Custom profile for testing"
  match       = "identity.email == \"test@example.com\""
  precedence  = %[3]d
}

resource "cloudflare_zero_trust_local_fallback_domain" "%[1]s" {
  account_id = "%[2]s"
  policy_id  = cloudflare_zero_trust_device_profiles.%[1]s_profile.id

  domains {
    suffix      = "custom.example.com"
    description = "Custom profile fallback domain"
    dns_server  = ["10.0.0.1"]
  }
}`, rnd, accountID, precedence)

	stateChecks := []statecheck.StateCheck{
		// Verify device profile migrated to custom profile
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd+"_profile", tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd+"_profile", tfjsonpath.New("name"), knownvalue.StringExact("Test Custom Profile")),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd+"_profile", tfjsonpath.New("match"), knownvalue.StringExact("identity.email == \"test@example.com\"")),
		// Migration transforms precedence by adding 900 to avoid conflicts
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd+"_profile", tfjsonpath.New("precedence"), knownvalue.Float64Exact(float64(precedence+900))),

		// Verify fallback domain migrated to custom profile fallback domain
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("policy_id"), knownvalue.NotNull()),

		// Verify domains converted from Set to List
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains"), knownvalue.ListSizeExact(1)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains").AtSliceIndex(0).AtMapKey("suffix"), knownvalue.StringExact("custom.example.com")),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains").AtSliceIndex(0).AtMapKey("description"), knownvalue.StringExact("Custom profile fallback domain")),
	}

	migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, v4Config, tmpDir, "4.40.0", "v4", "v5", stateChecks)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.40.0",
					},
				},
				Config: v4Config,
			},
		}, migrationSteps...),
	})
}

// TestMigrateZeroTrustDeviceCustomProfileLocalDomainFallback_MultipleDomains tests migration with multiple fallback domains
// Tests: Multiple domains in Set → List, preservation of all domain fields
func TestMigrateZeroTrustDeviceCustomProfileLocalDomainFallback_MultipleDomains(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	precedence := generateUniquePrecedence(rnd)
	tmpDir := t.TempDir()

	// v4 config with custom profile + multiple fallback domains
	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_device_profiles" "%[1]s_profile" {
  account_id  = "%[2]s"
  name        = "Multi Domain Profile"
  description = "Profile with multiple domains"
  match       = "identity.email == \"admin@example.com\""
  precedence  = %[3]d
}

resource "cloudflare_zero_trust_local_fallback_domain" "%[1]s" {
  account_id = "%[2]s"
  policy_id  = cloudflare_zero_trust_device_profiles.%[1]s_profile.id

  domains {
    suffix      = "dev.example.com"
    description = "Development environment"
    dns_server  = ["192.168.1.1"]
  }

  domains {
    suffix      = "staging.example.com"
    description = "Staging environment"
    dns_server  = ["192.168.2.1", "192.168.2.2"]
  }

  domains {
    suffix = "prod.example.com"
  }
}`, rnd, accountID, precedence)

	stateChecks := []statecheck.StateCheck{
		// Verify device profile migrated
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd+"_profile", tfjsonpath.New("id"), knownvalue.NotNull()),

		// Verify fallback domain migrated with all 3 domains
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains"), knownvalue.ListSizeExact(3)),

		// Note: We can't reliably test specific indices because Set → List conversion may change order
		// We just verify the list size is correct
	}

	migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, v4Config, tmpDir, "4.40.0", "v4", "v5", stateChecks)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.40.0",
					},
				},
				Config: v4Config,
			},
		}, migrationSteps...),
	})
}

// TestMigrateZeroTrustDeviceCustomProfileLocalDomainFallback_DeprecatedName tests migration from deprecated resource names
// Tests: Both deprecated device profile and deprecated fallback domain names
func TestMigrateZeroTrustDeviceCustomProfileLocalDomainFallback_DeprecatedName(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	precedence := generateUniquePrecedence(rnd)
	tmpDir := t.TempDir()

	// v4 config using DEPRECATED resource names
	v4Config := fmt.Sprintf(`
resource "cloudflare_device_settings_policy" "%[1]s_profile" {
  account_id  = "%[2]s"
  name        = "Deprecated Name Profile"
  description = "Using deprecated resource name"
  match       = "identity.email == \"legacy@example.com\""
  precedence  = %[3]d
}

resource "cloudflare_fallback_domain" "%[1]s" {
  account_id = "%[2]s"
  policy_id  = cloudflare_device_settings_policy.%[1]s_profile.id

  domains {
    suffix      = "deprecated.example.com"
    description = "Using deprecated fallback domain name"
  }
}`, rnd, accountID, precedence)

	stateChecks := []statecheck.StateCheck{
		// Verify device profile migrated to new custom profile name
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd+"_profile", tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd+"_profile", tfjsonpath.New("name"), knownvalue.StringExact("Deprecated Name Profile")),

		// Verify fallback domain migrated to new custom profile fallback domain name
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains"), knownvalue.ListSizeExact(1)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains").AtSliceIndex(0).AtMapKey("suffix"), knownvalue.StringExact("deprecated.example.com")),
	}

	migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, v4Config, tmpDir, "4.52.1", "v4", "v5", stateChecks)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: append([]resource.TestStep{
			{
				// Step 1: Create with v4 provider using deprecated names
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
		}, migrationSteps...),
	})
}
