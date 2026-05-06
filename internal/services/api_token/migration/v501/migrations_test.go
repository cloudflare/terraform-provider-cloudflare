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
// v5.16.0 and v5.18.0 both stored api_token state at schema_version=1
// (via `Version: 1` or `GetSchemaVersion(1, 500)` in production).
// The state uses Set-typed policies and permission_groups.
//
// The version 1 upgrader must use the Set-based PriorSchema (not the current
// List-based schema) to decode this state, then sort canonically for List ordering.
//
// GitHub issue #7077: upgrading from v5.18→v5.19 caused a nil pointer dereference
// because the version 1 PriorSchema incorrectly used the current List-based schema
// to decode Set-encoded state.

// TestMigrateAPITokenFromV1_MultiPolicy tests migration from schema_version=1
// (Set-based, v5.16.0) to v501 (List-based) with multiple policies.
// This verifies the version 1 upgrader correctly decodes Set state and sorts
// policies canonically for stable List ordering.
func TestMigrateAPITokenFromV1_MultiPolicy(t *testing.T) {
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
								fmt.Sprintf("cloudflare_api_token.%s", rnd),
								tfjsonpath.New("name"),
								knownvalue.StringExact(rnd),
							),
							statecheck.ExpectKnownValue(
								fmt.Sprintf("cloudflare_api_token.%s", rnd),
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

// TestMigrateAPITokenFromV1_MultiPermGroups tests migration from schema_version=1
// (Set-based) to v501 (List-based) with multiple permission groups.
// This verifies permission groups are sorted canonically within each policy.
func TestMigrateAPITokenFromV1_MultiPermGroups(t *testing.T) {
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
								fmt.Sprintf("cloudflare_api_token.%s", rnd),
								tfjsonpath.New("name"),
								knownvalue.StringExact(rnd),
							),
							statecheck.ExpectKnownValue(
								fmt.Sprintf("cloudflare_api_token.%s", rnd),
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

// TestMigrateAPITokenFromV1_Basic tests the simplest case: a single-policy
// token created in v5.18 upgrading to the current version.
// This is the minimal reproducer for GitHub issue #7077 — the nil pointer
// dereference when the version 1 PriorSchema was List-based but state was
// Set-encoded.
func TestMigrateAPITokenFromV1_Basic(t *testing.T) {
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

			// Single-policy token with one permission group — simplest possible case
			config := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

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
						// Step 1: Create with v5.16/v5.18 (stores schema_version=1, Set-based)
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: config,
					},
					{
						// Step 2: Upgrade to current (v501, List-based)
						// Before the fix, this step would panic with:
						//   nil pointer dereference in UpgradeFromV1
						// because the PriorSchema was List-based but state was Set-encoded.
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						Config:                   config,
						ConfigPlanChecks: resource.ConfigPlanChecks{
							PostApplyPostRefresh: []plancheck.PlanCheck{
								plancheck.ExpectEmptyPlan(),
							},
						},
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(
								fmt.Sprintf("cloudflare_api_token.%s", rnd),
								tfjsonpath.New("name"),
								knownvalue.StringExact(rnd),
							),
							statecheck.ExpectKnownValue(
								fmt.Sprintf("cloudflare_api_token.%s", rnd),
								tfjsonpath.New("policies"),
								knownvalue.ListSizeExact(1),
							),
							statecheck.ExpectKnownValue(
								fmt.Sprintf("cloudflare_api_token.%s", rnd),
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
