// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_caption_language

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*StreamCaptionLanguageDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"identifier": schema.StringAttribute{
				Description: "A Cloudflare-generated unique identifier for a media item.",
				Required:    true,
			},
			"language": schema.StringAttribute{
				Description: "The language tag in BCP 47 format.",
				Required:    true,
			},
			"generated": schema.BoolAttribute{
				Description: "Whether the caption was generated via AI.",
				Computed:    true,
			},
			"label": schema.StringAttribute{
				Description: "The language label displayed in the native language to users.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "The status of a generated caption.\navailable values: \"ready\", \"inprogress\", \"error\"",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"ready",
						"inprogress",
						"error",
					),
				},
			},
		},
	}
}

func (d *StreamCaptionLanguageDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *StreamCaptionLanguageDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
