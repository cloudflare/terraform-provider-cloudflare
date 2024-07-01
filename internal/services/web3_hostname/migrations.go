// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3_hostname

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r Web3HostnameResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:   "Identifier",
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"zone_identifier": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"target": schema.StringAttribute{
						Description: "Target gateway of the hostname.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("ethereum", "ipfs", "ipfs_universal_path"),
						},
					},
					"description": schema.StringAttribute{
						Description: "An optional description of the hostname.",
						Optional:    true,
					},
					"dnslink": schema.StringAttribute{
						Description: "DNSLink value used if the target is ipfs.",
						Optional:    true,
					},
					"created_on": schema.StringAttribute{
						Computed: true,
					},
					"modified_on": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Description: "The hostname that will point to the target gateway via CNAME.",
						Computed:    true,
					},
					"status": schema.StringAttribute{
						Description: "Status of the hostname's activation.",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("active", "pending", "deleting", "error"),
						},
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state Web3HostnameModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
