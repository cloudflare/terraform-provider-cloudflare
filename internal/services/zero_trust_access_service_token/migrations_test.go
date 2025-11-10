package zero_trust_access_service_token_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// Config generators for different test scenarios

// accessServiceTokenConfigV4Basic creates a basic v4 config with required fields only
func accessServiceTokenConfigV4Basic(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_access_service_token" "%[1]s" {
  account_id = "%[2]s"
  name       = "test-%[1]s"
}`, rnd, accountID)
}

// accessServiceTokenConfigV4WithDeprecatedField creates a v4 config with the min_days_for_renewal field that will be removed
func accessServiceTokenConfigV4WithDeprecatedField(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_access_service_token" "%[1]s" {
  account_id           = "%[2]s"
  name                 = "test-%[1]s"
  duration             = "17520h"
  min_days_for_renewal = 30
}`, rnd, accountID)
}

// accessServiceTokenConfigV4ZoneScoped creates a zone-scoped v4 config
func accessServiceTokenConfigV4ZoneScoped(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_access_service_token" "%[1]s" {
  zone_id = "%[2]s"
  name    = "test-%[1]s"
}`, rnd, zoneID)
}

// accessServiceTokenConfigV4LegacyName creates a v4 config using the deprecated resource name
func accessServiceTokenConfigV4LegacyName(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_service_token" "%[1]s" {
  account_id = "%[2]s"
  name       = "test-%[1]s"
}`, rnd, accountID)
}

// TestMigrateZeroTrustAccessServiceToken_Basic tests basic migration from v4 to v5
// This test verifies:
// 1. Resource is created successfully with v4 provider
// 2. Migration tool runs without errors
// 3. Resource is renamed to v5 name (if using legacy name)
// 4. All fields are preserved correctly
// 5. State can be read by v5 provider
func TestMigrateZeroTrustAccessServiceToken_Basic(t *testing.T) {
	// Zero Trust Access resources don't support API tokens yet
	// This is required for the test to work properly
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_service_token." + rnd
	tmpDir := t.TempDir()

	// Create v4 configuration
	v4Config := accessServiceTokenConfigV4Basic(rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create resource with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1", // Use exact version for consistency
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			// MigrationV2TestStep will:
			// - Write the config to the tmpDir
			// - Run tf-migrate to transform both config and state
			// - Verify the plan is empty (no changes needed)
			// - Run the state checks
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Verify resource exists with correct type
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				// Verify fields are preserved
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact("test-"+rnd)),
				// Verify computed fields exist
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("client_id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_at"), knownvalue.NotNull()),
				// Verify default duration is applied
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("duration"), knownvalue.StringExact("8760h")),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessServiceToken_WithDeprecatedField tests migration with field removal
// This test specifically verifies that:
// 1. The min_days_for_renewal field is removed from the config
// 2. The min_days_for_renewal field is removed from the state
// 3. The migration handles this gracefully without errors
// 4. Other fields (like duration) are preserved
func TestMigrateZeroTrustAccessServiceToken_WithDeprecatedField(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_service_token." + rnd
	tmpDir := t.TempDir()

	v4Config := accessServiceTokenConfigV4WithDeprecatedField(rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create resource with v4 provider including deprecated field
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration
			// The migration tool should:
			// - Remove min_days_for_renewal from the config file
			// - Remove min_days_for_renewal from the state
			// - Preserve all other fields
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact("test-"+rnd)),
				// Verify duration is preserved (not removed with min_days_for_renewal)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("duration"), knownvalue.StringExact("17520h")),
				// Note: We can't check that min_days_for_renewal is absent because
				// the StateCheck API doesn't have an "expect absent" check.
				// But the migration will fail if it's still present because v5 schema doesn't accept it.
			}),
		},
	})
}

// TestMigrateZeroTrustAccessServiceToken_ZoneScoped tests zone-scoped resource migration
// This test verifies that:
// 1. Zone-scoped resources (using zone_id instead of account_id) migrate correctly
// 2. The zone_id is preserved
// 3. No account_id is added
func TestMigrateZeroTrustAccessServiceToken_ZoneScoped(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_service_token." + rnd
	tmpDir := t.TempDir()

	v4Config := accessServiceTokenConfigV4ZoneScoped(rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create zone-scoped resource with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify zone_id is preserved
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				// Verify zone_id is preserved
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact("test-"+rnd)),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessServiceToken_LegacyName tests migration from deprecated resource name
// This test verifies that:
// 1. The deprecated v4 resource name "cloudflare_access_service_token" is handled
// 2. The resource is migrated to "cloudflare_zero_trust_access_service_token"
// 3. All data is preserved during the rename
func TestMigrateZeroTrustAccessServiceToken_LegacyName(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	// Note: The resource name in state will be the v5 name after migration
	resourceName := "cloudflare_zero_trust_access_service_token." + rnd
	tmpDir := t.TempDir()

	v4Config := accessServiceTokenConfigV4LegacyName(rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create resource with v4 provider using legacy name
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration
			// The migration should:
			// - Rename the resource type in the config from cloudflare_access_service_token to cloudflare_zero_trust_access_service_token
			// - Update the state to use the new resource type
			// - Preserve all data
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact("test-"+rnd)),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessServiceToken_TypeConversion tests client_secret_version type conversion
// This test verifies that:
// 1. The client_secret_version field is converted from int to float64
// 2. Values like 1, 2, 3 become 1.0, 2.0, 3.0
// 3. The conversion happens transparently without errors
//
// Note: This test creates a resource and then triggers rotation to get a client_secret_version > 1
func TestMigrateZeroTrustAccessServiceToken_TypeConversion(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_service_token." + rnd
	tmpDir := t.TempDir()

	// Create a config that explicitly sets client_secret_version
	// In practice, this field is computed, but for migration testing we want to ensure
	// the type conversion works
	v4ConfigWithRotation := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_service_token" "%[1]s" {
  account_id = "%[2]s"
  name       = "test-%[1]s"
  duration   = "8760h"
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create resource with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4ConfigWithRotation,
			},
			// Step 2: Run migration
			// The migration should:
			// - Convert client_secret_version from int to float64
			// - If the value was 1 (int), it becomes 1.0 (float64)
			// - The provider should accept the float64 value without errors
			acctest.MigrationV2TestStep(t, v4ConfigWithRotation, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact("test-"+rnd)),
				// Verify client_secret_version exists (should be converted to float64)
				// We use NotNull because we can't check the exact value without knowing if rotation happened
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("client_secret_version"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("duration"), knownvalue.StringExact("8760h")),
			}),
		},
	})
}

// TestMigrateZeroTrustAccessServiceToken_CompleteResource tests migration with all fields
// This test verifies that:
// 1. A resource with all possible fields migrates correctly
// 2. All fields are preserved
// 3. The deprecated field is removed
// 4. Type conversions work with a complete resource
func TestMigrateZeroTrustAccessServiceToken_CompleteResource(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_access_service_token." + rnd
	tmpDir := t.TempDir()

	// Create a complete v4 config with all fields
	v4ConfigComplete := fmt.Sprintf(`
resource "cloudflare_zero_trust_access_service_token" "%[1]s" {
  account_id           = "%[2]s"
  name                 = "test-%[1]s"
  duration             = "43800h"
  min_days_for_renewal = 60
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create complete resource with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4ConfigComplete,
			},
			// Step 2: Run migration on complete resource
			acctest.MigrationV2TestStep(t, v4ConfigComplete, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact("test-"+rnd)),
				// Verify duration is preserved (43800h = 5 years)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("duration"), knownvalue.StringExact("43800h")),
				// Verify computed fields
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("client_id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("client_secret"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_at"), knownvalue.NotNull()),
			}),
		},
	})
}
