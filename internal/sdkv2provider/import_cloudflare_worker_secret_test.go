package cloudflare

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccCloudflareWorkerSecret_Import(t *testing.T) {
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
