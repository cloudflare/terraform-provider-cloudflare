// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package image_variant

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ImageVariantDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Required:    true,
			},
			"variant_id": schema.StringAttribute{
				Required: true,
			},
			"variant": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[ImageVariantVariantDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"options": schema.SingleNestedAttribute{
						Description: "Allows you to define image resizing sizes for different use cases.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ImageVariantVariantOptionsDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"fit": schema.StringAttribute{
								Description: "The fit property describes how the width and height dimensions should be interpreted.\navailable values: \"scale-down\", \"contain\", \"cover\", \"crop\", \"pad\"",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"scale-down",
										"contain",
										"cover",
										"crop",
										"pad",
									),
								},
							},
							"height": schema.Float64Attribute{
								Description: "Maximum height in image pixels.",
								Computed:    true,
								Validators: []validator.Float64{
									float64validator.AtLeast(1),
								},
							},
							"metadata": schema.StringAttribute{
								Description: "What EXIF data should be preserved in the output image.\navailable values: \"keep\", \"copyright\", \"none\"",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"keep",
										"copyright",
										"none",
									),
								},
							},
							"width": schema.Float64Attribute{
								Description: "Maximum width in image pixels.",
								Computed:    true,
								Validators: []validator.Float64{
									float64validator.AtLeast(1),
								},
							},
						},
					},
					"never_require_signed_urls": schema.BoolAttribute{
						Description: "Indicates whether the variant can access an image without a signature, regardless of image access control.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *ImageVariantDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ImageVariantDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
