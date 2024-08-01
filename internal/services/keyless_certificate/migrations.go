// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keyless_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r KeylessCertificateResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:   "Keyless certificate identifier tag.",
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"zone_id": schema.StringAttribute{
						Description:   "Identifier",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"certificate": schema.StringAttribute{
						Description:   "The zone's SSL certificate or SSL certificate and intermediate(s).",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"bundle_method": schema.StringAttribute{
						Description: "A ubiquitous bundle has the highest probability of being verified everywhere, even by clients using outdated or unusual trust stores. An optimal bundle uses the shortest chain and newest intermediates. And the force bundle verifies the chain, but does not otherwise modify it.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("ubiquitous", "optimal", "force"),
						},
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						Default:       stringdefault.StaticString("ubiquitous"),
					},
					"host": schema.StringAttribute{
						Description: "The keyless SSL name.",
						Required:    true,
					},
					"enabled": schema.BoolAttribute{
						Description: "Whether or not the Keyless SSL is on or off.",
						Optional:    true,
					},
					"name": schema.StringAttribute{
						Description: "The keyless SSL name.",
						Optional:    true,
					},
					"tunnel": schema.SingleNestedAttribute{
						Description: "Configuration for using Keyless SSL through a Cloudflare Tunnel",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"private_ip": schema.StringAttribute{
								Description: "Private IP of the Key Server Host",
								Required:    true,
							},
							"vnet_id": schema.StringAttribute{
								Description: "Cloudflare Tunnel Virtual Network ID",
								Required:    true,
							},
						},
					},
					"port": schema.Float64Attribute{
						Description: "The keyless SSL port used to communicate between Cloudflare and the client's Keyless SSL server.",
						Computed:    true,
						Optional:    true,
						Default:     float64default.StaticFloat64(24008),
					},
					"created_on": schema.StringAttribute{
						Description: "When the Keyless SSL was created.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"modified_on": schema.StringAttribute{
						Description: "When the Keyless SSL was last modified.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"status": schema.StringAttribute{
						Description: "Status of the Keyless SSL.",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("active", "deleted"),
						},
					},
					"permissions": schema.ListAttribute{
						Description: "Available permissions for the Keyless SSL for the current user requesting the item.",
						Computed:    true,
						ElementType: jsontypes.NewNormalizedNull().Type(ctx),
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state KeylessCertificateModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
