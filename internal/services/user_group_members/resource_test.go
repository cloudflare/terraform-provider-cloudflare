package user_group_members_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/accounts"
	"github.com/cloudflare/cloudflare-go/v6/iam"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_user_group_members", &resource.Sweeper{
		Name: "cloudflare_user_group_members",
		F:    testSweepCloudflareUserGroupMembers,
	})
}

// testSweepCloudflareUserGroupMembers cleans up test user groups; deleting the parent
// group cascades to its members.
func testSweepCloudflareUserGroupMembers(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		tflog.Info(ctx, "Skipping user_group_members sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	// List user groups and delete any leftover test groups
	groups, err := client.IAM.UserGroups.List(
		ctx,
		iam.UserGroupListParams{
			AccountID: cloudflare.F(accountID),
		},
	)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list user groups: %s", err))
		return fmt.Errorf("failed to list user groups: %w", err)
	}

	for _, group := range groups.Result {
		// Only sweep test groups (both DSR check groups and test user groups)
		if !strings.HasPrefix(group.Name, "cftftest") &&
			!strings.HasPrefix(group.Name, "tf-") &&
			!utils.ShouldSweepResource(group.Name) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting test user group: %s (%s)", group.ID, group.Name))
		_, err := client.IAM.UserGroups.Delete(
			ctx,
			group.ID,
			iam.UserGroupDeleteParams{
				AccountID: cloudflare.F(accountID),
			},
		)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete user group %s: %s", group.ID, err))
		}
	}

	return nil
}

// tryAddMember attempts to add a member to a user group to verify DSR is enabled.
// Returns nil if successful, an error if the member cannot be added.
func tryAddMember(ctx context.Context, client *cloudflare.Client, accountID, userGroupID, memberID string) error {
	body := fmt.Sprintf(`[{"id":"%s"}]`, memberID)
	_, err := client.IAM.UserGroups.Members.New(ctx, userGroupID,
		iam.UserGroupMemberNewParams{
			AccountID: cloudflare.F(accountID),
		},
		option.WithRequestBody("application/json", []byte(body)),
	)
	return err
}

// findDSREnabledMembers returns up to `count` member IDs that can be added to a user
// group, skipping the test if none are found. DSR (zone_level_access_beta flag) is not
// exposed in the API, so we detect it by trying to add each member to a temp group.
func findDSREnabledMembers(t *testing.T, accountID string, count int) []string {
	t.Helper()
	ctx := context.Background()
	client := acctest.SharedClient()

	// Create a temporary user group for testing DSR status
	tempGroupName := fmt.Sprintf("tf-dsr-check-%s", utils.GenerateRandomResourceName())
	tempGroup, err := client.IAM.UserGroups.New(ctx, iam.UserGroupNewParams{
		AccountID: cloudflare.F(accountID),
		Name:      cloudflare.F(tempGroupName),
	})
	if err != nil {
		t.Skipf("Failed to create temp user group for DSR check: %s", err)
	}

	// Clean up temp group at the end
	defer func() {
		_, cleanupErr := client.IAM.UserGroups.Delete(ctx, tempGroup.ID, iam.UserGroupDeleteParams{
			AccountID: cloudflare.F(accountID),
		})
		if cleanupErr != nil {
			tflog.Warn(ctx, fmt.Sprintf("Failed to cleanup temp user group %s: %s", tempGroup.ID, cleanupErr))
		}
	}()

	dsrMembers := []string{}

	// Try CLOUDFLARE_MEMBER_ID first if set
	if envMemberID := os.Getenv("CLOUDFLARE_MEMBER_ID"); envMemberID != "" {
		if err := tryAddMember(ctx, client, accountID, tempGroup.ID, envMemberID); err == nil {
			tflog.Info(ctx, fmt.Sprintf("Using CLOUDFLARE_MEMBER_ID from env: %s", envMemberID))
			dsrMembers = append(dsrMembers, envMemberID)
			if len(dsrMembers) >= count {
				return dsrMembers
			}
		}
	}

	// List account members
	members, err := client.Accounts.Members.List(ctx, accounts.MemberListParams{
		AccountID: cloudflare.F(accountID),
		Status:    cloudflare.F(accounts.MemberListParamsStatusAccepted),
	})
	if err != nil {
		t.Skipf("Failed to list account members: %s", err)
	}
	if len(members.Result) == 0 {
		t.Skip("Account has no accepted members")
	}

	// Try each member - skip already found, collect DSR-enabled ones
	for _, member := range members.Result {
		// Skip if already in our list
		alreadyFound := false
		for _, id := range dsrMembers {
			if id == member.ID {
				alreadyFound = true
				break
			}
		}
		if alreadyFound {
			continue
		}

		err := tryAddMember(ctx, client, accountID, tempGroup.ID, member.ID)
		if err == nil {
			tflog.Info(ctx, fmt.Sprintf("Found DSR-enabled member: %s (%s)", member.ID, member.User.Email))
			dsrMembers = append(dsrMembers, member.ID)
			if len(dsrMembers) >= count {
				return dsrMembers
			}
			continue
		}
		if strings.Contains(err.Error(), "do not have DSR enabled") {
			tflog.Info(ctx, fmt.Sprintf("Member %s does not have DSR enabled, trying next", member.ID))
			continue
		}
		tflog.Warn(ctx, fmt.Sprintf("Failed to add member %s: %s", member.ID, err))
	}

	if len(dsrMembers) == 0 {
		t.Skip("No DSR-enabled member found in account; cannot test user_group_members")
	}

	return dsrMembers
}

// findDSREnabledMember is a convenience wrapper for a single DSR-enabled member.
func findDSREnabledMember(t *testing.T, accountID string) string {
	t.Helper()
	members := findDSREnabledMembers(t, accountID, 1)
	if len(members) == 0 {
		t.Skip("No DSR-enabled member found")
	}
	return members[0]
}

// Note: To represent a user group with no members, omit this resource entirely. Setting
// members = [] causes refresh drift; future versions may enforce min 1 member.

// TestAccCloudflareUserGroupMembers_Basic tests create + idempotency with one member.
func TestAccCloudflareUserGroupMembers_Basic(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for acceptance tests")
	}
	memberID := findDSREnabledMember(t, accountID)
	resourceName := "cloudflare_user_group_members." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareUserGroupMembersDestroy,
		Steps: []resource.TestStep{
			// Create with member
			{
				Config: testAccCloudflareUserGroupMembersBasicConfig(rnd, accountID, memberID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("user_group_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("members"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("members").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact(memberID)),
				},
			},
			// Idempotency: re-apply is a no-op
			{
				Config: testAccCloudflareUserGroupMembersBasicConfig(rnd, accountID, memberID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Import
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

// TestAccCloudflareUserGroupMembers_MultipleMembers tests creating a user group with multiple members.
func TestAccCloudflareUserGroupMembers_MultipleMembers(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for acceptance tests")
	}
	members := findDSREnabledMembers(t, accountID, 2)
	if len(members) < 2 {
		t.Skipf("Need at least 2 DSR-enabled members, found %d", len(members))
	}
	resourceName := "cloudflare_user_group_members." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareUserGroupMembersDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareUserGroupMembersMultipleConfig(rnd, accountID, members[0], members[1]),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("members"), knownvalue.ListSizeExact(2)),
				},
			},
			// Import (multi-member)
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

// TestAccCloudflareUserGroupMembers_UpdateAddMember tests adding a member via update.
func TestAccCloudflareUserGroupMembers_UpdateAddMember(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for acceptance tests")
	}
	members := findDSREnabledMembers(t, accountID, 2)
	if len(members) < 2 {
		t.Skipf("Need at least 2 DSR-enabled members, found %d", len(members))
	}
	resourceName := "cloudflare_user_group_members." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareUserGroupMembersDestroy,
		Steps: []resource.TestStep{
			// Start with 1 member
			{
				Config: testAccCloudflareUserGroupMembersBasicConfig(rnd, accountID, members[0]),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("members"), knownvalue.ListSizeExact(1)),
				},
			},
			// Add a second member
			{
				Config: testAccCloudflareUserGroupMembersMultipleConfig(rnd, accountID, members[0], members[1]),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("members"), knownvalue.ListSizeExact(2)),
				},
			},
			// Back to 1 member
			{
				Config: testAccCloudflareUserGroupMembersBasicConfig(rnd, accountID, members[0]),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("members"), knownvalue.ListSizeExact(1)),
				},
			},
			// Import (after multi-member updates)
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

// TestAccCloudflareUserGroupMembers_DestroyRemovesAllMembers verifies destroying just
// the members resource clears all members (via PUT []) while keeping the parent group.
func TestAccCloudflareUserGroupMembers_DestroyRemovesAllMembers(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for acceptance tests")
	}
	memberID := findDSREnabledMember(t, accountID)
	resourceName := "cloudflare_user_group_members." + rnd
	groupResourceName := "cloudflare_user_group." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck_AccountID(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareUserGroupMembersDestroy,
		Steps: []resource.TestStep{
			// Create group with members
			{
				Config: testAccCloudflareUserGroupMembersBasicConfig(rnd, accountID, memberID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("members"), knownvalue.ListSizeExact(1)),
				},
			},
			// Remove only the members resource; group must remain with zero members.
			{
				Config: testAccCloudflareUserGroupOnlyConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(groupResourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				},
				Check: testAccCheckCloudflareUserGroupHasNoMembers(accountID, groupResourceName),
			},
		},
	})
}

// Note: `members = []` causes "inconsistent result after apply" because the framework
// serializes empty `*[]*T` as null at the cty layer. Users wanting an empty group should
// omit the resource or `terraform destroy -target` it. Switching the schema to
// SetNestedAttribute may fix this and the permission_groups ordering drift.

func testAccCloudflareUserGroupMembersBasicConfig(rnd, accountID, memberID string) string {
	return acctest.LoadTestCase("usergroupmembersbasic.tf", rnd, accountID, memberID)
}

func testAccCloudflareUserGroupMembersMultipleConfig(rnd, accountID, memberID1, memberID2 string) string {
	return acctest.LoadTestCase("usergroupmembersmultiple.tf", rnd, accountID, memberID1, memberID2)
}

func testAccCloudflareUserGroupOnlyConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("usergrouponly.tf", rnd, accountID)
}

// testAccCheckCloudflareUserGroupMembersDestroy verifies the parent user group is gone
// from the API after destroy; removing the group cascades to its members.
func testAccCheckCloudflareUserGroupMembersDestroy(s *terraform.State) error {
	client := acctest.SharedClient()
	ctx := context.Background()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_user_group" {
			continue
		}

		_, err := client.IAM.UserGroups.Get(ctx, rs.Primary.ID, iam.UserGroupGetParams{
			AccountID: cloudflare.F(rs.Primary.Attributes[consts.AccountIDSchemaKey]),
		})
		if err == nil {
			return fmt.Errorf("user group %s still exists after destroy", rs.Primary.ID)
		}
	}

	return nil
}

// testAccCheckCloudflareUserGroupHasNoMembers asserts via the API that the given user
// group has zero members.
func testAccCheckCloudflareUserGroupHasNoMembers(accountID, groupResourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[groupResourceName]
		if !ok {
			return fmt.Errorf("user group resource %s not found in state", groupResourceName)
		}
		userGroupID := rs.Primary.ID
		client := acctest.SharedClient()
		ctx := context.Background()

		members, err := client.IAM.UserGroups.Members.List(ctx, userGroupID, iam.UserGroupMemberListParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			return fmt.Errorf("failed to list members for user group %s: %w", userGroupID, err)
		}
		if len(members.Result) > 0 {
			return fmt.Errorf("expected user group %s to have no members, found %d", userGroupID, len(members.Result))
		}
		return nil
	}
}
