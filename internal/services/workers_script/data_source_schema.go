// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkersScriptDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"script_name": schema.StringAttribute{
				Description: "Name of the script, used in URLs and route configuration.",
				Required:    true,
			},
		},
	}
}

func (d *WorkersScriptDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WorkersScriptDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
