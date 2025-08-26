package account_member_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareAccountMember_Basic(t *testing.T) {
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
				Config: testCloudflareAccountMemberBasicConfig(accountID, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_account_member.test_member", consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr("cloudflare_account_member.test_member", "email", email),
					resource.TestCheckResourceAttr("cloudflare_account_member.test_member", "status", "pending"),
					resource.TestCheckResourceAttr("cloudflare_account_member.test_member", "policies.#", "1"),
				),
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
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareAccountMemberBasicConfig(accountID, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudflare_account_member.test_member", consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr("cloudflare_account_member.test_member", "email", email),
					resource.TestCheckResourceAttr("cloudflare_account_member.test_member", "status", "pending"),
					resource.TestCheckResourceAttr("cloudflare_account_member.test_member", "policies.#", "1"),
				),
			},
			{
				ResourceName:        "cloudflare_account_member.test_member",
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
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

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_account_member." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Create with initial role
				Config: testCloudflareAccountMemberRolesConfig(rnd, fmt.Sprintf("%s@example.com", rnd), accountID, "05784afa30c1afe1440e79d9351c7430"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "email", fmt.Sprintf("%s@example.com", rnd)),
					resource.TestCheckResourceAttr(name, "roles.#", "1"),
					resource.TestCheckResourceAttr(name, "roles.0", "05784afa30c1afe1440e79d9351c7430"),
				),
			},
			{
				// Update role in-place (tests custom marshal logic)
				Config: testCloudflareAccountMemberRolesConfig(rnd, fmt.Sprintf("%s@example.com", rnd), accountID, "33666b9c79b9a5273fc7344ff42f953d"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "email", fmt.Sprintf("%s@example.com", rnd)),
					resource.TestCheckResourceAttr(name, "roles.#", "1"),
					resource.TestCheckResourceAttr(name, "roles.0", "33666b9c79b9a5273fc7344ff42f953d"),
				),
			},
		},
	})
}

func TestAccCloudflareAccountMember_RolesVsPolicies(t *testing.T) {
	// Test that using roles ignores auto-generated policies to avoid diffs
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_account_member." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareAccountMemberRolesConfig(rnd, fmt.Sprintf("%s@example.com", rnd), accountID, "05784afa30c1afe1440e79d9351c7430"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "email", fmt.Sprintf("%s@example.com", rnd)),
					resource.TestCheckResourceAttr(name, "roles.#", "1"),
					resource.TestCheckResourceAttr(name, "roles.0", "05784afa30c1afe1440e79d9351c7430"),
					// When using roles, policies are auto-generated but should not cause diffs
					// We verify this by the test completing without plan differences
				),
			},
			{
				// Second apply should not cause any changes (stable state)
				Config: testCloudflareAccountMemberRolesConfig(rnd, fmt.Sprintf("%s@example.com", rnd), accountID, "05784afa30c1afe1440e79d9351c7430"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "roles.#", "1"),
					resource.TestCheckResourceAttr(name, "roles.0", "05784afa30c1afe1440e79d9351c7430"),
				),
			},
		},
	})
}

func testCloudflareAccountMemberDirectAdd(accountID, emailAddress string) string {
	return acctest.LoadTestCase("cloudflareaccountmemberdirectadd.tf", accountID, emailAddress)
}

func testCloudflareAccountMemberRolesConfig(resourceID, emailAddress, accountID, roleID string) string {
	return acctest.LoadTestCase("cloudflareaccountmemberrolesconfig.tf", resourceID, emailAddress, accountID, roleID)
}

func TestAccCloudflareAccountMember_Policies(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_account_member." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareAccountMemberPoliciesConfig(accountID, rnd, accountID, fmt.Sprintf("%s@example.com", rnd), "8e23b19e4e0d44c29d239c5688ba8cbb"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "email", fmt.Sprintf("%s@example.com", rnd)),
					resource.TestCheckResourceAttr(name, "policies.#", "1"),
					resource.TestCheckResourceAttr(name, "policies.0.access", "allow"),
					resource.TestCheckResourceAttr(name, "policies.0.permission_groups.#", "1"),
					resource.TestCheckResourceAttr(name, "policies.0.permission_groups.0.id", "8e23b19e4e0d44c29d239c5688ba8cbb"),
					resource.TestCheckResourceAttr(name, "policies.0.resource_groups.#", "1"),
					resource.TestCheckResourceAttrSet(name, "policies.0.resource_groups.0.id"),
					// Roles should not be set when using policies
					resource.TestCheckNoResourceAttr(name, "roles.0"),
				),
			},
			{
				// Second apply should not cause any changes (stable state)
				Config: testCloudflareAccountMemberPoliciesConfig(accountID, rnd, accountID, fmt.Sprintf("%s@example.com", rnd), "8e23b19e4e0d44c29d239c5688ba8cbb"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "policies.#", "1"),
					resource.TestCheckResourceAttr(name, "policies.0.access", "allow"),
				),
			},
		},
	})
}

func testCloudflareAccountMemberPoliciesConfig(dataSourceAccountID, resourceID, accountID, emailAddress, permissionGroupID string) string {
	return acctest.LoadTestCase("cloudflareaccountmemberpoliciesconfig.tf", dataSourceAccountID, resourceID, accountID, emailAddress, permissionGroupID)
}
