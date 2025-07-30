// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippet_rules

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*SnippetRulesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "The unique ID of the zone.",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"rules": schema.ListNestedAttribute{
				Description: "A list of snippet rules.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[SnippetRuleDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique ID of the rule.",
							Computed:    true,
						},
						"expression": schema.StringAttribute{
							Description: "The expression defining which traffic will match the rule.",
							Computed:    true,
						},
						"last_updated": schema.StringAttribute{
							Description: "The timestamp of when the rule was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"snippet_name": schema.StringAttribute{
							Description: "The identifying name of the snippet.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "An informative description of the rule.",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the rule should be executed.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *SnippetRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *SnippetRulesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
