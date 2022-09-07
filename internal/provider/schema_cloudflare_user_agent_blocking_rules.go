package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareUserAgentBlockingRulesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"id": {
			Description: "The unique identifier of the User Agent Blocking rule.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"paused": {
			Description: "When true, indicates that the rule is currently paused.",
			Type:        schema.TypeBool,
			Required:    true,
		},
		"description": {
			Description:  "An informative summary of the rule.\n\n",
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringLenBetween(0, 1024),
		},
		"mode": {
			Description:  "The action to apply to a matched request.",
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"block", "challenge", "js_challenge", "managed_challenge"}, false),
		},
		"configuration": {
			Description: "The configuration object for the current rule.",
			Required:    true,
			Type:        schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"target": {
						Description:  "The configuration target for this rule. You must set the target to ua for User Agent Blocking rules.",
						Required:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"ua"}, false),
					},
					"value": {
						Description: "The exact user agent string to match. This value will be compared to the received User-Agent HTTP header value.",
						Required:    true,
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}
