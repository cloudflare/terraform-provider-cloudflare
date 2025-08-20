// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_transforms

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithUpgradeState = (*ManagedTransformsResource)(nil)

func (r *ManagedTransformsResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			// PriorSchema must work with BOTH v4 and v5 states
			// In v4: managed_request_headers and managed_response_headers were optional TypeSet blocks
			// In v5: they are required ListNestedAttribute
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"zone_id": schema.StringAttribute{
						Required: true,
					},
					// Use ListNestedAttribute for both v4 blocks and v5 attributes
					// This works because the structure is the same
					"managed_request_headers": schema.ListNestedAttribute{
						Optional: true, // Optional in v4, Required in v5
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Required: true,
								},
								"enabled": schema.BoolAttribute{
									Required: true,
								},
								"has_conflict": schema.BoolAttribute{
									Computed: true,
								},
								"conflicts_with": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
							},
						},
					},
					"managed_response_headers": schema.ListNestedAttribute{
						Optional: true, // Optional in v4, Required in v5
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Required: true,
								},
								"enabled": schema.BoolAttribute{
									Required: true,
								},
								"has_conflict": schema.BoolAttribute{
									Computed: true,
								},
								"conflicts_with": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
							},
						},
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				// Define struct matching PriorSchema
				var priorStateData struct {
					ID                     types.String                                     `tfsdk:"id"`
					ZoneID                 types.String                                     `tfsdk:"zone_id"`
					ManagedRequestHeaders  *[]*ManagedTransformsManagedRequestHeadersModel  `tfsdk:"managed_request_headers"`
					ManagedResponseHeaders *[]*ManagedTransformsManagedResponseHeadersModel `tfsdk:"managed_response_headers"`
				}

				resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)
				if resp.Diagnostics.HasError() {
					return
				}

				// Initialize new state
				newState := ManagedTransformsModel{
					ID:     priorStateData.ID,
					ZoneID: priorStateData.ZoneID,
				}

				// Handle managed_request_headers
				// In v4, this was optional and could be nil
				// In v5, it's required and must be at least an empty list
				if priorStateData.ManagedRequestHeaders != nil {
					newState.ManagedRequestHeaders = priorStateData.ManagedRequestHeaders
				} else {
					// If nil (not specified in v4), create an empty list
					emptyRequestHeaders := make([]*ManagedTransformsManagedRequestHeadersModel, 0)
					newState.ManagedRequestHeaders = &emptyRequestHeaders
				}

				// Handle managed_response_headers
				// Same logic as request headers
				if priorStateData.ManagedResponseHeaders != nil {
					newState.ManagedResponseHeaders = priorStateData.ManagedResponseHeaders
				} else {
					// If nil (not specified in v4), create an empty list
					emptyResponseHeaders := make([]*ManagedTransformsManagedResponseHeadersModel, 0)
					newState.ManagedResponseHeaders = &emptyResponseHeaders
				}

				// Marshal the upgraded state
				resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
			},
		},
	}
}
