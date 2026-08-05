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

//go:embed testdata/v4_with_caching.tf
var v4WithCachingConfig string

//go:embed testdata/v5_with_caching.tf
var v5WithCachingConfig string

// TestMigrateHyperdriveConfig_Basic tests v4 -> v5 migration for a basic
// hyperdrive config without caching. The v5 config must include
// caching = { disabled = false } to match the API default.
func TestMigrateHyperdriveConfig_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()
	resourceName := "cloudflare_hyperdrive_config." + rnd
	dbHost, dbName, dbPort, dbUser, dbPassword := hyperdriveEnv(t)

	v4Config := fmt.Sprintf(v4BasicConfig, rnd, accountID, dbName, dbHost, dbPort, dbUser, dbPassword)
	v5Config := fmt.Sprintf(v5BasicConfig, rnd, accountID, dbName, dbHost, dbPort, dbUser, dbPassword)

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
				// Step 2: Run tf-migrate, then apply with v5 provider
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
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-hyperdrive-%s", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
		},
	})
}

// TestMigrateHyperdriveConfig_WithCaching tests v4 -> v5 migration for a
// hyperdrive config with explicit caching. This exercises the types.Object ->
// struct pointer conversion in the state upgrader.
func TestMigrateHyperdriveConfig_WithCaching(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()
	resourceName := "cloudflare_hyperdrive_config." + rnd
	dbHost, dbName, dbPort, dbUser, dbPassword := hyperdriveEnv(t)

	v4Config := fmt.Sprintf(v4WithCachingConfig, rnd, accountID, dbName, dbHost, dbPort, dbUser, dbPassword)
	v5Config := fmt.Sprintf(v5WithCachingConfig, rnd, accountID, dbName, dbHost, dbPort, dbUser, dbPassword)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		CheckDestroy: nil,
		WorkingDir:   tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: acctest.GetLastV4Version(),
					},
				},
				Config: v4Config,
			},
			{
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
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-hyperdrive-%s", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("caching").AtMapKey("disabled"), knownvalue.Bool(true)),
				},
			},
		},
	})
}

// TestMigrateHyperdriveConfig_V5Idempotent verifies that a resource created
// with the current v5 provider (schema_version=500) can be re-applied without
// changes. Note: this does NOT exercise the UpgradeFromV0 handler because the
// resource is created at version 500 from the start. The v5 no-op path through
// UpgradeFromV0 is tested in handler_test.go (TestUpgradeFromV0_V5RawState).
func TestMigrateHyperdriveConfig_V5Idempotent(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_hyperdrive_config." + rnd
	dbHost, dbName, dbPort, dbUser, dbPassword := hyperdriveEnv(t)

	v5Config := fmt.Sprintf(v5BasicConfig, rnd, accountID, dbName, dbHost, dbPort, dbUser, dbPassword)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: v5Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-hyperdrive-%s", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
			{
				// Re-apply same config -- should produce no changes
				Config: v5Config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						acctest.ExpectEmptyPlanExceptFalseyToNull,
					},
				},
			},
		},
	})
}

// hyperdriveEnv loads and validates Hyperdrive database env vars.
func hyperdriveEnv(t *testing.T) (host, name, port, user, password string) {
	t.Helper()
	if os.Getenv("TF_ACC") == "" {
		t.Skip("acceptance tests skipped unless TF_ACC is set")
	}
	acctest.TestAccPreCheck_Hyperdrive(t)
	return os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_HOSTNAME"),
		os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_NAME"),
		os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_PORT"),
		os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_USER"),
		os.Getenv("CLOUDFLARE_HYPERDRIVE_DATABASE_PASSWORD")
}
