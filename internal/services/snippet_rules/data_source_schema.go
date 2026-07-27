// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippet_rules

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*SnippetRulesDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Snippets Read",
				"Snippets Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Use this field to specify the unique ID of the zone.",
				Computed:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Use this field to specify the unique ID of the zone.",
				Required:    true,
			},
		},
	}
}

func (d *SnippetRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *SnippetRulesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
