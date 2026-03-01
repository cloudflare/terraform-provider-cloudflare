package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
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
// - Resource renamed: cloudflare_authenticated_origin_pulls → cloudflare_authenticated_origin_pulls_settings
// - Fields removed in v5: hostname, authenticated_origin_pulls_certificate
// - Fields kept: zone_id, enabled, id

// Embed migration test configuration files
//
//go:embed testdata/v4_zone_wide.tf
var v4ZoneWideConfig string

//go:embed testdata/v5_zone_wide.tf
var v5ZoneWideConfig string

//go:embed testdata/v4_with_hostname.tf
var v4WithHostnameConfig string

//go:embed testdata/v4_with_certificate.tf
var v4WithCertificateConfig string

//go:embed testdata/v4_disabled.tf
var v4DisabledConfig string

//go:embed testdata/v5_disabled.tf
var v5DisabledConfig string

// aopMigrationTestStep creates a test step for authenticated_origin_pulls_settings migration
func aopMigrationTestStep(t *testing.T, testConfig string, tmpDir string, exactVersion string, sourceVersion string, targetVersion string, expectNonEmptyPlan bool, stateChecks []statecheck.StateCheck) resource.TestStep {
	return resource.TestStep{
		PreConfig: func() {
			acctest.WriteOutConfig(t, testConfig, tmpDir)
			acctest.RunMigrationV2Command(t, testConfig, tmpDir, sourceVersion, targetVersion)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		ConfigDirectory:          config.StaticDirectory(tmpDir),
		ExpectNonEmptyPlan:       expectNonEmptyPlan,
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				acctest.DebugNonEmptyPlan,
			},
		},
		ConfigStateChecks: stateChecks,
	}
}

// TestMigrateAuthenticatedOriginPullsSettingsBasic tests migration of basic zone-wide AOP from v4 to v5
func TestMigrateAuthenticatedOriginPullsSettingsBasic(t *testing.T) {
	testCases := []struct {
		name               string
		version            string
		configFn           func(rnd, zoneID string) string
		expectNonEmptyPlan bool
	}{
		{
			name:     "from_v4_latest",
			version:  os.Getenv("LAST_V4_VERSION"),
			configFn: func(rnd, zoneID string) string { return fmt.Sprintf(v4ZoneWideConfig, rnd, zoneID) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, zoneID string) string { return fmt.Sprintf(v5ZoneWideConfig, rnd, zoneID) },
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
					aopMigrationTestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, tc.expectNonEmptyPlan, []statecheck.StateCheck{
						// Resource should be renamed to cloudflare_authenticated_origin_pulls_settings
						statecheck.ExpectKnownValue("cloudflare_authenticated_origin_pulls_settings."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue("cloudflare_authenticated_origin_pulls_settings."+rnd, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					}),
				},
			})
		})
	}
}

// TestMigrateAuthenticatedOriginPullsSettingsDisabled tests migration of disabled AOP from v4 to v5
func TestMigrateAuthenticatedOriginPullsSettingsDisabled(t *testing.T) {
	testCases := []struct {
		name               string
		version            string
		configFn           func(rnd, zoneID string) string
		expectNonEmptyPlan bool
	}{
		{
			name:     "from_v4_latest",
			version:  os.Getenv("LAST_V4_VERSION"),
			configFn: func(rnd, zoneID string) string { return fmt.Sprintf(v4DisabledConfig, rnd, zoneID) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, zoneID string) string { return fmt.Sprintf(v5DisabledConfig, rnd, zoneID) },
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
					aopMigrationTestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, tc.expectNonEmptyPlan, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_authenticated_origin_pulls_settings."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue("cloudflare_authenticated_origin_pulls_settings."+rnd, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					}),
				},
			})
		})
	}
}

// TestMigrateAuthenticatedOriginPullsSettingsWithHostname tests migration with v4-only hostname field
// The hostname field should be removed in the v5 state
func TestMigrateAuthenticatedOriginPullsSettingsWithHostname(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V4 config with hostname (should be removed in v5)
	v4Config := fmt.Sprintf(v4WithHostnameConfig, rnd, zoneID)

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
						VersionConstraint: os.Getenv("LAST_V4_VERSION"),
					},
				},
				Config: v4Config,
			},
			aopMigrationTestStep(t, v4Config, tmpDir, os.Getenv("LAST_V4_VERSION"), "v4", "v5", false, []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_authenticated_origin_pulls_settings."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				// hostname should not exist in v5 state
			}),
		},
	})
}

// TestMigrateAuthenticatedOriginPullsSettingsWithCertificate tests migration with v4-only certificate field
// The authenticated_origin_pulls_certificate field should be removed in the v5 state
func TestMigrateAuthenticatedOriginPullsSettingsWithCertificate(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V4 config with certificate reference (should be removed in v5)
	v4Config := fmt.Sprintf(v4WithCertificateConfig, rnd, zoneID)

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
						VersionConstraint: os.Getenv("LAST_V4_VERSION"),
					},
				},
				Config: v4Config,
			},
			aopMigrationTestStep(t, v4Config, tmpDir, os.Getenv("LAST_V4_VERSION"), "v4", "v5", false, []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_authenticated_origin_pulls_settings."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				// authenticated_origin_pulls_certificate should not exist in v5 state
			}),
		},
	})
}
