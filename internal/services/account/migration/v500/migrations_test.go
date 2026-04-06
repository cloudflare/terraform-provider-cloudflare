package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Embed migration test configuration files
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

// TestMigrateAccount_V4ToV5_Basic tests migration of a basic account
// with minimal configuration. This verifies:
// 1. enforce_twofactor moves from top-level to settings.enforce_twofactor
// 2. All required fields are preserved (name)
// 3. Migration works from both v4 and v5
//
// Note: We test with enforce_twofactor=false only because the Cloudflare API
// ignores enforce_twofactor on the POST (create) endpoint. Testing with
// enforce_twofactor=true would require a create-then-update flow which is
// outside the scope of migration testing. See PT-792.
func TestMigrateAccount_V4ToV5_Basic(t *testing.T) {
	acctest.TestAccSkipForDefaultAccount(t, "Requires account creation permissions not available on default test account.")

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, name string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, name string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			name := fmt.Sprintf("tf-test-%s", rnd)
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			// For v5 tests, use local provider; for v4 tests, use external provider
			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				// Use local v5 provider (has GetSchemaVersion, will create version=1 state)
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				// Use external v4 provider (will create version=0 state)
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_account."+rnd,
								tfjsonpath.New("id"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_account."+rnd,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name),
							),
							// Verify enforce_twofactor is now in settings
							statecheck.ExpectKnownValue(
								"cloudflare_account."+rnd,
								tfjsonpath.New("settings").AtMapKey("enforce_twofactor"),
								knownvalue.Bool(false),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateAccount_V4ToV5_SpecificVersions tests migration from specific v4 versions
func TestMigrateAccount_V4ToV5_SpecificVersions(t *testing.T) {
	acctest.TestAccSkipForDefaultAccount(t, "Requires account creation permissions not available on default test account.")

	// Test migration from the last known v4 version
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("tf-test-v4-%s", rnd)
	tmpDir := t.TempDir()

	// V4 config - enforce_twofactor is a top-level attribute
	v4Config := fmt.Sprintf(v4BasicConfig, rnd, name)

	// V5 config - enforce_twofactor is nested in settings
	v5Config := fmt.Sprintf(v5BasicConfig, rnd, name)

	_ = os.Getenv("CLOUDFLARE_ACCOUNT_ID") // ensure env is available

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: acctest.GetLastV4Version(),
					},
				},
				Config: v4Config,
			},
			{
				// Step 2: Upgrade to latest provider with state upgrader
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   v5Config,
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify name is preserved
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account.%s", rnd),
						tfjsonpath.New("name"),
						knownvalue.StringExact(name),
					),
					// Verify enforce_twofactor moved to settings
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account.%s", rnd),
						tfjsonpath.New("settings").AtMapKey("enforce_twofactor"),
						knownvalue.Bool(false),
					),
				},
			},
		},
	})
}
