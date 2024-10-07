package list_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"

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
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return errors.New("CLOUDFLARE_ACCOUNT_ID must be set")
	}

	lists, err := client.ListLists(ctx, cfv1.AccountIdentifier(accountID), cfv1.ListListsParams{})
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
		client.DeleteList(ctx, cfv1.AccountIdentifier(accountID), list.ID)
	}

	return nil
}

func TestAccCloudflareList_Exists(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var list cfv1.List

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var list cfv1.List
	var initialID string

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
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
	rndIP := utils.GenerateRandomResourceName()
	rndRedirect := utils.GenerateRandomResourceName()
	rndASN := utils.GenerateRandomResourceName()
	rndHostname := utils.GenerateRandomResourceName()

	nameIP := fmt.Sprintf("cloudflare_list.%s", rndIP)
	nameRedirect := fmt.Sprintf("cloudflare_list.%s", rndRedirect)
	nameASN := fmt.Sprintf("cloudflare_list.%s", rndASN)
	nameHostname := fmt.Sprintf("cloudflare_list.%s", rndHostname)

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var list cfv1.List
	var initialID string

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareList(rndIP, rndIP, rndIP, accountID, "ip"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(nameIP, &list),
					resource.TestCheckResourceAttr(
						nameIP, "name", rndIP),
					resource.TestCheckResourceAttr(
						nameIP, "description", rndIP),
					resource.TestCheckResourceAttr(
						nameIP, "kind", "ip"),
					resource.TestCheckResourceAttr(
						nameIP, "account_id", accountID),
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

					resource.TestCheckResourceAttr(nameIP, "item.0.value.0.redirect.#", "0"),
					resource.TestCheckResourceAttr(nameIP, "item.0.value.0.hostanme.#", "0"),
					resource.TestCheckNoResourceAttr(nameIP, "item.0.value.0.asn"),
					resource.TestCheckResourceAttr(nameIP, "item.1.value.0.redirect.#", "0"),
					resource.TestCheckResourceAttr(nameIP, "item.1.value.0.hostname.#", "0"),
					resource.TestCheckNoResourceAttr(nameIP, "item.1.value.0.asn"),

					resource.TestCheckTypeSetElemNestedAttrs(nameIP, "item.*", map[string]string{
						"value.0.ip": "192.0.2.0",
						"comment":    "one",
					}),
					resource.TestCheckTypeSetElemNestedAttrs(nameIP, "item.*", map[string]string{
						"value.0.ip": "192.0.2.1",
						"comment":    "two",
					}),
				),
			},
			{
				Config: testAccCheckCloudflareList(rndRedirect, rndRedirect, rndRedirect, accountID, "redirect"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(nameRedirect, &list),
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
					resource.TestCheckResourceAttr(nameRedirect, "kind", "redirect"),
					resource.TestCheckResourceAttr(nameRedirect, "item.#", "2"),

					resource.TestCheckResourceAttr(nameRedirect, "item.0.value.0.hostname.#", "0"),
					resource.TestCheckNoResourceAttr(nameRedirect, "item.0.value.0.asn"),
					resource.TestCheckResourceAttr(nameRedirect, "item.1.value.0.hostname.#", "0"),
					resource.TestCheckNoResourceAttr(nameRedirect, "item.1.value.0.asn"),
					resource.TestCheckNoResourceAttr(nameRedirect, "item.0.value.0.ip"),
					resource.TestCheckNoResourceAttr(nameRedirect, "item.1.value.0.ip"),

					resource.TestCheckTypeSetElemNestedAttrs(nameRedirect, "item.*", map[string]string{
						"value.0.redirect.0.source_url": "cloudflare.com/blog",
						"value.0.redirect.0.target_url": "https://blog.cloudflare.com",
					}),
					resource.TestCheckTypeSetElemNestedAttrs(nameRedirect, "item.*", map[string]string{
						"value.0.redirect.0.source_url":            "cloudflare.com/foo",
						"value.0.redirect.0.target_url":            "https://foo.cloudflare.com",
						"value.0.redirect.0.include_subdomains":    "enabled",
						"value.0.redirect.0.subpath_matching":      "enabled",
						"value.0.redirect.0.status_code":           "301",
						"value.0.redirect.0.preserve_query_string": "enabled",
						"value.0.redirect.0.preserve_path_suffix":  "disabled",
					}),
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
					resource.TestCheckResourceAttr(nameRedirect, "item.0.value.0.redirect.0.source_url", "cloudflare.com/blog"),
					resource.TestCheckResourceAttr(nameRedirect, "item.0.value.0.redirect.0.target_url", "https://theblog.cloudflare.com"),
				),
			},
			{
				Config: testAccCheckCloudflareList(rndASN, rndASN, rndASN, accountID, "asn"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(nameASN, &list),
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
					resource.TestCheckResourceAttr(nameASN, "kind", "asn"),
					resource.TestCheckResourceAttr(nameASN, "item.#", "2"),

					resource.TestCheckResourceAttr(nameASN, "item.0.value.0.redirect.#", "0"),
					resource.TestCheckResourceAttr(nameASN, "item.0.value.0.hostname.#", "0"),
					resource.TestCheckNoResourceAttr(nameASN, "item.0.value.0.ip"),
					resource.TestCheckResourceAttr(nameASN, "item.1.value.0.redirect.#", "0"),
					resource.TestCheckResourceAttr(nameASN, "item.1.value.0.hostname.#", "0"),
					resource.TestCheckNoResourceAttr(nameASN, "item.1.value.0.ip"),
					resource.TestCheckTypeSetElemNestedAttrs(nameASN, "item.*", map[string]string{
						"value.0.asn": "345",
					}),
					resource.TestCheckTypeSetElemNestedAttrs(nameASN, "item.*", map[string]string{
						"value.0.asn": "567",
					}),
				),
			},
			{
				Config: testAccCheckCloudflareList(rndHostname, rndHostname, rndHostname, accountID, "hostname"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(nameHostname, &list),
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
					resource.TestCheckTypeSetElemNestedAttrs(nameHostname, "item.*", map[string]string{
						"comment":                         "hostname one",
						"value.0.hostname.0.url_hostname": "*.google.com",
					}),
					resource.TestCheckTypeSetElemNestedAttrs(nameHostname, "item.*", map[string]string{
						"comment":                         "hostname two",
						"value.0.hostname.0.url_hostname": "manutd.com",
					}),
				),
			},
		},
	})
}

func TestAccCloudflareList_UpdateIgnoreIPOrdering(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareListIPListOrdered(rnd, rnd, rnd, accountID),
			},
			{
				Config:             testAccCheckCloudflareListIPListUnordered(rnd, rnd, rnd, accountID),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func TestAccCloudflareList_RemoveInlineConfig(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var list cfv1.List

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareList(rnd, rnd, rnd, accountID, "ip"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListExists(name, &list),
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

func testAccCheckCloudflareListExists(n string, list *cfv1.List) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No List ID is set")
		}

		client, _ := acctest.SharedV1Client()
		foundList, err := client.GetList(context.Background(), cfv1.AccountIdentifier(accountID), rs.Primary.ID)
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
