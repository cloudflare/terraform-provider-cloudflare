package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v4_with_advertisement.tf
var v4WithAdvertisementConfig string

//go:embed testdata/v4_minimal.tf
var v4MinimalConfig string

// TestMigrateBYOIPPrefix_Basic tests basic config migration with description.
//
// This test verifies tf-migrate correctly transforms the config from v4 to v5:
// - Removes prefix_id field (v4 only)
// - Removes advertisement field (v4 only)
// - Adds MIGRATION WARNING comment with instructions for asn/cidr/loa_document_id
// - Preserves account_id and description fields
//
// IMPORTANT: This test only validates config transformation, NOT state migration.
//
// Why Config-Only Testing?
// ------------------------
// BYO IP Prefix resources have unique constraints that make E2E migration testing
// impractical:
//
// 1. Resources represent real IP addresses owned by users
// 2. Cloudflare API prevents deletion/recreation (returns: prefix_exists_for_cidr)
// 3. Critical fields (asn, cidr, loa_document_id) are immutable after creation
// 4. State upgrade approach fails (sets new required fields to null → plans replacement)
//
// Migration Path for Users:
// -------------------------
// Users must follow an IMPORT-BASED migration path (not state upgrade):
//
// 1. Query Cloudflare API for actual values (asn, cidr, loa_document_id)
// 2. Update config to v5 format with exact API values
// 3. Remove from v4 state: terraform state rm cloudflare_byo_ip_prefix.X
// 4. Import with v5: terraform import cloudflare_byo_ip_prefix.X account/prefix
// 5. Verify: terraform plan (should show no replacement if values match)
//
// See MIGRATION_GUIDE.md for detailed step-by-step instructions.
//
// Test Validation:
// ----------------
// Manual testing (test_import_with_actual_values.sh) has validated that the
// import-based migration path works when users provide correct values from the API.
//
// State upgrade logic is tested separately in the provider's state upgrader tests.
func TestMigrateBYOIPPrefix_Basic(t *testing.T) {
	tmpDir := t.TempDir()

	// Test data from environment variables
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	prefixID := os.Getenv("CLOUDFLARE_BYO_IP_PREFIX_ID")
	rnd := "test"

	// Skip if required environment variables are not set
	if accountID == "" || prefixID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID and CLOUDFLARE_BYO_IP_PREFIX_ID must be set for this test")
	}

	// Prepare v4 config
	testConfig := fmt.Sprintf(v4BasicConfig, rnd, accountID, prefixID)

	// Write v4 config
	acctest.WriteOutConfig(t, testConfig, tmpDir)

	// Run tf-migrate
	sourceVer, targetVer := acctest.InferMigrationVersions("4.52.5")
	acctest.RunMigrationV2Command(t, testConfig, tmpDir, sourceVer, targetVer)

	// Verify migration results
	verifyBYOIPPrefixMigration(t, tmpDir, accountID, true)
}

// TestMigrateBYOIPPrefix_WithAdvertisement tests that the v4 advertisement field is
// correctly dropped during migration and replaced by the v5 computed advertised field.
//
// The test verifies config transformation only (no apply/plan operations).
// See TestMigrateBYOIPPrefix_Basic for details on the testing approach.
func TestMigrateBYOIPPrefix_WithAdvertisement(t *testing.T) {
	tmpDir := t.TempDir()

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	prefixID := os.Getenv("CLOUDFLARE_BYO_IP_PREFIX_ID")
	rnd := "test"

	// Skip if required environment variables are not set
	if accountID == "" || prefixID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID and CLOUDFLARE_BYO_IP_PREFIX_ID must be set for this test")
	}

	testConfig := fmt.Sprintf(v4WithAdvertisementConfig, rnd, accountID, prefixID, "on")

	acctest.WriteOutConfig(t, testConfig, tmpDir)

	sourceVer, targetVer := acctest.InferMigrationVersions("4.52.5")
	acctest.RunMigrationV2Command(t, testConfig, tmpDir, sourceVer, targetVer)

	verifyBYOIPPrefixMigration(t, tmpDir, accountID, true)
}

// TestMigrateBYOIPPrefix_Minimal tests migration with only the required v4 fields,
// verifying that v5 migration warning is added for required fields.
//
// The test verifies config transformation only (no apply/plan operations).
// See TestMigrateBYOIPPrefix_Basic for details on the testing approach.
func TestMigrateBYOIPPrefix_Minimal(t *testing.T) {
	tmpDir := t.TempDir()

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	prefixID := os.Getenv("CLOUDFLARE_BYO_IP_PREFIX_ID")
	rnd := "test"

	// Skip if required environment variables are not set
	if accountID == "" || prefixID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID and CLOUDFLARE_BYO_IP_PREFIX_ID must be set for this test")
	}

	testConfig := fmt.Sprintf(v4MinimalConfig, rnd, accountID, prefixID)

	acctest.WriteOutConfig(t, testConfig, tmpDir)

	sourceVer, targetVer := acctest.InferMigrationVersions("4.52.5")
	acctest.RunMigrationV2Command(t, testConfig, tmpDir, sourceVer, targetVer)

	verifyBYOIPPrefixMigration(t, tmpDir, accountID, false)
}

// verifyBYOIPPrefixMigration verifies that tf-migrate correctly transformed the config:
// - Removed v4-only fields (prefix_id, advertisement)
// - Added MIGRATION WARNING comment
// - Preserved core fields (account_id, description)
func verifyBYOIPPrefixMigration(t *testing.T, tmpDir, accountID string, expectDescription bool) {
	// Find migrated files
	migratedFiles, err := filepath.Glob(filepath.Join(tmpDir, "*.tf"))
	require.NoError(t, err, "Failed to find migrated files")
	require.NotEmpty(t, migratedFiles, "No .tf files found after migration")

	// Read migrated config
	var migratedContent string
	for _, file := range migratedFiles {
		content, err := os.ReadFile(file)
		require.NoError(t, err, "Failed to read %s", file)

		if strings.Contains(string(content), "cloudflare_byo_ip_prefix") {
			migratedContent = string(content)
			break
		}
	}

	require.NotEmpty(t, migratedContent, "No file containing cloudflare_byo_ip_prefix found")

	// Verify v4-only fields are removed
	assert.NotContains(t, migratedContent, "prefix_id",
		"Config should not contain prefix_id after migration")
	assert.NotContains(t, migratedContent, "advertisement =",
		"Config should not contain advertisement field after migration")

	// Verify migration warning is present
	assert.Contains(t, migratedContent, "MIGRATION WARNING",
		"Config should contain MIGRATION WARNING comment")
	assert.Contains(t, migratedContent, "manual intervention",
		"Warning should mention manual intervention")
	assert.Contains(t, migratedContent, "asn",
		"Warning should mention asn field")
	assert.Contains(t, migratedContent, "cidr",
		"Warning should mention cidr field")
	assert.Contains(t, migratedContent, "Cloudflare Dashboard",
		"Warning should mention where to find values")

	// Verify core fields are preserved
	assert.Contains(t, migratedContent, "account_id",
		"Config should preserve account_id")
	assert.Contains(t, migratedContent, accountID,
		"Config should preserve account_id value")

	if expectDescription {
		assert.Contains(t, migratedContent, "description",
			"Config should preserve description field")
	}
}
