// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r AuthenticatedOriginPullsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"hostname": schema.StringAttribute{
				Description: "The hostname on the origin for which the client certificate uploaded will be used.",
				Optional:    true,
			},
			"config": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"cert_id": schema.StringAttribute{
							Description: "Certificate identifier tag.",
							Optional:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Indicates whether hostname-level authenticated origin pulls is enabled. A null value voids the association.",
							Optional:    true,
						},
						"hostname": schema.StringAttribute{
							Description: "The hostname on the origin for which the client certificate uploaded will be used.",
							Optional:    true,
						},
					},
				},
			},
		},
	}
}
