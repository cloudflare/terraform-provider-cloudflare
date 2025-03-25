package account_test

import (
	"fmt"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareAccount_Basic(t *testing.T) {
	acctest.TestAccSkipForDefaultAccount(t, "Pending PT-792 to address underlying issue.")

	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_account.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccountName(rnd, fmt.Sprintf("%s_old", rnd)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("%s_old", rnd)),
					resource.TestCheckResourceAttr(name, "enforce_twofactor", "false"),
				),
			},
			{
				Config: testAccCheckCloudflareAccountName(rnd, fmt.Sprintf("%s_new", rnd)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("%s_new", rnd)),
					resource.TestCheckResourceAttr(name, "enforce_twofactor", "false"),
				),
			},
			{
				Config: testAccCheckCloudflareAccountWith2FA(rnd, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "enforce_twofactor", "true"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccountName(rnd, name string) string {
	return acctest.LoadTestCase("accountname.tf", rnd, name)
}

func TestAccCloudflareAccount_2FAEnforced(t *testing.T) {
	acctest.TestAccSkipForDefaultAccount(t, "Pending PT-792 to address underlying issue.")

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_account.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// The POST endpoint ignores the settings on the create so we
			// need to first create and then update for 2FA enforcement.
			// Tracking with the service team via PT-792.
			{
				Config: testAccCheckCloudflareAccountName(rnd, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "enforce_twofactor", "false"),
				),
			},
			{
				Config: testAccCheckCloudflareAccountWith2FA(rnd, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "enforce_twofactor", "true"),
				),
			},
			{
				Config: testAccCheckCloudflareAccountName(rnd, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "enforce_twofactor", "false"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccountWith2FA(rnd, name string) string {
	return acctest.LoadTestCase("accountwith2fa.tf", rnd, name)
}
