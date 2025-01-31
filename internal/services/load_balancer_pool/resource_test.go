package load_balancer_pool_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	cfold "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/load_balancers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
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
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return errors.New("CLOUDFLARE_ACCOUNT_ID must be set")
	}

	pools, err := client.ListLoadBalancerPools(ctx, cfold.AccountIdentifier(accountID), cfold.ListLoadBalancerPoolParams{})
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
		client.DeleteLoadBalancerPool(ctx, cfold.AccountIdentifier(accountID), pool.ID)
	}

	return nil
}

func TestAccCloudflareLoadBalancerPool_Basic(t *testing.T) {
	// multiple instances of this config would conflict but we only use it once
	t.Parallel()
	testStartTime := time.Now().UTC()
	var loadBalancerPool cfold.LoadBalancerPool
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer_pool." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerPoolDestroy,
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
	var loadBalancerPool cfold.LoadBalancerPool
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer_pool." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerPoolConfigOriginSteeringLeastOutstandingRequests(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerPoolExists(name, &loadBalancerPool),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr(name, "check_regions.#", "0"),
					resource.TestCheckResourceAttr(name, "header.#", "0"),
					resource.TestCheckResourceAttr(name, "origin_steering.%", "1"),
					resource.TestCheckResourceAttr(name, "origin_steering.policy", "least_outstanding_requests"),
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
	var loadBalancerPool cfold.LoadBalancerPool
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer_pool." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerPoolConfigOriginSteeringLeastConnections(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerPoolExists(name, &loadBalancerPool),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr(name, "check_regions.#", "0"),
					resource.TestCheckResourceAttr(name, "header.#", "0"),
					resource.TestCheckResourceAttr(name, "origin_steering.%", "1"),
					resource.TestCheckResourceAttr(name, "origin_steering.policy", "least_connections"),
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

	var tunnelVirtualNetwork cfold.TunnelVirtualNetwork
	var loadBalancerPool cfold.LoadBalancerPool

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	vnetResID := utils.GenerateRandomResourceName()
	tunnelResID := utils.GenerateRandomResourceName()
	tunnelRouteResID := utils.GenerateRandomResourceName()
	poolResID := utils.GenerateRandomResourceName()

	vnetName := "cloudflare_zero_trust_tunnel_cloudflared_virtual_network." + vnetResID
	poolName := "cloudflare_load_balancer_pool." + poolResID

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerPoolDestroy,
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
					//// resource.TestCheckResourceAttr(poolName, "check_regions.#", "0"),
					//// resource.TestCheckResourceAttr(poolName, "header.#", "0"),
					// also expect api to generate some values
					testAccCheckCloudflareLoadBalancerPoolDates(poolName, &loadBalancerPool, testStartTime),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancerPool_FullySpecified(t *testing.T) {
	t.Parallel()
	var loadBalancerPool cfold.LoadBalancerPool
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer_pool." + rnd
	headerValue := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerPoolConfigFullySpecified(rnd, headerValue, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerPoolExists(name, &loadBalancerPool),
					// checking our overrides of default values worked
					resource.TestCheckResourceAttr(name, "enabled", "false"),
					resource.TestCheckResourceAttr(name, "load_shedding.%", "4"),
					resource.TestCheckResourceAttr(name, "load_shedding.default_percent", "55"),
					resource.TestCheckResourceAttr(name, "load_shedding.default_policy", "random"),
					resource.TestCheckResourceAttr(name, "load_shedding.session_percent", "12"),
					resource.TestCheckResourceAttr(name, "load_shedding.session_policy", "hash"),
					resource.TestCheckResourceAttr(name, "description", "tfacc-fully-specified"),
					resource.TestCheckResourceAttr(name, "check_regions.#", "1"),
					resource.TestCheckResourceAttr(name, "check_regions.0", "WEU"),
					resource.TestCheckResourceAttr(name, "minimum_origins", "2"),
					resource.TestCheckResourceAttr(name, "latitude", "12.3"),
					resource.TestCheckResourceAttr(name, "longitude", "55"),
					resource.TestCheckResourceAttr(name, "origin_steering.%", "1"),
					resource.TestCheckResourceAttr(name, "origin_steering.policy", "random"),
					resource.TestCheckResourceAttr(name, "origins.0.header.host.0", "test1.terraform.cfapi.net"),
					resource.TestCheckResourceAttr(name, "origins.1.header.host.0", "test2.terraform.cfapi.net"),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancerPool_CreateAfterManualDestroy(t *testing.T) {
	t.Parallel()
	var loadBalancerPool cfold.LoadBalancerPool
	var initialId string
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer_pool." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerPoolConfigBasic(rnd, accountID),
				Check:  resource.ComposeTestCheckFunc(
				// TODO: see if this is still actually needed
				// testAccCheckCloudflareLoadBalancerPoolExists(name, &loadBalancerPool),
				// testAccManuallyDeleteLoadBalancerPool(name, &loadBalancerPool, &initialId),
				),
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
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_load_balancer_pool" {
			continue
		}

		_, err := client.GetLoadBalancerPool(context.Background(), cfold.AccountIdentifier(accountID), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Load balancer pool still exists")
		}
	}

	return nil
}

func testAccCheckCloudflareLoadBalancerPoolExists(n string, loadBalancerPool *cfold.LoadBalancerPool) resource.TestCheckFunc {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Load Balancer ID is set")
		}

		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		foundLoadBalancerPool, err := client.GetLoadBalancerPool(context.Background(), cfold.AccountIdentifier(accountID), rs.Primary.ID)
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
		var tunnelVirtualNetwork cfold.TunnelVirtualNetwork
		var loadBalancerPool cfold.LoadBalancerPool

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

func testAccCheckCloudflareLoadBalancerPoolDates(n string, loadBalancerPool *cfold.LoadBalancerPool, testStartTime time.Time) resource.TestCheckFunc {
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

func testAccManuallyDeleteLoadBalancerPool(name string, loadBalancerPool *cfold.LoadBalancerPool, initialId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acctest.SharedClient()
		*initialId = loadBalancerPool.ID
		_, err := client.LoadBalancers.Pools.Delete(context.Background(), loadBalancerPool.ID, load_balancers.PoolDeleteParams{
			AccountID: cloudflare.F(os.Getenv("CLOUDFLARE_ACCOUNT_ID")),
		})
		if err != nil {
			return err
		}
		return nil
	}
}

// using IPs from 192.0.2.0/24 as per RFC5737.
func testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID string) string {
	return acctest.LoadTestCase("loadbalancerpoolconfigbasic.tf", id, accountID)
}

func testAccCheckCloudflareLoadBalancerPoolConfigOriginSteeringLeastOutstandingRequests(id, accountID string) string {
	return acctest.LoadTestCase("loadbalancerpoolconfigoriginsteeringleastoutstandingrequests.tf", id, accountID)
}

func testAccCheckCloudflareLoadBalancerPoolConfigOriginSteeringLeastConnections(id, accountID string) string {
	return acctest.LoadTestCase("loadbalancerpoolconfigoriginsteeringleastconnections.tf", id, accountID)
}

func testAccCheckCloudflareLoadBalancerPoolConfigVirtualNetworkID(accountID, vnetResID, tunnelResID, tunnelRouteResID, poolResID string) string {
	return acctest.LoadTestCase("loadbalancerpoolconfigvirtualnetworkid.tf", accountID, vnetResID, tunnelResID, tunnelRouteResID, poolResID)
}

func testAccCheckCloudflareLoadBalancerPoolConfigFullySpecified(id, headerValue, accountID string) string {
	return acctest.LoadTestCase("loadbalancerpoolconfigfullyspecified.tf", id, headerValue, accountID)
	// TODO add field to config after creating monitor resource
}

func testAccCheckCloudflareTunnelVirtualNetworkExists(name string, virtualNetwork *cfold.TunnelVirtualNetwork) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Tunnel Virtual Network is set")
		}

		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		foundTunnelVirtualNetworks, err := client.ListTunnelVirtualNetworks(context.Background(), cfold.AccountIdentifier(rs.Primary.Attributes[consts.AccountIDSchemaKey]), cfold.TunnelVirtualNetworksListParams{
			IsDeleted: cfold.BoolPtr(false),
			ID:        rs.Primary.ID,
		})

		if err != nil {
			return err
		}

		*virtualNetwork = foundTunnelVirtualNetworks[0]

		return nil
	}
}
