// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mtls_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r MTLSCertificateResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:   "Identifier",
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"account_id": schema.StringAttribute{
						Description:   "Identifier",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"ca": schema.BoolAttribute{
						Description:   "Indicates whether the certificate is a CA or leaf certificate.",
						Required:      true,
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
					},
					"certificates": schema.StringAttribute{
						Description:   "The uploaded root CA certificate.",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"name": schema.StringAttribute{
						Description:   "Optional unique name for the certificate. Only used for human readability.",
						Optional:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"private_key": schema.StringAttribute{
						Description:   "The private key for the certificate",
						Optional:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"expires_on": schema.StringAttribute{
						Description: "When the certificate expires.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"issuer": schema.StringAttribute{
						Description: "The certificate authority that issued the certificate.",
						Computed:    true,
					},
					"serial_number": schema.StringAttribute{
						Description: "The certificate serial number.",
						Computed:    true,
					},
					"signature": schema.StringAttribute{
						Description: "The type of hash used for the certificate.",
						Computed:    true,
					},
					"updated_at": schema.StringAttribute{
						Description: "This is the time the certificate was updated.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"uploaded_on": schema.StringAttribute{
						Description: "This is the time the certificate was uploaded.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state MTLSCertificateModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
