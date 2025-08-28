package regional_hostname_test

import (
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

// Config generators for different provider versions
func regionalHostnameConfigV4WithTimeouts(rnd, zoneID, hostname, regionKey string) string {
	return fmt.Sprintf(`
resource "cloudflare_regional_hostname" "%[1]s" {
  zone_id    = "%[2]s"
  hostname   = "%[3]s"
  region_key = "%[4]s"
  
  timeouts {
    create = "30s"
    update = "30s"
  }
}`, rnd, zoneID, hostname, regionKey)
}

func regionalHostnameConfigV4WithoutTimeouts(rnd, zoneID, hostname, regionKey string) string {
	return fmt.Sprintf(`
resource "cloudflare_regional_hostname" "%[1]s" {
  zone_id    = "%[2]s"
  hostname   = "%[3]s"
  region_key = "%[4]s"
}`, rnd, zoneID, hostname, regionKey)
}

func regionalHostnameConfigV5(rnd, zoneID, hostname, regionKey string) string {
	return fmt.Sprintf(`
resource "cloudflare_regional_hostname" "%[1]s" {
  zone_id    = "%[2]s"
  hostname   = "%[3]s"
  region_key = "%[4]s"
}`, rnd, zoneID, hostname, regionKey)
}

func TestMigrateRegionalHostnameMultiVersion(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")

	testCases := []struct {
		name          string
		version       string
		configFunc    func(rnd, zoneID, hostname, regionKey string) string
		skipPlanCheck bool // Some v5 versions have issues with refresh plan
	}{
		{
			name:       "from_v4_52_1_with_timeouts",
			version:    "4.52.1",
			configFunc: regionalHostnameConfigV4WithTimeouts,
		},
		{
			name:       "from_v4_52_1_without_timeouts",
			version:    "4.52.1",
			configFunc: regionalHostnameConfigV4WithoutTimeouts,
		},
		{
			name:       "from_v5_0_0",
			version:    "5.0.0",
			configFunc: regionalHostnameConfigV5,
		},
		{
			name:       "from_v5_8_4", // Current stable release
			version:    "5.8.4",
			configFunc: regionalHostnameConfigV5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			hostname := fmt.Sprintf("%s.%s", rnd, zoneName)
			resourceName := "cloudflare_regional_hostname." + rnd
			tmpDir := t.TempDir()
			regionKey := "ca"

			config := tc.configFunc(rnd, zoneID, hostname, regionKey)

			// Build test steps
			steps := []resource.TestStep{}

			// Step 1: Create with specific provider version
			step1 := resource.TestStep{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: tc.version,
					},
				},
				Config: config,
			}

			// Some v5 versions have issues, skip plan check if needed
			if tc.skipPlanCheck {
				step1.ExpectNonEmptyPlan = true
			}

			steps = append(steps, step1)

			// Step 2: Run migration (for v4) or just upgrade provider (for v5)
			// MigrationTestStep automatically detects v4 vs v5 and only runs migration for v4
			steps = append(steps,
				acctest.MigrationTestStep(t, config, tmpDir, tc.version, []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(hostname)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("region_key"), knownvalue.StringExact(regionKey)),
					// Verify routing has default value (handled by state upgrader)
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("routing"), knownvalue.StringExact("dns")),
				}),
			)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps:      steps,
			})
		})
	}
}

func TestMigrateRegionalHostnameEdgeCases(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")

	// Test edge cases with different region keys
	edgeCases := []struct {
		name           string
		version        string
		regionKey      string
		configFunc     func(rnd, zoneID, hostname, regionKey string) string
		expectedChecks func(resourceName, zoneID, hostname, regionKey string) []statecheck.StateCheck
	}{
		{
			name:       "v4_with_eu_region",
			version:    "4.52.1",
			regionKey:  "eu",
			configFunc: regionalHostnameConfigV4WithTimeouts,
			expectedChecks: func(resourceName, zoneID, hostname, regionKey string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(hostname)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("region_key"), knownvalue.StringExact(regionKey)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("routing"), knownvalue.StringExact("dns")),
				}
			},
		},
		{
			name:       "v5_with_us_region",
			version:    "5.0.0",
			regionKey:  "us",
			configFunc: regionalHostnameConfigV5,
			expectedChecks: func(resourceName, zoneID, hostname, regionKey string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(hostname)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("region_key"), knownvalue.StringExact(regionKey)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("routing"), knownvalue.StringExact("dns")),
				}
			},
		},
		{
			name:       "v4_with_us_region",
			version:    "4.52.1",
			regionKey:  "us",
			configFunc: regionalHostnameConfigV4WithoutTimeouts,
			expectedChecks: func(resourceName, zoneID, hostname, regionKey string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(hostname)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("region_key"), knownvalue.StringExact(regionKey)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("routing"), knownvalue.StringExact("dns")),
				}
			},
		},
	}

	for _, tc := range edgeCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			hostname := fmt.Sprintf("%s.%s", rnd, zoneName)
			resourceName := "cloudflare_regional_hostname." + rnd
			tmpDir := t.TempDir()

			config := tc.configFunc(rnd, zoneID, hostname, tc.regionKey)
			expectedChecks := tc.expectedChecks(resourceName, zoneID, hostname, tc.regionKey)

			// Build test steps
			steps := []resource.TestStep{}

			// Step 1: Create with specific provider version
			step1 := resource.TestStep{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: tc.version,
					},
				},
				Config: config,
			}

			steps = append(steps, step1)

			// Step 2: Run migration (for v4) or just upgrade provider (for v5)
			steps = append(steps,
				acctest.MigrationTestStep(t, config, tmpDir, tc.version, expectedChecks),
			)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps:      steps,
			})
		})
	}
}
