package regional_hostname_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareRegionalHostname_Migration_TimeoutsRemoval(t *testing.T) {
	// This test verifies that the migration tool properly removes the timeouts block
	// when upgrading from v4 to v5, since v5 provider no longer supports timeouts
	// configuration for regional_hostname resources.

	rnd := utils.GenerateRandomResourceName()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_regional_hostname." + rnd
	tmpDir := t.TempDir()

	// V4 config with timeouts block
	v4Config := testRegionalHostnameV4ConfigWithTimeouts(rnd, zoneName, "ca")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
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
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, zoneName))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("region_key"), knownvalue.StringExact("ca")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				},
			},
			// Step 2: Run migration from v4 to current version
			// This will run the migration tool which should remove the timeouts block
			// and the state upgrade function will handle the state transformation
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, zoneName))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("region_key"), knownvalue.StringExact("ca")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
			}),
			{
				// Step 3: Apply migrated config with current provider
				// Should succeed without any timeouts configuration
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, zoneName))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("region_key"), knownvalue.StringExact("ca")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					// Verify routing has default value (handled by state upgrader)
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("routing"), knownvalue.StringExact("dns")),
				},
			},
			{
				// Step 4: Import verification to ensure resource is properly accessible
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ResourceName:             resourceName,
				ImportState:              true,
				ImportStateVerify:        true,
				ImportStateIdFunc:        testAccCloudflareRegionalHostnameImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore:  []string{"created_on"}, // Computed field may vary
			},
		},
	})
}

func testRegionalHostnameV4ConfigWithTimeouts(name string, zoneName, regionKey string) string {
	// Use a random subdomain to avoid conflicts with existing regional hostnames
	hostname := fmt.Sprintf("%s.%s", name, zoneName)
	return acctest.LoadTestCase("regionalhostname_v4_with_timeouts.tf", name, zoneID, hostname, regionKey)
}