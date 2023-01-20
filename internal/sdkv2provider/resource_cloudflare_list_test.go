package sdkv2provider

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
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var list cloudflare.List

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareList(rnd, rnd, rnd, accountID, "ip"),
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
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var list cloudflare.List
	var initialID string

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareList(rnd, rnd, rnd, accountID, "ip"),
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
				Config: testAccCheckCloudflareList(rnd, rnd, rnd+"-updated", accountID, "ip"),
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
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rndIP := generateRandomResourceName()
	rndRedirect := generateRandomResourceName()

	nameIP := fmt.Sprintf("cloudflare_list.%s", rndIP)
	nameRedirect := fmt.Sprintf("cloudflare_list.%s", rndRedirect)

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var list cloudflare.List
	var initialID string

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareList(rndIP, rndIP, rndIP, accountID, "ip"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(nameIP, &list),
					resource.TestCheckResourceAttr(
						nameIP, "name", rndIP),
				),
			},
			{
				PreConfig: func() {
					initialID = list.ID
				},
				Config: testAccCheckCloudflareListIPUpdate(rndIP, rndIP, rndIP, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(nameIP, &list),
					func(state *terraform.State) error {
						if initialID != list.ID {
							return fmt.Errorf("wanted update but List got recreated (id changed %q -> %q)", initialID, list.ID)
						}
						return nil
					},
					resource.TestCheckResourceAttr(nameIP, "item.#", "2"),
				),
			},
			{
				Config: testAccCheckCloudflareList(rndRedirect, rndRedirect, rndRedirect, accountID, "redirect"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(nameRedirect, &list),
					resource.TestCheckResourceAttr(
						nameRedirect, "name", rndRedirect),
				),
			},
			{
				PreConfig: func() {
					initialID = list.ID
				},
				Config: testAccCheckCloudflareListRedirectUpdate(rndRedirect, rndRedirect, rndRedirect, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(nameRedirect, &list),
					func(state *terraform.State) error {
						if initialID != list.ID {
							return fmt.Errorf("wanted update but List got recreated (id changed %q -> %q)", initialID, list.ID)
						}
						return nil
					},
					resource.TestCheckResourceAttr(nameRedirect, "item.#", "2"),
				),
			},
			{
				PreConfig: func() {
					initialID = list.ID
				},
				Config: testAccCheckCloudflareListRedirectUpdateTargetUrl(rndRedirect, rndRedirect, rndRedirect, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(nameRedirect, &list),
					func(state *terraform.State) error {
						if initialID != list.ID {
							return fmt.Errorf("wanted update but List got recreated (id changed %q -> %q)", initialID, list.ID)
						}
						return nil
					},
					resource.TestCheckResourceAttr(nameRedirect, "item.#", "1"),
				),
			},
		},
	})
}

func TestAccCloudflareList_UpdateIgnoreIPOrdering(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the IP List
	// endpoint does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

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
				Config: testAccCheckCloudflareListIPListOrdered(rnd, rnd, rnd, accountID),
			},
			{
				Config: testAccCheckCloudflareListIPListUnordered(rnd, rnd, rnd, accountID),
			},
		},
	})
}

func testAccCheckCloudflareListIPListOrdered(ID, name, description, accountID string) string {
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

	item {
	  value {
		ip = "192.0.2.2"
	  }
	  comment = "three"
	}

	item {
	  value {
	    ip = "192.0.2.3"
	  }
	  comment = "four"
	}
  }`, ID, name, description, accountID)
}

func testAccCheckCloudflareListIPListUnordered(ID, name, description, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_list" "%[1]s" {
    account_id = "%[4]s"
    name = "%[2]s"
    description = "%[3]s"
    kind = "ip"

    item {
	  value {
	    ip = "192.0.2.2"
	  }
	  comment = "three"
	}

	item {
      value {
        ip = "192.0.2.0"
      }
      comment = "one"
    }

	item {
	  value {
		ip = "192.0.2.3"
	  }
	  comment = "four"
	}

    item {
      value {
        ip = "192.0.2.1"
      }
      comment = "two"
    }
  }`, ID, name, description, accountID)
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
		foundList, err := client.GetList(context.Background(), cloudflare.AccountIdentifier(accountID), rs.Primary.ID)
		if err != nil {
			return err
		}

		*list = foundList

		return nil
	}
}

func testAccCheckCloudflareList(ID, name, description, accountID, kind string) string {
	return fmt.Sprintf(`
  resource "cloudflare_list" "%[1]s" {
    account_id = "%[4]s"
    name = "%[2]s"
    description = "%[3]s"
    kind = "%[5]s"
  }`, ID, name, description, accountID, kind)
}

func testAccCheckCloudflareListIPUpdate(ID, name, description, accountID string) string {
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

func testAccCheckCloudflareListRedirectUpdate(ID, name, description, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_list" "%[1]s" {
    account_id = "%[4]s"
    name = "%[2]s"
    description = "%[3]s"
    kind = "redirect"

    item {
      value {
        redirect {
          source_url = "cloudflare.com/blog"
          target_url = "https://blog.cloudflare.com"
        }
      }
      comment = "one"
    }

    item {
      value {
        redirect {
          source_url = "cloudflare.com/foo"
          target_url = "https://foo.cloudflare.com"
          include_subdomains = "enabled"
          subpath_matching = "enabled"
          status_code = 301
          preserve_query_string = "enabled"
          preserve_path_suffix = "disabled"
		}
      }
      comment = "two"
    }
  }`, ID, name, description, accountID)
}

func testAccCheckCloudflareListRedirectUpdateTargetUrl(ID, name, description, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_list" "%[1]s" {
    account_id = "%[4]s"
    name = "%[2]s"
    description = "%[3]s"
    kind = "redirect"

    item {
      value {
        redirect {
          source_url = "cloudflare.com/blog"
          target_url = "https://theblog.cloudflare.com"
        }
      }
    }
  }`, ID, name, description, accountID)
}
