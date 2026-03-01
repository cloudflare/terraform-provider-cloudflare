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
//go:embed testdata/v4_enabled_true.tf
var v4EnabledTrueConfig string

//go:embed testdata/v5_enabled_true.tf
var v5EnabledTrueConfig string

//go:embed testdata/v4_enabled_false.tf
var v4EnabledFalseConfig string

//go:embed testdata/v5_enabled_false.tf
var v5EnabledFalseConfig string

// TestMigrateLogpullRetention_V4ToV5_EnabledTrue tests migration with enabled=true → flag=true.
// This test validates:
// - Field rename: enabled → flag
// - Boolean value preservation: true → true
func TestMigrateLogpullRetention_V4ToV5_EnabledTrue(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4EnabledTrueConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5EnabledTrueConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test setup
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

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
								"cloudflare_logpull_retention."+rnd,
								tfjsonpath.New("zone_id"),
								knownvalue.StringExact(zoneID),
							),
							// CRITICAL: Verify enabled → flag rename with value preserved (true)
							statecheck.ExpectKnownValue(
								"cloudflare_logpull_retention."+rnd,
								tfjsonpath.New("flag"),
								knownvalue.Bool(true),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateLogpullRetention_V4ToV5_EnabledFalse tests migration with enabled=false → flag=false.
// This test validates:
// - Field rename: enabled → flag
// - Boolean value preservation: false → false
func TestMigrateLogpullRetention_V4ToV5_EnabledFalse(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4EnabledFalseConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5EnabledFalseConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test setup
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

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
								"cloudflare_logpull_retention."+rnd,
								tfjsonpath.New("zone_id"),
								knownvalue.StringExact(zoneID),
							),
							// CRITICAL: Verify enabled → flag rename with value preserved (false)
							statecheck.ExpectKnownValue(
								"cloudflare_logpull_retention."+rnd,
								tfjsonpath.New("flag"),
								knownvalue.Bool(false),
							),
						},
					),
				},
			})
		})
	}
}
