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
	currentProviderVersion = internal.PackageVersion
)

// Embed migration test configuration files.
//
// v4 configs use cloudflare_zone_settings_override (the v4 resource).
//
// v5 configs come in two variants:
//   - plain (e.g. v5_on_off_plain.tf): just the cloudflare_zone_setting resource,
//     used as the first step in from_v5 tests to create the resource fresh.
//   - with import+removed blocks (e.g. v5_on_off.tf): the full tf-migrate output,
//     used as the v4Config argument to MigrationV2TestStep so tf-migrate can
//     transform it and verify the result.

//go:embed testdata/v4_on_off.tf
var v4OnOffConfig string

//go:embed testdata/v5_on_off.tf
var v5OnOffConfig string

//go:embed testdata/v5_on_off_plain.tf
var v5OnOffPlainConfig string

//go:embed testdata/v4_number.tf
var v4NumberConfig string

//go:embed testdata/v5_number.tf
var v5NumberConfig string

//go:embed testdata/v5_number_plain.tf
var v5NumberPlainConfig string

//go:embed testdata/v4_security_header.tf
var v4SecurityHeaderConfig string

//go:embed testdata/v5_security_header.tf
var v5SecurityHeaderConfig string

//go:embed testdata/v5_security_header_plain.tf
var v5SecurityHeaderPlainConfig string

//go:embed testdata/v4_multiple.tf
var v4MultipleConfig string

//go:embed testdata/v5_multiple.tf
var v5MultipleConfig string

//go:embed testdata/v5_multiple_plain.tf
var v5MultiplePlainConfig string

// TestMigrateZoneSetting_OnOff tests migration of a simple on/off string setting (http3).
//
// v4: cloudflare_zone_settings_override with settings { http3 = "on" }
// v5: cloudflare_zone_setting with setting_id = "http3", value = "on"
//
// The v5 config includes import + removed blocks so Terraform adopts the
// existing API resource rather than creating a new one.
func TestMigrateZoneSetting_OnOff(t *testing.T) {
	testCases := []struct {
		name       string
		version    string
		configFn   func(rnd, zoneID string) string // config for step 1 (create)
		v4ConfigFn func(rnd, zoneID string) string // config passed to MigrationV2TestStep (what tf-migrate transforms)
	}{
		{
			name:       "from_v4_latest",
			version:    acctest.GetLastV4Version(),
			configFn:   func(rnd, zoneID string) string { return fmt.Sprintf(v4OnOffConfig, rnd, zoneID) },
			v4ConfigFn: func(rnd, zoneID string) string { return fmt.Sprintf(v4OnOffConfig, rnd, zoneID) },
		},
		{
			name:       "from_v5",
			version:    currentProviderVersion,
			configFn:   func(rnd, zoneID string) string { return fmt.Sprintf(v5OnOffPlainConfig, rnd, zoneID) },
			v4ConfigFn: func(rnd, zoneID string) string { return fmt.Sprintf(v5OnOffConfig, rnd, zoneID) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			firstConfig := tc.configFn(rnd, zoneID)
			migrateConfig := tc.v4ConfigFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   firstConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: firstConfig,
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
					acctest.MigrationV2TestStepForZoneSetting(t, migrateConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"cloudflare_zone_setting."+rnd+"_http3",
							tfjsonpath.New("zone_id"),
							knownvalue.StringExact(zoneID),
						),
						statecheck.ExpectKnownValue(
							"cloudflare_zone_setting."+rnd+"_http3",
							tfjsonpath.New("setting_id"),
							knownvalue.StringExact("http3"),
						),
						statecheck.ExpectKnownValue(
							"cloudflare_zone_setting."+rnd+"_http3",
							tfjsonpath.New("value"),
							knownvalue.StringExact("on"),
						),
					}),
				},
			})
		})
	}
}

// TestMigrateZoneSetting_Number tests migration of an integer setting (browser_cache_ttl).
//
// v4: cloudflare_zone_settings_override with settings { browser_cache_ttl = 14400 }
// v5: cloudflare_zone_setting with setting_id = "browser_cache_ttl", value = 14400
func TestMigrateZoneSetting_Number(t *testing.T) {
	testCases := []struct {
		name       string
		version    string
		configFn   func(rnd, zoneID string) string
		v4ConfigFn func(rnd, zoneID string) string
	}{
		{
			name:       "from_v4_latest",
			version:    acctest.GetLastV4Version(),
			configFn:   func(rnd, zoneID string) string { return fmt.Sprintf(v4NumberConfig, rnd, zoneID) },
			v4ConfigFn: func(rnd, zoneID string) string { return fmt.Sprintf(v4NumberConfig, rnd, zoneID) },
		},
		{
			name:       "from_v5",
			version:    currentProviderVersion,
			configFn:   func(rnd, zoneID string) string { return fmt.Sprintf(v5NumberPlainConfig, rnd, zoneID) },
			v4ConfigFn: func(rnd, zoneID string) string { return fmt.Sprintf(v5NumberConfig, rnd, zoneID) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			firstConfig := tc.configFn(rnd, zoneID)
			migrateConfig := tc.v4ConfigFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   firstConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: firstConfig,
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
					acctest.MigrationV2TestStepForZoneSetting(t, migrateConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"cloudflare_zone_setting."+rnd+"_browser_cache_ttl",
							tfjsonpath.New("zone_id"),
							knownvalue.StringExact(zoneID),
						),
						statecheck.ExpectKnownValue(
							"cloudflare_zone_setting."+rnd+"_browser_cache_ttl",
							tfjsonpath.New("setting_id"),
							knownvalue.StringExact("browser_cache_ttl"),
						),
						statecheck.ExpectKnownValue(
							"cloudflare_zone_setting."+rnd+"_browser_cache_ttl",
							tfjsonpath.New("value"),
							knownvalue.Float64Exact(14400),
						),
					}),
				},
			})
		})
	}
}

// TestMigrateZoneSetting_SecurityHeader tests migration of the security_header nested block.
//
// v4: cloudflare_zone_settings_override with a security_header { ... } nested block
// v5: cloudflare_zone_setting with setting_id = "security_header" and a nested
//
//	value.strict_transport_security object (tf-migrate wraps the attributes)
func TestMigrateZoneSetting_SecurityHeader(t *testing.T) {
	testCases := []struct {
		name       string
		version    string
		configFn   func(rnd, zoneID string) string
		v4ConfigFn func(rnd, zoneID string) string
	}{
		{
			name:       "from_v4_latest",
			version:    acctest.GetLastV4Version(),
			configFn:   func(rnd, zoneID string) string { return fmt.Sprintf(v4SecurityHeaderConfig, rnd, zoneID) },
			v4ConfigFn: func(rnd, zoneID string) string { return fmt.Sprintf(v4SecurityHeaderConfig, rnd, zoneID) },
		},
		{
			name:       "from_v5",
			version:    currentProviderVersion,
			configFn:   func(rnd, zoneID string) string { return fmt.Sprintf(v5SecurityHeaderPlainConfig, rnd, zoneID) },
			v4ConfigFn: func(rnd, zoneID string) string { return fmt.Sprintf(v5SecurityHeaderConfig, rnd, zoneID) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			firstConfig := tc.configFn(rnd, zoneID)
			migrateConfig := tc.v4ConfigFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   firstConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: firstConfig,
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
					acctest.MigrationV2TestStepForZoneSetting(t, migrateConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"cloudflare_zone_setting."+rnd+"_security_header",
							tfjsonpath.New("zone_id"),
							knownvalue.StringExact(zoneID),
						),
						statecheck.ExpectKnownValue(
							"cloudflare_zone_setting."+rnd+"_security_header",
							tfjsonpath.New("setting_id"),
							knownvalue.StringExact("security_header"),
						),
						// Nested block becomes value.strict_transport_security object
						statecheck.ExpectKnownValue(
							"cloudflare_zone_setting."+rnd+"_security_header",
							tfjsonpath.New("value").AtMapKey("strict_transport_security").AtMapKey("enabled"),
							knownvalue.Bool(true),
						),
						statecheck.ExpectKnownValue(
							"cloudflare_zone_setting."+rnd+"_security_header",
							tfjsonpath.New("value").AtMapKey("strict_transport_security").AtMapKey("max_age"),
							knownvalue.Float64Exact(86400),
						),
						statecheck.ExpectKnownValue(
							"cloudflare_zone_setting."+rnd+"_security_header",
							tfjsonpath.New("value").AtMapKey("strict_transport_security").AtMapKey("include_subdomains"),
							knownvalue.Bool(true),
						),
					}),
				},
			})
		})
	}
}

// TestMigrateZoneSetting_Multiple tests the 1-to-N split: one cloudflare_zone_settings_override
// with multiple settings becomes multiple cloudflare_zone_setting resources.
//
// This is the core migration scenario. With import blocks generated by tf-migrate,
// Terraform adopts the existing API resources rather than creating new ones, so
// MigrationV2TestStep (empty plan) is correct — no AllowCreate needed.
func TestMigrateZoneSetting_Multiple(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	version := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(version)

	v4Config := fmt.Sprintf(v4MultipleConfig, rnd, zoneID)

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
						VersionConstraint: version,
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStepForZoneSetting(t, v4Config, tmpDir, version, sourceVer, targetVer, []statecheck.StateCheck{
				// http3 setting
				statecheck.ExpectKnownValue(
					"cloudflare_zone_setting."+rnd+"_http3",
					tfjsonpath.New("zone_id"),
					knownvalue.StringExact(zoneID),
				),
				statecheck.ExpectKnownValue(
					"cloudflare_zone_setting."+rnd+"_http3",
					tfjsonpath.New("setting_id"),
					knownvalue.StringExact("http3"),
				),
				statecheck.ExpectKnownValue(
					"cloudflare_zone_setting."+rnd+"_http3",
					tfjsonpath.New("value"),
					knownvalue.StringExact("on"),
				),
				// browser_cache_ttl setting (integer)
				statecheck.ExpectKnownValue(
					"cloudflare_zone_setting."+rnd+"_browser_cache_ttl",
					tfjsonpath.New("setting_id"),
					knownvalue.StringExact("browser_cache_ttl"),
				),
				statecheck.ExpectKnownValue(
					"cloudflare_zone_setting."+rnd+"_browser_cache_ttl",
					tfjsonpath.New("value"),
					knownvalue.Float64Exact(14400),
				),
				// min_tls_version setting
				statecheck.ExpectKnownValue(
					"cloudflare_zone_setting."+rnd+"_min_tls_version",
					tfjsonpath.New("setting_id"),
					knownvalue.StringExact("min_tls_version"),
				),
				statecheck.ExpectKnownValue(
					"cloudflare_zone_setting."+rnd+"_min_tls_version",
					tfjsonpath.New("value"),
					knownvalue.StringExact("1.2"),
				),
			}),
		},
	})
}
