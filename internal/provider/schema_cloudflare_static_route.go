package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareStaticRouteSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"prefix": {
			Type:     schema.TypeString,
			Required: true,
		},
		"nexthop": {
			Type:     schema.TypeString,
			Required: true,
		},
		"priority": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"weight": {
			Type:     schema.TypeInt,
			Optional: true,
			// API does not allow to reset weights when attribute isn't send. To avoid generating unnecessary changes
			// we will trigger a re-create when weights change
			ForceNew: true,
		},
		"colo_regions": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"colo_names": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}
