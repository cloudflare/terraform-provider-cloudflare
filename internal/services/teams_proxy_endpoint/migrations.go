// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_proxy_endpoint

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r TeamsProxyEndpointResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"account_id": schema.StringAttribute{
						Required: true,
					},
					"ips": schema.ListAttribute{
						Description: "A list of CIDRs to restrict ingress connections.",
						Required:    true,
						ElementType: types.StringType,
					},
					"name": schema.StringAttribute{
						Description: "The name of the proxy endpoint.",
						Required:    true,
					},
					"created_at": schema.StringAttribute{
						Computed: true,
					},
					"subdomain": schema.StringAttribute{
						Description: "The subdomain to be used as the destination in the proxy client.",
						Computed:    true,
					},
					"updated_at": schema.StringAttribute{
						Computed: true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state TeamsProxyEndpointModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
