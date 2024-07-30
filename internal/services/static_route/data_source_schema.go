// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package static_route

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &StaticRouteDataSource{}
var _ datasource.DataSourceWithValidateConfig = &StaticRouteDataSource{}

func (r StaticRouteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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

func (r *StaticRouteDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *StaticRouteDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
