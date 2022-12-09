package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareCustomSslSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"custom_ssl_priority": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"priority": {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		},
		"custom_ssl_options": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			ForceNew: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"certificate": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"private_key": {
						Type:      schema.TypeString,
						Optional:  true,
						Sensitive: true,
					},
					"bundle_method": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"ubiquitous", "optimal", "force"}, false),
					},
					"geo_restrictions": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"us", "eu", "highest_security"}, false),
					},
					"type": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"legacy_custom", "sni_custom"}, false),
					},
				},
			},
		},
		"hosts": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"issuer": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"signature": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"uploaded_on": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"modified_on": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"expires_on": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"priority": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}
