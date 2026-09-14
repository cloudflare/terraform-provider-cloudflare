// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package field_extractor

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*FieldExtractorDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Cloudflare account ID.",
				Required:    true,
			},
			"extractor": schema.StringAttribute{
				Description: "Extractor type.",
				Required:    true,
			},
			"rules": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[FieldExtractorRulesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"fields": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[FieldExtractorRulesFieldsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"expression": schema.StringAttribute{
										Description: "Wirefilter value expression.",
										Computed:    true,
									},
									"name": schema.StringAttribute{
										Description: "Field name.",
										Computed:    true,
									},
								},
							},
						},
						"ref": schema.StringAttribute{
							Description: "Stable rule identifier.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "Human-readable rule description.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *FieldExtractorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *FieldExtractorDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
