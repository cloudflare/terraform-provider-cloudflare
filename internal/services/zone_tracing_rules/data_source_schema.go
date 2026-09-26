// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_tracing_rules

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZoneTracingRulesDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Specify the zone ID.",
				Computed:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Specify the zone ID.",
				Required:    true,
			},
			"rules": schema.ListNestedAttribute{
				Description: "Trace rules in evaluation order.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZoneTracingRulesRulesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							Description: `Available values: "set_trace_settings".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("set_trace_settings"),
							},
						},
						"action_parameters": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZoneTracingRulesRulesActionParametersDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"sampling_ratio": schema.Float64Attribute{
									Description: "The ratio of requests sampled for tracing, from 0 to 1.",
									Computed:    true,
									Validators: []validator.Float64{
										float64validator.Between(0, 1),
									},
								},
							},
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"enabled": schema.BoolAttribute{
							Computed: true,
						},
						"expression": schema.StringAttribute{
							Description: "A Rules language expression that selects requests.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *ZoneTracingRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZoneTracingRulesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
