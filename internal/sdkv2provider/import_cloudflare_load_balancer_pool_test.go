package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareLoadBalancerPool_Import(t *testing.T) {
	t.Parallel()
	var loadBalancerPool cloudflare.LoadBalancerPool
	rnd := generateRandomResourceName()
	name := "cloudflare_load_balancer_pool." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerPoolConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerPoolExists(name, &loadBalancerPool),
				),
			},
			{
				ResourceName:        name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerPoolExists(name, &loadBalancerPool),
				),
			},
		},
	})
}
