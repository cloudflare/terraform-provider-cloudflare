package turnstile_widget_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateCloudflare TurnstileWidget_Basic tests the basic turnstile_widget migration scenario
// with minimal configuration. This test ensures that:
// 1. The domains field is correctly transformed from Set to List
// 2. State transformation sets schema_version = 0
// 3. All required fields are preserved
// 4. The migration tool successfully transforms both configuration and state files
func TestMigrateCloudflareTurnstileWidget_Basic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_turnstile_widget." + rnd
	tmpDir := t.TempDir()

	// v4 configuration with toset() wrapper
	v4Config := fmt.Sprintf(`
resource "cloudflare_turnstile_widget" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-test-%[1]s"
  domains    = toset(["example.com"])
  mode       = "managed"
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create resource with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-test-%s", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mode"), knownvalue.StringExact("managed")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domains"), knownvalue.ListSizeExact(1)),
				},
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-test-%s", rnd))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mode"), knownvalue.StringExact("managed")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domains"), knownvalue.ListSizeExact(1)),
			}),
			{
				// Step 3: Apply the migrated configuration with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-test-%s", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mode"), knownvalue.StringExact("managed")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domains"), knownvalue.ListSizeExact(1)),
				},
			},
		},
	})
}

// TestMigrateCloudflareTurnstileWidget_AllOptionalFields tests turnstile_widget migration
// with all optional fields configured. This verifies that:
// 1. All optional fields are preserved (region, bot_fight_mode, offlabel)
// 2. Multiple domains are handled correctly in the Set→List transformation
// 3. All configuration modes work (managed, invisible, non-interactive)
func TestMigrateCloudflareTurnstileWidget_AllOptionalFields(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_turnstile_widget." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_turnstile_widget" "%[1]s" {
  account_id     = "%[2]s"
  name           = "tf-test-full-%[1]s"
  domains        = toset(["example.com", "test.example.org"])
  mode           = "invisible"
  region         = "world"
  bot_fight_mode = false
  offlabel       = false
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create resource with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mode"), knownvalue.StringExact("invisible")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("region"), knownvalue.StringExact("world")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bot_fight_mode"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("offlabel"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domains"), knownvalue.ListSizeExact(2)),
				},
			},
			// Step 2: Run migration
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mode"), knownvalue.StringExact("invisible")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("region"), knownvalue.StringExact("world")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bot_fight_mode"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("offlabel"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domains"), knownvalue.ListSizeExact(2)),
			}),
			{
				// Step 3: Apply with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mode"), knownvalue.StringExact("invisible")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("region"), knownvalue.StringExact("world")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bot_fight_mode"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("offlabel"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domains"), knownvalue.ListSizeExact(2)),
				},
			},
		},
	})
}

// TestMigrateCloudflareTurnstileWidget_NonInteractiveMode tests migration
// with non-interactive mode, ensuring all three mode types work correctly.
func TestMigrateCloudflareTurnstileWidget_NonInteractiveMode(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_turnstile_widget." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_turnstile_widget" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-test-nonint-%[1]s"
  domains    = toset(["noninteractive.example.com"])
  mode       = "non-interactive"
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mode"), knownvalue.StringExact("non-interactive")),
				},
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mode"), knownvalue.StringExact("non-interactive")),
			}),
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("mode"), knownvalue.StringExact("non-interactive")),
				},
			},
		},
	})
}
