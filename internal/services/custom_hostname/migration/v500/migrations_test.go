package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

const currentProviderVersion = internal.PackageVersion

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

func TestMigrateCustomHostnameBasic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, domain string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID, domain string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, zoneID, rnd, domain, rnd, domain, rnd, domain)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID, domain string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, zoneID, rnd, domain, rnd, domain, rnd, domain)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_custom_hostname." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID, domain)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
					ExpectNonEmptyPlan:       true,
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

			migrationStep := acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(fmt.Sprintf("%s.%s", rnd, domain))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ssl").AtMapKey("method"), knownvalue.StringExact("txt")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ssl").AtMapKey("type"), knownvalue.StringExact("dv")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ssl").AtMapKey("settings").AtMapKey("tls_1_3"), knownvalue.StringExact("on")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("custom_metadata").AtMapKey("environment"), knownvalue.StringExact("migration")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("custom_origin_server"), knownvalue.StringExact(fmt.Sprintf("origin-%s.%s", rnd, domain))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("custom_origin_sni"), knownvalue.StringExact(fmt.Sprintf("origin-%s.%s", rnd, domain))),
			})

			if tc.version == currentProviderVersion {
				migrationStep.ExpectNonEmptyPlan = true
				migrationStep.ConfigPlanChecks = resource.ConfigPlanChecks{}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					migrationStep,
				},
			})
		})
	}
}
