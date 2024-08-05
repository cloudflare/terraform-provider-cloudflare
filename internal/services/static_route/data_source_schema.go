// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package static_route

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &StaticRouteDataSource{}

func (d *StaticRouteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"route_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"route": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (d *StaticRouteDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
