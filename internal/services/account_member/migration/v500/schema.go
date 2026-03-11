package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceV4Schema returns the v4 SDKv2 account_member schema for state parsing.
//
// Reference: cloudflare-terraform-v4/internal/sdkv2provider/schema_cloudflare_account_member.go
//
// Key differences from v5:
//   - email_address (v4) instead of email (v5)
//   - role_ids (v4) instead of roles (v5)
//   - No policies field (v5 only)
//   - No user field (v5 only)
func SourceV4Schema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"email_address": schema.StringAttribute{
				Required: true,
			},
			"role_ids": schema.SetAttribute{
				Required:    true,
				ElementType: types.StringType,
			},
			"status": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
		},
	}
}

// SourceV513Schema returns the v5.13.0 account_member schema for state parsing.
//
// This is needed because v5.13 had a different schema than current v5:
//   - policies was ListNestedAttribute (now SetNestedAttribute)
//   - policies had an 'id' field (now removed)
//   - permission_groups was ListNestedAttribute (now SetNestedAttribute)
//   - resource_groups was ListNestedAttribute (now SetNestedAttribute)
//   - roles was ListAttribute (now SetAttribute)
func SourceV513Schema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"email": schema.StringAttribute{
				Required: true,
			},
			"status": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			// v5.13 used ListAttribute for roles
			"roles": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			// v5.13 used ListNestedAttribute for policies with an 'id' field
			"policies": schema.ListNestedAttribute{
				Computed: true,
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"access": schema.StringAttribute{
							Required: true,
						},
						"permission_groups": schema.ListNestedAttribute{
							Required: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Required: true,
									},
								},
							},
						},
						"resource_groups": schema.ListNestedAttribute{
							Required: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Required: true,
									},
								},
							},
						},
					},
				},
			},
			"user": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"email": schema.StringAttribute{
						Computed: true,
					},
					"id": schema.StringAttribute{
						Computed: true,
					},
					"first_name": schema.StringAttribute{
						Computed: true,
					},
					"last_name": schema.StringAttribute{
						Computed: true,
					},
					"two_factor_authentication_enabled": schema.BoolAttribute{
						Computed: true,
						Default:  booldefault.StaticBool(false),
					},
				},
			},
		},
	}
}
