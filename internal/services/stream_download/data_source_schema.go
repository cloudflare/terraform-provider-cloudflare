// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_download

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*StreamDownloadDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"identifier": schema.StringAttribute{
				Description: "A Cloudflare-generated unique identifier for a media item.",
				Required:    true,
			},
			"audio": schema.SingleNestedAttribute{
				Description: "The audio-only download. Only present if this download type has been created.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[StreamDownloadAudioDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"percent_complete": schema.Float64Attribute{
						Description: "Indicates the progress as a percentage between 0 and 100.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 100),
						},
					},
					"status": schema.StringAttribute{
						Description: "The status of a generated download.\nAvailable values: \"ready\", \"inprogress\", \"error\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"ready",
								"inprogress",
								"error",
							),
						},
					},
					"url": schema.StringAttribute{
						Description: "The URL to access the generated download.",
						Computed:    true,
					},
				},
			},
			"default": schema.SingleNestedAttribute{
				Description: "The default video download. Only present if this download type has been created.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[StreamDownloadDefaultDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"percent_complete": schema.Float64Attribute{
						Description: "Indicates the progress as a percentage between 0 and 100.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 100),
						},
					},
					"status": schema.StringAttribute{
						Description: "The status of a generated download.\nAvailable values: \"ready\", \"inprogress\", \"error\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"ready",
								"inprogress",
								"error",
							),
						},
					},
					"url": schema.StringAttribute{
						Description: "The URL to access the generated download.",
						Computed:    true,
					},
				},
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
