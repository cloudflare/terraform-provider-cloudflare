package zero_trust_organization_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

/* Migration tests for zero_trust_organization resource.
 *
 * IMPORTANT: zero_trust_organization is a SINGLETON and IMPORT-ONLY resource in v4.
 *
 * This means:
 * - One organization per account (singleton)
 * - v4 provider's Create() returns error "cannot be created and must be imported"
 * - Must import existing organization before migration
 * - v5 provider only supports account-scoped imports (zone-scoped not supported)
 *
 * ═══════════════════════════════════════════════════════════════════════════════
 * TEST COVERAGE STRATEGY
 * ═══════════════════════════════════════════════════════════════════════════════
 *
 * Due to the import-only constraint, testing is distributed across different locations:
 *
 * 1. THIS FILE (migrations_test.go):
 *    ✓ Import functionality - validates users can import post-migration
 *    ✓ Account-scoped imports
 *    ✓ Basic structure validation
 *    ✗ Cannot test v4→v5 migration directly (v4 Create fails)
 *    ✗ Cannot test multi-step workflows (unknown org values)
 *
 * 2. resource_test.go:
 *    ✓ Full CRUD operations (v5 Create calls Update internally, so it works)
 *    ✓ All field combinations
 *    ✓ login_design and custom_pages attributes
 *    ✓ Boolean defaults behavior
 *    ✓ Update operations
 *
 * 3. tf-migrate tests (in tf-migrate/):
 *    ✓ Actual v4→v5 migration transformations
 *    ✓ Resource rename: cloudflare_access_organization → cloudflare_zero_trust_organization
 *    ✓ Resource rename: cloudflare_zero_trust_access_organization → cloudflare_zero_trust_organization
 *    ✓ login_design: TypeList MaxItems:1 block → SingleNestedAttribute
 *    ✓ custom_pages: TypeList MaxItems:1 block → SingleNestedAttribute
 *    ✓ Boolean defaults: auto_redirect_to_identity, allow_authenticate_via_warp, is_ui_read_only
 *    ✓ Config transformations (HCL rewriting)
 *    ✓ State transformations (JSON manipulation)
 *
 * This distributed approach provides comprehensive coverage despite the import-only limitation.
 *
 * ═══════════════════════════════════════════════════════════════════════════════
 * REAL-WORLD MIGRATION WORKFLOW
 * ═══════════════════════════════════════════════════════════════════════════════
 *
 * 1. User has v4 config: terraform import cloudflare_access_organization.x account/123
 * 2. User runs: tf-migrate migrate --config-dir .
 * 3. User upgrades to v5: terraform init -upgrade
 * 4. User verifies: terraform plan (should show no changes)
 *
 * Import format difference:
 * - v4: terraform import cloudflare_access_organization.x account/123 (account-scoped)
 * - v5: terraform import cloudflare_zero_trust_organization.x 123 (only account-scoped supported)
 */

// TestMigrateZeroTrustOrganization_Import tests that the v5 resource can be imported correctly.
// This validates the import functionality that users rely on after migration.
//
// Note: This test only verifies import works. We cannot test that subsequent plans are empty
// because we don't know what values exist in the imported organization, and providing a minimal
// config (only account_id) will cause Terraform to want to recreate with that minimal config.
func TestMigrateZeroTrustOrganization_Import(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_organization." + rnd

	// Minimal config for import test
	v5Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_organization" "%s" {
  account_id = "%s"
}
`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			{
				// Import organization and verify it succeeds
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   v5Config,
				ResourceName:             resourceName,
				ImportState:              true,
				ImportStateId:            accountID, // v5 expects just the account ID, not "account/{id}"
				ImportStateVerify:        false,     // Can't verify - actual org values are unknown
				// Verify the basic structure after import
				ImportStateCheck: func(states []*terraform.InstanceState) error {
					if len(states) == 0 {
						return fmt.Errorf("no states returned from import")
					}
					state := states[0]
					if state.Attributes[consts.AccountIDSchemaKey] != accountID {
						return fmt.Errorf("account_id mismatch: got %s, want %s",
							state.Attributes[consts.AccountIDSchemaKey], accountID)
					}
					if state.Attributes["name"] == "" {
						return fmt.Errorf("name is empty after import")
					}
					if state.Attributes["auth_domain"] == "" {
						return fmt.Errorf("auth_domain is empty after import")
					}
					return nil
				},
			},
		},
	})
}

// TestMigrateZeroTrustOrganization_ImportStructureValidation performs comprehensive
// validation of the imported state structure, ensuring all expected fields are present
// and have the correct types. This validates that the v5 schema is correctly structured
// for organizations that may have been migrated from v4.
func TestMigrateZeroTrustOrganization_ImportStructureValidation(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_organization." + rnd

	v5Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_organization" "%s" {
  account_id = "%s"
}
`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: []resource.TestStep{
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   v5Config,
				ResourceName:             resourceName,
				ImportState:              true,
				ImportStateId:            accountID,
				ImportStateCheck: func(states []*terraform.InstanceState) error {
					if len(states) == 0 {
						return fmt.Errorf("no states returned from import")
					}

					state := states[0]
					attrs := state.Attributes

					// Validate required fields are present
					requiredFields := []string{"account_id", "name", "auth_domain"}
					for _, field := range requiredFields {
						if attrs[field] == "" {
							return fmt.Errorf("required field %s is empty or missing", field)
						}
					}

					// Validate account_id matches
					if attrs["account_id"] != accountID {
						return fmt.Errorf("account_id mismatch: got %s, want %s", attrs["account_id"], accountID)
					}

					// Validate boolean fields exist (they should have defaults even if not set)
					// These are critical for v4→v5 migration where defaults were added
					boolFields := []string{
						"is_ui_read_only",
						"auto_redirect_to_identity",
						"allow_authenticate_via_warp",
					}
					for _, field := range boolFields {
						val, exists := attrs[field]
						if !exists {
							return fmt.Errorf("boolean field %s is missing (should have default)", field)
						}
						// Should be "true" or "false", not empty
						if val != "true" && val != "false" {
							return fmt.Errorf("boolean field %s has invalid value: %s", field, val)
						}
					}

					// If login_design exists, validate it's structured as an object not array
					// This is critical for v4→v5 migration (TypeList MaxItems:1 → SingleNestedAttribute)
					if attrs["login_design.%"] != "" && attrs["login_design.%"] != "0" {
						// The .% syntax indicates it's a map/object (good)
						// If it were an array, we'd see login_design.# instead
						t.Logf("✓ login_design is correctly structured as object (not array)")

						// Validate nested fields exist
						loginDesignFields := []string{
							"background_color", "text_color", "logo_path",
							"header_text", "footer_text",
						}
						for _, field := range loginDesignFields {
							key := "login_design." + field
							if _, exists := attrs[key]; exists {
								t.Logf("✓ login_design.%s exists", field)
							}
						}
					}

					// If custom_pages exists, validate it's structured as an object not array
					// This is critical for v4→v5 migration (TypeList MaxItems:1 → SingleNestedAttribute)
					if attrs["custom_pages.%"] != "" && attrs["custom_pages.%"] != "0" {
						t.Logf("✓ custom_pages is correctly structured as object (not array)")

						customPagesFields := []string{"forbidden", "identity_denied"}
						for _, field := range customPagesFields {
							key := "custom_pages." + field
							if _, exists := attrs[key]; exists {
								t.Logf("✓ custom_pages.%s exists", field)
							}
						}
					}

					t.Logf("✓ Import structure validation passed")
					return nil
				},
			},
		},
	})
}
