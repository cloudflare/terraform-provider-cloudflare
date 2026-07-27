// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flagship_app

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*FlagshipAppDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Flagship Read",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "App identifier.",
				Computed:    true,
			},
			"app_id": schema.StringAttribute{
				Description: "App identifier.",
				Required:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Cloudflare account ID.",
				Required:    true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"updated_by": schema.StringAttribute{
				Description: "Email of the actor who last modified the app, or `edge-gateway` for gateway-authenticated changes.",
				Computed:    true,
			},
		},
	}
}

func (d *FlagshipAppDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *FlagshipAppDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
