// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive_configs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r HyperdriveConfigsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"hyperdrive_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"origin": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"password": schema.StringAttribute{
						Description: "The password required to access your origin database. This value is write-only and never returned by the API.",
						Required:    true,
					},
				},
			},
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
		},
	}
}
