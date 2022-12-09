package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
data "cloudflare_load_balancer_pools" "%[1]s" {
	account_id = "%[2]s"
}

%[3]s`, name, accountID, testPools)
}

var testPools = fmt.Sprintf(`resource "cloudflare_load_balancer_pool" "pool1" {
	account_id = "%[1]s"
	name = "pool1"
	origins {
	  name = "example-1"
	  address = "192.0.2.1"
	  enabled = false
	  header {
		header = "Host"
		values = ["example-1"]
	  }
	}
	origins {
	  name = "example-2"
	  address = "192.0.2.2"
	  header {
		header = "Host"
		values = ["example-2"]
	  }
	}
	latitude = 55
	longitude = -12
	description = "example load balancer pool"
	enabled = false
	minimum_origins = 1
	notification_email = "someone@example.com"
	load_shedding {
	  default_percent = 55
	  default_policy = "random"
	  session_percent = 12
	  session_policy = "hash"
	}
  }
}

resource "cloudflare_load_balancer_pool" "pool2" {
	account_id = "%[1]s"
	name = "pool2"
	origins {
	  name = "example-3"
	  address = "192.0.2.3"
	  enabled = false
	  header {
		header = "Host"
		values = ["example-3"]
	  }
	}
	origins {
	  name = "example-4"
	  address = "192.0.2.4"
	  header {
		header = "Host"
		values = ["example-4"]
	  }
	}
	latitude = 55
	longitude = -12
	description = "example load balancer pool"
	enabled = false
	minimum_origins = 1
	notification_email = "someone@example.com"
	load_shedding {
	  default_percent = 55
	  default_policy = "random"
	  session_percent = 12
	  session_policy = "hash"
	}
  }
}`, os.Getenv("CLOUDFLARE_ACCOUNT_ID"))
