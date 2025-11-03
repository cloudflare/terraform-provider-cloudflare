package snippet_rules_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateSnippetRules_Basic tests basic migration from v4 to v5
// This test verifies that:
// - rules blocks are converted to rules list attribute
func TestMigrateSnippetRules_Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_snippet_rules." + rnd
	tmpDir := t.TempDir()

	// V4 config with rules blocks
	v4Config := fmt.Sprintf(`
resource "cloudflare_snippet_rules" "%[1]s" {
  zone_id = "%[2]s"

  rules {
    enabled      = true
    expression   = "true"
    snippet_name = "snippet_1"
    description  = "Test rule"
  }
}`, rnd, zoneID)

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
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				// Verify rules blocks were converted to rules list
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expression"), knownvalue.StringExact("true")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("snippet_name"), knownvalue.StringExact("snippet_1")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("description"), knownvalue.StringExact("Test rule")),
			}),
		},
	})
}

// TestMigrateSnippetRules_MultipleRules tests migration with multiple rules blocks
func TestMigrateSnippetRules_MultipleRules(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_snippet_rules." + rnd
	tmpDir := t.TempDir()

	// V4 config with multiple rules blocks
	v4Config := fmt.Sprintf(`
resource "cloudflare_snippet_rules" "%[1]s" {
  zone_id = "%[2]s"

  rules {
    enabled      = true
    expression   = "http.request.uri.path contains \"/api\""
    snippet_name = "api_snippet"
    description  = "API rule"
  }

  rules {
    enabled      = false
    expression   = "http.request.uri.path contains \"/admin\""
    snippet_name = "admin_snippet"
    description  = "Admin rule"
  }
}`, rnd, zoneID)

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
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				// Verify both rules blocks were converted to rules list
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(2)),
				// Verify first rule
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("snippet_name"), knownvalue.StringExact("api_snippet")),
				// Verify second rule
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("enabled"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("snippet_name"), knownvalue.StringExact("admin_snippet")),
			}),
		},
	})
}

// TestMigrateSnippetRules_OptionalAttributes tests migration with optional attributes
func TestMigrateSnippetRules_OptionalAttributes(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_snippet_rules." + rnd
	tmpDir := t.TempDir()

	// V4 config with minimal attributes (enabled and description are optional)
	v4Config := fmt.Sprintf(`
resource "cloudflare_snippet_rules" "%[1]s" {
  zone_id = "%[2]s"

  rules {
    expression   = "true"
    snippet_name = "minimal_snippet"
  }
}`, rnd, zoneID)

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
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expression"), knownvalue.StringExact("true")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("snippet_name"), knownvalue.StringExact("minimal_snippet")),
				// Verify default values are set
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("description"), knownvalue.StringExact("")),
			}),
		},
	})
}
