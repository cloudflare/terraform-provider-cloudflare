package zero_trust_list_test

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go"
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
