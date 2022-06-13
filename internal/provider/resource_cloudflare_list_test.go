package provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudflareList_Exists(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the IP List
	// endpoint does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		defer func(apiToken string) {
			os.Setenv("CLOUDFLARE_API_TOKEN", apiToken)
		}(os.Getenv("CLOUDFLARE_API_TOKEN"))
		os.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var list cloudflare.List

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckAccount(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareList(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(name, &list),
					resource.TestCheckResourceAttr(
						name, "name", rnd),
				),
			},
		},
	})
}

func TestAccCloudflareList_UpdateDescription(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the IP List
	// endpoint does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		defer func(apiToken string) {
			os.Setenv("CLOUDFLARE_API_TOKEN", apiToken)
		}(os.Getenv("CLOUDFLARE_API_TOKEN"))
		os.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var list cloudflare.List
	var initialID string

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckAccount(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareList(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(name, &list),
					resource.TestCheckResourceAttr(
						name, "description", rnd),
				),
			},
			{
				PreConfig: func() {
					initialID = list.ID
				},
				Config: testAccCheckCloudflareList(rnd, rnd, rnd+"-updated", accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(name, &list),
					func(state *terraform.State) error {
						if initialID != list.ID {
							return fmt.Errorf("wanted update but List got recreated (id changed %q -> %q)", initialID, list.ID)
						}
						return nil
					},
					resource.TestCheckResourceAttr(name, "description", rnd+"-updated"),
				),
			},
		},
	})
}

func TestAccCloudflareList_Update(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the IP List
	// endpoint does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		defer func(apiToken string) {
			os.Setenv("CLOUDFLARE_API_TOKEN", apiToken)
		}(os.Getenv("CLOUDFLARE_API_TOKEN"))
		os.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var list cloudflare.List
	var initialID string

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckAccount(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareList(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(name, &list),
					resource.TestCheckResourceAttr(
						name, "name", rnd),
				),
			},
			{
				PreConfig: func() {
					initialID = list.ID
				},
				Config: testAccCheckCloudflareListUpdate(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(name, &list),
					func(state *terraform.State) error {
						if initialID != list.ID {
							return fmt.Errorf("wanted update but List got recreated (id changed %q -> %q)", initialID, list.ID)
						}
						return nil
					},
					resource.TestCheckResourceAttr(name, "item.#", "2"),
				),
			},
		},
	})
}

func testAccCheckCloudflareListExists(n string, list *cloudflare.List) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No List ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundList, err := client.GetList(context.Background(), cloudflare.ListGetParams{
			AccountID: accountID,
			ID:        rs.Primary.ID,
		})
		if err != nil {
			return err
		}

		*list = foundList

		return nil
	}
}

func testAccCheckCloudflareList(ID, name, description, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_list" "%[1]s" {
    account_id = "%[4]s"
    name = "%[2]s"
    description = "%[3]s"
    kind = "ip"
  }`, ID, name, description, accountID)
}

func testAccCheckCloudflareListUpdate(ID, name, description, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_list" "%[1]s" {
    account_id = "%[4]s"
    name = "%[2]s"
    description = "%[3]s"
    kind = "ip"

    item {
      value {
        ip = "192.0.2.0"
     }
      comment = "one"
    }

    item {
      value {
        ip = "192.0.2.1"
      }
      comment = "two"
    }
  }`, ID, name, description, accountID)
}
