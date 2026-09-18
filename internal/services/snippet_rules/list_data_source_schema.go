// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippet_rules

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*SnippetRulesListDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Snippets Read",
				"Snippets Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Use this field to specify the unique ID of the zone.",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[SnippetRulesListResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
						Description: "Specify the unique ID of the rule.",
						Computed:    true,
					},
					"expression": schema.StringAttribute{
						Description: "Define the expression that determines which traffic matches the rule.",
						Computed:    true,
					},
					"last_updated": schema.StringAttribute{
						Description: "Specify the timestamp of when the rule was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"snippet_name": schema.StringAttribute{
						Description: "Identify the snippet.",
						Computed:    true,
					},
					"description": schema.StringAttribute{
						Description: "Provide an informative description of the rule.",
						Computed:    true,
					},
					"enabled": schema.BoolAttribute{
						Description: "Indicate whether to execute the rule.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *SnippetRulesListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *SnippetRulesListDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
