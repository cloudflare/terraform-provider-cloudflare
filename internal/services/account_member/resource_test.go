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

func testCloudflareAccountMemberDirectAdd(accountID, emailAddress string) string {
	return acctest.LoadTestCase("cloudflareaccountmemberdirectadd.tf", accountID, emailAddress)
}
