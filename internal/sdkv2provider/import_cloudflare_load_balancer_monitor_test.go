package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareLoadBalancerMonitor_Import(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()
	name := "cloudflare_load_balancer_monitor." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerMonitorConfigBasic(rnd, accountID),
			},
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}
