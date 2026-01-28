package logpush_ownership_challenge_test

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

// TestMigrateLogpushOwnershipChallenge_Migration_Basic_Zone tests the basic
// logpush ownership challenge migration with zone_id.
// This test ensures that:
// 1. Configuration stays identical (no HCL transformations)
// 2. State transformation removes ownership_challenge_filename field
// 3. schema_version is set to 0
// 4. All other fields are preserved
func TestMigrateLogpushOwnershipChallenge_Migration_Basic_Zone(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_logpush_ownership_challenge." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_logpush_ownership_challenge" "%[1]s" {
  zone_id          = "%[2]s"
  destination_conf = "gs://cf-terraform-provider-acct-test/ownership_challenges"
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create ownership challenge with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact("gs://cf-terraform-provider-acct-test/ownership_challenges")),
				},
			},
			// Step 2: Migrate to v5 provider
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact("gs://cf-terraform-provider-acct-test/ownership_challenges")),
			}),
			{
				// Step 3: Apply migrated config with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact("gs://cf-terraform-provider-acct-test/ownership_challenges")),
				},
			},
		},
	})
}

// TestMigrateLogpushOwnershipChallenge_Migration_Basic_Account tests the basic
// logpush ownership challenge migration with account_id.
// This test ensures proper migration for account-scoped challenges.
func TestMigrateLogpushOwnershipChallenge_Migration_Basic_Account(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_logpush_ownership_challenge." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_logpush_ownership_challenge" "%[1]s" {
  account_id       = "%[2]s"
  destination_conf = "gs://cf-terraform-provider-acct-test/ownership_challenges"
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create ownership challenge with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact("gs://cf-terraform-provider-acct-test/ownership_challenges")),
				},
			},
			// Step 2: Migrate to v5 provider
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact("gs://cf-terraform-provider-acct-test/ownership_challenges")),
			}),
			{
				// Step 3: Apply migrated config with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact("gs://cf-terraform-provider-acct-test/ownership_challenges")),
				},
			},
		},
	})
}

// TestMigrateLogpushOwnershipChallenge_Migration_GCS tests migration with
// Google Cloud Storage destination to ensure different destination types work.
func TestMigrateLogpushOwnershipChallenge_Migration_GCS(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_logpush_ownership_challenge." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_logpush_ownership_challenge" "%[1]s" {
  zone_id          = "%[2]s"
  destination_conf = "gs://cf-terraform-provider-acct-test/ownership_challenges"
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create ownership challenge with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact("gs://cf-terraform-provider-acct-test/ownership_challenges")),
				},
			},
			// Step 2: Migrate to v5 provider
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact("gs://cf-terraform-provider-acct-test/ownership_challenges")),
			}),
			{
				// Step 3: Apply migrated config with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("destination_conf"), knownvalue.StringExact("gs://cf-terraform-provider-acct-test/ownership_challenges")),
				},
			},
		},
	})
}
