// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_service_token

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &AccessServiceTokensDataSource{}
var _ datasource.DataSourceWithValidateConfig = &AccessServiceTokensDataSource{}

func (r AccessServiceTokensDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the service token.",
							Computed:    true,
							Optional:    true,
						},
						"client_id": schema.StringAttribute{
							Description: "The Client ID for the service token. Access will check for this value in the `CF-Access-Client-ID` request header.",
							Computed:    true,
							Optional:    true,
						},
						"created_at": schema.StringAttribute{
							Computed: true,
						},
						"duration": schema.StringAttribute{
							Description: "The duration for how long the service token will be valid. Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The default is 1 year in hours (8760h).",
							Computed:    true,
							Optional:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the service token.",
							Computed:    true,
							Optional:    true,
						},
						"updated_at": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (r *AccessServiceTokensDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *AccessServiceTokensDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
