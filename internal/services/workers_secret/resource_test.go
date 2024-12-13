package workers_secret_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/workers_for_platforms"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const scriptContentForSecret = `addEventListener('fetch', event => {event.respondWith(new Response('test 1'))});`

var workerSecretTestScriptName string

func TestAccCloudflareWorkerSecret_Basic(t *testing.T) {
	t.Parallel()

	name := utils.GenerateRandomResourceName()
	secretText := utils.GenerateRandomResourceName()
	workerSecretTestScriptName = utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWorkerSecretDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerSecretWithWorkerScript(workerSecretTestScriptName, name, secretText, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerSecretExists(workerSecretTestScriptName, name, accountID, name),
				),
			},
			// {
			// 	Config:                  testAccCheckCloudflareWorkerSecretWithWorkerScript(workerSecretTestScriptName, name, secretText, accountID),
			// 	ResourceName:            "cloudflare_workers_secret." + name,
			// 	ImportStateId:           fmt.Sprintf("%s/%s/%s", accountID, workerSecretTestScriptName, name),
			// 	ImportState:             true,
			// 	ImportStateVerify:       true,
			// 	ImportStateVerifyIgnore: []string{"secret_text"},
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testAccCheckCloudflareWorkerSecretExists(workerSecretTestScriptName, name, accountID),
			// 	),
			// },
		},
	})
}

func testAccCheckCloudflareWorkerSecretDestroy(s *terraform.State) error {
	accountId := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_workers_secret" {
			continue
		}
		client := acctest.SharedClient()

		namespace := rs.Primary.Attributes["dispatch_namespace"]
		scriptName := rs.Primary.Attributes["script_name"]
		secretResponse, err := client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Secrets.List(
			context.Background(),
			namespace,
			scriptName,
			workers_for_platforms.DispatchNamespaceScriptSecretListParams{AccountID: cloudflare.F(accountId)},
		)
		if err == nil {
			return fmt.Errorf("worker secret %q still exists against %q", secretResponse.Result[0].Name, scriptName)
		}
	}

	return nil
}

func testAccCheckCloudflareWorkerSecretWithWorkerScript(scriptName string, name string, secretText string, accountId string) string {
	return acctest.LoadTestCase("workersecretwithworkerscript.tf", scriptName, name, secretText, accountId, scriptContentForSecret)
}

func testAccCheckCloudflareWorkerSecretExists(scriptName, name, accountId, namespace string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acctest.SharedClient()
		secretResponse, err := client.WorkersForPlatforms.Dispatch.Namespaces.Scripts.Secrets.List(
			context.Background(),
			namespace,
			scriptName,
			workers_for_platforms.DispatchNamespaceScriptSecretListParams{AccountID: cloudflare.F(accountId)},
		)
		if err != nil {
			return err
		}

		for _, secret := range secretResponse.Result {
			if secret.Name == name {
				return nil
			}
		}

		return fmt.Errorf("worker secret with name %q not found against %s", name, scriptName)
	}
}
