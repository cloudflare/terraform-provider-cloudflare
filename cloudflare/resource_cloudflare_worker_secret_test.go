package cloudflare

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const scriptContentForSecret = `addEventListener('fetch', event => {event.respondWith(new Response('test 1'))});`

var workerSecretTestScriptName string

func TestAccCloudflareWorkerSecret_Basic(t *testing.T) {
	t.Parallel()

	name := generateRandomResourceName()
	secretText := generateRandomResourceName()
	workerSecretTestScriptName = generateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testPreCheckAccountCreateWorker(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareWorkerSecretDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerSecretWithWorkerScript(workerSecretTestScriptName, name, secretText),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerSecretExists(workerSecretTestScriptName, name, secretText),
				),
			},
		},
	})
}

func testPreCheckAccountCreateWorker(t *testing.T) {
	testAccPreCheckAccount(t)

	// Create the worker manually
	client := testAccProvider.Meta().(*cloudflare.API)

	scriptRequestParams := cloudflare.WorkerRequestParams{
		ScriptName: workerSecretTestScriptName,
	}

	client.UploadWorker(&scriptRequestParams, scriptContentForSecret)

}

func testAccCheckCloudflareWorkerSecretDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)
	var discoveredWorkerSecretName string

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_worker_secret`" {
			continue
		}

		params := getRequestParamsFromResource(rs)
		discoveredWorkerSecretName = params.ScriptName
		secretResponse, err := client.ListWorkersSecrets(context.Background(), params.ScriptName)

		// Cleanup the temporary non-terraform worker created in PreCheck step.
		scriptRequestParams := cloudflare.WorkerRequestParams{
			ScriptName: discoveredWorkerSecretName,
		}

		_, delErr := client.DeleteWorker(&scriptRequestParams)

		if err != nil || delErr != nil {
			return fmt.Errorf("Error deleting worker secret, and temp worker script %s %s", err, delErr)
		}

		if len(secretResponse.Result) > 0 {
			return fmt.Errorf("Worker secret with name %s still exists against Work Script %s", secretResponse.Result[0].Name, params.ScriptName)
		}
	}

	return nil
}

func testAccCheckCloudflareWorkerSecretWithWorkerScript(scriptName string, name string, secretText string) string {
	return fmt.Sprintf(`
	resource "cloudflare_worker_secret" "%[2]s" {
		script_name = "%[1]s"
		name 		= "%[2]s"
		secret_text	= "%[3]s"
	}`, scriptName, name, secretText)
}

func testAccCheckCloudflareWorkerSecretExists(scriptName string, name string, secretText string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cloudflare.API)
		secretResponse, err := client.ListWorkersSecrets(context.Background(), scriptName)
		if err != nil {
			return err
		}

		for _, secret := range secretResponse.Result {
			if secret.Name == name {
				return nil
			}
		}

		return fmt.Errorf("Worker secret with name %s not found against Worker Script %s", name, scriptName)
	}
}
