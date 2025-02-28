// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_audio_track

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*StreamAudioTrackDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The account identifier tag.",
				Required:    true,
			},
			"identifier": schema.StringAttribute{
				Description: "A Cloudflare-generated unique identifier for a media item.",
				Required:    true,
			},
			"default": schema.BoolAttribute{
				Description: "Denotes whether the audio track will be played by default in a player.",
				Computed:    true,
			},
			"label": schema.StringAttribute{
				Description: "A string to uniquely identify the track amongst other audio track labels for the specified video.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Specifies the processing status of the video.\nAvailable values: \"queued\", \"ready\", \"error\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"queued",
						"ready",
						"error",
					),
				},
			},
			"uid": schema.StringAttribute{
				Description: "A Cloudflare-generated unique identifier for a media item.",
				Computed:    true,
			},
		},
	}
}

func (d *StreamAudioTrackDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *StreamAudioTrackDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
