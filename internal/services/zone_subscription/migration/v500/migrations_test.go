package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Embed test configs
//
//go:embed testdata/v5_basic.tf
var v5BasicConfig string

// TestMigrateZoneSubscription_V5 tests zone_subscription state version bump within v5.
//
// cloudflare_zone_subscription is a new v5 resource with no v4 counterpart.
// This test verifies that existing v5 state at version 0 is correctly bumped
// to version 500 without any transformation.
func TestMigrateZoneSubscription_V5(t *testing.T) {
	if os.Getenv("migration mode") == "" {
		t.Skip("Skipping migration test: migration mode is not set")
	}

	zoneID := "5a870ecfe7d96ad6c056fdaf44a72556" //os.Getenv("CLOUDFLARE_ALT_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	testConfig := fmt.Sprintf(v5BasicConfig, rnd, zoneID)
	sourceVer, targetVer := acctest.InferMigrationVersions(currentProviderVersion)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with current dev provider (state at version 0/1)
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testConfig,
			},
			// Step 2: Run migration and verify state is preserved (no-op version bump)
			acctest.MigrationV2TestStep(t, testConfig, tmpDir, currentProviderVersion, sourceVer, targetVer,
				[]statecheck.StateCheck{
					// Verify zone_id is preserved after version bump
					statecheck.ExpectKnownValue(
						"cloudflare_zone_subscription."+rnd,
						tfjsonpath.New("zone_id"),
						knownvalue.StringExact(zoneID),
					),
					// Verify id is preserved (id == zone_id for zone_subscription)
					statecheck.ExpectKnownValue(
						"cloudflare_zone_subscription."+rnd,
						tfjsonpath.New("id"),
						knownvalue.StringExact(zoneID),
					),
				},
			),
		},
	})
}
