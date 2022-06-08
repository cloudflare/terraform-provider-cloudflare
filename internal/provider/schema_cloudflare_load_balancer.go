package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareLoadBalancerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"fallback_pool_id": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringLenBetween(1, 32),
		},

		"default_pool_ids": {
			Type:     schema.TypeList,
			Required: true,
			MinItems: 1,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(1, 32),
			},
		},

		"session_affinity": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "none",
			ValidateFunc: validation.StringInSlice([]string{"", "none", "cookie", "ip_cookie"}, false),
		},

		"proxied": {
			Type:          schema.TypeBool,
			Optional:      true,
			Default:       false,
			ConflictsWith: []string{"ttl"},
		},

		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},

		"ttl": {
			Type:          schema.TypeInt,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{"proxied"}, // this is set to zero regardless of config when proxied=true
		},

		"description": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(0, 1024),
		},

		"steering_policy": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"off", "geo", "dynamic_latency", "random", "proximity", ""}, false),
			Computed:     true,
		},

		"session_affinity_ttl": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      nil,
			ValidateFunc: validation.IntBetween(1800, 604800),
		},

		"session_affinity_attributes": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},

		"rules": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     rulesElem,
		},

		// nb enterprise only
		"pop_pools": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     popPoolElem,
		},

		"region_pools": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     regionPoolElem,
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
