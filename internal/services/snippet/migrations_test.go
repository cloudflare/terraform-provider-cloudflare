package snippet_test

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

//go:embed testdata/migration_basic.tf
var testAccCloudflareSnippetMigrationConfigBasic string

//go:embed testdata/migration_multiple_files.tf
var testAccCloudflareSnippetMigrationConfigMultipleFiles string

//go:embed testdata/migration_complex_content.tf
var testAccCloudflareSnippetMigrationConfigComplexContent string

//go:embed testdata/migration_url_rewrite.tf
var testAccCloudflareSnippetMigrationConfigURLRewrite string

//go:embed testdata/migration_header_manipulation.tf
var testAccCloudflareSnippetMigrationConfigHeaderManipulation string

//go:embed testdata/migration_edge_cases.tf
var testAccCloudflareSnippetMigrationConfigEdgeCases string

// TestMigrateCloudflareSnippetBasic tests migration of basic snippet with only required attributes
// Tests the core attribute renames: name → snippet_name
// Tests structural change: main_module → metadata.main_module
// Tests type conversion: files block → list attribute
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
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: fmt.Sprintf(testAccCloudflareSnippetMigrationConfigBasic, rnd, zoneID),
			},
			// Step 2: Run migration and verify state
			{
				PreConfig: func() {
					// Write out config
					acctest.WriteOutConfig(t, fmt.Sprintf(testAccCloudflareSnippetMigrationConfigBasic, rnd, zoneID), tmpDir)

					// Run V2 migration
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
					// Verify renamed attribute
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("snippet_name"), knownvalue.StringExact(rnd)),
					// Verify restructured metadata
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("metadata").AtMapKey("main_module"), knownvalue.StringExact("main.js")),
					// Verify files converted to list attribute
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

// TestMigrateCloudflareSnippetMultipleFiles tests migration with multiple files
// Ensures all files are properly migrated from blocks to list attributes
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
				// Step 1: Create with v4 provider
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
					// Write out config
					acctest.WriteOutConfig(t, fmt.Sprintf(testAccCloudflareSnippetMigrationConfigMultipleFiles, rnd, zoneID), tmpDir)

					// Run V2 migration
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
					// Verify all files are migrated
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files"), knownvalue.ListSizeExact(3)),
					// Check first file
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files").AtSliceIndex(0), knownvalue.ObjectPartial(map[string]knownvalue.Check{
						"name": knownvalue.StringExact("main.js"),
					})),
					// Check second file
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files").AtSliceIndex(1), knownvalue.ObjectPartial(map[string]knownvalue.Check{
						"name": knownvalue.StringExact("helper.js"),
					})),
					// Check third file
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files").AtSliceIndex(2), knownvalue.ObjectPartial(map[string]knownvalue.Check{
						"name":    knownvalue.StringExact("utils.js"),
						"content": knownvalue.StringExact("export const VERSION = '1.0.0';"),
					})),
				},
			},
		},
	})
}

// TestMigrateCloudflareSnippetComplexContent tests migration with complex JavaScript content
// Ensures content integrity during migration including special characters, multi-line strings, comments
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
				// Step 1: Create with v4 provider
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
					// Write out config
					acctest.WriteOutConfig(t, fmt.Sprintf(testAccCloudflareSnippetMigrationConfigComplexContent, rnd, zoneID), tmpDir)

					// Run V2 migration
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
					// Verify the complex content is preserved
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files").AtSliceIndex(0), knownvalue.ObjectPartial(map[string]knownvalue.Check{
						"name": knownvalue.StringExact("worker.js"),
						// Check that content contains expected patterns
						"content": knownvalue.StringRegexp(regexp.MustCompile("Complex worker with multiple features")),
					})),
				},
			},
		},
	})
}

// TestMigrateCloudflareSnippetWithImport tests migration followed by import
// Ensures import functionality works correctly after migration
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
				// Step 1: Create with v4 provider
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
					// Write out config
					acctest.WriteOutConfig(t, fmt.Sprintf(testAccCloudflareSnippetMigrationConfigBasic, rnd, zoneID), tmpDir)

					// Run V2 migration
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

// TestMigrateCloudflareSnippetURLRewrite tests migration of snippet with URL rewrite logic
// Tests realistic use case of URL path manipulation
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
				// Step 1: Create with v4 provider
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
					// Write out config
					acctest.WriteOutConfig(t, fmt.Sprintf(testAccCloudflareSnippetMigrationConfigURLRewrite, rnd, zoneID), tmpDir)

					// Run V2 migration
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
						"name": knownvalue.StringExact("rewrite.js"),
						// Verify URL rewrite logic is preserved
						"content": knownvalue.StringRegexp(regexp.MustCompile(`pathname\.match\(/\^\\/old`)),
					})),
				},
			},
		},
	})
}

// TestMigrateCloudflareSnippetHeaderManipulation tests migration of snippet with header manipulation
// Tests preservation of security-related header modifications
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
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: fmt.Sprintf(testAccCloudflareSnippetMigrationConfigHeaderManipulation, rnd, zoneID),
			},
			// Step 2: Run migration and verify state
			{
				PreConfig: func() {
					// Write out config
					acctest.WriteOutConfig(t, fmt.Sprintf(testAccCloudflareSnippetMigrationConfigHeaderManipulation, rnd, zoneID), tmpDir)

					// Run V2 migration
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
						"name": knownvalue.StringExact("headers.js"),
						// Check header operations are preserved
						"content": knownvalue.StringRegexp(regexp.MustCompile("headers\\.set\\(\"X-Custom-Header\"")),
					})),
				},
			},
		},
	})
}

// TestMigrateCloudflareSnippetEdgeCases tests migration with special characters and edge cases
// Tests handling of quotes, escapes, unicode, and other special content
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
				// Step 1: Create with v4 provider
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
					// Write out config
					acctest.WriteOutConfig(t, fmt.Sprintf(testAccCloudflareSnippetMigrationConfigEdgeCases, rnd, zoneID), tmpDir)

					// Run V2 migration
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
						"name": knownvalue.StringExact("edge.js"),
						// Check special characters are preserved
						"content": knownvalue.StringRegexp(regexp.MustCompile("emoji 🚀")),
					})),
				},
			},
		},
	})
}
