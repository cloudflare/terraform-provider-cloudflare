package regional_hostname_test

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

func TestMigrateCloudflareRegionalHostname_Migration_TimeoutsRemoval(t *testing.T) {
	// This test verifies that the migration tool properly removes the timeouts block
	// when upgrading from v4 to v5, since v5 provider no longer supports timeouts
	// configuration for regional_hostname resources.

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	hostname := fmt.Sprintf("%s.%s", rnd, zoneName)
	resourceName := "cloudflare_regional_hostname." + rnd
	tmpDir := t.TempDir()

	// V4 config with timeouts block (should be removed during migration)
	v4Config := fmt.Sprintf(`
resource "cloudflare_regional_hostname" "%[1]s" {
  zone_id    = "%[2]s"
  hostname   = "%[3]s"
  region_key = "ca"

  timeouts {
    create = "30s"
    update = "30s"
  }
}`, rnd, zoneID, hostname)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider including timeouts
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1", // Use exact v4 version
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(hostname)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("region_key"), knownvalue.StringExact("ca")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				},
			},
			// Step 2: Run migration from v4 to v5
			// This will run the migration tool which should remove the timeouts block
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(hostname)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("region_key"), knownvalue.StringExact("ca")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
			}),
		},
	})
}
