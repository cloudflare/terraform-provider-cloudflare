// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hostname_tls_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r HostnameTLSSettingResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"zone_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"setting_id": schema.StringAttribute{
						Description: "The TLS Setting name.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("ciphers", "min_tls_version", "http2"),
						},
					},
					"hostname": schema.StringAttribute{
						Description: "The hostname for which the tls settings are set.",
						Optional:    true,
					},
					"value": schema.Float64Attribute{
						Description: "The tls setting value.",
						Required:    true,
					},
					"created_at": schema.StringAttribute{
						Description: "This is the time the tls setting was originally created for this hostname.",
						Computed:    true,
					},
					"status": schema.StringAttribute{
						Computed: true,
					},
					"updated_at": schema.StringAttribute{
						Description: "This is the time the tls setting was updated.",
						Computed:    true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state HostnameTLSSettingModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
