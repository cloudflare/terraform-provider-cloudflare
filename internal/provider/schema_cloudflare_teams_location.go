package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareTeamsLocationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"networks": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"network": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"client_default": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"policy_ids": {
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Computed: true,
		},
		"ip": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"doh_subdomain": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"anonymized_logs_enabled": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"ipv4_destination": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
