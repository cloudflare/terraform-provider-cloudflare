package workers_for_platforms_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_workers_for_platforms_namespace", &resource.Sweeper{
		Name: "cloudflare_workers_for_platforms_namespace",
		F: func(region string) error {
			client, err := acctest.SharedV1Client()
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

			if err != nil {
				return fmt.Errorf("error establishing client: %w", err)
			}

			ctx := context.Background()
			resp, err := client.ListWorkersForPlatformsDispatchNamespaces(ctx, cfv1.AccountIdentifier(accountID))
			if err != nil {
				return err
			}

			for _, namespace := range resp.Result {
				err := client.DeleteWorkersForPlatformsDispatchNamespace(ctx, cfv1.AccountIdentifier(accountID), namespace.NamespaceName)
				if err != nil {
					return err
				}
			}

			return nil
		},
	})
}

func TestAccCloudflareWorkersForPlatforms_NamespaceManagement(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_workers_for_platforms_namespace." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersForPlatformsNamespaceManagement(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:        resourceName,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func testAccCheckCloudflareWorkersForPlatformsNamespaceManagement(rnd, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_workers_for_platforms_namespace" "%[1]s" {
    account_id = "%[2]s"
    name       = "%[1]s"
  }`, rnd, accountID)
}

func TestAccCloudflareWorkersForPlatforms_UploadUserWorker(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_workers_for_platforms_namespace." + rnd
	workerResource := "cloudflare_worker_script.script_" + rnd

	scriptContent := `<<EOT
	export default {
		fetch() {
			return new Response("Hello, World!")
		}
	}
	EOT
	`

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWorkerScriptDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersForPlatformsUploadUserWorker(rnd, accountID, scriptContent, "2024-01-01", []string{"free"}),
				Check: resource.ComposeTestCheckFunc(
					// Check namespace
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttrSet(resourceName, "id"),

					// Check Worker
					resource.TestCheckResourceAttr(workerResource, "name", "script_"+rnd),
					resource.TestCheckResourceAttr(workerResource, "compatibility_date", "2024-01-01"),
					resource.TestCheckResourceAttr(workerResource, "dispatch_namespace", rnd),
					resource.TestCheckResourceAttr(workerResource, "tags.#", "1"),
					resource.TestCheckResourceAttr(workerResource, "tags.0", "free"),
				),
			},
			{
				ResourceName:        resourceName,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func testAccCheckCloudflareWorkersForPlatformsUploadUserWorker(rnd, accountID, moduleContent, compatibilityDate string, tags []string) string {
	return fmt.Sprintf(`
	  resource "cloudflare_workers_for_platforms_namespace" "%[1]s" {
		account_id = "%[2]s"
		name       = "%[1]s"
	  }

	  resource "cloudflare_worker_script" "script_%[1]s" {
		account_id          = "%[2]s"
		name                = "script_%[1]s"
		content             = %[3]s
		module              = true
		compatibility_date  = "%[4]s"
		dispatch_namespace  = "%[1]s"
		tags                = %[5]q

		depends_on = [cloudflare_workers_for_platforms_namespace.%[1]s]
	  }`, rnd, accountID, moduleContent, compatibilityDate, tags)
}

func testAccCheckCloudflareWorkerScriptDestroy(s *terraform.State) error {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_worker_script" {
			continue
		}

		client, err := acctest.SharedV1Client()
		if err != nil {
			return fmt.Errorf("error establishing client: %w", err)
		}
		r, _ := client.GetWorkerWithDispatchNamespace(
			context.Background(),
			cfv1.AccountIdentifier(accountID),
			rs.Primary.Attributes["name"],
			rs.Primary.Attributes["dispatch_namespace"],
		)

		if r.Script != "" {
			return fmt.Errorf("namespaced worker script with id %s still exists", rs.Primary.ID)
		}
	}

	return nil
}
