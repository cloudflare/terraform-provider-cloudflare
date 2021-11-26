package cloudflare

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareLoadBalancerPoolSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
			Type:     schema.TypeString,
			Required: true,
		},

		"address": {
			Type:     schema.TypeString,
			Required: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validateStringIP,
			},
		},

		"weight": {
			Type:         schema.TypeFloat,
			Optional:     true,
			Default:      1.0,
			ValidateFunc: validation.FloatBetween(0.0, 1.0),
		},

		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},

		"header": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"header": {
						Type:     schema.TypeString,
						Required: true,
					},
					"values": {
						Type:     schema.TypeSet,
						Required: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
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
		},

		"default_policy": {
			Type:         schema.TypeString,
			Default:      "",
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"", "hash", "random"}, false),
		},

		"session_percent": {
			Type:         schema.TypeFloat,
			Default:      0,
			Optional:     true,
			ValidateFunc: validation.FloatBetween(0, 100),
		},

		"session_policy": {
			Type:         schema.TypeString,
			Default:      "",
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"", "hash"}, false),
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
		},
	},
}
