package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

//go:embed testdata/v4_basic.tf
var testAccCloudflareSnippetMigrationConfigBasic string

//go:embed testdata/v4_multiple_files.tf
var testAccCloudflareSnippetMigrationConfigMultipleFiles string

//go:embed testdata/v4_complex_content.tf
var testAccCloudflareSnippetMigrationConfigComplexContent string

//go:embed testdata/v4_url_rewrite.tf
var testAccCloudflareSnippetMigrationConfigURLRewrite string

//go:embed testdata/v4_header_manipulation.tf
var testAccCloudflareSnippetMigrationConfigHeaderManipulation string

//go:embed testdata/v4_edge_cases.tf
var testAccCloudflareSnippetMigrationConfigEdgeCases string

// Configs for _Migration_ tests (zoneID=%[1]s, rnd=%[2]s arg order)

//go:embed testdata/v4_migration_basic.tf
var v4MigrationBasicConfig string

//go:embed testdata/v4_migration_multiple_files.tf
var v4MigrationMultipleFilesConfig string

//go:embed testdata/v4_migration_complex_content.tf
var v4MigrationComplexContentConfig string

// ============================================================================
// Tests ported from migrations_test.go (original PreConfig/RunMigrationV2Command style)
// ============================================================================

func TestMigrateCloudflareSnippetBasic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_snippet.%s", rnd)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: fmt.Sprintf(testAccCloudflareSnippetMigrationConfigBasic, rnd, zoneID),
			},
			{
				PreConfig: func() {
					acctest.WriteOutConfig(t, fmt.Sprintf(testAccCloudflareSnippetMigrationConfigBasic, rnd, zoneID), tmpDir)
					acctest.RunMigrationV2Command(t, fmt.Sprintf(testAccCloudflareSnippetMigrationConfigBasic, rnd, zoneID), tmpDir, "v4", "v5")
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						acctest.DebugNonEmptyPlan,
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("snippet_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("metadata").AtMapKey("main_module"), knownvalue.StringExact("main.js")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files").AtSliceIndex(0), knownvalue.ObjectExact(map[string]knownvalue.Check{
						"name":    knownvalue.StringExact("main.js"),
						"content": knownvalue.StringExact("export default {async fetch(request) {return fetch(request)}};"),
					})),
				},
			},
		},
	})
}

func TestMigrateCloudflareSnippetMultipleFiles(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_snippet.%s", rnd)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: fmt.Sprintf(testAccCloudflareSnippetMigrationConfigMultipleFiles, rnd, zoneID),
			},
			{
				PreConfig: func() {
					acctest.WriteOutConfig(t, fmt.Sprintf(testAccCloudflareSnippetMigrationConfigMultipleFiles, rnd, zoneID), tmpDir)
					acctest.RunMigrationV2Command(t, fmt.Sprintf(testAccCloudflareSnippetMigrationConfigMultipleFiles, rnd, zoneID), tmpDir, "v4", "v5")
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						acctest.DebugNonEmptyPlan,
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("snippet_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("metadata").AtMapKey("main_module"), knownvalue.StringExact("main.js")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files"), knownvalue.ListSizeExact(3)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files").AtSliceIndex(0), knownvalue.ObjectPartial(map[string]knownvalue.Check{
						"name": knownvalue.StringExact("main.js"),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files").AtSliceIndex(1), knownvalue.ObjectPartial(map[string]knownvalue.Check{
						"name": knownvalue.StringExact("helper.js"),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files").AtSliceIndex(2), knownvalue.ObjectPartial(map[string]knownvalue.Check{
						"name":    knownvalue.StringExact("utils.js"),
						"content": knownvalue.StringExact("export const VERSION = '1.0.0';"),
					})),
				},
			},
		},
	})
}

func TestMigrateCloudflareSnippetComplexContent(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_snippet.%s", rnd)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: fmt.Sprintf(testAccCloudflareSnippetMigrationConfigComplexContent, rnd, zoneID),
			},
			{
				PreConfig: func() {
					acctest.WriteOutConfig(t, fmt.Sprintf(testAccCloudflareSnippetMigrationConfigComplexContent, rnd, zoneID), tmpDir)
					acctest.RunMigrationV2Command(t, fmt.Sprintf(testAccCloudflareSnippetMigrationConfigComplexContent, rnd, zoneID), tmpDir, "v4", "v5")
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						acctest.DebugNonEmptyPlan,
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("snippet_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("metadata").AtMapKey("main_module"), knownvalue.StringExact("worker.js")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files").AtSliceIndex(0), knownvalue.ObjectPartial(map[string]knownvalue.Check{
						"name":    knownvalue.StringExact("worker.js"),
						"content": knownvalue.StringRegexp(regexp.MustCompile("Complex worker with multiple features")),
					})),
				},
			},
		},
	})
}

func TestMigrateCloudflareSnippetWithImport(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_snippet.%s", rnd)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: fmt.Sprintf(testAccCloudflareSnippetMigrationConfigBasic, rnd, zoneID),
			},
			{
				PreConfig: func() {
					acctest.WriteOutConfig(t, fmt.Sprintf(testAccCloudflareSnippetMigrationConfigBasic, rnd, zoneID), tmpDir)
					acctest.RunMigrationV2Command(t, fmt.Sprintf(testAccCloudflareSnippetMigrationConfigBasic, rnd, zoneID), tmpDir, "v4", "v5")
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						acctest.DebugNonEmptyPlan,
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("snippet_name"), knownvalue.StringExact(rnd)),
				},
			},
		},
	})
}

func TestMigrateCloudflareSnippetURLRewrite(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_snippet.%s", rnd)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: fmt.Sprintf(testAccCloudflareSnippetMigrationConfigURLRewrite, rnd, zoneID),
			},
			{
				PreConfig: func() {
					acctest.WriteOutConfig(t, fmt.Sprintf(testAccCloudflareSnippetMigrationConfigURLRewrite, rnd, zoneID), tmpDir)
					acctest.RunMigrationV2Command(t, fmt.Sprintf(testAccCloudflareSnippetMigrationConfigURLRewrite, rnd, zoneID), tmpDir, "v4", "v5")
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						acctest.DebugNonEmptyPlan,
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("snippet_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("metadata").AtMapKey("main_module"), knownvalue.StringExact("rewrite.js")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files").AtSliceIndex(0), knownvalue.ObjectPartial(map[string]knownvalue.Check{
						"name":    knownvalue.StringExact("rewrite.js"),
						"content": knownvalue.StringRegexp(regexp.MustCompile(`pathname\.match\(/\^\\/old`)),
					})),
				},
			},
		},
	})
}

func TestMigrateCloudflareSnippetHeaderManipulation(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_snippet.%s", rnd)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: fmt.Sprintf(testAccCloudflareSnippetMigrationConfigHeaderManipulation, rnd, zoneID),
			},
			{
				PreConfig: func() {
					acctest.WriteOutConfig(t, fmt.Sprintf(testAccCloudflareSnippetMigrationConfigHeaderManipulation, rnd, zoneID), tmpDir)
					acctest.RunMigrationV2Command(t, fmt.Sprintf(testAccCloudflareSnippetMigrationConfigHeaderManipulation, rnd, zoneID), tmpDir, "v4", "v5")
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						acctest.DebugNonEmptyPlan,
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("snippet_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("metadata").AtMapKey("main_module"), knownvalue.StringExact("headers.js")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files").AtSliceIndex(0), knownvalue.ObjectPartial(map[string]knownvalue.Check{
						"name":    knownvalue.StringExact("headers.js"),
						"content": knownvalue.StringRegexp(regexp.MustCompile(`headers\.set\("X-Custom-Header"`)),
					})),
				},
			},
		},
	})
}

func TestMigrateCloudflareSnippetEdgeCases(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_snippet.%s", rnd)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: fmt.Sprintf(testAccCloudflareSnippetMigrationConfigEdgeCases, rnd, zoneID),
			},
			{
				PreConfig: func() {
					acctest.WriteOutConfig(t, fmt.Sprintf(testAccCloudflareSnippetMigrationConfigEdgeCases, rnd, zoneID), tmpDir)
					acctest.RunMigrationV2Command(t, fmt.Sprintf(testAccCloudflareSnippetMigrationConfigEdgeCases, rnd, zoneID), tmpDir, "v4", "v5")
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						acctest.DebugNonEmptyPlan,
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("snippet_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("metadata").AtMapKey("main_module"), knownvalue.StringExact("edge.js")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files").AtSliceIndex(0), knownvalue.ObjectPartial(map[string]knownvalue.Check{
						"name":    knownvalue.StringExact("edge.js"),
						"content": knownvalue.StringRegexp(regexp.MustCompile("emoji 🚀")),
					})),
				},
			},
		},
	})
}

// ============================================================================
// Tests ported from migrations_v4_to_v5_test.go (MigrationV2TestStepWithStateNormalization style)
// ============================================================================

func TestMigrateCloudflareSnippet_Migration_Basic_MultiVersion(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(zoneID, rnd string) string
	}{
		{
			name:     "from_v4_52_1",
			version:  "4.52.1",
			configFn: func(zoneID, rnd string) string { return fmt.Sprintf(v4MigrationBasicConfig, zoneID, rnd) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_snippet." + rnd
			testConfig := tc.configFn(zoneID, rnd)
			tmpDir := t.TempDir()

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: append(
					[]resource.TestStep{
						{
							ExternalProviders: map[string]resource.ExternalProvider{
								"cloudflare": {
									VersionConstraint: tc.version,
									Source:            "cloudflare/cloudflare",
								},
							},
							Config: testConfig,
							ConfigStateChecks: []statecheck.StateCheck{
								statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
								statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
								statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("main_module"), knownvalue.StringExact("main.js")),
							},
						},
					},
					acctest.MigrationV2TestStepWithStateNormalization(t, testConfig, tmpDir, tc.version, "v4", "v5", []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("snippet_name"), knownvalue.StringExact(rnd)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("metadata").AtMapKey("main_module"), knownvalue.StringExact("main.js")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files"), knownvalue.ListSizeExact(1)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("main.js")),
					})...,
				),
			})
		})
	}
}

func TestMigrateCloudflareSnippet_Migration_WithMultipleFiles(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_snippet." + rnd
	v4Config := fmt.Sprintf(v4MigrationMultipleFilesConfig, zoneID, rnd)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: append(
			[]resource.TestStep{
				{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							VersionConstraint: "4.52.1",
							Source:            "cloudflare/cloudflare",
						},
					},
					Config: v4Config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					},
				},
			},
			acctest.MigrationV2TestStepWithStateNormalization(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("snippet_name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("metadata").AtMapKey("main_module"), knownvalue.StringExact("main.js")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files"), knownvalue.ListSizeExact(3)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("main.js")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files").AtSliceIndex(1).AtMapKey("name"), knownvalue.StringExact("helper.js")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files").AtSliceIndex(2).AtMapKey("name"), knownvalue.StringExact("utils.js")),
			})...,
		),
	})
}

func TestMigrateCloudflareSnippet_Migration_ComplexContent(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_snippet." + rnd
	v4Config := fmt.Sprintf(v4MigrationComplexContentConfig, zoneID, rnd)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: append(
			[]resource.TestStep{
				{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							VersionConstraint: "4.52.1",
							Source:            "cloudflare/cloudflare",
						},
					},
					Config: v4Config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					},
				},
			},
			acctest.MigrationV2TestStepWithStateNormalization(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("snippet_name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("metadata").AtMapKey("main_module"), knownvalue.StringExact("worker.js")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files"), knownvalue.ListSizeExact(1)),
			})...,
		),
	})
}

