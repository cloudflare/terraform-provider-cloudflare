// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_ssl

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithUpgradeState = &CustomSSLResource{}

func (r *CustomSSLResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"zone_id": schema.StringAttribute{
						Description:   "Identifier",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"custom_certificate_id": schema.StringAttribute{
						Description:   "Identifier",
						Optional:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"type": schema.StringAttribute{
						Description: "The type 'legacy_custom' enables support for legacy clients which do not include SNI in the TLS handshake.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("legacy_custom", "sni_custom"),
						},
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						Default:       stringdefault.StaticString("legacy_custom"),
					},
					"certificate": schema.StringAttribute{
						Description: "The zone's SSL certificate or certificate and the intermediate(s).",
						Required:    true,
					},
					"private_key": schema.StringAttribute{
						Description: "The zone's private key.",
						Required:    true,
					},
					"policy": schema.StringAttribute{
						Description: "Specify the policy that determines the region where your private key will be held locally. HTTPS connections to any excluded data center will still be fully encrypted, but will incur some latency while Keyless SSL is used to complete the handshake with the nearest allowed data center. Any combination of countries, specified by their two letter country code (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements) can be chosen, such as 'country: IN', as well as 'region: EU' which refers to the EU region. If there are too few data centers satisfying the policy, it will be rejected.",
						Optional:    true,
					},
					"geo_restrictions": schema.SingleNestedAttribute{
						Description: "Specify the region where your private key can be held locally for optimal TLS performance. HTTPS connections to any excluded data center will still be fully encrypted, but will incur some latency while Keyless SSL is used to complete the handshake with the nearest allowed data center. Options allow distribution to only to U.S. data centers, only to E.U. data centers, or only to highest security data centers. Default distribution is to all Cloudflare datacenters, for optimal performance.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"label": schema.StringAttribute{
								Optional: true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"us",
										"eu",
										"highest_security",
									),
								},
							},
						},
					},
					"bundle_method": schema.StringAttribute{
						Description: "A ubiquitous bundle has the highest probability of being verified everywhere, even by clients using outdated or unusual trust stores. An optimal bundle uses the shortest chain and newest intermediates. And the force bundle verifies the chain, but does not otherwise modify it.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"ubiquitous",
								"optimal",
								"force",
							),
						},
						Default: stringdefault.StaticString("ubiquitous"),
					},
					"id": schema.StringAttribute{
						Description: "Identifier",
						Computed:    true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state CustomSSLModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
