// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package content_scanning

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ContentScanningDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Defines an identifier.",
				Required:    true,
			},
			"modified": schema.StringAttribute{
				Description: "Defines the last modification date (ISO 8601) of the Content Scanning status.",
				Computed:    true,
			},
			"value": schema.StringAttribute{
				Description: "Defines the status of Content Scanning.",
				Computed:    true,
			},
		},
	}
}

func (d *ContentScanningDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ContentScanningDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
