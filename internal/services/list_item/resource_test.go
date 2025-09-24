package list_item_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

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
	redirect1Resource := fmt.Sprintf("cloudflare_list_item.%s[\"redirect1\"]", rnd)
	redirect2Resource := fmt.Sprintf("cloudflare_list_item.%s[\"redirect2\"]", rnd)
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
					resource.TestCheckResourceAttr(redirect1Resource, "redirect.source_url", "example1.com/"),
					resource.TestCheckResourceAttr(redirect1Resource, "redirect.target_url", "https://target1.com"),
					resource.TestCheckResourceAttr(redirect1Resource, "redirect.status_code", "301"),
					resource.TestCheckResourceAttr(redirect1Resource, "comment", rnd+"-redirect1"),
					resource.TestCheckResourceAttr(redirect2Resource, "redirect.source_url", "example2.com/"),
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
