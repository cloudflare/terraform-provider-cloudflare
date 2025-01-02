// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippet_rules

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*SnippetRulesDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				Computed: true,
			},
			"enabled": schema.BoolAttribute{
				Computed: true,
			},
			"expression": schema.StringAttribute{
				Computed: true,
			},
			"snippet_name": schema.StringAttribute{
				Description: "Snippet identifying name",
				Computed:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"zone_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
				},
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
