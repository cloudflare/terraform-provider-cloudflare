package v501_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

//go:embed testdata/v501_multi_policy.tf
var v501MultiPolicyConfig string

//go:embed testdata/v501_multi_permgroups.tf
var v501MultiPermgroupsConfig string

// TestMigrateAccountTokenFromV500_MultiPolicy tests migration from v500 (Set-based,
// schema_version=500 via v5.16.0) to v501 (List-based) with multiple policies.
// This verifies that the Set→List migration properly converts policy collections
// and that the provider can plan/apply without drift.
func TestMigrateAccountTokenFromV500_MultiPolicy(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	config := fmt.Sprintf(v501MultiPolicyConfig, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Create with v5.16.0 which has schema_version=500 (Set-based)
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.16.0",
					},
				},
				Config: config,
			},
			{
				// Migrate to current (v501, List-based) and verify no drift
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("name"),
						knownvalue.StringExact(rnd),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies"),
						knownvalue.ListSizeExact(2),
					),
				},
			},
		},
	})
}

// TestMigrateAccountTokenFromV500_MultiPermGroups tests migration from v500
// (Set-based) to v501 (List-based) with a policy that has multiple permission
// groups. This is the scenario that triggers the O(n²) performance issue —
// verifying that migration works correctly and produces no drift.
func TestMigrateAccountTokenFromV500_MultiPermGroups(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	tmpDir := t.TempDir()

	config := fmt.Sprintf(v501MultiPermgroupsConfig, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Create with v5.16.0 which has schema_version=500 (Set-based)
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.16.0",
					},
				},
				Config: config,
			},
			{
				// Migrate to current (v501, List-based) and verify no drift
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("name"),
						knownvalue.StringExact(rnd),
					),
					statecheck.ExpectKnownValue(
						fmt.Sprintf("cloudflare_account_token.%s", rnd),
						tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups"),
						knownvalue.ListSizeExact(3),
					),
				},
			},
		},
	})
}
