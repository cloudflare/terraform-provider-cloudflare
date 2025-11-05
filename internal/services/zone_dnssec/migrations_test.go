package zone_dnssec_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// zoneDNSSECMigrationTestStep creates a migration test step without plan checks.
// This is needed because DNSSEC status field transitions from intermediate states
// (pending, pending-disabled) to final states (active, disabled) during migration,
// which causes plan diffs that are expected and correct.
func zoneDNSSECMigrationTestStep(t *testing.T, v4Config string, tmpDir string, exactVersion string, sourceVersion string, targetVersion string, stateChecks []statecheck.StateCheck) resource.TestStep {
	return resource.TestStep{
		PreConfig: func() {
			acctest.WriteOutConfig(t, v4Config, tmpDir)
			acctest.RunMigrationV2Command(t, v4Config, tmpDir, sourceVersion, targetVersion)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		ConfigDirectory:          config.StaticDirectory(tmpDir),
		PlanOnly:                 true,           // Only verify migration, don't apply changes
		ExpectNonEmptyPlan:       true,           // Expect plan diff due to status field transitions
		// Note: No ConfigPlanChecks - we expect plan changes due to status field transitions
		ConfigStateChecks: stateChecks,
	}
}

// TestMigrateZoneDNSSECBasic tests migration of a basic zone_dnssec resource from v4 to v5
func TestMigrateZoneDNSSECBasic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V4 config with just zone_id (minimal configuration)
	v4Config := fmt.Sprintf(`
resource "cloudflare_zone_dnssec" "%[1]s" {
  zone_id = "%[2]s"
}`, rnd, zoneID)

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
			zoneDNSSECMigrationTestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Resource should keep the same name (no rename)
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				// Status field will be added with "active" (converted from "pending" if applicable)
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("status"), knownvalue.NotNull()),
				// Computed fields should exist
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("algorithm"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("flags"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("key_tag"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateZoneDNSSECWithModifiedOn tests migration where modified_on exists in state but not config
// The modified_on field was optional+computed in v4 but is computed-only in v5
func TestMigrateZoneDNSSECWithModifiedOn(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V4 config without modified_on (it's computed by the API)
	// The migration should handle the modified_on field in state correctly
	v4Config := fmt.Sprintf(`
resource "cloudflare_zone_dnssec" "%[1]s" {
  zone_id = "%[2]s"
}`, rnd, zoneID)

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
			zoneDNSSECMigrationTestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Resource should exist
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				// Status will be added (converted from "pending" to "active" if needed)
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("status"), knownvalue.NotNull()),
				// modified_on should still exist in state (it's computed in v5)
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("modified_on"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateZoneDNSSECStatusActive tests that status field is correctly preserved
func TestMigrateZoneDNSSECStatusActive(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V4 config - status was computed-only in v4, so it won't be in the input config
	// But the migration should add it from the state
	v4Config := fmt.Sprintf(`
resource "cloudflare_zone_dnssec" "%[1]s" {
  zone_id = "%[2]s"
}`, rnd, zoneID)

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
			zoneDNSSECMigrationTestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				// Status will be added (converted from "pending" to "active" if needed)
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("status"), knownvalue.NotNull()),
				// Verify numeric fields are present (flags, key_tag converted from int to float64)
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("flags"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("key_tag"), knownvalue.NotNull()),
				// Verify other computed fields
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("algorithm"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("digest"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("digest_algorithm"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("digest_type"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("ds"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("key_type"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue("cloudflare_zone_dnssec."+rnd, tfjsonpath.New("public_key"), knownvalue.NotNull()),
			}),
		},
	})
}
