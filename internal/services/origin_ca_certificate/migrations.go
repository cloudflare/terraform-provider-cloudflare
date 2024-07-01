// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_ca_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r OriginCACertificateResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"certificate_id": schema.StringAttribute{
						Description: "Identifier",
						Optional:    true,
					},
					"csr": schema.StringAttribute{
						Description: "The Certificate Signing Request (CSR). Must be newline-encoded.",
						Optional:    true,
					},
					"hostnames": schema.ListAttribute{
						Description: "Array of hostnames or wildcard names (e.g., *.example.com) bound to the certificate.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"request_type": schema.StringAttribute{
						Description: "Signature type desired on certificate (\"origin-rsa\" (rsa), \"origin-ecc\" (ecdsa), or \"keyless-certificate\" (for Keyless SSL servers).",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("origin-rsa", "origin-ecc", "keyless-certificate"),
						},
					},
					"requested_validity": schema.Float64Attribute{
						Description: "The number of days for which the certificate should be valid.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.OneOf(7, 30, 90, 365, 730, 1095, 5475),
						},
						Default: float64default.StaticFloat64(5475),
					},
					"id": schema.StringAttribute{
						Description: "Identifier",
						Computed:    true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state OriginCACertificateModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
