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
// Key changes for workers_for_platforms_dispatch_namespace:
//   - namespace_name field added (v5 uses it for Read/Delete; migration copies id → namespace_name)
//   - Multiple new computed fields added (created_by, created_on, modified_by, modified_on, etc.)
//   - cloudflare_workers_for_platforms_namespace (deprecated) renamed to cloudflare_workers_for_platforms_dispatch_namespace

// Embed migration test configuration files

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_deprecated.tf
var v4DeprecatedConfig string

//go:embed testdata/v5_deprecated.tf
var v5DeprecatedConfig string

// TestMigrateWorkersForPlatformsDispatchNamespaceBasic tests migration from v4 dispatch_namespace to v5.
// This covers the UpgradeState path (no resource rename; same type in v4 and v5).
func TestMigrateWorkersForPlatformsDispatchNamespaceBasic(t *testing.T) {
	t.Skip("Migration not enabled yet")
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string { return fmt.Sprintf(v4BasicConfig, rnd, accountID, name) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID, name string) string { return fmt.Sprintf(v5BasicConfig, rnd, accountID, name) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("test-namespace-%s", rnd)
			tmpDir := t.TempDir()
			resourceName := "cloudflare_workers_for_platforms_dispatch_namespace." + rnd
			testConfig := tc.configFn(rnd, accountID, name)
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
						// Verify id and namespace_name are both populated (key transformation)
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("namespace_name"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(name)),
					}),
				},
			})
		})
	}
}

// TestMigrateWorkersForPlatformsNamespaceDeprecated tests migration from the deprecated
// cloudflare_workers_for_platforms_namespace resource type (v4) to
// cloudflare_workers_for_platforms_dispatch_namespace (v5).
// This covers the MoveState path (resource rename triggered by `moved` block).
func TestMigrateWorkersForPlatformsNamespaceDeprecated(t *testing.T) {
	t.Skip("Migration not enabled yet")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("test-namespace-%s", rnd)
	tmpDir := t.TempDir()
	// After migration, resource is referenced as cloudflare_workers_for_platforms_dispatch_namespace
	resourceName := "cloudflare_workers_for_platforms_dispatch_namespace." + rnd
	v4Config := fmt.Sprintf(v4DeprecatedConfig, rnd, accountID, name)
	v4Version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(v4Version)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider using deprecated resource type
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: v4Version,
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration (renames type + generates moved block) and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, v4Version, sourceVer, targetVer, []statecheck.StateCheck{
				// Verify resource moved to new type and state transformed correctly
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("namespace_name"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(name)),
			}),
		},
	})
}
