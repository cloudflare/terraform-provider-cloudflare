package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAccessGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description:   consts.AccountIDSchemaDescription,
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{consts.ZoneIDSchemaKey},
		},
		consts.ZoneIDSchemaKey: {
			Description:   consts.ZoneIDSchemaDescription,
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
		"email_list": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"ip": {
			Type:        schema.TypeList,
			Description: "An IPv4 or IPv6 CIDR block.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"ip_list": {
			Type:        schema.TypeList,
			Description: "The ID of an existing IP list to reference.",
			Optional:    true,
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
						Type:        schema.TypeList,
						Description: "The ID of the Azure group or user.",
						Optional:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"identity_provider_id": {
						Type:        schema.TypeString,
						Description: "The ID of the Azure Identity provider",
						Optional:    true,
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
		"auth_context": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Description: "The ID of the Authentication Context.",
						Required:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"identity_provider_id": {
						Type:        schema.TypeString,
						Description: "The ID of the Azure Identity provider",
						Required:    true,
					},
					"ac_id": {
						Type:        schema.TypeString,
						Description: "The ACID of the Authentication Context",
						Required:    true,
					},
				},
			},
		},
		"common_names": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Overflow field if you need to have multiple common_name rules in a single policy.  Use in place of the singular common_name field.",
		},
	},
}
