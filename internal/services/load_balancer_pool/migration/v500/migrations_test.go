package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

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

// Embed test configs
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_full.tf
var v4FullConfig string

//go:embed testdata/v5_full.tf
var v5FullConfig string

//go:embed testdata/v4_check_regions.tf
var v4CheckRegionsConfig string

//go:embed testdata/v5_check_regions.tf
var v5CheckRegionsConfig string

// TestMigrateLoadBalancerPool_V4ToV5_Basic tests basic field migrations with DUAL test cases
func TestMigrateLoadBalancerPool_V4ToV5_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test setup
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := rnd // Name suffix for pool
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name)
			// Infer source/target versions from test version
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)
			resourceName := fmt.Sprintf("cloudflare_load_balancer_pool.%s", rnd)

			// For v5 tests, use local provider; for v4 tests, use external provider
			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				// Use local v5 provider (will create state)
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				// Use external v4 provider (will create version=0 state)
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify required fields
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("account_id"),
								knownvalue.StringExact(accountID),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("name"),
								knownvalue.StringExact(fmt.Sprintf("my-tf-pool-basic-%s", name)),
							),
							// Verify origins transformed correctly (Set → List)
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("origins"),
								knownvalue.ListSizeExact(1),
							),
							// Verify id exists (computed)
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("id"),
								knownvalue.NotNull(),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateLoadBalancerPool_V4ToV5_FullConfig tests all optional attributes with complex transformations
// Note: Only testing from_v4_latest because from_v5 with load_shedding/origin_steering creates a schema
// conflict (v5 state uses objects, but v4 schema expects arrays). This is not a real-world scenario
// since v5 users already have correct state format.
func TestMigrateLoadBalancerPool_V4ToV5_FullConfig(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name, domain string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name, domain string) string {
				return fmt.Sprintf(v4FullConfig, rnd, accountID, name, domain, domain)
			},
		},
		// Skipping from_v5 test case due to schema version collision:
		// Both v4 and v5 use schema_version=0, but have different formats for load_shedding/origin_steering
		// v5 users already have correct state format and don't need migration
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			rnd := utils.GenerateRandomResourceName()
			name := rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name, domain)
			// Infer source/target versions from test version
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)
			resourceName := fmt.Sprintf("cloudflare_load_balancer_pool.%s", rnd)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
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
							// Basic fields
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("name"),
								knownvalue.StringExact(fmt.Sprintf("my-tf-pool-full-%s", name)),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("enabled"),
								knownvalue.Bool(false),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("minimum_origins"),
								knownvalue.Int64Exact(2),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("description"),
								knownvalue.StringExact("tfacc-fully-specified"),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("latitude"),
								knownvalue.Float64Exact(12.3),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("longitude"),
								knownvalue.Float64Exact(55),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("notification_email"),
								knownvalue.StringExact("someone@example.com"),
							),
							// Origins: Set → List (2 origins)
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("origins"),
								knownvalue.ListSizeExact(2),
							),
							// Origins.header: Complex transformation {header, values} → {host: []}
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("origins").AtSliceIndex(0).AtMapKey("header").AtMapKey("host"),
								knownvalue.ListSizeExact(1),
							),
							// Load shedding: Array[0] → Object
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("load_shedding").AtMapKey("default_percent"),
								knownvalue.Float64Exact(55),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("load_shedding").AtMapKey("default_policy"),
								knownvalue.StringExact("random"),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("load_shedding").AtMapKey("session_percent"),
								knownvalue.Float64Exact(12),
							),
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("load_shedding").AtMapKey("session_policy"),
								knownvalue.StringExact("hash"),
							),
							// Origin steering: Array[0] → Object
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("origin_steering").AtMapKey("policy"),
								knownvalue.StringExact("random"),
							),
							// Check regions: Set → List
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("check_regions"),
								knownvalue.ListSizeExact(1),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateLoadBalancerPool_V4ToV5_CheckRegions tests Set → List transformation for check_regions
func TestMigrateLoadBalancerPool_V4ToV5_CheckRegions(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v4CheckRegionsConfig, rnd, accountID, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, name string) string {
				return fmt.Sprintf(v5CheckRegionsConfig, rnd, accountID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			name := rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, name)
			// Infer source/target versions from test version
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)
			resourceName := fmt.Sprintf("cloudflare_load_balancer_pool.%s", rnd)

			// For v5 tests, use local provider; for v4 tests, use external provider
			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				// Use local v5 provider (will create state)
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				// Use external v4 provider (will create version=0 state)
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("name"),
								knownvalue.StringExact(fmt.Sprintf("my-tf-pool-regions-%s", name)),
							),
							// Check regions: Set → List (3 regions)
							statecheck.ExpectKnownValue(
								resourceName,
								tfjsonpath.New("check_regions"),
								knownvalue.ListSizeExact(3),
							),
						},
					),
				},
			})
		})
	}
}
