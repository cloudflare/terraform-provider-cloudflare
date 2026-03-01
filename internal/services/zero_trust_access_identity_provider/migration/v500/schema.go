package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceAccessIdentityProviderSchema returns the legacy cloudflare_access_identity_provider schema (schema_version=0).
// This is used by MoveState and UpgradeFromV4 to parse state from the legacy SDKv2 provider.
// Only Required/Optional/Computed/ElementType are included — no validators, descriptions, or plan modifiers.
func SourceAccessIdentityProviderSchema() schema.Schema {
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
			"name": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Required: true,
			},
		},
		Blocks: map[string]schema.Block{
			// v4 config is a TypeList block with MaxItems:1
			"config": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"claims": schema.ListAttribute{
							Optional:    true,
							ElementType: types.StringType,
						},
						"client_id": schema.StringAttribute{
							Optional: true,
						},
						"client_secret": schema.StringAttribute{
							Optional: true,
						},
						"conditional_access_enabled": schema.BoolAttribute{
							Optional: true,
						},
						"directory_id": schema.StringAttribute{
							Optional: true,
						},
						"email_claim_name": schema.StringAttribute{
							Optional: true,
						},
						"prompt": schema.StringAttribute{
							Optional: true,
						},
						"support_groups": schema.BoolAttribute{
							Optional: true,
						},
						"centrify_account": schema.StringAttribute{
							Optional: true,
						},
						"centrify_app_id": schema.StringAttribute{
							Optional: true,
						},
						"apps_domain": schema.StringAttribute{
							Optional: true,
						},
						"auth_url": schema.StringAttribute{
							Optional: true,
						},
						"certs_url": schema.StringAttribute{
							Optional: true,
						},
						"pkce_enabled": schema.BoolAttribute{
							Optional: true,
						},
						"scopes": schema.ListAttribute{
							Optional:    true,
							ElementType: types.StringType,
						},
						"token_url": schema.StringAttribute{
							Optional: true,
						},
						"authorization_server_id": schema.StringAttribute{
							Optional: true,
						},
						"okta_account": schema.StringAttribute{
							Optional: true,
						},
						"onelogin_account": schema.StringAttribute{
							Optional: true,
						},
						"ping_env_id": schema.StringAttribute{
							Optional: true,
						},
						"attributes": schema.ListAttribute{
							Optional:    true,
							ElementType: types.StringType,
						},
						"email_attribute_name": schema.StringAttribute{
							Optional: true,
						},
						"header_attributes": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"attribute_name": schema.StringAttribute{
										Optional: true,
									},
									"header_name": schema.StringAttribute{
										Optional: true,
									},
								},
							},
						},
						"issuer_url": schema.StringAttribute{
							Optional: true,
						},
						"sign_request": schema.BoolAttribute{
							Optional: true,
						},
						"sso_target_url": schema.StringAttribute{
							Optional: true,
						},
						"redirect_url": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
						// Deprecated fields (removed in v5)
						"api_token": schema.StringAttribute{
							Optional: true,
						},
						"idp_public_cert": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			// v4 scim_config is a TypeList block with MaxItems:1
			"scim_config": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"enabled": schema.BoolAttribute{
							Optional: true,
							Computed: true,
						},
						"identity_update_behavior": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
						"scim_base_url": schema.StringAttribute{
							Computed: true,
						},
						"seat_deprovision": schema.BoolAttribute{
							Optional: true,
							Computed: true,
						},
						"secret": schema.StringAttribute{
							Computed: true,
						},
						"user_deprovision": schema.BoolAttribute{
							Optional: true,
							Computed: true,
						},
						// Deprecated field (removed in v5)
						"group_member_deprovision": schema.BoolAttribute{
							Optional: true,
						},
					},
				},
			},
		},
	}
}
