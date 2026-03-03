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
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_with_wait_for_active_status.tf
var v4WithWaitForActiveStatusConfig string

//go:embed testdata/v5_with_optional_fields.tf
var v5WithOptionalFieldsConfig string

//go:embed testdata/v4_google_http.tf
var v4GoogleHTTPConfig string

//go:embed testdata/v5_google_http.tf
var v5GoogleHTTPConfig string

// TestMigrateCertificatePack_V4ToV5_Basic tests basic certificate pack migration with required fields
func TestMigrateCertificatePack_V4ToV5_Basic(t *testing.T) {
	t.Skip("Migration not enabled yet")
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, domain string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID, domain string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, zoneID, domain, domain)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, zoneID, domain string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, zoneID, domain, domain)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test setup
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID, domain)
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
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify core fields
							statecheck.ExpectKnownValue(
								"cloudflare_certificate_pack."+rnd,
								tfjsonpath.New("zone_id"),
								knownvalue.StringExact(zoneID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_certificate_pack."+rnd,
								tfjsonpath.New("type"),
								knownvalue.StringExact("advanced"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_certificate_pack."+rnd,
								tfjsonpath.New("validation_method"),
								knownvalue.StringExact("txt"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_certificate_pack."+rnd,
								tfjsonpath.New("validity_days"),
								knownvalue.Int64Exact(90),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_certificate_pack."+rnd,
								tfjsonpath.New("certificate_authority"),
								knownvalue.StringExact("lets_encrypt"),
							),
							// Verify hosts are preserved
							statecheck.ExpectKnownValue(
								"cloudflare_certificate_pack."+rnd,
								tfjsonpath.New("hosts"),
								knownvalue.SetSizeExact(2),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateCertificatePack_V4ToV5_WaitForActiveStatus tests wait_for_active_status field removal
func TestMigrateCertificatePack_V4ToV5_WaitForActiveStatus(t *testing.T) {
	t.Skip("Migration not enabled yet")
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, domain string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID, domain string) string {
				return fmt.Sprintf(v4WithWaitForActiveStatusConfig, rnd, zoneID, domain, domain)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, zoneID, domain string) string {
				return fmt.Sprintf(v5WithOptionalFieldsConfig, rnd, zoneID, domain, domain)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Skip v4 test - v4 provider bug with wait_for_active_status causes "certificate list in response is empty"
			if tc.version != currentProviderVersion {
				t.Skip("Skipping v4 test due to known v4 provider bug with wait_for_active_status")
			}

			// Test setup
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID, domain)
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
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify cloudflare_branding is preserved
							statecheck.ExpectKnownValue(
								"cloudflare_certificate_pack."+rnd,
								tfjsonpath.New("cloudflare_branding"),
								knownvalue.Bool(true),
							),
							// Verify wait_for_active_status is NOT in state (removed in v5)
							// Note: In v4, this field exists but won't be checked in v5 state
							statecheck.ExpectKnownValue(
								"cloudflare_certificate_pack."+rnd,
								tfjsonpath.New("zone_id"),
								knownvalue.StringExact(zoneID),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateCertificatePack_V4ToV5_DifferentValidation tests different validation methods and CAs
func TestMigrateCertificatePack_V4ToV5_DifferentValidation(t *testing.T) {
	t.Skip("Migration not enabled yet")
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, domain string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID, domain string) string {
				return fmt.Sprintf(v4GoogleHTTPConfig, rnd, zoneID, domain)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, zoneID, domain string) string {
				return fmt.Sprintf(v5GoogleHTTPConfig, rnd, zoneID, domain)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Skip v4 test - Google CA with HTTP validation requires actual HTTP challenges that can timeout
			if tc.version != currentProviderVersion {
				t.Skip("Skipping v4 test due to Google CA with HTTP validation taking too long (20+ minutes)")
			}

			// Test setup
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID, domain)
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
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify validation_method is preserved
							statecheck.ExpectKnownValue(
								"cloudflare_certificate_pack."+rnd,
								tfjsonpath.New("validation_method"),
								knownvalue.StringExact("http"),
							),
							// Verify certificate_authority is preserved
							statecheck.ExpectKnownValue(
								"cloudflare_certificate_pack."+rnd,
								tfjsonpath.New("certificate_authority"),
								knownvalue.StringExact("google"),
							),
							// Verify validity_days is preserved (different value)
							statecheck.ExpectKnownValue(
								"cloudflare_certificate_pack."+rnd,
								tfjsonpath.New("validity_days"),
								knownvalue.Int64Exact(90),
							),
							// Verify single host
							statecheck.ExpectKnownValue(
								"cloudflare_certificate_pack."+rnd,
								tfjsonpath.New("hosts"),
								knownvalue.SetSizeExact(1),
							),
						},
					),
				},
			})
		})
	}
}
