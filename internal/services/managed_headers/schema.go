// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_headers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r ManagedHeadersResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"managed_request_headers": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Human-readable identifier of the Managed Transform.",
							Optional:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "When true, the Managed Transform is enabled.",
							Optional:    true,
						},
					},
				},
			},
			"managed_response_headers": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Human-readable identifier of the Managed Transform.",
							Optional:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "When true, the Managed Transform is enabled.",
							Optional:    true,
						},
					},
				},
			},
		},
	}
}
