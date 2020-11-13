package cloudflare

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccCloudflareWorkerSecret_Import(t *testing.T) {

	client, err := createCloudflareClient()

	if err != nil {
		panic(err)
	}

	scriptRequestParams := cloudflare.WorkerRequestParams{
		ScriptName: workerSecretTestScriptName,
	}

	client.UploadWorker(&scriptRequestParams, scriptContentForSecret)

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

	client.DeleteWorker(&scriptRequestParams)
}
