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
// Singular name configs (cloudflare_worker_script → cloudflare_workers_script rename)

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_with_binding.tf
var v4WithBindingConfig string

//go:embed testdata/v5_with_binding.tf
var v5WithBindingConfig string

//go:embed testdata/v4_multiple_bindings.tf
var v4MultipleBindingsConfig string

//go:embed testdata/v5_multiple_bindings.tf
var v5MultipleBindingsConfig string

// Plural name configs (cloudflare_workers_script — no rename, just attribute changes)

//go:embed testdata/v4_basic_plural.tf
var v4BasicPluralConfig string

//go:embed testdata/v4_with_binding_plural.tf
var v4WithBindingPluralConfig string

//go:embed testdata/v4_multiple_bindings_plural.tf
var v4MultipleBindingsPluralConfig string

//go:embed testdata/v4_complex_bindings.tf
var v4ComplexBindingsConfig string

// ============================================================================
// Ported from migrations_test.go (service root) — original v4-only tests
// using cloudflare_workers_script (plural, no rename)
// ============================================================================

// TestMigrateWorkersScript_FromV4Basic tests basic migration from v4 to v5.
// Ported from TestMigrateWorkersScriptMigrationFromV4Basic.
// Uses plural name (cloudflare_workers_script) — no resource rename, just name→script_name.
func TestMigrateWorkersScript_FromV4Basic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script." + rnd
	tmpDir := t.TempDir()
	scriptName := fmt.Sprintf("test-script-%s", rnd)
	version := acctest.GetLastV4Version()
	testConfig := fmt.Sprintf(v4BasicPluralConfig, rnd, accountID, scriptName)
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: version,
					},
				},
				Config: testConfig,
			},
			acctest.MigrationV2TestStep(t, testConfig, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(scriptName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("content"), knownvalue.StringExact("addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });")),
			}),
		},
	})
}

// TestMigrateWorkersScript_FromV4WithBindings tests migration with bindings.
// Ported from TestMigrateWorkersScriptMigrationFromV4WithBindings.
func TestMigrateWorkersScript_FromV4WithBindings(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script." + rnd
	tmpDir := t.TempDir()
	scriptName := fmt.Sprintf("test-script-%s", rnd)
	version := acctest.GetLastV4Version()
	testConfig := fmt.Sprintf(v4WithBindingPluralConfig, rnd, accountID, scriptName)
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: version,
					},
				},
				Config: testConfig,
			},
			acctest.MigrationV2TestStep(t, testConfig, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(scriptName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("MY_VAR")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("type"), knownvalue.StringExact("plain_text")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("text"), knownvalue.StringExact("my-value")),
			}),
		},
	})
}

// TestMigrateWorkersScript_FromV4SingleResource tests migration using singular resource name.
// Ported from TestMigrateWorkersScriptMigrationFromV4SingleResource.
// Uses cloudflare_worker_script (singular) — tests the resource rename path.
func TestMigrateWorkersScript_FromV4SingleResource(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script." + rnd
	tmpDir := t.TempDir()
	scriptName := fmt.Sprintf("test-script-%s", rnd)
	version := acctest.GetLastV4Version()
	testConfig := fmt.Sprintf(v4BasicConfig, rnd, accountID, scriptName)
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: version,
					},
				},
				Config: testConfig,
			},
			acctest.MigrationV2TestStep(t, testConfig, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(scriptName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("content"), knownvalue.StringExact("addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });")),
			}),
		},
	})
}

// TestMigrateWorkersScript_FromV4MultipleBindingTypes tests migration with multiple binding types.
// Ported from TestMigrateWorkersScriptMigrationFromV4MultipleBindingTypes.
func TestMigrateWorkersScript_FromV4MultipleBindingTypes(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script." + rnd
	tmpDir := t.TempDir()
	scriptName := fmt.Sprintf("test-script-%s", rnd)
	version := acctest.GetLastV4Version()
	testConfig := fmt.Sprintf(v4MultipleBindingsPluralConfig, rnd, accountID, scriptName)
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: version,
					},
				},
				Config: testConfig,
			},
			acctest.MigrationV2TestStep(t, testConfig, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(scriptName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings"), knownvalue.ListSizeExact(2)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("MY_VAR")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("type"), knownvalue.StringExact("plain_text")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("text"), knownvalue.StringExact("my-value")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(1).AtMapKey("name"), knownvalue.StringExact("MY_SECRET")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(1).AtMapKey("type"), knownvalue.StringExact("secret_text")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(1).AtMapKey("text"), knownvalue.StringExact("secret-value")),
			}),
		},
	})
}

// TestMigrateWorkersScript_FromV4ComplexBindings tests migration with complex binding attribute mappings.
// Ported from TestMigrateWorkersScriptMigrationFromV4ComplexBindings.
func TestMigrateWorkersScript_FromV4ComplexBindings(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script." + rnd
	tmpDir := t.TempDir()
	scriptName := fmt.Sprintf("test-script-%s", rnd)
	version := acctest.GetLastV4Version()
	testConfig := fmt.Sprintf(v4ComplexBindingsConfig, rnd, accountID, scriptName)
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: version,
					},
				},
				Config: testConfig,
			},
			acctest.MigrationV2TestStep(t, testConfig, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(scriptName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("type"), knownvalue.StringExact("plain_text")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("SIMPLE_VAR")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("text"), knownvalue.StringExact("simple-value")),
			}),
		},
	})
}

// ============================================================================
// New dual test cases (from_v4_latest / from_v5) using singular name
// Tests the resource rename path (cloudflare_worker_script → cloudflare_workers_script)
// ============================================================================

// TestMigrateWorkersScript_Basic tests basic migration with resource rename.
// Dual test: from_v4_latest (singular name) and from_v5 (plural name).
func TestMigrateWorkersScript_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, scriptName string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, scriptName string) string { return fmt.Sprintf(v4BasicConfig, rnd, accountID, scriptName) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID, scriptName string) string { return fmt.Sprintf(v5BasicConfig, rnd, accountID, scriptName) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			scriptName := fmt.Sprintf("test-script-%s", rnd)
			resourceName := "cloudflare_workers_script." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, scriptName)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
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
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(scriptName)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("content"), knownvalue.StringExact("addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });")),
					}),
				},
			})
		})
	}
}

// TestMigrateWorkersScript_WithBinding tests migration with binding and resource rename.
// Dual test: verifies binding consolidation works through both migration paths.
func TestMigrateWorkersScript_WithBinding(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, scriptName string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, scriptName string) string { return fmt.Sprintf(v4WithBindingConfig, rnd, accountID, scriptName) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID, scriptName string) string { return fmt.Sprintf(v5WithBindingConfig, rnd, accountID, scriptName) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			scriptName := fmt.Sprintf("test-script-%s", rnd)
			resourceName := "cloudflare_workers_script." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, scriptName)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
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
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(scriptName)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings"), knownvalue.ListSizeExact(1)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("type"), knownvalue.StringExact("plain_text")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("MY_VAR")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("text"), knownvalue.StringExact("my-value")),
					}),
				},
			})
		})
	}
}

// TestMigrateWorkersScript_MultipleBindingTypes tests migration with multiple bindings and resource rename.
// Dual test: verifies consolidation of plain_text + secret_text works through both paths.
func TestMigrateWorkersScript_MultipleBindingTypes(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, scriptName string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, scriptName string) string { return fmt.Sprintf(v4MultipleBindingsConfig, rnd, accountID, scriptName) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID, scriptName string) string { return fmt.Sprintf(v5MultipleBindingsConfig, rnd, accountID, scriptName) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			scriptName := fmt.Sprintf("test-script-%s", rnd)
			resourceName := "cloudflare_workers_script." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, scriptName)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
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
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(scriptName)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings"), knownvalue.ListSizeExact(2)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("type"), knownvalue.StringExact("plain_text")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("MY_VAR")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(1).AtMapKey("type"), knownvalue.StringExact("secret_text")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bindings").AtSliceIndex(1).AtMapKey("name"), knownvalue.StringExact("MY_SECRET")),
					}),
				},
			})
		})
	}
}
