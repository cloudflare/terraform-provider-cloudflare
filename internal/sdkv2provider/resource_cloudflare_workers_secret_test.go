package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	cleanhttp "github.com/hashicorp/go-cleanhttp"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const scriptContentForSecret = `addEventListener('fetch', event => {event.respondWith(new Response('test 1'))});`

var workerSecretTestScriptName string

func TestAccCloudflareWorkerSecret_Basic(t *testing.T) {
	t.Parallel()

	name := generateRandomResourceName()
	secretText := generateRandomResourceName()
	workerSecretTestScriptName = generateRandomResourceName()
	accountId := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareWorkerSecretDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerSecretWithWorkerScript(workerSecretTestScriptName, name, secretText, accountId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerSecretExists(workerSecretTestScriptName, name, accountId),
				),
			},
		},
	})

	deleteWorker()
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

		if err != nil {
			return err
		}

		if len(secretResponse.Result) > 0 {
			return fmt.Errorf("Worker secret with name %s still exists against Work Script %s", secretResponse.Result[0].Name, params.ScriptName)
		}
	}

	return nil
}

func testAccCheckCloudflareWorkerSecretWithWorkerScript(scriptName string, name string, secretText string, accountId string) string {
	err := createWorker()

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf(`
	resource "cloudflare_worker_secret" "%[2]s" {
		account_id  = "%[4]s"
		script_name = "%[1]s"
		name 		= "%[2]s"
		secret_text	= "%[3]s"
	}`, scriptName, name, secretText, accountId)
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

		return fmt.Errorf("Worker secret with name %s not found against Worker Script %s", name, scriptName)
	}
}

func createWorker() error {
	client, err := createCloudflareClient()
	accountId := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if err != nil {
		panic(err)
	}

	_, err = client.UploadWorker(context.Background(), cloudflare.AccountIdentifier(accountId), cloudflare.CreateWorkerParams{
		ScriptName: workerSecretTestScriptName,
		Script:     scriptContentForSecret,
	})

	if err != nil {
		return err
	}

	return nil
}

func deleteWorker() error {
	client, err := createCloudflareClient()
	accountId := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if err != nil {
		return err
	}

	err = client.DeleteWorker(context.Background(), cloudflare.AccountIdentifier(accountId), cloudflare.DeleteWorkerParams{
		ScriptName: workerSecretTestScriptName,
	})

	if err != nil {
		return err
	}

	return nil
}

func createCloudflareClient() (*cloudflare.API, error) {
	options := []cloudflare.Option{}

	httpClient := cleanhttp.DefaultClient()
	options = append(options, cloudflare.HTTPClient(httpClient))

	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")

	client, err := cloudflare.NewWithAPIToken(apiToken, options...)

	if err != nil {
		return nil, fmt.Errorf("Error creating Cloudflare client directly: %s ", err)
	}

	return client, nil
}
