// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel_virtual_network

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r TunnelVirtualNetworkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Cloudflare account ID",
				Required:    true,
			},
			"virtual_network_id": schema.StringAttribute{
				Description: "UUID of the virtual network.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "A user-friendly name for the virtual network.",
				Required:    true,
			},
			"comment": schema.StringAttribute{
				Description: "Optional remark describing the virtual network.",
				Optional:    true,
			},
			"is_default": schema.BoolAttribute{
				Description: "If `true`, this virtual network is the default for the account.",
				Optional:    true,
			},
		},
	}
}
