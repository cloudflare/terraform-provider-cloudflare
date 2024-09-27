// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_audio_track

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
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
		},
	}
}

func (d *StreamAudioTrackDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *StreamAudioTrackDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
