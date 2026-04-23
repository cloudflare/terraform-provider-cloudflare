package custom_hostname_fallback_origin_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateCustomHostnameFallbackOriginBasic tests basic fallback origin migration from v4 to v5
// This is a simple pass-through migration - no field renames, no type conversions, no structural changes
func TestMigrateCustomHostnameFallbackOriginBasic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_hostname_fallback_origin." + rnd
	tmpDir := t.TempDir()
	originHostname := fmt.Sprintf("cftftest-fallback-%s.cf-tf-test.com", rnd)

	// V4 config - basic fallback origin with required fields
	v4Config := fmt.Sprintf(`
resource "cloudflare_custom_hostname_fallback_origin" "%[1]s" {
  zone_id = "%[2]s"
  origin  = "%[3]s"
}`, rnd, zoneID, originHostname)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Verify resource exists (no type rename for this resource)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				// Verify required fields are preserved
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin"), knownvalue.StringExact(originHostname)),
				// Verify computed fields exist (status is computed by API)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateCustomHostnameFallbackOriginMinimal tests migration with absolutely minimal configuration
func TestMigrateCustomHostnameFallbackOriginMinimal(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_hostname_fallback_origin." + rnd
	tmpDir := t.TempDir()
	originHostname := fmt.Sprintf("cftftest-minimal-%s.cf-tf-test.com", rnd)

	// V4 config - absolutely minimal configuration (only required fields)
	v4Config := fmt.Sprintf(`
resource "cloudflare_custom_hostname_fallback_origin" "%[1]s" {
  zone_id = "%[2]s"
  origin  = "%[3]s"
}`, rnd, zoneID, originHostname)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Verify the migration preserves all user-provided fields
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin"), knownvalue.StringExact(originHostname)),
				// Verify computed field (status) is present after migration
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.NotNull()),
			}),
		},
	})
}
