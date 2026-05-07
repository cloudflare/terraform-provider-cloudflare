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

// Migration Test Configuration — v1 → v501 (Set→List)
//
// v5.16.0 and v5.18.0 both stored account_token state at schema_version=1
// with Set-typed policies and permission_groups.
//
// The version 1 upgrader must decode Set state and sort canonically for stable
// List ordering. Before the fix for GitHub issue #7077, slot 1 routed through
// v500.UpgradeFromV1 which performed a no-op (deserialize/re-serialize without
// sorting), causing permission_groups to appear in arbitrary order and triggering
// "inconsistent state" errors during apply.

// TestMigrateAccountTokenFromV1_Basic tests the simplest case: a single-policy
// account_token created in v5.18 upgrading to the current version.
// Before the fix, the version 1 upgrader performed a no-op without sorting,
// causing List-based state with arbitrary Set ordering.
func TestMigrateAccountTokenFromV1_Basic(t *testing.T) {
	testCases := []struct {
		name    string
		version string
	}{
		{
			name:    "from_v5.16",
			version: "5.16.0",
		},
		{
			name:    "from_v5.18",
			version: "5.18.0",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			tmpDir := t.TempDir()

			config := fmt.Sprintf(`
resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"

  policies = [
    {
      effect = "allow"
      permission_groups = [{
        id = "82e64a83756745bbbb1c9c2701bf816b"
      }]
      resources = jsonencode({
        "com.cloudflare.api.account.%[2]s" = "*"
      })
    }
  ]
}`, rnd, accountID)

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
								VersionConstraint: tc.version,
							},
						},
						Config: config,
					},
					{
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
								knownvalue.ListSizeExact(1),
							),
							statecheck.ExpectKnownValue(
								fmt.Sprintf("cloudflare_account_token.%s", rnd),
								tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("effect"),
								knownvalue.StringExact("allow"),
							),
						},
					},
				},
			})
		})
	}
}

// TestMigrateAccountTokenFromV1_MultiPolicy tests migration from schema_version=1
// (Set-based) to v501 (List-based) with multiple policies.
func TestMigrateAccountTokenFromV1_MultiPolicy(t *testing.T) {
	testCases := []struct {
		name    string
		version string
	}{
		{
			name:    "from_v5.16",
			version: "5.16.0",
		},
		{
			name:    "from_v5.18",
			version: "5.18.0",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
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
						// Create with v5.16/v5.18 (stores schema_version=1, Set-based)
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: config,
					},
					{
						// Upgrade to current (v501, List-based)
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
		})
	}
}

// TestMigrateAccountTokenFromV1_MultiPermGroups tests migration from schema_version=1
// (Set-based) to v501 (List-based) with a policy that has multiple permission groups.
//
// This is the exact scenario from GitHub issue #7077 where the "inconsistent state"
// error occurred: permission_groups were returned by the API in a different order
// than stored in state, and without canonical sorting during the v1→v501 upgrade,
// the List-based state had arbitrary ordering.
func TestMigrateAccountTokenFromV1_MultiPermGroups(t *testing.T) {
	testCases := []struct {
		name    string
		version string
	}{
		{
			name:    "from_v5.16",
			version: "5.16.0",
		},
		{
			name:    "from_v5.18",
			version: "5.18.0",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
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
						// Create with v5.16/v5.18 (stores schema_version=1, Set-based)
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: config,
					},
					{
						// Upgrade to current (v501, List-based) — before the fix,
						// this could produce unsorted permission_groups causing
						// "inconsistent state" on the next apply
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
		})
	}
}
