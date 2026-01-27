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
	"github.com/hashicorp/terraform-plugin-testing/config"
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

// testAccPreCheck_ZeroTrustOrganization verifies a Zero Trust organization exists for the account
// and can be read by both the v6 API (used by v5 provider) and v1 API (used by v4 provider).
// This is required because zero_trust_organization is a singleton resource that must be imported,
// not created. The organization must exist before the test can run.
func testAccPreCheck_ZeroTrustOrganization(t *testing.T) {
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
	// Include terraform block to specify provider SOURCE (but not version - ExternalProviders handles that)
	// This prevents terraform from defaulting to "hashicorp/cloudflare" instead of "cloudflare/cloudflare"
	v4Config := fmt.Sprintf(`
terraform {
  required_providers {
    cloudflare = {
      source = "cloudflare/cloudflare"
    }
  }
}

provider "cloudflare" {}

resource "cloudflare_access_organization" "%s" {
  account_id = "%s"
}
`, rnd, accountID)

	// Helper function to build config with imported values from state
	buildConfigFromState := func() string {
		// Find the state file - test framework creates subdirectories like 001/, 002/, etc.
		var stateFile string
		entries, err := os.ReadDir(tmpDir)
		if err != nil {
			t.Fatalf("Failed to read test directory: %v", err)
		}

		// Look for terraform.tfstate in subdirectories (sorted by name, so 001 comes first)
		for _, entry := range entries {
			if entry.IsDir() {
				candidateState := filepath.Join(tmpDir, entry.Name(), "terraform.tfstate")
				if _, err := os.Stat(candidateState); err == nil {
					stateFile = candidateState
					break
				}
			}
		}

		if stateFile == "" {
			t.Fatalf("Could not find terraform.tfstate in any subdirectory of %s", tmpDir)
		}

		stateBytes, err := os.ReadFile(stateFile)
		if err != nil {
			t.Fatalf("Failed to read state file %s: %v", stateFile, err)
		}

		// Extract the imported values from state
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

		// Build config dynamically - include all fields from imported state to prevent
		// Terraform from trying to apply defaults during migration test
		configTemplate := `
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

		return fmt.Sprintf(configTemplate, rnd, accountID, importedName, importedAuthDomain, additionalLines)
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			testAccPreCheck_ZeroTrustOrganization(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Import with v4 provider using test framework
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config:             v4Config,
				ResourceName:       v4ResourceName,
				ImportState:        true,
				ImportStateId:      accountID,
				ImportStatePersist: true, // Keep imported state for migration step
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
			{
				// Step 2: Build config from imported state, then run migration
				PreConfig: func() {
					// Build config that matches imported state
					v4ConfigWithValues := buildConfigFromState()

					// Write config to the main tmpDir (not subdirectory)
					// Migration tool will look here and also find state in subdirectories
					configPath := filepath.Join(tmpDir, "test_migration.tf")
					if err := os.WriteFile(configPath, []byte(v4ConfigWithValues), 0644); err != nil {
						t.Fatalf("Failed to write config: %v", err)
					}

					// Run migration
					acctest.RunMigrationV2Command(t, v4ConfigWithValues, tmpDir, "v4", "v5")
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir), // Read config from tmpDir
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify resource was renamed from v4 to v5 name
					statecheck.ExpectKnownValue(v5ResourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					// Verify core fields exist
					statecheck.ExpectKnownValue(v5ResourceName, tfjsonpath.New("name"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(v5ResourceName, tfjsonpath.New("auth_domain"), knownvalue.NotNull()),
				},
			},
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
			testAccPreCheck_ZeroTrustOrganization(t)
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
			testAccPreCheck_ZeroTrustOrganization(t)
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

					return nil
				},
			},
		},
	})
}
