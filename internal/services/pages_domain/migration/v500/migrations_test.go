package v500_test

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zones"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Migration Test Configuration
//
// Version is read from LAST_V4_VERSION environment variable (set in .github/workflows/migration-tests.yml)
// - Last stable v4 release: default 4.52.5
// - Current v5 release: auto-updates with releases (internal.PackageVersion)
//
// Based on breaking changes analysis:
// - All breaking changes happened between 4.x and 5.0.0
// - No breaking changes between v5 releases (testing against latest v5)
// - Key changes: domain → name field rename

// Embed migration test configuration files
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_with_variables.tf
var v4WithVariablesConfig string

//go:embed testdata/v5_with_variables.tf
var v5WithVariablesConfig string

//go:embed testdata/v4_subdomain.tf
var v4SubdomainConfig string

//go:embed testdata/v5_subdomain.tf
var v5SubdomainConfig string

//go:embed testdata/v4_multiple.tf
var v4MultipleConfig string

//go:embed testdata/v5_multiple.tf
var v5MultipleConfig string

// getTestZoneDomain returns a zone domain from the test account for use in tests
func getTestZoneDomain(t *testing.T, accountID string) string {
	client := acctest.SharedClient()
	zonePage, err := client.Zones.List(context.Background(), zones.ZoneListParams{
		Account: cloudflare.F(zones.ZoneListParamsAccount{
			ID: cloudflare.F(accountID),
		}),
	})
	if err != nil || zonePage == nil || len(zonePage.Result) == 0 {
		t.Skip("No zones available in account for testing")
	}
	return zonePage.Result[0].Name
}

// TestMigratePagesDomain_Basic tests basic migration from v4 to v5
// Migrates domain field to name field
func TestMigratePagesDomain_Basic(t *testing.T) {
	testCases := []struct {
		name                 string
		version              string
		projectResourceName  string
		configFn             func(rnd, accountID, projectName, domainName, projectResourceName string) string
	}{
		{
			name:                "from_v4_latest",
			version:             acctest.GetLastV4Version(),
			projectResourceName: "project",
			configFn: func(rnd, accountID, projectName, domainName, projectResourceName string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, projectResourceName, accountID, projectName, domainName)
			},
		},
		{
			name:                "from_v5",
			version:             currentProviderVersion,
			projectResourceName: "v5_project",
			configFn: func(rnd, accountID, projectName, domainName, projectResourceName string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, projectResourceName, accountID, projectName, domainName)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			projectName := fmt.Sprintf("tf-acc-test-pages-project-%s", rnd)
			zoneDomain := getTestZoneDomain(t, accountID)
			domainName := fmt.Sprintf("%s.%s", rnd, zoneDomain)
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, projectName, domainName, tc.projectResourceName)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resourceName := fmt.Sprintf("cloudflare_pages_domain.%s", rnd)

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

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					// Step 2: Run migration and verify field rename (domain → name)
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("project_name"), knownvalue.StringExact(projectName)),
						// Verify domain field was renamed to name
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(domainName)),
					}),
				},
			})
		})
	}
}

// TestMigratePagesDomain_WithVariables tests migration with variable references
func TestMigratePagesDomain_WithVariables(t *testing.T) {
	testCases := []struct {
		name                string
		version             string
		projectResourceName string
		configFn            func(rnd, accountID, projectName, domainName, projectResourceName string) string
	}{
		{
			name:                "from_v4_latest",
			version:             acctest.GetLastV4Version(),
			projectResourceName: "project",
			configFn: func(rnd, accountID, projectName, domainName, projectResourceName string) string {
				return fmt.Sprintf(v4WithVariablesConfig, rnd, projectResourceName, accountID, projectName, domainName)
			},
		},
		{
			name:                "from_v5",
			version:             currentProviderVersion,
			projectResourceName: "v5_project",
			configFn: func(rnd, accountID, projectName, domainName, projectResourceName string) string {
				return fmt.Sprintf(v5WithVariablesConfig, rnd, projectResourceName, accountID, projectName, domainName)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			projectName := fmt.Sprintf("tf-acc-test-pages-project-%s", rnd)
			zoneDomain := getTestZoneDomain(t, accountID)
			domainName := fmt.Sprintf("%s.%s", rnd, zoneDomain)
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, projectName, domainName, tc.projectResourceName)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resourceName := fmt.Sprintf("cloudflare_pages_domain.%s", rnd)

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

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					// Step 2: Run migration and verify variable references preserved
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("project_name"), knownvalue.StringExact(projectName)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(domainName)),
					}),
				},
			})
		})
	}
}

// TestMigratePagesDomain_Subdomain tests migration with subdomain
func TestMigratePagesDomain_Subdomain(t *testing.T) {
	testCases := []struct {
		name                string
		version             string
		projectResourceName string
		configFn            func(rnd, accountID, projectName, domainName, projectResourceName string) string
	}{
		{
			name:                "from_v4_latest",
			version:             acctest.GetLastV4Version(),
			projectResourceName: "project",
			configFn: func(rnd, accountID, projectName, domainName, projectResourceName string) string {
				return fmt.Sprintf(v4SubdomainConfig, rnd, projectResourceName, accountID, projectName, domainName)
			},
		},
		{
			name:                "from_v5",
			version:             currentProviderVersion,
			projectResourceName: "v5_project",
			configFn: func(rnd, accountID, projectName, domainName, projectResourceName string) string {
				return fmt.Sprintf(v5SubdomainConfig, rnd, projectResourceName, accountID, projectName, domainName)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			projectName := fmt.Sprintf("tf-acc-test-pages-project-%s", rnd)
			zoneDomain := getTestZoneDomain(t, accountID)
			domainName := fmt.Sprintf("blog.%s.%s", rnd, zoneDomain)
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, projectName, domainName, tc.projectResourceName)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resourceName := fmt.Sprintf("cloudflare_pages_domain.%s", rnd)

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

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					// Step 2: Run migration and verify subdomain preserved
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("project_name"), knownvalue.StringExact(projectName)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(domainName)),
					}),
				},
			})
		})
	}
}

// TestMigratePagesDomain_MultipleResources tests migration with multiple domains
func TestMigratePagesDomain_MultipleResources(t *testing.T) {
	testCases := []struct {
		name                string
		version             string
		projectResourceName string
		configFn            func(rnd, accountID, projectName, domain1, domain2, projectResourceName string) string
	}{
		{
			name:                "from_v4_latest",
			version:             acctest.GetLastV4Version(),
			projectResourceName: "project",
			configFn: func(rnd, accountID, projectName, domain1, domain2, projectResourceName string) string {
				return fmt.Sprintf(v4MultipleConfig, rnd, projectResourceName, accountID, projectName, domain1, domain2)
			},
		},
		{
			name:                "from_v5",
			version:             currentProviderVersion,
			projectResourceName: "v5_project",
			configFn: func(rnd, accountID, projectName, domain1, domain2, projectResourceName string) string {
				return fmt.Sprintf(v5MultipleConfig, rnd, projectResourceName, accountID, projectName, domain1, domain2)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			projectName := fmt.Sprintf("tf-acc-test-pages-project-%s", rnd)
			zoneDomain := getTestZoneDomain(t, accountID)
			domain1 := fmt.Sprintf("%s-1.%s", rnd, zoneDomain)
			domain2 := fmt.Sprintf("%s-2.%s", rnd, zoneDomain)
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, projectName, domain1, domain2, tc.projectResourceName)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource1Name := fmt.Sprintf("cloudflare_pages_domain.%s_domain1", rnd)
			resource2Name := fmt.Sprintf("cloudflare_pages_domain.%s_domain2", rnd)

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

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					// Step 2: Run migration and verify both domains migrated
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						// Verify first domain
						statecheck.ExpectKnownValue(resource1Name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resource1Name, tfjsonpath.New("project_name"), knownvalue.StringExact(projectName)),
						statecheck.ExpectKnownValue(resource1Name, tfjsonpath.New("name"), knownvalue.StringExact(domain1)),
						// Verify second domain
						statecheck.ExpectKnownValue(resource2Name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resource2Name, tfjsonpath.New("project_name"), knownvalue.StringExact(projectName)),
						statecheck.ExpectKnownValue(resource2Name, tfjsonpath.New("name"), knownvalue.StringExact(domain2)),
					}),
				},
			})
		})
	}
}
