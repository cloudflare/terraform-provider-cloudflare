package cloudflare

import (
	"errors"
	"fmt"
	"testing"

	"time"

	"os"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"regexp"
)

func TestAccCloudFlareLoadBalancer_Basic(t *testing.T) {
	// multiple instances of this config would conflict but we only use it once
	t.Parallel()
	testStartTime := time.Now().UTC()
	var loadBalancer cloudflare.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := acctest.RandString(10)
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareLoadBalancerConfigBasic(zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudFlareLoadBalancerIDIsValid(name, zone),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr(name, "pop_pools.#", "0"),
					resource.TestCheckResourceAttr(name, "region_pools.#", "0"),
					// also expect api to generate some values
					testAccCheckCloudFlareLoadBalancerDates(name, &loadBalancer, testStartTime),
					resource.TestCheckResourceAttr(name, "proxied", "false"), // default value
					resource.TestCheckResourceAttr(name, "ttl", "30"),
				),
			},
		},
	})
}

func TestAccCloudFlareLoadBalancer_GeoBalanced(t *testing.T) {
	t.Parallel()
	var loadBalancer cloudflare.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := acctest.RandString(10)
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareLoadBalancerConfigGeoBalanced(zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudFlareLoadBalancerIDIsValid(name, zone),
					// checking our overrides of default values worked
					resource.TestCheckResourceAttr(name, "description", "tf-acctest load balancer using geo-balancing"),
					resource.TestCheckResourceAttr(name, "proxied", "true"),
					resource.TestCheckResourceAttr(name, "ttl", "0"),
					resource.TestCheckResourceAttr(name, "pop_pools.#", "1"),
					resource.TestCheckResourceAttr(name, "region_pools.#", "1"),
				),
			},
		},
	})
}

func TestAccCloudFlareLoadBalancer_DuplicatePool(t *testing.T) {
	t.Parallel()
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckCloudFlareLoadBalancerConfigDuplicatePool(zone, rnd),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("duplicate entry specified for pop pool in location \"LAX\". each location must only be specified once")),
			},
		},
	})
}

/**
Any change to a load balancer  results in a new resource
Although the API client contains a modify method, this always results in 405 status
*/
func TestAccCloudFlareLoadBalancer_Update(t *testing.T) {
	t.Parallel()
	var loadBalancer cloudflare.LoadBalancer
	var initialId string
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := acctest.RandString(10)
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareLoadBalancerConfigBasic(zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudFlareLoadBalancerIDIsValid(name, zone),
				),
			},
			{
				PreConfig: func() {
					initialId = loadBalancer.ID
				},
				Config: testAccCheckCloudFlareLoadBalancerConfigGeoBalanced(zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudFlareLoadBalancerIDIsValid(name, zone),
					func(state *terraform.State) error {
						if initialId != loadBalancer.ID {
							// want in place update
							return fmt.Errorf("load balancer id is different after second config applied ( %s -> %s )",
								initialId, loadBalancer.ID)
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccCloudFlareLoadBalancer_CreateAfterManualDestroy(t *testing.T) {
	t.Parallel()
	var loadBalancer cloudflare.LoadBalancer
	var initialId string
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := acctest.RandString(10)
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFlareLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudFlareLoadBalancerConfigBasic(zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudFlareLoadBalancerIDIsValid(name, zone),
					testAccManuallyDeleteLoadBalancer(name, &loadBalancer, &initialId),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCloudFlareLoadBalancerConfigBasic(zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudFlareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudFlareLoadBalancerIDIsValid(name, zone),
					func(state *terraform.State) error {
						if initialId == loadBalancer.ID {
							return fmt.Errorf("load balancer id is unchanged even after we thought we deleted it ( %s )",
								loadBalancer.ID)
						}
						return nil
					},
				),
			},
		},
	})
}

func testAccCheckCloudFlareLoadBalancerDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_load_balancer" {
			continue
		}

		_, err := client.LoadBalancerDetails(rs.Primary.Attributes["zone_id"], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("load balancer still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckCloudFlareLoadBalancerExists(n string, loadBalancer *cloudflare.LoadBalancer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Load Balancer ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundLoadBalancer, err := client.LoadBalancerDetails(rs.Primary.Attributes["zone_id"], rs.Primary.ID)
		if err != nil {
			return err
		}

		*loadBalancer = foundLoadBalancer

		return nil
	}
}

func testAccCheckCloudFlareLoadBalancerIDIsValid(n, expectedZone string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		if len(rs.Primary.ID) != 32 {
			return fmt.Errorf("invalid id %q, should be a string of length 32", rs.Primary.ID)
		}

		if rs.Primary.Attributes["zone"] != expectedZone {
			return fmt.Errorf("zone attribute %q doesn't match the expected value %q", rs.Primary.Attributes["zone"], expectedZone)
		}

		if zoneId, ok := rs.Primary.Attributes["zone_id"]; !ok || len(zoneId) < 1 {
			return errors.New("zone_id is unset, should always be set with id")
		}
		return nil
	}
}

func testAccCheckCloudFlareLoadBalancerDates(n string, loadBalancer *cloudflare.LoadBalancer, testStartTime time.Time) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, _ := s.RootModule().Resources[n]

		for timeStampAttr, serverVal := range map[string]time.Time{"created_on": *loadBalancer.CreatedOn, "modified_on": *loadBalancer.ModifiedOn} {
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

func testAccManuallyDeleteLoadBalancer(name string, loadBalancer *cloudflare.LoadBalancer, initialId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		client := testAccProvider.Meta().(*cloudflare.API)
		*initialId = loadBalancer.ID
		err := client.DeleteLoadBalancer(rs.Primary.Attributes["zone_id"], rs.Primary.ID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckCloudFlareLoadBalancerConfigBasic(zone, id string) string {
	return testAccCheckCloudFlareLoadBalancerPoolConfigBasic(id) + fmt.Sprintf(`
resource "cloudflare_load_balancer" "%[2]s" {
  zone = "%[1]s"
  name = "tf-testacc-lb-%[2]s"
  fallback_pool_id = "${cloudflare_load_balancer_pool.%[2]s.id}"
  default_pool_ids = ["${cloudflare_load_balancer_pool.%[2]s.id}"]
}`, zone, id)
}

func testAccCheckCloudFlareLoadBalancerConfigGeoBalanced(zone, id string) string {
	return testAccCheckCloudFlareLoadBalancerPoolConfigBasic(id) + fmt.Sprintf(`
resource "cloudflare_load_balancer" "%[2]s" {
  zone = "%[1]s"
  name = "tf-testacc-lb-%[2]s"
  fallback_pool_id = "${cloudflare_load_balancer_pool.%[2]s.id}"
  default_pool_ids = ["${cloudflare_load_balancer_pool.%[2]s.id}"]
  description = "tf-acctest load balancer using geo-balancing"
  proxied = true // can't set ttl with proxied
  pop_pools {
    pop = "LAX"
    pool_ids = ["${cloudflare_load_balancer_pool.%[2]s.id}"]
  }
  region_pools {
    region = "WNAM"
    pool_ids = ["${cloudflare_load_balancer_pool.%[2]s.id}"]
  }
}`, zone, id)
}

func testAccCheckCloudFlareLoadBalancerConfigDuplicatePool(zone, id string) string {
	return testAccCheckCloudFlareLoadBalancerPoolConfigBasic(id) + fmt.Sprintf(`
resource "cloudflare_load_balancer" "%[2]s" {
  zone = "%[1]s"
  name = "tf-testacc-lb-%[2]s"
  fallback_pool_id = "${cloudflare_load_balancer_pool.%[2]s.id}"
  default_pool_ids = ["${cloudflare_load_balancer_pool.%[2]s.id}"]
  pop_pools {
    pop = "LAX"
    pool_ids = ["i_am_an_invalid_pool_id"]
  }
  pop_pools {
    pop = "LAX"
    pool_ids = ["${cloudflare_load_balancer_pool.%[2]s.id}"]
  }
}`, zone, id)
}
