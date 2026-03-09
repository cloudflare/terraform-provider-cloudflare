package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"
	"time"

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

// Migration Scope
//
// The v4 cloudflare_authenticated_origin_pulls resource had 3 modes:
// 1. Global AOP: zone_id + enabled (NO hostname, NO cert)
//    → Migrates to cloudflare_authenticated_origin_pulls_settings
// 2. Per-Zone AOP: zone_id + cert + enabled (NO hostname)
//    → Migrates to cloudflare_authenticated_origin_pulls_settings
// 3. Per-Hostname AOP: zone_id + hostname + cert + enabled
//    → Migrates to cloudflare_authenticated_origin_pulls (THIS RESOURCE)
//
// This test suite ONLY covers mode 3 (Per-Hostname AOP).
// Modes 1 and 2 are tested in authenticated_origin_pulls_settings/migration/v500/migrations_test.go

// Embed test configs
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_disabled.tf
var v4DisabledConfig string

//go:embed testdata/v5_disabled.tf
var v5DisabledConfig string

//go:embed testdata/v4_multiple.tf
var v4MultipleConfig string

//go:embed testdata/v5_multiple.tf
var v5MultipleConfig string

// TestMigrateAuthenticatedOriginPulls_V4ToV5_Basic tests basic Per-Hostname AOP migration with DUAL test cases
//
// This test validates:
// - Flat v4 structure → Nested v5 config structure
// - Field rename: authenticated_origin_pulls_certificate → config[0].cert_id
// - Field move: hostname → config[0].hostname
// - Field move: enabled → config[0].enabled
func TestMigrateAuthenticatedOriginPulls_V4ToV5_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(certRnd, zoneID, cert, key, rnd, zoneID2, hostname, certRnd2 string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(certRnd, zoneID, cert, key, rnd, zoneID2, hostname, certRnd2 string) string {
				return fmt.Sprintf(v4BasicConfig, certRnd, zoneID, cert, key, rnd, zoneID2, hostname, certRnd2)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(certRnd, zoneID, cert, key, rnd, zoneID2, hostname, certRnd2 string) string {
				return fmt.Sprintf(v5BasicConfig, certRnd, zoneID, cert, key, rnd, zoneID2, hostname, certRnd2)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test setup
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			if domain == "" {
				t.Skip("CLOUDFLARE_DOMAIN must be set for this test")
			}
			rnd := utils.GenerateRandomResourceName()
			certRnd := rnd + "_hostname_cert"
			hostname := "tf-test-" + rnd + "." + domain

			// Generate valid test certificate
			expiry := time.Now().Add(time.Hour * 24 * 365)
			cert, key, err := utils.GenerateEphemeralCertAndKey([]string{hostname}, expiry)
			if err != nil {
				t.Fatalf("Failed to generate certificate: %s", err)
			}

			tmpDir := t.TempDir()
			testConfig := tc.configFn(certRnd, zoneID, cert, key, rnd, zoneID, hostname, certRnd)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			// For v5 tests, use local provider; for v4 tests, use external provider
			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				// Use local v5 provider (will create state)
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
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify critical identifier
							statecheck.ExpectKnownValue(
								"cloudflare_authenticated_origin_pulls."+rnd,
								tfjsonpath.New("zone_id"),
								knownvalue.StringExact(zoneID),
							),
							// Verify nested config structure (key transformation)
							statecheck.ExpectKnownValue(
								"cloudflare_authenticated_origin_pulls."+rnd,
								tfjsonpath.New("config").AtSliceIndex(0).AtMapKey("hostname"),
								knownvalue.StringExact(hostname),
							),
							// Note: cert_id is dynamically generated by the API, so we just check it's not null
							statecheck.ExpectKnownValue(
								"cloudflare_authenticated_origin_pulls."+rnd,
								tfjsonpath.New("config").AtSliceIndex(0).AtMapKey("cert_id"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_authenticated_origin_pulls."+rnd,
								tfjsonpath.New("config").AtSliceIndex(0).AtMapKey("enabled"),
								knownvalue.Bool(true),
							),
							// Verify top-level computed fields are populated from API
							statecheck.ExpectKnownValue(
								"cloudflare_authenticated_origin_pulls."+rnd,
								tfjsonpath.New("hostname"),
								knownvalue.StringExact(hostname),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_authenticated_origin_pulls."+rnd,
								tfjsonpath.New("cert_id"),
								knownvalue.NotNull(),
							),
							// Note: We don't validate other computed fields like cert_status, created_at, etc.
							// These are populated by the API and may vary
						},
					),
				},
			})
		})
	}
}

// TestMigrateAuthenticatedOriginPulls_V4ToV5_Disabled tests Per-Hostname AOP migration with disabled state
//
// This test validates:
// - enabled = false is correctly preserved through migration
// - Nested config structure is created correctly even when disabled
// - Top-level computed fields (enabled, hostname) reflect disabled state
func TestMigrateAuthenticatedOriginPulls_V4ToV5_Disabled(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(certRnd, zoneID, cert, key, rnd, zoneID2, hostname, certRnd2 string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(certRnd, zoneID, cert, key, rnd, zoneID2, hostname, certRnd2 string) string {
				return fmt.Sprintf(v4DisabledConfig, certRnd, zoneID, cert, key, rnd, zoneID2, hostname, certRnd2)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(certRnd, zoneID, cert, key, rnd, zoneID2, hostname, certRnd2 string) string {
				return fmt.Sprintf(v5DisabledConfig, certRnd, zoneID, cert, key, rnd, zoneID2, hostname, certRnd2)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			if domain == "" {
				t.Skip("CLOUDFLARE_DOMAIN must be set for this test")
			}
			rnd := utils.GenerateRandomResourceName()
			certRnd := rnd + "_hostname_cert"
			hostname := "tf-test-" + rnd + "." + domain

			expiry := time.Now().Add(time.Hour * 24 * 365)
			cert, key, err := utils.GenerateEphemeralCertAndKey([]string{hostname}, expiry)
			if err != nil {
				t.Fatalf("Failed to generate certificate: %s", err)
			}

			tmpDir := t.TempDir()
			testConfig := tc.configFn(certRnd, zoneID, cert, key, rnd, zoneID, hostname, certRnd)
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
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								"cloudflare_authenticated_origin_pulls."+rnd,
								tfjsonpath.New("zone_id"),
								knownvalue.StringExact(zoneID),
							),
							// Verify nested config with disabled state
							statecheck.ExpectKnownValue(
								"cloudflare_authenticated_origin_pulls."+rnd,
								tfjsonpath.New("config").AtSliceIndex(0).AtMapKey("hostname"),
								knownvalue.StringExact(hostname),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_authenticated_origin_pulls."+rnd,
								tfjsonpath.New("config").AtSliceIndex(0).AtMapKey("cert_id"),
								knownvalue.NotNull(),
							),
							// Critical: verify enabled = false is preserved
							statecheck.ExpectKnownValue(
								"cloudflare_authenticated_origin_pulls."+rnd,
								tfjsonpath.New("config").AtSliceIndex(0).AtMapKey("enabled"),
								knownvalue.Bool(false),
							),
							// Verify top-level computed fields
							statecheck.ExpectKnownValue(
								"cloudflare_authenticated_origin_pulls."+rnd,
								tfjsonpath.New("hostname"),
								knownvalue.StringExact(hostname),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_authenticated_origin_pulls."+rnd,
								tfjsonpath.New("enabled"),
								knownvalue.Bool(false),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateAuthenticatedOriginPulls_V4ToV5_Multiple tests migration of multiple Per-Hostname AOP resources
//
// This test validates:
// - Multiple hostname associations in same zone migrate correctly
// - Each resource maintains its own state independently
// - Mix of enabled/disabled states is preserved
// - No cross-contamination between resources
func TestMigrateAuthenticatedOriginPulls_V4ToV5_Multiple(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(certRnd1, zoneID1, cert1, key1, certRnd2, zoneID2, cert2, key2, rnd1, zoneID3, hostname1, certRnd1Ref, rnd2, zoneID4, hostname2, certRnd2Ref string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(certRnd1, zoneID1, cert1, key1, certRnd2, zoneID2, cert2, key2, rnd1, zoneID3, hostname1, certRnd1Ref, rnd2, zoneID4, hostname2, certRnd2Ref string) string {
				return fmt.Sprintf(v4MultipleConfig, certRnd1, zoneID1, cert1, key1, certRnd2, zoneID2, cert2, key2, rnd1, zoneID3, hostname1, certRnd1Ref, rnd2, zoneID4, hostname2, certRnd2Ref)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(certRnd1, zoneID1, cert1, key1, certRnd2, zoneID2, cert2, key2, rnd1, zoneID3, hostname1, certRnd1Ref, rnd2, zoneID4, hostname2, certRnd2Ref string) string {
				return fmt.Sprintf(v5MultipleConfig, certRnd1, zoneID1, cert1, key1, certRnd2, zoneID2, cert2, key2, rnd1, zoneID3, hostname1, certRnd1Ref, rnd2, zoneID4, hostname2, certRnd2Ref)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			if domain == "" {
				t.Skip("CLOUDFLARE_DOMAIN must be set for this test")
			}
			rnd1 := utils.GenerateRandomResourceName()
			rnd2 := utils.GenerateRandomResourceName()
			certRnd1 := rnd1 + "_hostname_cert"
			certRnd2 := rnd2 + "_hostname_cert"
			hostname1 := "tf-test-" + rnd1 + "." + domain
			hostname2 := "tf-test-" + rnd2 + "." + domain

			// Generate certificates for both hostnames
			expiry := time.Now().Add(time.Hour * 24 * 365)
			cert1, key1, err := utils.GenerateEphemeralCertAndKey([]string{hostname1}, expiry)
			if err != nil {
				t.Fatalf("Failed to generate certificate 1: %s", err)
			}
			cert2, key2, err := utils.GenerateEphemeralCertAndKey([]string{hostname2}, expiry)
			if err != nil {
				t.Fatalf("Failed to generate certificate 2: %s", err)
			}

			tmpDir := t.TempDir()
			testConfig := tc.configFn(certRnd1, zoneID, cert1, key1, certRnd2, zoneID, cert2, key2, rnd1, zoneID, hostname1, certRnd1, rnd2, zoneID, hostname2, certRnd2)
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
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify first resource (enabled = true)
							statecheck.ExpectKnownValue(
								"cloudflare_authenticated_origin_pulls."+rnd1,
								tfjsonpath.New("zone_id"),
								knownvalue.StringExact(zoneID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_authenticated_origin_pulls."+rnd1,
								tfjsonpath.New("config").AtSliceIndex(0).AtMapKey("hostname"),
								knownvalue.StringExact(hostname1),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_authenticated_origin_pulls."+rnd1,
								tfjsonpath.New("config").AtSliceIndex(0).AtMapKey("enabled"),
								knownvalue.Bool(true),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_authenticated_origin_pulls."+rnd1,
								tfjsonpath.New("hostname"),
								knownvalue.StringExact(hostname1),
							),
							// Verify second resource (enabled = false)
							statecheck.ExpectKnownValue(
								"cloudflare_authenticated_origin_pulls."+rnd2,
								tfjsonpath.New("zone_id"),
								knownvalue.StringExact(zoneID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_authenticated_origin_pulls."+rnd2,
								tfjsonpath.New("config").AtSliceIndex(0).AtMapKey("hostname"),
								knownvalue.StringExact(hostname2),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_authenticated_origin_pulls."+rnd2,
								tfjsonpath.New("config").AtSliceIndex(0).AtMapKey("enabled"),
								knownvalue.Bool(false),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_authenticated_origin_pulls."+rnd2,
								tfjsonpath.New("hostname"),
								knownvalue.StringExact(hostname2),
							),
						},
					),
				},
			})
		})
	}
}
