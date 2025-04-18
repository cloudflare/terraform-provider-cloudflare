// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_route

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkersRouteDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier.",
				Computed:    true,
			},
			"route_id": schema.StringAttribute{
				Description: "Identifier.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"pattern": schema.StringAttribute{
				Description: "Pattern to match incoming requests against. [Learn more](https://developers.cloudflare.com/workers/configuration/routing/routes/#matching-behavior).",
				Computed:    true,
			},
			"script": schema.StringAttribute{
				Description: "Name of the script to run if the route matches.",
				Computed:    true,
			},
		},
	}
}

func (d *WorkersRouteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WorkersRouteDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
