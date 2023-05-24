package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareStaticRouteSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Description of the static route.",
		},
		"prefix": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Your network prefix using CIDR notation.",
		},
		"nexthop": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The nexthop IP address where traffic will be routed to.",
		},
		"priority": {
			Type:        schema.TypeInt,
			Required:    true,
			Description: "The priority for the static route.",
		},
		"weight": {
			Type:     schema.TypeInt,
			Optional: true,
			// API does not allow to reset weights when attribute isn't send. To avoid generating unnecessary changes
			// we will trigger a re-create when weights change
			ForceNew:    true,
			Description: "The optional weight for ECMP routes.",
		},
		"colo_regions": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "List of Cloudflare colocation names for this static route.",
		},
		"colo_names": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "List of Cloudflare colocation regions for this static route.",
		},
	}
}
