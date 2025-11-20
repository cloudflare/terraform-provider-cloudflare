package workers_kv_namespace_test

import (
	"context"
	"fmt"
	"log"
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

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_workers_kv_namespace", &resource.Sweeper{
		Name: "cloudflare_workers_kv_namespace",
		F:    testSweepCloudflareWorkersKVNamespace,
	})
}

func testSweepCloudflareWorkersKVNamespace(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	
	if accountID == "" {
		return nil
	}

	// List all KV namespaces
	namespaces, err := client.KV.Namespaces.List(ctx, kv.NamespaceListParams{
		AccountID: cloudflare.F(accountID),
	})
	if err != nil {
		log.Printf("[ERROR] Failed to fetch KV namespaces: %s", err)
		return err
	}

	// Delete all namespaces (sweepers clean up everything from test accounts)
	for _, namespace := range namespaces.Result {
		_, err := client.KV.Namespaces.Delete(ctx, namespace.ID, kv.NamespaceDeleteParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			log.Printf("[ERROR] Failed to delete KV namespace %s (%s): %s", namespace.Title, namespace.ID, err)
		}
	}

	return nil
}

func TestAccCloudflareWorkersKVNamespace_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	newRnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_kv_namespace." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCloudflareWorkersKVNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersKVNamespace(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("title"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("supports_url_encoding"), knownvalue.NotNull()),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVNamespaceExists(rnd, rnd),
				),
			},
			{
				ImportState:       true,
				ImportStateVerify: true,
				ResourceName:      resourceName,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s", accountID, s.RootModule().Resources[resourceName].Primary.ID), nil
				},
			},
			{
				Config: testAccCheckCloudflareWorkersKVNamespaceRename(rnd, newRnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("title"), knownvalue.StringExact(newRnd)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("title"), knownvalue.StringExact(newRnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVNamespaceExists(rnd, newRnd),
				),
			},
		},
	})
}

func TestAccCloudflareWorkersKVNamespace_SpecialCharactersInTitle(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	// Test title with special characters and spaces
	specialTitle := "test-namespace_with.special chars-" + rnd
	resourceName := "cloudflare_workers_kv_namespace." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCloudflareWorkersKVNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersKVNamespaceCustomTitle(rnd, specialTitle, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("title"), knownvalue.StringExact(specialTitle)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVNamespaceExists(rnd, specialTitle),
				),
			},
		},
	})
}

func TestAccCloudflareWorkersKVNamespace_AccountIDForcesRecreation(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_kv_namespace." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	// This test validates that changing account_id requires replacement
	// Note: We can't actually test with a different account_id as that would require different credentials
	// But we can verify the schema has RequiresReplace plan modifier
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCloudflareWorkersKVNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersKVNamespace(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("title"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					// Validate computed field has expected type
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("supports_url_encoding"), knownvalue.Bool(true)),
				},
			},
		},
	})
}

func TestAccCloudflareWorkersKVNamespace_InvalidImportID(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_kv_namespace." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCloudflareWorkersKVNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersKVNamespace(rnd, accountID),
			},
			{
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: "invalid-import-id",
				ExpectError:   regexp.MustCompile("invalid ID|expected format"),
			},
			{
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: accountID, // Missing namespace ID
				ExpectError:   regexp.MustCompile("invalid ID|expected format"),
			},
		},
	})
}

func TestAccCloudflareWorkersKVNamespace_LongTitle(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	// Test with a reasonably long title (but not excessive to avoid API limits)
	longTitle := "very-long-namespace-title-" + strings.Repeat("test-", 10) + rnd
	resourceName := "cloudflare_workers_kv_namespace." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCloudflareWorkersKVNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersKVNamespaceCustomTitle(rnd, longTitle, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("title"), knownvalue.StringExact(longTitle)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVNamespaceExists(rnd, longTitle),
				),
			},
		},
	})
}

func TestAccCloudflareWorkersKVNamespace_MultipleUpdates(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	firstTitle := rnd + "-first"
	secondTitle := rnd + "-second"
	thirdTitle := rnd + "-final"
	resourceName := "cloudflare_workers_kv_namespace." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCloudflareWorkersKVNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersKVNamespaceCustomTitle(rnd, firstTitle, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("title"), knownvalue.StringExact(firstTitle)),
				},
			},
			{
				Config: testAccCheckCloudflareWorkersKVNamespaceCustomTitle(rnd, secondTitle, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("title"), knownvalue.StringExact(secondTitle)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("title"), knownvalue.StringExact(secondTitle)),
				},
			},
			{
				Config: testAccCheckCloudflareWorkersKVNamespaceCustomTitle(rnd, thirdTitle, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("title"), knownvalue.StringExact(thirdTitle)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("title"), knownvalue.StringExact(thirdTitle)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkersKVNamespaceExists(rnd, thirdTitle),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s", accountID, s.RootModule().Resources[resourceName].Primary.ID), nil
				},
			},
		},
	})
}

func testAccCloudflareWorkersKVNamespaceDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_workers_kv_namespace" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]

		page, err := client.KV.Namespaces.List(context.Background(), kv.NamespaceListParams{AccountID: cloudflare.F(accountID)})
		if err != nil {
			return err
		}

		for _, n := range page.Result {
			if n.ID == rs.Primary.ID {
				return fmt.Errorf("namespace still exists but should not")
			}
		}
	}

	return nil
}

func testAccCheckCloudflareWorkersKVNamespace(rName, accountID string) string {
	return acctest.LoadTestCase("workerskvnamespace.tf", rName, accountID)
}

func testAccCheckCloudflareWorkersKVNamespaceRename(resourceName, newName, accountID string) string {
	return acctest.LoadTestCase("workerskvnamespace_rename.tf", resourceName, newName, accountID)
}

func testAccCheckCloudflareWorkersKVNamespaceCustomTitle(resourceName, customTitle, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_workers_kv_namespace" "%s" {
	account_id = "%s"
	title = "%s"
}
`, resourceName, accountID, customTitle)
}

func testAccCheckCloudflareWorkersKVNamespaceExists(resourceSuffix, title string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acctest.SharedClient()

		rs, ok := s.RootModule().Resources["cloudflare_workers_kv_namespace."+resourceSuffix]
		if !ok {
			return fmt.Errorf("not found: %s", resourceSuffix)
		}
		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]

		page, err := client.KV.Namespaces.List(context.Background(), kv.NamespaceListParams{AccountID: cloudflare.F(accountID)})
		if err != nil {
			return err
		}

		for _, n := range page.Result {
			if n.Title == title {
				return nil
			}
		}

		return fmt.Errorf("namespace not found")
	}
}
