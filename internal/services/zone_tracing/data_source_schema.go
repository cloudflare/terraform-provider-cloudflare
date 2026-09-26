// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_tracing

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZoneTracingDataSource)(nil)

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
			"enabled": schema.BoolAttribute{
				Description: "Whether Cloudflare Traces is enabled for the zone.",
				Computed:    true,
			},
			"forward_context": schema.BoolAttribute{
				Description: "Whether trace context is sent externally or across a zone boundary.",
				Computed:    true,
			},
			"persist": schema.BoolAttribute{
				Description: "Whether traces are persisted in Cloudflare.",
				Computed:    true,
			},
			"propagation_policy": schema.StringAttribute{
				Description: "When inbound trace context may be continued. Authenticated propagation is not supported yet.\nAvailable values: \"accept\", \"authenticated\", \"reject\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"accept",
						"authenticated",
						"reject",
					),
				},
			},
			"sampling_ratio": schema.Float64Attribute{
				Description: "The ratio of requests sampled for tracing, from 0 to 1.",
				Computed:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"destinations": schema.ListAttribute{
				Description: "Up to 100 OpenTelemetry destination identifiers that receive traces.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
		},
	}
}

func (d *ZoneTracingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZoneTracingDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
