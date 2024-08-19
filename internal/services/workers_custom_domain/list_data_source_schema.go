// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_custom_domain

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &WorkersCustomDomainsDataSource{}

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifer of the account.",
				Required:    true,
			},
			"environment": schema.StringAttribute{
				Description: "Worker environment associated with the zone and hostname.",
				Optional:    true,
			},
			"hostname": schema.StringAttribute{
				Description: "Hostname of the Worker Domain.",
				Optional:    true,
			},
			"service": schema.StringAttribute{
				Description: "Worker service associated with the zone and hostname.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier of the zone.",
				Optional:    true,
			},
			"zone_name": schema.StringAttribute{
				Description: "Name of the zone.",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifer of the Worker Domain.",
							Computed:    true,
							Optional:    true,
						},
						"environment": schema.StringAttribute{
							Description: "Worker environment associated with the zone and hostname.",
							Computed:    true,
							Optional:    true,
						},
						"hostname": schema.StringAttribute{
							Description: "Hostname of the Worker Domain.",
							Computed:    true,
							Optional:    true,
						},
						"service": schema.StringAttribute{
							Description: "Worker service associated with the zone and hostname.",
							Computed:    true,
							Optional:    true,
						},
						"zone_id": schema.StringAttribute{
							Description: "Identifier of the zone.",
							Computed:    true,
							Optional:    true,
						},
						"zone_name": schema.StringAttribute{
							Description: "Name of the zone.",
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func (d *WorkersCustomDomainsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *WorkersCustomDomainsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
