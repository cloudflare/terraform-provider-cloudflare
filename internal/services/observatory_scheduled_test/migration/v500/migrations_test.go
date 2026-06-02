package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release

	// useLocalV5Provider: Toggle to use local v5 build or download from registry
	//   true  = Use local build (ProtoV6ProviderFactories) for from_v5 tests
	//   false = Download from registry (ExternalProviders) for from_v5 tests
	useLocalV5Provider = true
)

// Note: Tests are intentionally minimal due to per-zone API throttling limits on this resource.

// Migration Test Configuration
//
// Version is read from LAST_V4_VERSION environment variable (set in .github/workflows/migration-tests.yml)
// - Last stable v4 release: default 4.52.5
// - Current v5 release: auto-updates with releases (internal.PackageVersion)

// Embed migration test configuration files
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

// TestMigrateObservatoryScheduledTest_V4ToV5_Basic tests basic migration with all v4 fields
func TestMigrateObservatoryScheduledTest_V4ToV5_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(zoneID, domain string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: getLastV4Version(),
			configFn: func(zoneID, domain string) string {
				return fmt.Sprintf(v4BasicConfig, zoneID, domain)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(zoneID, domain string) string {
				return fmt.Sprintf(v5BasicConfig, zoneID, domain)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			url := normalizeObservatoryURL(domain)
			resourceName := "cloudflare_observatory_scheduled_test.test"
			tmpDir := t.TempDir()
			testConfig := tc.configFn(zoneID, url)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			// Determine provider source based on version and toggle
			var firstStep resource.TestStep
			if tc.version == currentProviderVersion && useLocalV5Provider {
				// Use local v5 build
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				// Download from registry
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
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact(url)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("region"), knownvalue.StringExact("us-central1")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("frequency"), knownvalue.StringExact("DAILY")),
					}),
				},
			})
		})
	}
}

// TestMigrateObservatoryScheduledTest_V4ToV5_URLWithPath and V4ToV5_WEEKLY and V4ToV5_AllRegions removed.
// The migration is a pure pass-through (all fields are direct string copies); varying region/frequency
// values exercises identical code paths, so a single test case is sufficient.

// getLastV4Version returns the last v4 version from environment or default
func getLastV4Version() string {
	if v := os.Getenv("LAST_V4_VERSION"); v != "" {
		return v
	}
	return "4.52.5" // Default last v4 release
}

func normalizeObservatoryURL(input string) string {
	value := strings.TrimSpace(input)
	value = strings.TrimPrefix(value, "https://")
	value = strings.TrimPrefix(value, "http://")
	value = strings.TrimSuffix(value, "/")
	if value == "" {
		return "/"
	}
	return value + "/"
}
