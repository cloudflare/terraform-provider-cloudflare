package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareEmailRoutingCatchAllSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"tag": {
			Description: "Routing rule identifier.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"name": {
			Description: "Routing rule name.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"enabled": {
			Description: "Routing rule status.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"matcher": {
			Description: "Matching patterns to forward to your actions.",
			Type:        schema.TypeSet,
			Required:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"type": {
						Description:  "Type of matcher.",
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringInSlice([]string{"all"}, true),
					},
				},
			},
		},
		"action": {
			Description: "List actions patterns.",
			Type:        schema.TypeSet,
			Required:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"type": {
						Description:  "Type of supported action.",
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringInSlice([]string{"drop", "forward", "worker"}, true),
					},
					"value": {
						Description: "An array with items in the following form.",
						Type:        schema.TypeList,
						Required:    true,
						Elem: &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 90),
						},
					},
				},
			},
		},
	}
}
