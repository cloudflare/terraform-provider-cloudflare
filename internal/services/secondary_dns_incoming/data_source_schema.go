// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns_incoming

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*SecondaryDNSIncomingDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			"auto_refresh_seconds": schema.Float64Attribute{
				Description: "How often should a secondary zone auto refresh regardless of DNS NOTIFY.\nNot applicable for primary zones.",
				Optional:    true,
			},
			"checked_time": schema.StringAttribute{
				Description: "The time for a specific event.",
				Optional:    true,
			},
			"created_time": schema.StringAttribute{
				Description: "The time for a specific event.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Optional: true,
			},
			"modified_time": schema.StringAttribute{
				Description: "The time for a specific event.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Zone name.",
				Optional:    true,
			},
			"soa_serial": schema.Float64Attribute{
				Description: "The serial number of the SOA for the given zone.",
				Optional:    true,
			},
			"peers": schema.ListAttribute{
				Description: "A list of peer tags.",
				Optional:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (d *SecondaryDNSIncomingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *SecondaryDNSIncomingDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
