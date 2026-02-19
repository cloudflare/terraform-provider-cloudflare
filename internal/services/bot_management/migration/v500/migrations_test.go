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

//go:embed testdata/v4_ai_bots_protection.tf
var v4AIBotsProtectionConfig string

//go:embed testdata/v5_ai_bots_protection.tf
var v5AIBotsProtectionConfig string

//go:embed testdata/v4_multiple_fields.tf
var v4MultipleFieldsConfig string

//go:embed testdata/v5_multiple_fields.tf
var v5MultipleFieldsConfig string

// TestMigrateBotManagement_V4ToV5_Basic tests basic bot management migration with minimal fields
func TestMigrateBotManagement_V4ToV5_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_bot_management." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

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
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_update_model"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
						},
					),
				},
			})
		})
	}
}

// TestMigrateBotManagement_V4ToV5_AIBotsProtection tests migration with AI bots protection
func TestMigrateBotManagement_V4ToV5_AIBotsProtection(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4AIBotsProtectionConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5AIBotsProtectionConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_bot_management." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

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
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_bots_protection"), knownvalue.StringExact("block")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_update_model"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("suppress_session_score"), knownvalue.Bool(false)),
						},
					),
				},
			})
		})
	}
}

// TestMigrateBotManagement_V4ToV5_MultipleFields tests migration with multiple fields
func TestMigrateBotManagement_V4ToV5_MultipleFields(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4MultipleFieldsConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5MultipleFieldsConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_bot_management." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

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
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_js"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_update_model"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("suppress_session_score"), knownvalue.Bool(false)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ai_bots_protection"), knownvalue.StringExact("block")),
						},
					),
				},
			})
		})
	}
}
