package provider

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"name": {
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringMatch(regexp.MustCompile("^[0-9a-z_]+$"), "List name must only contain lowercase letters, numbers and underscores"),
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"kind": {
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice([]string{"ip"}, false),
			Required:     true,
		},
		"item": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     listItemElem,
		},
	}
}

var listItemElem = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"value": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 1,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"ip": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"comment": {
			Type:     schema.TypeString,
			Optional: true,
		},
	},
}
