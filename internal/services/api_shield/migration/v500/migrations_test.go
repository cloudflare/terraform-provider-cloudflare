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

// Migration Test Configuration
//
// Version is read from LAST_V4_VERSION environment variable (set in .github/workflows/migration-tests.yml)
// - Last stable v4 release: default 4.52.5
// - Current v5 release: auto-updates with releases (internal.PackageVersion)
//
// For this resource:
// - Breaking change: auth_id_characteristics block syntax → attribute syntax
// - Simple transformation: array structure identical, only HCL syntax differs
// - Optional → Required: empty array handling for missing auth_id_characteristics

//go:embed testdata/v4_single_header.tf
var v4SingleHeaderConfig string

//go:embed testdata/v4_single_cookie.tf
var v4SingleCookieConfig string

//go:embed testdata/v4_multiple_mixed.tf
var v4MultipleMixedConfig string

//go:embed testdata/v4_special_chars.tf
var v4SpecialCharsConfig string

//go:embed testdata/v4_oauth_flow.tf
var v4OAuthFlowConfig string

//go:embed testdata/v4_empty.tf
var v4EmptyConfig string

// TestMigrateAPIShield_V4ToV5_SingleHeader tests migration with single header characteristic.
//
// This is the most common use case: a single Authorization header for API authentication.
// Validates:
// - Block syntax → Attribute syntax conversion
// - zone_id preservation
// - Single array element transformation
// - Type and name field preservation
func TestMigrateAPIShield_V4ToV5_SingleHeader(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_shield." + rnd
	tmpDir := t.TempDir()

	configFn := func(config string) string {
		return fmt.Sprintf(config, rnd, zoneID)
	}

	legacyProviderVersion := os.Getenv("LAST_V4_VERSION")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: legacyProviderVersion,
					},
				},
				Config: configFn(v4SingleHeaderConfig),
			},
			// Run migration from v4 to v5
			acctest.MigrationV2TestStep(t, configFn(v4SingleHeaderConfig), tmpDir, legacyProviderVersion, "v4", "v5", []statecheck.StateCheck{
				// Verify user-configured fields preserved
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("auth_id_characteristics").AtSliceIndex(0).AtMapKey("type"),
					knownvalue.StringExact("header")),
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("auth_id_characteristics").AtSliceIndex(0).AtMapKey("name"),
					knownvalue.StringExact("authorization")),
				// Verify computed field exists
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateAPIShield_V4ToV5_SingleCookie tests migration with single cookie characteristic.
//
// Validates alternative type value (cookie instead of header).
func TestMigrateAPIShield_V4ToV5_SingleCookie(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_shield." + rnd
	tmpDir := t.TempDir()

	configFn := func(config string) string {
		return fmt.Sprintf(config, rnd, zoneID)
	}

	legacyProviderVersion := os.Getenv("LAST_V4_VERSION")

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
						VersionConstraint: legacyProviderVersion,
					},
				},
				Config: configFn(v4SingleCookieConfig),
			},
			acctest.MigrationV2TestStep(t, configFn(v4SingleCookieConfig), tmpDir, legacyProviderVersion, "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("auth_id_characteristics").AtSliceIndex(0).AtMapKey("type"),
					knownvalue.StringExact("cookie")),
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("auth_id_characteristics").AtSliceIndex(0).AtMapKey("name"),
					knownvalue.StringExact("session_id")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateAPIShield_V4ToV5_MultipleMixed tests migration with multiple characteristics.
//
// Validates:
// - Array transformation with multiple elements (3 characteristics)
// - Mixed types (headers + cookies)
// - Ordering preservation
func TestMigrateAPIShield_V4ToV5_MultipleMixed(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_shield." + rnd
	tmpDir := t.TempDir()

	configFn := func(config string) string {
		return fmt.Sprintf(config, rnd, zoneID)
	}

	legacyProviderVersion := os.Getenv("LAST_V4_VERSION")

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
						VersionConstraint: legacyProviderVersion,
					},
				},
				Config: configFn(v4MultipleMixedConfig),
			},
			acctest.MigrationV2TestStep(t, configFn(v4MultipleMixedConfig), tmpDir, legacyProviderVersion, "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics"), knownvalue.ListSizeExact(3)),
				// Verify element [0]: type="header", name="authorization"
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("auth_id_characteristics").AtSliceIndex(0).AtMapKey("type"),
					knownvalue.StringExact("header")),
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("auth_id_characteristics").AtSliceIndex(0).AtMapKey("name"),
					knownvalue.StringExact("authorization")),
				// Verify element [1]: type="cookie", name="session_id"
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("auth_id_characteristics").AtSliceIndex(1).AtMapKey("type"),
					knownvalue.StringExact("cookie")),
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("auth_id_characteristics").AtSliceIndex(1).AtMapKey("name"),
					knownvalue.StringExact("session_id")),
				// Verify element [2]: type="header", name="x-api-key"
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("auth_id_characteristics").AtSliceIndex(2).AtMapKey("type"),
					knownvalue.StringExact("header")),
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("auth_id_characteristics").AtSliceIndex(2).AtMapKey("name"),
					knownvalue.StringExact("x-api-key")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateAPIShield_V4ToV5_SpecialChars tests migration with special characters in names.
//
// Validates name preservation with various patterns:
// - Hyphens (X-API-Key, user-session-token)
// - Underscores (X_Custom_Header)
// - Case sensitivity (SessionID, CamelCase)
//
// Critical because HTTP headers often use hyphens and mixed case.
func TestMigrateAPIShield_V4ToV5_SpecialChars(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_shield." + rnd
	tmpDir := t.TempDir()

	configFn := func(config string) string {
		return fmt.Sprintf(config, rnd, zoneID)
	}

	legacyProviderVersion := os.Getenv("LAST_V4_VERSION")

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
						VersionConstraint: legacyProviderVersion,
					},
				},
				Config: configFn(v4SpecialCharsConfig),
			},
			acctest.MigrationV2TestStep(t, configFn(v4SpecialCharsConfig), tmpDir, legacyProviderVersion, "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics"), knownvalue.ListSizeExact(4)),
				// Verify names with special characters preserved exactly
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("auth_id_characteristics").AtSliceIndex(0).AtMapKey("name"),
					knownvalue.StringExact("X-API-Key")),
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("auth_id_characteristics").AtSliceIndex(1).AtMapKey("name"),
					knownvalue.StringExact("X_Custom_Header")),
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("auth_id_characteristics").AtSliceIndex(2).AtMapKey("name"),
					knownvalue.StringExact("SessionID")),
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("auth_id_characteristics").AtSliceIndex(3).AtMapKey("name"),
					knownvalue.StringExact("user-session-token")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateAPIShield_V4ToV5_OAuthFlow tests migration with OAuth-style authentication.
//
// Simulates real-world OAuth implementation with:
// - Multiple headers (Authorization, X-OAuth-Token, X-Request-ID)
// - OAuth state cookie
// - Various naming patterns
func TestMigrateAPIShield_V4ToV5_OAuthFlow(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_shield." + rnd
	tmpDir := t.TempDir()

	configFn := func(config string) string {
		return fmt.Sprintf(config, rnd, zoneID)
	}

	legacyProviderVersion := os.Getenv("LAST_V4_VERSION")

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
						VersionConstraint: legacyProviderVersion,
					},
				},
				Config: configFn(v4OAuthFlowConfig),
			},
			acctest.MigrationV2TestStep(t, configFn(v4OAuthFlowConfig), tmpDir, legacyProviderVersion, "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics"), knownvalue.ListSizeExact(4)),
				// Verify OAuth-style characteristics
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("auth_id_characteristics").AtSliceIndex(0).AtMapKey("name"),
					knownvalue.StringExact("Authorization")),
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("auth_id_characteristics").AtSliceIndex(1).AtMapKey("name"),
					knownvalue.StringExact("X-OAuth-Token")),
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("auth_id_characteristics").AtSliceIndex(2).AtMapKey("type"),
					knownvalue.StringExact("cookie")),
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("auth_id_characteristics").AtSliceIndex(2).AtMapKey("name"),
					knownvalue.StringExact("oauth_state")),
				statecheck.ExpectKnownValue(resourceName,
					tfjsonpath.New("auth_id_characteristics").AtSliceIndex(3).AtMapKey("name"),
					knownvalue.StringExact("X-Request-ID")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateAPIShield_V4ToV5_Empty tests migration with empty auth_id_characteristics.
//
// This is a CRITICAL edge case:
// - v4: auth_id_characteristics is Optional (can be omitted)
// - v5: auth_id_characteristics is Required (must have value)
// - Migration must set empty array [] for v5 schema compliance
//
// Validates the Optional → Required field migration pattern.
func TestMigrateAPIShield_V4ToV5_Empty(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_shield." + rnd
	tmpDir := t.TempDir()

	configFn := func(config string) string {
		return fmt.Sprintf(config, rnd, zoneID)
	}

	legacyProviderVersion := os.Getenv("LAST_V4_VERSION")

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
						VersionConstraint: legacyProviderVersion,
					},
				},
				Config: configFn(v4EmptyConfig),
			},
			acctest.MigrationV2TestStep(t, configFn(v4EmptyConfig), tmpDir, legacyProviderVersion, "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				// Verify empty array handling: Optional → Required migration
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics"), knownvalue.ListSizeExact(0)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
			}),
		},
	})
}
