package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareEmailRoutingRuleSchema() map[string]*schema.Schema {
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
		"priority": {
			Description:  "Priority of the routing rule.",
			Type:         schema.TypeInt,
			ValidateFunc: validation.IntAtLeast(0),
			Optional:     true,
		},
		"enabled": {
			Description: "Routing rule status.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"matchers": {
			Description: "Matching patterns to forward to your actions.",
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"type": {
						Description:  "Type of matcher.",
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringInSlice([]string{"literal"}, true),
					},
					"field": {
						Description: "Field for type matcher.",
						Type:        schema.TypeString,
						Required:    true,
					},
					"value": {
						Description:  "Value for matcher.",
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringLenBetween(0, 90),
					},
				},
			},
		},
		"actions": {
			Description: "List actions patterns.",
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"type": {
						Description:  "Type of supported action.",
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringInSlice([]string{"forward", "worker"}, true),
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
