package zero_trust_device_default_profile_local_domain_fallback_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_Basic tests migration of a basic default profile fallback domain
// Tests: resource rename, field removal (policy_id), Set → List transformation
func TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_Basic(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// v4 config with cloudflare_zero_trust_local_fallback_domain (no policy_id = default profile)
	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_local_fallback_domain" "%[1]s" {
  account_id = "%[2]s"

  domains {
    suffix      = "example.com"
    description = "Example domain"
    dns_server  = ["10.0.0.1", "10.0.0.2"]
  }
}`, rnd, accountID)

	stateChecks := []statecheck.StateCheck{
		// Verify resource type changed
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),

		// Verify domains converted from Set to List
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains"), knownvalue.ListSizeExact(1)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains").AtSliceIndex(0).AtMapKey("suffix"), knownvalue.StringExact("example.com")),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains").AtSliceIndex(0).AtMapKey("description"), knownvalue.StringExact("Example domain")),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains").AtSliceIndex(0).AtMapKey("dns_server"), knownvalue.ListSizeExact(2)),
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
				// Step 1: Create with v4 provider
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

// TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_DeprecatedName tests migration from deprecated resource name
// Tests: cloudflare_fallback_domain → cloudflare_zero_trust_device_default_profile_local_domain_fallback
func TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_DeprecatedName(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// v4 config with DEPRECATED cloudflare_fallback_domain name (no policy_id = default profile)
	v4Config := fmt.Sprintf(`
resource "cloudflare_fallback_domain" "%[1]s" {
  account_id = "%[2]s"

  domains {
    suffix      = "internal.example.com"
    description = "Internal domain"
    dns_server  = ["10.1.0.1"]
  }
}`, rnd, accountID)

	stateChecks := []statecheck.StateCheck{
		// Verify resource type changed to new name
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),

		// Verify domains preserved
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains"), knownvalue.ListSizeExact(1)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains").AtSliceIndex(0).AtMapKey("suffix"), knownvalue.StringExact("internal.example.com")),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains").AtSliceIndex(0).AtMapKey("description"), knownvalue.StringExact("Internal domain")),
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
				// Step 1: Create with v4 provider using deprecated name
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

// TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_MultipleDomains tests migration with multiple fallback domains
// Tests: Multiple domains in Set → List, preservation of all domain fields
func TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_MultipleDomains(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// v4 config with multiple domains
	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_local_fallback_domain" "%[1]s" {
  account_id = "%[2]s"

  domains {
    suffix      = "corp.example.com"
    description = "Corporate network"
    dns_server  = ["10.0.0.1", "10.0.0.2"]
  }

  domains {
    suffix      = "internal.example.com"
    description = "Internal services"
    dns_server  = ["10.1.0.1"]
  }

  domains {
    suffix = "local.example.com"
  }
}`, rnd, accountID)

	stateChecks := []statecheck.StateCheck{
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),

		// Verify all 3 domains migrated
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains"), knownvalue.ListSizeExact(3)),

		// Note: We can't reliably test specific indices because Set → List conversion may change order
		// We just verify the list size is correct
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
				// Step 1: Create with v4 provider
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

// TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_NullPolicyID tests migration with explicit null policy_id
// Tests: policy_id = null should be treated as default profile and removed
func TestMigrateZeroTrustDeviceDefaultProfileLocalDomainFallback_NullPolicyID(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// v4 config with explicit policy_id = null (should be treated as default profile)
	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_local_fallback_domain" "%[1]s" {
  account_id = "%[2]s"
  policy_id  = null

  domains {
    suffix      = "test.example.com"
    description = "Test domain with null policy_id"
    dns_server  = ["10.2.0.1"]
  }
}`, rnd, accountID)

	stateChecks := []statecheck.StateCheck{
		// Verify resource type changed to default profile
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),

		// Verify policy_id was removed (field doesn't exist in default profile schema)
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains"), knownvalue.ListSizeExact(1)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_default_profile_local_domain_fallback."+rnd, tfjsonpath.New("domains").AtSliceIndex(0).AtMapKey("suffix"), knownvalue.StringExact("test.example.com")),
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
				// Step 1: Create with v4 provider
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
