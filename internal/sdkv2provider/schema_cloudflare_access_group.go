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
			Type:        schema.TypeList,
			Description: "The email of the user.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"email_domain": {
			Type:        schema.TypeList,
			Description: "The email domain to match.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"email_list": {
			Type:        schema.TypeList,
			Description: "The ID of a previously created email list.",
			Optional:    true,
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
			Description: "The ID of a previously created IP list.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"service_token": {
			Type:        schema.TypeList,
			Description: "The ID of an Access service token.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"any_valid_service_token": {
			Type:        schema.TypeBool,
			Description: "Matches any valid Access service token.",
			Optional:    true,
		},
		"group": {
			Type:        schema.TypeList,
			Description: "The ID of a previously created Access group.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"everyone": {
			Type:        schema.TypeBool,
			Description: "Matches everyone.",
			Optional:    true,
		},
		"certificate": {
			Type:        schema.TypeBool,
			Description: "Matches any valid client certificate.",
			Optional:    true,
		},
		"common_name": {
			Type:        schema.TypeString,
			Description: "Matches a valid client certificate common name.",
			Optional:    true,
		},
		"auth_method": {
			Type:        schema.TypeString,
			Description: "The type of authentication method. Refer to https://datatracker.ietf.org/doc/html/rfc8176#section-2 for possible types.",
			Optional:    true,
		},
		"geo": {
			Type:        schema.TypeList,
			Description: "Matches a specific country.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"login_method": {
			Type:        schema.TypeList,
			Description: "The ID of a configured identity provider.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"device_posture": {
			Type:        schema.TypeList,
			Description: "The ID of a device posture integration.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"gsuite": {
			Type:        schema.TypeList,
			Description: "Matches a group in Google Workspace. Requires a Google Workspace identity provider.",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"email": {
						Type:        schema.TypeList,
						Description: "The email of the Google Workspace group.",
						Required:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"identity_provider_id": {
						Type:        schema.TypeString,
						Description: "The ID of your Google Workspace identity provider.",
						Required:    true,
					},
				},
			},
		},
		"github": {
			Type:        schema.TypeList,
			Description: "Matches a Github organization. Requires a Github identity provider.",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:        schema.TypeString,
						Description: "The name of the organization.",
						Optional:    true,
					},
					"teams": {
						Type:        schema.TypeList,
						Description: "The teams that should be matched.",
						Optional:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"identity_provider_id": {
						Type:        schema.TypeString,
						Description: "The ID of your Github identity provider.",
						Optional:    true,
					},
				},
			},
		},
		"azure": {
			Type:        schema.TypeList,
			Description: "Matches an Azure group. Requires an Azure identity provider.",
			Optional:    true,
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
						Description: "The ID of the Azure identity provider",
						Optional:    true,
					},
				},
			},
		},
		"okta": {
			Type:        schema.TypeList,
			Description: "Matches an Okta group. Requires an Okta identity provider.",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:        schema.TypeList,
						Description: "The name of the Okta Group",
						Optional:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"identity_provider_id": {
						Type:        schema.TypeString,
						Description: "The ID of your Okta identity provider.",
						Optional:    true,
					},
				},
			},
		},
		"saml": {
			Type:        schema.TypeList,
			Description: "Matches a SAML group. Requires a SAML identity provider.",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"attribute_name": {
						Type:        schema.TypeString,
						Description: "The name of the SAML attribute.",
						Optional:    true,
					},
					"attribute_value": {
						Type:        schema.TypeString,
						Description: "The SAML attribute value to look for.",
						Optional:    true,
					},
					"identity_provider_id": {
						Type:        schema.TypeString,
						Description: "The ID of your SAML identity provider.",
						Optional:    true,
					},
				},
			},
		},
		"external_evaluation": {
			Type:        schema.TypeList,
			Description: "Create Allow or Block policies which evaluate the user based on custom criteria. https://developers.cloudflare.com/cloudflare-one/policies/access/external-evaluation/",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"evaluate_url": {
						Type:        schema.TypeString,
						Description: "The API endpoint containing your business logic.",
						Optional:    true,
					},
					"keys_url": {
						Type:        schema.TypeString,
						Description: "The API endpoint containing the key that Access uses to verify that the response came from your API.",
						Optional:    true,
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
						Description: "The ID of the Azure identity provider",
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
