package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Migration Test Configuration
//
// Version is read from LAST_V4_VERSION environment variable (set in .github/workflows/migration-tests.yml)
// - Last stable v4 release: default 4.52.5
// - Current v5 release: auto-updates with releases (internal.PackageVersion)
//
// Based on breaking changes analysis:
// - All breaking changes happened between 4.x and 5.0.0
// - No breaking changes between v5 releases (testing against latest v5)
// - Key changes for workers_kv_namespace: none (pass-through migration)

// Embed migration test configuration files
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_special_chars.tf
var v4SpecialCharsConfig string

//go:embed testdata/v5_special_chars.tf
var v5SpecialCharsConfig string

//go:embed testdata/v4_multiple.tf
var v4MultipleConfig string

//go:embed testdata/v5_multiple.tf
var v5MultipleConfig string

// TestMigrateWorkersKVNamespaceBasic tests basic migration from v4 to v5
// This tests both v4→v5 migration and v5→v5 version bump
func TestMigrateWorkersKVNamespaceBasic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, title string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, title string) string { return fmt.Sprintf(v4BasicConfig, rnd, accountID, title) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID, title string) string { return fmt.Sprintf(v5BasicConfig, rnd, accountID, title) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			title := fmt.Sprintf("test-kv-namespace-%s", rnd)
			tmpDir := t.TempDir()
			resourceName := "cloudflare_workers_kv_namespace." + rnd
			testConfig := tc.configFn(rnd, accountID, title)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with specific version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						// Resource name stays the same (not renamed)
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("title"), knownvalue.StringExact(title)),
						// Verify new computed field is present in v5
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("supports_url_encoding"), knownvalue.NotNull()),
					}),
				},
			})
		})
	}
}

// TestMigrateWorkersKVNamespaceWithSpecialChars tests migration with special characters in title
// This tests both v4→v5 migration and v5→v5 version bump
func TestMigrateWorkersKVNamespaceWithSpecialChars(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, title string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, title string) string {
				return fmt.Sprintf(v4SpecialCharsConfig, rnd, accountID, title)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, title string) string {
				return fmt.Sprintf(v5SpecialCharsConfig, rnd, accountID, title)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			// Title with spaces, dashes, and underscores
			title := fmt.Sprintf("Test KV Namespace_%s-2024", rnd)
			tmpDir := t.TempDir()
			resourceName := "cloudflare_workers_kv_namespace." + rnd
			testConfig := tc.configFn(rnd, accountID, title)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with specific version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("title"), knownvalue.StringExact(title)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("supports_url_encoding"), knownvalue.NotNull()),
					}),
				},
			})
		})
	}
}

// TestMigrateWorkersKVNamespaceMultiple tests migration of multiple KV namespaces in one config
// This tests both v4→v5 migration and v5→v5 version bump
func TestMigrateWorkersKVNamespaceMultiple(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd1, rnd2, accountID, title1, title2 string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd1, rnd2, accountID, title1, title2 string) string {
				return fmt.Sprintf(v4MultipleConfig, rnd1, rnd2, accountID, title1, title2)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd1, rnd2, accountID, title1, title2 string) string {
				return fmt.Sprintf(v5MultipleConfig, rnd1, rnd2, accountID, title1, title2)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			rnd1 := rnd + "1"
			rnd2 := rnd + "2"
			title1 := fmt.Sprintf("test-kv-namespace-1-%s", rnd)
			title2 := fmt.Sprintf("test-kv-namespace-2-%s", rnd)
			tmpDir := t.TempDir()
			resourceName1 := "cloudflare_workers_kv_namespace." + rnd1
			resourceName2 := "cloudflare_workers_kv_namespace." + rnd2
			testConfig := tc.configFn(rnd1, rnd2, accountID, title1, title2)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with specific version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						// Verify first namespace
						statecheck.ExpectKnownValue(resourceName1, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName1, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName1, tfjsonpath.New("title"), knownvalue.StringExact(title1)),
						statecheck.ExpectKnownValue(resourceName1, tfjsonpath.New("supports_url_encoding"), knownvalue.NotNull()),
						// Verify second namespace
						statecheck.ExpectKnownValue(resourceName2, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName2, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName2, tfjsonpath.New("title"), knownvalue.StringExact(title2)),
						statecheck.ExpectKnownValue(resourceName2, tfjsonpath.New("supports_url_encoding"), knownvalue.NotNull()),
					}),
				},
			})
		})
	}
}
