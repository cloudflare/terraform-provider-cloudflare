// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_managed_networks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r DeviceManagedNetworksResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "API UUID.",
						Computed:    true,
					},
					"network_id": schema.StringAttribute{
						Description:   "API UUID.",
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"account_id": schema.StringAttribute{
						Required: true,
					},
					"config": schema.SingleNestedAttribute{
						Description: "The configuration object containing information for the WARP client to detect the managed network.",
						Required:    true,
						Attributes: map[string]schema.Attribute{
							"tls_sockaddr": schema.StringAttribute{
								Description: "A network address of the form \"host:port\" that the WARP client will use to detect the presence of a TLS host.",
								Required:    true,
							},
							"sha256": schema.StringAttribute{
								Description: "The SHA-256 hash of the TLS certificate presented by the host found at tls_sockaddr. If absent, regular certificate verification (trusted roots, valid timestamp, etc) will be used to validate the certificate.",
								Optional:    true,
							},
						},
					},
					"name": schema.StringAttribute{
						Description: "The name of the device managed network. This name must be unique.",
						Required:    true,
					},
					"type": schema.StringAttribute{
						Description: "The type of device managed network.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("tls"),
						},
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state DeviceManagedNetworksModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
