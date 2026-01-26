package zero_trust_organization_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/tidwall/gjson"

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
 *    ✓ V4→V5 migration flow (import with v4, run migration, verify with v5)
 *    ✓ Import functionality - validates users can import post-migration
 *    ✓ Account-scoped imports
 *    ✓ Basic structure validation
 *    ✓ Resource rename verification
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

// TestAccPreCheck_ZeroTrustOrganization verifies a Zero Trust organization exists for the account
// and can be read by both the v6 API (used by v5 provider) and v1 API (used by v4 provider).
// This is required because zero_trust_organization is a singleton resource that must be imported,
// not created. The organization must exist before the test can run.
func TestAccPreCheck_ZeroTrustOrganization(t *testing.T) {
	t.Helper()

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Fatal("CLOUDFLARE_ACCOUNT_ID must be set for this acceptance test")
	}

	// Test 1: Check with v6 API (used by v5 provider)
	client := cloudflare.NewClient(
		option.WithAPIKey(os.Getenv("CLOUDFLARE_API_KEY")),
		option.WithAPIEmail(os.Getenv("CLOUDFLARE_EMAIL")),
		option.WithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN")),
	)

	ctx := context.Background()
	org, err := client.ZeroTrust.Organizations.List(ctx, zero_trust.OrganizationListParams{
		AccountID: cloudflare.F(accountID),
	})

	if err != nil {
		t.Skipf("Skipping test: Failed to list Zero Trust organization with v6 API for account %s. Error: %v", accountID, err)
	}

	if org == nil {
		t.Skipf("Skipping test: Zero Trust organization response is nil for account %s", accountID)
	}

	t.Logf("✓ Zero Trust organization found with v6 API for account %s", accountID)

	// Test 2: Check with v1 API (used by v4 provider) using the Access Organization endpoint
	// This is critical because v4 uses cloudflare-go v1 which calls the Access Organization API
	v1Client, err := cfv1.New(os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_EMAIL"))
	if err != nil {
		// If API key/email not available, try with token
		v1Client, err = cfv1.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
		if err != nil {
			t.Skipf("Skipping test: Failed to create v1 API client: %v", err)
		}
	}

	// Try to get the Access Organization using the v1 API (this is what v4 import does)
	v1Org, _, err := v1Client.GetAccessOrganization(ctx, cfv1.AccountIdentifier(accountID), cfv1.GetAccessOrganizationParams{})
	if err != nil {
		t.Skipf("Skipping test: v4 provider cannot read Access Organization for account %s. Error: %v. This might mean the organization was created through Zero Trust UI instead of Access UI, making it incompatible with v4 import.", accountID, err)
	}

	if v1Org.Name == "" {
		t.Skipf("Skipping test: Access Organization exists but has empty name for account %s", accountID)
	}

	t.Logf("✓ Access Organization readable with v1 API (v4 compatible) for account %s: Name=%s", accountID, v1Org.Name)
}

// TestMigrateZeroTrustOrganization_V4ToV5 tests the complete v4→v5 migration flow:
// 1. Import organization with v4 provider (cloudflare_access_organization)
// 2. Run migration tool to transform config and state
// 3. Verify resource works with v5 provider (cloudflare_zero_trust_organization)
//
// This tests the actual migration workflow users will follow.
func TestMigrateZeroTrustOrganization_V4ToV5(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	v4ResourceName := "cloudflare_access_organization." + rnd
	v5ResourceName := "cloudflare_zero_trust_organization." + rnd
	tmpDir := t.TempDir()

	// V4 config using deprecated resource name - minimal config for import
	// Must include terraform block to specify provider source
	// Must include provider block (even if empty) for v4 to recognize environment variables
	v4Config := fmt.Sprintf(`
terraform {
  required_providers {
    cloudflare = {
      source = "cloudflare/cloudflare"
      version = "4.52.1"
    }
  }
}

provider "cloudflare" {
  # Credentials will be read from environment variables:
  # CLOUDFLARE_API_TOKEN or CLOUDFLARE_API_KEY + CLOUDFLARE_EMAIL
}

resource "cloudflare_access_organization" "%s" {
  account_id = "%s"
}
`, rnd, accountID)

	// Step 0: Import with v4 provider before test steps begin
	acctest.ImportResourceWithV4Provider(t, v4Config, tmpDir, "4.52.1", v4ResourceName, accountID)

	// Read the imported state to get the actual values returned from the API
	// This is necessary because import-only resources bring in all fields from the API,
	// and we need to include them in the config to prevent Terraform from trying to null them out
	stateFile := filepath.Join(tmpDir, "terraform.tfstate")
	stateBytes, err := os.ReadFile(stateFile)
	if err != nil {
		t.Fatalf("Failed to read state file: %v", err)
	}

	// Extract the imported values from state
	// We need to include all fields that have non-default values to prevent Terraform
	// from trying to apply defaults during the migration test
	attrs := gjson.Get(string(stateBytes), "resources.0.instances.0.attributes")

	// Required fields (always present)
	importedName := attrs.Get("name").String()
	importedAuthDomain := attrs.Get("auth_domain").String()

	// Optional string fields
	importedSessionDuration := attrs.Get("session_duration").String()
	importedUIReadOnlyReason := attrs.Get("ui_read_only_toggle_reason").String()
	importedUserSeatExpiration := attrs.Get("user_seat_expiration_inactive_time").String()
	importedWARPAuthDuration := attrs.Get("warp_auth_session_duration").String()

	// Boolean fields (may have non-default values)
	importedAllowAuthViaWarp := attrs.Get("allow_authenticate_via_warp").Bool()
	importedAutoRedirect := attrs.Get("auto_redirect_to_identity").Bool()
	importedIsUIReadOnly := attrs.Get("is_ui_read_only").Bool()

	// Nested objects (check if they exist in state)
	// In v4 state, these are arrays with MaxItems:1, so access with .0
	loginDesign := attrs.Get("login_design.0")
	customPages := attrs.Get("custom_pages.0")

	// Clean up v4 provider cache and lockfile so the test framework can inject v5 provider
	// Keep the tfstate file which contains the imported resource
	terraformDir := filepath.Join(tmpDir, ".terraform")
	if err := os.RemoveAll(terraformDir); err != nil && !os.IsNotExist(err) {
		t.Fatalf("Failed to remove .terraform directory: %v", err)
	}
	lockFile := filepath.Join(tmpDir, ".terraform.lock.hcl")
	if err := os.Remove(lockFile); err != nil && !os.IsNotExist(err) {
		t.Fatalf("Failed to remove .terraform.lock.hcl: %v", err)
	}

	// Update config to remove version constraint and include imported values
	// so test framework can inject v5 provider and config matches state
	configPath := filepath.Join(tmpDir, "test_migration.tf")

	// Build config dynamically - include all fields from imported state to prevent
	// Terraform from trying to apply defaults during migration test
	configTemplate := `
provider "cloudflare" {
  # Credentials will be read from environment variables:
  # CLOUDFLARE_API_TOKEN or CLOUDFLARE_API_KEY + CLOUDFLARE_EMAIL
}

resource "cloudflare_access_organization" "%s" {
  account_id  = "%s"
  name        = "%s"
  auth_domain = "%s"%s
}
`
	var additionalLines string

	// Include optional string fields if present
	if importedSessionDuration != "" {
		additionalLines += fmt.Sprintf("\n  session_duration = \"%s\"", importedSessionDuration)
	}
	if importedUIReadOnlyReason != "" {
		additionalLines += fmt.Sprintf("\n  ui_read_only_toggle_reason = \"%s\"", importedUIReadOnlyReason)
	}
	if importedUserSeatExpiration != "" {
		additionalLines += fmt.Sprintf("\n  user_seat_expiration_inactive_time = \"%s\"", importedUserSeatExpiration)
	}
	if importedWARPAuthDuration != "" {
		additionalLines += fmt.Sprintf("\n  warp_auth_session_duration = \"%s\"", importedWARPAuthDuration)
	}

	// Include boolean fields (always include to match state)
	additionalLines += fmt.Sprintf("\n  allow_authenticate_via_warp = %t", importedAllowAuthViaWarp)
	additionalLines += fmt.Sprintf("\n  auto_redirect_to_identity = %t", importedAutoRedirect)
	additionalLines += fmt.Sprintf("\n  is_ui_read_only = %t", importedIsUIReadOnly)

	// Include nested objects if present in state
	// In v4 config, these use block syntax, but we need to preserve them
	if loginDesign.Exists() {
		additionalLines += "\n  login_design {"
		if bg := loginDesign.Get("background_color").String(); bg != "" {
			additionalLines += fmt.Sprintf("\n    background_color = %q", bg)
		}
		if tc := loginDesign.Get("text_color").String(); tc != "" {
			additionalLines += fmt.Sprintf("\n    text_color = %q", tc)
		}
		if lp := loginDesign.Get("logo_path").String(); lp != "" {
			additionalLines += fmt.Sprintf("\n    logo_path = %q", lp)
		}
		if ht := loginDesign.Get("header_text").String(); ht != "" {
			additionalLines += fmt.Sprintf("\n    header_text = %q", ht)
		}
		if ft := loginDesign.Get("footer_text").String(); ft != "" {
			additionalLines += fmt.Sprintf("\n    footer_text = %q", ft)
		}
		additionalLines += "\n  }"
	}

	if customPages.Exists() {
		additionalLines += "\n  custom_pages {"
		if f := customPages.Get("forbidden").String(); f != "" {
			additionalLines += fmt.Sprintf("\n    forbidden = %q", f)
		}
		if id := customPages.Get("identity_denied").String(); id != "" {
			additionalLines += fmt.Sprintf("\n    identity_denied = %q", id)
		}
		additionalLines += "\n  }"
	}

	v4ConfigWithoutVersion := fmt.Sprintf(configTemplate, rnd, accountID, importedName, importedAuthDomain, additionalLines)
	if err := os.WriteFile(configPath, []byte(v4ConfigWithoutVersion), 0644); err != nil {
		t.Fatalf("Failed to update config: %v", err)
	}
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			TestAccPreCheck_ZeroTrustOrganization(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			// Step 1: Run migration and verify with v5
			acctest.MigrationV2TestStep(t, v4ConfigWithoutVersion, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Verify resource was renamed from v4 to v5 name
				statecheck.ExpectKnownValue(v5ResourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				// Verify core fields exist (we can't check exact values since we imported an unknown org)
				statecheck.ExpectKnownValue(v5ResourceName, tfjsonpath.New("name"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(v5ResourceName, tfjsonpath.New("auth_domain"), knownvalue.NotNull()),
			}),
		},
	})
}

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
			TestAccPreCheck_ZeroTrustOrganization(t)
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
			TestAccPreCheck_ZeroTrustOrganization(t)
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
