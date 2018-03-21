package cloudflare

import (
	"fmt"
	"testing"

	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccCloudFlareLoadBalancerPool_Basic(t *testing.T) {
	// multiple instances of this config would conflict but we only use it once
	t.Parallel()
	testStartTime := time.Now().UTC()
	var loadBalancerPool cloudflare.LoadBalancerPool
	rnd := acctest.RandString(10)
	name := "cloudflare_load_balancer_pool." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareLoadBalancerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareLoadBalancerPoolConfigBasic(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareLoadBalancerPoolExists(name, &loadBalancerPool),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr(name, "check_regions.#", "0"),
					// also expect api to generate some values
					testAccCheckCloudFlareLoadBalancerPoolDates(name, &loadBalancerPool, testStartTime),
				),
			},
		},
	})
}

func TestAccCloudFlareLoadBalancerPool_FullySpecified(t *testing.T) {
	t.Parallel()
	var loadBalancerPool cloudflare.LoadBalancerPool
	rnd := acctest.RandString(10)
	name := "cloudflare_load_balancer_pool." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareLoadBalancerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareLoadBalancerPoolConfigFullySpecified(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareLoadBalancerPoolExists(name, &loadBalancerPool),
					// checking our overrides of default values worked
					resource.TestCheckResourceAttr(name, "enabled", "false"),
					resource.TestCheckResourceAttr(name, "description", "tfacc-fully-specified"),
					resource.TestCheckResourceAttr(name, "check_regions.#", "1"),
					resource.TestCheckResourceAttr(name, "minimum_origins", "2"),
				),
			},
		},
	})
}

/**
Any change to a load balancer pool results in a new resource
Although the API client contains a modify method, this always results in 405 status
*/
func TestAccCloudFlareLoadBalancerPool_ForceNew(t *testing.T) {
	t.Parallel()
	var loadBalancerPool cloudflare.LoadBalancerPool
	var initialId string
	rnd := acctest.RandString(10)
	name := "cloudflare_load_balancer_pool." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareLoadBalancerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareLoadBalancerPoolConfigBasic(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareLoadBalancerPoolExists(name, &loadBalancerPool),
				),
			},
			{
				PreConfig: func() {
					initialId = loadBalancerPool.ID
				},
				Config: testAccCheckCloudFlareLoadBalancerPoolConfigFullySpecified(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareLoadBalancerPoolExists(name, &loadBalancerPool),
					func(state *terraform.State) error {
						if initialId == loadBalancerPool.ID {
							return fmt.Errorf("id should be different after recreation, but is unchanged: %s ",
								loadBalancerPool.ID)
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccCloudFlareLoadBalancerPool_CreateAfterManualDestroy(t *testing.T) {
	t.Parallel()
	var loadBalancerPool cloudflare.LoadBalancerPool
	var initialId string
	rnd := acctest.RandString(10)
	name := "cloudflare_load_balancer_pool." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareLoadBalancerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareLoadBalancerPoolConfigBasic(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareLoadBalancerPoolExists(name, &loadBalancerPool),
					testAccManuallyDeleteLoadBalancerPool(name, &loadBalancerPool, &initialId),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCloudFlareLoadBalancerPoolConfigBasic(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareLoadBalancerPoolExists(name, &loadBalancerPool),
					func(state *terraform.State) error {
						if initialId == loadBalancerPool.ID {
							return fmt.Errorf("load balancer pool id is unchanged even after we thought we deleted it ( %s )",
								loadBalancerPool.ID)
						}
						return nil
					},
				),
			},
		},
	})
}

func testAccCheckCloudFlareLoadBalancerPoolDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_load_balancer_pool" {
			continue
		}

		_, err := client.LoadBalancerPoolDetails(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Load balancer pool still exists")
		}
	}

	return nil
}

func testAccCheckCloudFlareLoadBalancerPoolExists(n string, loadBalancerPool *cloudflare.LoadBalancerPool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Load Balancer ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundLoadBalancerPool, err := client.LoadBalancerPoolDetails(rs.Primary.ID)
		if err != nil {
			return err
		}

		*loadBalancerPool = foundLoadBalancerPool

		return nil
	}
}

func testAccCheckCloudFlareLoadBalancerPoolDates(n string, loadBalancerPool *cloudflare.LoadBalancerPool, testStartTime time.Time) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, _ := s.RootModule().Resources[n]

		for timeStampAttr, serverVal := range map[string]time.Time{"created_on": *loadBalancerPool.CreatedOn, "modified_on": *loadBalancerPool.ModifiedOn} {
			timeStamp, err := time.Parse(time.RFC3339Nano, rs.Primary.Attributes[timeStampAttr])
			if err != nil {
				return err
			}

			if timeStamp != serverVal {
				return fmt.Errorf("state value of %s: %s is different than server created value: %s",
					timeStampAttr, rs.Primary.Attributes[timeStampAttr], serverVal.Format(time.RFC3339Nano))
			}

			// check retrieved values are reasonable
			// note this could fail if local time is out of sync with server time
			if timeStamp.Before(testStartTime) {
				return fmt.Errorf("State value of %s: %s should be greater than test start time: %s",
					timeStampAttr, timeStamp.Format(time.RFC3339Nano), testStartTime.Format(time.RFC3339Nano))
			}
		}

		return nil
	}
}

func testAccManuallyDeleteLoadBalancerPool(name string, loadBalancerPool *cloudflare.LoadBalancerPool, initialId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cloudflare.API)
		*initialId = loadBalancerPool.ID
		err := client.DeleteLoadBalancerPool(loadBalancerPool.ID)
		if err != nil {
			return err
		}
		return nil
	}
}

// using IPs from 192.0.2.0/24 as per RFC5737
func testAccCheckCloudFlareLoadBalancerPoolConfigBasic(id string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[1]s" {
  name = "my-tf-pool-basic-%[1]s"
  origins {
    name = "example-1"
    address = "192.0.2.1"
    enabled = true
  }
}`, id)
}

func testAccCheckCloudFlareLoadBalancerPoolConfigFullySpecified(id string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[1]s" {
  name = "my-tf-pool-basic-%[1]s"
  origins {
    name = "example-1"
    address = "192.0.2.1"
    enabled = false
  }
  origins {
    name = "example-2"
    address = "192.0.2.2"
  }
  check_regions = ["WEU"]
  description = "tfacc-fully-specified"
  enabled = false
  minimum_origins = 2
  // monitor = abcd TODO: monitor resource
  notification_email = "someone@example.com"
}`, id)
	// TODO add field to config after creating monitor resource
}
