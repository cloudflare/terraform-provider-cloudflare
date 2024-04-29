package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func init() {
	resource.AddTestSweepers("cloudflare_list", &resource.Sweeper{
		Name: "cloudflare_list",
		F:    testSweepCloudflareList,
	})
}

func testSweepCloudflareList(r string) error {
	ctx := context.Background()
	client, clientErr := sharedClient()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return errors.New("CLOUDFLARE_ACCOUNT_ID must be set")
	}

	lists, err := client.ListLists(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListListsParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Lists: %s", err))
	}

	if len(lists) == 0 {
		log.Print("[DEBUG] No Cloudflare Lists to sweep")
		return nil
	}

	for _, list := range lists {
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare List ID: %s", list.ID))
		//nolint:errcheck
		client.DeleteList(ctx, cloudflare.AccountIdentifier(accountID), list.ID)
	}

	return nil
}

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
	rndASN := generateRandomResourceName()
	rndHostname := generateRandomResourceName()

	nameIP := fmt.Sprintf("cloudflare_list.%s", rndIP)
	nameRedirect := fmt.Sprintf("cloudflare_list.%s", rndRedirect)
	nameASN := fmt.Sprintf("cloudflare_list.%s", rndASN)
	nameHostname := fmt.Sprintf("cloudflare_list.%s", rndHostname)

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
			{
				Config: testAccCheckCloudflareList(rndASN, rndASN, rndASN, accountID, "asn"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(nameASN, &list),
					resource.TestCheckResourceAttr(
						nameASN, "name", rndASN),
				),
			},
			{
				PreConfig: func() {
					initialID = list.ID
				},
				Config: testAccCheckCloudflareListASNUpdate(rndASN, rndASN, rndASN, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(nameASN, &list),
					func(state *terraform.State) error {
						if initialID != list.ID {
							return fmt.Errorf("wanted update but List got recreated (id changed %q -> %q)", initialID, list.ID)
						}
						return nil
					},
					resource.TestCheckResourceAttr(nameASN, "item.#", "2"),
				),
			},
			{
				Config: testAccCheckCloudflareList(rndHostname, rndHostname, rndHostname, accountID, "hostname"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(nameHostname, &list),
					resource.TestCheckResourceAttr(
						nameHostname, "name", rndHostname),
				),
			},
			{
				PreConfig: func() {
					initialID = list.ID
				},
				Config: testAccCheckCloudflareListHostnameUpdate(rndHostname, rndHostname, rndHostname, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(nameHostname, &list),
					func(state *terraform.State) error {
						if initialID != list.ID {
							return fmt.Errorf("wanted update but List got recreated (id changed %q -> %q)", initialID, list.ID)
						}
						return nil
					},
					resource.TestCheckResourceAttr(nameHostname, "item.#", "2"),
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

func TestAccCloudflareList_RemoveInlineConfig(t *testing.T) {
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
					resource.TestCheckResourceAttr(name, "item.#", "0"),
				),
			},
			{
				Config: testAccCheckCloudflareListBasicIP(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(name, &list),
					resource.TestCheckResourceAttr(name, "item.#", "1"),
				),
			},
			{
				Config: testAccCheckCloudflareList(rnd, rnd, rnd, accountID, "ip"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(name, &list),
					resource.TestCheckResourceAttr(name, "item.#", "0"),
				),
			},
		},
	})
}

func TestAccCloudflareList_Import(t *testing.T) {
	t.Skip("Pending investigation into item.0.value.0.asn being imported incorrectly")

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
				Config: testAccCheckCloudflareListBasicIP(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(name, &list),
					resource.TestCheckResourceAttr(name, "item.#", "1"),
				),
			},
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "item[0].value[0].ip"),
				),
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

func testAccCheckCloudflareListBasicIP(ID, name, description, accountID string) string {
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
  }`, ID, name, description, accountID)
}

func testAccCheckCloudflareListASNUpdate(ID, name, description, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_list" "%[1]s" {
    account_id = "%[4]s"
    name = "%[2]s"
    description = "%[3]s"
    kind = "asn"

    item {
      value {
        asn = 345
      }
      comment = "ASN test"
    }

    item {
      value {
        asn = 567
      }
      comment = "ASN test two"
    }
  }`, ID, name, description, accountID)
}

func testAccCheckCloudflareListHostnameUpdate(ID, name, description, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_list" "%[1]s" {
    account_id = "%[4]s"
    name = "%[2]s"
    description = "%[3]s"
    kind = "hostname"

    item {
      value {
        hostname {
		  url_hostname = "*.google.com"
		}
      }
      comment = "hostname one"
    }

    item {
	  value {
		hostname {
		  url_hostname = "manutd.com"
		}
	  }
      comment = "hostname two"
    }
  }`, ID, name, description, accountID)
}
