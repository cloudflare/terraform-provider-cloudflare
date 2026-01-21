package url_normalization_settings_test

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

func TestMigrateURLNormalizationSettings_Cloudflare(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()
	resourceName := "cloudflare_url_normalization_settings." + rnd

	// v4 config - no changes needed, pass-through migration
	v4Config := fmt.Sprintf(`
resource "cloudflare_url_normalization_settings" "%[1]s" {
  zone_id = "%[2]s"
  type    = "cloudflare"
  scope   = "incoming"
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		CheckDestroy: nil, // Settings resource doesn't have destroy check
		WorkingDir:   tmpDir,
		Steps: append([]resource.TestStep{
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
		}, // Step 2: Run migration and verify state
			acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("cloudflare")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scope"), knownvalue.StringExact("incoming")),
			})...),
	})
}

func TestMigrateURLNormalizationSettings_RFC3986(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()
	resourceName := "cloudflare_url_normalization_settings." + rnd

	// v4 config with RFC3986 type
	v4Config := fmt.Sprintf(`
resource "cloudflare_url_normalization_settings" "%[1]s" {
  zone_id = "%[2]s"
  type    = "rfc3986"
  scope   = "both"
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		CheckDestroy: nil, // Settings resource doesn't have destroy check
		WorkingDir:   tmpDir,
		Steps: append([]resource.TestStep{
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
		}, // Step 2: Run migration and verify state
			acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("rfc3986")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scope"), knownvalue.StringExact("both")),
			})...),
	})
}

func TestMigrateURLNormalizationSettings_NoneScope(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()
	resourceName := "cloudflare_url_normalization_settings." + rnd

	// v4 config with "none" scope
	v4Config := fmt.Sprintf(`
resource "cloudflare_url_normalization_settings" "%[1]s" {
  zone_id = "%[2]s"
  type    = "cloudflare"
  scope   = "none"
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		CheckDestroy: nil, // Settings resource doesn't have destroy check
		WorkingDir:   tmpDir,
		Steps: append([]resource.TestStep{
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
		}, // Step 2: Run migration and verify state
			acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("cloudflare")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("scope"), knownvalue.StringExact("none")),
			})...),
	})
}
