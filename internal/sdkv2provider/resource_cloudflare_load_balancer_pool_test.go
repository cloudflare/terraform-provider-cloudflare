package sdkv2provider

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/pkg/errors"
)

func init() {
	resource.AddTestSweepers("cloudflare_load_balancer_pool", &resource.Sweeper{
		Name: "cloudflare_load_balancer_pool",
		F:    testSweepCloudflareLoadBalancerPool,
	})
}

func testSweepCloudflareLoadBalancerPool(r string) error {
	ctx := context.Background()
	client, clientErr := sharedClient()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return errors.New("CLOUDFLARE_ACCOUNT_ID must be set")
	}

	pools, err := client.ListLoadBalancerPools(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListLoadBalancerPoolParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Load Balancer Pools: %s", err))
	}

	if len(pools) == 0 {
		log.Print("[DEBUG] No Cloudflare Load Balancer Pools to sweep")
		return nil
	}

	for _, pool := range pools {
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Load Balancer Pool ID: %s", pool.ID))
		//nolint:errcheck
		client.DeleteLoadBalancerPool(ctx, cloudflare.AccountIdentifier(accountID), pool.ID)
	}

	return nil
}

func TestAccCloudflareLoadBalancerPool_Basic(t *testing.T) {
	// multiple instances of this config would conflict but we only use it once
	t.Parallel()
	testStartTime := time.Now().UTC()
	var loadBalancerPool cloudflare.LoadBalancerPool
	rnd := generateRandomResourceName()
	name := "cloudflare_load_balancer_pool." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareLoadBalancerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerPoolConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerPoolExists(name, &loadBalancerPool),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr(name, "check_regions.#", "0"),
					resource.TestCheckResourceAttr(name, "header.#", "0"),
					// also expect api to generate some values
					testAccCheckCloudflareLoadBalancerPoolDates(name, &loadBalancerPool, testStartTime),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancerPool_OriginSteeringLeastOutstandingRequests(t *testing.T) {
	// multiple instances of this config would conflict but we only use it once
	t.Parallel()
	testStartTime := time.Now().UTC()
	var loadBalancerPool cloudflare.LoadBalancerPool
	rnd := generateRandomResourceName()
	name := "cloudflare_load_balancer_pool." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareLoadBalancerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerPoolConfigOriginSteeringLeastOutstandingRequests(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerPoolExists(name, &loadBalancerPool),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr(name, "check_regions.#", "0"),
					resource.TestCheckResourceAttr(name, "header.#", "0"),
					resource.TestCheckResourceAttr(name, "origin_steering.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(name, "origin_steering.*", map[string]string{
						"policy": "least_outstanding_requests",
					}),
					// also expect api to generate some values
					testAccCheckCloudflareLoadBalancerPoolDates(name, &loadBalancerPool, testStartTime),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancerPool_OriginSteeringLeastConnections(t *testing.T) {
	// multiple instances of this config would conflict but we only use it once
	t.Parallel()
	testStartTime := time.Now().UTC()
	var loadBalancerPool cloudflare.LoadBalancerPool
	rnd := generateRandomResourceName()
	name := "cloudflare_load_balancer_pool." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareLoadBalancerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerPoolConfigOriginSteeringLeastConnections(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerPoolExists(name, &loadBalancerPool),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr(name, "check_regions.#", "0"),
					resource.TestCheckResourceAttr(name, "header.#", "0"),
					resource.TestCheckResourceAttr(name, "origin_steering.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(name, "origin_steering.*", map[string]string{
						"policy": "least_connections",
					}),
					// also expect api to generate some values
					testAccCheckCloudflareLoadBalancerPoolDates(name, &loadBalancerPool, testStartTime),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancerPool_VirtualNetworkID(t *testing.T) {
	//
	// Note: We need to first set up a valid vnet that covers the address "192.0.2.1" or the LB API will complain with:
	// --> "virtual_network_id does not belong to tunnel route that covers origin IP: validation failed".
	//

	// multiple instances of this config would conflict but we only use it once
	t.Parallel()
	testStartTime := time.Now().UTC()

	var tunnelVirtualNetwork cloudflare.TunnelVirtualNetwork
	var loadBalancerPool cloudflare.LoadBalancerPool

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	vnetResID := generateRandomResourceName()
	tunnelResID := generateRandomResourceName()
	tunnelRouteResID := generateRandomResourceName()
	poolResID := generateRandomResourceName()

	vnetName := "cloudflare_zero_trust_tunnel_cloudflared_virtual_network." + vnetResID
	poolName := "cloudflare_load_balancer_pool." + poolResID

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareLoadBalancerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerPoolConfigVirtualNetworkID(accountID, vnetResID, tunnelResID, tunnelRouteResID, poolResID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareTunnelVirtualNetworkExists(vnetName, &tunnelVirtualNetwork),
					testAccCheckCloudflareLoadBalancerPoolExists(poolName, &loadBalancerPool),
					// check that virtual network ID is the same on the virtual network and load balancer pool
					testAccCheckCloudflareLoadBalancerPoolVirtualNetworkMatch(vnetName, poolName),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr(poolName, "check_regions.#", "0"),
					resource.TestCheckResourceAttr(poolName, "header.#", "0"),
					// also expect api to generate some values
					testAccCheckCloudflareLoadBalancerPoolDates(poolName, &loadBalancerPool, testStartTime),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancerPool_FullySpecified(t *testing.T) {
	t.Parallel()
	var loadBalancerPool cloudflare.LoadBalancerPool
	rnd := generateRandomResourceName()
	name := "cloudflare_load_balancer_pool." + rnd
	headerValue := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareLoadBalancerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerPoolConfigFullySpecified(rnd, headerValue, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerPoolExists(name, &loadBalancerPool),
					// checking our overrides of default values worked
					resource.TestCheckResourceAttr(name, "enabled", "false"),
					resource.TestCheckResourceAttr(name, "load_shedding.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(name, "load_shedding.*", map[string]string{
						"default_percent": "55",
						"default_policy":  "random",
						"session_percent": "12",
						"session_policy":  "hash",
					}),
					resource.TestCheckResourceAttr(name, "description", "tfacc-fully-specified"),
					resource.TestCheckResourceAttr(name, "check_regions.#", "1"),
					resource.TestCheckResourceAttr(name, "minimum_origins", "2"),
					resource.TestCheckResourceAttr(name, "latitude", "12.3"),
					resource.TestCheckResourceAttr(name, "longitude", "55"),
					resource.TestCheckResourceAttr(name, "origin_steering.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(name, "origin_steering.*", map[string]string{
						"policy": "random",
					}),
					func(state *terraform.State) error {
						for _, rs := range state.RootModule().Resources {
							for k, v := range rs.Primary.Attributes {
								r, _ := regexp.Compile("origins.*\\.header.*\\.header")

								if r.MatchString(k) {
									if v == "Host" {
										return nil
									}
								}
							}
						}
						return errors.New("Not equal")
					},
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancerPool_CreateAfterManualDestroy(t *testing.T) {
	t.Parallel()
	var loadBalancerPool cloudflare.LoadBalancerPool
	var initialId string
	rnd := generateRandomResourceName()
	name := "cloudflare_load_balancer_pool." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareLoadBalancerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerPoolConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerPoolExists(name, &loadBalancerPool),
					testAccManuallyDeleteLoadBalancerPool(name, &loadBalancerPool, &initialId),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCloudflareLoadBalancerPoolConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerPoolExists(name, &loadBalancerPool),
					func(state *terraform.State) error {
						if initialId == loadBalancerPool.ID {
							return fmt.Errorf("load balancer pool id is unchanged even after we thought we deleted it ( %s )", loadBalancerPool.ID)
						}
						return nil
					},
				),
			},
		},
	})
}

func testAccCheckCloudflareLoadBalancerPoolDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_load_balancer_pool" {
			continue
		}

		_, err := client.GetLoadBalancerPool(context.Background(), cloudflare.AccountIdentifier(accountID), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Load balancer pool still exists")
		}
	}

	return nil
}

func testAccCheckCloudflareLoadBalancerPoolExists(n string, loadBalancerPool *cloudflare.LoadBalancerPool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Load Balancer ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundLoadBalancerPool, err := client.GetLoadBalancerPool(context.Background(), cloudflare.AccountIdentifier(accountID), rs.Primary.ID)
		if err != nil {
			return err
		}

		*loadBalancerPool = foundLoadBalancerPool

		return nil
	}
}
func testAccCheckCloudflareLoadBalancerPoolVirtualNetworkMatch(vnetName, poolName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// fetch virtual network and pool and make sure they both reference the same virtual network ID
		var tunnelVirtualNetwork cloudflare.TunnelVirtualNetwork
		var loadBalancerPool cloudflare.LoadBalancerPool

		if err := testAccCheckCloudflareTunnelVirtualNetworkExists(vnetName, &tunnelVirtualNetwork)(s); err != nil {
			return err
		}

		if err := testAccCheckCloudflareLoadBalancerPoolExists(poolName, &loadBalancerPool)(s); err != nil {
			return err
		}

		tunnelVnet := tunnelVirtualNetwork.ID
		if tunnelVnet == "" {
			return fmt.Errorf("No Tunnel Virtual Network ID set")
		}

		originVnet := loadBalancerPool.Origins[0].VirtualNetworkID
		if originVnet == "" {
			return fmt.Errorf("No Origin Virtual Network ID set")
		}

		if tunnelVnet != originVnet {
			return fmt.Errorf("Tunnel Virtual Network %q does not match Origin Virtual Network %q", tunnelVnet, originVnet)
		}

		// inspect the pool's terraform attribute directly and make sure it matches
		if err := resource.TestCheckResourceAttr(poolName, "origins.0.virtual_network_id", tunnelVnet)(s); err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckCloudflareLoadBalancerPoolDates(n string, loadBalancerPool *cloudflare.LoadBalancerPool, testStartTime time.Time) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[n]

		for timeStampAttr, serverVal := range map[string]time.Time{"created_on": *loadBalancerPool.CreatedOn, "modified_on": *loadBalancerPool.ModifiedOn} {
			timeStamp, err := time.Parse(time.RFC3339Nano, rs.Primary.Attributes[timeStampAttr])
			if err != nil {
				return err
			}

			if timeStamp != serverVal {
				return fmt.Errorf("state value of %s: %s is different than server created value: %s", timeStampAttr, rs.Primary.Attributes[timeStampAttr], serverVal.Format(time.RFC3339Nano))
			}

			// check retrieved values are reasonable
			// note this could fail if local time is out of sync with server time
			if timeStamp.Before(testStartTime) {
				return fmt.Errorf("state value of %s: %s should be greater than test start time: %s", timeStampAttr, timeStamp.Format(time.RFC3339Nano), testStartTime.Format(time.RFC3339Nano))
			}
		}

		return nil
	}
}

func testAccManuallyDeleteLoadBalancerPool(name string, loadBalancerPool *cloudflare.LoadBalancerPool, initialId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cloudflare.API)
		*initialId = loadBalancerPool.ID
		err := client.DeleteLoadBalancerPool(context.Background(), cloudflare.AccountIdentifier(os.Getenv("CLOUDFLARE_ACCOUNT_ID")), loadBalancerPool.ID)
		if err != nil {
			return err
		}
		return nil
	}
}

// using IPs from 192.0.2.0/24 as per RFC5737.
func testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[1]s" {
  account_id = "%[2]s"
  name = "my-tf-pool-basic-%[1]s"
  latitude = 12.3
  longitude = 55
  origins {
    name = "example-1"
    address = "192.0.2.1"
    enabled = true
  }
}`, id, accountID)
}

func testAccCheckCloudflareLoadBalancerPoolConfigOriginSteeringLeastOutstandingRequests(id, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[1]s" {
  account_id = "%[2]s"
  name = "my-tf-pool-basic-%[1]s"
  latitude = 12.3
  longitude = 55
  origins {
    name = "example-1"
    address = "192.0.2.1"
    enabled = true
  }
  origin_steering {
    policy = "least_outstanding_requests"
  }
}`, id, accountID)
}

func testAccCheckCloudflareLoadBalancerPoolConfigOriginSteeringLeastConnections(id, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[1]s" {
  account_id = "%[2]s"
  name = "my-tf-pool-basic-%[1]s"
  latitude = 12.3
  longitude = 55
  origins {
    name = "example-1"
    address = "192.0.2.1"
    enabled = true
  }
  origin_steering {
    policy = "least_connections"
  }
}`, id, accountID)
}

func testAccCheckCloudflareLoadBalancerPoolConfigVirtualNetworkID(accountID, vnetResID, tunnelResID, tunnelRouteResID, poolResID string) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_tunnel_cloudflared_virtual_network" "%[2]s" {
	account_id = "%[1]s"
	name       = "my-tf-vnet-for-pool-%[2]s"
	comment     = "test"
}

resource "cloudflare_zero_trust_tunnel_cloudflared" "%[3]s" {
	account_id = "%[1]s"
	name       = "my-tf-tunnel-for-pool-%[3]s"
	secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}

resource "cloudflare_zero_trust_tunnel_cloudflared_route" "%[4]s" {
    account_id = "%[1]s"
    network = "192.0.2.1/32"
    virtual_network_id = cloudflare_zero_trust_tunnel_cloudflared_virtual_network.%[2]s.id
    tunnel_id = cloudflare_zero_trust_tunnel_cloudflared.%[3]s.id
    comment = "test"
}

resource "cloudflare_load_balancer_pool" "%[5]s" {
  account_id = "%[1]s"
  name = "my-tf-pool-with-vnet-%[5]s"
  latitude = 12.3
  longitude = 55
  origins {
    name = "example-1"
    address = "192.0.2.1"
    virtual_network_id = cloudflare_zero_trust_tunnel_cloudflared_route.%[4]s.virtual_network_id
    enabled = true
  }
}`, accountID, vnetResID, tunnelResID, tunnelRouteResID, poolResID)
}

func testAccCheckCloudflareLoadBalancerPoolConfigFullySpecified(id, headerValue, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_load_balancer_pool" "%[1]s" {
  name = "my-tf-pool-basic-%[1]s"
  account_id = "%[3]s"
  origins {
    name = "example-1"
    address = "192.0.2.1"
    enabled = false
    weight = 1.0
    header {
      header = "Host"
      values = ["test1.%[2]s"]
     }
  }

  origins {
    name = "example-2"
    address = "192.0.2.2"
    weight = 0.5
    header {
      header = "Host"
      values = ["test2.%[2]s"]
    }
  }

  load_shedding {
    default_percent = 55
    default_policy = "random"
    session_percent = 12
    session_policy = "hash"
  }

  latitude = 12.3
  longitude = 55

  origin_steering {
    policy = "random"
  }

  check_regions = ["WEU"]
  description = "tfacc-fully-specified"
  enabled = false
  minimum_origins = 2
  // monitor = abcd TODO: monitor resource
  notification_email = "someone@example.com"
}`, id, headerValue, accountID)
	// TODO add field to config after creating monitor resource
}
