package v500_test

import (
	_ "embed"
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

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v4_with_settings.tf
var v4WithSettingsConfig string

func TestMigrateQueue_Basic(t *testing.T) {
	t.Skip("Migration not enabled yet")
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()
	resourceName := "cloudflare_queue." + rnd

	// v4 config uses "name" attribute
	v4Config := fmt.Sprintf(v4BasicConfig, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
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
		},
			// Step 2: Run tf-migrate and apply with v5 provider, verify state
			acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("queue_name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-queue-%s", rnd))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("queue_id"), knownvalue.NotNull()),
			})...),
	})
}

func TestMigrateQueue_WithSettings(t *testing.T) {
	t.Skip("Migration not enabled yet")
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()
	resourceName := "cloudflare_queue." + rnd

	// v4 config - settings are not supported in v4, so this tests the base migration path
	// with a queue name that includes "settings" in it for distinguishability
	v4Config := fmt.Sprintf(v4WithSettingsConfig, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
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
		},
			// Step 2: Run tf-migrate and apply with v5 provider, verify state
			acctest.MigrationV2TestStepWithPlan(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("queue_name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-queue-settings-%s", rnd))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("queue_id"), knownvalue.NotNull()),
			})...),
	})
}
