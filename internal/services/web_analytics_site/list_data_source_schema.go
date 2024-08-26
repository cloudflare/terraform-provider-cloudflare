// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web_analytics_site

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*WebAnalyticsSitesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
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
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[WebAnalyticsSitesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"auto_install": schema.BoolAttribute{
							Description: "If enabled, the JavaScript snippet is automatically injected for orange-clouded sites.",
							Computed:    true,
							Optional:    true,
						},
						"created": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"rules": schema.ListNestedAttribute{
							Description: "A list of rules.",
							Computed:    true,
							Optional:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "The Web Analytics rule identifier.",
										Computed:    true,
										Optional:    true,
									},
									"created": schema.StringAttribute{
										Computed:   true,
										CustomType: timetypes.RFC3339Type{},
									},
									"host": schema.StringAttribute{
										Description: "The hostname the rule will be applied to.",
										Computed:    true,
										Optional:    true,
									},
									"inclusive": schema.BoolAttribute{
										Description: "Whether the rule includes or excludes traffic from being measured.",
										Computed:    true,
										Optional:    true,
									},
									"is_paused": schema.BoolAttribute{
										Description: "Whether the rule is paused or not.",
										Computed:    true,
										Optional:    true,
									},
									"paths": schema.ListAttribute{
										Description: "The paths the rule will be applied to.",
										Computed:    true,
										Optional:    true,
										ElementType: types.StringType,
									},
									"priority": schema.Float64Attribute{
										Computed: true,
										Optional: true,
									},
								},
							},
						},
						"ruleset": schema.SingleNestedAttribute{
							Computed: true,
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The Web Analytics ruleset identifier.",
									Computed:    true,
									Optional:    true,
								},
								"enabled": schema.BoolAttribute{
									Description: "Whether the ruleset is enabled.",
									Computed:    true,
									Optional:    true,
								},
								"zone_name": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"zone_tag": schema.StringAttribute{
									Description: "The zone identifier.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"site_tag": schema.StringAttribute{
							Description: "The Web Analytics site identifier.",
							Computed:    true,
							Optional:    true,
						},
						"site_token": schema.StringAttribute{
							Description: "The Web Analytics site token.",
							Computed:    true,
							Optional:    true,
						},
						"snippet": schema.StringAttribute{
							Description: "Encoded JavaScript snippet.",
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func (d *WebAnalyticsSitesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *WebAnalyticsSitesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
