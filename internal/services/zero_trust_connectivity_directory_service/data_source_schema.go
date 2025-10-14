// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_connectivity_directory_service

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustConnectivityDirectoryServiceDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"service_id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (d *ZeroTrustConnectivityDirectoryServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustConnectivityDirectoryServiceDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
