package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareListItem_Exists(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var ListItem cloudflare.ListItem

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
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

func TestAccCloudflareListItem_Update(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var listItem cloudflare.ListItem

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
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
	rnd := generateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckCloudflareBadListItemType(rnd, rnd, rnd, accountID),
				ExpectError: regexp.MustCompile(" can not be added to lists of type "),
			},
		},
	})
}

func testAccCheckCloudflareIPListItem(ID, name, comment, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_list" "%[2]s" {
    account_id          = "%[4]s"
    name                = "%[2]s"
    description         = "list named %[2]s"
    kind                = "ip"
  }
  
  resource "cloudflare_list_item" "%[1]s" {
    account_id = "%[4]s"
	list_id    = cloudflare_list.%[2]s.id
	ip         = "192.0.2.0"
	comment    = "%[3]s"
  } `, ID, name, comment, accountID)
}

func testAccCheckCloudflareListItemExists(n string, name string, listItem *cloudflare.ListItem) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
		listRS := s.RootModule().Resources["cloudflare_list."+name]

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No List ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundList, err := client.GetListItem(context.Background(), cloudflare.AccountIdentifier(accountID), listRS.Primary.ID, rs.Primary.ID)
		if err != nil {
			return err
		}

		*listItem = foundList

		return nil
	}
}

func testAccCheckCloudflareBadListItemType(ID, name, comment, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_list" "%[2]s" {
    account_id          = "%[4]s"
    name                = "%[2]s"
    description         = "list named %[2]s"
    kind                = "redirect"
  }
  
  resource "cloudflare_list_item" "%[2]s" {
    account_id = "%[4]s"
	list_id    = cloudflare_list.%[2]s.id
	ip         = "192.0.2.0"
	comment    = "%[3]s"
  } `, ID, name, comment, accountID)
}
