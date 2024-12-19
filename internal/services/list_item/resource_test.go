package list_item_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareListItem_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.ParallelTest(t, resource.TestCase{
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

func TestAccCloudflareListItem_Update(t *testing.T) {
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
			{
				Config: testAccCheckCloudflareIPListItem(rnd, rnd, rnd+"-updated", accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "comment", rnd+"-updated"),
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

func testAccCheckCloudflareIPListItem(ID, name, comment, accountID string) string {
	return acctest.LoadTestCase("iplistitem.tf", ID, name, comment, accountID)
}

func testAccCheckCloudflareIPListItemMultipleEntries(ID, name, comment, accountID string) string {
	return acctest.LoadTestCase("iplistitemmultipleentries.tf", ID, name, comment, accountID)
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
