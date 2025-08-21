package snippet_rules_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareSnippetRules_Migration_BasicBlockToAttribute(t *testing.T) {
	// This test verifies that the migration tool properly converts the rules block
	// structure from v4 to the rules attribute structure in v5.
	// v4: rules { } (list block)
	// v5: rules = [ ] (list attribute)

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_snippet_rules." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()

	// V4 config using block syntax - we'll reference an existing snippet name
	// that should already exist in the test zone
	v4Config := fmt.Sprintf(`
resource "cloudflare_snippet_rules" "%[1]s" {
  zone_id = "%[2]s"
  
  rules {
    expression   = "true"
    snippet_name = "rules_set_snippet"
    enabled      = true
    description  = "Test snippet rule"
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
				// Step 1: Create snippet rules with v4 provider using block syntax
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("snippet_name"), knownvalue.StringExact("rules_set_snippet")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expression"), knownvalue.StringExact("true")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("description"), knownvalue.StringExact("Test snippet rule")),
				},
			},
			// Step 2: Migrate to v5 provider - block should convert to attribute
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("snippet_name"), knownvalue.StringExact("rules_set_snippet")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expression"), knownvalue.StringExact("true")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("description"), knownvalue.StringExact("Test snippet rule")),
			}),
		},
	})
}

func TestAccCloudflareSnippetRules_Migration_MultipleRules(t *testing.T) {
	// This test verifies migration with multiple rules including different optional attributes
	
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_snippet_rules." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	tmpDir := t.TempDir()

	// V4 config with multiple rules using block syntax
	v4Config := fmt.Sprintf(`
resource "cloudflare_snippet_rules" "%[1]s" {
  zone_id = "%[2]s"
  
  rules {
    expression   = "http.request.uri.path contains \"/api\""
    snippet_name = "rules_set_snippet"
    enabled      = false
  }

  rules {
    expression   = "http.request.uri.path contains \"/admin\""
    snippet_name = "rules_set_snippet"
    enabled      = false
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
				// Step 1: Create snippet rules with v4 provider using block syntax
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					// First rule - with explicit enabled value
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("snippet_name"), knownvalue.StringExact("rules_set_snippet")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expression"), knownvalue.StringExact("http.request.uri.path contains \"/api\"")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(false)),
					// Second rule - all attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("snippet_name"), knownvalue.StringExact("rules_set_snippet")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("expression"), knownvalue.StringExact("http.request.uri.path contains \"/admin\"")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("description"), knownvalue.StringExact("Admin rule")),
				},
			},
			// Step 2: Migrate to v5 provider
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				// Both rules should preserve their explicit values from v4
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("snippet_name"), knownvalue.StringExact("rules_set_snippet")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expression"), knownvalue.StringExact("http.request.uri.path contains \"/api\"")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("description"), knownvalue.StringExact("")),  // v5 default for unset description
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("snippet_name"), knownvalue.StringExact("rules_set_snippet")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("expression"), knownvalue.StringExact("http.request.uri.path contains \"/admin\"")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("enabled"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("description"), knownvalue.StringExact("Admin rule")),
			}),
		},
	})
}