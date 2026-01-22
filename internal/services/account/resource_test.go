package account_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	cfaccounts "github.com/cloudflare/cloudflare-go/v6/accounts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
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
		tflog.Error(ctx, fmt.Sprintf("Failed to list accounts: %s", err))
		return err
	}

	if len(accountsResp.Result) == 0 {
		tflog.Info(ctx, "No Cloudflare accounts to sweep")
		return nil
	}

	for _, account := range accountsResp.Result {
		if !utils.ShouldSweepResource(account.Name) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting account: %s (%s)", account.Name, account.ID))
		_, err := client.Accounts.Delete(ctx, cfaccounts.AccountDeleteParams{
			AccountID: cloudflare.F(account.ID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete account %s (%s): %s", account.Name, account.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted account: %s (%s)", account.Name, account.ID))
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
			// Create an enterprise account
			{
				Config: testAccCheckCloudflareAccountWithType(rnd, fmt.Sprintf("%s_old", rnd), "enterprise"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s_old", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("enterprise")),
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
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("enterprise")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("enforce_twofactor"), knownvalue.Bool(false)),
				},
			},
			// 2FA update step
			{
				Config: testAccCheckCloudflareAccountWith2FA(rnd, rnd, true),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("enforce_twofactor"), knownvalue.Bool(true)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("enterprise")),
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
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("standard")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("enforce_twofactor"), knownvalue.Bool(false)),
				},
			},
			{
				Config: testAccCheckCloudflareAccountWith2FA(rnd, rnd, true),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("enforce_twofactor"), knownvalue.Bool(true)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("enforce_twofactor"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("standard")),
				},
			},
			{
				Config: testAccCheckCloudflareAccountWith2FA(rnd, rnd, false),
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

func TestAccCloudflareAccount_WithMulti(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_account.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccountDestroy,
		Steps: []resource.TestStep{
			// Create an enterprise account
			{
				Config: testAccCheckCloudflareAccountWithType(rnd, rnd, "enterprise"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("enterprise")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("enforce_twofactor"), knownvalue.Bool(false)),
				},
			},
			// Update to add abuse email
			{
				Config: testAccCheckCloudflareAccountWithMulti(rnd, rnd, "enterprise", false, "abuse@example.com"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("enterprise")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("abuse_contact_email"), knownvalue.StringExact("abuse@example.com")),
				},
			},
			// Update to enable 2FA
			{
				Config: testAccCheckCloudflareAccountWithMulti(rnd, rnd, "enterprise", true, "abuse@example.com"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("enterprise")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("enforce_twofactor"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("abuse_contact_email"), knownvalue.StringExact("abuse@example.com")),
				},
			},
			// Update email
			{
				Config: testAccCheckCloudflareAccountWithMulti(rnd, rnd, "enterprise", true, "updated@example.com"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("enterprise")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("settings").AtMapKey("abuse_contact_email"), knownvalue.StringExact("updated@example.com")),
				},
			},
			// Import
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareAccount_WithUnit(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_account.%s", rnd)

	unitID := os.Getenv("CLOUDFLARE_UNIT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccountWithUnit(rnd, rnd, unitID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("standard")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("unit").AtMapKey("id"), knownvalue.StringExact(unitID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_by").AtMapKey("parent_org_id"), knownvalue.StringExact(unitID)),
				},
			},
		},
	})
}

func testAccCheckCloudflareAccountName(rnd, name string) string {
	return acctest.LoadTestCase("accountname.tf", rnd, name)
}

func testAccCheckCloudflareAccountWithType(rnd, name, account_type string) string {
	return acctest.LoadTestCase("accountwithtype.tf", rnd, name, account_type)
}

func testAccCheckCloudflareAccountWith2FA(rnd, name string, enforce_twofactor bool) string {
	return acctest.LoadTestCase("accountwith2fa.tf", rnd, name, enforce_twofactor)
}

func testAccCheckCloudflareAccountWithMulti(rnd, name, accountType string, enforce2FA bool, email string) string {
	return acctest.LoadTestCase("accountwithmulti.tf", rnd, name, accountType, enforce2FA, email)
}

func testAccCheckCloudflareAccountWithUnit(rnd, name, unitID string) string {
	return acctest.LoadTestCase("accountwithunit.tf", rnd, name, unitID)
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
