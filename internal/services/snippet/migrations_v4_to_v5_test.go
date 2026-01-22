package snippet_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateCloudflareSnippet_Migration_Basic_MultiVersion tests the snippet
// migration from v4 to v5. This test ensures that:
// 1. name field is renamed to snippet_name
// 2. files blocks are converted to files array attribute
// 3. main_module is moved to metadata.main_module
// 4. The migration tool successfully transforms both configuration and state files
// 5. Resources remain functional after migration without requiring manual intervention
func TestMigrateCloudflareSnippet_Migration_Basic_MultiVersion(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(zoneID, rnd string) string
	}{
		{
			name:     "from_v4_52_1", // Last v4 release
			version:  "4.52.1",
			configFn: testAccCloudflareSnippetMigrationConfigV4Basic,
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
							// Step 1: Create snippet with v4 provider
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
					// Steps 2-4: Migrate to v5 provider with state normalization
					// The state normalization is needed because files is not stored in v4 state
					// and needs to be refreshed from the API after migration
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

// TestMigrateCloudflareSnippet_Migration_WithMultipleFiles tests migration with multiple files
// to ensure all files blocks are properly converted to the array attribute format
func TestMigrateCloudflareSnippet_Migration_WithMultipleFiles(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_snippet." + rnd
	v4Config := testAccCloudflareSnippetMigrationConfigV4MultipleFiles(zoneID, rnd)
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
					// Step 1: Create snippet with v4 provider
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
			// Steps 2-4: Migrate to v5 provider with state normalization
			acctest.MigrationV2TestStepWithStateNormalization(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("snippet_name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("metadata").AtMapKey("main_module"), knownvalue.StringExact("main.js")),
				// Verify all 3 files are migrated
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files"), knownvalue.ListSizeExact(3)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("main.js")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files").AtSliceIndex(1).AtMapKey("name"), knownvalue.StringExact("helper.js")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files").AtSliceIndex(2).AtMapKey("name"), knownvalue.StringExact("utils.js")),
			})...,
		),
	})
}

// TestMigrateCloudflareSnippet_Migration_ComplexContent tests migration with complex JavaScript content
// containing special characters, escape sequences, and multi-line code
func TestMigrateCloudflareSnippet_Migration_ComplexContent(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_snippet." + rnd
	v4Config := testAccCloudflareSnippetMigrationConfigV4ComplexContent(zoneID, rnd)
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
					// Step 1: Create snippet with v4 provider
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
			// Steps 2-4: Migrate to v5 provider with state normalization
			acctest.MigrationV2TestStepWithStateNormalization(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("snippet_name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("metadata").AtMapKey("main_module"), knownvalue.StringExact("worker.js")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("files"), knownvalue.ListSizeExact(1)),
			})...,
		),
	})
}

// V4 Configuration Functions

func testAccCloudflareSnippetMigrationConfigV4Basic(zoneID, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_snippet" "%[2]s" {
  zone_id     = "%[1]s"
  name        = "%[2]s"
  main_module = "main.js"

  files {
    name    = "main.js"
    content = "export default {async fetch(request) {return fetch(request)}};"
  }
}
`, zoneID, rnd)
}

func testAccCloudflareSnippetMigrationConfigV4MultipleFiles(zoneID, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_snippet" "%[2]s" {
  zone_id     = "%[1]s"
  name        = "%[2]s"
  main_module = "main.js"

  files {
    name    = "main.js"
    content = "import { helper } from './helper.js'; export default {async fetch(request) {return helper(request)}};"
  }

  files {
    name    = "helper.js"
    content = "export function helper(request) { return fetch(request); }"
  }

  files {
    name    = "utils.js"
    content = "export const VERSION = '1.0.0';"
  }
}
`, zoneID, rnd)
}

func testAccCloudflareSnippetMigrationConfigV4ComplexContent(zoneID, rnd string) string {
	return fmt.Sprintf(`
resource "cloudflare_snippet" "%[2]s" {
  zone_id     = "%[1]s"
  name        = "%[2]s"
  main_module = "worker.js"

  files {
    name    = "worker.js"
    content = <<-EOT
      // Complex worker with multiple features
      export default {
        async fetch(request, env, ctx) {
          const url = new URL(request.url);

          // Handle different paths
          if (url.pathname === '/api/data') {
            return new Response(JSON.stringify({
              message: "Hello from Cloudflare",
              path: url.pathname,
              timestamp: Date.now()
            }), {
              headers: {
                'Content-Type': 'application/json',
                'X-Custom-Header': 'test-value'
              }
            });
          }

          // Default response
          return fetch(request);
        }
      };
    EOT
  }
}
`, zoneID, rnd)
}
