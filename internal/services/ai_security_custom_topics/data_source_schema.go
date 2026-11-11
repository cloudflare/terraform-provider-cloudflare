// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_security_custom_topics

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*AISecurityCustomTopicsDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Defines the zone.",
				Computed:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Defines the zone.",
				Required:    true,
			},
			"topics": schema.ListNestedAttribute{
				Description: "Custom topic categories for AI Security for Apps content detection.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[AISecurityCustomTopicsTopicsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"label": schema.StringAttribute{
							Description: "Unique label identifier. Must contain only lowercase letters (a–z), digits (0–9), and hyphens.",
							Computed:    true,
						},
						"topic": schema.StringAttribute{
							Description: "Description of the topic category. Must contain 2–50 UTF-8 bytes. Text in any language is supported, including spaces, numbers, punctuation, and symbols. Tabs, line breaks, invalid text, and invisible formatting characters are not allowed.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *AISecurityCustomTopicsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AISecurityCustomTopicsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
