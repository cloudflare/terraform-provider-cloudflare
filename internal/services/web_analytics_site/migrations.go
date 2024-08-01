// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web_analytics_site

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r WebAnalyticsSiteResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "The Web Analytics site identifier.",
						Computed:    true,
					},
					"site_tag": schema.StringAttribute{
						Description:   "The Web Analytics site identifier.",
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"account_id": schema.StringAttribute{
						Description:   "Identifier",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"auto_install": schema.BoolAttribute{
						Description: "If enabled, the JavaScript snippet is automatically injected for orange-clouded sites.",
						Optional:    true,
					},
					"host": schema.StringAttribute{
						Description: "The hostname to use for gray-clouded sites.",
						Optional:    true,
					},
					"zone_tag": schema.StringAttribute{
						Description: "The zone identifier.",
						Optional:    true,
					},
					"created": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"site_token": schema.StringAttribute{
						Description: "The Web Analytics site token.",
						Computed:    true,
					},
					"snippet": schema.StringAttribute{
						Description: "Encoded JavaScript snippet.",
						Computed:    true,
					},
					"rules": schema.ListNestedAttribute{
						Description: "A list of rules.",
						Computed:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The Web Analytics rule identifier.",
									Optional:    true,
								},
								"created": schema.StringAttribute{
									Computed:   true,
									CustomType: timetypes.RFC3339Type{},
								},
								"host": schema.StringAttribute{
									Description: "The hostname the rule will be applied to.",
									Optional:    true,
								},
								"inclusive": schema.BoolAttribute{
									Description: "Whether the rule includes or excludes traffic from being measured.",
									Optional:    true,
								},
								"is_paused": schema.BoolAttribute{
									Description: "Whether the rule is paused or not.",
									Optional:    true,
								},
								"paths": schema.ListAttribute{
									Description: "The paths the rule will be applied to.",
									Optional:    true,
									ElementType: types.StringType,
								},
								"priority": schema.Float64Attribute{
									Optional: true,
								},
							},
						},
					},
					"ruleset": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[WebAnalyticsSiteRulesetModel](ctx),
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description: "The Web Analytics ruleset identifier.",
								Optional:    true,
							},
							"enabled": schema.BoolAttribute{
								Description: "Whether the ruleset is enabled.",
								Optional:    true,
							},
							"zone_name": schema.StringAttribute{
								Optional: true,
							},
							"zone_tag": schema.StringAttribute{
								Description: "The zone identifier.",
								Optional:    true,
							},
						},
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state WebAnalyticsSiteModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
