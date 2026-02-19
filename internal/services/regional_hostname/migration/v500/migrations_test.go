package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

//go:embed testdata/v4_timeouts.tf
var v4TimeoutsConfig string

// TestMigrateRegionalHostname_TimeoutsRemoval tests v4→v5 migration with timeouts block removal
func TestMigrateRegionalHostname_TimeoutsRemoval(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	hostname := fmt.Sprintf("%s.%s", rnd, zoneName)
	resourceName := "cloudflare_regional_hostname." + rnd
	tmpDir := t.TempDir()
	testConfig := fmt.Sprintf(v4TimeoutsConfig, rnd, zoneID, hostname)
	version := acctest.GetLastV4Version()

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
						VersionConstraint: version,
					},
				},
				Config: testConfig,
			},
		},
			acctest.MigrationV2TestStepWithPlan(t, testConfig, tmpDir, version, "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(hostname)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("region_key"), knownvalue.StringExact("ca")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
			})...),
	})
}
