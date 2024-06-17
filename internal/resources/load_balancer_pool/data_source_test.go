package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareLoadBalancerPools(t *testing.T) {
	t.Parallel()

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_load_balancer_pools.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareLoadBalancerPoolsConfig(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "pools.#", "2"),
				),
			},
		},
	})
}

func testAccCloudflareLoadBalancerPoolsConfig(name, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "pool1" {
	account_id = "%[2]s"
	name = "pool1"
	origins {
		name = "example-1"
		address = "example.com"
		enabled = true
	}
}

resource "cloudflare_load_balancer_pool" "pool2" {
	account_id = "%[2]s"
	name = "pool2"
	origins {
		name = "example-2"
		address = "example.com"
		enabled = true
	}
}

data "cloudflare_load_balancer_pools" "%[1]s" {
	account_id = "%[2]s"

	depends_on = ["cloudflare_load_balancer_pool.pool1", "cloudflare_load_balancer_pool.pool2"]
}
`, name, accountID)
}
