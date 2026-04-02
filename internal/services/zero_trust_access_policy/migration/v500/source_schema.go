package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceAccessPolicySchema returns the v4 cloudflare_access_policy schema.
// This is used by MoveState to parse the source state from v4 provider.
func SourceAccessPolicySchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"zone_id": schema.StringAttribute{
				Optional: true,
			},
			"application_id": schema.StringAttribute{
				Optional: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"precedence": schema.Int64Attribute{
				Optional: true,
			},
			"decision": schema.StringAttribute{
				Required: true,
			},
			"session_duration": schema.StringAttribute{
				Optional: true,
			},
			"isolation_required": schema.BoolAttribute{
				Optional: true,
			},
			"purpose_justification_required": schema.BoolAttribute{
				Optional: true,
			},
			"purpose_justification_prompt": schema.StringAttribute{
				Optional: true,
			},
			"approval_required": schema.BoolAttribute{
				Optional: true,
			},
		},
		Blocks: map[string]schema.Block{
			// v4 include/exclude/require are list blocks
			"include": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: sourceConditionGroupAttributes(),
					Blocks:     sourceConditionGroupBlocks(),
				},
			},
			"exclude": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: sourceConditionGroupAttributes(),
					Blocks:     sourceConditionGroupBlocks(),
				},
			},
			"require": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: sourceConditionGroupAttributes(),
					Blocks:     sourceConditionGroupBlocks(),
				},
			},
			// v4 approval_group is a list block
			"approval_group": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"approvals_needed": schema.Int64Attribute{
							Required: true,
						},
						"email_addresses": schema.ListAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"email_list_uuid": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			// v4 connection_rules is a list block with MaxItems:1
			"connection_rules": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Blocks: map[string]schema.Block{
						"ssh": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"usernames": schema.ListAttribute{
										ElementType: types.StringType,
										Required:    true,
									},
									"allow_email_alias": schema.BoolAttribute{
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// sourceConditionGroupAttributes returns the attributes for condition groups
func sourceConditionGroupAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"everyone": schema.BoolAttribute{
			Optional: true,
		},
		"certificate": schema.BoolAttribute{
			Optional: true,
		},
		"any_valid_service_token": schema.BoolAttribute{
			Optional: true,
		},
		"email": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"email_domain": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"ip": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"group": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"geo": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"login_method": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"common_name": schema.StringAttribute{
			Optional: true,
		},
		"common_names": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"auth_method": schema.StringAttribute{
			Optional: true,
		},
		// v4 uses simple lists for these
		"device_posture": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"email_list": schema.ListAttribute{
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
	}
}

// sourceConditionGroupBlocks returns the nested blocks for condition groups
func sourceConditionGroupBlocks() map[string]schema.Block {
	return map[string]schema.Block{
		"saml": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
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
		"oidc": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"claim_name": schema.StringAttribute{
						Optional: true,
					},
					"claim_value": schema.StringAttribute{
						Optional: true,
					},
					"identity_provider_id": schema.StringAttribute{
						Optional: true,
					},
				},
			},
		},
		"azure": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
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
		"okta": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
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
		"gsuite": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"email": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"identity_provider_id": schema.StringAttribute{
						Optional: true,
					},
				},
			},
		},
		"github": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Optional: true,
					},
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
		"external_evaluation": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
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
		"auth_context": schema.ListNestedBlock{
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Required: true,
					},
					"ac_id": schema.StringAttribute{
						Required: true,
					},
					"identity_provider_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
	}
}
