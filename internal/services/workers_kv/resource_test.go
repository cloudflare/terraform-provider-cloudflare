package workers_kv_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/kv"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareWorkersKV_Basic(t *testing.T) {
	t.Parallel()
	name := utils.GenerateRandomResourceName()
	key := utils.GenerateRandomResourceName()
	value := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_kv." + name
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCloudflareWorkersKVDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersKVWithAccount(name, key, value, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVExists(key),
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
	name := utils.GenerateRandomResourceName()
	key := utils.GenerateRandomResourceName()
	value := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_kv." + name
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCloudflareWorkersKVDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersKVWithAccount(name, key, value, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVExists(key),
					resource.TestCheckResourceAttr(
						resourceName, "value", value,
					),
				),
			},
			{
				Config: testAccCheckCloudflareWorkersKVWithAccount(name, key+"-updated", value, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVExists(key+"-updated"),
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
	name := utils.GenerateRandomResourceName()
	key := utils.GenerateRandomResourceName()
	value := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_workers_kv." + name

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCloudflareWorkersKVDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersKVWithAccount(name, key, value, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVExists(key),
					resource.TestCheckResourceAttr(resourceName, "value", value),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
				),
			},
		},
	})
}

func testAccCloudflareWorkersKVDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_workers_kv" {
			continue
		}

		namespaceID := rs.Primary.Attributes["namespace_id"]
		key := rs.Primary.Attributes["key"]

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]

		_, err := client.KV.Namespaces.Values.Get(context.Background(), namespaceID, key, kv.NamespaceValueGetParams{AccountID: cloudflare.F(accountID)})

		if err == nil {
			return fmt.Errorf("workers kv pair still exists")
		}
	}

	return nil
}

func testAccCheckCloudflareWorkersKV(rName, key, value, accountID string) string {
	return acctest.LoadTestCase("workerskv.tf", rName, key, value, accountID)
}

func testAccCheckCloudflareWorkersKVWithAccount(rName string, key string, value string, accountID string) string {
	return acctest.LoadTestCase("workerskvwithaccount.tf", rName, key, value, accountID)
}

func testAccCheckCloudflareWorkersKVExists(key string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acctest.SharedClient()

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "cloudflare_workers_kv" {
				continue
			}

			accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
			namespaceID := rs.Primary.Attributes["namespace_id"]
			value, err := client.KV.Namespaces.Values.Get(context.Background(), namespaceID, key, kv.NamespaceValueGetParams{AccountID: cloudflare.F(accountID)})
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
