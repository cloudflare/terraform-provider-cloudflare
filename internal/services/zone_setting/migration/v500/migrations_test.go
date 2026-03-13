package v500_test

import (
	_ "embed"
	"fmt"
	"math/big"
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

// Migration Test Configuration
//
// Version is read from LAST_V4_VERSION environment variable (set in .github/workflows/migration-tests.yml)
// - Last stable v4 release: default 4.52.5
// - Current v5 release: auto-updates with releases (internal.PackageVersion)
//
// Based on breaking changes analysis:
// - cloudflare_zone_settings_override (v4) → N × cloudflare_zone_setting (v5)
// - This is a one-to-many transformation: tf-migrate splits one v4 resource into multiple v5 resources
// - The v4 state is deleted; v5 resources are created via terraform apply after migration
// - Key changes: settings block → individual setting_id resources, zero_rtt → 0rtt, security_header wrapping
//
// Test Strategy:
// - v4→v5 tests: Create with v4 provider, run tf-migrate, verify each generated v5 resource
// - v5→v5 tests: Create with v5 provider, verify idempotency (no migration needed)
//
// TF_MIG_TEST must be set to "true" to run these tests (they require TF_ACC and real credentials).

// Embed migration test configuration files

//go:embed testdata/v4_http3.tf
var v4HTTP3Config string

//go:embed testdata/v5_http3.tf
var v5HTTP3Config string

//go:embed testdata/v4_tls_min.tf
var v4TLSMinConfig string

//go:embed testdata/v5_tls_min.tf
var v5TLSMinConfig string

//go:embed testdata/v4_security_header.tf
var v4SecurityHeaderConfig string

//go:embed testdata/v5_security_header.tf
var v5SecurityHeaderConfig string

//go:embed testdata/v4_multiple.tf
var v4MultipleConfig string

//go:embed testdata/v5_multiple.tf
var v5MultipleConfig string

//go:embed testdata/v4_zero_rtt.tf
var v4ZeroRTTConfig string

//go:embed testdata/v5_zero_rtt.tf
var v5ZeroRTTConfig string

// skipUnlessMigTest skips the test unless TF_MIG_TEST is set.
// These tests require real Cloudflare credentials and the tf-migrate binary.
func skipUnlessMigTest(t *testing.T) {
	t.Helper()
	if os.Getenv("TF_MIG_TEST") == "" {
		t.Skip("Skipping migration test: TF_MIG_TEST not set. Set TF_MIG_TEST=true to run.")
	}
}

// TestMigrateZoneSettingHTTP3 tests migration of a single on/off setting (http3) from v4 to v5.
//
// v4: cloudflare_zone_settings_override with settings { http3 = "on" }
// v5: cloudflare_zone_setting with setting_id = "http3", value = "on"
//
// This is the simplest migration case: one setting → one resource.
func TestMigrateZoneSettingHTTP3(t *testing.T) {
	skipUnlessMigTest(t)

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	legacyProviderVersion := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(legacyProviderVersion)

	v4Config := fmt.Sprintf(v4HTTP3Config, rnd, zoneID)
	resourceName := fmt.Sprintf("cloudflare_zone_setting.%s_http3", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: append(
			[]resource.TestStep{
				{
					// Step 1: Create with v4 provider (zone_settings_override)
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: legacyProviderVersion,
						},
					},
					Config: v4Config,
				},
			},
			// Step 2+: Run tf-migrate, apply with v5 provider, verify state
			// MigrationV2TestStepAllowCreate is used because the v4 state is deleted
			// and v5 resources are created fresh via terraform apply.
			acctest.MigrationV2TestStepAllowCreate(t, v4Config, tmpDir, legacyProviderVersion, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("setting_id"), knownvalue.StringExact("http3")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
			})...,
		),
	})
}

// TestMigrateZoneSettingHTTP3FromV5 tests v5→v5 idempotency for http3.
//
// Creates the resource with the current v5 provider and verifies no migration is needed.
func TestMigrateZoneSettingHTTP3FromV5(t *testing.T) {
	skipUnlessMigTest(t)

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	v5Config := fmt.Sprintf(v5HTTP3Config, rnd, zoneID)
	resourceName := fmt.Sprintf("cloudflare_zone_setting.%s_http3", rnd)
	sourceVer, targetVer := acctest.InferMigrationVersions(currentProviderVersion)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with current v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   v5Config,
			},
			// Step 2: Run migration (no-op for v5→v5), verify state unchanged
			acctest.MigrationV2TestStep(t, v5Config, tmpDir, currentProviderVersion, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("setting_id"), knownvalue.StringExact("http3")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
			}),
		},
	})
}

// TestMigrateZoneSettingTLSMin tests migration of min_tls_version from v4 to v5.
//
// v4: cloudflare_zone_settings_override with settings { min_tls_version = "1.2" }
// v5: cloudflare_zone_setting with setting_id = "min_tls_version", value = "1.2"
func TestMigrateZoneSettingTLSMin(t *testing.T) {
	skipUnlessMigTest(t)

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	legacyProviderVersion := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(legacyProviderVersion)

	v4Config := fmt.Sprintf(v4TLSMinConfig, rnd, zoneID)
	resourceName := fmt.Sprintf("cloudflare_zone_setting.%s_min_tls_version", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: append(
			[]resource.TestStep{
				{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: legacyProviderVersion,
						},
					},
					Config: v4Config,
				},
			},
			acctest.MigrationV2TestStepAllowCreate(t, v4Config, tmpDir, legacyProviderVersion, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("setting_id"), knownvalue.StringExact("min_tls_version")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("1.2")),
			})...,
		),
	})
}

// TestMigrateZoneSettingSecurityHeader tests migration of the security_header nested block.
//
// v4: cloudflare_zone_settings_override with settings { security_header { ... } }
// v5: cloudflare_zone_setting with setting_id = "security_header", value = { strict_transport_security = { ... } }
//
// This is a special case: the v4 security_header block is wrapped in strict_transport_security
// during migration to match the v5 API structure.
func TestMigrateZoneSettingSecurityHeader(t *testing.T) {
	skipUnlessMigTest(t)

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	legacyProviderVersion := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(legacyProviderVersion)

	v4Config := fmt.Sprintf(v4SecurityHeaderConfig, rnd, zoneID)
	resourceName := fmt.Sprintf("cloudflare_zone_setting.%s_security_header", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: append(
			[]resource.TestStep{
				{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: legacyProviderVersion,
						},
					},
					Config: v4Config,
				},
			},
			acctest.MigrationV2TestStepAllowCreate(t, v4Config, tmpDir, legacyProviderVersion, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("setting_id"), knownvalue.StringExact("security_header")),
				// Verify the strict_transport_security wrapping
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("value").AtMapKey("strict_transport_security").AtMapKey("enabled"),
					knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("value").AtMapKey("strict_transport_security").AtMapKey("max_age"),
					knownvalue.NumberExact(big.NewFloat(86400))),
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("value").AtMapKey("strict_transport_security").AtMapKey("include_subdomains"),
					knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("value").AtMapKey("strict_transport_security").AtMapKey("preload"),
					knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("value").AtMapKey("strict_transport_security").AtMapKey("nosniff"),
					knownvalue.Bool(false)),
			})...,
		),
	})
}

// TestMigrateZoneSettingZeroRTT tests migration of zero_rtt → 0rtt name mapping.
//
// v4: cloudflare_zone_settings_override with settings { zero_rtt = "on" }
// v5: cloudflare_zone_setting with setting_id = "0rtt", value = "on"
//
// The v4 provider used "zero_rtt" but the API expects "0rtt".
// tf-migrate handles this name mapping automatically.
func TestMigrateZoneSettingZeroRTT(t *testing.T) {
	skipUnlessMigTest(t)

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	legacyProviderVersion := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(legacyProviderVersion)

	v4Config := fmt.Sprintf(v4ZeroRTTConfig, rnd, zoneID)
	// Resource name uses the v4 attribute name (zero_rtt), setting_id uses the API name (0rtt)
	resourceName := fmt.Sprintf("cloudflare_zone_setting.%s_zero_rtt", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: append(
			[]resource.TestStep{
				{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: legacyProviderVersion,
						},
					},
					Config: v4Config,
				},
			},
			acctest.MigrationV2TestStepAllowCreate(t, v4Config, tmpDir, legacyProviderVersion, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				// setting_id must be "0rtt" (API name), not "zero_rtt" (v4 name)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("setting_id"), knownvalue.StringExact("0rtt")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
			})...,
		),
	})
}

// TestMigrateZoneSettingMultiple tests migration of multiple settings from one v4 resource.
//
// v4: cloudflare_zone_settings_override with settings { http3 = "on", min_tls_version = "1.2", brotli = "on" }
// v5: three separate cloudflare_zone_setting resources (one per setting)
//
// This is the core one-to-many transformation: one v4 resource → N v5 resources.
func TestMigrateZoneSettingMultiple(t *testing.T) {
	skipUnlessMigTest(t)

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	legacyProviderVersion := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(legacyProviderVersion)

	v4Config := fmt.Sprintf(v4MultipleConfig, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: append(
			[]resource.TestStep{
				{
					// Step 1: Create one v4 zone_settings_override with multiple settings
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: legacyProviderVersion,
						},
					},
					Config: v4Config,
				},
			},
			// Step 2+: Run tf-migrate (splits into 3 resources), apply, verify all three
			acctest.MigrationV2TestStepAllowCreate(t, v4Config, tmpDir, legacyProviderVersion, sourceVer, targetVer, []statecheck.StateCheck{
				// brotli resource
				statecheck.ExpectKnownValue(
					fmt.Sprintf("cloudflare_zone_setting.%s_brotli", rnd),
					tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(
					fmt.Sprintf("cloudflare_zone_setting.%s_brotli", rnd),
					tfjsonpath.New("setting_id"), knownvalue.StringExact("brotli")),
				statecheck.ExpectKnownValue(
					fmt.Sprintf("cloudflare_zone_setting.%s_brotli", rnd),
					tfjsonpath.New("value"), knownvalue.StringExact("on")),

				// http3 resource
				statecheck.ExpectKnownValue(
					fmt.Sprintf("cloudflare_zone_setting.%s_http3", rnd),
					tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(
					fmt.Sprintf("cloudflare_zone_setting.%s_http3", rnd),
					tfjsonpath.New("setting_id"), knownvalue.StringExact("http3")),
				statecheck.ExpectKnownValue(
					fmt.Sprintf("cloudflare_zone_setting.%s_http3", rnd),
					tfjsonpath.New("value"), knownvalue.StringExact("on")),

				// min_tls_version resource
				statecheck.ExpectKnownValue(
					fmt.Sprintf("cloudflare_zone_setting.%s_min_tls_version", rnd),
					tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(
					fmt.Sprintf("cloudflare_zone_setting.%s_min_tls_version", rnd),
					tfjsonpath.New("setting_id"), knownvalue.StringExact("min_tls_version")),
				statecheck.ExpectKnownValue(
					fmt.Sprintf("cloudflare_zone_setting.%s_min_tls_version", rnd),
					tfjsonpath.New("value"), knownvalue.StringExact("1.2")),
			})...,
		),
	})
}
