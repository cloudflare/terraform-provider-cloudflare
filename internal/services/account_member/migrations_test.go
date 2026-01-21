package account_member_test

import (
	"fmt"
	"net/url"
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

// Note: Account members are challenging to test with CheckDestroy because:
// 1. The API requires special permissions that may not be available with test tokens
// 2. Account members are typically persistent and should not be deleted automatically
// 3. Test members with fake emails may cause API errors when trying to create/manage them
//
// For migration testing, we rely on the built-in Terraform test framework validation.

// TestMigrateCloudflareAccountMember_Migration_Basic_MultiVersion tests the account member
// migration with simple field renames. This test ensures that:
// 1. email_address field is renamed to email
// 2. role_ids field is renamed to roles
// 3. The migration tool successfully transforms both configuration and state files
// 4. Resources remain functional after migration without requiring manual intervention
func TestMigrateCloudflareAccountMember_Migration_Basic_MultiVersion(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(accountID, rnd, email, roleID string) string
	}{
		{
			name:     "from_v4_52_1", // Last v4 release
			version:  "4.52.1",
			configFn: testAccCloudflareAccountMemberMigrationConfigV4Basic,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Skip if acceptance tests are not enabled
			if os.Getenv("TF_ACC") == "" {
				t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
			}

			// Temporarily unset CLOUDFLARE_API_TOKEN as the API token won't have
			// permission to manage account members.
			if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
				t.Setenv("CLOUDFLARE_API_TOKEN", "")
			}

			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			if accountID == "" {
				t.Fatal("CLOUDFLARE_ACCOUNT_ID must be set for this acceptance test.")
			}
			rnd := utils.GenerateRandomResourceName()
			// Use real test user email that exists in Cloudflare system
			email := "terraform-test-user-a@cfapi.net"
			resourceName := "cloudflare_account_member." + rnd
			roleID := getRoleId(t, accountID, "Administrator Read Only")
			testConfig := tc.configFn(accountID, rnd, email, roleID)
			tmpDir := t.TempDir()

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
					baseUrl := os.Getenv("CLOUDFLARE_BASE_URL")
					if baseUrl != "" {
						u, err := url.Parse(baseUrl)
						if err != nil {
							t.Fatal(err)
						}
						hostname := u.Hostname()
						// legacy env var for base url
						os.Setenv("CLOUDFLARE_API_HOSTNAME", hostname)
					}
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create account member with v4 provider
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								VersionConstraint: tc.version,
								Source:            "cloudflare/cloudflare",
							},
						},
						Config: testConfig,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email_address"), knownvalue.StringExact(email)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("role_ids"), knownvalue.NotNull()),
						},
					},
					// Step 2: Migrate to v5 provider
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, "v4", "v5", []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(email)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("roles"), knownvalue.NotNull()),
					}),
					{
						// Step 3: Apply the migrated configuration with v5 provider
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(email)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("roles"), knownvalue.NotNull()),
						},
					},
				},
			})
		})
	}
}

// TestMigrateCloudflareAccountMember_Migration_WithStatus tests migration of account members
// with the optional status field to ensure all fields are properly migrated.
func TestMigrateCloudflareAccountMember_Migration_WithStatus(t *testing.T) {
	// Skip if acceptance tests are not enabled
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
	}

	// Temporarily unset CLOUDFLARE_API_TOKEN as the API token won't have
	// permission to manage account members.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Fatal("CLOUDFLARE_ACCOUNT_ID must be set for this acceptance test.")
	}
	rnd := utils.GenerateRandomResourceName()
	// Use real test user email that exists in Cloudflare system
	email := "terraform-test-user-b@cfapi.net"
	resourceName := "cloudflare_account_member." + rnd
	roleID := getRoleId(t, accountID, "Administrator Read Only")
	v4Config := testAccCloudflareAccountMemberMigrationConfigV4WithStatus(accountID, rnd, email, roleID)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			baseUrl := os.Getenv("CLOUDFLARE_BASE_URL")
			if baseUrl != "" {
				u, err := url.Parse(baseUrl)
				if err != nil {
					t.Fatal(err)
				}
				hostname := u.Hostname()
				// legacy env var for base url
				os.Setenv("CLOUDFLARE_API_HOSTNAME", hostname)
			}
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create account member with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email_address"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.StringExact("accepted")),
				},
			},
			// Step 2: Migrate to v5 provider
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(email)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.StringExact("accepted")),
			}),
			{
				// Step 3: Apply migrated config with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.StringExact("accepted")),
				},
			},
		},
	})
}

// V4 Configuration Functions

func testAccCloudflareAccountMemberMigrationConfigV4Basic(accountID, rnd, email, roleID string) string {
	return fmt.Sprintf(`
resource "cloudflare_account_member" "%[2]s" {
  account_id    = "%[1]s"
  email_address = "%[3]s"
  role_ids      = ["%[4]s"]
}
`, accountID, rnd, email, roleID)
}

func testAccCloudflareAccountMemberMigrationConfigV4WithStatus(accountID, rnd, email, roleID string) string {
	return fmt.Sprintf(`
resource "cloudflare_account_member" "%[2]s" {
  account_id    = "%[1]s"
  email_address = "%[3]s"
  status        = "accepted"
  role_ids      = ["%[4]s"]
}
`, accountID, rnd, email, roleID)
}

// TestMigrateAccountMemberFromV5_13 tests migration from v5.13.0
func TestMigrateAccountMemberFromV5_13(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	email := fmt.Sprintf("%s@example.com", rnd)
	tmpDir := t.TempDir()

	// config in both test changes is the same, we are just ensuring that the
	// state file can be migrated
	config := fmt.Sprintf(`
data "cloudflare_resource_groups" "all" {
  account_id = "%[1]s"
  name       = "com.cloudflare.api.account.%[1]s"
}

data "cloudflare_account_permission_groups" "all" {
  account_id = "%[1]s"
}

locals {
  api_token_permissions_groups_map = {
    for perm in data.cloudflare_account_permission_groups.all.result :
    perm.name => perm.id
  }
}

resource "cloudflare_account_member" "test_member" {
  account_id = "%[1]s"
  email      = "%[2]s"
  policies = [{
    access = "allow"
    resource_groups = [{
      id : data.cloudflare_resource_groups.all.result[0].id
    }]
    permission_groups = [{
      id : local.api_token_permissions_groups_map["Administrator"]
    }]
  }]
}

`, accountID, email)

	resourceName := "cloudflare_account_member.test_member"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5.13 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.13.0",
					},
				},
				Config: config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
					// we should have a policy ID
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("access"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resource_groups"), knownvalue.ListSizeExact(1)),
				},
			},
			{
				// Step 2: Upgrade to latest provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("access"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resource_groups"), knownvalue.ListSizeExact(1)),
				},
			},
		},
	})
}
