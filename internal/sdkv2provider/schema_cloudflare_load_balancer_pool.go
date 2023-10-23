package sdkv2provider

import (
	"fmt"
	"regexp"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareLoadBalancerPoolSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},

		"name": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringMatch(regexp.MustCompile("[-_a-zA-Z0-9]+"), "Only alphanumeric characters, hyphens and underscores are allowed."),
			Description:  "A short name (tag) for the pool.",
		},

		"origins": {
			Type:        schema.TypeSet,
			Required:    true,
			Elem:        originsElem,
			Description: "The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy origins, provided the pool itself is healthy.",
		},

		"enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Whether to enable (the default) this pool. Disabled pools will not receive traffic and are excluded from health checks. Disabling a pool will cause any load balancers using it to failover to the next pool (if any).",
		},

		"minimum_origins": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     1,
			Description: "The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins falls below this number, the pool will be marked unhealthy and we will failover to the next available pool.",
		},

		"latitude": {
			Type:         schema.TypeFloat,
			Optional:     true,
			ValidateFunc: validation.FloatBetween(-90, 90),
			Description:  "The latitude this pool is physically located at; used for proximity steering.",
		},

		"longitude": {
			Type:         schema.TypeFloat,
			Optional:     true,
			ValidateFunc: validation.FloatBetween(-180, 180),
			Description:  "The longitude this pool is physically located at; used for proximity steering.",
		},

		"check_regions": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "A list of regions (specified by region code) from which to run health checks. Empty means every Cloudflare data center (the default), but requires an Enterprise plan. Region codes can be found [here](https://developers.cloudflare.com/load-balancing/reference/region-mapping-api).",
		},

		"description": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(0, 1024),
			Description:  "Free text description.",
		},

		"monitor": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(0, 32),
			Description:  "The ID of the Monitor to use for health checking origins within this pool.",
		},

		"notification_email": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The email address to send health status notifications to. This can be an individual mailbox or a mailing list. Multiple emails can be supplied as a comma delimited list.",
		},

		"load_shedding": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        loadShedElem,
			Description: "Setting for controlling load shedding for this pool.",
		},

		"origin_steering": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        originSteeringElem,
			Description: "Set an origin steering policy to control origin selection within a pool.",
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

var originsElem = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "A human-identifiable name for the origin.",
		},

		"address": {
			Type:     schema.TypeString,
			Required: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validateStringIP,
			},
			Description: "The IP address (IPv4 or IPv6) of the origin, or the publicly addressable hostname.",
		},

		"weight": {
			Type:         schema.TypeFloat,
			Optional:     true,
			Default:      1.0,
			ValidateFunc: validation.FloatBetween(0.0, 1.0),
			Description:  "The weight (0.01 - 1.00) of this origin, relative to other origins in the pool. Equal values mean equal weighting. A weight of 0 means traffic will not be sent to this origin, but health is still checked. When [`origin_steering.policy=\"least_outstanding_requests\"`](#policy), weight is used to scale the origin's outstanding requests. When [`origin_steering.policy=\"least_connections\"`](#policy), weight is used to scale the origin's open connections.",
		},

		"enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Whether this origin is enabled. Disabled origins will not receive traffic and are excluded from health checks.",
		},

		"header": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "HTTP request headers.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"header": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "HTTP Header name.",
					},
					"values": {
						Type:     schema.TypeSet,
						Required: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Description: "Values for the HTTP headers.",
					},
				},
			},
			Set: HashByMapKey("header"),
		},
	},
}

var loadShedElem = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"default_percent": {
			Type:         schema.TypeFloat,
			Default:      0,
			Optional:     true,
			ValidateFunc: validation.FloatBetween(0, 100),
			Description:  "Percent of traffic to shed 0 - 100.",
		},

		"default_policy": {
			Type:         schema.TypeString,
			Default:      "",
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"", "hash", "random"}, false),
			Description:  fmt.Sprintf("Method of shedding traffic. %s", renderAvailableDocumentationValuesStringSlice([]string{"", "hash", "random"})),
		},

		"session_percent": {
			Type:         schema.TypeFloat,
			Default:      0,
			Optional:     true,
			ValidateFunc: validation.FloatBetween(0, 100),
			Description:  "Percent of session traffic to shed 0 - 100.",
		},

		"session_policy": {
			Type:         schema.TypeString,
			Default:      "",
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"", "hash"}, false),
			Description:  fmt.Sprintf("Method of shedding traffic. %s", renderAvailableDocumentationValuesStringSlice([]string{"", "hash"})),
		},
	},
}

var originSteeringElem = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"policy": {
			Type:         schema.TypeString,
			Default:      "random",
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"", "hash", "random", "least_outstanding_requests", "least_connections"}, false),
			Description:  fmt.Sprintf("Origin steering policy to be used. Value `random` selects an origin randomly. Value `hash` selects an origin by computing a hash over the CF-Connecting-IP address. Value `least_outstanding_requests` selects an origin by taking into consideration origin weights, as well as each origin's number of outstanding requests. Origins with more pending requests are weighted proportionately less relative to others. Value `least_connections` selects an origin by taking into consideration origin weights, as well as each origin's number of open connections. Origins with more open connections are weighted proportionately less relative to others. Supported for HTTP/1 and HTTP/2 connections. %s", renderAvailableDocumentationValuesStringSlice([]string{"", "hash", "random", "least_outstanding_requests", "least_connections"})),
		},
	},
}
