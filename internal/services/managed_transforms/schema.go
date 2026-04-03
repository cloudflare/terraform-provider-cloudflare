// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_transforms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*ManagedTransformsResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The unique ID of the zone.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "The unique ID of the zone.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"managed_request_headers": schema.SetNestedAttribute{
				Description: "The list of Managed Request Transforms.",
				Required:    true,
				CustomType:  customfield.NewNestedObjectSetType[ManagedTransformsManagedRequestHeadersModel](ctx),
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
					},
				},
			},
			"managed_response_headers": schema.SetNestedAttribute{
				Description: "The list of Managed Response Transforms.",
				Required:    true,
				CustomType:  customfield.NewNestedObjectSetType[ManagedTransformsManagedResponseHeadersModel](ctx),
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
