package load_balancer_pool_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMigrateCloudflareLoadBalancerPool_Basic_MultiVersion(t *testing.T) {
	// Based on breaking changes analysis:
	// - All breaking changes happened between 4.x and 5.0.0
	// - No breaking changes between v5 releases
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:     "from_v4_52_1", // Last v4 release
			version:  "4.52.1",
			configFn: testAccCloudflareLoadBalancerPoolMigrationConfigV4Basic,
		},
		{
			name:     "from_v5_0_0", // First v5 release, after breaking changes
			version:  "5.0.0",
			configFn: testAccCloudflareLoadBalancerPoolMigrationConfigV5Basic, // v5 uses list syntax for origins
		},
		{
			name:     "from_v5_7_1", // Recent v5 release
			version:  "5.7.1",
			configFn: testAccCloudflareLoadBalancerPoolMigrationConfigV5Basic, // v5 uses list syntax for origins
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := fmt.Sprintf("cloudflare_load_balancer_pool.%s", rnd)
			testConfig := tc.configFn(rnd, accountID)
			tmpDir := t.TempDir()

		// Build test steps
		steps := []resource.TestStep{
					{
						// Step 1: Create pool with specific version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("my-tf-pool-basic-%s", rnd))),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("minimum_origins"), knownvalue.Int64Exact(1)),
						},
					},
		}

		// Step 2: Migrate to v5 provider
		migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, testConfig, tmpDir, tc.version, "v4", "v5", []statecheck.StateCheck{
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("my-tf-pool-basic-%s", rnd))),
		})
		steps = append(steps, migrationSteps...)

		resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
			Steps:        steps,
		})
		})
	}
}

func TestMigrateCloudflareLoadBalancerPool_AllOptionalAttributes_MultiVersion(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, domain string) string
	}{
		{
			name:     "from_v4_52_1",
			version:  "4.52.1",
			configFn: testAccCloudflareLoadBalancerPoolMigrationConfigV4AllOptional,
		},
		{
			name:     "from_v5_0_0",
			version:  "5.0.0",
			configFn: testAccCloudflareLoadBalancerPoolMigrationConfigV5AllOptional,
		},
		{
			name:     "from_v5_7_1",
			version:  "5.7.1",
			configFn: testAccCloudflareLoadBalancerPoolMigrationConfigV5AllOptional,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			rnd := utils.GenerateRandomResourceName()
			resourceName := fmt.Sprintf("cloudflare_load_balancer_pool.%s", rnd)
			testConfig := tc.configFn(rnd, accountID, domain)
			tmpDir := t.TempDir()

		// Build test steps
		steps := []resource.TestStep{
					{
						// Step 1: Create pool with specific version with all optional attributes
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("my-tf-pool-full-%s", rnd))),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("minimum_origins"), knownvalue.Int64Exact(2)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("tfacc-fully-specified")),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("latitude"), knownvalue.Float64Exact(12.3)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("longitude"), knownvalue.Float64Exact(55)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("notification_email"), knownvalue.StringExact("someone@example.com")),
						},
					},
		}

		// Step 2: Migrate to v5 provider
		migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, testConfig, tmpDir, tc.version, "v4", "v5", []statecheck.StateCheck{
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("my-tf-pool-full-%s", rnd))),
		})
		steps = append(steps, migrationSteps...)

		resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				WorkingDir: tmpDir,
			Steps:        steps,
		})
		})
	}
}

func TestMigrateCloudflareLoadBalancerPool_OriginSteering_MultiVersion(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, policy string) string
	}{
		{
			name:     "from_v4_52_1",
			version:  "4.52.1",
			configFn: testAccCloudflareLoadBalancerPoolMigrationConfigV4OriginSteering,
		},
		{
			name:     "from_v5_0_0",
			version:  "5.0.0",
			configFn: testAccCloudflareLoadBalancerPoolMigrationConfigV5OriginSteering,
		},
		{
			name:     "from_v5_7_1",
			version:  "5.7.1",
			configFn: testAccCloudflareLoadBalancerPoolMigrationConfigV5OriginSteering,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := fmt.Sprintf("cloudflare_load_balancer_pool.%s", rnd)
			testConfig := tc.configFn(rnd, accountID, "least_connections")
			tmpDir := t.TempDir()

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create pool with specific version with origin steering
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("my-tf-pool-steering-%s", rnd))),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						},
					},
					{
						// Step 2: Migrate to latest provider
						PreConfig: func() {
							// Run the migration command
							acctest.WriteOutConfig(t, testConfig, tmpDir)
							acctest.RunMigrationCommand(t, testConfig, tmpDir)
						},
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						// Verify no changes needed after migration
						ConfigPlanChecks: resource.ConfigPlanChecks{
							PreApply: []plancheck.PlanCheck{
								plancheck.ExpectEmptyPlan(),
							},
						},
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("my-tf-pool-steering-%s", rnd))),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							// Verify origin_steering structure changed correctly
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin_steering").AtMapKey("policy"), knownvalue.StringExact("least_connections")),
						},
					},
					{
						// Step 3 - Import and verify
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ResourceName:             resourceName,
						ImportState:              true,
						ImportStateIdFunc: func(s *terraform.State) (string, error) {
							rs, ok := s.RootModule().Resources[resourceName]
							if !ok {
								return "", fmt.Errorf("resource not found: %s", resourceName)
							}
							return fmt.Sprintf("%s/%s", accountID, rs.Primary.ID), nil
						},
						ImportStateVerify: true,
					},
				},
			})
		})
	}
}

func TestMigrateCloudflareLoadBalancerPool_CheckRegions_MultiVersion(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID string) string
	}{
		{
			name:     "from_v4_52_1",
			version:  "4.52.1",
			configFn: testAccCloudflareLoadBalancerPoolMigrationConfigV4CheckRegions,
		},
		{
			name:     "from_v5_0_0",
			version:  "5.0.0",
			configFn: testAccCloudflareLoadBalancerPoolMigrationConfigV5CheckRegions,
		},
		{
			name:     "from_v5_7_1",
			version:  "5.7.1",
			configFn: testAccCloudflareLoadBalancerPoolMigrationConfigV5CheckRegions,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := fmt.Sprintf("cloudflare_load_balancer_pool.%s", rnd)
			testConfig := tc.configFn(rnd, accountID)
			tmpDir := t.TempDir()

		// Build test steps
		steps := []resource.TestStep{
					{
						// Step 1: Create pool with specific version with multiple check regions
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("my-tf-pool-regions-%s", rnd))),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						},
					},
		}

		// Step 2: Migrate to v5 provider
		migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, testConfig, tmpDir, tc.version, "v4", "v5", []statecheck.StateCheck{
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("my-tf-pool-regions-%s", rnd))),
		})
		steps = append(steps, migrationSteps...)

		resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
			Steps:        steps,
		})
		})
	}
}

func testAccCloudflareLoadBalancerPoolMigrationConfigV4Basic(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[1]s" {
  account_id = "%[2]s"
  name = "my-tf-pool-basic-%[1]s"
  
  origins {
    name = "example-1"
    address = "192.0.2.1"
    enabled = true
  }
}
`, rnd, accountID)
}

func testAccCloudflareLoadBalancerPoolMigrationConfigV5Basic(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[1]s" {
  account_id = "%[2]s"
  name = "my-tf-pool-basic-%[1]s"
  
  origins = [{
    name = "example-1"
    address = "192.0.2.1"
    enabled = true
  }]
}
`, rnd, accountID)
}

func testAccCloudflareLoadBalancerPoolMigrationConfigV4AllOptional(rnd, accountID, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[1]s" {
  account_id = "%[2]s"
  name = "my-tf-pool-full-%[1]s"
  
  origins {
    name = "example-1"
    address = "192.0.2.1"
    enabled = false
    weight = 1.0
    header {
      header = "Host"
      values = ["test1.%[3]s"]
    }
  }
  
  origins {
    name = "example-2"
    address = "192.0.2.2"
    weight = 0.5
    header {
      header = "Host"
      values = ["test2.%[3]s"]
    }
  }
  
  load_shedding {
    default_percent = 55
    default_policy = "random"
    session_percent = 12
    session_policy = "hash"
  }
  
  latitude = 12.3
  longitude = 55
  
  origin_steering {
    policy = "random"
  }
  
  check_regions = ["WEU"]
  description = "tfacc-fully-specified"
  enabled = false
  minimum_origins = 2
  notification_email = "someone@example.com"
}
`, rnd, accountID, domain)
}

func testAccCloudflareLoadBalancerPoolMigrationConfigV5AllOptional(rnd, accountID, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[1]s" {
  account_id = "%[2]s"
  name = "my-tf-pool-full-%[1]s"
  
  origins = [
    {
      name = "example-1"
      address = "192.0.2.1"
      enabled = false
      weight = 1.0
      header = {
        host = ["test1.%[3]s"]
      }
    },
    {
      name = "example-2"
      address = "192.0.2.2"
      weight = 0.5
      header = {
        host = ["test2.%[3]s"]
      }
    }
  ]
  
  load_shedding = {
    default_percent = 55
    default_policy = "random"
    session_percent = 12
    session_policy = "hash"
  }
  
  latitude = 12.3
  longitude = 55
  
  origin_steering = {
    policy = "random"
  }
  
  check_regions = ["WEU"]
  description = "tfacc-fully-specified"
  enabled = false
  minimum_origins = 2
  notification_email = "someone@example.com"
}
`, rnd, accountID, domain)
}

func testAccCloudflareLoadBalancerPoolMigrationConfigV4OriginSteering(rnd, accountID, policy string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[1]s" {
  account_id = "%[2]s"
  name = "my-tf-pool-steering-%[1]s"
  
  origins {
    name = "example-1"
    address = "192.0.2.1"
    weight = 0.8
  }
  
  origins {
    name = "example-2"
    address = "192.0.2.2"
    weight = 0.2
  }
  
  origin_steering {
    policy = "%[3]s"
  }
}
`, rnd, accountID, policy)
}

func testAccCloudflareLoadBalancerPoolMigrationConfigV5OriginSteering(rnd, accountID, policy string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[1]s" {
  account_id = "%[2]s"
  name = "my-tf-pool-steering-%[1]s"
  
  origins = [
    {
      name = "example-1"
      address = "192.0.2.1"
      weight = 0.8
    },
    {
      name = "example-2"
      address = "192.0.2.2"
      weight = 0.2
    }
  ]
  
  origin_steering = {
    policy = "%[3]s"
  }
}
`, rnd, accountID, policy)
}

func testAccCloudflareLoadBalancerPoolMigrationConfigV4CheckRegions(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[1]s" {
  account_id = "%[2]s"
  name = "my-tf-pool-regions-%[1]s"
  
  origins {
    name = "example-1"
    address = "192.0.2.1"
  }
  
  check_regions = ["WEU", "ENAM", "WNAM"]
}
`, rnd, accountID)
}

func testAccCloudflareLoadBalancerPoolMigrationConfigV5CheckRegions(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[1]s" {
  account_id = "%[2]s"
  name = "my-tf-pool-regions-%[1]s"
  
  origins = [{
    name = "example-1"
    address = "192.0.2.1"
  }]
  
  check_regions = ["WEU", "ENAM", "WNAM"]
}
`, rnd, accountID)
}

func TestMigrateCloudflareLoadBalancerPool_DynamicOrigins_MultiVersion(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, domain string) string
	}{
		{
			name:     "from_v4_52_1",
			version:  "4.52.1",
			configFn: testAccCloudflareLoadBalancerPoolMigrationConfigV4DynamicOrigins,
		},
		{
			name:     "from_v5_0_0",
			version:  "5.0.0",
			configFn: testAccCloudflareLoadBalancerPoolMigrationConfigV5DynamicOrigins,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			rnd := utils.GenerateRandomResourceName()
			resourceName := fmt.Sprintf("cloudflare_load_balancer_pool.%s", rnd)
			testConfig := tc.configFn(rnd, accountID, domain)
			tmpDir := t.TempDir()

		// Build test steps
		steps := []resource.TestStep{
					{
						// Step 1: Create pool with specific version using dynamic origins block
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("my-tf-pool-dynamic-%s", rnd))),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							// Check that origins exist (will be 3 based on our config)
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origins"), knownvalue.ListSizeExact(3)),
						},
					},
		}

		// Step 2: Migrate to v5 provider
		migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, testConfig, tmpDir, tc.version, "v4", "v5", []statecheck.StateCheck{
			statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("my-tf-pool-dynamic-%s", rnd))),
		})
		steps = append(steps, migrationSteps...)

		resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				WorkingDir: tmpDir,
			Steps:        steps,
		})
		})
	}
}

func testAccCloudflareLoadBalancerPoolMigrationConfigV4DynamicOrigins(rnd, accountID, domain string) string {
	return fmt.Sprintf(`
locals {
  origin_configs = [
    {
      name    = "origin-0"
      address = "192.0.2.1"
      host    = "test0.%[3]s"
    },
    {
      name    = "origin-1"
      address = "192.0.2.2"
      host    = "test1.%[3]s"
    },
    {
      name    = "origin-2"
      address = "192.0.2.3"
      host    = "test2.%[3]s"
    }
  ]
}

resource "cloudflare_load_balancer_pool" "%[1]s" {
  account_id = "%[2]s"
  name = "my-tf-pool-dynamic-%[1]s"
  
  dynamic "origins" {
    for_each = local.origin_configs
    content {
      name    = origins.value.name
      address = origins.value.address
      enabled = true
      
      header {
        header = "Host"
        values = [origins.value.host]
      }
    }
  }
}
`, rnd, accountID, domain)
}

func testAccCloudflareLoadBalancerPoolMigrationConfigV5DynamicOrigins(rnd, accountID, domain string) string {
	return fmt.Sprintf(`
locals {
  origin_configs = [
    {
      name    = "origin-0"
      address = "192.0.2.1"
      host    = "test0.%[3]s"
    },
    {
      name    = "origin-1"
      address = "192.0.2.2"
      host    = "test1.%[3]s"
    },
    {
      name    = "origin-2"
      address = "192.0.2.3"
      host    = "test2.%[3]s"
    }
  ]
}

resource "cloudflare_load_balancer_pool" "%[1]s" {
  account_id = "%[2]s"
  name = "my-tf-pool-dynamic-%[1]s"
  
  origins = [for value in local.origin_configs : {
    address = value.address
    enabled = true
    header  = { host = [value.host] }
    name    = value.name
  }]
}
`, rnd, accountID, domain)
}
