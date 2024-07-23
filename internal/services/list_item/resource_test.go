package list_item_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareListItem_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var ListItem cfv1.ListItem

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareIPListItem(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListItemExists(name, rnd, &ListItem),
					resource.TestCheckResourceAttr(
						name, "ip", "192.0.2.0"),
				),
			},
		},
	})
}

func TestAccCloudflareListItem_MultipleItems(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var ListItem cfv1.ListItem

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareIPListItemMultipleEntries(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListItemExists(name+"_1", rnd, &ListItem),
					resource.TestCheckResourceAttr(name+"_1", "ip", "192.0.2.0"),

					testAccCheckCloudflareListItemExists(name+"_2", rnd, &ListItem),
					resource.TestCheckResourceAttr(name+"_2", "ip", "192.0.2.1"),

					testAccCheckCloudflareListItemExists(name+"_3", rnd, &ListItem),
					resource.TestCheckResourceAttr(name+"_3", "ip", "192.0.2.2"),

					testAccCheckCloudflareListItemExists(name+"_4", rnd, &ListItem),
					resource.TestCheckResourceAttr(name+"_4", "ip", "192.0.2.3"),
				),
			},
		},
	})
}

func TestAccCloudflareListItem_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var listItem cfv1.ListItem

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareIPListItem(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListItemExists(name, rnd, &listItem),
					resource.TestCheckResourceAttr(
						name, "ip", "192.0.2.0"),
				),
			},
			{
				Config: testAccCheckCloudflareIPListItem(rnd, rnd, rnd+"-updated", accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListItemExists(name, rnd, &listItem),
					resource.TestCheckResourceAttr(name, "comment", rnd+"-updated"),
				),
			},
		},
	})
}

func TestAccCloudflareListItem_BadListItemType(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckCloudflareBadListItemType(rnd, rnd, rnd, accountID),
				ExpectError: regexp.MustCompile(" can not be added to lists of type "),
			},
		},
	})
}

func TestAccCloudflareListItem_ASN(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var ListItem cfv1.ListItem

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareASNListItem(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListItemExists(name, rnd, &ListItem),
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

	var ListItem cfv1.ListItem

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHostnameListItem(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListItemExists(name, rnd, &ListItem),
					resource.TestCheckResourceAttr(
						name, "hostname.#", "1"),
					resource.TestCheckResourceAttr(
						name, "hostname.0.url_hostname", "example.com"),
				),
			},
		},
	})
}

func TestAccCloudflareListItem_Redirect(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var ListItem cfv1.ListItem

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHostnameRedirectItem(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListItemExists(name, rnd, &ListItem),
					resource.TestCheckResourceAttr(
						name, "redirect.#", "1"),
					resource.TestCheckResourceAttr(
						name, "redirect.0.source_url", "example.com/"),
					resource.TestCheckResourceAttr(
						name, "redirect.0.target_url", "https://example1.com"),
					resource.TestCheckResourceAttr(
						name, "redirect.0.status_code", "301"),
				),
			},
		},
	})
}

func testAccCheckCloudflareIPListItem(ID, name, comment, accountID string) string {
	return acctest.LoadTestCase("iplistitem.tf", ID, name, comment, accountID)
}

func testAccCheckCloudflareIPListItemMultipleEntries(ID, name, comment, accountID string) string {
	return acctest.LoadTestCase("iplistitemmultipleentries.tf", ID, name, comment, accountID)
}

func testAccCheckCloudflareListItemExists(n string, name string, listItem *cfv1.ListItem) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
		listRS := s.RootModule().Resources["cloudflare_list."+name]

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no List ID is set")
		}

		client, err := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if err != nil {
			return fmt.Errorf("error establishing client: %w", err)
		}
		foundList, err := client.GetListItem(context.Background(), cfv1.AccountIdentifier(accountID), listRS.Primary.ID, rs.Primary.ID)
		if err != nil {
			return err
		}

		*listItem = foundList

		return nil
	}
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
					resource.TestCheckResourceAttr(firstResource, "redirect.#", "1"),
					resource.TestCheckResourceAttr(firstResource, "redirect.0.source_url", "www.site1.com/"),
					resource.TestCheckResourceAttr(firstResource, "redirect.0.target_url", "https://example.com"),
					resource.TestCheckResourceAttr(firstResource, "redirect.0.status_code", "301"),

					resource.TestCheckResourceAttr(secondResource, "redirect.#", "1"),
					resource.TestCheckResourceAttr(secondResource, "redirect.0.source_url", "www.site1.com/test"),
					resource.TestCheckResourceAttr(secondResource, "redirect.0.target_url", "https://cloudflare.com"),
					resource.TestCheckResourceAttr(secondResource, "redirect.0.status_code", "301"),
				),
			},
		},
	})
}

func testAccCheckCloudflareHostnameRedirectWithOverlappingSourceURL(ID, name, comment, accountID string) string {
	return acctest.LoadTestCase("hostnameredirectwithoverlappingsourceurl.tf", ID, name, comment, accountID)
}
