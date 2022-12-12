package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareAccountMemberBasic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN as the API token won't have
	// permission to manage account members.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := "cloudflare_account_member." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckAccount(t)
			testAccPreCheckEmail(t)
			testAccPreCheckApiKey(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareAccountMemberBasicConfig(rnd, fmt.Sprintf("%s@example.com", rnd), accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "email_address", fmt.Sprintf("%s@example.com", rnd)),
					resource.TestCheckResourceAttr(name, "role_ids.#", "1"),
					resource.TestCheckResourceAttr(name, "role_ids.0", "05784afa30c1afe1440e79d9351c7430"),
				),
			},
		},
	})
}

func testCloudflareAccountMemberBasicConfig(resourceID, emailAddress, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_account_member" "%[1]s" {
	account_id = "%[3]s"
    email_address = "%[2]s"
    role_ids = [ "05784afa30c1afe1440e79d9351c7430" ]
  }`, resourceID, emailAddress, accountID)
}
