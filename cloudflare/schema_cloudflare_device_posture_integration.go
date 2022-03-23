package cloudflare

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareDevicePostureIntegrationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{ws1}, false),
		},
		"identifier": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"interval": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"config": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"auth_url": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.IsURLWithHTTPS,
					},
					"api_url": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.IsURLWithHTTPS,
					},
					"client_id": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"client_secret": {
						Type:      schema.TypeString,
						Optional:  true,
						Sensitive: true,
					},
				},
			},
		},
	}
}
