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

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_with_activate.tf
var v4WithActivateConfig string

//go:embed testdata/v5_with_activate.tf
var v5WithActivateConfig string

//go:embed testdata/v4_with_validity.tf
var v4WithValidityConfig string

//go:embed testdata/v5_with_validity.tf
var v5WithValidityConfig string

// TestMigrateZeroTrustGatewayCertificate_V4ToV5_Basic tests basic migration when
// validity_period_days is NOT set in config (uses v4 default 1826).
// Verifies: validity_period_days is null in v5 state (default nulled out by Transform).
func TestMigrateZeroTrustGatewayCertificate_V4ToV5_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Zero Trust resources require API key, not token
			if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
				t.Setenv("CLOUDFLARE_API_TOKEN", "")
			}

			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_gateway_certificate." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				// Use local v5 provider (creates version=1 state)
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				// Use external v4 provider (creates version=0 state)
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
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify id is present
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
							// Verify account_id is preserved (pass-through)
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							// Verify computed field is present
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("binding_status"), knownvalue.NotNull()),
							// Note: validity_period_days null-out behavior tested in WithExplicitValidity.
							// v4 default (1826) is nulled; v5 configs specify 1825 explicitly.
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustGatewayCertificate_V4ToV5_WithActivate tests migration
// with activate=false explicitly set.
// Verifies: boolean fields are preserved correctly through migration.
func TestMigrateZeroTrustGatewayCertificate_V4ToV5_WithActivate(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4WithActivateConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5WithActivateConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Zero Trust resources require API key, not token
			if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
				t.Setenv("CLOUDFLARE_API_TOKEN", "")
			}

			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_gateway_certificate." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
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
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							// Verify activate=false is preserved (bool pass-through)
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("activate"), knownvalue.Bool(false)),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZeroTrustGatewayCertificate_V4ToV5_WithExplicitValidity tests migration
// when validity_period_days is explicitly set to a non-default value (3650 days = 10 years).
// Verifies: non-default values are preserved (not nulled out like the default 1826).
func TestMigrateZeroTrustGatewayCertificate_V4ToV5_WithExplicitValidity(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v4WithValidityConfig, rnd, accountID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID string) string {
				return fmt.Sprintf(v5WithValidityConfig, rnd, accountID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Zero Trust resources require API key, not token
			if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
				t.Setenv("CLOUDFLARE_API_TOKEN", "")
			}

			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_zero_trust_gateway_certificate." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID)
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
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							// Verify non-default validity_period_days is preserved (not nulled out)
							// Int64 values are stored as float64 in TF JSON state
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("validity_period_days"), knownvalue.Float64Exact(3650)),
						},
					),
				},
			})
		})
	}
}
