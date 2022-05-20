package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareApiTokenSchema() map[string]*schema.Schema {
	p := schema.Resource{
		Schema: map[string]*schema.Schema{
			"resources": {
				Type:     schema.TypeMap,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"permission_groups": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"effect": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "allow",
				ValidateFunc: validation.StringInSlice([]string{"allow", "deny"}, false),
			},
		},
	}

	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"policy": {
			Type:     schema.TypeSet,
			Set:      schema.HashResource(&p),
			Required: true,
			Elem:     &p,
		},
		"condition": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"request_ip": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"in": {
									Type:     schema.TypeSet,
									Elem:     &schema.Schema{Type: schema.TypeString},
									Optional: true,
								},
								"not_in": {
									Type:     schema.TypeSet,
									Elem:     &schema.Schema{Type: schema.TypeString},
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
		"value": {
			Type:      schema.TypeString,
			Computed:  true,
			Sensitive: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"issued_on": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"modified_on": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
