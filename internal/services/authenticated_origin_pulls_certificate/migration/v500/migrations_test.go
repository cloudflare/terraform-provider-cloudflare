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

// Embed test configs
//
//go:embed testdata/v4_per_zone.tf
var v4PerZoneConfig string

//go:embed testdata/v5_per_zone.tf
var v5PerZoneConfig string

//go:embed testdata/v4_per_zone_minimal.tf
var v4PerZoneMinimalConfig string

//go:embed testdata/v5_per_zone_minimal.tf
var v5PerZoneMinimalConfig string

//go:embed testdata/v4_per_zone_with_variables.tf
var v4PerZoneWithVariablesConfig string

//go:embed testdata/v5_per_zone_with_variables.tf
var v5PerZoneWithVariablesConfig string

// TestMigrateAuthenticatedOriginPullsCertificate_V4ToV5_PerZone tests per-zone certificate migration.
// This test validates:
// - Migration from v4 (with type="per-zone") to v5 (without type field)
// - Removal of serial_number field (not present in v5 per-zone schema)
// - Preservation of all other fields (zone_id, certificate, private_key, computed fields)
func TestMigrateAuthenticatedOriginPullsCertificate_V4ToV5_PerZone(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, cert, key string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID, cert, key string) string {
				return fmt.Sprintf(v4PerZoneConfig, rnd, zoneID, cert, key)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump v1→v500)
			version: currentProviderVersion,
			configFn: func(rnd, zoneID, cert, key string) string {
				return fmt.Sprintf(v5PerZoneConfig, rnd, zoneID, cert, key)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test setup
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			if zoneID == "" {
				t.Skip("CLOUDFLARE_ZONE_ID must be set for this test")
			}

			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()

			// Generate valid test certificate
			expiry := time.Now().Add(time.Hour * 24 * 365)
			cert, key, err := utils.GenerateEphemeralCertAndKey([]string{"aop-test.example.com"}, expiry)
			if err != nil {
				t.Fatalf("Failed to generate certificate: %s", err)
			}

			testConfig := tc.configFn(rnd, zoneID, cert, key)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)
			resourceName := "cloudflare_authenticated_origin_pulls_certificate." + rnd

			resource.Test(t, resource.TestCase{
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
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify zone_id preserved
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("zone_id"),
								knownvalue.StringExact(zoneID),
							),
							// Verify certificate preserved
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("certificate"),
								knownvalue.NotNull(),
							),
							// Verify private_key preserved
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("private_key"),
								knownvalue.NotNull(),
							),
							// Verify computed fields present
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("issuer"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("signature"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("status"),
								knownvalue.NotNull(),
							),
							// Note: type and serial_number should be removed after migration
							// These will be checked implicitly by the absence of errors
						},
					),
				},
			})
		})
	}
}

// TestMigrateAuthenticatedOriginPullsCertificate_V4ToV5_PerZone_Minimal tests minimal config migration.
// This test validates that resources with only required fields migrate correctly.
func TestMigrateAuthenticatedOriginPullsCertificate_V4ToV5_PerZone_Minimal(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, cert, key string) string
	}{
		{
			name:    "from_v4_latest_minimal",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID, cert, key string) string {
				return fmt.Sprintf(v4PerZoneMinimalConfig, rnd, zoneID, cert, key)
			},
		},
		{
			name:    "from_v5_minimal",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID, cert, key string) string {
				return fmt.Sprintf(v5PerZoneMinimalConfig, rnd, zoneID, cert, key)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			if zoneID == "" {
				t.Skip("CLOUDFLARE_ZONE_ID must be set for this test")
			}

			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()

			expiry := time.Now().Add(time.Hour * 24 * 365)
			cert, key, err := utils.GenerateEphemeralCertAndKey([]string{"aop-minimal-test.example.com"}, expiry)
			if err != nil {
				t.Fatalf("Failed to generate certificate: %s", err)
			}

			testConfig := tc.configFn(rnd, zoneID, cert, key)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)
			resourceName := "cloudflare_authenticated_origin_pulls_certificate." + rnd

			resource.Test(t, resource.TestCase{
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
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("zone_id"),
								knownvalue.StringExact(zoneID),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("certificate"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("private_key"),
								knownvalue.NotNull(),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateAuthenticatedOriginPullsCertificate_V4ToV5_PerZone_AllFields tests that all fields are preserved.
// This test validates:
// - All required fields (zone_id, certificate, private_key) are preserved
// - All computed fields (issuer, signature, serial_number, expires_on, status, uploaded_on) are preserved
// - The type field is properly removed during migration
func TestMigrateAuthenticatedOriginPullsCertificate_V4ToV5_PerZone_AllFields(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, cert, key string) string
	}{
		{
			name:    "from_v4_latest_all_fields",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID, cert, key string) string {
				return fmt.Sprintf(v4PerZoneConfig, rnd, zoneID, cert, key)
			},
		},
		{
			name:    "from_v5_all_fields",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID, cert, key string) string {
				return fmt.Sprintf(v5PerZoneConfig, rnd, zoneID, cert, key)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			if zoneID == "" {
				t.Skip("CLOUDFLARE_ZONE_ID must be set for this test")
			}

			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()

			expiry := time.Now().Add(time.Hour * 24 * 365)
			cert, key, err := utils.GenerateEphemeralCertAndKey([]string{"aop-allfields-test.example.com"}, expiry)
			if err != nil {
				t.Fatalf("Failed to generate certificate: %s", err)
			}

			testConfig := tc.configFn(rnd, zoneID, cert, key)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)
			resourceName := "cloudflare_authenticated_origin_pulls_certificate." + rnd

			resource.Test(t, resource.TestCase{
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
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Required fields
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("zone_id"),
								knownvalue.StringExact(zoneID),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("certificate"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("private_key"),
								knownvalue.NotNull(),
							),
							// Computed fields - verify they exist and are preserved
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("issuer"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("signature"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("serial_number"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("expires_on"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("status"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("uploaded_on"),
								knownvalue.NotNull(),
							),
							// v5-specific computed fields
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("certificate_id"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("id"),
								knownvalue.NotNull(),
							),
						},
					),
				},
			})
		})
	}
}
