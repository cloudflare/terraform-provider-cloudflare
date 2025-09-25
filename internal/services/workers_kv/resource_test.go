package workers_kv_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/kv"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareWorkersKV_Basic(t *testing.T) {
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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("key_name"), knownvalue.StringExact(key)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact(value)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("namespace_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(key)),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVExists(key),
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
		},
	})
}

func TestAccCloudflareWorkersKV_NameForcesRecreation(t *testing.T) {
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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("key_name"), knownvalue.StringExact(key)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact(value)),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVExists(key),
				),
			},
			{
				Config: testAccCheckCloudflareWorkersKV(name, key+"-updated", value, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("key_name"), knownvalue.StringExact(key+"-updated")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("key_name"), knownvalue.StringExact(key+"-updated")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact(value)),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVExists(key+"-updated"),
				),
			},
		},
	})
}

func TestAccCloudflareWorkersKV_ValueUpdate(t *testing.T) {
	name := utils.GenerateRandomResourceName()
	key := utils.GenerateRandomResourceName()
	value := utils.GenerateRandomResourceName()
	updatedValue := utils.GenerateRandomResourceName()
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
				Config: testAccCheckCloudflareWorkersKV(name, key, value, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact(value)),
				},
			},
			{
				Config: testAccCheckCloudflareWorkersKV(name, key, updatedValue, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact(updatedValue)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact(updatedValue)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("key_name"), knownvalue.StringExact(key)),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVExists(key),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					namespaceResourceName := fmt.Sprintf("cloudflare_workers_kv_namespace.%s", name)
					return fmt.Sprintf("%s/%s/%s", accountID, s.RootModule().Resources[namespaceResourceName].Primary.ID, key), nil
				},
			},
		},
	})
}

func TestAccCloudflareWorkersKV_EmptyValue(t *testing.T) {
	name := utils.GenerateRandomResourceName()
	key := utils.GenerateRandomResourceName()
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
				Config: testAccCheckCloudflareWorkersKV(name, key, "", accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("key_name"), knownvalue.StringExact(key)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("namespace_id"), knownvalue.NotNull()),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVExists(key),
				),
			},
		},
	})
}

func TestAccCloudflareWorkersKV_LargeValue(t *testing.T) {
	name := utils.GenerateRandomResourceName()
	key := utils.GenerateRandomResourceName()
	// Create a 1MB value (well within the 25MB limit)
	largeValue := strings.Repeat("a", 1024*1024)
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
				Config: testAccCheckCloudflareWorkersKV(name, key, largeValue, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("key_name"), knownvalue.StringExact(key)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact(largeValue)),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVExists(key),
				),
			},
		},
	})
}

func TestAccCloudflareWorkersKV_SpecialCharactersInKey(t *testing.T) {
	name := utils.GenerateRandomResourceName()
	// Test key with simple special characters (avoid URL encoding issues)
	key := "test-key_with.special-chars." + utils.GenerateRandomResourceName()
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
				Config: testAccCheckCloudflareWorkersKV(name, key, value, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("key_name"), knownvalue.StringExact(key)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact(value)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(key)),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVExists(key),
				),
			},
		},
	})
}

func TestAccCloudflareWorkersKV_InvalidJSONMetadata(t *testing.T) {
	name := utils.GenerateRandomResourceName()
	key := utils.GenerateRandomResourceName()
	value := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	// Invalid JSON - missing closing brace
	invalidMetadata := `{\"key\": \"value\"`

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCloudflareWorkersKVDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckCloudflareWorkersKVWithMetadata(name, key, value, accountID, invalidMetadata),
				ExpectError: regexp.MustCompile("Invalid JSON String Value|not valid JSON string format"),
			},
		},
	})
}

func TestAccCloudflareWorkersKV_InvalidImportID(t *testing.T) {
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
				Config: testAccCheckCloudflareWorkersKV(name, key, value, accountID),
			},
			{
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: "invalid-import-id",
				ExpectError:   regexp.MustCompile("invalid ID|expected urlencoded segments"),
			},
			{
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s/namespace_id", accountID), // Missing key
				ExpectError:   regexp.MustCompile("invalid ID|expected urlencoded segments"),
			},
		},
	})
}

func TestAccCloudflareWorkersKV_MetadataUpdate(t *testing.T) {
	name := utils.GenerateRandomResourceName()
	key := utils.GenerateRandomResourceName()
	value := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_workers_kv." + name

	initialMetadata := `{\"version\": \"1.0\", \"env\": \"test\"}`
	updatedMetadata := `{\"version\": \"1.1\", \"env\": \"production\", \"tags\": [\"api\", \"v1\"]}`

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCloudflareWorkersKVDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersKVWithMetadata(name, key, value, accountID, initialMetadata),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("metadata"), knownvalue.StringExact(strings.ReplaceAll(initialMetadata, "\\\"", "\""))),
				},
			},
			{
				Config: testAccCheckCloudflareWorkersKVWithMetadata(name, key, value, accountID, updatedMetadata),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("metadata"), knownvalue.StringExact(strings.ReplaceAll(updatedMetadata, "\\\"", "\""))),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("metadata"), knownvalue.StringExact(strings.ReplaceAll(updatedMetadata, "\\\"", "\""))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("key_name"), knownvalue.StringExact(key)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact(value)),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVMetadataExists(key),
				),
			},
		},
	})
}

func TestAccCloudflareWorkersKV_WithMetadata(t *testing.T) {
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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("key_name"), knownvalue.StringExact(key)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact(value)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("metadata"), knownvalue.StringExact(strings.ReplaceAll(metadata, "\\\"", "\""))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("namespace_id"), knownvalue.NotNull()),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVMetadataExists(key),
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
