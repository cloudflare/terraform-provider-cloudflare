package list_item_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	_ "github.com/cloudflare/terraform-provider-cloudflare/internal/services/list"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_list_item", &resource.Sweeper{
		Name: "cloudflare_list_item",
		F:    testSweepCloudflareListItems,
	})
}

func testSweepCloudflareListItems(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		tflog.Info(ctx, "Skipping list items sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	// List all lists using v6 client
	listsResp, err := client.Rules.Lists.List(ctx, rules.ListListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch lists: %s", err))
		return fmt.Errorf("failed to fetch lists: %w", err)
	}

	if len(listsResp.Result) == 0 {
		tflog.Info(ctx, "No lists found, skipping list items sweep")
		return nil
	}

	// For each list, get and delete its items
	for _, list := range listsResp.Result {
		// Only process lists that would be swept themselves
		if !utils.ShouldSweepResource(list.Name) {
			continue
		}

		// List items for this list using auto-paging
		itemsPager := client.Rules.Lists.Items.ListAutoPaging(ctx, list.ID, rules.ListItemListParams{
			AccountID: cloudflare.F(accountID),
		})

		type deleteItem struct {
			ID string `json:"id"`
		}
		type deletePayload struct {
			Items []deleteItem `json:"items"`
		}

		var itemsToDelete []deleteItem
		for itemsPager.Next() {
			item := itemsPager.Current()
			itemsToDelete = append(itemsToDelete, deleteItem{
				ID: item.ID,
			})
		}

		if err := itemsPager.Err(); err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to fetch items for list %s: %s", list.Name, err))
			continue
		}

		if len(itemsToDelete) == 0 {
			continue
		}

		payload := deletePayload{Items: itemsToDelete}
		deleteBody, _ := json.Marshal(payload)

		tflog.Info(ctx, fmt.Sprintf("Deleting %d items from list: %s (account: %s)", len(itemsToDelete), list.Name, accountID))
		_, err = client.Rules.Lists.Items.Delete(ctx, list.ID, rules.ListItemDeleteParams{
			AccountID: cloudflare.F(accountID),
		}, option.WithRequestBody("application/json", deleteBody))
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete items from list %s: %s", list.Name, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted %d items from list: %s", len(itemsToDelete), list.Name))
	}

	return nil
}

func TestAccCloudflareListItem_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareIPListItem(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "ip", "192.0.2.0"),
				),
			},
		},
	})
}

func TestAccCloudflareListItem_Import(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	itemName := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	listName := fmt.Sprintf("cloudflare_list.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	ip := "192.0.2.0"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareIPListItem(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(itemName, "ip", ip),
					resource.TestCheckResourceAttr(itemName, "comment", rnd),
				),
			},
			{
				ResourceName:      itemName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources[listName]
					if !ok {
						return "", fmt.Errorf("list resource not found: %s", listName)
					}
					listID := rs.Primary.ID
					rs, ok = s.RootModule().Resources[itemName]
					if !ok {
						return "", fmt.Errorf("list_item resource not found: %s", itemName)
					}
					itemID := rs.Primary.ID
					return fmt.Sprintf("%s/%s/%s", accountID, listID, itemID), nil
				},
			},
			{
				ResourceName:    itemName,
				ImportState:     true,
				ImportStateKind: resource.ImportBlockWithID,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources[listName]
					if !ok {
						return "", fmt.Errorf("list resource not found: %s", listName)
					}
					listID := rs.Primary.ID
					rs, ok = s.RootModule().Resources[itemName]
					if !ok {
						return "", fmt.Errorf("list_item resource not found: %s", itemName)
					}
					itemID := rs.Primary.ID
					return fmt.Sprintf("%s/%s/%s", accountID, listID, itemID), nil
				},
			},
		},
	})
}

func TestAccCloudflareListItem_MultipleItems(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareIPListItemMultipleEntries(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name+"_1", "ip", "192.0.2.0"),
					resource.TestCheckResourceAttr(name+"_2", "ip", "192.0.2.1"),
					resource.TestCheckResourceAttr(name+"_3", "ip", "192.0.2.2"),
					resource.TestCheckResourceAttr(name+"_4", "ip", "192.0.2.3"),
				),
			},
		},
	})
}

func TestAccCloudflareListItem_MultipleItemsHostname(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHostnameListItemMultipleEntries(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name+"_1", "hostname.url_hostname", "a.example.com"),
					resource.TestCheckResourceAttr(name+"_2", "hostname.url_hostname", "example.com"),
				),
			},
			{
				Config: testAccCheckCloudflareHostnameListItemMultipleEntries(rnd, rnd, rnd+"-updated", accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name+"_1", "hostname.url_hostname", "a.example.com"),
					resource.TestCheckResourceAttr(name+"_1", "comment", rnd+"-updated"),
					resource.TestCheckResourceAttr(name+"_2", "hostname.url_hostname", "example.com"),
					resource.TestCheckResourceAttr(name+"_2", "comment", rnd+"-updated"),
				),
			},
		},
	})
}

func TestAccCloudflareListItem_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var listItemID string

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareIPListItem(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "ip", "192.0.2.0"),
					func(s *terraform.State) error {
						listItemID = s.RootModule().Resources[name].Primary.Attributes["id"]
						return nil
					},
				),
			},
			{
				Config: testAccCheckCloudflareIPListItem(rnd, rnd, rnd+"-updated", accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "comment", rnd+"-updated"),
					func(s *terraform.State) error {
						newID := s.RootModule().Resources[name].Primary.Attributes["id"]
						if newID == listItemID {
							return fmt.Errorf("ID of list item did not change when updating comment")
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccCloudflareListItem_UpdateHostname(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var listItemID string

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHostnameListItem(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "hostname.url_hostname", "example.com"),
					resource.TestCheckResourceAttr(name, "comment", rnd),
					func(s *terraform.State) error {
						listItemID = s.RootModule().Resources[name].Primary.Attributes["id"]
						return nil
					},
				),
			},
			{
				Config: testAccCheckCloudflareHostnameListItem(rnd, rnd, rnd+"-updated", accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "hostname.url_hostname", "example.com"),
					resource.TestCheckResourceAttr(name, "comment", rnd+"-updated"),
					func(s *terraform.State) error {
						newID := s.RootModule().Resources[name].Primary.Attributes["id"]
						if newID == listItemID {
							return fmt.Errorf("ID of list item did not change when updating comment")
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccCloudflareListItem_UpdateReplace(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareIPListItemNewIp(rnd, rnd, rnd, accountID, "192.0.2.0"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "ip", "192.0.2.0"),
				),
			},
			{
				Config: testAccCheckCloudflareIPListItemNewIp(rnd, rnd, rnd, accountID, "192.0.2.1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "ip", "192.0.2.1"),
				),
			},
		},
	})
}

func TestAccCloudflareListItem_ASN(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareASNListItem(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						name, "asn", "1"),
				),
			},
		},
	})
}

func TestAccCloudflareListItem_Hostname(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHostnameListItem(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "hostname.url_hostname", "example.com"),
				),
			},
		},
	})
}

func TestAccCloudflareListItem_Redirect(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHostnameRedirectItem(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "redirect.source_url", "example.com/"),
					resource.TestCheckResourceAttr(name, "redirect.target_url", "https://example1.com"),
					resource.TestCheckResourceAttr(name, "redirect.status_code", "301"),
				),
			},
		},
	})
}

func TestAccCloudflareListItem_RedirectWithLocalsMap(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	redirect1Resource := fmt.Sprintf("cloudflare_list_item.%s_redirect1", rnd)
	redirect2Resource := fmt.Sprintf("cloudflare_list_item.%s_redirect2", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHostnameRedirectItemLocalsMap(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(redirect1Resource, "redirect.source_url", rnd+"-redirect1.cfapi.net/"),
					resource.TestCheckResourceAttr(redirect1Resource, "redirect.target_url", "https://target1.com"),
					resource.TestCheckResourceAttr(redirect1Resource, "redirect.status_code", "301"),
					resource.TestCheckResourceAttr(redirect1Resource, "comment", rnd+"-redirect1"),
					resource.TestCheckResourceAttr(redirect2Resource, "redirect.source_url", rnd+"-redirect2.cfapi.net/"),
					resource.TestCheckResourceAttr(redirect2Resource, "redirect.target_url", "https://target2.com"),
					resource.TestCheckResourceAttr(redirect2Resource, "redirect.status_code", "302"),
					resource.TestCheckResourceAttr(redirect2Resource, "comment", rnd+"-redirect2"),
				),
			},
		},
	})
}

func testAccCheckCloudflareIPListItem(ID, name, comment, accountID string) string {
	return acctest.LoadTestCase("iplistitem.tf", ID, name, comment, accountID)
}

func testAccCheckCloudflareIPListItemNewIp(ID, name, comment, accountID, ip string) string {
	return acctest.LoadTestCase("iplistitem_newip.tf", ID, name, comment, accountID, ip)
}

func testAccCheckCloudflareIPListItemMultipleEntries(ID, name, comment, accountID string) string {
	return acctest.LoadTestCase("iplistitemmultipleentries.tf", ID, name, comment, accountID)
}

func testAccCheckCloudflareHostnameListItemMultipleEntries(ID, name, comment, accountID string) string {
	return acctest.LoadTestCase("hostnamelistitemmultipleentries.tf", ID, name, comment, accountID)
}

func testAccCheckCloudflareBadListItemType(ID, name, comment, accountID string) string {
	return acctest.LoadTestCase("badlistitemtype.tf", ID, name, comment, accountID)
}

func testAccCheckCloudflareASNListItem(ID, name, comment, accountID string) string {
	return acctest.LoadTestCase("asnlistitem.tf", ID, name, comment, accountID)
}

func testAccCheckCloudflareHostnameListItem(ID, name, comment, accountID string) string {
	return acctest.LoadTestCase("hostnamelistitem.tf", ID, name, comment, accountID)
}

func testAccCheckCloudflareHostnameRedirectItem(ID, name, comment, accountID string) string {
	return acctest.LoadTestCase("hostnameredirectitem.tf", ID, name, comment, accountID)
}

func TestAccCloudflareListItem_RedirectWithOverlappingSourceURL(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	firstResource := fmt.Sprintf("cloudflare_list_item.%s_1", rnd)
	secondResource := fmt.Sprintf("cloudflare_list_item.%s_2", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHostnameRedirectWithOverlappingSourceURL(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(firstResource, "redirect.source_url", "www.site1.com/"),
					resource.TestCheckResourceAttr(firstResource, "redirect.target_url", "https://example.com"),
					resource.TestCheckResourceAttr(firstResource, "redirect.status_code", "301"),
					resource.TestCheckResourceAttr(secondResource, "redirect.source_url", "www.site1.com/test"),
					resource.TestCheckResourceAttr(secondResource, "redirect.target_url", "https://cloudflare.com"),
					resource.TestCheckResourceAttr(secondResource, "redirect.status_code", "301"),
				),
			},
		},
	})
}

func testAccCheckCloudflareHostnameRedirectWithOverlappingSourceURL(ID, name, comment, accountID string) string {
	return acctest.LoadTestCase("hostnameredirectwithoverlappingsourceurl.tf", ID, name, comment, accountID)
}

func testAccCheckCloudflareHostnameRedirectItemLocalsMap(ID, name, comment, accountID string) string {
	return acctest.LoadTestCase("hostnameredirectitemlocalsmap.tf", ID, name, comment, accountID)
}
