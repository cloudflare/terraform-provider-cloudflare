package workers_cron_trigger_test

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

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// TestMigrateWorkersCronTrigger_Basic tests basic migration from v4 to v5
// This test verifies that:
// - cloudflare_worker_cron_trigger (singular) is renamed to cloudflare_workers_cron_trigger (plural)
// - all attributes are preserved
func TestMigrateWorkersCronTrigger_Basic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	scriptName := "test-script-" + rnd
	resourceName := "cloudflare_workers_cron_trigger." + rnd
	tmpDir := t.TempDir()

	// V4 config using cloudflare_worker_cron_trigger (singular)
	v4Config := fmt.Sprintf(`
resource "cloudflare_worker_cron_trigger" "%[1]s" {
  account_id  = "%[2]s"
  script_name = "%[3]s"
  schedules   = [
    {
      cron = "*/5 * * * *"
    }
  ]
}`, rnd, accountID, scriptName)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(scriptName)),
				// Verify schedules attribute is preserved
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("schedules"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("schedules").AtSliceIndex(0).AtMapKey("cron"), knownvalue.StringExact("*/5 * * * *")),
			}),
		},
	})
}

// TestMigrateWorkersCronTrigger_MultipleSchedules tests migration with multiple schedules
func TestMigrateWorkersCronTrigger_MultipleSchedules(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	scriptName := "test-script-" + rnd
	resourceName := "cloudflare_workers_cron_trigger." + rnd
	tmpDir := t.TempDir()

	// V4 config with multiple schedules
	v4Config := fmt.Sprintf(`
resource "cloudflare_worker_cron_trigger" "%[1]s" {
  account_id  = "%[2]s"
  script_name = "%[3]s"
  schedules   = [
    {
      cron = "*/5 * * * *"
    },
    {
      cron = "10 7 * * mon-fri"
    },
    {
      cron = "0 0 * * *"
    }
  ]
}`, rnd, accountID, scriptName)

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
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(scriptName)),
				// Verify all schedules are preserved
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("schedules"), knownvalue.ListSizeExact(3)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("schedules").AtSliceIndex(0).AtMapKey("cron"), knownvalue.StringExact("*/5 * * * *")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("schedules").AtSliceIndex(1).AtMapKey("cron"), knownvalue.StringExact("10 7 * * mon-fri")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("schedules").AtSliceIndex(2).AtMapKey("cron"), knownvalue.StringExact("0 0 * * *")),
			}),
		},
	})
}
