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

const (
	legacyProviderVersion  = internal.LastV4Version  // Last v4 release
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Embed test configs
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_cache_key_fields.tf
var v4CacheKeyFieldsConfig string

//go:embed testdata/v5_cache_key_fields.tf
var v5CacheKeyFieldsConfig string

// TestMigratePageRule_V4ToV5_Basic tests basic field migrations with DUAL test cases.
// This test validates:
// - status default preservation (v4="active" → v5="active", not v5 default "disabled")
// - actions extraction (TypeList MaxItems:1 → SingleNestedAttribute)
func TestMigratePageRule_V4ToV5_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, target string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: legacyProviderVersion,
			configFn: func(rnd, zoneID, target string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, zoneID, target)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, zoneID, target string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, zoneID, target)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test setup
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			target := "tf-test-" + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID, target)
			sourceVer, targetVer := acctest.DetermineSourceTargetVersion(tc.version)

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
								"cloudflare_page_rule."+rnd,
								tfjsonpath.New("zone_id"),
								knownvalue.StringExact(zoneID),
							),
							// Verify target preserved
							statecheck.ExpectKnownValue(
								"cloudflare_page_rule."+rnd,
								tfjsonpath.New("target"),
								knownvalue.StringExact(target+".example.com/*"),
							),
							// CRITICAL: Verify status is "active" (v4 default preserved, not v5 default "disabled")
							statecheck.ExpectKnownValue(
								"cloudflare_page_rule."+rnd,
								tfjsonpath.New("status"),
								knownvalue.StringExact("active"),
							),
							// Verify actions.cache_level (tests actions extraction from array)
							statecheck.ExpectKnownValue(
								"cloudflare_page_rule."+rnd,
								tfjsonpath.New("actions").AtMapKey("cache_level"),
								knownvalue.StringExact("bypass"),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigratePageRule_V4ToV5_CacheKeyFields tests 5-level deep nested structure transformation.
// This test validates:
// - cache_key_fields extraction (5 levels: actions → cache_key_fields → host/query_string/user)
// - user.lang field addition (v4 may not have this, must add with default false)
// - Set[String] → List[String] conversions
func TestMigratePageRule_V4ToV5_CacheKeyFields(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, target string) string
	}{
		{
			name:    "from_v4_latest",
			version: legacyProviderVersion,
			configFn: func(rnd, zoneID, target string) string {
				return fmt.Sprintf(v4CacheKeyFieldsConfig, rnd, zoneID, target)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID, target string) string {
				return fmt.Sprintf(v5CacheKeyFieldsConfig, rnd, zoneID, target)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test setup
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			target := "tf-test-" + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID, target)
			sourceVer, targetVer := acctest.DetermineSourceTargetVersion(tc.version)

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
							// Verify cache_key_fields.host.resolved
							statecheck.ExpectKnownValue(
								"cloudflare_page_rule."+rnd,
								tfjsonpath.New("actions").AtMapKey("cache_key_fields").AtMapKey("host").AtMapKey("resolved"),
								knownvalue.Bool(true),
							),
							// Verify cache_key_fields.query_string.exclude
							statecheck.ExpectKnownValue(
								"cloudflare_page_rule."+rnd,
								tfjsonpath.New("actions").AtMapKey("cache_key_fields").AtMapKey("query_string").AtMapKey("exclude").AtSliceIndex(0),
								knownvalue.StringExact("utm_source"),
							),
							// Verify cache_key_fields.user.device_type
							statecheck.ExpectKnownValue(
								"cloudflare_page_rule."+rnd,
								tfjsonpath.New("actions").AtMapKey("cache_key_fields").AtMapKey("user").AtMapKey("device_type"),
								knownvalue.Bool(true),
							),
							// Verify cache_key_fields.user.geo
							statecheck.ExpectKnownValue(
								"cloudflare_page_rule."+rnd,
								tfjsonpath.New("actions").AtMapKey("cache_key_fields").AtMapKey("user").AtMapKey("geo"),
								knownvalue.Bool(false),
							),
							// CRITICAL: Verify cache_key_fields.user.lang = false (added during migration)
							statecheck.ExpectKnownValue(
								"cloudflare_page_rule."+rnd,
								tfjsonpath.New("actions").AtMapKey("cache_key_fields").AtMapKey("user").AtMapKey("lang"),
								knownvalue.Bool(false),
							),
						},
					),
				},
			})
		})
	}
}
