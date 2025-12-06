package account_member_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	cloudflarev6 "github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/accounts"
	"github.com/cloudflare/cloudflare-go/v6/iam"
	"github.com/cloudflare/cloudflare-go/v6/zones"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

// Note: Account members are challenging to test with sweepers/CheckDestroy because:
// 1. The API requires special permissions that may not be available with test tokens
// 2. Account members are typically persistent and should not be deleted automatically
// 3. Test members with fake emails may cause API errors when trying to create/manage them
//
// For comprehensive resource lifecycle testing, we rely on the built-in Terraform
// test framework validation and the resource's own Delete implementation.

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_account_member", &resource.Sweeper{
		Name: "cloudflare_account_member",
		F:    testSweepCloudflareAccountMembers,
	})
}

func testSweepCloudflareAccountMembers(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		tflog.Info(ctx, "Skipping account members sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	members, _, err := client.AccountMembers(ctx, accountID, cloudflare.PaginationOptions{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch account members: %s", err))
		return fmt.Errorf("failed to fetch account members: %w", err)
	}

	if len(members) == 0 {
		tflog.Info(ctx, "No account members to sweep")
		return nil
	}

	for _, member := range members {
		// Only sweep test members with @example.com emails or emails matching test patterns
		if !strings.HasSuffix(member.User.Email, "@example.com") && !utils.ShouldSweepResource(member.User.Email) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting account member: %s (email: %s, account: %s)", member.ID, member.User.Email, accountID))
		err := client.DeleteAccountMember(ctx, accountID, member.ID)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete account member %s: %s", member.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted account member: %s", member.ID))
	}

	// List all Resource Groups
	client2 := acctest.SharedClient()
	resourceGroups, err := client2.IAM.ResourceGroups.List(ctx, iam.ResourceGroupListParams{
		AccountID: cloudflarev6.String(accountID),
	})
	if err != nil {
		return fmt.Errorf("failed to fetch account tokens: %w", err)
	}

	for _, resourceGroup := range resourceGroups.Result {
		// Only sweep test resource groups with names matching test patterns
		if utils.ShouldSweepResource(resourceGroup.Name) {
			err = deleteDomainGroup(accountID, resourceGroup.ID)
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete resource group %s: %s", resourceGroup.ID, err))
				continue
			}
			tflog.Info(ctx, fmt.Sprintf("Deleted domain group: %s", resourceGroup.ID))
		}
	}

	return nil
}

func TestAccCloudflareAccountMember_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN as the API token won't have
	// permission to manage account members.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	email := fmt.Sprintf("%s@example.com", rnd)
	resourceName := "cloudflare_account_member.test_member"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareAccountMemberBasicConfig(accountID, email),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.StringExact("pending")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
				},
			},
		},
	})
}

func TestAccCloudflareAccountMember_Import(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN as the API token won't have
	// permission to manage account members.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	email := fmt.Sprintf("%s@example.com", rnd)
	resourceName := "cloudflare_account_member.test_member"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareAccountMemberBasicConfig(accountID, email),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.StringExact("pending")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"policies.0.resource_groups.0.id"},
			},
		},
	})
}

func TestAccCloudflareAccountMember_DirectAdd(t *testing.T) {

	t.Skip("API now throws if the user doesn't exist. We will have to see if we can easily create test users for this test.")

	// Temporarily unset CLOUDFLARE_API_TOKEN as the API token won't have
	// permission to manage account members.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	email := fmt.Sprintf("%s@example.com", rnd)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareAccountMemberDirectAdd(accountID, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_account_member.test_member", consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr("cloudflare_account_member.test_member", "email", email),
					resource.TestCheckResourceAttr("cloudflare_account_member.test_member", "status", "accepted"),
				),
			},
		},
	})
}

func testCloudflareAccountMemberBasicConfig(accountID, emailAddress string) string {
	return acctest.LoadTestCase("cloudflareaccountmemberbasicconfig.tf", accountID, emailAddress)
}

func TestAccCloudflareAccountMember_RolesUpdate(t *testing.T) {
	// Test that roles can be updated in-place without replacement
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
	}
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_account_member.test_member"
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	email := fmt.Sprintf("%s@example.com", rnd)

	initialRoleID := getRoleId(t, accountID, "Administrator")
	updatedRoleID := getRoleId(t, accountID, "Super Administrator - All Privileges")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Create with initial role
				Config: testCloudflareAccountMemberRolesConfig(email, accountID, initialRoleID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("roles"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("roles").AtSliceIndex(0), knownvalue.StringExact(initialRoleID)),
				},
			},
			{
				// Update role in-place (tests custom marshal logic)
				Config: testCloudflareAccountMemberRolesConfig(email, accountID, updatedRoleID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("roles").AtSliceIndex(0), knownvalue.StringExact(updatedRoleID)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("roles"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("roles").AtSliceIndex(0), knownvalue.StringExact(updatedRoleID)),
				},
			},
		},
	})
}

// Test that using roles ignores auto-generated policies to avoid diffs
func TestAccCloudflareAccountMember_RolesVsPolicies(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	if os.Getenv("TF_ACC") == "" {
		t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_account_member.test_member"
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	email := fmt.Sprintf("%s@example.com", rnd)
	roleID := getRoleId(t, accountID, "Administrator")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareAccountMemberRolesConfig(email, accountID, roleID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("roles"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("roles").AtSliceIndex(0), knownvalue.StringExact(roleID)),
					// When using roles, policies are auto-generated but should not cause diffs
					// We verify this by the test completing without plan differences
				},
			},
			{
				// Second apply should not cause any changes (stable state)
				Config: testCloudflareAccountMemberRolesConfig(email, accountID, roleID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("roles"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("roles").AtSliceIndex(0), knownvalue.StringExact(roleID)),
				},
			},
		},
	})
}

func testCloudflareAccountMemberDirectAdd(accountID, emailAddress string) string {
	return acctest.LoadTestCase("cloudflareaccountmemberdirectadd.tf", accountID, emailAddress)
}

func testCloudflareAccountMemberRolesConfig(emailAddress, accountID, roleID string) string {
	return acctest.LoadTestCase("cloudflareaccountmemberrolesconfig.tf", emailAddress, accountID, roleID)
}

func TestAccCloudflareAccountMember_Policies(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN as the API token won't have
	// permission to manage account members.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	if os.Getenv("TF_ACC") == "" {
		t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_account_member.test_member"
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	email := fmt.Sprintf("%s@example.com", rnd)
	permissionGroupID := getPermissionGroupId(t, accountID, "all_privileges")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareAccountMemberPoliciesConfig(accountID, email, permissionGroupID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("access"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact(permissionGroupID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resource_groups"), knownvalue.ListSizeExact(1)),
				},
			},
			{
				// Second apply should not cause any changes (stable state)
				Config: testCloudflareAccountMemberPoliciesConfig(accountID, email, permissionGroupID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("access"), knownvalue.StringExact("allow")),
				},
			},
		},
	})
}

func TestAccCloudflareAccountMember_PoliciesAddResourceGroup(t *testing.T) {
	t.Skip("Needs a DSR enabled user")
	// Temporarily unset CLOUDFLARE_API_TOKEN as the API token won't have
	// permission to manage account members.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	if os.Getenv("TF_ACC") == "" {
		t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_account_member.test_member"
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	email := fmt.Sprintf("%s@example.com", rnd)
	permissionGroupID := getPermissionGroupId(t, accountID, "domain_admin_readonly")

	zones := getDomains(t, accountID)
	if len(zones) < 2 {
		t.Skip("Not enough domains found, need 2 for this test")
	}
	domainGroupID1 := createDomainGroup(t, rnd, accountID, zones[0].ID)
	domainGroupID2 := createDomainGroup(t, rnd, accountID, zones[1].ID)

	t.Cleanup(func() {
		deleteDomainGroup(accountID, domainGroupID1)
		deleteDomainGroup(accountID, domainGroupID2)
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("cloudflare_account_member-add-resource-group1.tf", accountID, email, permissionGroupID, domainGroupID1),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("access"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact(permissionGroupID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resource_groups"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resource_groups").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact(domainGroupID1)),
				},
			},
			{
				// Another apply should not cause any changes (stable state)
				Config: acctest.LoadTestCase("cloudflare_account_member-add-resource-group1.tf", accountID, email, permissionGroupID, domainGroupID1),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("access"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact(permissionGroupID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resource_groups"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resource_groups").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact(domainGroupID1)),
				},
			},
			{
				Config: acctest.LoadTestCase("cloudflare_account_member-add-resource-group2.tf", accountID, email, permissionGroupID, domainGroupID1, domainGroupID2),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("access"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact(permissionGroupID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resource_groups"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resource_groups").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact(domainGroupID1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resource_groups").AtSliceIndex(1).AtMapKey("id"), knownvalue.StringExact(domainGroupID2)),
				},
			},
			{
				// Another apply should not cause any changes (stable state)
				Config: acctest.LoadTestCase("cloudflare_account_member-add-resource-group2.tf", accountID, email, permissionGroupID, domainGroupID1, domainGroupID2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("access"), knownvalue.StringExact("allow")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact(permissionGroupID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resource_groups"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resource_groups").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact(domainGroupID1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resource_groups").AtSliceIndex(1).AtMapKey("id"), knownvalue.StringExact(domainGroupID2)),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

func testCloudflareAccountMemberPoliciesConfig(accountID, emailAddress, permgroupId string) string {
	return acctest.LoadTestCase("cloudflareaccountmemberpoliciesconfig.tf", accountID, emailAddress, permgroupId)
}

func getPermissionGroupId(t *testing.T, accountID string, label string) string {
	ctx := context.Background()
	client := acctest.SharedClient()
	res, err := client.IAM.PermissionGroups.List(ctx, iam.PermissionGroupListParams{
		AccountID: cloudflarev6.String(accountID),
		Label:     cloudflarev6.String(label),
	})
	if err != nil {
		t.Fatalf("Failed to list permission groups: %v for account %s", err, accountID)
	}
	if len(res.Result) == 0 {
		t.Fatalf("Expected at least one permission group with label '%s' but got none for account %s", label, accountID)
	}
	return res.Result[0].ID
}

func getRoleId(t *testing.T, accountID string, label string) string {
	ctx := context.Background()
	client := acctest.SharedClient()
	roles, err := client.Accounts.Roles.List(ctx, accounts.RoleListParams{
		AccountID: cloudflarev6.String(accountID),
		// this may eventually become a problem if there are too many roles
		PerPage: cloudflarev6.Float(100),
	})
	if err != nil {
		t.Fatalf("failed to list roles for account %s: %s", accountID, err)
	}
	if len(roles.Result) == 0 {
		t.Fatalf("no roles available for testing for account %s", accountID)
	}
	var roleID string
	for i, role := range roles.Result {
		if role.Name == label {
			roleID = roles.Result[i].ID
			break
		}
	}
	if roleID == "" {
		t.Fatalf("failed to find '%s' role for testing", label)
	}
	return roleID
}

func getDomains(t *testing.T, accountID string) []zones.Zone {
	ctx := context.Background()
	client := acctest.SharedClient()
	res, err := client.Zones.List(ctx, zones.ZoneListParams{
		Account: cloudflarev6.F(zones.ZoneListParamsAccount{
			ID: cloudflarev6.F(accountID),
		}),
	})
	if err != nil {
		t.Fatalf("Failed to list domains: %v", err)
	}
	return res.Result
}

func createDomainGroup(t *testing.T, rnd, accountID, domainID string) string {
	ctx := context.Background()
	client := acctest.SharedClient()
	domainGroup, err := client.IAM.ResourceGroups.New(ctx, iam.ResourceGroupNewParams{
		AccountID: cloudflarev6.String(accountID),
		Name:      cloudflarev6.String(fmt.Sprintf("%s-%s", rnd, domainID)),
		Scope: cloudflarev6.F(iam.ResourceGroupNewParamsScope{
			Key: cloudflarev6.String(fmt.Sprintf("com.cloudflare.api.account.%s", accountID)),
			Objects: cloudflarev6.F([]iam.ResourceGroupNewParamsScopeObject{
				{
					Key: cloudflarev6.String(fmt.Sprintf("com.cloudflare.api.account.zone.%s", domainID)),
				},
			}),
		}),
	})
	if err != nil {
		t.Fatalf("Failed to create domain group: %v", err)
	}
	return domainGroup.ID
}

func deleteDomainGroup(accountID, domainID string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	_, err := client.IAM.ResourceGroups.Delete(ctx, domainID, iam.ResourceGroupDeleteParams{
		AccountID: cloudflarev6.String(accountID),
	})
	return err
}
