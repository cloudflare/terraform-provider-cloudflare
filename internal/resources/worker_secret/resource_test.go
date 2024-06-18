package worker_secret_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stainless-sdks/cloudflare-terraform/internal/acctest"
	"github.com/stainless-sdks/cloudflare-terraform/internal/utils"
)

const scriptContentForSecret = `addEventListener('fetch', event => {event.respondWith(new Response('test 1'))});`

var workerSecretTestScriptName string

func TestAccCloudflareWorkerSecret_Basic(t *testing.T) {
	t.Parallel()

	name := utils.GenerateRandomResourceName()
	secretText := utils.GenerateRandomResourceName()
	workerSecretTestScriptName = utils.GenerateRandomResourceName()
	accountId := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWorkerSecretDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerSecretWithWorkerScript(workerSecretTestScriptName, name, secretText, accountId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerSecretExists(workerSecretTestScriptName, name, accountId),
				),
			},
			{
				Config:                  testAccCheckCloudflareWorkerSecretWithWorkerScript(workerSecretTestScriptName, name, secretText, accountId),
				ResourceName:            "cloudflare_worker_secret." + name,
				ImportStateId:           fmt.Sprintf("%s/%s/%s", accountID, workerSecretTestScriptName, name),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"secret_text"},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerSecretExists(workerSecretTestScriptName, name, accountID),
				),
			},
		},
	})
}

func testAccCheckCloudflareWorkerSecretDestroy(s *terraform.State) error {
	accountId := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_worker_secret" {
			continue
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		params := cloudflare.ListWorkersSecretsParams{
			ScriptName: rs.Primary.Attributes["script_name"],
		}

		secretResponse, err := client.ListWorkersSecrets(context.Background(), cloudflare.AccountIdentifier(accountId), params)

		if err == nil {
			return fmt.Errorf("worker secret %q still exists against %q", secretResponse.Result[0].Name, params.ScriptName)
		}
	}

	return nil
}

func testAccCheckCloudflareWorkerSecretWithWorkerScript(scriptName string, name string, secretText string, accountId string) string {
	return fmt.Sprintf(`
	resource "cloudflare_worker_script" "%[2]s" {
		account_id = "%[4]s"
		name       = "%[1]s"
		content    = "%[5]s"
	}

	resource "cloudflare_worker_secret" "%[2]s" {
		account_id  = "%[4]s"
		script_name = cloudflare_worker_script.%[2]s.name
		name 		= "%[2]s"
		secret_text	= "%[3]s"
	}`, scriptName, name, secretText, accountId, scriptContentForSecret)
}

func testAccCheckCloudflareWorkerSecretExists(scriptName string, name string, accountId string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cloudflare.API)
		params := cloudflare.ListWorkersSecretsParams{
			ScriptName: scriptName,
		}

		secretResponse, err := client.ListWorkersSecrets(context.Background(), cloudflare.AccountIdentifier(accountId), params)

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
