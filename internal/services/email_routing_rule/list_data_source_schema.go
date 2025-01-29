// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_rule

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*EmailRoutingRulesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Filter by enabled routing rules.",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[EmailRoutingRulesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Routing rule identifier.",
							Computed:    true,
						},
						"actions": schema.ListNestedAttribute{
							Description: "List actions patterns.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[EmailRoutingRulesActionsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										Description: "Type of supported action.",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"drop",
												"forward",
												"worker",
											),
										},
									},
									"value": schema.ListAttribute{
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
								},
							},
						},
						"enabled": schema.BoolAttribute{
							Description: "Routing rule status.",
							Computed:    true,
						},
						"matchers": schema.ListNestedAttribute{
							Description: "Matching patterns to forward to your actions.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[EmailRoutingRulesMatchersDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"field": schema.StringAttribute{
										Description: "Field for type matcher.",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("to"),
										},
									},
									"type": schema.StringAttribute{
										Description: "Type of matcher.",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("literal"),
										},
									},
									"value": schema.StringAttribute{
										Description: "Value for matcher.",
										Computed:    true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "Routing rule name.",
							Computed:    true,
						},
						"priority": schema.Float64Attribute{
							Description: "Priority of the routing rule.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.AtLeast(0),
							},
						},
						"tag": schema.StringAttribute{
							Description: "Routing rule tag. (Deprecated, replaced by routing rule identifier)",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *EmailRoutingRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *EmailRoutingRulesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
