package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

// Migration Test Configuration
//
// IMPORTANT: zero_trust_organization is an IMPORT-ONLY resource in v4.
// Tests must use import annotations rather than resource creation.
//
// Version is read from LAST_V4_VERSION environment variable (set in .github/workflows/migration-tests.yml)
// - Last stable v4 release: default 4.52.5
// - Current v5 release: auto-updates with releases (internal.PackageVersion)

// Embed migration test configuration files
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_with_custom_pages.tf
var v4WithCustomPagesConfig string

//go:embed testdata/v5_with_custom_pages.tf
var v5WithCustomPagesConfig string

//go:embed testdata/v4_minimal.tf
var v4MinimalConfig string

//go:embed testdata/v5_minimal.tf
var v5MinimalConfig string

func TestMigrateZeroTrustOrganization_V4ToV5_Basic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for this test")
	}

	// TF_MIG_TEST must be set to activate migration testing
	t.Setenv("TF_MIG_TEST", "1")

	tmpDir := t.TempDir()
	accountIDConfig := fmt.Sprintf(`
variable "account_id" {
  type    = string
  default = "%s"
}
`, accountID)

	resourceName := "cloudflare_zero_trust_organization.test"

	acctest.MigrationV2TestStep(
		t,
		v4BasicConfig+accountIDConfig,
		tmpDir,
		acctest.GetLastV4Version(),
		"v4",
		"v5",
		[]statecheck.StateCheck{
			// Verify core identity fields
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact("Test Organization")),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_domain"), knownvalue.StringExact("test.cloudflareaccess.com")),

			// Verify boolean defaults were added
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_authenticate_via_warp"), knownvalue.Bool(false)),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_redirect_to_identity"), knownvalue.Bool(false)),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_ui_read_only"), knownvalue.Bool(false)),

			// Verify login_design was converted from array to object
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("login_design").AtMapKey("background_color"), knownvalue.StringExact("#000000")),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("login_design").AtMapKey("text_color"), knownvalue.StringExact("#FFFFFF")),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("login_design").AtMapKey("header_text"), knownvalue.StringExact("Welcome")),

			// Verify pass-through fields
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("warp_auth_session_duration"), knownvalue.StringExact("12h")),
		},
	)
}

// TestMigrateZeroTrustOrganization_V5ToV5_Basic verifies that v5 resources
// with login_design object syntax work correctly without migration (v5→v5 pass-through).
func TestMigrateZeroTrustOrganization_V5ToV5_Basic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for this test")
	}

	// TF_MIG_TEST must be set to activate migration testing
	t.Setenv("TF_MIG_TEST", "1")

	tmpDir := t.TempDir()
	accountIDConfig := fmt.Sprintf(`
variable "account_id" {
  type    = string
  default = "%s"
}
`, accountID)

	resourceName := "cloudflare_zero_trust_organization.test"

	// Use v5 config directly to test v5→v5 pass-through (no migration needed)
	acctest.MigrationV2TestStep(
		t,
		v5BasicConfig+accountIDConfig,
		tmpDir,
		acctest.GetLastV4Version(),
		"v5",
		"v5",
		[]statecheck.StateCheck{
			// Verify core identity fields
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact("Test Organization")),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_domain"), knownvalue.StringExact("test.cloudflareaccess.com")),

			// Verify login_design object syntax works in v5
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("login_design").AtMapKey("background_color"), knownvalue.StringExact("#000000")),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("login_design").AtMapKey("text_color"), knownvalue.StringExact("#FFFFFF")),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("login_design").AtMapKey("header_text"), knownvalue.StringExact("Welcome")),

			// Verify pass-through fields
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_duration"), knownvalue.StringExact("24h")),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("warp_auth_session_duration"), knownvalue.StringExact("12h")),
		},
	)
}

func TestMigrateZeroTrustOrganization_V4ToV5_WithCustomPages(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for this test")
	}

	// TF_MIG_TEST must be set to activate migration testing
	t.Setenv("TF_MIG_TEST", "1")

	tmpDir := t.TempDir()
	accountIDConfig := fmt.Sprintf(`
variable "account_id" {
  type    = string
  default = "%s"
}
`, accountID)

	resourceName := "cloudflare_zero_trust_organization.test"

	acctest.MigrationV2TestStep(
		t,
		v4WithCustomPagesConfig+accountIDConfig,
		tmpDir,
		acctest.GetLastV4Version(),
		"v4",
		"v5",
		[]statecheck.StateCheck{
			// Verify custom_pages was converted from array to object
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("custom_pages").AtMapKey("forbidden"), knownvalue.StringExact("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("custom_pages").AtMapKey("identity_denied"), knownvalue.StringExact("yyyyyyyy-yyyy-yyyy-yyyy-yyyyyyyyyyyy")),

			// Verify boolean defaults were added even when not specified
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_authenticate_via_warp"), knownvalue.Bool(false)),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_redirect_to_identity"), knownvalue.Bool(false)),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_ui_read_only"), knownvalue.Bool(false)),
		},
	)
}

// TestMigrateZeroTrustOrganization_V5ToV5_WithCustomPages verifies that v5 resources
// with custom_pages object syntax work correctly without migration (v5→v5 pass-through).
func TestMigrateZeroTrustOrganization_V5ToV5_WithCustomPages(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for this test")
	}

	// TF_MIG_TEST must be set to activate migration testing
	t.Setenv("TF_MIG_TEST", "1")

	tmpDir := t.TempDir()
	accountIDConfig := fmt.Sprintf(`
variable "account_id" {
  type    = string
  default = "%s"
}
`, accountID)

	resourceName := "cloudflare_zero_trust_organization.test"

	// Use v5 config directly to test v5→v5 pass-through (no migration needed)
	acctest.MigrationV2TestStep(
		t,
		v5WithCustomPagesConfig+accountIDConfig,
		tmpDir,
		acctest.GetLastV4Version(),
		"v5",
		"v5",
		[]statecheck.StateCheck{
			// Verify custom_pages object syntax works in v5
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("custom_pages").AtMapKey("forbidden"), knownvalue.StringExact("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("custom_pages").AtMapKey("identity_denied"), knownvalue.StringExact("yyyyyyyy-yyyy-yyyy-yyyy-yyyyyyyyyyyy")),

			// Verify core fields
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact("Test Organization")),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_domain"), knownvalue.StringExact("test.cloudflareaccess.com")),
		},
	)
}

func TestMigrateZeroTrustOrganization_V4ToV5_Minimal(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for this test")
	}

	// TF_MIG_TEST must be set to activate migration testing
	t.Setenv("TF_MIG_TEST", "1")

	tmpDir := t.TempDir()
	accountIDConfig := fmt.Sprintf(`
variable "account_id" {
  type    = string
  default = "%s"
}
`, accountID)

	resourceName := "cloudflare_zero_trust_organization.test"

	acctest.MigrationV2TestStep(
		t,
		v4MinimalConfig+accountIDConfig,
		tmpDir,
		acctest.GetLastV4Version(),
		"v4",
		"v5",
		[]statecheck.StateCheck{
			// Verify minimal required fields
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact("Minimal Organization")),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_domain"), knownvalue.StringExact("minimal.cloudflareaccess.com")),

			// Critical: Verify boolean defaults were added for minimal config
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_authenticate_via_warp"), knownvalue.Bool(false)),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_redirect_to_identity"), knownvalue.Bool(false)),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_ui_read_only"), knownvalue.Bool(false)),
		},
	)
}

// TestMigrateZeroTrustOrganization_V5ToV5_Minimal verifies that minimal v5 resources
// work correctly without migration (v5→v5 pass-through).
func TestMigrateZeroTrustOrganization_V5ToV5_Minimal(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for this test")
	}

	// TF_MIG_TEST must be set to activate migration testing
	t.Setenv("TF_MIG_TEST", "1")

	tmpDir := t.TempDir()
	accountIDConfig := fmt.Sprintf(`
variable "account_id" {
  type    = string
  default = "%s"
}
`, accountID)

	resourceName := "cloudflare_zero_trust_organization.test"

	// Use v5 config directly to test v5→v5 pass-through (no migration needed)
	acctest.MigrationV2TestStep(
		t,
		v5MinimalConfig+accountIDConfig,
		tmpDir,
		acctest.GetLastV4Version(),
		"v5",
		"v5",
		[]statecheck.StateCheck{
			// Verify minimal required fields
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact("Minimal Organization")),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_domain"), knownvalue.StringExact("minimal.cloudflareaccess.com")),

			// Verify boolean defaults exist (computed defaults in v5)
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_authenticate_via_warp"), knownvalue.Bool(false)),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_redirect_to_identity"), knownvalue.Bool(false)),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("is_ui_read_only"), knownvalue.Bool(false)),
		},
	)
}

// TestMigrateZeroTrustOrganization_V4ToV5_NoDrift verifies that the migration doesn't introduce drift.
// After migrating from v4 to v5, running terraform plan should show no changes.
func TestMigrateZeroTrustOrganization_V4ToV5_NoDrift(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for this test")
	}

	// TF_MIG_TEST must be set to activate migration testing
	t.Setenv("TF_MIG_TEST", "1")

	tmpDir := t.TempDir()
	accountIDConfig := fmt.Sprintf(`
variable "account_id" {
  type    = string
  default = "%s"
}
`, accountID)

	resourceName := "cloudflare_zero_trust_organization.test"

	// Use MigrationV2TestStep which already verifies no drift via ExpectEmptyPlan
	acctest.MigrationV2TestStep(
		t,
		v4BasicConfig+accountIDConfig,
		tmpDir,
		acctest.GetLastV4Version(),
		"v4",
		"v5",
		[]statecheck.StateCheck{
			// Just verify the resource exists with expected fields
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact("Test Organization")),
		},
	)
}
