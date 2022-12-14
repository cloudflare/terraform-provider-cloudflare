package provider

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareLoadBalancerPoolSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringMatch(regexp.MustCompile("[-_a-zA-Z0-9]+"), "Only alphanumeric characters, hyphens and underscores are allowed."),
		},

		"origins": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     originsElem,
		},

		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},

		"minimum_origins": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},

		"latitude": {
			Type:         schema.TypeFloat,
			Optional:     true,
			ValidateFunc: validation.FloatBetween(-90, 90),
		},

		"longitude": {
			Type:         schema.TypeFloat,
			Optional:     true,
			ValidateFunc: validation.FloatBetween(-180, 180),
		},

		"check_regions": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},

		"description": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(0, 1024),
		},

		"monitor": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(0, 32),
		},

		"notification_email": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"load_shedding": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     loadShedElem,
		},

		"origin_steering": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     originSteeringElem,
		},

		"created_on": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"modified_on": {
			Type:     schema.TypeString,
			Computed: true,
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
			Description:  "The weight (0.01 - 1.00) of this origin, relative to other origins in the pool. Equal values mean equal weighting. A weight of 0 means traffic will not be sent to this origin, but health is still checked.",
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
			ValidateFunc: validation.StringInSlice([]string{"", "hash", "random"}, false),
			Description:  fmt.Sprintf("Origin steering policy to be used. %s", renderAvailableDocumentationValuesStringSlice([]string{"", "hash", "random"})),
		},
	},
}
