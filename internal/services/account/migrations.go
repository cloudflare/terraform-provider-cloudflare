// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithUpgradeState = &AccountResource{}

func (r *AccountResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"account_id": schema.StringAttribute{
						Required: true,
					},
					"name": schema.StringAttribute{
						Description: "Account name",
						Required:    true,
					},
					"settings": schema.SingleNestedAttribute{
						Description: "Account settings",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"abuse_contact_email": schema.StringAttribute{
								Description: "Sets an abuse contact email to notify for abuse reports.",
								Optional:    true,
							},
							"default_nameservers": schema.StringAttribute{
								Description: "Specifies the default nameservers to be used for new zones added to this account.\n\n- `cloudflare.standard` for Cloudflare-branded nameservers\n- `custom.account` for account custom nameservers\n- `custom.tenant` for tenant custom nameservers\n\nSee [Custom Nameservers](https://developers.cloudflare.com/dns/additional-options/custom-nameservers/)\nfor more information.",
								Computed:    true,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"cloudflare.standard",
										"custom.account",
										"custom.tenant",
									),
								},
								Default: stringdefault.StaticString("cloudflare.standard"),
							},
							"enforce_twofactor": schema.BoolAttribute{
								Description: "Indicates whether membership in this account requires that\nTwo-Factor Authentication is enabled",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
							"use_account_custom_ns_by_default": schema.BoolAttribute{
								Description: "Indicates whether new zones should use the account-level custom\nnameservers by default.\n\nDeprecated in favor of `default_nameservers`.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
						},
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state AccountModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
