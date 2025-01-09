// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_transforms

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ManagedTransformsResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The unique ID of the zone.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "The unique ID of the zone.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"managed_request_headers": schema.ListNestedAttribute{
				Description: "The list of Managed Request Transforms.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The human-readable identifier of the Managed Transform.",
							Required:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the Managed Transform is enabled.",
							Required:    true,
						},
						"has_conflict": schema.BoolAttribute{
							Description: "Whether the Managed Transform conflicts with the currently-enabled Managed Transforms.",
							Computed:    true,
						},
						"conflicts_with": schema.ListAttribute{
							Description: "The Managed Transforms that this Managed Transform conflicts with.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
					},
				},
			},
			"managed_response_headers": schema.ListNestedAttribute{
				Description: "The list of Managed Response Transforms.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The human-readable identifier of the Managed Transform.",
							Required:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the Managed Transform is enabled.",
							Required:    true,
						},
						"has_conflict": schema.BoolAttribute{
							Description: "Whether the Managed Transform conflicts with the currently-enabled Managed Transforms.",
							Computed:    true,
						},
						"conflicts_with": schema.ListAttribute{
							Description: "The Managed Transforms that this Managed Transform conflicts with.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
					},
				},
			},
		},
	}
}

func (r *ManagedTransformsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ManagedTransformsResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
