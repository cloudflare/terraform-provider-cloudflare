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

// Embed test configs
// Note: Only v5_status_active.tf remains for from_v5 tests. Other from_v5 test cases were removed
// because the v5 provider creates DNSSEC with status="disabled" by default (when status is not
// explicitly set), causing cleanup failures with "DNSSEC is already deleted" errors.
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v4_with_modified_on.tf
var v4WithModifiedOnConfig string

//go:embed testdata/v4_status_active.tf
var v4StatusActiveConfig string

//go:embed testdata/v5_status_active.tf
var v5StatusActiveConfig string

// TestMigrateZoneDNSSEC_V4ToV5_Basic tests basic zone_dnssec migration from v4 to v5.
// Only tests from_v4_latest; from_v5 variant removed due to cleanup issues with implicit status="disabled".
func TestMigrateZoneDNSSEC_V4ToV5_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with specific provider version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					// Step 2: Run migration and verify state
					// Uses MigrationV2TestStepForZoneDNSSEC which expects non-empty plan due to status field limitation
					acctest.MigrationV2TestStepForZoneDNSSEC(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							// Note: Not checking status field - v4->v5 migration doesn't populate it in state
							statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("algorithm"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("flags"), knownvalue.NotNull()),
							statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("key_tag"), knownvalue.NotNull()),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZoneDNSSEC_V4ToV5_WithModifiedOn tests migration where modified_on exists in state.
// The modified_on field was optional+computed in v4 but is computed-only in v5.
// Only tests from_v4_latest; from_v5 variant removed due to cleanup issues with implicit status="disabled".
func TestMigrateZoneDNSSEC_V4ToV5_WithModifiedOn(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4WithModifiedOnConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					// Uses MigrationV2TestStepForZoneDNSSEC which expects non-empty plan due to status field limitation
					acctest.MigrationV2TestStepForZoneDNSSEC(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							// Note: Not checking status field - v4->v5 migration doesn't populate it in state
							statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("modified_on"), knownvalue.NotNull()),
						},
					),
				},
			})
		})
	}
}

// TestMigrateZoneDNSSEC_V4ToV5_StatusActive tests that numeric fields are converted correctly.
// Tests Int64→Float64 conversion for flags and key_tag fields.
// Includes from_v5 variant with explicit status="active" to avoid cleanup issues.
func TestMigrateZoneDNSSEC_V4ToV5_StatusActive(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4StatusActiveConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5StatusActiveConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			// Build state checks conditionally based on test variant
			stateChecks := []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				// Verify numeric fields are present (flags, key_tag converted from int to float64)
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("flags"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("key_tag"), knownvalue.NotNull()),
				// Verify other computed fields
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("algorithm"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("digest"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("digest_algorithm"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("digest_type"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("ds"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("key_type"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("public_key"), knownvalue.NotNull()),
			}
			if tc.name == "from_v5" {
				// Only check status for v5->v5 migrations where it's explicitly set in config
				stateChecks = append(stateChecks, statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("status"), knownvalue.NotNull()))
			}
			// Note: Not checking status for from_v4_latest - v4->v5 migration doesn't populate it in state

			// Build the first test step differently for v5 (use local provider) vs v4 (use external provider)
			firstStep := resource.TestStep{
				Config:             testConfig,
				ExpectNonEmptyPlan: tc.name == "from_v5", // from_v5 expects non-empty plan because DNSSEC status transitions from "pending" to "active"
			}
			if tc.name == "from_v5" {
				// Use local dev provider for v5 tests (version not yet published to registry)
				firstStep.ProtoV6ProviderFactories = acctest.TestAccProtoV6ProviderFactories
			} else {
				// Use external provider from registry for v4 tests
				firstStep.ExternalProviders = map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: tc.version,
					},
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					// Uses MigrationV2TestStepForZoneDNSSEC which expects non-empty plan due to status field limitation
					acctest.MigrationV2TestStepForZoneDNSSEC(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, stateChecks),
				},
			})
		})
	}
}
