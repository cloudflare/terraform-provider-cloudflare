// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web_analytics_site

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*WebAnalyticsSiteResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
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
				CustomType:  customfield.NewNestedObjectListType[WebAnalyticsSiteRulesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The Web Analytics rule identifier.",
							Computed:    true,
						},
						"created": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"host": schema.StringAttribute{
							Description: "The hostname the rule will be applied to.",
							Computed:    true,
						},
						"inclusive": schema.BoolAttribute{
							Description: "Whether the rule includes or excludes traffic from being measured.",
							Computed:    true,
						},
						"is_paused": schema.BoolAttribute{
							Description: "Whether the rule is paused or not.",
							Computed:    true,
						},
						"paths": schema.ListAttribute{
							Description: "The paths the rule will be applied to.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"priority": schema.Float64Attribute{
							Computed: true,
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
						Computed:    true,
					},
					"enabled": schema.BoolAttribute{
						Description: "Whether the ruleset is enabled.",
						Computed:    true,
					},
					"zone_name": schema.StringAttribute{
						Computed: true,
					},
					"zone_tag": schema.StringAttribute{
						Description: "The zone identifier.",
						Computed:    true,
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
