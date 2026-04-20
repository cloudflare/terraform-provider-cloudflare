package v500_test

import (
	"context"
	_ "embed"
	"fmt"
	"net/url"
	"os"
	"testing"

	cloudflarev6 "github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/accounts"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Embed migration test configuration files
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_with_status.tf
var v4WithStatusConfig string

//go:embed testdata/v5_with_status.tf
var v5WithStatusConfig string

//go:embed testdata/v5_with_policies.tf
var v5WithPoliciesConfig string

// TestMigrateAccountMember_V4ToV5_Basic tests basic field migrations.
// This verifies:
// 1. Field renames: email_address → email, role_ids → roles
// 2. Both migration paths: from v4 and from v5
//
// Both v4_basic.tf and v5_basic.tf explicitly set status="pending" to prevent
// tf-migrate from injecting status="accepted" (which it does when status is absent).
// Without that explicit value, the injected "accepted" would cause a replace for
// from_v4_latest (RequiresReplaceIfConfigured) and a persistent update drift for
// from_v5 (API returns "pending" after invite, creating a perpetual diff).
func TestMigrateAccountMember_V4ToV5_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, email, roleID string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5 
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, email, roleID string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, accountID, email, roleID)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, accountID, email, roleID string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, accountID, email, roleID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if os.Getenv("TF_ACC") == "" {
				t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
			}

			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()

			// Use random email with cfapi.net domain
			testEmail := fmt.Sprintf("tf-test-%s@cfapi.net", rnd)
			resourceName := "cloudflare_account_member." + rnd

			// Fetch the role ID BEFORE unsetting CLOUDFLARE_API_TOKEN; newCleanupClient
			// falls back to the token when key+email are absent, so clearing it first
			// would leave the call unauthenticated.
			roleID := getTestRoleID(t, accountID, "Administrator Read Only")

			// Temporarily unset CLOUDFLARE_API_TOKEN as the API token won't have
			// permission to manage account members.
			if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
				t.Setenv("CLOUDFLARE_API_TOKEN", "")
			}

			testConfig := tc.configFn(rnd, accountID, testEmail, roleID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			// For v5 tests, use local provider; for v4 tests, use external provider
			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				// Use local v5 provider
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				// Use external v4 provider
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
					setupBaseURLEnv(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					// Step 1: Create resource with source provider version
					firstStep,
					// Step 2: Run migration and verify no-diff state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(testEmail)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("roles"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					}),
				},
			})
		})
	}
}

// TestMigrateAccountMember_V4ToV5_WithStatus tests migration of account members
// with the optional status field to ensure all fields are properly migrated.
func TestMigrateAccountMember_V4ToV5_WithStatus(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, email, roleID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, email, roleID string) string {
				return fmt.Sprintf(v4WithStatusConfig, rnd, accountID, email, roleID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, email, roleID string) string {
				return fmt.Sprintf(v5WithStatusConfig, rnd, accountID, email, roleID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if os.Getenv("TF_ACC") == "" {
				t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
			}

			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

			// Use real test user email that exists in Cloudflare system
			// (required because status="accepted" needs an existing user)
			testEmail := "terraform-test-user-b@cfapi.net"

			// Perform API calls that may need the token BEFORE unsetting it.
			// cleanupExistingMember and getTestRoleID both use newCleanupClient which
			// falls back to CLOUDFLARE_API_TOKEN when key+email are not set; t.Setenv
			// would clear the token making those calls unauthenticated.
			cleanupExistingMember(t, accountID, testEmail)
			roleID := getTestRoleID(t, accountID, "Administrator Read Only")

			// Temporarily unset CLOUDFLARE_API_TOKEN as the API token won't have
			// permission to manage account members.
			if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
				t.Setenv("CLOUDFLARE_API_TOKEN", "")
			}

			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()

			resourceName := "cloudflare_account_member." + rnd
			testConfig := tc.configFn(rnd, accountID, testEmail, roleID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
					setupBaseURLEnv(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					// Step 2: Migrate to v5 provider
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(testEmail)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.StringExact("accepted")),
					}),
					{
						// Step 3: Apply migrated config with v5 provider
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(testEmail)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.StringExact("accepted")),
						},
					},
				},
			})
		})
	}
}

// TestMigrateAccountMember_FromV5_13 tests migration from v5.13.0 (internal v5 state upgrade).
// This tests the stepping stone pattern where early v5 state needs to be upgraded to v500.
func TestMigrateAccountMember_FromV5_13(t *testing.T) {
	t.Skip("API returning id:00000000000000000000000000000000 for the test user.")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	email := fmt.Sprintf("%s@example.com", rnd)
	tmpDir := t.TempDir()
	resourceName := "cloudflare_account_member." + rnd

	// Config uses policies (v5 feature, not available in v4)
	testConfig := fmt.Sprintf(v5WithPoliciesConfig, accountID, accountID, accountID, rnd, accountID, email)

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
				Config: testConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("access"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resource_groups"), knownvalue.ListSizeExact(1)),
				},
			},
			{
				// Step 2: Upgrade to latest provider (state upgrader handles v5 internal migration)
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testConfig,
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

// Helper functions

// newCleanupClient returns a cloudflare-go v6 client for cleanup operations.
// It prefers API key+email auth (same as SharedClient), falling back to API
// token when CLOUDFLARE_API_KEY is not set. The CLOUDFLARE_API_TOKEN env var
// may have been unset earlier in the test to prevent the v5 Terraform provider
// from using it (account-member management requires key+email or a privileged
// token), so we read both env vars without relying on the provider's credential
// chain.
func newCleanupClient() *cloudflarev6.Client {
	apiKey := os.Getenv("CLOUDFLARE_API_KEY")
	apiEmail := os.Getenv("CLOUDFLARE_EMAIL")
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")

	if apiKey != "" && apiEmail != "" {
		return cloudflarev6.NewClient(
			option.WithAPIKey(apiKey),
			option.WithAPIEmail(apiEmail),
		)
	}
	// Fall back to API token when key+email are not configured.
	return cloudflarev6.NewClient(option.WithAPIToken(apiToken))
}

// cleanupExistingMember deletes an existing account member with the given email
// to prevent "Account member already exists" errors from previous test runs.
// It pages through all members (using ListAutoPaging) and matches on either the
// top-level Email field or User.Email, since the Cloudflare API populates these
// fields differently depending on whether the member has accepted the invitation.
func cleanupExistingMember(t *testing.T, accountID, email string) {
	t.Helper()

	ctx := context.Background()
	client := newCleanupClient()

	iter := client.Accounts.Members.ListAutoPaging(ctx, accounts.MemberListParams{
		AccountID: cloudflarev6.String(accountID),
	})

	for iter.Next() {
		member := iter.Current()
		if member.Email == email || member.User.Email == email {
			t.Logf("Deleting existing account member %s (%s) from previous test run", member.ID, email)
			_, err := client.Accounts.Members.Delete(ctx, member.ID, accounts.MemberDeleteParams{
				AccountID: cloudflarev6.String(accountID),
			})
			if err != nil {
				t.Fatalf("Failed to delete existing member %s for cleanup: %v", member.ID, err)
			}
			t.Logf("Successfully deleted existing member %s", member.ID)
			return
		}
	}

	if err := iter.Err(); err != nil {
		t.Fatalf("Failed to list account members for cleanup: %v", err)
	}
}

// getTestRoleID fetches a valid role ID for the account by role name.
func getTestRoleID(t *testing.T, accountID, roleName string) string {
	t.Helper()

	ctx := context.Background()
	client := newCleanupClient()
	roles, err := client.Accounts.Roles.List(ctx, accounts.RoleListParams{
		AccountID: cloudflarev6.String(accountID),
		PerPage:   cloudflarev6.Float(100),
	})
	if err != nil {
		t.Fatalf("Failed to get account roles: %v", err)
	}

	for _, role := range roles.Result {
		if role.Name == roleName {
			return role.ID
		}
	}

	t.Fatalf("Role %q not found for account %s", roleName, accountID)
	return ""
}

// setupBaseURLEnv sets up the legacy CLOUDFLARE_API_HOSTNAME env var if needed.
func setupBaseURLEnv(t *testing.T) {
	t.Helper()
	baseUrl := os.Getenv("CLOUDFLARE_BASE_URL")
	if baseUrl != "" {
		u, err := url.Parse(baseUrl)
		if err != nil {
			t.Fatal(err)
		}
		hostname := u.Hostname()
		os.Setenv("CLOUDFLARE_API_HOSTNAME", hostname)
	}
}
