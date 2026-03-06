package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Embed test config templates.
// The configs contain %[N]s placeholders for rnd, zoneID, cert, and key.

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_geo_restrictions.tf
var v4GeoRestrictionsConfig string

//go:embed testdata/v5_geo_restrictions.tf
var v5GeoRestrictionsConfig string

// skipIfNoCustomSSL skips the test unless CLOUDFLARE_CUSTOM_SSL_TEST=1.
// The test zone has a limited quota for custom SSL certificates.
func skipIfNoCustomSSL(t *testing.T) {
	t.Helper()
	if os.Getenv("CLOUDFLARE_CUSTOM_SSL_TEST") != "1" {
		t.Skip("Skipping custom SSL migration test. Set CLOUDFLARE_CUSTOM_SSL_TEST=1 to run.")
	}
}

// generateCertAndKey generates a short-lived certificate and private key for testing.
func generateCertAndKey(t *testing.T) (cert, key string) {
	t.Helper()
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	expiry := time.Now().Add(time.Hour * 1)
	c, k, err := utils.GenerateEphemeralCertAndKey([]string{domain}, expiry)
	if err != nil {
		t.Fatalf("Failed to generate certificate: %s", err)
	}
	return c, k
}

// TestMigrateCustomSSL_V4ToV5_Basic validates the core migration:
// custom_ssl_options block → flat top-level attributes.
//
// Covers:
//   - certificate, private_key, bundle_method, type (hoisted from block)
//   - status, hosts (computed fields preserved)
func TestMigrateCustomSSL_V4ToV5_Basic(t *testing.T) {
	skipIfNoCustomSSL(t)

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	cert, key := generateCertAndKey(t)

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, zoneID, cert, key)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, zoneID, cert, key)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_custom_ssl." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd)
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
					acctest.TestAccPreCheck_Credentials(t)
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						// Core ID preserved.
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
						// zone_id preserved.
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
						// Flat fields hoisted from custom_ssl_options block.
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bundle_method"), knownvalue.StringExact("force")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("legacy_custom")),
						// Computed fields populated from API.
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hosts"), knownvalue.NotNull()),
					}),
				},
			})
		})
	}
}

// TestMigrateCustomSSL_V4ToV5_WithGeoRestrictions validates that
// geo_restrictions is transformed from a plain string to { label = "..." }.
func TestMigrateCustomSSL_V4ToV5_WithGeoRestrictions(t *testing.T) {
	skipIfNoCustomSSL(t)

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	cert, key := generateCertAndKey(t)

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd string) string {
				return fmt.Sprintf(v4GeoRestrictionsConfig, rnd, zoneID, cert, key)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd string) string {
				return fmt.Sprintf(v5GeoRestrictionsConfig, rnd, zoneID, cert, key)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_custom_ssl." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd)
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
					acctest.TestAccPreCheck_Credentials(t)
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						// Core fields.
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
						// geo_restrictions transformed to nested object.
						statecheck.ExpectKnownValue(resourceName,
							tfjsonpath.New("geo_restrictions").AtMapKey("label"),
							knownvalue.StringExact("us"),
						),
						// Other hoisted fields.
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bundle_method"), knownvalue.StringExact("force")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("legacy_custom")),
					}),
				},
			})
		})
	}
}
