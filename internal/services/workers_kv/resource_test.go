package workers_kv_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/kv"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareWorkersKV_Basic(t *testing.T) {
	t.Parallel()
	name := utils.GenerateRandomResourceName()
	key := utils.GenerateRandomResourceName()
	value := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_kv." + name
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	var namespaceID string

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCloudflareWorkersKVDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersKV(name, key, value, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVExists(key),
					resource.TestCheckResourceAttr(
						resourceName, "value", value,
					),
					getNamespaceID(resourceName, &namespaceID),
				),
			},
			{
				ImportState:       true,
				ImportStateVerify: true,
				ResourceName:      resourceName,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					namespaceResourceName := fmt.Sprintf("cloudflare_workers_kv_namespace.%s", name)
					return fmt.Sprintf("%s/%s/%s", accountID, s.RootModule().Resources[namespaceResourceName].Primary.ID, key), nil
				},
			},
			// test refresh behavior
			{
				PreConfig: func() {
					client := acctest.SharedClient()
					result, err := client.KV.Namespaces.Values.Update(context.Background(), namespaceID, key, kv.NamespaceValueUpdateParams{AccountID: cloudflare.F(accountID), Value: cloudflare.String("foo")})
					if err != nil {
						t.Errorf("Error updating KV value out-of-band to test drift detection: %s", err)
					}
					if result == nil {
						t.Error("Could not update KV value out-of-band to test drift detection.")
					}
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
				RefreshPlanChecks: resource.RefreshPlanChecks{
					PostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
					},
				},
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
				Config: testAccCheckCloudflareWorkersKV(name, key, value, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVExists(key),
					resource.TestCheckResourceAttr(
						resourceName, "value", value,
					),
				),
			},
			{
				Config: testAccCheckCloudflareWorkersKV(name, key+"-updated", value, accountID),
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

func TestAccCloudflareWorkersKV_WithMetadata(t *testing.T) {
	t.Parallel()
	name := utils.GenerateRandomResourceName()
	key := utils.GenerateRandomResourceName()
	value := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_workers_kv." + name
	metadataKey := utils.GenerateRandomResourceName()
	metadataValue := utils.GenerateRandomResourceName()
	metadata := fmt.Sprintf("{\\\"%s\\\": \\\"%s\\\"}", metadataKey, metadataValue)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCloudflareWorkersKVDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersKVWithMetadata(name, key, value, accountID, metadata),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVMetadataExists(key),
					resource.TestCheckResourceAttr(resourceName, "value", value),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "metadata", strings.ReplaceAll(metadata, "\\\"", "\"")),
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

func testAccCheckCloudflareWorkersKVWithAccount(rName, key, value, accountID string) string {
	return acctest.LoadTestCase("workerskvwithaccount.tf", rName, key, value, accountID)
}

func testAccCheckCloudflareWorkersKVWithMetadata(rName, key, value, accountID, metadata string) string {
	return acctest.LoadTestCase("workerskvwithmetadata.tf", rName, key, value, accountID, metadata)
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

func testAccCheckCloudflareWorkersKVMetadataExists(key string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acctest.SharedClient()

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "cloudflare_workers_kv" {
				continue
			}

			accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
			namespaceID := rs.Primary.Attributes["namespace_id"]
			value, err := client.KV.Namespaces.Metadata.Get(context.Background(), namespaceID, key, kv.NamespaceMetadataGetParams{AccountID: cloudflare.F(accountID)})
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

func getNamespaceID(resourceName string, namespaceId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Resource Not found: %s", resourceName)
		}

		if namespaceId != nil {
			*namespaceId = rs.Primary.Attributes["namespace_id"]
		}

		return nil
	}
}
