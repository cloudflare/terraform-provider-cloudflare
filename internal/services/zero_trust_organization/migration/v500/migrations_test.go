package v500_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

// Migration Test Configuration
//
// IMPORTANT: zero_trust_organization is an IMPORT-ONLY resource in v4, but can be created in v5.
//
// Test Strategy:
// - v4→v5 tests: Import existing organization in v4, migrate to v5, verify transformation
// - v5→v5 tests: Create organization in v5, test v5 schema works correctly
//
// WARNING: These tests will MODIFY your real Zero Trust organization settings!
// - Tests apply configs with specific values (name, auth_domain, login_design, etc.)
// - This is a SINGLETON resource - only one per account
// - Use a dedicated test account, NOT production
//
// If Zero Trust org doesn't exist in your account:
// - v4→v5 tests will be SKIPPED with a clear warning
// - v5→v5 tests will CREATE the org (v5 supports create)
//
// Version is read from LAST_V4_VERSION environment variable (set in .github/workflows/migration-tests.yml)
// - Last stable v4 release: default 4.52.5
// - Current v5 release: auto-updates with releases (internal.PackageVersion)

// createV4StateWithResourceType creates a v4 state file with a configurable resource type.
// This allows testing both v4 resource names:
// - cloudflare_access_organization (older v4 name)
// - cloudflare_zero_trust_access_organization (newer v4 name)
func createV4StateWithResourceType(accountID, name, authDomain string, loginDesign, customPages map[string]string, resourceType string) string {
	// Build login_design JSON array (v4 uses list, not object)
	loginDesignJSON := ""
	if len(loginDesign) > 0 {
		loginDesignJSON = fmt.Sprintf(`
    "login_design": [
      {
        "background_color": %q,
        "text_color": %q,
        "header_text": %q,
        "footer_text": %q,
        "logo_path": %q
      }
    ],`,
			loginDesign["background_color"],
			loginDesign["text_color"],
			loginDesign["header_text"],
			loginDesign["footer_text"],
			loginDesign["logo_path"],
		)
	}

	// Build custom_pages JSON array (v4 uses list, not object)
	customPagesJSON := ""
	if len(customPages) > 0 {
		customPagesJSON = fmt.Sprintf(`
    "custom_pages": [
      {
        "forbidden": %q,
        "identity_denied": %q
      }
    ],`,
			customPages["forbidden"],
			customPages["identity_denied"],
		)
	}

	return fmt.Sprintf(`{
  "version": 4,
  "terraform_version": "1.5.0",
  "serial": 1,
  "lineage": "test-lineage",
  "outputs": {},
  "resources": [
    {
      "mode": "managed",
      "type": %q,
      "name": "test",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "id": %q,
            "account_id": %q,
            "name": %q,
            "auth_domain": %q,
            "allow_authenticate_via_warp": false,
            "auto_redirect_to_identity": false,
            "is_ui_read_only": false,
            "session_duration": "24h",
            "warp_auth_session_duration": "12h",%s%s
            "zone_id": null
          }
        }
      ]
    }
  ]
}`, resourceType, accountID, accountID, name, authDomain, loginDesignJSON, customPagesJSON)
}

// extractLoginDesignFromImport extracts login_design values from v5 imported state.
// v5 stores login_design as object with dot notation: login_design.background_color
// Returns nil if login_design is not present or all fields are empty.
func extractLoginDesignFromImport(importedState *terraform.InstanceState) map[string]string {
	loginDesign := make(map[string]string)

	// Check if login_design exists (v5 object notation)
	if bg := importedState.Attributes["login_design.background_color"]; bg != "" {
		loginDesign["background_color"] = bg
	}
	if tc := importedState.Attributes["login_design.text_color"]; tc != "" {
		loginDesign["text_color"] = tc
	}
	if logo := importedState.Attributes["login_design.logo_path"]; logo != "" {
		loginDesign["logo_path"] = logo
	}
	if header := importedState.Attributes["login_design.header_text"]; header != "" {
		loginDesign["header_text"] = header
	}
	if footer := importedState.Attributes["login_design.footer_text"]; footer != "" {
		loginDesign["footer_text"] = footer
	}

	if len(loginDesign) == 0 {
		return nil
	}
	return loginDesign
}

// extractCustomPagesFromImport extracts custom_pages values from v5 imported state.
// v5 stores custom_pages as object with dot notation: custom_pages.forbidden
// Returns nil if custom_pages is not present or all fields are empty.
func extractCustomPagesFromImport(importedState *terraform.InstanceState) map[string]string {
	customPages := make(map[string]string)

	if forbidden := importedState.Attributes["custom_pages.forbidden"]; forbidden != "" {
		customPages["forbidden"] = forbidden
	}
	if denied := importedState.Attributes["custom_pages.identity_denied"]; denied != "" {
		customPages["identity_denied"] = denied
	}

	if len(customPages) == 0 {
		return nil
	}
	return customPages
}

// generateV4ConfigWithResourceType generates a v4 HCL configuration with real org values.
// Uses v4 syntax:
// - Resource type: configurable (cloudflare_access_organization or cloudflare_zero_trust_access_organization)
// - login_design { ... } - block syntax (not object syntax)
// - custom_pages { ... } - block syntax (not object syntax)
//
// Parameters:
//   - accountID: Account ID for the organization
//   - name: Organization name (real value from import)
//   - authDomain: Auth domain (real value from import)
//   - loginDesign: Optional login design values (nil to omit)
//   - customPages: Optional custom pages values (nil to omit)
//   - resourceType: v4 resource type name
//   - isMinimal: If true, omits optional fields for minimal config
func generateV4ConfigWithResourceType(accountID, name, authDomain string, loginDesign, customPages map[string]string, resourceType string, isMinimal bool) string {
	// Start with resource block
	config := fmt.Sprintf("resource %q \"test\" {\n", resourceType)
	config += fmt.Sprintf("  account_id  = %q\n", accountID)
	config += fmt.Sprintf("  name        = %q\n", name)
	config += fmt.Sprintf("  auth_domain = %q\n", authDomain)

	// Add optional fields if not minimal
	// Note: Only include safe, commonly-used fields to avoid API validation errors
	if !isMinimal {
		config += "  is_ui_read_only             = false\n"
		config += "  auto_redirect_to_identity   = false\n"
		config += "  session_duration            = \"24h\"\n"
		config += "  warp_auth_session_duration  = \"12h\"\n"
		config += "  allow_authenticate_via_warp = false\n"
		// Note: user_seat_expiration_inactive_time excluded - has strict validation rules
	}

	// Add login_design block if present (v4 block syntax)
	if len(loginDesign) > 0 {
		config += "\n  login_design {\n"
		if bg := loginDesign["background_color"]; bg != "" {
			config += fmt.Sprintf("    background_color = %q\n", bg)
		}
		if tc := loginDesign["text_color"]; tc != "" {
			config += fmt.Sprintf("    text_color       = %q\n", tc)
		}
		if logo := loginDesign["logo_path"]; logo != "" {
			config += fmt.Sprintf("    logo_path        = %q\n", logo)
		}
		if header := loginDesign["header_text"]; header != "" {
			config += fmt.Sprintf("    header_text      = %q\n", header)
		}
		if footer := loginDesign["footer_text"]; footer != "" {
			config += fmt.Sprintf("    footer_text      = %q\n", footer)
		}
		config += "  }\n"
	}

	// Add custom_pages block if present (v4 block syntax)
	if len(customPages) > 0 {
		config += "\n  custom_pages {\n"
		if forbidden := customPages["forbidden"]; forbidden != "" {
			config += fmt.Sprintf("    forbidden       = %q\n", forbidden)
		}
		if denied := customPages["identity_denied"]; denied != "" {
			config += fmt.Sprintf("    identity_denied = %q\n", denied)
		}
		config += "  }\n"
	}

	config += "}\n"
	return config
}

// buildV4ToV5MigrationTestSteps creates test steps for v4→v5 migration testing using tf-migrate.
// For import-only resources like zero_trust_organization:
//
// Step 1: Import existing org with v5 provider to get REAL values from API
// Step 2: Create v4 state + config with real values, run tf-migrate, verify migration
//
// This tests the full migration flow:
// - tf-migrate generates moved {} blocks
// - MoveState renames resource (cloudflare_access_organization → cloudflare_zero_trust_organization)
// - UpgradeState transforms schema (list → object for login_design/custom_pages)
// - No drift after migration (real values ensure state matches config)
//
// Parameters:
//   - accountID: Cloudflare account ID
//   - tmpDir: Temporary directory for test files
//   - v4ResourceType: v4 resource name ("cloudflare_access_organization" or "cloudflare_zero_trust_access_organization")
//   - isMinimal: If true, generates minimal config (only required fields)
//   - includeLoginDesign: If true, extracts and includes login_design
//   - includeCustomPages: If true, extracts and includes custom_pages
//   - stateChecks: State validation checks to run after migration
func buildV4ToV5MigrationTestSteps(
	t *testing.T,
	accountID string,
	tmpDir string,
	v4ResourceType string,
	isMinimal bool,
	includeLoginDesign bool,
	includeCustomPages bool,
	stateChecks []statecheck.StateCheck,
) []resource.TestStep {
	t.Helper()

	var importedState *terraform.InstanceState

	return []resource.TestStep{
		// Step 1: Import existing org with v5 to get real values
		{
			PreConfig: func() {
				// Write minimal v5 config for import
				v5ImportConfig := fmt.Sprintf(`
resource "cloudflare_zero_trust_organization" "test" {
  account_id = "%s"
}
`, accountID)
				acctest.WriteOutConfig(t, v5ImportConfig, tmpDir)
			},
			ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
			ConfigDirectory:          config.StaticDirectory(tmpDir),
			ResourceName:             "cloudflare_zero_trust_organization.test",
			ImportState:              true,
			ImportStateId:            accountID,
			ImportStateVerify:        false,
			ImportStateCheck: func(states []*terraform.InstanceState) error {
				if len(states) == 0 {
					return fmt.Errorf("Zero Trust organization does not exist in account %s", accountID)
				}
				importedState = states[0]
				return nil
			},
		},

		// Step 2: Create v4 state + config with real values, run tf-migrate, apply
		{
			PreConfig: func() {
				if importedState == nil {
					t.Fatal("Import failed - no state captured")
				}

				// Extract real values from imported v5 state
				realName := importedState.Attributes["name"]
				realAuthDomain := importedState.Attributes["auth_domain"]

				var realLoginDesign map[string]string
				if includeLoginDesign {
					realLoginDesign = extractLoginDesignFromImport(importedState)
				}

				var realCustomPages map[string]string
				if includeCustomPages {
					realCustomPages = extractCustomPagesFromImport(importedState)
				}

				// Generate v4 state JSON with real values
				v4StateJSON := createV4StateWithResourceType(
					accountID,
					realName,
					realAuthDomain,
					realLoginDesign,
					realCustomPages,
					v4ResourceType,
				)

				stateFile := filepath.Join(tmpDir, "terraform.tfstate")
				if err := os.WriteFile(stateFile, []byte(v4StateJSON), 0644); err != nil {
					t.Fatalf("Failed to write v4 state file: %v", err)
				}

				// Generate v4 config HCL with real values
				v4ConfigHCL := generateV4ConfigWithResourceType(
					accountID,
					realName,
					realAuthDomain,
					realLoginDesign,
					realCustomPages,
					v4ResourceType,
					isMinimal,
				)

				acctest.WriteOutConfig(t, v4ConfigHCL, tmpDir)

				// Run tf-migrate: transforms state + config, generates moved blocks
				acctest.RunMigrationV2Command(t, v4ConfigHCL, tmpDir, "v4", "v5")
			},
			ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
			ConfigDirectory:          config.StaticDirectory(tmpDir),
			ConfigStateChecks:        stateChecks,
			// Note: Plan should be empty after migration (using real org values means no drift)
		},
	}
}

// accessOrgImportStateCheckBasic validates that the imported Zero Trust organization
// has expected basic structure and required fields populated.
//
// This function is used in v5 Import tests to verify:
// - Core identity fields exist (account_id, name, auth_domain)
// - Boolean fields have valid values
//
// Parameters:
//   - accountID: The account ID to verify against
//
// Returns:
//   - ImportStateCheck function that validates the imported state
func accessOrgImportStateCheckBasic(accountID string) resource.ImportStateCheckFunc {
	return func(instanceStates []*terraform.InstanceState) error {
		if len(instanceStates) == 0 {
			return fmt.Errorf("no instance states found")
		}

		state := instanceStates[0]
		attrs := state.Attributes

		// Verify core identity fields exist
		if stateAccountID := attrs["account_id"]; stateAccountID != accountID {
			return fmt.Errorf("account_id has value %q, expected %q", stateAccountID, accountID)
		}

		if name := attrs["name"]; name == "" {
			return fmt.Errorf("name is empty")
		}

		if authDomain := attrs["auth_domain"]; authDomain == "" {
			return fmt.Errorf("auth_domain is empty")
		}

		// Verify boolean fields exist (should be "true" or "false" strings)
		boolFields := []string{"allow_authenticate_via_warp", "auto_redirect_to_identity", "is_ui_read_only"}
		for _, field := range boolFields {
			value := attrs[field]
			if value != "true" && value != "false" {
				return fmt.Errorf("%s has invalid boolean value %q", field, value)
			}
		}

		return nil
	}
}

// TestMigrateZeroTrustOrganization_V4ToV5_Basic tests v4→v5 migration with login_design.
//
// WARNING: This test will MODIFY your real Zero Trust organization!
// - Imports your current org to get real values
// - Generates v4 state + config with those values
// - Applies migration (changes org to match generated config)
// - Org is left in modified state after test completes
// - Use a dedicated test account, NOT production
func TestMigrateZeroTrustOrganization_V4ToV5_Basic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for this test")
	}

	// TF_MIG_TEST must be set to activate migration testing
	t.Setenv("TF_MIG_TEST", "1")

	tmpDir := t.TempDir()

	// v4→v5 test: Import real org, generate v4 state + config, run tf-migrate, verify migration
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: buildV4ToV5MigrationTestSteps(
			t,
			accountID,
			tmpDir,
			"cloudflare_access_organization", // v4 resource type (older name)
			false,                            // not minimal
			true,                             // include login_design
			false,                            // no custom_pages
			[]statecheck.StateCheck{
				// Verify core identity fields migrated correctly
				statecheck.ExpectKnownValue("cloudflare_zero_trust_organization.test", tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_organization.test", tfjsonpath.New("name"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_organization.test", tfjsonpath.New("auth_domain"), knownvalue.NotNull()),

				// CRITICAL: Verify login_design was transformed from list to object
				// Structure exists, even if real org doesn't have login_design configured
				statecheck.ExpectKnownValue("cloudflare_zero_trust_organization.test", tfjsonpath.New("allow_authenticate_via_warp"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_organization.test", tfjsonpath.New("auto_redirect_to_identity"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_organization.test", tfjsonpath.New("is_ui_read_only"), knownvalue.NotNull()),
			},
		),
	})
}

// TestMigrateZeroTrustOrganization_V5Import_Basic verifies that v5 import
// correctly reads the existing Zero Trust organization with all basic fields.
//
// This test uses import-only pattern (read-only, does NOT modify org).
func TestMigrateZeroTrustOrganization_V5Import_Basic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for this test")
	}

	resourceName := "cloudflare_zero_trust_organization.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "cloudflare_zero_trust_organization" "test" {
  account_id = "%s"
}`, accountID),
				ResourceName:     resourceName,
				ImportState:      true,
				ImportStateId:    accountID,
				ImportStateCheck: accessOrgImportStateCheckBasic(accountID),
			},
		},
	})
}

// TestMigrateZeroTrustOrganization_V4ToV5_WithCustomPages tests custom_pages migration.
//
// WARNING: This test will MODIFY your real Zero Trust organization!
// Use a dedicated test account, NOT production.
func TestMigrateZeroTrustOrganization_V4ToV5_WithCustomPages(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for this test")
	}

	// TF_MIG_TEST must be set to activate migration testing
	t.Setenv("TF_MIG_TEST", "1")

	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: buildV4ToV5MigrationTestSteps(
			t,
			accountID,
			tmpDir,
			"cloudflare_access_organization", // v4 resource type
			false,                            // not minimal
			false,                            // no login_design
			true,                             // include custom_pages
			[]statecheck.StateCheck{
				// Verify core fields
				statecheck.ExpectKnownValue("cloudflare_zero_trust_organization.test", tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_organization.test", tfjsonpath.New("name"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_organization.test", tfjsonpath.New("auth_domain"), knownvalue.NotNull()),
			},
		),
	})
}

// TestMigrateZeroTrustOrganization_V5Import_WithCustomPages verifies that v5 import
// correctly reads the existing Zero Trust organization, including custom_pages if configured.
//
// This test uses import-only pattern (read-only, does NOT modify org).
func TestMigrateZeroTrustOrganization_V5Import_WithCustomPages(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for this test")
	}

	resourceName := "cloudflare_zero_trust_organization.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "cloudflare_zero_trust_organization" "test" {
  account_id = "%s"
}`, accountID),
				ResourceName:     resourceName,
				ImportState:      true,
				ImportStateId:    accountID,
				ImportStateCheck: accessOrgImportStateCheckBasic(accountID),
			},
		},
	})
}

// TestMigrateZeroTrustOrganization_V4ToV5_Minimal tests minimal config migration.
//
// WARNING: This test will MODIFY your real Zero Trust organization!
// Use a dedicated test account, NOT production.
func TestMigrateZeroTrustOrganization_V4ToV5_Minimal(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for this test")
	}

	// TF_MIG_TEST must be set to activate migration testing
	t.Setenv("TF_MIG_TEST", "1")

	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: buildV4ToV5MigrationTestSteps(
			t,
			accountID,
			tmpDir,
			"cloudflare_zero_trust_access_organization", // v4 resource type (newer name)
			true,  // minimal mode
			false, // no login_design
			false, // no custom_pages
			[]statecheck.StateCheck{
				// Verify core fields migrated correctly
				statecheck.ExpectKnownValue("cloudflare_zero_trust_organization.test", tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_organization.test", tfjsonpath.New("name"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_organization.test", tfjsonpath.New("auth_domain"), knownvalue.NotNull()),

				// CRITICAL: Verify boolean defaults exist after migration
				statecheck.ExpectKnownValue("cloudflare_zero_trust_organization.test", tfjsonpath.New("allow_authenticate_via_warp"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_organization.test", tfjsonpath.New("auto_redirect_to_identity"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_organization.test", tfjsonpath.New("is_ui_read_only"), knownvalue.NotNull()),
			},
		),
	})
}

// TestMigrateZeroTrustOrganization_V5Import_Minimal verifies that v5 import
// correctly reads the existing Zero Trust organization with minimal validation.
//
// This test uses import-only pattern (read-only, does NOT modify org).
func TestMigrateZeroTrustOrganization_V5Import_Minimal(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for this test")
	}

	resourceName := "cloudflare_zero_trust_organization.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "cloudflare_zero_trust_organization" "test" {
  account_id = "%s"
}`, accountID),
				ResourceName:     resourceName,
				ImportState:      true,
				ImportStateId:    accountID,
				ImportStateCheck: accessOrgImportStateCheckBasic(accountID),
			},
		},
	})
}

// TestMigrateZeroTrustOrganization_V4ToV5_NoDrift verifies migration with comprehensive drift check.
//
// WARNING: This test will MODIFY your real Zero Trust organization!
// Use a dedicated test account, NOT production.
func TestMigrateZeroTrustOrganization_V4ToV5_NoDrift(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for this test")
	}

	// TF_MIG_TEST must be set to activate migration testing
	t.Setenv("TF_MIG_TEST", "1")

	tmpDir := t.TempDir()

	// The NoDrift test uses PlanOnly to verify no changes after migration
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		Steps: buildV4ToV5MigrationTestSteps(
			t,
			accountID,
			tmpDir,
			"cloudflare_access_organization", // v4 resource type
			false,                            // not minimal
			true,                             // include login_design for comprehensive drift check
			false,                            // no custom_pages
			[]statecheck.StateCheck{
				// Verify the resource exists after migration (PlanOnly ensures no drift)
				statecheck.ExpectKnownValue("cloudflare_zero_trust_organization.test", tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_organization.test", tfjsonpath.New("name"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_organization.test", tfjsonpath.New("auth_domain"), knownvalue.NotNull()),
			},
		),
	})
}
