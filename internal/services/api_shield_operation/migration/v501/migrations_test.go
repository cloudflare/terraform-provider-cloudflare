package v501_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

//go:embed testdata/v526_basic.tf
var v526BasicConfig string

// TestMigrateAPIShieldOperation_V526ToCurrentSchema tests that state created by
// provider v5.26.0 — which incorrectly stored the GET query-parameter attributes
// feature (null) and with_schemas (false) — is transparently upgraded to the
// v501 schema by the StateUpgrader, and that the resulting plan is a strict
// no-op with no spurious attribute changes.
func TestMigrateAPIShieldOperation_V526ToCurrentSchema(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_shield_operation." + rnd
	tmpDir := t.TempDir()
	cfg := fmt.Sprintf(v526BasicConfig, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: create with v5.26.0 — the version that shipped with
				// feature and with_schemas in the schema. The state file will
				// contain schema_version=500 with both attributes present.
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.26.0",
					},
				},
				Config: cfg,
			},
			{
				// Step 2: upgrade to the current provider.
				// The 500→501 StateUpgrader strips feature and with_schemas.
				// The plan must be completely empty: neither attribute should
				// appear as a pending change.
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   cfg,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("method"), knownvalue.StringExact("GET")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("host"), knownvalue.StringExact("api.example.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("endpoint"), knownvalue.StringExact("/api/v1/users")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("operation_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
		},
	})
}
