package load_balancer_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
	cfv2 "github.com/cloudflare/cloudflare-go/v2"

	"github.com/cloudflare/cloudflare-go/v2/load_balancers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/pkg/errors"
)

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
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}

	lbs, err := client.ListLoadBalancers(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.ListLoadBalancerParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Load Balancers: %s", err))
	}

	if len(lbs) == 0 {
		log.Print("[DEBUG] No Cloudflare Load Balancers to sweep")
		return nil
	}

	for _, lb := range lbs {
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Load Balancer ID: %s", lb.ID))
		//nolint:errcheck
		client.DeleteLoadBalancer(ctx, cloudflare.ZoneIdentifier(zoneID), lb.ID)
	}

	return nil
}

func TestAccCloudflareLoadBalancer_Basic(t *testing.T) {
	testStartTime := time.Now().UTC()
	var loadBalancer cloudflare.LoadBalancer
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
	var loadBalancer cloudflare.LoadBalancer
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
					resource.TestCheckResourceAttr(name, "session_affinity_attributes.%", "4"),
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
	var loadBalancer cloudflare.LoadBalancer
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
	var loadBalancer cloudflare.LoadBalancer
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
	var loadBalancer cloudflare.LoadBalancer
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
	var loadBalancer cloudflare.LoadBalancer
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
	var loadBalancer cloudflare.LoadBalancer
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
					resource.TestCheckResourceAttr(name, "random_steering.%", "2"),                     // random_steering appears once
					resource.TestCheckResourceAttr(name, "random_steering.0.pool_weights.%", "1"),      // one pool configured
					resource.TestCheckTypeSetElemAttr(name, "random_steering.0.pool_weights.*", "0.4"), // pool weight of 0.4
					resource.TestCheckResourceAttr(name, "random_steering.default_weight", "0.8"),      // default weight of 0.8
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
	var loadBalancer cloudflare.LoadBalancer
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
	var loadBalancer cloudflare.LoadBalancer
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
	var loadBalancer cloudflare.LoadBalancer
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
	var loadBalancer cloudflare.LoadBalancer
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
					resource.TestCheckResourceAttr(name, "rules.0.overrides.%", "1"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.steering_policy", "least_outstanding_requests"),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancer_LeastConnectionsBalanced(t *testing.T) {
	var loadBalancer cloudflare.LoadBalancer
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
					resource.TestCheckResourceAttr(name, "rules.0.overrides.%", "1"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.steering_policy", "least_connections"),
				),
			},
		},
	})
}

func TestAccCloudflareLoadBalancer_Rules(t *testing.T) {
	var loadBalancer cloudflare.LoadBalancer
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
					resource.TestCheckResourceAttr(name, "rules.0.overrides.%", "5"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.steering_policy", "geo"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.session_affinity_attributes.%", "4"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.session_affinity_attributes.samesite", "Auto"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.session_affinity_attributes.secure", "Auto"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.session_affinity_attributes.zero_downtime_failover", "sticky"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.adaptive_routing.%", "1"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.adaptive_routing.failover_across_pools", "true"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.location_strategy.%", "2"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.location_strategy.prefer_ecs", "always"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.location_strategy.mode", "resolver_ip"),
					resource.TestCheckResourceAttr(name, "rules.0.overrides.random_steering.%", "1"),
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
	var loadBalancer cloudflare.LoadBalancer
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

func TestAccCloudflareLoadBalancer_CreateAfterManualDestroy(t *testing.T) {
	var loadBalancer cloudflare.LoadBalancer
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
					testAccManuallyDeleteLoadBalancer(name, &loadBalancer, &initialId),
				),
			},
			{
				Config: testAccCheckCloudflareLoadBalancerConfigBasic(zoneID, zone, rnd),
				Check: resource.ComposeTestCheckFunc(
					// testAccCheckCloudflareLoadBalancerExists(name, &loadBalancer),
					// testAccCheckCloudflareLoadBalancerIDIsValid(name, zoneID),
					func(state *terraform.State) error {
						if initialId == loadBalancer.ID {
							return fmt.Errorf("load balancer id is unchanged even after we thought we deleted it ( %s )", loadBalancer.ID)
						}
						return nil
					},
				),
			},
		},
	})
}

func testAccCheckCloudflareLoadBalancerDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_load_balancer" {
			continue
		}

		_, err := client.GetLoadBalancer(context.Background(), cloudflare.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("load balancer still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckCloudflareLoadBalancerExists(n string, loadBalancer *cloudflare.LoadBalancer) resource.TestCheckFunc {
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
		foundLoadBalancer, err := client.GetLoadBalancer(context.Background(), cloudflare.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]), rs.Primary.ID)
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

func testAccCheckCloudflareLoadBalancerDates(n string, loadBalancer *cloudflare.LoadBalancer, testStartTime time.Time) resource.TestCheckFunc {
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

func testAccManuallyDeleteLoadBalancer(name string, loadBalancer *cloudflare.LoadBalancer, initialId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		_, err := acctest.SharedV2Client().LoadBalancers.Delete(
			context.Background(),
			loadBalancer.ID,
			load_balancers.LoadBalancerDeleteParams{
				ZoneID: cfv2.F(rs.Primary.Attributes[consts.ZoneIDSchemaKey]),
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

func testAccCheckCloudflareLoadBalancerConfigRules(zoneID, zone, id string) string {
	return testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID) + acctest.LoadTestCase("loadbalancerconfigrules.tf", zoneID, zone, id)
}

func testAccCheckCloudflareLoadBalancerPoolConfigBasic(id, accountID string) string {
	return acctest.LoadTestCase("loadbalancerpoolconfigbasic.tf", id, accountID)
}
