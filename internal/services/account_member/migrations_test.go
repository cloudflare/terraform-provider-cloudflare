package account_member_test

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
		configFn func(accountID, rnd, email string) string
	}{
		{
			name:     "from_v4_52_1", // Last v4 release
			version:  "4.52.1",
			configFn: testAccCloudflareAccountMemberMigrationConfigV4Basic,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Temporarily unset CLOUDFLARE_API_TOKEN as the API token won't have
			// permission to manage account members.
			if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
				t.Setenv("CLOUDFLARE_API_TOKEN", "")
			}

			accountID := acctest.TestAccCloudflareAccountID
			rnd := utils.GenerateRandomResourceName()
			email := fmt.Sprintf("test-%s@example.com", rnd)
			resourceName := "cloudflare_account_member." + rnd
			testConfig := tc.configFn(accountID, rnd, email)
			tmpDir := t.TempDir()

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
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
					acctest.MigrationTestStep(t, testConfig, tmpDir, tc.version, []statecheck.StateCheck{
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
	// Temporarily unset CLOUDFLARE_API_TOKEN as the API token won't have
	// permission to manage account members.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := acctest.TestAccCloudflareAccountID
	rnd := utils.GenerateRandomResourceName()
	email := fmt.Sprintf("test-%s@example.com", rnd)
	resourceName := "cloudflare_account_member." + rnd
	v4Config := testAccCloudflareAccountMemberMigrationConfigV4WithStatus(accountID, rnd, email)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
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
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
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

func testAccCloudflareAccountMemberMigrationConfigV4Basic(accountID, rnd, email string) string {
	// Using a standard "Administrator Read Only" role ID from Cloudflare's predefined roles
	roleID := "05784afa30c1afe1440e79d9351c7430"
	return fmt.Sprintf(`
resource "cloudflare_account_member" "%[2]s" {
  account_id    = "%[1]s"
  email_address = "%[3]s"
  role_ids      = ["%[4]s"]
}
`, accountID, rnd, email, roleID)
}

func testAccCloudflareAccountMemberMigrationConfigV4WithStatus(accountID, rnd, email string) string {
	// Using a standard "Administrator Read Only" role ID from Cloudflare's predefined roles
	roleID := "05784afa30c1afe1440e79d9351c7430"
	return fmt.Sprintf(`
resource "cloudflare_account_member" "%[2]s" {
  account_id    = "%[1]s"
  email_address = "%[3]s"
  status        = "accepted"
  role_ids      = ["%[4]s"]
}
`, accountID, rnd, email, roleID)
}
