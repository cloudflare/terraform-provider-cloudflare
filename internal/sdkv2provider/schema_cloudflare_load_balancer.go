package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// "" or "none" = session affinity is disabled.
var sessionAffinityPolicies = []string{"", "none", "cookie", "ip_cookie", "header"}

var (
	loadBalancerSessionAffinityAttributesElem = &schema.Resource{
		Schema: map[string]*schema.Schema{
			"samesite": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Auto", "Lax", "None", "Strict"}, false),
				Default:      "Auto",
				Description:  fmt.Sprintf("Configures the SameSite attribute on session affinity cookie. Value `Auto` will be translated to `Lax` or `None` depending if Always Use HTTPS is enabled. Note: when using value `None`, then you can not set [`secure=\"Never\"`](#secure). %s", renderAvailableDocumentationValuesStringSlice([]string{"Auto", "Lax", "None", "Strict"})),
			},

			"secure": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "Auto",
				ValidateFunc: validation.StringInSlice([]string{"Auto", "Always", "Never"}, false),
				Description:  fmt.Sprintf("Configures the Secure attribute on session affinity cookie. Value `Always` indicates the Secure attribute will be set in the Set-Cookie header, `Never` indicates the Secure attribute will not be set, and `Auto` will set the Secure attribute depending if Always Use HTTPS is enabled. %s", renderAvailableDocumentationValuesStringSlice([]string{"Auto", "Always", "Never"})),
			},

			"drain_duration": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      0,
				ValidateFunc: validation.IntBetween(0, 86400), // less than 24 hours
				Description:  "Configures the drain duration in seconds. This field is only used when session affinity is enabled on the load balancer.",
			},

			"zero_downtime_failover": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: validation.StringInSlice([]string{"none", "temporary", "sticky"}, false),
				Description:  fmt.Sprintf("Configures the zero-downtime failover between origins within a pool when session affinity is enabled. Value `none` means no failover takes place for sessions pinned to the origin. Value `temporary` means traffic will be sent to another other healthy origin until the originally pinned origin is available; note that this can potentially result in heavy origin flapping. Value `sticky` means the session affinity cookie is updated and subsequent requests are sent to the new origin. This feature is currently incompatible with Argo, Tiered Cache, and Bandwidth Alliance. %s", renderAvailableDocumentationValuesStringSlice([]string{"none", "temporary", "sticky"})),
			},

			"headers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringLenBetween(1, 100),
				},
				Description: "Configures the HTTP header names to use when header session affinity is enabled.",
			},

			"require_all_headers": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Configures how headers are used when header session affinity is enabled. Set to true to require all headers to be present on requests in order for sessions to be created or false to require at least one header to be present.",
			},
		},
	}

	loadBalancerOverridesSessionAffinityAttributesElem = &schema.Resource{
		Schema: map[string]*schema.Schema{
			"samesite": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Auto", "Lax", "None", "Strict"}, false),
				Description:  fmt.Sprintf("Configures the SameSite attribute on session affinity cookie. Value `Auto` will be translated to `Lax` or `None` depending if Always Use HTTPS is enabled. Note: when using value `None`, then you can not set [`secure=\"Never\"`](#secure). %s", renderAvailableDocumentationValuesStringSlice([]string{"Auto", "Lax", "None", "Strict"})),
			},

			"secure": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Auto", "Always", "Never"}, false),
				Description:  fmt.Sprintf("Configures the Secure attribute on session affinity cookie. Value `Always` indicates the Secure attribute will be set in the Set-Cookie header, `Never` indicates the Secure attribute will not be set, and `Auto` will set the Secure attribute depending if Always Use HTTPS is enabled. %s", renderAvailableDocumentationValuesStringSlice([]string{"Auto", "Always", "Never"})),
			},

			//
			// "drain_duration" not currently supported as a rule override
			//

			"zero_downtime_failover": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"none", "temporary", "sticky"}, false),
				Description:  fmt.Sprintf("Configures the zero-downtime failover between origins within a pool when session affinity is enabled. Value `none` means no failover takes place for sessions pinned to the origin. Value `temporary` means traffic will be sent to another other healthy origin until the originally pinned origin is available; note that this can potentially result in heavy origin flapping. Value `sticky` means the session affinity cookie is updated and subsequent requests are sent to the new origin. This feature is currently incompatible with Argo, Tiered Cache, and Bandwidth Alliance. %s", renderAvailableDocumentationValuesStringSlice([]string{"none", "temporary", "sticky"})),
			},

			"headers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringLenBetween(1, 100),
				},
				Description: "Configures the HTTP header names to use when header session affinity is enabled.",
			},

			"require_all_headers": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Configures how headers are used when header session affinity is enabled. Set to true to require all headers to be present on requests in order for sessions to be created or false to require at least one header to be present.",
			},
		},
	}

	loadBalancerAdaptiveRoutingElem = &schema.Resource{
		Schema: map[string]*schema.Schema{
			"failover_across_pools": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Extends zero-downtime failover of requests to healthy origins from alternate pools, when no healthy alternate exists in the same pool, according to the failover order defined by traffic and origin steering. When set `false`, zero-downtime failover will only occur between origins within the same pool.",
			},
		},
	}

	loadBalancerOverridesAdaptiveRoutingElem = &schema.Resource{
		Schema: map[string]*schema.Schema{
			"failover_across_pools": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Extends zero-downtime failover of requests to healthy origins from alternate pools, when no healthy alternate exists in the same pool, according to the failover order defined by traffic and origin steering. When set `false`, zero-downtime failover will only occur between origins within the same pool.",
			},
		},
	}

	loadBalancerLocationStrategyElem = &schema.Resource{
		Schema: map[string]*schema.Schema{
			"prefer_ecs": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"always", "never", "proximity", "geo"}, false),
				Default:      "proximity",
				Description:  fmt.Sprintf("Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the authoritative location. Value `always` will always prefer ECS, `never` will never prefer ECS, `proximity` will prefer ECS only when [`steering_policy=\"proximity\"`](#steering_policy), and `geo` will prefer ECS only when [`steering_policy=\"geo\"`](#steering_policy). %s", renderAvailableDocumentationValuesStringSlice([]string{"always", "never", "proximity", "geo"})),
			},

			"mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "pop",
				ValidateFunc: validation.StringInSlice([]string{"pop", "resolver_ip"}, false),
				Description:  fmt.Sprintf("Determines the authoritative location when ECS is not preferred, does not exist in the request, or its GeoIP lookup is unsuccessful. Value `pop` will use the Cloudflare PoP location. Value `resolver_ip` will use the DNS resolver GeoIP location. If the GeoIP lookup is unsuccessful, it will use the Cloudflare PoP location. %s", renderAvailableDocumentationValuesStringSlice([]string{"pop", "resolver_ip"})),
			},
		},
	}

	loadBalancerOverridesLocationStrategyElem = &schema.Resource{
		Schema: map[string]*schema.Schema{
			"prefer_ecs": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"always", "never", "proximity", "geo"}, false),
				Description:  fmt.Sprintf("Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the authoritative location. Value `always` will always prefer ECS, `never` will never prefer ECS, `proximity` will prefer ECS only when [`steering_policy=\"proximity\"`](#steering_policy), and `geo` will prefer ECS only when [`steering_policy=\"geo\"`](#steering_policy). %s", renderAvailableDocumentationValuesStringSlice([]string{"always", "never", "proximity", "geo"})),
			},

			"mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"pop", "resolver_ip"}, false),
				Description:  fmt.Sprintf("Determines the authoritative location when ECS is not preferred, does not exist in the request, or its GeoIP lookup is unsuccessful. Value `pop` will use the Cloudflare PoP location. Value `resolver_ip` will use the DNS resolver GeoIP location. If the GeoIP lookup is unsuccessful, it will use the Cloudflare PoP location. %s", renderAvailableDocumentationValuesStringSlice([]string{"pop", "resolver_ip"})),
			},
		},
	}

	loadBalancerRandomSteeringElem = &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pool_weights": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeFloat,
					ValidateFunc: validation.FloatBetween(0, 1),
				},
				Description: "A mapping of pool IDs to custom weights. The weight is relative to other pools in the load balancer.",
			},

			"default_weight": {
				Type:         schema.TypeFloat,
				Optional:     true,
				ValidateFunc: validation.FloatBetween(0, 1),
				Description:  "The default weight for pools in the load balancer that are not specified in the [`pool_weights`](#pool_weights) map.",
			},
		},
	}

	loadBalancerOverridesRandomSteeringElem = &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pool_weights": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeFloat,
					ValidateFunc: validation.FloatBetween(0, 1),
				},
				Description: "A mapping of pool IDs to custom weights. The weight is relative to other pools in the load balancer.",
			},

			"default_weight": {
				Type:         schema.TypeFloat,
				Optional:     true,
				ValidateFunc: validation.FloatBetween(0, 1),
				Description:  "The default weight for pools in the load balancer that are not specified in the [`pool_weights`](#pool_weights) map.",
			},
		},
	}

	loadBalancerPopPoolElem = &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pop": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A 3-letter code for the Point-of-Presence. Allowed values can be found in the list of datacenters on the [status page](https://www.cloudflarestatus.com/). Multiple entries should not be specified with the same PoP.",
				// let the api handle validating pops
			},

			"pool_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringLenBetween(1, 32),
				},
				Description: "A list of pool IDs in failover priority to use for traffic reaching the given PoP.",
			},
		},
	}

	loadBalancerOverridesPopPoolElem = &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pop": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A 3-letter code for the Point-of-Presence. Allowed values can be found in the list of datacenters on the [status page](https://www.cloudflarestatus.com/). Multiple entries should not be specified with the same PoP.",
			},

			"pool_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringLenBetween(1, 32),
				},
				Description: "A list of pool IDs in failover priority to use for traffic reaching the given PoP.",
			},
		},
	}

	loadBalancerCountryPoolElem = &schema.Resource{
		Schema: map[string]*schema.Schema{
			"country": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A country code which can be determined with the Load Balancing Regions API described [here](https://developers.cloudflare.com/load-balancing/reference/region-mapping-api/). Multiple entries should not be specified with the same country.",
				// let the api handle validating countries
			},

			"pool_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringLenBetween(1, 32),
				},
				Description: "A list of pool IDs in failover priority to use in the given country.",
			},
		},
	}

	loadBalancerOverridesCountryPoolElem = &schema.Resource{
		Schema: map[string]*schema.Schema{
			"country": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A country code which can be determined with the Load Balancing Regions API described [here](https://developers.cloudflare.com/load-balancing/reference/region-mapping-api/). Multiple entries should not be specified with the same country.",
			},

			"pool_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringLenBetween(1, 32),
				},
				Description: "A list of pool IDs in failover priority to use in the given country.",
			},
		},
	}

	loadBalancerRegionPoolElem = &schema.Resource{
		Description: "A set containing mappings of region codes to a list of pool IDs (ordered by their failover priority) for the given region. Fields documented below.",
		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A region code which must be in the list defined [here](https://developers.cloudflare.com/load-balancing/reference/region-mapping-api/#list-of-load-balancer-regions). Multiple entries should not be specified with the same region.",
				// let the api handle validating regions
			},

			"pool_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringLenBetween(1, 32),
				},
				Description: "A list of pool IDs in failover priority to use in the given region.",
			},
		},
	}

	loadBalancerOverridesRegionPoolElem = &schema.Resource{
		Description: "A set containing mappings of region codes to a list of pool IDs (ordered by their failover priority) for the given region. Fields documented below.",
		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A region code which must be in the list defined [here](https://developers.cloudflare.com/load-balancing/reference/region-mapping-api/#list-of-load-balancer-regions). Multiple entries should not be specified with the same region.",
			},

			"pool_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringLenBetween(1, 32),
				},
				Description: "A list of pool IDs in failover priority to use in the given region.",
			},
		},
	}

	loadBalancerLocalPoolElems = map[string]*schema.Resource{
		"pop":     loadBalancerPopPoolElem,
		"region":  loadBalancerRegionPoolElem,
		"country": loadBalancerCountryPoolElem,
	}

	loadBalancerOverridesLocalPoolElems = map[string]*schema.Resource{
		"pop":     loadBalancerOverridesPopPoolElem,
		"region":  loadBalancerOverridesRegionPoolElem,
		"country": loadBalancerOverridesCountryPoolElem,
	}

	rulesElem = &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 200),
				Description:  "Human readable name for this rule.",
			},

			"priority": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Priority used when determining the order of rule execution. Lower values are executed first. If not provided, the list order will be used.",
			},

			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "A disabled rule will not be executed.",
			},

			"condition": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The statement to evaluate to determine if this rule's effects should be applied. An empty condition is always true. See [load balancing rules](https://developers.cloudflare.com/load-balancing/understand-basics/load-balancing-rules).",
			},

			"terminates": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Terminates indicates that if this rule is true no further rules should be executed. Note: setting a [`fixed_response`](#fixed_response) forces this field to `true`.",
			},

			"overrides": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The load balancer settings to alter if this rule's [`condition`](#condition) is true. Note: [`overrides`](#overrides) or [`fixed_response`](#fixed_response) must be set.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"session_affinity": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice(sessionAffinityPolicies, false),
							Description:  "Configure attributes for session affinity.",
						},

						"session_affinity_ttl": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validation.IntBetween(1800, 604800),
							Description:  "Time, in seconds, until this load balancer's session affinity cookie expires after being created. This parameter is ignored unless a supported session affinity policy is set. The current default of `82800` (23 hours) will be used unless [`session_affinity_ttl`](#session_affinity_ttl) is explicitly set. Once the expiry time has been reached, subsequent requests may get sent to a different origin server. Valid values are between `1800` and `604800`.",
						},

						"session_affinity_attributes": {
							Type:        schema.TypeSet,
							Optional:    true,
							Elem:        loadBalancerOverridesSessionAffinityAttributesElem,
							Description: "Configure attributes for session affinity. Note that the property [`drain_duration`](#drain_duration) is not currently supported as a rule override.",
						},

						"adaptive_routing": {
							Type:        schema.TypeSet,
							Optional:    true,
							Elem:        loadBalancerOverridesAdaptiveRoutingElem,
							Description: "Controls features that modify the routing of requests to pools and origins in response to dynamic conditions, such as during the interval between active health monitoring requests.",
						},

						"location_strategy": {
							Type:        schema.TypeSet,
							Optional:    true,
							Elem:        loadBalancerOverridesLocationStrategyElem,
							Description: "Controls location-based steering for non-proxied requests.",
						},

						"random_steering": {
							Type:        schema.TypeSet,
							Optional:    true,
							Elem:        loadBalancerOverridesRandomSteeringElem,
							Description: "Configures pool weights. When [`steering_policy=\"random\"`](#steering_policy), a random pool is selected with probability proportional to pool weights. When [`steering_policy=\"least_outstanding_requests\"`](#steering_policy), pool weights are used to scale each pool's outstanding requests.",
						},

						"ttl": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Time to live (TTL) of the DNS entry for the IP address returned by this load balancer. This cannot be set for proxied load balancers. Defaults to `30`.",
						},

						"steering_policy": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"off", "geo", "dynamic_latency", "random", "proximity", "least_outstanding_requests", ""}, false),
							Description:  fmt.Sprintf("The method the load balancer uses to determine the route to your origin. Value `off` uses [`default_pool_ids`](#default_pool_ids). Value `geo` uses [`pop_pools`](#pop_pools)/[`country_pools`](#country_pools)/[`region_pools`](#region_pools). For non-proxied requests, the [`country`](#country) for [`country_pools`](#country_pools) is determined by [`location_strategy`](#location_strategy). Value `random` selects a pool randomly. Value `dynamic_latency` uses round trip time to select the closest pool in [`default_pool_ids`](#default_pool_ids) (requires pool health checks). Value `proximity` uses the pools' latitude and longitude to select the closest pool using the Cloudflare PoP location for proxied requests or the location determined by [`location_strategy`](#location_strategy) for non-proxied requests. Value `least_outstanding_requests` selects a pool by taking into consideration [`random_steering`](#random_steering) weights, as well as each pool's number of outstanding requests. Pools with more pending requests are weighted proportionately less relative to others. Value `\"\"` maps to `geo` if you use [`pop_pools`](#pop_pools)/[`country_pools`](#country_pools)/[`region_pools`](#region_pools) otherwise `off`. %s Defaults to `\"\"`.", renderAvailableDocumentationValuesStringSlice([]string{"off", "geo", "dynamic_latency", "random", "proximity", "least_outstanding_requests", `""`})),
						},

						"fallback_pool": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The pool ID to use when all other pools are detected as unhealthy.",
						},

						"default_pools": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "A list of pool IDs ordered by their failover priority. Used whenever [`pop_pools`](#pop_pools)/[`country_pools`](#country_pools)/[`region_pools`](#region_pools) are not defined.",
						},

						"pop_pools": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Elem:        loadBalancerOverridesPopPoolElem,
							Description: "A set containing mappings of Cloudflare Point-of-Presence (PoP) identifiers to a list of pool IDs (ordered by their failover priority) for the PoP (datacenter). This feature is only available to enterprise customers.",
						},

						"country_pools": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Elem:        loadBalancerOverridesCountryPoolElem,
							Description: "A set containing mappings of country codes to a list of pool IDs (ordered by their failover priority) for the given country.",
						},

						"region_pools": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Elem:        loadBalancerOverridesRegionPoolElem,
							Description: "A set containing mappings of region codes to a list of pool IDs (ordered by their failover priority) for the given region.",
						},
					},
				},
			},

			"fixed_response": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Settings for a HTTP response to return directly to the eyeball if the condition is true. Note: [`overrides`](#overrides) or [`fixed_response`](#fixed_response) must be set.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"message_body": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringLenBetween(0, 1024),
							Description:  "The text used as the html body for this fixed response.",
						},

						"status_code": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The HTTP status code used for this fixed response.",
						},

						"content_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringLenBetween(0, 32),
							Description:  "The value of the HTTP context-type header for this fixed response.",
						},

						"location": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringLenBetween(0, 2048),
							Description:  "The value of the HTTP location header for this fixed response.",
						},
					},
				},
			},
		},
	}
)

func resourceCloudflareLoadBalancerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The zone ID to add the load balancer to.",
		},

		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The DNS hostname to associate with your load balancer. If this hostname already exists as a DNS record in Cloudflare's DNS, the load balancer will take precedence and the DNS record will not be used.",
		},

		"fallback_pool_id": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringLenBetween(1, 32),
			Description:  "The pool ID to use when all other pools are detected as unhealthy.",
		},

		"default_pool_ids": {
			Type:     schema.TypeList,
			Required: true,
			MinItems: 1,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(1, 32),
			},
			Description: "A list of pool IDs ordered by their failover priority. Used whenever [`pop_pools`](#pop_pools)/[`country_pools`](#country_pools)/[`region_pools`](#region_pools) are not defined.",
		},

		"session_affinity": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "none",
			ValidateFunc: validation.StringInSlice(sessionAffinityPolicies, false),
			Description:  fmt.Sprintf("Specifies the type of session affinity the load balancer should use unless specified as `none` or `\"\"` (default). With value `cookie`, on the first request to a proxied load balancer, a cookie is generated, encoding information of which origin the request will be forwarded to. Subsequent requests, by the same client to the same load balancer, will be sent to the origin server the cookie encodes, for the duration of the cookie and as long as the origin server remains healthy. If the cookie has expired or the origin server is unhealthy then a new origin server is calculated and used. Value `ip_cookie` behaves the same as `cookie` except the initial origin selection is stable and based on the client's IP address. %s", renderAvailableDocumentationValuesStringSlice(sessionAffinityPolicies)),
		},

		"proxied": {
			Type:          schema.TypeBool,
			Optional:      true,
			Default:       false,
			ConflictsWith: []string{"ttl"},
			Description:   "Whether the hostname gets Cloudflare's origin protection.",
		},

		"enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Enable or disable the load balancer.",
		},

		"ttl": {
			Type:          schema.TypeInt,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{"proxied"}, // this is set to zero regardless of config when proxied=true
			Description:   "Time to live (TTL) of the DNS entry for the IP address returned by this load balancer. This cannot be set for proxied load balancers. Defaults to `30`.",
		},

		"description": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(0, 1024),
			Description:  "Free text description.",
		},

		"steering_policy": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validation.StringInSlice([]string{"off", "geo", "dynamic_latency", "random", "proximity", "least_outstanding_requests", ""}, false),
			Description:  fmt.Sprintf("The method the load balancer uses to determine the route to your origin. Value `off` uses [`default_pool_ids`](#default_pool_ids). Value `geo` uses [`pop_pools`](#pop_pools)/[`country_pools`](#country_pools)/[`region_pools`](#region_pools). For non-proxied requests, the [`country`](#country) for [`country_pools`](#country_pools) is determined by [`location_strategy`](#location_strategy). Value `random` selects a pool randomly. Value `dynamic_latency` uses round trip time to select the closest pool in [`default_pool_ids`](#default_pool_ids) (requires pool health checks). Value `proximity` uses the pools' latitude and longitude to select the closest pool using the Cloudflare PoP location for proxied requests or the location determined by [`location_strategy`](#location_strategy) for non-proxied requests. Value `least_outstanding_requests` selects a pool by taking into consideration [`random_steering`](#random_steering) weights, as well as each pool's number of outstanding requests. Pools with more pending requests are weighted proportionately less relative to others. Value `\"\"` maps to `geo` if you use [`pop_pools`](#pop_pools)/[`country_pools`](#country_pools)/[`region_pools`](#region_pools) otherwise `off`. %s Defaults to `\"\"`.", renderAvailableDocumentationValuesStringSlice([]string{"off", "geo", "dynamic_latency", "random", "proximity", "least_outstanding_requests", `""`})),
		},

		"session_affinity_ttl": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntBetween(0, 604800),
			Description:  "Time, in seconds, until this load balancer's session affinity cookie expires after being created. This parameter is ignored unless a supported session affinity policy is set. The current default of `82800` (23 hours) will be used unless [`session_affinity_ttl`](#session_affinity_ttl) is explicitly set. Once the expiry time has been reached, subsequent requests may get sent to a different origin server. Valid values are between `1800` and `604800`.",
		},

		"session_affinity_attributes": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        loadBalancerSessionAffinityAttributesElem,
			Description: "Configure attributes for session affinity.",
		},

		"adaptive_routing": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        loadBalancerAdaptiveRoutingElem,
			Description: "Controls features that modify the routing of requests to pools and origins in response to dynamic conditions, such as during the interval between active health monitoring requests.",
		},

		"location_strategy": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        loadBalancerLocationStrategyElem,
			Description: "Controls location-based steering for non-proxied requests.",
		},

		"random_steering": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        loadBalancerRandomSteeringElem,
			Description: "Configures pool weights. When [`steering_policy=\"random\"`](#steering_policy), a random pool is selected with probability proportional to pool weights. When [`steering_policy=\"least_outstanding_requests\"`](#steering_policy), pool weights are used to scale each pool's outstanding requests.",
		},

		"rules": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        rulesElem,
			Description: "A list of rules for this load balancer to execute.",
		},

		// nb enterprise only
		"pop_pools": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        loadBalancerPopPoolElem,
			Description: "A set containing mappings of Cloudflare Point-of-Presence (PoP) identifiers to a list of pool IDs (ordered by their failover priority) for the PoP (datacenter). This feature is only available to enterprise customers.",
		},

		"country_pools": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        loadBalancerCountryPoolElem,
			Description: "A set containing mappings of country codes to a list of pool IDs (ordered by their failover priority) for the given country.",
		},

		"region_pools": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        loadBalancerRegionPoolElem,
			Description: "A set containing mappings of region codes to a list of pool IDs (ordered by their failover priority) for the given region.",
		},

		"created_on": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The RFC3339 timestamp of when the load balancer was created.",
		},

		"modified_on": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The RFC3339 timestamp of when the load balancer was last modified.",
		},
	}
}
