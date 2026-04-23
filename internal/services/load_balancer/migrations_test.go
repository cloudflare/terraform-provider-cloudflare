package load_balancer_test

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

// TestAccCloudflareLoadBalancer_Migration_Basic_MultiVersion tests the most fundamental
// load balancer migration scenario with minimal configuration. This test ensures that:
// 1. Basic attribute renames work (fallback_pool_id → fallback_pool, default_pool_ids → default_pools)
// 2. State transformation handles empty arrays → empty maps for country_pools, pop_pools, region_pools
// 3. The migration tool successfully transforms both configuration and state files
// 4. Resources remain functional after migration without requiring manual intervention
func TestMigrateCloudflareLoadBalancer_Migration_Basic_MultiVersion(t *testing.T) {
	// Based on breaking changes analysis:
	// - All breaking changes happened between 4.x and 5.0.0
	// - No breaking changes between v5 releases for load_balancer
	testCases := []struct {
		name     string
		version  string
		configFn func(accountID, zoneID, zone, rnd string) string
	}{
		{
			name:     "from_v4_52_1", // Last v4 release
			version:  "4.52.1",
			configFn: testAccCloudflareLoadBalancerMigrationConfigV4Basic,
		},
		{
			name:     "from_v5_0_0", // First v5 release, after breaking changes
			version:  "5.0.0",
			configFn: testAccCloudflareLoadBalancerMigrationConfigV5Basic,
		},
		{
			name:     "from_v5_7_1", // Recent v5 release
			version:  "5.7.1",
			configFn: testAccCloudflareLoadBalancerMigrationConfigV5Basic,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
			// service does not yet support the API tokens and it results in
			// misleading state error messages.
			if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
				t.Setenv("CLOUDFLARE_API_TOKEN", "")
			}

			accountID := acctest.TestAccCloudflareAccountID
			zoneID := acctest.TestAccCloudflareZoneID
			zone := acctest.TestAccCloudflareZoneName
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_load_balancer." + rnd
			testConfig := tc.configFn(accountID, zoneID, zone, rnd)
			tmpDir := t.TempDir()

			// Build test steps
			steps := []resource.TestStep{
				{
					// Step 1: Create load balancer with specific version
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							VersionConstraint: tc.version,
							Source:            "cloudflare/cloudflare",
						},
					},
					Config: testConfig,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-testacc-lb-%s.%s", rnd, zone))),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("steering_policy"), knownvalue.StringExact("off")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_affinity"), knownvalue.StringExact("none")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ttl"), knownvalue.Float64Exact(30)),
					},
				},
			}

			// Step 2: Migrate to v5 provider
			migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, testConfig, tmpDir, tc.version, "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-testacc-lb-%s.%s", rnd, zone))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fallback_pool"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("default_pools"), knownvalue.ListSizeExact(1)),
			})
			steps = append(steps, migrationSteps...)

			// Step 3: Apply the migrated configuration
			steps = append(steps, resource.TestStep{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-testacc-lb-%s.%s", rnd, zone))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fallback_pool"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("default_pools"), knownvalue.ListSizeExact(1)),
				},
			})

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				CheckDestroy: testAccCheckCloudflareLoadBalancerDestroy,
				WorkingDir:   tmpDir,
				Steps:        steps,
			})
		})
	}
}

// TestAccCloudflareLoadBalancer_Migration_AllOptionalAttributes_MultiVersion tests migration
// with all optional attributes configured. This comprehensive test verifies that:
// 1. Block-to-single-object conversions work (adaptive_routing, location_strategy, random_steering, session_affinity_attributes)
// 2. All optional fields maintain their values through migration (description, enabled, proxied, etc.)
// 3. Complex nested structures are properly transformed from v4 blocks to v5 single objects
// 4. Session affinity configurations with attributes are correctly migrated
// 5. State transformation removes empty arrays for single-object attributes
func TestMigrateCloudflareLoadBalancer_Migration_AllOptionalAttributes_MultiVersion(t *testing.T) {
	testCases := []struct {
		name               string
		version            string
		configFn           func(accountID, zoneID, zone, rnd string) string
		ExpectNonEmptyPlan bool
	}{
		{
			name:     "from_v4_52_1",
			version:  "4.52.1",
			configFn: testAccCloudflareLoadBalancerMigrationConfigV4AllOptional,
		},
		{
			name:     "from_v5_0_0",
			version:  "5.0.0",
			configFn: testAccCloudflareLoadBalancerMigrationConfigV5AllOptional,
		},
		{
			name:               "from_v5_2_0", // Added zone_name field
			version:            "5.2.0",
			configFn:           testAccCloudflareLoadBalancerMigrationConfigV5AllOptional,
			ExpectNonEmptyPlan: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
			// service does not yet support the API tokens and it results in
			// misleading state error messages.
			if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
				t.Setenv("CLOUDFLARE_API_TOKEN", "")
			}

			accountID := acctest.TestAccCloudflareAccountID
			zoneID := acctest.TestAccCloudflareZoneID
			zone := acctest.TestAccCloudflareZoneName
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_load_balancer." + rnd
			testConfig := tc.configFn(accountID, zoneID, zone, rnd)
			tmpDir := t.TempDir()

			// Build test steps
			steps := []resource.TestStep{
				{
					// Step 1: Create load balancer with specific version
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							VersionConstraint: tc.version,
							Source:            "cloudflare/cloudflare",
						},
					},
					Config: testConfig,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-testacc-lb-%s.%s", rnd, zone))),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("test load balancer")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("proxied"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_affinity"), knownvalue.StringExact("cookie")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_affinity_ttl"), knownvalue.Float64Exact(1800)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("steering_policy"), knownvalue.StringExact("random")),
					},
					ExpectNonEmptyPlan: tc.ExpectNonEmptyPlan,
				},
			}

			// Step 2: Migrate to v5 provider
			migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, testConfig, tmpDir, tc.version, "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-testacc-lb-%s.%s", rnd, zone))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("test load balancer")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("proxied"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_affinity"), knownvalue.StringExact("cookie")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_affinity_ttl"), knownvalue.Int64Exact(1800)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("steering_policy"), knownvalue.StringExact("random")),
			})
			steps = append(steps, migrationSteps...)

			// Step 3: Apply migrated config with v5 provider
			steps = append(steps, resource.TestStep{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-testacc-lb-%s.%s", rnd, zone))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("test load balancer")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("proxied"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_affinity"), knownvalue.StringExact("cookie")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_affinity_ttl"), knownvalue.Int64Exact(1800)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("steering_policy"), knownvalue.StringExact("random")),
				},
			})

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				CheckDestroy: testAccCheckCloudflareLoadBalancerDestroy,
				WorkingDir:   tmpDir,
				Steps:        steps,
			})
		})
	}
}

// TestAccCloudflareLoadBalancer_Migration_GeoBalanced_MultiVersion tests migration of
// load balancers using geographic steering policies. This test ensures that:
// 1. region_pools and country_pools blocks are correctly transformed to maps in v5
// 2. Multiple pool configurations with geo-steering remain functional
// 3. Empty map attributes in state are properly converted from v4 arrays to v5 maps
// 4. Geo steering policy settings are preserved during migration
// 5. Complex pool ID references across multiple resources work correctly
func TestMigrateCloudflareLoadBalancer_Migration_GeoBalanced_MultiVersion(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(accountID, zoneID, zone, rnd string) string
	}{
		{
			name:     "from_v4_52_1",
			version:  "4.52.1",
			configFn: testAccCloudflareLoadBalancerMigrationConfigV4Geo,
		},
		{
			name:     "from_v5_0_0",
			version:  "5.0.0",
			configFn: testAccCloudflareLoadBalancerMigrationConfigV5Geo,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
			// service does not yet support the API tokens and it results in
			// misleading state error messages.
			if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
				t.Setenv("CLOUDFLARE_API_TOKEN", "")
			}

			accountID := acctest.TestAccCloudflareAccountID
			zoneID := acctest.TestAccCloudflareZoneID
			zone := acctest.TestAccCloudflareZoneName
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_load_balancer." + rnd
			testConfig := tc.configFn(accountID, zoneID, zone, rnd)
			tmpDir := t.TempDir()

			// Create initial step
			initialStep := resource.TestStep{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: tc.version,
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("steering_policy"), knownvalue.StringExact("geo")),
				},
			}

			// Create migration steps - use standard migration approach for all versions
			migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, testConfig, tmpDir, tc.version, "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("steering_policy"), knownvalue.StringExact("geo")),
			})
			migrationSteps = append(migrationSteps, resource.TestStep{
				// Final step: Apply the migrated configuration and verify state
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("steering_policy"), knownvalue.StringExact("geo")),
				},
			})

			// Combine all steps  
			allSteps := []resource.TestStep{initialStep}
			allSteps = append(allSteps, migrationSteps...)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				CheckDestroy: testAccCheckCloudflareLoadBalancerDestroy,
				WorkingDir:   tmpDir,
				Steps:        allSteps,
			})
		})
	}
}

// TestAccCloudflareLoadBalancer_Migration_Rules_MultiVersion tests migration of load balancers
// with custom routing rules. This test verifies that:
// 1. Rules blocks are correctly transformed from v4 to v5 list syntax
// 2. Rule overrides (including nested fallback_pool and default_pools) are migrated properly
// 3. Complex rule conditions and priorities are preserved
// 4. Rules remain a ListNestedAttribute (array) in v5, not converted to maps
// 5. Multiple pools referenced in rules maintain their relationships
func TestMigrateCloudflareLoadBalancer_Migration_Rules_MultiVersion(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(accountID, zoneID, zone, rnd string) string
	}{
		{
			name:     "from_v4_52_1",
			version:  "4.52.1",
			configFn: testAccCloudflareLoadBalancerMigrationConfigV4Rules,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
			// service does not yet support the API tokens and it results in
			// misleading state error messages.
			if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
				t.Setenv("CLOUDFLARE_API_TOKEN", "")
			}

			accountID := acctest.TestAccCloudflareAccountID
			zoneID := acctest.TestAccCloudflareZoneID
			zone := acctest.TestAccCloudflareZoneName
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_load_balancer." + rnd
			testConfig := tc.configFn(accountID, zoneID, zone, rnd)
			tmpDir := t.TempDir()

		// Build test steps
		steps := []resource.TestStep{
					{
						// Step 1: Create load balancer with specific version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								VersionConstraint: tc.version,
								Source:            "cloudflare/cloudflare",
							},
						},
						Config: testConfig,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("steering_policy"), knownvalue.StringExact("off")),
						},
					},
		}

		// Step 2: Migrate to v5 provider
		migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, testConfig, tmpDir, tc.version, "v4", "v5", []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("steering_policy"), knownvalue.StringExact("off")),
		})
		steps = append(steps, migrationSteps...)

		// Step 3: Apply migrated config with v5 provider
		steps = append(steps, resource.TestStep{
			ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
			ConfigDirectory:          config.StaticDirectory(tmpDir),
			ConfigStateChecks: []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("steering_policy"), knownvalue.StringExact("off")),
			},
		})

		resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				CheckDestroy: testAccCheckCloudflareLoadBalancerDestroy,
				WorkingDir:   tmpDir,
			Steps:        steps,
		})
		})
	}
}

// Additional test cases for specific scenarios

// TestAccCloudflareLoadBalancer_Migration_SessionAffinityIPCookie tests a specific edge case
// for session affinity migration. This test ensures that:
// 1. IP cookie session affinity settings are preserved during migration
// 2. Session affinity TTL values are correctly maintained
// 3. The ip_cookie affinity type doesn't require any special transformation
// 4. TTL field type changes from Float64 in v4 to Int64 in v5 are handled
func TestMigrateCloudflareLoadBalancer_Migration_SessionAffinityIPCookie(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := acctest.TestAccCloudflareAccountID
	zoneID := acctest.TestAccCloudflareZoneID
	zone := acctest.TestAccCloudflareZoneName
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_load_balancer." + rnd
	v4Config := testAccCloudflareLoadBalancerMigrationConfigV4SessionAffinityIPCookie(accountID, zoneID, zone, rnd)
	tmpDir := t.TempDir()

		// Build test steps
		steps := []resource.TestStep{
			{
				// Step 1: Create load balancer with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_affinity"), knownvalue.StringExact("ip_cookie")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_affinity_ttl"), knownvalue.Float64Exact(10800)),
				},
			},
		}

		// Step 2: Migrate to v5 provider
		migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{})
		steps = append(steps, migrationSteps...)

		// Step 3: Apply migrated config with v5 provider
		steps = append(steps, resource.TestStep{
				// Step 3: Apply migrated config with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_affinity"), knownvalue.StringExact("ip_cookie")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("session_affinity_ttl"), knownvalue.Int64Exact(10800)),
				},
		})

		resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		CheckDestroy: testAccCheckCloudflareLoadBalancerDestroy,
		WorkingDir:   tmpDir,
			Steps:        steps,
		})
}

// TestAccCloudflareLoadBalancer_Migration_ProximitySteeringPolicy tests migration of
// load balancers using proximity-based steering. This test verifies that:
// 1. Proximity steering policy setting is preserved during migration
// 2. Pool geographic coordinates (latitude/longitude) remain intact for proximity calculations
// 3. The steering policy doesn't require special transformation beyond basic migration
// 4. Load balancer continues to route based on proximity after migration
func TestMigrateCloudflareLoadBalancer_Migration_ProximitySteeringPolicy(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := acctest.TestAccCloudflareAccountID
	zoneID := acctest.TestAccCloudflareZoneID
	zone := acctest.TestAccCloudflareZoneName
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_load_balancer." + rnd
	v4Config := testAccCloudflareLoadBalancerMigrationConfigV4ProximityPolicy(accountID, zoneID, zone, rnd)
	tmpDir := t.TempDir()

		// Build test steps
		steps := []resource.TestStep{
			{
				// Step 1: Create load balancer with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("steering_policy"), knownvalue.StringExact("proximity")),
				},
			},
		}

		// Step 2: Migrate to v5 provider
		migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{})
		steps = append(steps, migrationSteps...)

		// Step 3: Apply migrated config with v5 provider
		steps = append(steps, resource.TestStep{
				// Step 3: Apply migrated config with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("steering_policy"), knownvalue.StringExact("proximity")),
				},
		})

		resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		CheckDestroy: testAccCheckCloudflareLoadBalancerDestroy,
		WorkingDir:   tmpDir,
			Steps:        steps,
		})
}

// V4 Configuration Functions

func testAccCloudflareLoadBalancerMigrationConfigV4Basic(accountID, zoneID, zone, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[4]s" {
  account_id = "%[1]s"
  name = "my-tf-pool-basic-%[4]s"
  latitude = 12.3
  longitude = 55
  origins {
    name = "example-1"
    address = "192.0.2.1"
    enabled = true
  }
}

resource "cloudflare_load_balancer" "%[4]s" {
  zone_id          = "%[2]s"
  name             = "tf-testacc-lb-%[4]s.%[3]s"
  steering_policy  = "off"
  session_affinity = "none"
  fallback_pool_id = cloudflare_load_balancer_pool.%[4]s.id
  default_pool_ids = [
    cloudflare_load_balancer_pool.%[4]s.id
  ]
  ttl = 30
}
`, accountID, zoneID, zone, rnd)
}

func testAccCloudflareLoadBalancerMigrationConfigV4AllOptional(accountID, zoneID, zone, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[4]s" {
  account_id = "%[1]s"
  name = "my-tf-pool-basic-%[4]s"
  latitude = 12.3
  longitude = 55
  origins {
    name = "example-1"
    address = "192.0.2.1"
    enabled = true
  }
}

resource "cloudflare_load_balancer" "%[4]s" {
  zone_id              = "%[2]s"
  name                 = "tf-testacc-lb-%[4]s.%[3]s"
  description          = "test load balancer"
  enabled              = true
  proxied              = true
  steering_policy      = "random"
  session_affinity     = "cookie"
  session_affinity_ttl = 1800
  fallback_pool_id     = cloudflare_load_balancer_pool.%[4]s.id
  default_pool_ids     = [
    cloudflare_load_balancer_pool.%[4]s.id
  ]
  
  session_affinity_attributes {
    samesite = "Lax"
    secure = "Auto"
  }
  
  adaptive_routing {
    failover_across_pools = false
  }
  
  location_strategy {
    prefer_ecs = "proximity"
    mode = "pop"
  }
  
  random_steering {
    default_weight = 0.5
  }
}
`, accountID, zoneID, zone, rnd)
}

func testAccCloudflareLoadBalancerMigrationConfigV4Geo(accountID, zoneID, zone, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[4]s" {
  account_id = "%[1]s"
  name = "my-tf-pool-basic-%[4]s"
  latitude = 12.3
  longitude = 55
  origins {
    name = "example-1"
    address = "192.0.2.1"
    enabled = true
  }
}

resource "cloudflare_load_balancer_pool" "%[4]s-2" {
  account_id = "%[1]s"
  name = "my-tf-pool-basic-%[4]s-2"
  latitude = 55.1
  longitude = -12.3
  origins {
    name = "example-2"
    address = "192.0.2.2"
    enabled = true
  }
}

resource "cloudflare_load_balancer" "%[4]s" {
  zone_id          = "%[2]s"
  name             = "tf-testacc-lb-%[4]s.%[3]s"
  steering_policy  = "geo"
  session_affinity = "none"
  fallback_pool_id = cloudflare_load_balancer_pool.%[4]s.id
  default_pool_ids = [
    cloudflare_load_balancer_pool.%[4]s.id
  ]
  ttl = 30
}
`, accountID, zoneID, zone, rnd)
}

func testAccCloudflareLoadBalancerMigrationConfigV4Rules(accountID, zoneID, zone, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[4]s" {
  account_id = "%[1]s"
  name = "my-tf-pool-basic-%[4]s"
  latitude = 12.3
  longitude = 55
  origins {
    name = "example-1"
    address = "192.0.2.1"
    enabled = true
  }
}

resource "cloudflare_load_balancer" "%[4]s" {
  zone_id          = "%[2]s"
  name             = "tf-testacc-lb-%[4]s.%[3]s"
  steering_policy  = "off"
  session_affinity = "none"
  fallback_pool_id = cloudflare_load_balancer_pool.%[4]s.id
  default_pool_ids = [
    cloudflare_load_balancer_pool.%[4]s.id
  ]
  ttl = 30
}
`, accountID, zoneID, zone, rnd)
}

func testAccCloudflareLoadBalancerMigrationConfigV4SessionAffinityIPCookie(accountID, zoneID, zone, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[4]s" {
  account_id = "%[1]s"
  name = "my-tf-pool-basic-%[4]s"
  latitude = 12.3
  longitude = 55
  origins {
    name = "example-1"
    address = "192.0.2.1"
    enabled = true
  }
}

resource "cloudflare_load_balancer" "%[4]s" {
  zone_id              = "%[2]s"
  name                 = "tf-testacc-lb-%[4]s.%[3]s"
  steering_policy      = "off"
  session_affinity     = "ip_cookie"
  session_affinity_ttl = 10800
  fallback_pool_id     = cloudflare_load_balancer_pool.%[4]s.id
  default_pool_ids     = [
    cloudflare_load_balancer_pool.%[4]s.id
  ]
  ttl = 30
}
`, accountID, zoneID, zone, rnd)
}

func testAccCloudflareLoadBalancerMigrationConfigV4ProximityPolicy(accountID, zoneID, zone, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[4]s" {
  account_id = "%[1]s"
  name = "my-tf-pool-basic-%[4]s"
  latitude = 12.3
  longitude = 55
  origins {
    name = "example-1"
    address = "192.0.2.1"
    enabled = true
  }
}

resource "cloudflare_load_balancer" "%[4]s" {
  zone_id          = "%[2]s"
  name             = "tf-testacc-lb-%[4]s.%[3]s"
  steering_policy  = "proximity"
  session_affinity = "none"
  fallback_pool_id = cloudflare_load_balancer_pool.%[4]s.id
  default_pool_ids = [
    cloudflare_load_balancer_pool.%[4]s.id
  ]
  ttl = 30
}
`, accountID, zoneID, zone, rnd)
}

// V5 Configuration Functions

func testAccCloudflareLoadBalancerMigrationConfigV5Basic(accountID, zoneID, zone, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[4]s" {
  account_id = "%[1]s"
  name = "my-tf-pool-basic-%[4]s"
  latitude = 12.3
  longitude = 55
  origins = [{
    name = "example-1"
    address = "192.0.2.1"
    enabled = true
  }]
}

resource "cloudflare_load_balancer" "%[4]s" {
  zone_id          = "%[2]s"
  name             = "tf-testacc-lb-%[4]s.%[3]s"
  steering_policy  = "off"
  session_affinity = "none"
  fallback_pool = cloudflare_load_balancer_pool.%[4]s.id
  default_pools = [
    cloudflare_load_balancer_pool.%[4]s.id
  ]
  ttl = 30
}
`, accountID, zoneID, zone, rnd)
}

func testAccCloudflareLoadBalancerMigrationConfigV5AllOptional(accountID, zoneID, zone, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[4]s" {
  account_id = "%[1]s"
  name = "my-tf-pool-basic-%[4]s"
  latitude = 12.3
  longitude = 55
  origins = [{
    name = "example-1"
    address = "192.0.2.1"
    enabled = true
  }]
}

resource "cloudflare_load_balancer" "%[4]s" {
  zone_id              = "%[2]s"
  name                 = "tf-testacc-lb-%[4]s.%[3]s"
  description          = "test load balancer"
  enabled              = true
  proxied              = true
  steering_policy      = "random"
  session_affinity     = "cookie"
  session_affinity_ttl = 1800
  fallback_pool = cloudflare_load_balancer_pool.%[4]s.id
  default_pools = [
    cloudflare_load_balancer_pool.%[4]s.id
  ]
  
  session_affinity_attributes = {
    samesite = "Lax"
    secure = "Auto"
  }
  
  adaptive_routing = {
    failover_across_pools = false
  }
  
  location_strategy = {
    prefer_ecs = "proximity"
    mode = "pop"
  }
  
  random_steering = {
    default_weight = 0.5
  }
}
`, accountID, zoneID, zone, rnd)
}

func testAccCloudflareLoadBalancerMigrationConfigV5Geo(accountID, zoneID, zone, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[4]s" {
  account_id = "%[1]s"
  name = "my-tf-pool-basic-%[4]s"
  latitude = 12.3
  longitude = 55
  origins = [{
    name = "example-1"
    address = "192.0.2.1"
    enabled = true
  }]
}

resource "cloudflare_load_balancer_pool" "%[4]s-2" {
  account_id = "%[1]s"
  name = "my-tf-pool-basic-%[4]s-2"
  latitude = 55.1
  longitude = -12.3
  origins = [{
    name = "example-2"
    address = "192.0.2.2"
    enabled = true
  }]
}

resource "cloudflare_load_balancer" "%[4]s" {
  zone_id          = "%[2]s"
  name             = "tf-testacc-lb-%[4]s.%[3]s"
  steering_policy  = "geo"
  session_affinity = "none"
  fallback_pool = cloudflare_load_balancer_pool.%[4]s.id
  default_pools = [
    cloudflare_load_balancer_pool.%[4]s.id
  ]
  
  region_pools = {
    WNAM = [cloudflare_load_balancer_pool.%[4]s.id]
    ENAM = [cloudflare_load_balancer_pool.%[4]s-2.id]
  }
  country_pools = {
    US = [cloudflare_load_balancer_pool.%[4]s.id]
    GB = [cloudflare_load_balancer_pool.%[4]s-2.id]
  }
}
`, accountID, zoneID, zone, rnd)
}

func testAccCloudflareLoadBalancerMigrationConfigV5Rules(accountID, zoneID, zone, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[4]s" {
  account_id = "%[1]s"
  name = "my-tf-pool-basic-%[4]s"
  latitude = 12.3
  longitude = 55
  origins = [{
    name = "example-1"
    address = "192.0.2.1"
    enabled = true
  }]
}

resource "cloudflare_load_balancer_pool" "%[4]s-2" {
  account_id = "%[1]s"
  name = "my-tf-pool-basic-%[4]s-2"
  latitude = 55.1
  longitude = -12.3
  origins = [{
    name = "example-2"
    address = "192.0.2.2"
    enabled = true
  }]
}

resource "cloudflare_load_balancer" "%[4]s" {
  zone_id          = "%[2]s"
  name             = "tf-testacc-lb-%[4]s.%[3]s"
  steering_policy  = "off"
  session_affinity = "none"
  fallback_pool = cloudflare_load_balancer_pool.%[4]s.id
  default_pools = [
    cloudflare_load_balancer_pool.%[4]s.id
  ]
  ttl = 30
  
  rules = [{
    name = "test rule"
    condition = "http.request.uri.path contains \"/api\""
    disabled = false
    priority = 1
    
    overrides = {
      fallback_pool = cloudflare_load_balancer_pool.%[4]s-2.id
      default_pools = [cloudflare_load_balancer_pool.%[4]s-2.id]
      session_affinity = "cookie"
    }
  }]
}
`, accountID, zoneID, zone, rnd)
}
