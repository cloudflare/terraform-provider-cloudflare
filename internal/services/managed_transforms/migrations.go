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
			// v4 schema had managed_headers resource with blocks
			// v5 schema has managed_transforms resource with list attributes
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"zone_id": schema.StringAttribute{
						Required: true,
					},
					"managed_request_headers": schema.SetNestedAttribute{
						Optional:   true,
						CustomType: customfield.NewNestedObjectSetType[ManagedTransformsManagedRequestHeadersModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Required: true,
								},
								"enabled": schema.BoolAttribute{
									Required: true,
								},
							},
						},
					},
					"managed_response_headers": schema.SetNestedAttribute{
						Optional:   true,
						CustomType: customfield.NewNestedObjectSetType[ManagedTransformsManagedResponseHeadersModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Required: true,
								},
								"enabled": schema.BoolAttribute{
									Required: true,
								},
							},
						},
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var priorStateData struct {
					ID                     types.String                                     `tfsdk:"id"`
					ZoneID                 types.String                                     `tfsdk:"zone_id"`
					ManagedRequestHeaders  *[]*ManagedTransformsManagedRequestHeadersModel  `tfsdk:"managed_request_headers" json:"managed_request_headers,required"`
					ManagedResponseHeaders *[]*ManagedTransformsManagedResponseHeadersModel `tfsdk:"managed_response_headers" json:"managed_response_headers,required"`
				}

				resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)
				if resp.Diagnostics.HasError() {
					return
				}

				// Initialize new state
				newState := ManagedTransformsModel{
					ID:                     priorStateData.ID,
					ZoneID:                 priorStateData.ZoneID,
					ManagedRequestHeaders:  priorStateData.ManagedRequestHeaders,
					ManagedResponseHeaders: priorStateData.ManagedResponseHeaders,
				}

				// Ensure both required attributes are present (even if empty)
				if newState.ManagedRequestHeaders == nil {
					emptyHeaders := make([]*ManagedTransformsManagedRequestHeadersModel, 0)
					newState.ManagedRequestHeaders = &emptyHeaders
				}

				if newState.ManagedResponseHeaders == nil {
					emptyHeaders := make([]*ManagedTransformsManagedResponseHeadersModel, 0)
					newState.ManagedResponseHeaders = &emptyHeaders
				}

				// Marshal the upgraded state
				resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
			},
		},
	}
}
