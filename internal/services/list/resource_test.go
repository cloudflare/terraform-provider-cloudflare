package list_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const listTestPrefix = "tf_test_list_"

func init() {
	resource.AddTestSweepers("cloudflare_list", &resource.Sweeper{
		Name: "cloudflare_list",
		F:    testSweepCloudflareList,
	})
}

func testSweepCloudflareList(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return errors.New("CLOUDFLARE_ACCOUNT_ID must be set")
	}

	lists, err := client.Rules.Lists.List(ctx, rules.ListListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Lists: %s", err))
	}

	if len(lists.Result) == 0 {
		log.Print("[DEBUG] No Cloudflare Lists to sweep")
		return nil
	}

	for _, list := range lists.Result {
		if !strings.HasPrefix(list.Name, listTestPrefix) {
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare List ID: %s", list.ID))
		//nolint:errcheck
		client.Rules.Lists.Delete(ctx, list.ID, rules.ListDeleteParams{
			AccountID: cloudflare.F(accountID),
		})
	}

	return nil
}

func TestAccCloudflareList(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the IP List
	// endpoint does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rndIP := utils.GenerateRandomResourceName()
	rndRedirect := utils.GenerateRandomResourceName()
	rndASN := utils.GenerateRandomResourceName()
	rndHostname := utils.GenerateRandomResourceName()

	resourceNameIP := fmt.Sprintf("cloudflare_list.%s", rndIP)
	resourceNameRedirect := fmt.Sprintf("cloudflare_list.%s", rndRedirect)
	resourceNameASN := fmt.Sprintf("cloudflare_list.%s", rndASN)
	resourceNameHostname := fmt.Sprintf("cloudflare_list.%s", rndHostname)

	dataResourceNameIP := fmt.Sprintf("data.cloudflare_list.%s", rndIP)
	dataResourceNameRedirect := fmt.Sprintf("data.cloudflare_list.%s", rndRedirect)
	dataResourceNameASN := fmt.Sprintf("data.cloudflare_list.%s", rndASN)
	dataResourceNameHostname := fmt.Sprintf("data.cloudflare_list.%s", rndHostname)

	descriptionIP := fmt.Sprintf("description.%s", rndIP)
	descriptionRedirect := fmt.Sprintf("description.%s", rndRedirect)
	descriptionASN := fmt.Sprintf("description.%s", rndASN)
	descriptionHostname := fmt.Sprintf("description.%s", rndHostname)

	descriptionIPNew := fmt.Sprintf("%s.new", descriptionIP)
	descriptionRedirectNew := fmt.Sprintf("%s.new", descriptionRedirect)
	descriptionASNNew := fmt.Sprintf("%s.new", descriptionASN)
	descriptionHostnameNew := fmt.Sprintf("%s.new", descriptionHostname)

	listNameIP := fmt.Sprintf("%s%s", listTestPrefix, rndIP)
	listNameRedirect := fmt.Sprintf("%s%s", listTestPrefix, rndRedirect)
	listNameASN := fmt.Sprintf("%s%s", listTestPrefix, rndASN)
	listNameHostname := fmt.Sprintf("%s%s", listTestPrefix, rndHostname)

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var list rules.ListsList
	var initialID string

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareList(rndIP, listNameIP, descriptionIP, accountID, "ip"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceNameIP, "name", listNameIP),
					resource.TestCheckResourceAttr(resourceNameIP, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceNameIP, "description", descriptionIP),
					resource.TestCheckResourceAttr(resourceNameIP, "kind", "ip"),
					checkListAndPopulate(resourceNameIP, &list),
				),
			},
			{
				PreConfig: func() {
					initialID = list.ID
				},
				Config: testAccCheckCloudflareList(rndIP, listNameIP, descriptionIPNew, accountID, "ip"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceNameIP, "description", descriptionIPNew),
					checkListAndPopulate(resourceNameIP, &list),
					func(state *terraform.State) error {
						if initialID != list.ID {
							return fmt.Errorf("wanted update but List got recreated (id changed %q -> %q)", initialID, list.ID)
						}
						return nil
					},
				),
			},
			{
				Config: testAccCheckCloudflareListDataSource(rndIP, accountID, listNameIP, descriptionIP, "ip"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataResourceNameIP, "name", listNameIP),
					resource.TestCheckResourceAttr(dataResourceNameIP, "account_id", accountID),
					resource.TestCheckResourceAttr(dataResourceNameIP, "description", descriptionIP),
					resource.TestCheckResourceAttr(dataResourceNameIP, "kind", "ip"),
				),
			},
			{
				ImportState:  true,
				ResourceName: resourceNameIP,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s", accountID, s.RootModule().Resources[resourceNameIP].Primary.ID), nil
				},
				ImportStateVerify: true,
			},
			{
				ImportState:  true,
				ResourceName: resourceNameIP,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s", accountID, s.RootModule().Resources[resourceNameIP].Primary.ID), nil
				},
				ImportStateKind: resource.ImportBlockWithID,
			},
			{
				Config: testAccCheckCloudflareList(rndRedirect, listNameRedirect, descriptionRedirect, accountID, "redirect"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceNameRedirect, "name", listNameRedirect),
					resource.TestCheckResourceAttr(resourceNameRedirect, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceNameRedirect, "description", descriptionRedirect),
					resource.TestCheckResourceAttr(resourceNameRedirect, "kind", "redirect"),
					checkListAndPopulate(resourceNameRedirect, &list),
					func(s *terraform.State) error {
						if _, exists := s.RootModule().Resources[resourceNameIP]; exists {
							return fmt.Errorf("Expected old list to be destroyed and removed from state, %q", resourceNameIP)
						}
						return nil
					},
				),
			},
			{
				PreConfig: func() {
					initialID = list.ID
				},
				Config: testAccCheckCloudflareList(rndRedirect, listNameRedirect, descriptionRedirectNew, accountID, "redirect"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceNameRedirect, "description", descriptionRedirectNew),
					checkListAndPopulate(resourceNameRedirect, &list),
					func(state *terraform.State) error {
						if initialID != list.ID {
							return fmt.Errorf("wanted update but List got recreated (id changed %q -> %q)", initialID, list.ID)
						}
						return nil
					},
				),
			},
			{
				Config: testAccCheckCloudflareListDataSource(rndRedirect, accountID, listNameRedirect, descriptionRedirect, "redirect"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataResourceNameRedirect, "name", listNameRedirect),
					resource.TestCheckResourceAttr(dataResourceNameRedirect, "account_id", accountID),
					resource.TestCheckResourceAttr(dataResourceNameRedirect, "description", descriptionRedirect),
					resource.TestCheckResourceAttr(dataResourceNameRedirect, "kind", "redirect"),
				),
			},
			{
				ImportState:  true,
				ResourceName: resourceNameRedirect,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s", accountID, s.RootModule().Resources[resourceNameRedirect].Primary.ID), nil
				},
				ImportStateVerify: true,
			},
			{
				ImportState:  true,
				ResourceName: resourceNameRedirect,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s", accountID, s.RootModule().Resources[resourceNameRedirect].Primary.ID), nil
				},
				ImportStateKind: resource.ImportBlockWithID,
			},
			{
				Config: testAccCheckCloudflareList(rndASN, listNameASN, descriptionASN, accountID, "asn"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceNameASN, "name", listNameASN),
					resource.TestCheckResourceAttr(resourceNameASN, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceNameASN, "description", descriptionASN),
					resource.TestCheckResourceAttr(resourceNameASN, "kind", "asn"),
					checkListAndPopulate(resourceNameASN, &list),
					func(s *terraform.State) error {
						if _, exists := s.RootModule().Resources[resourceNameRedirect]; exists {
							return fmt.Errorf("Expected old list to be destroyed and removed from state, %q", resourceNameRedirect)
						}
						return nil
					},
				),
			},
			{
				PreConfig: func() {
					initialID = list.ID
				},
				Config: testAccCheckCloudflareList(rndASN, listNameASN, descriptionASNNew, accountID, "asn"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceNameASN, "description", descriptionASNNew),
					checkListAndPopulate(resourceNameASN, &list),
					func(state *terraform.State) error {
						if initialID != list.ID {
							return fmt.Errorf("wanted update but List got recreated (id changed %q -> %q)", initialID, list.ID)
						}
						return nil
					},
				),
			},
			{
				Config: testAccCheckCloudflareListDataSource(rndASN, accountID, listNameASN, descriptionASN, "asn"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataResourceNameASN, "name", listNameASN),
					resource.TestCheckResourceAttr(dataResourceNameASN, "account_id", accountID),
					resource.TestCheckResourceAttr(dataResourceNameASN, "description", descriptionASN),
					resource.TestCheckResourceAttr(dataResourceNameASN, "kind", "asn"),
				),
			},
			{
				ImportState:  true,
				ResourceName: resourceNameASN,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s", accountID, s.RootModule().Resources[resourceNameASN].Primary.ID), nil
				},
				ImportStateVerify: true,
			},
			{
				ImportState:  true,
				ResourceName: resourceNameASN,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s", accountID, s.RootModule().Resources[resourceNameASN].Primary.ID), nil
				},
				ImportStateKind: resource.ImportBlockWithID,
			},
			{
				Config: testAccCheckCloudflareList(rndHostname, listNameHostname, descriptionHostname, accountID, "hostname"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceNameHostname, "name", listNameHostname),
					resource.TestCheckResourceAttr(resourceNameHostname, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceNameHostname, "description", descriptionHostname),
					resource.TestCheckResourceAttr(resourceNameHostname, "kind", "hostname"),
					checkListAndPopulate(resourceNameHostname, &list),
					func(s *terraform.State) error {
						if _, exists := s.RootModule().Resources[resourceNameASN]; exists {
							return fmt.Errorf("Expected old list to be destroyed and removed from state, %q", resourceNameASN)
						}
						return nil
					},
				),
			},
			{
				Config: testAccCheckCloudflareListDataSource(rndHostname, accountID, listNameHostname, descriptionHostname, "hostname"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataResourceNameHostname, "name", listNameHostname),
					resource.TestCheckResourceAttr(dataResourceNameHostname, "account_id", accountID),
					resource.TestCheckResourceAttr(dataResourceNameHostname, "description", descriptionHostname),
					resource.TestCheckResourceAttr(dataResourceNameHostname, "kind", "hostname"),
				),
			},
			{
				PreConfig: func() {
					initialID = list.ID
				},
				Config: testAccCheckCloudflareList(rndHostname, listNameHostname, descriptionHostnameNew, accountID, "hostname"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceNameHostname, "description", descriptionHostnameNew),
					checkListAndPopulate(resourceNameHostname, &list),
					func(state *terraform.State) error {
						if initialID != list.ID {
							return fmt.Errorf("wanted update but List got recreated (id changed %q -> %q)", initialID, list.ID)
						}
						return nil
					},
				),
			},
			{
				ImportState:  true,
				ResourceName: resourceNameHostname,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s", accountID, s.RootModule().Resources[resourceNameHostname].Primary.ID), nil
				},
				ImportStateVerify: true,
			},
			{
				ImportState:  true,
				ResourceName: resourceNameHostname,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s", accountID, s.RootModule().Resources[resourceNameHostname].Primary.ID), nil
				},
				ImportStateKind: resource.ImportBlockWithID,
			},
		},
	})
}

func testAccCheckCloudflareList(resourceName, listName, description, accountID, kind string) string {
	return acctest.LoadTestCase("list.tf", resourceName, listName, description, accountID, kind)
}

func testAccCheckCloudflareListDataSource(resourceName, accountID, listName, description, kind string) string {
	return acctest.LoadTestCase("listdatasource.tf", resourceName, accountID, listName, description, kind)
}

func checkListAndPopulate(resourceName string, list *rules.ListsList) func(*terraform.State) error {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("List ID is not set")
		}

		numItems, err := strconv.Atoi(rs.Primary.Attributes["num_items"])
		if err != nil {
			return fmt.Errorf("failed parsing num_items: %w", err)
		}
		numReferencingFilters, err := strconv.Atoi(rs.Primary.Attributes["num_referencing_filters"])
		if err != nil {
			return fmt.Errorf("failed parsing num_referencing_filters: %w", err)
		}

		var kind rules.ListsListKind
		switch rs.Primary.Attributes["kind"] {
		case "ip":
			kind = rules.ListsListKindIP
		case "asn":
			kind = rules.ListsListKindASN
		case "redirect":
			kind = rules.ListsListKindRedirect
		case "hostname":
			kind = rules.ListsListKindHostname
		}

		*list = rules.ListsList{
			ID:                    rs.Primary.ID,
			Name:                  rs.Primary.Attributes["name"],
			Description:           rs.Primary.Attributes["description"],
			Kind:                  kind,
			NumItems:              float64(numItems),
			NumReferencingFilters: float64(numReferencingFilters),
			CreatedOn:             rs.Primary.Attributes["created_on"],
			ModifiedOn:            rs.Primary.Attributes["modified_on"],
		}

		return nil
	}
}
