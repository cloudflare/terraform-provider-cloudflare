// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_mutual_tls_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r AccessMutualTLSCertificateResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:   "The ID of the application that will use this certificate.",
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"account_id": schema.StringAttribute{
						Description:   "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
						Optional:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"zone_id": schema.StringAttribute{
						Description:   "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
						Optional:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"name": schema.StringAttribute{
						Description: "The name of the certificate.",
						Required:    true,
					},
					"associated_hostnames": schema.ListAttribute{
						Description: "The hostnames of the applications that will use this certificate.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"certificate": schema.StringAttribute{
						Description:   "The certificate content.",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"created_at": schema.StringAttribute{
						Computed: true,
					},
					"expires_on": schema.StringAttribute{
						Computed: true,
					},
					"fingerprint": schema.StringAttribute{
						Description: "The MD5 fingerprint of the certificate.",
						Computed:    true,
					},
					"updated_at": schema.StringAttribute{
						Computed: true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state AccessMutualTLSCertificateModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
