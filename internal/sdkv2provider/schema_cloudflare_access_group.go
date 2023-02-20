package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAccessGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description:   "The account identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{consts.ZoneIDSchemaKey},
		},
		consts.ZoneIDSchemaKey: {
			Description:   "The zone identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{consts.AccountIDSchemaKey},
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"require": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     AccessGroupOptionSchemaElement,
		},
		"exclude": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     AccessGroupOptionSchemaElement,
		},
		"include": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     AccessGroupOptionSchemaElement,
		},
	}
}

// AccessGroupOptionSchemaElement is used by `require`, `exclude` and `include`
// attributes to build out the expected access conditions.
var AccessGroupOptionSchemaElement = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"email": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"email_domain": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"ip": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"ip_list": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"service_token": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"any_valid_service_token": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"group": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"everyone": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"certificate": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"common_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"auth_method": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"geo": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"login_method": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"device_posture": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"gsuite": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"email": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"identity_provider_id": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"github": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"teams": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"identity_provider_id": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"azure": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"identity_provider_id": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"okta": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"identity_provider_id": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"saml": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"attribute_name": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"attribute_value": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"identity_provider_id": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"external_evaluation": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"evaluate_url": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"keys_url": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
	},
}
