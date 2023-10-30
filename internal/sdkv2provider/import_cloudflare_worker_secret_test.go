package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareWorkerSecret_Import(t *testing.T) {
	secretName := generateRandomResourceName()
	resourceName := "cloudflare_worker_secret." + secretName
	secretText := generateRandomResourceName()
	workerSecretTestScriptName = generateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareWorkerSecretDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerSecretWithWorkerScript(workerSecretTestScriptName, secretName, secretText, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerSecretExists(workerSecretTestScriptName, secretName, accountID),
				),
			},
			{
				ResourceName:      resourceName,
				ImportStateId:     fmt.Sprintf("%s/%s/%s", accountID, workerSecretTestScriptName, secretName),
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerSecretExists(workerSecretTestScriptName, secretName, accountID),
				),
			},
		},
	})
}
