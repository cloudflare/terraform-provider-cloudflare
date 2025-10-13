package zero_trust_list_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	cfv6 "github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
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
	resource.AddTestSweepers("cloudflare_zero_trust_list", &resource.Sweeper{
		Name: "cloudflare_zero_trust_list",
		F:    testSweepCloudflareZeroTrustList,
	})
}

func testSweepCloudflareZeroTrustList(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	
	if accountID == "" {
		return nil
	}
	
	// List all zero trust lists
	page, err := client.ZeroTrust.Gateway.Lists.List(ctx, zero_trust.GatewayListListParams{
		AccountID: cfv6.F(accountID),
	})
	if err != nil {
		log.Printf("[ERROR] Failed to fetch zero trust lists: %s", err)
		return err
	}

	// Delete all lists (sweepers clean up everything from test accounts)
	for _, list := range page.Result {
		_, err := client.ZeroTrust.Gateway.Lists.Delete(ctx, list.ID, zero_trust.GatewayListDeleteParams{
			AccountID: cfv6.F(accountID),
		})
		if err != nil {
			log.Printf("[ERROR] Failed to delete zero trust list %s (%s): %s", list.Name, list.ID, err)
		}
	}

	return nil
}

func TestAccCloudflareTeamsList_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_list.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsListConfigBasic(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("SERIAL")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("My description")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.ListExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"description": knownvalue.Null(),
							"value":       knownvalue.StringExact("asdf-1234"),
						}),
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"description": knownvalue.Null(),
							"value":       knownvalue.StringExact("asdf-5678"),
						}),
					})),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"list_count"},
			},
		},
	})
}

func TestAccCloudflareTeamsList_LottaListItems(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_list.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsListConfigBigItemCount(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("SERIAL")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("My description")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.ListSizeExact(1000)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"list_count"},
			},
		},
	})
}

func TestAccCloudflareTeamsList_Reordered(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_list.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsListConfigBasic(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.ListExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"description": knownvalue.Null(),
							"value":       knownvalue.StringExact("asdf-1234"),
						}),
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"description": knownvalue.Null(),
							"value":       knownvalue.StringExact("asdf-5678"),
						}),
					})),
				},
			},
			{
				Config: testAccCloudflareTeamsListConfigReorderedItems(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.ListExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"description": knownvalue.Null(),
							"value":       knownvalue.StringExact("asdf-1234"),
						}),
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"description": knownvalue.Null(),
							"value":       knownvalue.StringExact("asdf-5678"),
						}),
					})),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"list_count"},
			},
		},
	})
}

func TestAccCloudflareTeamsList_Update(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_list.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsListConfigReorderedItems(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "description", "My description"),
					resource.TestCheckResourceAttr(resourceName, "type", "SERIAL"),
					resource.TestCheckResourceAttr(resourceName, "items.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "items.0.value", "asdf-1234"),
					resource.TestCheckResourceAttr(resourceName, "items.1.value", "asdf-5678"),
				),
			},
			{
				Config: testAccCloudflareTeamsListConfigReorderedItemsUpdate(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "description", "My description update"),
					resource.TestCheckResourceAttr(resourceName, "type", "SERIAL"),
					resource.TestCheckResourceAttr(resourceName, "items.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "items.0.value", "asdf-1234"),
					resource.TestCheckResourceAttr(resourceName, "items.1.value", "bsdf-5678"),
					resource.TestCheckResourceAttr(resourceName, "items.2.value", "csdf-5678"),
				),
			},

			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"list_count"},
			},
		},
	})
}

func TestAccCloudflareTeamsList_DomainType(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_list.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsListConfigDomain(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("DOMAIN")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Domain list for testing")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.SetSizeExact(3)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"list_count"},
			},
		},
	})
}

func TestAccCloudflareTeamsList_URLType(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_list.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsListConfigURL(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("URL")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("URL list for testing")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.SetSizeExact(3)),
				},
				ExpectNonEmptyPlan: true, // URL normalization by API can cause plan drift
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"list_count", "items", "created_at", "updated_at"},
			},
		},
	})
}

func TestAccCloudflareTeamsList_EmailType(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_list.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsListConfigEmail(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("EMAIL")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Email list for testing")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.SetSizeExact(3)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"list_count"},
			},
		},
	})
}

func TestAccCloudflareTeamsList_IPType(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_list.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsListConfigIP(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("IP")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("IP list for testing")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.SetSizeExact(3)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"list_count"},
			},
		},
	})
}

func TestAccCloudflareTeamsList_EmptyList(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_list.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsListConfigEmpty(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("DOMAIN")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Empty list for testing")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.Null()),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"list_count"},
			},
		},
	})
}

func TestAccCloudflareTeamsList_ItemLifecycle(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_list.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsListDestroy,
		Steps: []resource.TestStep{
			// Start with DOMAIN list
			{
				Config: testAccCloudflareTeamsListConfigDomain(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("DOMAIN")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Domain list for testing")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.SetSizeExact(3)),
				},
			},
			// Update items and descriptions
			{
				Config: testAccCloudflareTeamsListConfigUpdatedItems(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("DOMAIN")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Updated list for testing")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.SetSizeExact(2)),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"list_count", "items", "created_at", "updated_at"},
			},
		},
	})
}

func TestAccCloudflareTeamsList_DescriptionUpdates(t *testing.T) {
	// Test for issue #4119: cloudflare_teams_list not updating items_with_description
	// Validates that item descriptions can be added, updated, and removed correctly
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_list.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsListDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create list with mixed items (some with descriptions, some without)
			{
				Config: testAccCloudflareTeamsListConfigDescUpdate1(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("SERIAL")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Testing description updates")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"value":       knownvalue.StringExact("device-001"),
							"description": knownvalue.Null(),
						}),
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"value":       knownvalue.StringExact("device-002"),
							"description": knownvalue.StringExact("Original description"),
						}),
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"value":       knownvalue.StringExact("device-003"),
							"description": knownvalue.StringExact("Will be updated"),
						}),
					})),
				},
			},
			// Step 2: Update descriptions - add, modify, remove, and add new item
			{
				Config: testAccCloudflareTeamsListConfigDescUpdate2(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("SERIAL")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Testing description updates")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"value":       knownvalue.StringExact("device-001"),
							"description": knownvalue.StringExact("Added description"), // Added description
						}),
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"value":       knownvalue.StringExact("device-002"),
							"description": knownvalue.Null(), // Removed description
						}),
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"value":       knownvalue.StringExact("device-003"),
							"description": knownvalue.StringExact("Updated description"), // Modified description
						}),
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"value":       knownvalue.StringExact("device-004"),
							"description": knownvalue.StringExact("New item with description"), // New item
						}),
					})),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"list_count", "created_at", "updated_at"},
			},
		},
	})
}

func testAccCloudflareTeamsListConfigBasic(rnd, accountID string) string {
	return acctest.LoadTestCase("teamslistconfigbasic.tf", rnd, accountID)
}

func testAccCloudflareTeamsListConfigBigItemCount(rnd, accountID string) string {
	items := []string{}
	for i := 0; i < 1000; i++ {
		items = append(items, `{value = "example-`+strconv.Itoa(i)+`"}`)
	}

	return acctest.LoadTestCase("teamslistconfigbigitemcount.tf", rnd, accountID, strings.Join(items, ","))
}

func testAccCloudflareTeamsListConfigReorderedItems(rnd, accountID string) string {
	return acctest.LoadTestCase("teamslistconfigreordereditems.tf", rnd, accountID)
}

func testAccCloudflareTeamsListConfigReorderedItemsUpdate(rnd, accountID string) string {
	return acctest.LoadTestCase("teamslistconfigreordereditems-update.tf", rnd, accountID)
}

func testAccCheckCloudflareTeamsListDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_list" {
			continue
		}

		identifier := cloudflare.AccountIdentifier(rs.Primary.Attributes[consts.AccountIDSchemaKey])
		_, err := client.GetTeamsList(context.Background(), identifier, rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Teams List still exists")
		}
	}

	return nil
}

func testAccCloudflareTeamsListConfigDomain(rnd, accountID string) string {
	return acctest.LoadTestCase("teamslistconfigdomain.tf", rnd, accountID)
}

func testAccCloudflareTeamsListConfigURL(rnd, accountID string) string {
	return acctest.LoadTestCase("teamslistconfigurl.tf", rnd, accountID)
}

func testAccCloudflareTeamsListConfigEmail(rnd, accountID string) string {
	return acctest.LoadTestCase("teamslistconfigemail.tf", rnd, accountID)
}

func testAccCloudflareTeamsListConfigIP(rnd, accountID string) string {
	return acctest.LoadTestCase("teamslistconfigip.tf", rnd, accountID)
}

func testAccCloudflareTeamsListConfigEmpty(rnd, accountID string) string {
	return acctest.LoadTestCase("teamslistconfigempty.tf", rnd, accountID)
}

func testAccCloudflareTeamsListConfigUpdatedItems(rnd, accountID string) string {
	return acctest.LoadTestCase("teamslistconfigupdateditems.tf", rnd, accountID)
}

func testAccCloudflareTeamsListConfigDescUpdate1(rnd, accountID string) string {
	return acctest.LoadTestCase("teamslistconfigdescupdate1.tf", rnd, accountID)
}

func testAccCloudflareTeamsListConfigDescUpdate2(rnd, accountID string) string {
	return acctest.LoadTestCase("teamslistconfigdescupdate2.tf", rnd, accountID)
}
