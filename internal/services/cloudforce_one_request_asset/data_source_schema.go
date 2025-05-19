// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_asset

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*CloudforceOneRequestAssetDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"asset_id": schema.StringAttribute{
				Description: "UUID.",
				Required:    true,
			},
			"request_id": schema.StringAttribute{
				Description: "UUID.",
				Required:    true,
			},
			"created": schema.StringAttribute{
				Description: "Defines the asset creation time.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"description": schema.StringAttribute{
				Description: "Asset description.",
				Computed:    true,
			},
			"file_type": schema.StringAttribute{
				Description: "Asset file type.",
				Computed:    true,
			},
			"id": schema.Int64Attribute{
				Description: "Asset ID.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Asset name.",
				Computed:    true,
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
