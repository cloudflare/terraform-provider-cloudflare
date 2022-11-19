package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceChallengeWidgetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"site_key": {
			Description: "Widget item identifier tag.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"secret": {
			Description: "Secret key for this widget.",
			Type:        schema.TypeString,
			Sensitive:   true,
			Computed:    true,
		},
		"name": {
			Description: "Human readable widget name. Not unique. Cloudflare suggests that you set this to a meaningful string to make it easier to identify your widget, and where it is used.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"domains": {
			Description:  "Hosts as a hostname or IPv4/IPv6 address represented by strings. The widget will only work on these domains, and their subdomains.",
			Type:         schema.TypeSet,
			Elem:         &schema.Schema{Type: schema.TypeString},
			Optional:     true,
		},
		"type": {
			Description:  "Type of widget.",
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice([]string{"non_interactive", "invisible", "managed"}, false),
			Optional:     true,
		},
	}
}
