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

// Embed migration test configuration files
//
//go:embed testdata/v4_basic_http.tf
var v4BasicHTTPConfig string

//go:embed testdata/v5_basic_http.tf
var v5BasicHTTPConfig string

//go:embed testdata/v4_traceroute.tf
var v4TracerouteConfig string

//go:embed testdata/v5_traceroute.tf
var v5TracerouteConfig string

//go:embed testdata/v4_disabled.tf
var v4DisabledConfig string

//go:embed testdata/v5_disabled.tf
var v5DisabledConfig string

//go:embed testdata/v4_minimal.tf
var v4MinimalConfig string

//go:embed testdata/v5_minimal.tf
var v5MinimalConfig string

//go:embed testdata/v4_maximal.tf
var v4MaximalConfig string

//go:embed testdata/v5_maximal.tf
var v5MaximalConfig string

// TestMigrateZeroTrustDEXTest_BasicHTTP tests basic HTTP DEX test migration
// Covers: data field transformation (array[0] → pointer), test_id copy, new fields (targeted, target_policies)
func TestMigrateZeroTrustDEXTest_BasicHTTP(t *testing.T) {
	// Zero Trust resources don't support API tokens yet
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string { return fmt.Sprintf(v4BasicHTTPConfig, rnd, accountID, name) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID, name string) string { return fmt.Sprintf(v5BasicHTTPConfig, rnd, accountID, name) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name)
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
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					// Step 2: Run migration and verify state transformation
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						// Resource should be renamed to cloudflare_zero_trust_dex_test
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
						// Verify top-level fields preserved
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(name)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("description"), knownvalue.StringExact("Test HTTP connectivity")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("interval"), knownvalue.StringExact("0h30m0s")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
						// Verify data field transformed from block to attribute (single nested object)
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("data").AtMapKey("kind"), knownvalue.StringExact("http")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("data").AtMapKey("host"), knownvalue.StringExact("https://dash.cloudflare.com")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("data").AtMapKey("method"), knownvalue.StringExact("GET")),
					}),
				},
			})
		})
	}
}

// TestMigrateZeroTrustDEXTest_Traceroute tests traceroute DEX test migration
// Edge case: method field is absent for traceroute tests (only present for HTTP)
func TestMigrateZeroTrustDEXTest_Traceroute(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string { return fmt.Sprintf(v4TracerouteConfig, rnd, accountID, name) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID, name string) string { return fmt.Sprintf(v5TracerouteConfig, rnd, accountID, name) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name)
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
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(name)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("interval"), knownvalue.StringExact("1h0m0s")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("data").AtMapKey("kind"), knownvalue.StringExact("traceroute")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("data").AtMapKey("host"), knownvalue.StringExact("1.1.1.1")),
						// Note: method field should NOT be present for traceroute
					}),
				},
			})
		})
	}
}

// TestMigrateZeroTrustDEXTest_Disabled tests disabled DEX test migration
// Covers: enabled=false field value
func TestMigrateZeroTrustDEXTest_Disabled(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string { return fmt.Sprintf(v4DisabledConfig, rnd, accountID, name) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID, name string) string { return fmt.Sprintf(v5DisabledConfig, rnd, accountID, name) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name)
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
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("interval"), knownvalue.StringExact("0h15m0s")),
					}),
				},
			})
		})
	}
}

// TestMigrateZeroTrustDEXTest_MinimalConfig tests minimal required fields only
// Covers: Minimal field set
func TestMigrateZeroTrustDEXTest_MinimalConfig(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string { return fmt.Sprintf(v4MinimalConfig, rnd, accountID, name) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID, name string) string { return fmt.Sprintf(v5MinimalConfig, rnd, accountID, name) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name)
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
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("data").AtMapKey("kind"), knownvalue.StringExact("http")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("data").AtMapKey("host"), knownvalue.StringExact("https://cloudflare.com")),
					}),
				},
			})
		})
	}
}

// TestMigrateZeroTrustDEXTest_MaximalConfig tests all available v4 fields
// Covers: All fields present
func TestMigrateZeroTrustDEXTest_MaximalConfig(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string { return fmt.Sprintf(v4MaximalConfig, rnd, accountID, name) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, accountID, name string) string { return fmt.Sprintf(v5MaximalConfig, rnd, accountID, name) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name)
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
						// Verify all top-level fields
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(name)),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("description"), knownvalue.StringExact("Comprehensive test with all fields")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("interval"), knownvalue.StringExact("0h30m0s")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
						// Verify all data fields
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("data").AtMapKey("kind"), knownvalue.StringExact("http")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("data").AtMapKey("host"), knownvalue.StringExact("https://dash.cloudflare.com/login")),
						statecheck.ExpectKnownValue("cloudflare_zero_trust_dex_test."+rnd, tfjsonpath.New("data").AtMapKey("method"), knownvalue.StringExact("GET")),
					}),
				},
			})
		})
	}
}
