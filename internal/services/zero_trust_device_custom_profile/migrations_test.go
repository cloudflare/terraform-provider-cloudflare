package zero_trust_device_custom_profile_test

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

// TestMigrateZeroTrustDeviceCustomProfile_Basic tests migration of a basic custom profile
// Tests: resource rename, field preservation (name, description, match, precedence), field removal (default, enabled)
func TestMigrateZeroTrustDeviceCustomProfile_Basic(t *testing.T) {
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

	// v4 config with cloudflare_zero_trust_device_profiles + match + precedence (custom profile)
	// Include service_mode_v2 to avoid API 500 errors with minimal configs
	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_device_profiles" "%[1]s" {
  account_id  = "%[2]s"
  name        = "Custom Test Profile"
  description = "Test custom device profile for migration"
  match       = "identity.email == \"test@example.com\""
  precedence  = %[3]d

  allow_mode_switch    = true
  auto_connect         = 0
  captive_portal       = 180
  service_mode_v2_mode = "proxy"
  service_mode_v2_port = 8080
}`, rnd, accountID, precedence)

	stateChecks := []statecheck.StateCheck{
		// Verify resource type changed to custom profile
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),

		// Verify custom profile fields preserved
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("Custom Test Profile")),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("description"), knownvalue.StringExact("Test custom device profile for migration")),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("match"), knownvalue.StringExact("identity.email == \"test@example.com\"")),
		// Migration transforms precedence by adding 900 to avoid conflicts
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("precedence"), knownvalue.Float64Exact(float64(precedence+900))),

		// Verify other fields preserved
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("allow_mode_switch"), knownvalue.Bool(true)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(0)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("captive_portal"), knownvalue.Float64Exact(180)),
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

// TestMigrateZeroTrustDeviceCustomProfile_WithServiceModeV2 tests migration with service_mode_v2 transformation
// Tests: flat fields (service_mode_v2_mode, service_mode_v2_port) → nested object for custom profiles
func TestMigrateZeroTrustDeviceCustomProfile_WithServiceModeV2(t *testing.T) {
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

	// v4 config with service_mode_v2 flat fields
	// Note: When mode is "warp", don't set port (API uses default 0)
	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_device_profiles" "%[1]s" {
  account_id  = "%[2]s"
  name        = "Custom Profile with Service Mode"
  description = "Test custom profile with service_mode_v2 fields"
  match       = "identity.email == \"admin@example.com\""
  precedence  = %[3]d

  allow_mode_switch    = false
  auto_connect         = 15
  captive_portal       = 300
  service_mode_v2_mode = "warp"
}`, rnd, accountID, precedence)

	stateChecks := []statecheck.StateCheck{
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),

		// Verify custom profile fields
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("Custom Profile with Service Mode")),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("match"), knownvalue.StringExact("identity.email == \"admin@example.com\"")),
		// Migration transforms precedence by adding 900 to avoid conflicts
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("precedence"), knownvalue.Float64Exact(float64(precedence+900))),

		// Verify service_mode_v2 nested object created
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("service_mode_v2").AtMapKey("mode"), knownvalue.StringExact("warp")),
		// When mode is "warp", port is not used (should be null or 0)
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("service_mode_v2").AtMapKey("port"), knownvalue.Null()),

		// Verify other fields
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("allow_mode_switch"), knownvalue.Bool(false)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(15)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("captive_portal"), knownvalue.Float64Exact(300)),
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

// TestMigrateZeroTrustDeviceCustomProfile_Maximal tests migration with all optional fields
// Tests: comprehensive field preservation, type conversions for custom profiles
func TestMigrateZeroTrustDeviceCustomProfile_Maximal(t *testing.T) {
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

	// v4 config with all optional fields for custom profile
	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_device_profiles" "%[1]s" {
  account_id  = "%[2]s"
  name        = "Maximal Custom Profile"
  description = "Test custom profile with all optional fields"
  match       = "identity.email == \"maximal@example.com\""
  precedence  = %[3]d

  allow_mode_switch     = false
  allow_updates         = true
  allowed_to_leave      = false
  auto_connect          = 30
  captive_portal        = 600
  disable_auto_fallback = true
  switch_locked         = true
  support_url           = "https://support.custom.cf-tf-test.com"

  service_mode_v2_mode = "proxy"
  service_mode_v2_port = 443
}`, rnd, accountID, precedence)

	stateChecks := []statecheck.StateCheck{
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),

		// Verify custom profile fields
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("Maximal Custom Profile")),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("description"), knownvalue.StringExact("Test custom profile with all optional fields")),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("match"), knownvalue.StringExact("identity.email == \"maximal@example.com\"")),
		// Migration transforms precedence by adding 900 to avoid conflicts
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("precedence"), knownvalue.Float64Exact(float64(precedence+900))),

		// Verify boolean fields
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("allow_mode_switch"), knownvalue.Bool(false)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("allow_updates"), knownvalue.Bool(true)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("allowed_to_leave"), knownvalue.Bool(false)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("disable_auto_fallback"), knownvalue.Bool(true)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("switch_locked"), knownvalue.Bool(true)),

		// Verify numeric fields (converted to Float64)
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(30)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("captive_portal"), knownvalue.Float64Exact(600)),

		// Verify string fields
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("support_url"), knownvalue.StringExact("https://support.custom.cf-tf-test.com")),

		// Verify service_mode_v2 nested object
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("service_mode_v2").AtMapKey("mode"), knownvalue.StringExact("proxy")),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("service_mode_v2").AtMapKey("port"), knownvalue.Float64Exact(443)),
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

// TestMigrateZeroTrustDeviceCustomProfile_OldResourceName tests migration from old resource name
// Tests: cloudflare_device_settings_policy (old name with match+precedence) → cloudflare_zero_trust_device_custom_profile
func TestMigrateZeroTrustDeviceCustomProfile_OldResourceName(t *testing.T) {
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

	// v4 config with OLD resource name (cloudflare_device_settings_policy) + match + precedence
	v4Config := fmt.Sprintf(`
resource "cloudflare_device_settings_policy" "%[1]s" {
  account_id  = "%[2]s"
  name        = "Old Name Custom Profile"
  description = "Test custom profile with old resource name"
  match       = "identity.email == \"legacy@example.com\""
  precedence  = %[3]d

  allow_mode_switch    = true
  auto_connect         = 0
  captive_portal       = 300
  service_mode_v2_mode = "proxy"
  service_mode_v2_port = 8080
}`, rnd, accountID, precedence)

	stateChecks := []statecheck.StateCheck{
		// Verify resource type changed to new custom profile name
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),

		// Verify custom profile fields
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("Old Name Custom Profile")),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("match"), knownvalue.StringExact("identity.email == \"legacy@example.com\"")),
		// Migration transforms precedence by adding 900 to avoid conflicts
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("precedence"), knownvalue.Float64Exact(float64(precedence+900))),

		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("allow_mode_switch"), knownvalue.Bool(true)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("auto_connect"), knownvalue.Float64Exact(0)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("captive_portal"), knownvalue.Float64Exact(300)),
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
				// Step 1: Create with v4 provider using old name
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

// TestMigrateZeroTrustDeviceCustomProfile_HighPrecedence tests migration with high precedence value
// Tests: precedence field handling for custom profiles (higher precedence = higher priority)
func TestMigrateZeroTrustDeviceCustomProfile_HighPrecedence(t *testing.T) {
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

	// v4 config with unique precedence value
	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_device_profiles" "%[1]s" {
  account_id  = "%[2]s"
  name        = "High Priority Profile"
  description = "Custom profile with high precedence"
  match       = "identity.email == \"vip@example.com\""
  precedence  = %[3]d

  allow_mode_switch    = true
  auto_connect         = 5
  captive_portal       = 120
  service_mode_v2_mode = "proxy"
  service_mode_v2_port = 8080
}`, rnd, accountID, precedence)

	stateChecks := []statecheck.StateCheck{
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),

		// Verify precedence value transformed by adding 900 to avoid conflicts
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("precedence"), knownvalue.Float64Exact(float64(precedence+900))),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("name"), knownvalue.StringExact("High Priority Profile")),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_custom_profile."+rnd, tfjsonpath.New("match"), knownvalue.StringExact("identity.email == \"vip@example.com\"")),
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
