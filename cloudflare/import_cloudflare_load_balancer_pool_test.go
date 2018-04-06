package cloudflare

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCloudFlareLoadBalancerPool_Import(t *testing.T) {
	t.Parallel()
	var loadBalancerPool cloudflare.LoadBalancerPool
	rnd := acctest.RandString(10)
	name := "cloudflare_load_balancer_pool." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareLoadBalancerPoolConfigBasic(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareLoadBalancerPoolExists(name, &loadBalancerPool),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareLoadBalancerPoolExists(name, &loadBalancerPool),
				),
			},
		},
	})
}
