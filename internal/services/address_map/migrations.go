// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address_map

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r AddressMapResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:   "Identifier",
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"description": schema.StringAttribute{
						Description: "An optional description field which may be used to describe the types of IPs or zones on the map.",
						Optional:    true,
					},
					"enabled": schema.BoolAttribute{
						Description: "Whether the Address Map is enabled or not. Cloudflare's DNS will not respond with IP addresses on an Address Map until the map is enabled.",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
					},
					"ips": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
					},
					"memberships": schema.ListNestedAttribute{
						Description: "Zones and Accounts which will be assigned IPs on this Address Map. A zone membership will take priority over an account membership.",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"can_delete": schema.BoolAttribute{
									Description: "Controls whether the membership can be deleted via the API or not.",
									Computed:    true,
								},
								"created_at": schema.StringAttribute{
									Optional: true,
								},
								"identifier": schema.StringAttribute{
									Description: "The identifier for the membership (eg. a zone or account tag).",
									Optional:    true,
								},
								"kind": schema.StringAttribute{
									Description: "The type of the membership.",
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("zone", "account"),
									},
								},
							},
						},
					},
					"can_delete": schema.BoolAttribute{
						Description: "If set to false, then the Address Map cannot be deleted via API. This is true for Cloudflare-managed maps.",
						Computed:    true,
					},
					"can_modify_ips": schema.BoolAttribute{
						Description: "If set to false, then the IPs on the Address Map cannot be modified via the API. This is true for Cloudflare-managed maps.",
						Computed:    true,
					},
					"created_at": schema.StringAttribute{
						Computed: true,
					},
					"default_sni": schema.StringAttribute{
						Description: "If you have legacy TLS clients which do not send the TLS server name indicator, then you can specify one default SNI on the map. If Cloudflare receives a TLS handshake from a client without an SNI, it will respond with the default SNI on those IPs. The default SNI can be any valid zone or subdomain owned by the account.",
						Computed:    true,
					},
					"modified_at": schema.StringAttribute{
						Computed: true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state AddressMapModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
