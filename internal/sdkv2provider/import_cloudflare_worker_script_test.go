package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareWorkerScript_Import(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Workers
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	var script cloudflare.WorkerScript
	rnd := generateRandomResourceName()
	name := "cloudflare_worker_script." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareWorkerScriptDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerScriptConfigMultiScriptInitial(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerScriptExists(name, &script, nil),
				),
			},
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerScriptExists(name, &script, nil),
				),
			},
		},
	})
}
