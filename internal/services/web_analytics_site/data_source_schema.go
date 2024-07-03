// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web_analytics_site

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &WebAnalyticsSiteDataSource{}
var _ datasource.DataSourceWithValidateConfig = &WebAnalyticsSiteDataSource{}

func (r WebAnalyticsSiteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"site_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"auto_install": schema.BoolAttribute{
				Description: "If enabled, the JavaScript snippet is automatically injected for orange-clouded sites.",
				Optional:    true,
			},
			"created": schema.StringAttribute{
				Optional: true,
			},
			"rules": schema.ListNestedAttribute{
				Description: "A list of rules.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The Web Analytics rule identifier.",
							Optional:    true,
						},
						"created": schema.StringAttribute{
							Computed: true,
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
						"paths": schema.StringAttribute{
							Description: "The paths the rule will be applied to.",
							Optional:    true,
						},
						"priority": schema.Float64Attribute{
							Optional: true,
						},
					},
				},
			},
			"ruleset": schema.SingleNestedAttribute{
				Optional: true,
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
			"site_tag": schema.StringAttribute{
				Description: "The Web Analytics site identifier.",
				Optional:    true,
			},
			"site_token": schema.StringAttribute{
				Description: "The Web Analytics site token.",
				Optional:    true,
			},
			"snippet": schema.StringAttribute{
				Description: "Encoded JavaScript snippet.",
				Optional:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"order_by": schema.StringAttribute{
						Description: "The property used to sort the list of results.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("host", "created"),
						},
					},
					"page": schema.Float64Attribute{
						Description: "Current page within the paginated list of results.",
						Optional:    true,
					},
					"per_page": schema.Float64Attribute{
						Description: "Number of items to return per page of results.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (r *WebAnalyticsSiteDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *WebAnalyticsSiteDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
