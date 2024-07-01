// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_address

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r EmailRoutingAddressResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:   "Destination address identifier.",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"account_identifier": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"email": schema.StringAttribute{
						Description: "The contact email address of the user.",
						Required:    true,
					},
					"created": schema.StringAttribute{
						Description: "The date and time the destination address has been created.",
						Computed:    true,
					},
					"modified": schema.StringAttribute{
						Description: "The date and time the destination address was last modified.",
						Computed:    true,
					},
					"tag": schema.StringAttribute{
						Description: "Destination address tag. (Deprecated, replaced by destination address identifier)",
						Computed:    true,
					},
					"verified": schema.StringAttribute{
						Description: "The date and time the destination address has been verified. Null means not verified yet.",
						Computed:    true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state EmailRoutingAddressModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
