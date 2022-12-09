package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareGRETunnelSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"customer_gre_endpoint": {
			Type:     schema.TypeString,
			Required: true,
		},
		"cloudflare_gre_endpoint": {
			Type:     schema.TypeString,
			Required: true,
		},
		"interface_address": {
			Type:     schema.TypeString,
			Required: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"ttl": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"mtu": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"health_check_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"health_check_target": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"health_check_type": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validation.StringInSlice([]string{"request", "reply"}, false),
		},
	}
}
