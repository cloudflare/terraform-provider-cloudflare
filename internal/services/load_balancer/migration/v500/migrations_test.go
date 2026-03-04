package v500_test

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"testing"

	cfold "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Embed migration test configuration files
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_with_nested.tf
var v4WithNestedConfig string

//go:embed testdata/v5_with_nested.tf
var v5WithNestedConfig string

//go:embed testdata/v4_with_pools.tf
var v4WithPoolsConfig string

//go:embed testdata/v5_with_pools.tf
var v5WithPoolsConfig string

// TestMigrateLoadBalancer_V4ToV5_Basic tests basic field migrations with DUAL test cases.
// This verifies:
// 1. Field renames: fallback_pool_id → fallback_pool, default_pool_ids → default_pools
// 2. Type conversions: ttl and session_affinity_ttl (Int64 → Float64)
// 3. Both migration paths: from v4 and from v5
func TestMigrateLoadBalancer_V4ToV5_Basic(t *testing.T) {
	legacyProviderVersion := os.Getenv("LAST_V4_VERSION")
	if legacyProviderVersion == "" {
		legacyProviderVersion = "4.52.5" // Default to last known v4
	}

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, fallbackPool, defaultPool string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: legacyProviderVersion,
			configFn: func(rnd, zoneID, fallbackPool, defaultPool string) string {
				domain := os.Getenv("CLOUDFLARE_DOMAIN")
				lbName := fmt.Sprintf("tf-lb-%s.%s", rnd, domain)
				return fmt.Sprintf(v4BasicConfig, rnd, zoneID, lbName, fallbackPool, defaultPool)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, zoneID, fallbackPool, defaultPool string) string {
				domain := os.Getenv("CLOUDFLARE_DOMAIN")
				lbName := fmt.Sprintf("tf-lb-%s.%s", rnd, domain)
				return fmt.Sprintf(v5BasicConfig, rnd, zoneID, lbName, fallbackPool, defaultPool)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if os.Getenv("migration mode") != "1" {
				t.Skip("migration mode must be set to 1 to run migration tests")
			}

			// Setup
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()

			// Create pools for the load balancer
			fallbackPoolID, defaultPoolID := createTestPools(t, accountID)

			testConfig := tc.configFn(rnd, zoneID, fallbackPoolID, defaultPoolID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with specific provider version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify renamed fields
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer."+rnd,
								tfjsonpath.New("fallback_pool"),
								knownvalue.StringExact(fallbackPoolID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer."+rnd,
								tfjsonpath.New("default_pools"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.StringExact(defaultPoolID),
								}),
							),
							// Verify type conversions (Int64 → Float64)
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer."+rnd,
								tfjsonpath.New("ttl"),
								knownvalue.Float64Exact(30),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer."+rnd,
								tfjsonpath.New("session_affinity_ttl"),
								knownvalue.Float64Exact(1800),
							),
							// Verify basic fields preserved
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer."+rnd,
								tfjsonpath.New("zone_id"),
								knownvalue.StringExact(zoneID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer."+rnd,
								tfjsonpath.New("description"),
								knownvalue.StringExact("Load balancer for tf-migrate test"),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateLoadBalancer_V4ToV5_WithNested tests nested object transformations.
// This verifies:
// 1. session_affinity_attributes: Array[0] → single object
// 2. adaptive_routing: Array[0] → single object
// 3. location_strategy: Array[0] → single object
// 4. random_steering: Array[0] → single object
// 5. Type conversion for drain_duration (Int64 → Float64)
func TestMigrateLoadBalancer_V4ToV5_WithNested(t *testing.T) {
	legacyProviderVersion := os.Getenv("LAST_V4_VERSION")
	if legacyProviderVersion == "" {
		legacyProviderVersion = "4.52.5"
	}

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, fallbackPool, defaultPool string) string
	}{
		{
			name:    "from_v4_latest",
			version: legacyProviderVersion,
			configFn: func(rnd, zoneID, fallbackPool, defaultPool string) string {
				domain := os.Getenv("CLOUDFLARE_DOMAIN")
				lbName := fmt.Sprintf("tf-lb-nested-%s.%s", rnd, domain)
				return fmt.Sprintf(v4WithNestedConfig, rnd, zoneID, lbName, fallbackPool, defaultPool)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID, fallbackPool, defaultPool string) string {
				domain := os.Getenv("CLOUDFLARE_DOMAIN")
				lbName := fmt.Sprintf("tf-lb-nested-%s.%s", rnd, domain)
				return fmt.Sprintf(v5WithNestedConfig, rnd, zoneID, lbName, fallbackPool, defaultPool)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if os.Getenv("migration mode") != "1" {
				t.Skip("migration mode must be set to 1 to run migration tests")
			}

			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()

			fallbackPoolID, defaultPoolID := createTestPools(t, accountID)
			testConfig := tc.configFn(rnd, zoneID, fallbackPoolID, defaultPoolID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify session_affinity_attributes transformation
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer."+rnd,
								tfjsonpath.New("session_affinity_attributes").AtMapKey("samesite"),
								knownvalue.StringExact("Lax"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer."+rnd,
								tfjsonpath.New("session_affinity_attributes").AtMapKey("secure"),
								knownvalue.StringExact("Always"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer."+rnd,
								tfjsonpath.New("session_affinity_attributes").AtMapKey("drain_duration"),
								knownvalue.Float64Exact(100),
							),
							// Verify adaptive_routing transformation
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer."+rnd,
								tfjsonpath.New("adaptive_routing").AtMapKey("failover_across_pools"),
								knownvalue.Bool(true),
							),
							// Verify location_strategy transformation
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer."+rnd,
								tfjsonpath.New("location_strategy").AtMapKey("prefer_ecs"),
								knownvalue.StringExact("proximity"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer."+rnd,
								tfjsonpath.New("location_strategy").AtMapKey("mode"),
								knownvalue.StringExact("pop"),
							),
							// Verify random_steering transformation
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer."+rnd,
								tfjsonpath.New("random_steering").AtMapKey("default_weight"),
								knownvalue.Float64Exact(0.5),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateLoadBalancer_V4ToV5_WithPools tests pool block transformations.
// This verifies:
// 1. region_pools: Set of blocks → map
// 2. pop_pools: Set of blocks → map
// 3. country_pools: Set of blocks → map
func TestMigrateLoadBalancer_V4ToV5_WithPools(t *testing.T) {
	legacyProviderVersion := os.Getenv("LAST_V4_VERSION")
	if legacyProviderVersion == "" {
		legacyProviderVersion = "4.52.5"
	}

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, fallbackPool, defaultPool, regionPool, popPool, countryPool string) string
	}{
		{
			name:    "from_v4_latest",
			version: legacyProviderVersion,
			configFn: func(rnd, zoneID, fallbackPool, defaultPool, regionPool, popPool, countryPool string) string {
				domain := os.Getenv("CLOUDFLARE_DOMAIN")
				lbName := fmt.Sprintf("tf-lb-pools-%s.%s", rnd, domain)
				return fmt.Sprintf(v4WithPoolsConfig, rnd, zoneID, lbName,
					fallbackPool, defaultPool, regionPool, popPool, countryPool)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID, fallbackPool, defaultPool, regionPool, popPool, countryPool string) string {
				domain := os.Getenv("CLOUDFLARE_DOMAIN")
				lbName := fmt.Sprintf("tf-lb-pools-%s.%s", rnd, domain)
				return fmt.Sprintf(v5WithPoolsConfig, rnd, zoneID, lbName,
					fallbackPool, defaultPool, regionPool, popPool, countryPool)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if os.Getenv("migration mode") != "1" {
				t.Skip("migration mode must be set to 1 to run migration tests")
			}

			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()

			fallbackPoolID, defaultPoolID := createTestPools(t, accountID)
			regionPoolID := createTestPool(t, accountID, fmt.Sprintf("tf-lb-region-%s", rnd))
			popPoolID := createTestPool(t, accountID, fmt.Sprintf("tf-lb-pop-%s", rnd))
			countryPoolID := createTestPool(t, accountID, fmt.Sprintf("tf-lb-country-%s", rnd))

			testConfig := tc.configFn(rnd, zoneID, fallbackPoolID, defaultPoolID, regionPoolID, popPoolID, countryPoolID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify region_pools transformation (block → map)
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer."+rnd,
								tfjsonpath.New("region_pools").AtMapKey("WNAM"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.StringExact(regionPoolID),
								}),
							),
							// Verify pop_pools transformation (block → map)
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer."+rnd,
								tfjsonpath.New("pop_pools").AtMapKey("LAX"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.StringExact(popPoolID),
								}),
							),
							// Verify country_pools transformation (block → map)
							statecheck.ExpectKnownValue(
								"cloudflare_load_balancer."+rnd,
								tfjsonpath.New("country_pools").AtMapKey("US"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.StringExact(countryPoolID),
								}),
							),
						},
					),
				},
			})
		})
	}
}

// createTestPool creates a single test pool for use in load balancer tests.
// The pool is automatically cleaned up using t.Cleanup().
func createTestPool(t *testing.T, accountID, poolName string) string {
	t.Helper()

	client, err := acctest.SharedV1Client()
	if err != nil {
		t.Fatalf("Failed to create Cloudflare client: %v", err)
	}

	ctx := context.Background()

	pool, err := client.CreateLoadBalancerPool(ctx, cfold.AccountIdentifier(accountID), cfold.CreateLoadBalancerPoolParams{
		LoadBalancerPool: cfold.LoadBalancerPool{
			Name:    poolName,
			Enabled: true,
			Origins: []cfold.LoadBalancerOrigin{
				{
					Name:    fmt.Sprintf("%s-origin", poolName),
					Address: "192.0.2.1",
					Enabled: true,
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("Failed to create pool %s: %v", poolName, err)
	}

	// Register cleanup to delete pool when test completes
	t.Cleanup(func() {
		_ = client.DeleteLoadBalancerPool(context.Background(), cfold.AccountIdentifier(accountID), pool.ID)
	})

	return pool.ID
}

// createTestPools creates two test pools for use in load balancer tests.
// Returns (fallbackPoolID, defaultPoolID).
// Pools are automatically cleaned up using t.Cleanup().
func createTestPools(t *testing.T, accountID string) (string, string) {
	t.Helper()

	rnd := utils.GenerateRandomResourceName()
	fallbackPoolID := createTestPool(t, accountID, fmt.Sprintf("tf-lb-fallback-%s", rnd))
	defaultPoolID := createTestPool(t, accountID, fmt.Sprintf("tf-lb-default-%s", rnd))

	return fallbackPoolID, defaultPoolID
}
