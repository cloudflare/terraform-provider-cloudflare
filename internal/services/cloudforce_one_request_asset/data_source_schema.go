// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_asset

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*CloudforceOneRequestAssetDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_identifier": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"asset_identifer": schema.StringAttribute{
				Description: "UUID",
				Required:    true,
			},
			"request_identifier": schema.StringAttribute{
				Description: "UUID",
				Required:    true,
			},
		},
	}
}

func (d *CloudforceOneRequestAssetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CloudforceOneRequestAssetDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
