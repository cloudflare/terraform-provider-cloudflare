package account_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	cfaccounts "github.com/cloudflare/cloudflare-go/v6/accounts"
	"github.com/cloudflare/cloudflare-go/v6/organizations"
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
	acctest.TestAccSkipForDefaultAccount(t, "Requires account creation permissions not available on default test account.")

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
	acctest.TestAccSkipForDefaultAccount(t, "Requires account creation permissions not available on default test account.")

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
	t.Skip(`Skipped: 403 Forbidden {"success":false,"errors":[{"code":1002,"message":"Forbidden. Account creation is not allowed"}],"messages":[],"result":null}`)
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
	t.Skip(`Skipped: 403 Forbidden {"success":false,"errors":[{"code":1002,"message":"Forbidden. Account creation is not allowed"}],"messages":[],"result":null}`)
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_account.%s", rnd)

	// Get organization IDs to determine unit_id and alternate_unit_id
	orgIDs, err := getOrganizationIDs(t)
	if err != nil {
		t.Fatalf("Failed to get organization IDs: %v", err)
	}
	unitID := orgIDs[0]

	// For alternate: use second org if exists, otherwise "invalid-unit-id"
	alternateUnitID := "invalid-unit-id"
	if len(orgIDs) > 1 {
		alternateUnitID = orgIDs[1]
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccountWithUnit(rnd, rnd, unitID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "type", "standard"),
					resource.TestCheckResourceAttrSet(resourceName, "unit.id"),
					resource.TestCheckResourceAttrSet(resourceName, "managed_by.parent_org_id"),
					// Verify that unit.id and managed_by.parent_org_id match
					func(s *terraform.State) error {
						rs := s.RootModule().Resources[resourceName]
						unitIDAttr := rs.Primary.Attributes["unit.id"]
						parentOrgIDAttr := rs.Primary.Attributes["managed_by.parent_org_id"]
						if unitIDAttr != parentOrgIDAttr {
							return fmt.Errorf("unit.id (%s) does not match managed_by.parent_org_id (%s)", unitIDAttr, parentOrgIDAttr)
						}
						return nil
					},
				),
			},
			// Import step
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Changing unit.id should force replacement (destroy before create)
			// Use second org (if exists) or "invalid-unit-id"
			func() resource.TestStep {
				step := resource.TestStep{
					Config: testAccCheckCloudflareAccountWithUnit(rnd, rnd, alternateUnitID),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
						},
					},
				}

				if len(orgIDs) < 2 {
					// Only 1 org: using "invalid-unit-id", expect error
					step.ExpectError = regexp.MustCompile(`failed to make http request|403|not found|invalid|Forbidden`)
				} else {
					// 2+ orgs: using second org, verify successful update
					step.Check = resource.ComposeTestCheckFunc(
						func(s *terraform.State) error {
							rs := s.RootModule().Resources[resourceName]
							unitIDAttr := rs.Primary.Attributes["unit.id"]
							parentOrgIDAttr := rs.Primary.Attributes["managed_by.parent_org_id"]

							if unitIDAttr != alternateUnitID {
								return fmt.Errorf("expected unit.id to be %s, got %s", alternateUnitID, unitIDAttr)
							}
							if parentOrgIDAttr != alternateUnitID {
								return fmt.Errorf("expected managed_by.parent_org_id to be %s, got %s", alternateUnitID, parentOrgIDAttr)
							}
							return nil
						},
					)
				}

				return step
			}(),
		},
	})
}

func TestAccCloudflareAccount_InvalidType(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckCloudflareAccountWithType(rnd, rnd, "invalid_type"),
				ExpectError: regexp.MustCompile(`Attribute type value must be one of`),
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

// getOrganizationIDs fetches all organizations and returns their IDs
func getOrganizationIDs(t *testing.T) ([]string, error) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
	}

	client := acctest.SharedClient()
	orgsResp, err := client.Organizations.List(context.Background(), organizations.OrganizationListParams{})
	if err != nil {
		return nil, fmt.Errorf("failed to list organizations: %w", err)
	}

	if len(orgsResp.Result) == 0 {
		return nil, fmt.Errorf("no organization IDs found - this test requires a tenant account with organizations")
	}

	var orgIDs []string
	for _, org := range orgsResp.Result {
		orgIDs = append(orgIDs, org.ID)
	}
	return orgIDs, nil
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

func TestAccUpgradeAccount_FromPublishedV5(t *testing.T) {
	t.Skip(`Skipped: 403 Forbidden {"success":false,"errors":[{"code":1002,"message":"Forbidden. Account creation is not allowed"}],"messages":[],"result":null}`)
	rnd := utils.GenerateRandomResourceName()

	config := testAccCheckCloudflareAccountWithType(rnd, fmt.Sprintf("%s_old", rnd), "enterprise")

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.16.0",
					},
				},
				Config: config,
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}
