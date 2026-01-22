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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_workers_kv", &resource.Sweeper{
		Name: "cloudflare_workers_kv",
		F:    testSweepCloudflareWorkersKV,
	})
}

func testSweepCloudflareWorkersKV(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if accountID == "" {
		tflog.Info(ctx, "Skipping Workers KV sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	// List all KV namespaces
	namespaces, err := client.KV.Namespaces.List(ctx, kv.NamespaceListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch KV namespaces: %s", err))
		return fmt.Errorf("failed to fetch KV namespaces: %w", err)
	}

	if len(namespaces.Result) == 0 {
		tflog.Info(ctx, "No KV namespaces to sweep")
		return nil
	}

	for _, namespace := range namespaces.Result {
		// Use standard filtering helper to only sweep test namespaces
		if !utils.ShouldSweepResource(namespace.Title) {
			continue
		}

		// List keys in this namespace
		keys, err := client.KV.Namespaces.Keys.List(ctx, namespace.ID, kv.NamespaceKeyListParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to fetch KV keys for namespace %s (%s): %s", namespace.Title, namespace.ID, err))
			continue
		}

		// Delete all keys in the test namespace
		tflog.Info(ctx, fmt.Sprintf("Deleting %d keys from KV namespace: %s (account: %s)", len(keys.Result), namespace.Title, accountID))
		for _, key := range keys.Result {
			_, err := client.KV.Namespaces.Values.Delete(ctx, namespace.ID, key.Name, kv.NamespaceValueDeleteParams{
				AccountID: cloudflare.F(accountID),
			})
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete KV key %s in namespace %s: %s", key.Name, namespace.Title, err))
				continue
			}
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted keys from KV namespace: %s", namespace.Title))
	}

	return nil
}

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
					testAccCheckCloudflareWorkersKVExists(key + "-updated"),
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
				// Import immediately after creation (more reliable than after update due to KV eventual consistency).
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					namespaceResourceName := fmt.Sprintf("cloudflare_workers_kv_namespace.%s", name)
					return fmt.Sprintf("%s/%s/%s", accountID, s.RootModule().Resources[namespaceResourceName].Primary.ID, key), nil
				},
			},
			{
				// PlanOnly to avoid post-apply refresh issues from KV eventual consistency.
				Config:             testAccCheckCloudflareWorkersKV(name, key, updatedValue, accountID),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
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
	key := "test-key_with.special-chars/" + utils.GenerateRandomResourceName()
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
				// PlanOnly to avoid post-apply refresh issues from KV eventual consistency.
				Config:             testAccCheckCloudflareWorkersKVWithMetadata(name, key, value, accountID, updatedMetadata),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
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
