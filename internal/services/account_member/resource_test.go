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
	acctest.TestAccSkipForDefaultAccount(t, "Account is using domain scoped roles and cannot be used for legacy permissions.")

	// Temporarily unset CLOUDFLARE_API_TOKEN as the API token won't have
	// permission to manage account members.
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
				Config: testCloudflareAccountMemberBasicConfig(rnd, fmt.Sprintf("%s@example.com", rnd), accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "email_address", fmt.Sprintf("%s@example.com", rnd)),
					resource.TestCheckResourceAttr(name, "role_ids.#", "1"),
					resource.TestCheckResourceAttr(name, "role_ids.0", "05784afa30c1afe1440e79d9351c7430"),
				),
			},
		},
	})
}

func TestAccCloudflareAccountMember_DirectAdd(t *testing.T) {
	acctest.TestAccSkipForDefaultAccount(t, "Account is using domain scoped roles and cannot be used for legacy permissions.")

	// Temporarily unset CLOUDFLARE_API_TOKEN as the API token won't have
	// permission to manage account members.
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
				Config: testCloudflareAccountMemberDirectAdd(rnd, "millie@cloudflare.com", accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "email_address", "millie@cloudflare.com"),
					resource.TestCheckResourceAttr(name, "role_ids.#", "1"),
					resource.TestCheckResourceAttr(name, "role_ids.0", "05784afa30c1afe1440e79d9351c7430"),
					resource.TestCheckResourceAttr(name, "status", "accepted"),
				),
			},
		},
	})
}

func testCloudflareAccountMemberBasicConfig(resourceID, emailAddress, accountID string) string {
	return acctest.LoadTestCase("cloudflareaccountmemberbasicconfig.tf", resourceID, emailAddress, accountID)
}

func testCloudflareAccountMemberDirectAdd(resourceID, emailAddress, accountID string) string {
	return acctest.LoadTestCase("cloudflareaccountmemberdirectadd.tf", resourceID, emailAddress, accountID)
}
