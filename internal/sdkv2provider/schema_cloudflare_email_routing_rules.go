package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareEmailRoutingRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
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
			Computed:     true,
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
						ValidateFunc: validation.StringInSlice([]string{"literal", "all"}, true),
					},
					"field": {
						Description: "Field for type matcher.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"value": {
						Description:  "Value for matcher.",
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringLenBetween(0, 90),
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
