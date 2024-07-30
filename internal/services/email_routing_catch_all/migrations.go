// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_catch_all

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r EmailRoutingCatchAllResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:   "Routing rule identifier.",
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"zone_identifier": schema.StringAttribute{
						Description:   "Identifier",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"actions": schema.ListNestedAttribute{
						Description: "List actions for the catch-all routing rule.",
						Required:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									Description: "Type of action for catch-all rule.",
									Required:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("drop", "forward", "worker"),
									},
								},
								"value": schema.ListAttribute{
									Optional:    true,
									ElementType: types.StringType,
								},
							},
						},
						PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
					},
					"matchers": schema.ListNestedAttribute{
						Description: "List of matchers for the catch-all routing rule.",
						Required:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									Description: "Type of matcher. Default is 'all'.",
									Required:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("all"),
									},
								},
							},
						},
						PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
					},
					"enabled": schema.BoolAttribute{
						Description:   "Routing rule status.",
						Computed:      true,
						Optional:      true,
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
						Default:       booldefault.StaticBool(true),
					},
					"name": schema.StringAttribute{
						Description:   "Routing rule name.",
						Optional:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"tag": schema.StringAttribute{
						Description: "Routing rule tag. (Deprecated, replaced by routing rule identifier)",
						Computed:    true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state EmailRoutingCatchAllModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
