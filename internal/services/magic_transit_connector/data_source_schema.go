// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_connector

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*MagicTransitConnectorDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"connector_id": schema.StringAttribute{
				Optional: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"activated": schema.BoolAttribute{
				Computed: true,
			},
			"interrupt_window_duration_hours": schema.Float64Attribute{
				Computed: true,
			},
			"interrupt_window_hour_of_day": schema.Float64Attribute{
				Computed: true,
			},
			"last_heartbeat": schema.StringAttribute{
				Computed: true,
			},
			"last_seen_version": schema.StringAttribute{
				Computed: true,
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"notes": schema.StringAttribute{
				Computed: true,
			},
			"timezone": schema.StringAttribute{
				Computed: true,
			},
			"device": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[MagicTransitConnectorDeviceDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"serial_number": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *MagicTransitConnectorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MagicTransitConnectorDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
