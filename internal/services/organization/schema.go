// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organization

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*OrganizationResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"User Details Write",
			},
			MarkdownDescription: "User API tokens can list and create Organizations, but reading, updating, or deleting an Organization requires a Global API key and your registered email address. The `cloudflare_organization` Terraform resource therefore requires a Global API key for its full lifecycle.\n\nSee [Organizations authentication limitations](https://developers.cloudflare.com/fundamentals/organizations/limitations/#api-authentication) for credential compatibility by operation. Each endpoint’s **Security** section lists its supported authentication methods, with **Accepted Permissions** where applicable.",
		}.String(),
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"parent": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Required: true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"profile": schema.SingleNestedAttribute{
				Description: "Optional profile fields, each limited to 1,000 UTF-8 bytes. An empty profile is ignored.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"business_address": schema.StringAttribute{
						Optional: true,
					},
					"business_email": schema.StringAttribute{
						Optional: true,
					},
					"business_name": schema.StringAttribute{
						Optional: true,
					},
					"business_phone": schema.StringAttribute{
						Optional: true,
					},
					"external_metadata": schema.StringAttribute{
						Optional: true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "The Organization name, limited to 65 UTF-8 bytes.",
				Required:    true,
			},
			"parent_id": schema.StringAttribute{
				Description: "The current parent Organization ID. A different value is rejected; this endpoint cannot move an Organization.",
				Optional:    true,
			},
			"create_time": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"meta": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[OrganizationMetaModel](ctx),
				Attributes: map[string]schema.Attribute{
					"are_child_accounts_accessible_by_org_members": schema.BoolAttribute{
						Computed: true,
					},
					"hierarchy_tags": schema.ListAttribute{
						Description: "Ordered chain of organization tags from the root organization down to\n(and including) this organization itself. Root organizations return a\nsingle-element array containing their own tag; sub-organizations return\n`[rootTag, ...intermediateTags, parentTag, selfTag]`. Useful for\nconstructing authorization scopes that need to cover every ancestor\nin the hierarchy.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"managed_by": schema.StringAttribute{
						Computed: true,
					},
					"tenant_flags": schema.SingleNestedAttribute{
						Description: "Enable features for Organizations.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[OrganizationMetaTenantFlagsModel](ctx),
						Attributes: map[string]schema.Attribute{
							"account_creation": schema.StringAttribute{
								Computed: true,
							},
							"account_creation_applies_tenant_defaults": schema.StringAttribute{
								Computed: true,
							},
							"account_deletion": schema.StringAttribute{
								Computed: true,
							},
							"account_migration": schema.StringAttribute{
								Computed: true,
							},
							"account_mobility": schema.StringAttribute{
								Computed: true,
							},
							"enterprise_capability": schema.StringAttribute{
								Computed: true,
							},
							"member_management": schema.StringAttribute{
								Computed: true,
							},
							"sub_org_creation": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"tenant_type": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (r *OrganizationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *OrganizationResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
