package load_balancer_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	cfold "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/load_balancers"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

var (
	accountID = os.Getenv("CLOUDFLARE_ACCOUNT_ID")
)

func init() {
	resource.AddTestSweepers("cloudflare_load_balancer", &resource.Sweeper{
		Name: "cloudflare_load_balancer",
		F:    testSweepCloudflareLoadBalancer,
	})
}

func testSweepCloudflareLoadBalancer(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping load balancers sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	lbs, err := client.ListLoadBalancers(ctx, cfold.ZoneIdentifier(zoneID), cfold.ListLoadBalancerParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Load Balancers: %s", err))
		return err
	}

	if len(lbs) == 0 {
		tflog.Debug(ctx, "[DEBUG] No Cloudflare Load Balancers to sweep")
		return nil
	}

	tflog.Debug(ctx, fmt.Sprintf("[DEBUG] Found %d Cloudflare Load Balancers to sweep", len(lbs)))

	// Track deletion results
	deleted := 0
	failed := 0

	for _, lb := range lbs {
		// Use standard filtering helper
		if !utils.ShouldSweepResource(lb.Name) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Load Balancer: %s (%s)", lb.Name, lb.ID))

		err := client.DeleteLoadBalancer(ctx, cfold.ZoneIdentifier(zoneID), lb.ID)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete Load Balancer %s (%s): %s", lb.Name, lb.ID, err))
			failed++
			// Continue with other load balancers
		} else {
			tflog.Info(ctx, fmt.Sprintf("Deleted Load Balancer: %s (%s)", lb.Name, lb.ID))
			deleted++
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("[DEBUG] Load Balancer sweep completed: %d deleted, %d failed", deleted, failed))
	return nil
}

func TestAccCloudflareLoadBalancer_Basic(t *testing.T) {
	testStartTime := time.Now().UTC()
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerConfigBasic(zoneID, zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr(name, "pop_pools.#", "0"),
					resource.TestCheckResourceAttr(name, "country_pools.#", "0"),
					resource.TestCheckResourceAttr(name, "region_pools.#", "0"),
					// also expect api to generate some values
					testAccCheckCloudflareLoadBalancerDates(name, &loadBalancer, testStartTime),
					resource.TestCheckResourceAttr(name, "proxied", "false"), // default value
					resource.TestCheckResourceAttr(name, "ttl", "30"),
					resource.TestCheckResourceAttr(name, "steering_policy", "off"),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancer_SessionAffinity(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerConfigSessionAffinity(zoneID, zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
					// explicitly verify that our session_affinity has been set
					resource.TestCheckResourceAttr(name, "session_affinity", "cookie"),
					resource.TestCheckResourceAttr(name, "session_affinity_ttl", "1800"),
					resource.TestCheckResourceAttr(name, "session_affinity_attributes.samesite", "Auto"),
					resource.TestCheckResourceAttr(name, "session_affinity_attributes.secure", "Auto"),
					resource.TestCheckResourceAttr(name, "session_affinity_attributes.drain_duration", "60"),
					resource.TestCheckResourceAttr(name, "session_affinity_attributes.zero_downtime_failover", "sticky"),
					// dont check that other specified values are set, this will be evident by lack
					// of plan diff some values will get empty values
					resource.TestCheckResourceAttr(name, "pop_pools.#", "0"),
					resource.TestCheckResourceAttr(name, "country_pools.#", "0"),
					resource.TestCheckResourceAttr(name, "region_pools.#", "0"),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancer_SessionAffinityIPCookie(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerConfigSessionAffinityIPCookie(zoneID, zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
					// explicitly verify that our session_affinity has been set
					resource.TestCheckResourceAttr(name, "session_affinity", "ip_cookie"),
					// session_affinity_ttl should be present and set to its default
					resource.TestCheckResourceAttr(name, "session_affinity_ttl", "82800"),
					// session_affinity_attributes should be present and defaults set
					//resource.TestCheckResourceAttr(name, "session_affinity_attributes.%", "4"),
					resource.TestCheckResourceAttr(name, "session_affinity_attributes.drain_duration", "0"),
					resource.TestCheckResourceAttr(name, "session_affinity_attributes.samesite", "Auto"),
					resource.TestCheckResourceAttr(name, "session_affinity_attributes.secure", "Auto"),
					resource.TestCheckResourceAttr(name, "session_affinity_attributes.zero_downtime_failover", "none"),
					// dont check that other specified values are set, this will be evident by lack
					// of plan diff some values will get empty values
					resource.TestCheckResourceAttr(name, "pop_pools.#", "0"),
					resource.TestCheckResourceAttr(name, "region_pools.#", "0"),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancer_SessionAffinityHeader(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerConfigSessionAffinityHeader(zoneID, zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
					// explicitly verify that our session_affinity has been set
					resource.TestCheckResourceAttr(name, "session_affinity", "header"),
					resource.TestCheckResourceAttr(name, "session_affinity_ttl", "1800"),
					resource.TestCheckResourceAttr(name, "session_affinity_attributes.%", "6"),
					resource.TestCheckResourceAttr(name, "session_affinity_attributes.samesite", "Auto"),
					resource.TestCheckResourceAttr(name, "session_affinity_attributes.secure", "Auto"),
					resource.TestCheckResourceAttr(name, "session_affinity_attributes.drain_duration", "60"),
					resource.TestCheckResourceAttr(name, "session_affinity_attributes.zero_downtime_failover", "temporary"),
					resource.TestCheckResourceAttr(name, "session_affinity_attributes.headers.#", "1"),
					resource.TestCheckResourceAttr(name, "session_affinity_attributes.headers.0", "x-custom"),
					resource.TestCheckResourceAttr(name, "session_affinity_attributes.require_all_headers", "true"),
					// dont check that other specified values are set, this will be evident by lack
					// of plan diff some values will get empty values
					resource.TestCheckResourceAttr(name, "pop_pools.#", "0"),
					resource.TestCheckResourceAttr(name, "country_pools.#", "0"),
					resource.TestCheckResourceAttr(name, "region_pools.#", "0"),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancer_AdaptiveRouting(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerConfigAdaptiveRouting(zoneID, zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
					// explicitly verify that adaptive_routing has been set
					resource.TestCheckResourceAttr(name, "adaptive_routing.%", "1"),
					resource.TestCheckResourceAttr(name, "adaptive_routing.failover_across_pools", "true"),
					// dont check that other specified values are set, this will be evident by lack
					// of plan diff some values will get empty values
					resource.TestCheckResourceAttr(name, "pop_pools.#", "0"),
					resource.TestCheckResourceAttr(name, "country_pools.#", "0"),
					resource.TestCheckResourceAttr(name, "region_pools.#", "0"),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancer_LocationStrategy(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerConfigLocationStrategy(zoneID, zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
					// explicitly verify that location_strategy has been set
					resource.TestCheckResourceAttr(name, "location_strategy.%", "2"),
					resource.TestCheckResourceAttr(name, "location_strategy.prefer_ecs", "proximity"),
					resource.TestCheckResourceAttr(name, "location_strategy.mode", "pop"),
					// dont check that other specified values are set, this will be evident by lack
					// of plan diff some values will get empty values
					resource.TestCheckResourceAttr(name, "pop_pools.#", "0"),
					resource.TestCheckResourceAttr(name, "country_pools.#", "0"),
					resource.TestCheckResourceAttr(name, "region_pools.#", "0"),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancer_RandomSteering(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerConfigRandomSteering(zoneID, zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
					// explicitly verify that random_steering has been set
					resource.TestCheckResourceAttr(name, "random_steering.%", "2"),                   // random_steering appears once
					resource.TestCheckResourceAttr(name, "random_steering.pool_weights.%", "1"),      // one pool configured
					resource.TestCheckTypeSetElemAttr(name, "random_steering.pool_weights.*", "0.3"), // pool weight of 0.3
					resource.TestCheckResourceAttr(name, "random_steering.default_weight", "0.9"),    // default weight of 0.9
					// dont check that other specified values are set, this will be evident by lack
					// of plan diff some values will get empty values
					resource.TestCheckResourceAttr(name, "pop_pools.#", "0"),
					resource.TestCheckResourceAttr(name, "country_pools.#", "0"),
					resource.TestCheckResourceAttr(name, "region_pools.#", "0"),
				),
			},
			{
				Config: testAccCheckCloudflareLoadBalancerConfigRandomSteeringUpdate(zoneID, zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
					// explicitly verify that random_steering has been set
					resource.TestCheckResourceAttr(name, "random_steering.%", "2"),                   // random_steering appears once
					resource.TestCheckResourceAttr(name, "random_steering.pool_weights.%", "1"),      // one pool configured
					resource.TestCheckTypeSetElemAttr(name, "random_steering.pool_weights.*", "0.4"), // pool weight of 0.4
					resource.TestCheckResourceAttr(name, "random_steering.default_weight", "0.8"),    // default weight of 0.8
					// dont check that other specified values are set, this will be evident by lack
					// of plan diff some values will get empty values
					resource.TestCheckResourceAttr(name, "pop_pools.#", "0"),
					resource.TestCheckResourceAttr(name, "country_pools.#", "0"),
					resource.TestCheckResourceAttr(name, "region_pools.#", "0"),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancer_GeoBalancedUpdate(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerConfigGeoBalancedPoPCountry(zoneID, zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
					// checking our overrides of default values worked
					resource.TestCheckResourceAttr(name, "description", "tf-acctest load balancer using pop/country geo-balancing"),
					resource.TestCheckResourceAttr(name, "proxied", "true"),
					resource.TestCheckResourceAttr(name, "ttl", "30"),
					resource.TestCheckResourceAttr(name, "steering_policy", "geo"),
					resource.TestCheckResourceAttr(name, "pop_pools.%", "1"),
					resource.TestCheckResourceAttr(name, "country_pools.%", "1"),
					resource.TestCheckResourceAttr(name, "region_pools.%", "0"),
				),
			},
			{
				Config: testAccCheckCloudflareLoadBalancerConfigGeoBalancedPoPCountryToRegionUpdate(zoneID, zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
					// checking our updates to geo steering with only a region defined worked
					resource.TestCheckResourceAttr(name, "description", "tf-acctest load balancer using pop/country geo-balancing updated to region geo-balancing"),
					resource.TestCheckResourceAttr(name, "proxied", "true"),
					resource.TestCheckResourceAttr(name, "ttl", "30"),
					resource.TestCheckResourceAttr(name, "steering_policy", "geo"),
					resource.TestCheckResourceAttr(name, "pop_pools.%", "0"),
					resource.TestCheckResourceAttr(name, "country_pools.%", "0"),
					resource.TestCheckResourceAttr(name, "region_pools.%", "1"),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancer_GeoBalanced(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerConfigGeoBalanced(zoneID, zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
					// checking our overrides of default values worked
					resource.TestCheckResourceAttr(name, "description", "tf-acctest load balancer using geo-balancing"),
					resource.TestCheckResourceAttr(name, "proxied", "true"),
					resource.TestCheckResourceAttr(name, "ttl", "30"),
					resource.TestCheckResourceAttr(name, "steering_policy", "geo"),
					resource.TestCheckResourceAttr(name, "pop_pools.%", "1"),
					resource.TestCheckResourceAttr(name, "country_pools.%", "1"),
					resource.TestCheckResourceAttr(name, "region_pools.%", "1"),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancer_ProximityBalanced(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerConfigProximityBalanced(zoneID, zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
					// checking our overrides of default values worked
					resource.TestCheckResourceAttr(name, "description", "tf-acctest load balancer using proximity-balancing"),
					resource.TestCheckResourceAttr(name, "proxied", "true"),
					resource.TestCheckResourceAttr(name, "ttl", "30"),
					resource.TestCheckResourceAttr(name, "steering_policy", "proximity"),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancer_LeastOutstandingRequestsBalanced(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerConfigLeastOutstandingRequestsBalanced(zoneID, zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
					// checking our overrides of default values worked
					resource.TestCheckResourceAttr(name, "description", "tf-acctest load balancer using least outstanding requests steering"),
					resource.TestCheckResourceAttr(name, "proxied", "true"),
					resource.TestCheckResourceAttr(name, "ttl", "30"),
					resource.TestCheckResourceAttr(name, "steering_policy", "least_outstanding_requests"),
					resource.TestCheckResourceAttr(name, "rules.0.name", "test rule 1"),
					resource.TestCheckResourceAttr(name, "rules.0.condition", "dns.qry.type == 28"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.steering_policy", "least_outstanding_requests"),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancer_LeastConnectionsBalanced(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerConfigLeastConnectionsBalanced(zoneID, zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
					// checking our overrides of default values worked
					resource.TestCheckResourceAttr(name, "description", "tf-acctest load balancer using least connections steering"),
					resource.TestCheckResourceAttr(name, "proxied", "true"),
					resource.TestCheckResourceAttr(name, "ttl", "30"),
					resource.TestCheckResourceAttr(name, "steering_policy", "least_connections"),
					resource.TestCheckResourceAttr(name, "rules.0.name", "test rule 1"),
					resource.TestCheckResourceAttr(name, "rules.0.condition", "dns.qry.type == 28"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.steering_policy", "least_connections"),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancer_AdaptiveRoutingFailoverFalse(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{{
			Config: testAccCheckCloudflareLoadBalancerConfigAdaptiveRoutingFailoverFalse(zoneID, zone, rnd),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
				testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
			),
		}},
	})
}

func TestAccCloudflareLoadBalancer_AdaptiveRoutingFailoverTrue(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{{
			Config: testAccCheckCloudflareLoadBalancerConfigAdaptiveRoutingFailoverTrue(zoneID, zone, rnd),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
				testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
			),
		}},
	})
}

func TestAccCloudflareLoadBalancer_CountryPools(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{{
			Config: testAccCheckCloudflareLoadBalancerConfigCountryPools(zoneID, zone, rnd),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
				testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
			),
		}},
	})
}

func TestAccCloudflareLoadBalancer_CustomLocationStrategy(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{{
			Config: testAccCheckCloudflareLoadBalancerConfigCustomLocationStrategy(zoneID, zone, rnd),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
				testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
			),
		}},
	})
}

func TestAccCloudflareLoadBalancer_CustomPort(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{{
			Config: testAccCheckCloudflareLoadBalancerConfigCustomPort(zoneID, zone, rnd),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
				testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
			),
		}},
	})
}

func TestAccCloudflareLoadBalancer_CustomSessionAffinityAttributes(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{{
			Config: testAccCheckCloudflareLoadBalancerConfigCustomSessionAffinityAttributes(zoneID, zone, rnd),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
				testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
			),
		}},
	})
}

func TestAccCloudflareLoadBalancer_CustomTTL(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{{
			Config: testAccCheckCloudflareLoadBalancerConfigCustomTTL(zoneID, zone, rnd),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
				testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
			),
		}},
	})
}

func TestAccCloudflareLoadBalancer_LocationStrategyAlwaysResolverIP(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{{
			Config: testAccCheckCloudflareLoadBalancerConfigLocationStrategyAlwaysResolverIP(zoneID, zone, rnd),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
				testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
			),
		}},
	})
}

func TestAccCloudflareLoadBalancer_LocationStrategyNeverPop(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{{
			Config: testAccCheckCloudflareLoadBalancerConfigLocationStrategyNeverPop(zoneID, zone, rnd),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
				testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
			),
		}},
	})
}

func TestAccCloudflareLoadBalancer_StandardZone(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{{
			Config: testAccCheckCloudflareLoadBalancerConfigStandardZone(zoneID, zone, rnd),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
				testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
			),
		}},
	})
}

func TestAccCloudflareLoadBalancer_Rules(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerConfigRules(zoneID, zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
					// checking our overrides of default values worked
					resource.TestCheckResourceAttr(name, "description", "rules lb"),
					resource.TestCheckResourceAttr(name, "rules.0.name", "test rule 1"),
					resource.TestCheckResourceAttr(name, "rules.0.condition", "dns.qry.type == 28"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.steering_policy", "geo"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.session_affinity_attributes.samesite", "Auto"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.session_affinity_attributes.secure", "Auto"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.session_affinity_attributes.zero_downtime_failover", "sticky"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.adaptive_routing.%", "1"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.adaptive_routing.failover_across_pools", "true"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.location_strategy.%", "2"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.location_strategy.prefer_ecs", "always"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.location_strategy.mode", "resolver_ip"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.random_steering.pool_weights.%", "1"),
					resource.TestCheckTypeSetElemAttr(name, "rules.0.overrides.random_steering.pool_weights.*", "0.4"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.random_steering.default_weight", "0.2"),
					resource.TestCheckResourceAttr(name, "rules.#", "3"),
					resource.TestCheckResourceAttr(name, "rules.1.fixed_response.message_body", "hello"),
					resource.TestCheckResourceAttr(name, "rules.2.overrides.region_pools.%", "1"),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancer_Update(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	var initialId string
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareLoadBalancerConfigBasic(zoneID, zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
				),
			},
			{
				PreConfig: func() {
					initialId = loadBalancer.ID
				},
				Config: testAccCheckCloudflareLoadBalancerConfigGeoBalanced(zoneID, zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
					func(state *terraform.State) error {
						if initialId != loadBalancer.ID {
							// want in place update
							return fmt.Errorf("load balancer id is different after second config applied ( %s -> %s )", initialId, loadBalancer.ID)
						}
						return nil
					},
				),
			},
		},
	})
}

func testAccCheckCloudflareLoadBalancerDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_load_balancer" {
			continue
		}

		_, err := client.LoadBalancers.Get(context.Background(), rs.Primary.ID, load_balancers.LoadBalancerGetParams{
			ZoneID: cloudflare.F(rs.Primary.Attributes[consts.ZoneIDSchemaKey]),
		})
		if err == nil {
			return fmt.Errorf("load balancer still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckCloudflareLoadBalancerExists(n string, loadBalancer *cfold.LoadBalancer) resource.TestCheckFunc {
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
		foundLoadBalancer, err := client.GetLoadBalancer(context.Background(), cfold.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]), rs.Primary.ID)
		if err != nil {
			return err
		}

		*loadBalancer = foundLoadBalancer

		return nil
	}
}

func testAccCheckCloudflareLoadBalancerIDIsValid(n, expectedZoneID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		if len(rs.Primary.ID) != 32 {
			return fmt.Errorf("invalid id %q, should be a string of length 32", rs.Primary.ID)
		}

		if rs.Primary.Attributes[consts.ZoneIDSchemaKey] != expectedZoneID {
			return fmt.Errorf("zoneID attribute %q doesn't match the expected value %q", rs.Primary.Attributes[consts.ZoneIDSchemaKey], expectedZoneID)
		}

		return nil
	}
}

func testAccCheckCloudflareLoadBalancerDates(n string, loadBalancer *cfold.LoadBalancer, testStartTime time.Time) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[n]

		for timeStampAttr, serverVal := range map[string]time.Time{"created_on": *loadBalancer.CreatedOn, "modified_on": *loadBalancer.ModifiedOn} {
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

func testAccManuallyDeleteLoadBalancer(name string, loadBalancerID string, initialId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		_, err := acctest.SharedClient().LoadBalancers.Delete(
			context.Background(),
			loadBalancerID,
			load_balancers.LoadBalancerDeleteParams{
				ZoneID: cloudflare.F(rs.Primary.Attributes[consts.ZoneIDSchemaKey]),
			},
		)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckCloudflareLoadBalancerConfigBasic(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfigbasic.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigSessionAffinity(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfigsessionaffinity.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigSessionAffinityIPCookie(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfigsessionaffinityipcookie.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigSessionAffinityHeader(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfigsessionaffinityheader.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigAdaptiveRouting(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfigadaptiverouting.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigLocationStrategy(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfiglocationstrategy.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigRandomSteering(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfigrandomsteering.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigRandomSteeringUpdate(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfigrandomsteeringupdate.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigGeoBalanced(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfiggeobalanced.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigGeoBalancedPoPCountry(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfiggeobalancedpopcountry.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigGeoBalancedPoPCountryToRegionUpdate(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfiggeobalancedpopcountrytoregionupdate.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigProximityBalanced(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfigproximitybalanced.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigLeastOutstandingRequestsBalanced(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfigleastoutstandingrequestsbalanced.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigLeastConnectionsBalanced(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfigleastconnectionsbalanced.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigDuplicatePool(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfigduplicatepool.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigAdaptiveRoutingFailoverFalse(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfigadaptiveroutingfailoverfalse.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigAdaptiveRoutingFailoverTrue(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfigadaptiveroutingfailovertrue.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigCountryPools(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfigcountrypools.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigCustomLocationStrategy(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfigcustomlocationstrategy.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigCustomPort(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfigcustomport.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigCustomSessionAffinityAttributes(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfigcustomsessionaffinityattributes.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigCustomTTL(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfigcustomttl.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigLocationStrategyAlwaysResolverIP(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfiglocationstrategyalwaysresolverip.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigLocationStrategyNeverPop(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfiglocationstrategyneverpop.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigStandardZone(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfigstandardzone.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerConfigRules(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfigrules.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID string) string {
	return acctest.LoadTestCase("loadbalancerpoolconfigbasic.tf", id, accountID)
}

// TestAccLoadBalancer_RegionPoolsRemovalWithPoolDeletion tests the fix for the bug
// where removing region_pools and deleting pools in a single apply causes error 1005.
// This is User Story 1 (P1) - the core bug fix.
func TestAccLoadBalancer_RegionPoolsRemovalWithPoolDeletion(t *testing.T) {
	var loadBalancer cfold.LoadBalancer
	zone := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_load_balancer." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareLoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: Create load balancer with region_pools and the pools
				Config: testAccLoadBalancerConfigWithRegionPools(zoneID, zone, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
					resource.TestCheckResourceAttr(name, "description", "tf-acctest load balancer with region pools"),
					resource.TestCheckResourceAttr(name, "steering_policy", "geo"),
					resource.TestCheckResourceAttr(name, "region_pools.%", "2"), // WNAM and OC regions
				),
			},
			{
				// Step 2: Remove region_pools and delete the pools in a single apply
				// This is the critical test - should NOT fail with error 1005
				Config: testAccLoadBalancerConfigWithoutRegionPools(zoneID, zone, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
					testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
					resource.TestCheckResourceAttr(name, "description", "tf-acctest load balancer without region pools"),
					resource.TestCheckResourceAttr(name, "steering_policy", "dynamic"),
					resource.TestCheckNoResourceAttr(name, "region_pools.%"), // region_pools removed
				),
			},
		},
	})
}

// testAccLoadBalancerConfigWithRegionPools creates a load balancer with region_pools
// configuration that references two pools (CA and SG) across different regions.
func testAccLoadBalancerConfigWithRegionPools(zoneID, zone, id, accountID string) string {
	return fmt.Sprintf(`
# Create pools that will be referenced in region_pools
resource "cloudflare_load_balancer_pool" "test_pool_ca_%[3]s" {
  account_id = "%[4]s"
  name       = "tf-acctest-pool-ca-%[3]s"
  description = "CA region pool for region_pools test"
  enabled    = true
  minimum_origins = 1

  origins {
    name    = "origin-ca"
    address = "192.0.2.1"
    enabled = true
  }

  check_regions = ["WEU"]
}

resource "cloudflare_load_balancer_pool" "test_pool_sg_%[3]s" {
  account_id = "%[4]s"
  name       = "tf-acctest-pool-sg-%[3]s"
  description = "SG region pool for region_pools test"
  enabled    = true
  minimum_origins = 1

  origins {
    name    = "origin-sg"
    address = "192.0.2.2"
    enabled = true
  }

  check_regions = ["WEU"]
}

# Create a default pool (remains throughout the test)
resource "cloudflare_load_balancer_pool" "test_pool_default_%[3]s" {
  account_id = "%[4]s"
  name       = "tf-acctest-pool-default-%[3]s"
  description = "Default pool for region_pools test"
  enabled    = true
  minimum_origins = 1

  origins {
    name    = "origin-default"
    address = "192.0.2.10"
    enabled = true
  }

  check_regions = ["WEU"]
}

# Load balancer with region_pools configuration
resource "cloudflare_load_balancer" "%[3]s" {
  zone_id        = "%[1]s"
  name           = "tf-acctest-lb-%[3]s.%[2]s"
  description    = "tf-acctest load balancer with region pools"
  default_pools  = [cloudflare_load_balancer_pool.test_pool_default_%[3]s.id]
  fallback_pool  = cloudflare_load_balancer_pool.test_pool_default_%[3]s.id
  steering_policy = "geo"
  proxied        = true
  ttl            = 30

  # Geographic steering with region_pools
  region_pools {
    region   = "WNAM"
    pool_ids = [
      cloudflare_load_balancer_pool.test_pool_ca_%[3]s.id,
      cloudflare_load_balancer_pool.test_pool_default_%[3]s.id
    ]
  }

  region_pools {
    region   = "OC"
    pool_ids = [
      cloudflare_load_balancer_pool.test_pool_sg_%[3]s.id,
      cloudflare_load_balancer_pool.test_pool_default_%[3]s.id
    ]
  }
}
`, zoneID, zone, id, accountID)
}

// testAccLoadBalancerConfigWithoutRegionPools creates the same load balancer
// WITHOUT region_pools and WITHOUT the CA/SG pool resources.
// This tests the scenario where region_pools is removed and pools are deleted
// in a single apply - the core bug fix scenario.
func testAccLoadBalancerConfigWithoutRegionPools(zoneID, zone, id, accountID string) string {
	return fmt.Sprintf(`
# Default pool remains (not deleted)
resource "cloudflare_load_balancer_pool" "test_pool_default_%[3]s" {
  account_id = "%[4]s"
  name       = "tf-acctest-pool-default-%[3]s"
  description = "Default pool for region_pools test"
  enabled    = true
  minimum_origins = 1

  origins {
    name    = "origin-default"
    address = "192.0.2.10"
    enabled = true
  }

  check_regions = ["WEU"]
}

# NOTE: test_pool_ca and test_pool_sg resources are REMOVED from configuration
# This causes them to be deleted in the same apply where region_pools is removed

# Load balancer WITHOUT region_pools configuration
resource "cloudflare_load_balancer" "%[3]s" {
  zone_id        = "%[1]s"
  name           = "tf-acctest-lb-%[3]s.%[2]s"
  description    = "tf-acctest load balancer without region pools"
  default_pools  = [cloudflare_load_balancer_pool.test_pool_default_%[3]s.id]
  fallback_pool  = cloudflare_load_balancer_pool.test_pool_default_%[3]s.id
  steering_policy = "dynamic"  # Changed from "geo" to "dynamic"
  proxied        = true
  ttl            = 30

  # region_pools configuration REMOVED
  # The pools referenced in region_pools are also DELETED from configuration
  # This is the scenario that triggers error 1005 before the fix
}
`, zoneID, zone, id, accountID)
}
