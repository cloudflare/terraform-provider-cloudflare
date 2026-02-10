package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceV4ZeroTrustAccessGroupSchema returns the schema for v4 zero_trust_access_group resources (SDKv2 format).
// This schema represents the state structure from the v4 provider (schema_version=0).
//
// In v4, this resource was named cloudflare_access_group.
// In v5, it's renamed to cloudflare_zero_trust_access_group.
//
// Major structural difference:
// - v4: List fields contain multiple string values (e.g., email = ["a@b.com", "c@d.com"])
// - v5: Each value becomes a separate object (e.g., two separate email blocks)
func SourceV4ZeroTrustAccessGroupSchema() schema.Schema {
	return schema.Schema{
		Version: 0,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"zone_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},

			// In v4 SDKv2, include/exclude/require are TypeList with complex nested attributes
			"include": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: accessGroupOptionAttributes(),
				},
			},
			"exclude": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: accessGroupOptionAttributes(),
				},
			},
			"require": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: accessGroupOptionAttributes(),
				},
			},
		},
	}
}

// accessGroupOptionAttributes defines the selector attributes within include/exclude/require blocks.
// In v4, most simple selectors are TypeList of strings (arrays).
// Complex selectors (github, gsuite, etc.) are TypeList MaxItems:1 (array with single object).
func accessGroupOptionAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// Simple list selectors (stored as arrays of strings in v4)
		"email": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"email_domain": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"email_list": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"ip": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"ip_list": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"service_token": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"group": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"device_posture": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"login_method": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"geo": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},

		// String scalar selectors
		"common_name": schema.StringAttribute{
			Optional: true,
		},
		"auth_method": schema.StringAttribute{
			Optional: true,
		},

		// Boolean selectors (converted to empty objects in v5)
		"everyone": schema.BoolAttribute{
			Optional: true,
		},
		"certificate": schema.BoolAttribute{
			Optional: true,
		},
		"any_valid_service_token": schema.BoolAttribute{
			Optional: true,
		},

		// Special case: common_names overflow array (removed in v5)
		"common_names": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},

		// Complex nested object selectors (TypeList MaxItems:1 in v4)
		// These are stored as arrays with single element in v4 state

		// github → renamed to github_organization in v5
		"github": schema.ListNestedAttribute{
			Optional: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Optional: true,
					},
					// teams (plural) → becomes team (singular) in v5
					"teams": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"identity_provider_id": schema.StringAttribute{
						Optional: true,
					},
				},
			},
		},

		// gsuite
		"gsuite": schema.ListNestedAttribute{
			Optional: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"email": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"identity_provider_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},

		// azure → renamed to azure_ad in v5
		"azure": schema.ListNestedAttribute{
			Optional: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"id": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"identity_provider_id": schema.StringAttribute{
						Optional: true,
					},
				},
			},
		},

		// okta
		"okta": schema.ListNestedAttribute{
			Optional: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"identity_provider_id": schema.StringAttribute{
						Optional: true,
					},
				},
			},
		},

		// saml
		"saml": schema.ListNestedAttribute{
			Optional: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"attribute_name": schema.StringAttribute{
						Optional: true,
					},
					"attribute_value": schema.StringAttribute{
						Optional: true,
					},
					"identity_provider_id": schema.StringAttribute{
						Optional: true,
					},
				},
			},
		},

		// external_evaluation
		"external_evaluation": schema.ListNestedAttribute{
			Optional: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"evaluate_url": schema.StringAttribute{
						Optional: true,
					},
					"keys_url": schema.StringAttribute{
						Optional: true,
					},
				},
			},
		},

		// auth_context
		"auth_context": schema.ListNestedAttribute{
			Optional: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Required: true,
					},
					"identity_provider_id": schema.StringAttribute{
						Required: true,
					},
					"ac_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
	}
}
