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

// Embed migration test configuration files
//
//go:embed testdata/v4_cloudflare_incoming.tf
var v4CloudflareIncomingConfig string

//go:embed testdata/v5_cloudflare_incoming.tf
var v5CloudflareIncomingConfig string

//go:embed testdata/v4_rfc3986_both.tf
var v4RFC3986BothConfig string

//go:embed testdata/v5_rfc3986_both.tf
var v5RFC3986BothConfig string

//go:embed testdata/v4_cloudflare_none.tf
var v4CloudflareNoneConfig string

//go:embed testdata/v5_cloudflare_none.tf
var v5CloudflareNoneConfig string

// TestMigrateURLNormalizationSettings_CloudflareIncoming tests migration with type=cloudflare, scope=incoming
func TestMigrateURLNormalizationSettings_CloudflareIncoming(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4CloudflareIncomingConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5CloudflareIncomingConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_url_normalization_settings." + rnd
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: append([]resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
				},
					acctest.MigrationV2TestStepWithPlan(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("cloudflare")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scope"), knownvalue.StringExact("incoming")),
					})...),
			})
		})
	}
}

// TestMigrateURLNormalizationSettings_RFC3986Both tests migration with type=rfc3986, scope=both
func TestMigrateURLNormalizationSettings_RFC3986Both(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4RFC3986BothConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5RFC3986BothConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_url_normalization_settings." + rnd
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: append([]resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
				},
					acctest.MigrationV2TestStepWithPlan(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("rfc3986")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scope"), knownvalue.StringExact("both")),
					})...),
			})
		})
	}
}

// TestMigrateURLNormalizationSettings_CloudflareNone tests migration with type=cloudflare, scope=none
func TestMigrateURLNormalizationSettings_CloudflareNone(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4CloudflareNoneConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5CloudflareNoneConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_url_normalization_settings." + rnd
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: append([]resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
				},
					acctest.MigrationV2TestStepWithPlan(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("cloudflare")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scope"), knownvalue.StringExact("none")),
					})...),
			})
		})
	}
}
