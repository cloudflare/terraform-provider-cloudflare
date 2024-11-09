// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_cache_variants

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZoneCacheVariantsDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"id": schema.StringAttribute{
				Description: "ID of the zone setting.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("variants"),
				},
			},
			"modified_on": schema.StringAttribute{
				Description: "last time this setting was modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"value": schema.SingleNestedAttribute{
				Description: "Value of the zone setting.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZoneCacheVariantsValueDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"avif": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for avif.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"bmp": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for bmp.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"gif": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for gif.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"jp2": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for jp2.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"jpeg": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for jpeg.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"jpg": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for jpg.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"jpg2": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for jpg2.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"png": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for png.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"tif": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for tif.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"tiff": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for tiff.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"webp": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for webp.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
				},
			},
		},
	}
}

func (d *ZoneCacheVariantsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZoneCacheVariantsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
