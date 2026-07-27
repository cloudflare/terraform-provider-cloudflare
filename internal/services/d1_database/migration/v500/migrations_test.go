package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

func TestMigrateD1Database_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()
	resourceName := "cloudflare_d1_database." + rnd

	v4Config := fmt.Sprintf(v4BasicConfig, rnd, accountID)
	v5Config := fmt.Sprintf(v5BasicConfig, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: nil,
		WorkingDir:   tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: acctest.GetLastV4Version(),
					},
				},
				Config: v4Config,
			},
			{
				// Step 2: Run tf-migrate, then apply with v5 provider using
				// the v5 config that includes read_replication to match the
				// API default. Without it the provider sends null which the
				// D1 API rejects.
				PreConfig: func() {
					acctest.WriteOutConfig(t, v5Config, tmpDir)
					acctest.RunMigrationV2Command(t, v5Config, tmpDir, "v4", "v5")
					providerTF := filepath.Join(tmpDir, "provider.tf")
					os.Remove(providerTF)
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						acctest.DebugNonEmptyPlan,
						acctest.ExpectEmptyPlanExceptFalseyToNull,
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-d1-%s", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
		},
	})
}
