package sdkv2provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareWorkerSecret_Import(t *testing.T) {
	name := generateRandomResourceName()
	secretText := generateRandomResourceName()
	workerSecretTestScriptName = generateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheckAccount(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareWorkerSecretDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerSecretWithWorkerScript(workerSecretTestScriptName, name, secretText, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerSecretExists(workerSecretTestScriptName, name, accountID),
				),
			},
		},
	})

	deleteWorker()
}
