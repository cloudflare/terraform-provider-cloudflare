// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_watermark

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*StreamWatermarkDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The account identifier tag.",
				Required:    true,
			},
			"identifier": schema.StringAttribute{
				Description: "The unique identifier for a watermark profile.",
				Required:    true,
			},
			"created": schema.StringAttribute{
				Description: "The date and a time a watermark profile was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"downloaded_from": schema.StringAttribute{
				Description: "The source URL for a downloaded image. If the watermark profile was created via direct upload, this field is null.",
				Computed:    true,
			},
			"height": schema.Int64Attribute{
				Description: "The height of the image in pixels.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "A short description of the watermark profile.",
				Computed:    true,
			},
			"opacity": schema.Float64Attribute{
				Description: "The translucency of the image. A value of `0.0` makes the image completely transparent, and `1.0` makes the image completely opaque. Note that if the image is already semi-transparent, setting this to `1.0` will not make the image completely opaque.",
				Computed:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"padding": schema.Float64Attribute{
				Description: "The whitespace between the adjacent edges (determined by position) of the video and the image. `0.0` indicates no padding, and `1.0` indicates a fully padded video width or length, as determined by the algorithm.",
				Computed:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"position": schema.StringAttribute{
				Description: "The location of the image. Valid positions are: `upperRight`, `upperLeft`, `lowerLeft`, `lowerRight`, and `center`. Note that `center` ignores the `padding` parameter.",
				Computed:    true,
			},
			"scale": schema.Float64Attribute{
				Description: "The size of the image relative to the overall size of the video. This parameter will adapt to horizontal and vertical videos automatically. `0.0` indicates no scaling (use the size of the image as-is), and `1.0 `fills the entire video.",
				Computed:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"size": schema.Float64Attribute{
				Description: "The size of the image in bytes.",
				Computed:    true,
			},
			"uid": schema.StringAttribute{
				Description: "The unique identifier for a watermark profile.",
				Computed:    true,
			},
			"width": schema.Int64Attribute{
				Description: "The width of the image in pixels.",
				Computed:    true,
			},
		},
	}
}

func (d *StreamWatermarkDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *StreamWatermarkDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
