package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareIPsecTunnelSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"customer_endpoint": {
			Type:     schema.TypeString,
			Required: true,
		},
		"cloudflare_endpoint": {
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
		"psk": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"hex_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"user_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"fqdn_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"remote_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}
