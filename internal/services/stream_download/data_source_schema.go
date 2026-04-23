// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_download

import (
	"context"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*StreamDownloadDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Stream Read",
				"Stream Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Description: "A Cloudflare-generated unique identifier for a media item.",
				Required:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Optional:    true,
			},
		},
	}
}

func (d *StreamDownloadDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *StreamDownloadDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
