// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script_subdomain

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkersScriptSubdomainDataSource)(nil)

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
			"enabled": schema.BoolAttribute{
				Description: "Whether the Worker is available on the workers.dev subdomain.",
				Computed:    true,
			},
			"previews_enabled": schema.BoolAttribute{
				Description: "Whether the Worker's Preview URLs should be available on the workers.dev subdomain.",
				Computed:    true,
			},
		},
	}
}

func (d *WorkersScriptSubdomainDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WorkersScriptSubdomainDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
