package list_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

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
		tflog.Info(ctx, "Skipping lists sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	lists, err := client.Rules.Lists.List(ctx, rules.ListListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Lists: %s", err))
		return err
	}
	if len(lists.Result) == 0 {
		tflog.Info(ctx, "No Cloudflare Lists to sweep")
		return nil
	}

	for _, list := range lists.Result {
		// Check both standard test naming and legacy list test prefix
		if !utils.ShouldSweepResource(list.Name) && !strings.HasPrefix(list.Name, listTestPrefix) {
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare List: %s (%s)", list.Name, list.ID))
		_, err := client.Rules.Lists.Delete(ctx, list.ID, rules.ListDeleteParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete Cloudflare List %s (%s): %s", list.Name, list.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted Cloudflare List: %s (%s)", list.Name, list.ID))
	}

	return nil
}

func TestAccCloudflareList(t *testing.T) {

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

func TestAccCloudflareListWithItems_IP(t *testing.T) {

	rndIP := utils.GenerateRandomResourceName()

	resourceNameIP := fmt.Sprintf("cloudflare_list.%s", rndIP)

	descriptionIP := fmt.Sprintf("description.%s", rndIP)

	descriptionIPNew := fmt.Sprintf("%s.new", descriptionIP)

	listNameIP := fmt.Sprintf("%s%s", listTestPrefix, rndIP)

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
				Config:      testAccCheckCloudflareListWithCustomItem(rndIP, listNameIP, descriptionIP, accountID, `ip = "0f::"`),
				ExpectError: regexp.MustCompile(`IP address "0f::" must be normalized: "f::"`),
			},
			{
				Config:      testAccCheckCloudflareListWithCustomItem(rndIP, listNameIP, descriptionIP, accountID, `ip = "1.1.1.1/32"`),
				ExpectError: regexp.MustCompile(`CIDR "1\.1\.1\.1/32" must be represented as "1\.1\.1\.1"`),
			},
			{
				Config:      testAccCheckCloudflareListWithCustomItem(rndIP, listNameIP, descriptionIP, accountID, `ip = "f::/128"`),
				ExpectError: regexp.MustCompile(`CIDR "f::/128" must be represented as "f::"`),
			},
			{
				Config:      testAccCheckCloudflareListWithCustomItem(rndIP, listNameIP, descriptionIP, accountID, `ip = "1.1.1.1/24"`),
				ExpectError: regexp.MustCompile(`CIDR "1\.1\.1\.1/24" must be normalized: "1\.1\.1\.0/24"`),
			},
			{
				Config:      testAccCheckCloudflareListWithCustomItem(rndIP, listNameIP, descriptionIP, accountID, `ip = "f::1/64"`),
				ExpectError: regexp.MustCompile(`CIDR "f::1/64" must be normalized: "f::/64"`),
			},
			{
				Config: testAccCheckCloudflareListWithIPItems(rndIP, listNameIP, descriptionIP, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceNameIP, "name", listNameIP),
					resource.TestCheckResourceAttr(resourceNameIP, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceNameIP, "description", descriptionIP),
					resource.TestCheckResourceAttr(resourceNameIP, "kind", "ip"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceNameIP, "items.*", map[string]string{
						"ip": "1.1.1.1",
					}),
					resource.TestCheckTypeSetElemNestedAttrs(resourceNameIP, "items.*", map[string]string{
						"ip": "1.1.1.2",
					}),
					checkListAndPopulate(resourceNameIP, &list),
				),
			},
			{
				PreConfig: func() {
					initialID = list.ID
				},
				Config: testAccCheckCloudflareListWithIPItems(rndIP, listNameIP, descriptionIPNew, accountID),
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
			// unset items to clear list
			{
				Config: testAccCheckCloudflareListWithNullIPItems(rndIP, listNameIP, descriptionIP, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr(resourceNameIP, "items"),
				),
			},
			{
				ImportState:  true,
				ResourceName: resourceNameIP,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s", accountID, s.RootModule().Resources[resourceNameIP].Primary.ID), nil
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr(resourceNameIP, "items"),
				),
				ImportStateKind: resource.ImportBlockWithID,
			},
		},
	})
}

func TestAccCloudflareListWithItems_Hostname(t *testing.T) {

	rndIP := utils.GenerateRandomResourceName()

	resourceNameIP := fmt.Sprintf("cloudflare_list.%s", rndIP)

	descriptionIP := fmt.Sprintf("description.%s", rndIP)

	descriptionIPNew := fmt.Sprintf("%s.new", descriptionIP)

	listNameIP := fmt.Sprintf("%s%s", listTestPrefix, rndIP)

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
				Config: testAccCheckCloudflareListWithHostnameItems(rndIP, listNameIP, descriptionIP, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceNameIP, "name", listNameIP),
					resource.TestCheckResourceAttr(resourceNameIP, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceNameIP, "description", descriptionIP),
					resource.TestCheckResourceAttr(resourceNameIP, "kind", "hostname"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceNameIP, "items.*", map[string]string{
						"hostname.url_hostname": "example.com",
					}),
					resource.TestCheckTypeSetElemNestedAttrs(resourceNameIP, "items.*", map[string]string{
						"hostname.url_hostname": "*.a.example.com",
					}),
					checkListAndPopulate(resourceNameIP, &list),
				),
			},
			{
				PreConfig: func() {
					initialID = list.ID
				},
				Config: testAccCheckCloudflareListWithHostnameItems(rndIP, listNameIP, descriptionIPNew, accountID),
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
		},
	})
}

func TestAccCloudflareListWithItems_Redirect(t *testing.T) {

	rndIP := utils.GenerateRandomResourceName()

	resourceNameIP := fmt.Sprintf("cloudflare_list.%s", rndIP)

	descriptionIP := fmt.Sprintf("description.%s", rndIP)

	descriptionIPNew := fmt.Sprintf("%s.new", descriptionIP)

	listNameIP := fmt.Sprintf("%s%s", listTestPrefix, rndIP)

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
				Config: testAccCheckCloudflareListWithRedirectItems(rndIP, listNameIP, descriptionIP, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceNameIP, "name", listNameIP),
					resource.TestCheckResourceAttr(resourceNameIP, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceNameIP, "description", descriptionIP),
					resource.TestCheckResourceAttr(resourceNameIP, "kind", "redirect"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceNameIP, "items.*", map[string]string{
						"redirect.source_url": "example.com/1",
					}),
					resource.TestCheckTypeSetElemNestedAttrs(resourceNameIP, "items.*", map[string]string{
						"redirect.source_url": "example.com/2",
					}),
					checkListAndPopulate(resourceNameIP, &list),
				),
			},
			{
				PreConfig: func() {
					initialID = list.ID
				},
				Config: testAccCheckCloudflareListWithRedirectItems(rndIP, listNameIP, descriptionIPNew, accountID),
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
		},
	})
}

func testAccCheckCloudflareList(resourceName, listName, description, accountID, kind string) string {
	return acctest.LoadTestCase("list.tf", resourceName, listName, description, accountID, kind)
}

func testAccCheckCloudflareListWithIPItems(resourceName, listName, description, accountID string) string {
	return acctest.LoadTestCase("listwithipitems.tf", resourceName, listName, description, accountID)
}

func testAccCheckCloudflareListWithCustomItem(resourceName, listName, description, accountID, itemTf string) string {
	return acctest.LoadTestCase("listwithcustomitem.tf", resourceName, listName, description, accountID, itemTf)
}

func testAccCheckCloudflareListWithNullIPItems(resourceName, listName, description, accountID string) string {
	return acctest.LoadTestCase("listwithnullipitems.tf", resourceName, listName, description, accountID)
}

func testAccCheckCloudflareListWithHostnameItems(resourceName, listName, description, accountID string) string {
	return acctest.LoadTestCase("listwithhostnameitems.tf", resourceName, listName, description, accountID)
}

func testAccCheckCloudflareListWithRedirectItems(resourceName, listName, description, accountID string) string {
	return acctest.LoadTestCase("listwithredirectitems.tf", resourceName, listName, description, accountID)
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
