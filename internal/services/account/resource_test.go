package account_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	cfaccounts "github.com/cloudflare/cloudflare-go/v6/accounts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
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
	resource.AddTestSweepers("cloudflare_account", &resource.Sweeper{
		Name: "cloudflare_account",
		F:    testSweepCloudflareAccount,
	})
}

func testSweepCloudflareAccount(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	// List all accounts to find test accounts to sweep
	accountsResp, err := client.Accounts.List(ctx, cfaccounts.AccountListParams{})
	if err != nil {
		return fmt.Errorf("failed to list accounts: %w", err)
	}

	for _, account := range accountsResp.Result {
		// Only delete accounts that look like test accounts (contain "tf-acc-test" prefix)
		if strings.Contains(account.Name, "tf-acc-test") {
			_, err := client.Accounts.Delete(ctx, cfaccounts.AccountDeleteParams{
				AccountID: cloudflare.F(account.ID),
			})
			if err != nil {
				return fmt.Errorf("failed to delete account %s: %w", account.ID, err)
			}
		}
	}

	return nil
}

func TestAccCloudflareAccount_Basic(t *testing.T) {
	acctest.TestAccSkipForDefaultAccount(t, "Pending PT-792 to address underlying issue.")

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_account.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccountDestroy,
		Steps: []resource.TestStep{
			// Create step
			{
				Config: testAccCheckCloudflareAccountName(rnd, fmt.Sprintf("%s_old", rnd)),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s_old", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("enforce_twofactor"), knownvalue.Bool(false)),
				},
			},
			// Update step
			{
				Config: testAccCheckCloudflareAccountName(rnd, fmt.Sprintf("%s_new", rnd)),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s_new", rnd))),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s_new", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("enforce_twofactor"), knownvalue.Bool(false)),
				},
			},
			// 2FA update step
			{
				Config: testAccCheckCloudflareAccountWith2FA(rnd, rnd),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("enforce_twofactor"), knownvalue.Bool(true)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("enforce_twofactor"), knownvalue.Bool(true)),
				},
			},
			// Import step
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
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
	resourceName := fmt.Sprintf("cloudflare_account.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccountDestroy,
		Steps: []resource.TestStep{
			// The POST endpoint ignores the settings on the create so we
			// need to first create and then update for 2FA enforcement.
			// Tracking with the service team via PT-792.
			{
				Config: testAccCheckCloudflareAccountName(rnd, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("enforce_twofactor"), knownvalue.Bool(false)),
				},
			},
			{
				Config: testAccCheckCloudflareAccountWith2FA(rnd, rnd),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("enforce_twofactor"), knownvalue.Bool(true)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("enforce_twofactor"), knownvalue.Bool(true)),
				},
			},
			{
				Config: testAccCheckCloudflareAccountName(rnd, rnd),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("enforce_twofactor"), knownvalue.Bool(false)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("enforce_twofactor"), knownvalue.Bool(false)),
				},
			},
			// Import step
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCloudflareAccountWith2FA(rnd, name string) string {
	return acctest.LoadTestCase("accountwith2fa.tf", rnd, name)
}

func testAccCheckCloudflareAccountDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_account" {
			continue
		}

		_, err := client.Accounts.Get(context.Background(), cfaccounts.AccountGetParams{
			AccountID: cloudflare.F(rs.Primary.ID),
		})
		if err == nil {
			return fmt.Errorf("account %s still exists", rs.Primary.ID)
		}
	}

	return nil
}
