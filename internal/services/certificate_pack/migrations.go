// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_pack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithUpgradeState = &CertificatePackResource{}

func (r *CertificatePackResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:   "Identifier",
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"certificate_pack_id": schema.StringAttribute{
						Description:   "Identifier",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"zone_id": schema.StringAttribute{
						Description:   "Identifier",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"certificate_authority": schema.StringAttribute{
						Description: "Certificate Authority selected for the order.  For information on any certificate authority specific details or restrictions [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"google",
								"lets_encrypt",
								"ssl_com",
							),
						},
					},
					"cloudflare_branding": schema.BoolAttribute{
						Description: "Whether or not to add Cloudflare Branding for the order.  This will add sni.cloudflaressl.com as the Common Name if set true.",
						Computed:    true,
					},
					"status": schema.StringAttribute{
						Description: "Status of certificate pack.",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"initializing",
								"pending_validation",
								"deleted",
								"pending_issuance",
								"pending_deployment",
								"pending_deletion",
								"pending_expiration",
								"expired",
								"active",
								"initializing_timed_out",
								"validation_timed_out",
								"issuance_timed_out",
								"deployment_timed_out",
								"deletion_timed_out",
								"pending_cleanup",
								"staging_deployment",
								"staging_active",
								"deactivating",
								"inactive",
								"backup_issued",
								"holding_deployment",
							),
						},
					},
					"type": schema.StringAttribute{
						Description: "Type of certificate pack.",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("advanced"),
						},
					},
					"validation_method": schema.StringAttribute{
						Description: "Validation Method selected for the order.",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"txt",
								"http",
								"email",
							),
						},
					},
					"validity_days": schema.Int64Attribute{
						Description: "Validity Days selected for the order.",
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.OneOf(
								14,
								30,
								90,
								365,
							),
						},
					},
					"hosts": schema.ListAttribute{
						Description: "Comma separated list of valid host names for the certificate packs. Must contain the zone apex, may not contain more than 50 hosts, and may not be empty.",
						Computed:    true,
						ElementType: types.StringType,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state CertificatePackModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
