package cloudflare

import (
	"context"
	"fmt"
	"os"
	"testing"

	cleanhttp "github.com/hashicorp/go-cleanhttp"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/version"
	"github.com/hashicorp/terraform-plugin-sdk/helper/logging"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/httpclient"
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
		PreCheck:     func() { testAccPreCheckAccount(t) },
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

	deleteWorker()
}

func testAccCheckCloudflareWorkerSecretDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_worker_secret`" {
			continue
		}

		params := getRequestParamsFromResource(rs)
		secretResponse, err := client.ListWorkersSecrets(context.Background(), params.ScriptName)

		if err != nil {
			return err
		}

		if len(secretResponse.Result) > 0 {
			return fmt.Errorf("Worker secret with name %s still exists against Work Script %s", secretResponse.Result[0].Name, params.ScriptName)
		}
	}

	return nil
}

func testAccCheckCloudflareWorkerSecretWithWorkerScript(scriptName string, name string, secretText string) string {
	err := createWorker()

	if err != nil {
		panic(err)
	}

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

func createWorker() error {
	client, err := createCloudflareClient()

	if err != nil {
		panic(err)
	}

	scriptRequestParams := cloudflare.WorkerRequestParams{
		ScriptName: workerSecretTestScriptName,
	}

	_, err = client.UploadWorker(&scriptRequestParams, scriptContentForSecret)

	if err != nil {
		return err
	}

	return nil
}

func deleteWorker() error {
	client, err := createCloudflareClient()

	if err != nil {
		return err
	}

	scriptRequestParams := cloudflare.WorkerRequestParams{
		ScriptName: workerSecretTestScriptName,
	}

	_, err = client.DeleteWorker(&scriptRequestParams)

	if err != nil {
		return err
	}

	return nil
}

func createCloudflareClient() (*cloudflare.API, error) {
	options := []cloudflare.Option{}

	httpClient := cleanhttp.DefaultClient()
	httpClient.Transport = logging.NewTransport("Cloudflare", httpClient.Transport)
	options = append(options, cloudflare.HTTPClient(httpClient))

	tfUserAgent := httpclient.TerraformUserAgent("0.11+compatible")
	providerUserAgent := fmt.Sprintf("terraform-provider-cloudflare/%s", version.ProviderVersion)
	ua := fmt.Sprintf("%s %s", tfUserAgent, providerUserAgent)

	options = append(options, cloudflare.UserAgent(ua))
	options = append(options, cloudflare.UsingAccount(os.Getenv("CLOUDFLARE_ACCOUNT_ID")))

	apiToken := os.Getenv("CLOUDFLARE_TOKEN")

	client, err := cloudflare.NewWithAPIToken(apiToken, options...)

	if err != nil {
		return nil, fmt.Errorf("Error creating Cloudflare client directly: %s ", err)
	}

	return client, nil
}
