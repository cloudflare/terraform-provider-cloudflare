// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web_analytics_site

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*WebAnalyticsSiteResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The Web Analytics site identifier.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"site_tag": schema.StringAttribute{
				Description:   "The Web Analytics site identifier.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"auto_install": schema.BoolAttribute{
				Description:   "If enabled, the JavaScript snippet is automatically injected for orange-clouded sites.",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"enabled": schema.BoolAttribute{
				Description:   "Enables or disables RUM. This option can be used only when auto_install is set to true.",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"host": schema.StringAttribute{
				Description:   "The hostname to use for gray-clouded sites.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"lite": schema.BoolAttribute{
				Description:   "If enabled, the JavaScript snippet will not be injected for visitors from the EU.",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"zone_tag": schema.StringAttribute{
				Description: "The zone identifier.",
				Optional:    true,
			},
			"created": schema.StringAttribute{
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"site_token": schema.StringAttribute{
				Description:   "The Web Analytics site token.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"snippet": schema.StringAttribute{
				Description:   "Encoded JavaScript snippet.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"rules": schema.ListNestedAttribute{
				Description:   "A list of rules.",
				Computed:      true,
				CustomType:    customfield.NewNestedObjectListType[WebAnalyticsSiteRulesModel](ctx),
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description:   "The Web Analytics rule identifier.",
							Computed:      true,
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"created": schema.StringAttribute{
							Computed:      true,
							CustomType:    timetypes.RFC3339Type{},
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"host": schema.StringAttribute{
							Description:   "The hostname the rule will be applied to.",
							Computed:      true,
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"inclusive": schema.BoolAttribute{
							Description:   "Whether the rule includes or excludes traffic from being measured.",
							Computed:      true,
							PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
						},
						"is_paused": schema.BoolAttribute{
							Description:   "Whether the rule is paused or not.",
							Computed:      true,
							PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
						},
						"paths": schema.ListAttribute{
							Description:   "The paths the rule will be applied to.",
							Computed:      true,
							CustomType:    customfield.NewListType[types.String](ctx),
							ElementType:   types.StringType,
							PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						},
						"priority": schema.Float64Attribute{
							Computed:      true,
							PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown()},
						},
					},
				},
			},
			"ruleset": schema.SingleNestedAttribute{
				Computed:      true,
				CustomType:    customfield.NewNestedObjectType[WebAnalyticsSiteRulesetModel](ctx),
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:   "The Web Analytics ruleset identifier.",
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"enabled": schema.BoolAttribute{
						Description:   "Whether the ruleset is enabled.",
						Computed:      true,
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"zone_name": schema.StringAttribute{
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"zone_tag": schema.StringAttribute{
						Description:   "The zone identifier.",
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
		},
	}
}

func (r *WebAnalyticsSiteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *WebAnalyticsSiteResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
