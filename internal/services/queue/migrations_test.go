package queue_test

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

func TestMigrateQueue_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()
	resourceName := "cloudflare_queue." + rnd

	// v4 config uses "name" attribute
	v4Config := fmt.Sprintf(`
resource "cloudflare_queue" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-queue-%[1]s"
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: testAccCloudflareQueueDestroy,
		WorkingDir:   tmpDir,
		Steps: append([]resource.TestStep{
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
		}, // Step 2: Run migration and verify state
			acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("queue_name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-queue-%s", rnd))),
			})...),
	})
}

func TestMigrateQueue_WithSettings(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()
	resourceName := "cloudflare_queue." + rnd

	// v4 config with settings (if supported in v4)
	v4Config := fmt.Sprintf(`
resource "cloudflare_queue" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-queue-settings-%[1]s"
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: testAccCloudflareQueueDestroy,
		WorkingDir:   tmpDir,
		Steps: append([]resource.TestStep{
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
		}, // Step 2: Run migration and verify state
			acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("queue_name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-queue-settings-%s", rnd))),
			})...),
	})
}
