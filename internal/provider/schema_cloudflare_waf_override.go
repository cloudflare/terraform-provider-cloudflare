package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareWAFOverrideSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"urls": {
			Required: true,
			Type:     schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"rules": {
			Optional: true,
			Type:     schema.TypeMap,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"paused": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"priority": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntBetween(-1000000000, 1000000000),
		},
		"groups": {
			Optional: true,
			Type:     schema.TypeMap,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"rewrite_action": {
			Optional: true,
			Type:     schema.TypeMap,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"override_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
