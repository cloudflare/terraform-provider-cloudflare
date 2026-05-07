package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

const defaultWorkersService = "mute-truth-fdb1"

func testWorkersService() string {
	if v := os.Getenv("WORKERS_CUSTOM_DOMAIN_TEST_SERVICE"); v != "" {
		return v
	}
	return defaultWorkersService
}

func testWorkersEnvironment() string {
	if v, ok := os.LookupEnv("WORKERS_CUSTOM_DOMAIN_TEST_ENVIRONMENT"); ok {
		return v
	}
	return "production"
}

var (
	currentProviderVersion = internal.PackageVersion
)

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_without_environment.tf
var v4WithoutEnvironmentConfig string

//go:embed testdata/v5_without_environment.tf
var v5WithoutEnvironmentConfig string

func TestMigrateWorkersCustomDomain_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func() string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func() string { return v4BasicConfig },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func() string { return v5BasicConfig },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if os.Getenv("TF_ACC") == "" {
				t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
			}

			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
			rnd := utils.GenerateRandomResourceName()
			hostname := fmt.Sprintf("%s.%s", rnd, zoneName)
			workersService := testWorkersService()
			workersEnvironment := testWorkersEnvironment()
			resourceName := "cloudflare_workers_custom_domain.test"
			tmpDir := t.TempDir()
			testConfig := fmt.Sprintf(tc.configFn(), accountID, zoneID, hostname, workersService, workersEnvironment)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
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

			checks := []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(hostname)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("service"), knownvalue.StringExact(workersService)),
			}
			if workersEnvironment != "" {
				checks = append(checks, statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("environment"), knownvalue.StringExact(workersEnvironment)))
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					// Step 1: Create with source provider version.
					firstStep,
					// Step 2: Run migration and verify state parity.
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, checks),
				},
			})
		})
	}
}

func TestMigrateWorkersCustomDomain_WithoutEnvironment(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func() string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func() string { return v4WithoutEnvironmentConfig },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func() string { return v5WithoutEnvironmentConfig },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if os.Getenv("TF_ACC") == "" {
				t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
			}

			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
			rnd := utils.GenerateRandomResourceName()
			hostname := fmt.Sprintf("%s.%s", rnd, zoneName)
			workersService := testWorkersService()
			resourceName := "cloudflare_workers_custom_domain.test"
			tmpDir := t.TempDir()
			testConfig := fmt.Sprintf(tc.configFn(), accountID, zoneID, hostname, workersService)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PostApplyPostRefresh: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
						},
					},
				}
			} else {
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
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					// Step 1: Create with source provider version.
					firstStep,
					// Step 2: Run migration and expect known environment replacement drift.
					migrationStepExpectEnvironmentReplacement(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, resourceName, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(hostname)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("service"), knownvalue.StringExact(workersService)),
					}),
				},
			})
		})
	}
}

func migrationStepExpectEnvironmentReplacement(t *testing.T, cfg string, tmpDir string, exactVersion string, sourceVersion string, targetVersion string, resourceName string, stateChecks []statecheck.StateCheck) resource.TestStep {
	return resource.TestStep{
		PreConfig: func() {
			acctest.WriteOutConfig(t, cfg, tmpDir)
			acctest.RunMigrationV2Command(t, cfg, tmpDir, sourceVersion, targetVersion)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		ConfigDirectory:          config.StaticDirectory(tmpDir),
		ExpectNonEmptyPlan:       false,
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				acctest.DebugNonEmptyPlan,
			},
			PostApplyPostRefresh: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
			},
		},
		ConfigStateChecks: stateChecks,
	}
}
