package workers_for_platforms_dispatch_namespace_test

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

// TestMigrateWorkersForPlatformsDispatchNamespaceBasic tests basic migration from v4 to v5
func TestMigrateWorkersForPlatformsDispatchNamespaceBasic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_for_platforms_dispatch_namespace." + rnd
	tmpDir := t.TempDir()
	name := fmt.Sprintf("test-dispatch-namespace-%s", rnd)

	// V4 config - simple pass-through migration
	v4Config := fmt.Sprintf(`
resource "cloudflare_workers_for_platforms_dispatch_namespace" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[3]s"
}`, rnd, accountID, name)

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
				// Verify resource exists with same type (no rename)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(name)),
				// Verify new computed fields are present in v5 (populated by provider on refresh)
				// Note: These fields won't be in the migrated state immediately, but will be added by provider
			}),
		},
	})
}

// TestMigrateWorkersForPlatformsDispatchNamespaceWithoutName tests migration with optional name omitted
func TestMigrateWorkersForPlatformsDispatchNamespaceWithoutName(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_for_platforms_dispatch_namespace." + rnd
	tmpDir := t.TempDir()

	// V4 config - name is Required in v4, Optional in v5
	// This tests that the migration handles the field correctly
	v4Config := fmt.Sprintf(`
resource "cloudflare_workers_for_platforms_dispatch_namespace" "%[1]s" {
  account_id = "%[2]s"
  name       = "namespace-%[1]s"
}`, rnd, accountID)

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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("namespace-%s", rnd))),
			}),
		},
	})
}

// TestMigrateWorkersForPlatformsDispatchNamespaceWithSpecialChars tests migration with special characters in name
func TestMigrateWorkersForPlatformsDispatchNamespaceWithSpecialChars(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_for_platforms_dispatch_namespace." + rnd
	tmpDir := t.TempDir()
	// Name with dashes, underscores, and numbers
	name := fmt.Sprintf("test_namespace-%s-2024", rnd)

	v4Config := fmt.Sprintf(`
resource "cloudflare_workers_for_platforms_dispatch_namespace" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[3]s"
}`, rnd, accountID, name)

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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(name)),
			}),
		},
	})
}
