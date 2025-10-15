// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organization

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*OrganizationResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"profile": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"business_address": schema.StringAttribute{
						Required: true,
					},
					"business_email": schema.StringAttribute{
						Required: true,
					},
					"business_name": schema.StringAttribute{
						Required: true,
					},
					"business_phone": schema.StringAttribute{
						Required: true,
					},
					"external_metadata": schema.StringAttribute{
						Required: true,
					},
				},
			},
			"parent": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[OrganizationParentModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Required: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"create_time": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"meta": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[OrganizationMetaModel](ctx),
				Attributes: map[string]schema.Attribute{
					"flags": schema.SingleNestedAttribute{
						Description: "Organization flags for feature enablement",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[OrganizationMetaFlagsModel](ctx),
						Attributes: map[string]schema.Attribute{
							"account_creation": schema.StringAttribute{
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
							"sub_org_creation": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"managed_by": schema.StringAttribute{
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
