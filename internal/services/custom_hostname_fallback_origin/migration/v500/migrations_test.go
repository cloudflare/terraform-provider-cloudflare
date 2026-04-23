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

// Migration Test Configuration
//
// Version is read from LAST_V4_VERSION environment variable (set in .github/workflows/migration-tests.yml)
// - Last stable v4 release: default 4.52.5
// - Current v5 release: auto-updates with releases (internal.PackageVersion)
//
// For this resource:
// - NO breaking changes between v4 and v5 (fields identical)
// - Simple pass-through migration (direct copy)
// - No field renames, no type changes, no structure changes

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

// TestMigrateCustomHostnameFallbackOriginBasic tests migration of custom_hostname_fallback_origin resource.
//
// This is an ultra-simple migration:
// - No field renames
// - No type conversions
// - No structural changes
// - Models are identical between v4 and v5
//
// The migration tests verify that:
// 1. v4→v5 migration preserves all user-configured fields (zone_id, origin)
// 2. v5→v5 version bump works correctly (no-op upgrade)
// 3. Computed fields (id, status) are present after migration
func TestMigrateCustomHostnameFallbackOriginBasic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_hostname_fallback_origin." + rnd
	tmpDir := t.TempDir()
	originHostname := fmt.Sprintf("cftftest-fallback-%s.cf-tf-test.com", rnd)

	// Configuration function for parameterized configs
	configFn := func(config string) string {
		return fmt.Sprintf(config, rnd, zoneID, originHostname)
	}

	legacyProviderVersion := acctest.GetLastV4Version()
	sourceVer, targetVer := acctest.InferMigrationVersions(legacyProviderVersion)

	// Cleanup: Wait for resource to fully delete from Cloudflare API
	// This is a singleton resource, and the API needs time to complete deletion
	// before another test can create a new fallback origin in the same zone.
	t.Cleanup(func() {
		t.Logf("Waiting 5 seconds for fallback origin to fully delete from Cloudflare API...")
		time.Sleep(5 * time.Second)
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Test Case 1: Migration from v4 latest
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: legacyProviderVersion,
					},
				},
				Config: configFn(v4BasicConfig),
			},
			// Run migration from v4 to v5
			acctest.MigrationV2TestStep(t, configFn(v4BasicConfig), tmpDir, legacyProviderVersion, sourceVer, targetVer, []statecheck.StateCheck{
				// Verify user-configured fields are preserved exactly
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin"), knownvalue.StringExact(originHostname)),
				// Verify computed fields exist (API-assigned)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateCustomHostnameFallbackOriginFromV5 tests v5→v5 version bump (no-op upgrade).
//
// This test verifies that when TF_MIG_TEST=1, the resource can upgrade from version=1 to version=500
// without any data loss or transformation issues.
func TestMigrateCustomHostnameFallbackOriginFromV5(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_hostname_fallback_origin." + rnd
	tmpDir := t.TempDir()
	originHostname := fmt.Sprintf("cftftest-v5-%s.cf-tf-test.com", rnd)

	// Configuration function for parameterized configs
	configFn := func(config string) string {
		return fmt.Sprintf(config, rnd, zoneID, originHostname)
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Test Case 2: Create with current v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   configFn(v5BasicConfig),
			},
			// Run no-op migration from v5 (version=1) to v5 (version=500)
			acctest.MigrationV2TestStep(t, configFn(v5BasicConfig), tmpDir, currentProviderVersion, "v5", "v5", []statecheck.StateCheck{
				// Verify all fields are preserved during no-op upgrade
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin"), knownvalue.StringExact(originHostname)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.NotNull()),
			}),
		},
	})
}
