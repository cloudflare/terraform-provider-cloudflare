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

func TestAccCloudflareWorkersKV_Basic(t *testing.T) {
	t.Parallel()
	var kvPair cloudflare.WorkersKVPair
	name := generateRandomResourceName()
	key := generateRandomResourceName()
	value := generateRandomResourceName()
	resourceName := "cloudflare_workers_kv." + name

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCloudflareWorkersKVDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersKV(name, key, value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVExists(key, &kvPair),
					resource.TestCheckResourceAttr(
						resourceName, "value", value,
					),
				),
			},
		},
	})
}

func TestAccCloudflareWorkersKV_NameForcesRecreation(t *testing.T) {
	t.Parallel()
	var kvPair cloudflare.WorkersKVPair
	name := generateRandomResourceName()
	key := generateRandomResourceName()
	value := generateRandomResourceName()
	resourceName := "cloudflare_workers_kv." + name

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCloudflareWorkersKVDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersKV(name, key, value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVExists(key, &kvPair),
					resource.TestCheckResourceAttr(
						resourceName, "value", value,
					),
				),
			},
			{
				Config: testAccCheckCloudflareWorkersKV(name, key+"-updated", value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVExists(key, &kvPair),
					resource.TestCheckResourceAttr(
						resourceName, "value", value,
					),
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func TestAccCloudflareWorkersKV_WithAccountID(t *testing.T) {
	t.Parallel()
	var kvPair cloudflare.WorkersKVPair
	name := generateRandomResourceName()
	key := generateRandomResourceName()
	value := generateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_workers_kv." + name

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCloudflareWorkersKVDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersKVWithAccount(name, key, value, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVExists(key, &kvPair),
					resource.TestCheckResourceAttr(resourceName, "value", value),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
				),
			},
		},
	})
}

func testAccCloudflareWorkersKVDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_workers_kv" {
			continue
		}

		namespaceID := rs.Primary.Attributes["namespace_id"]
		key := rs.Primary.Attributes["key"]

		accountID := rs.Primary.Attributes["account_id"]
		if accountID == "" {
			accountID = client.AccountID
		}

		_, err := client.GetWorkersKV(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.GetWorkersKVParams{NamespaceID: namespaceID, Key: key})

		if err == nil {
			return fmt.Errorf("workers kv pair still exists")
		}
	}

	return nil
}

func testAccCheckCloudflareWorkersKV(rName string, key string, value string) string {
	return testAccCheckCloudflareWorkersKVNamespace(rName) + fmt.Sprintf(`
resource "cloudflare_workers_kv" "%[1]s" {
	namespace_id = cloudflare_workers_kv_namespace.%[1]s.id
	key = "%[2]s"
	value = "%[3]s"
}`, rName, key, value)
}

func testAccCheckCloudflareWorkersKVWithAccount(rName string, key string, value string, accountID string) string {
	return testAccCheckCloudflareWorkersKVNamespace(rName) + fmt.Sprintf(`
resource "cloudflare_workers_kv" "%[1]s" {
	account_id = "%[4]s"
	namespace_id = cloudflare_workers_kv_namespace.%[1]s.id
	key = "%[2]s"
	value = "%[3]s"
}`, rName, key, value, accountID)
}

func testAccCheckCloudflareWorkersKVExists(key string, kv *cloudflare.WorkersKVPair) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cloudflare.API)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "cloudflare_workers_kv" {
				continue
			}

			accountID := rs.Primary.Attributes["account_id"]
			if accountID == "" {
				accountID = client.AccountID
			}

			namespaceID := rs.Primary.Attributes["namespace_id"]
			value, err := client.GetWorkersKV(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.GetWorkersKVParams{NamespaceID: namespaceID, Key: key})
			if err != nil {
				return err
			}

			if value == nil {
				return fmt.Errorf("workers kv key %s not found in namespace %s", key, namespaceID)
			}
		}

		return nil
	}
}
