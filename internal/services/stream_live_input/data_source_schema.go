// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_live_input

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*StreamLiveInputDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"live_input_identifier": schema.StringAttribute{
				Description: "A unique identifier for a live input.",
				Required:    true,
			},
			"created": schema.StringAttribute{
				Description: "The date and time the live input was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"delete_recording_after_days": schema.Float64Attribute{
				Description: "Indicates the number of days after which the live inputs recordings will be deleted. When a stream completes and the recording is ready, the value is used to calculate a scheduled deletion date for that recording. Omit the field to indicate no change, or include with a `null` value to remove an existing scheduled deletion.",
				Computed:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(30),
				},
			},
			"modified": schema.StringAttribute{
				Description: "The date and time the live input was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"status": schema.StringAttribute{
				Description: "The connection status of a live input.\nAvailable values: \"connected\", \"reconnected\", \"reconnecting\", \"client_disconnect\", \"ttl_exceeded\", \"failed_to_connect\", \"failed_to_reconnect\", \"new_configuration_accepted\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"connected",
						"reconnected",
						"reconnecting",
						"client_disconnect",
						"ttl_exceeded",
						"failed_to_connect",
						"failed_to_reconnect",
						"new_configuration_accepted",
					),
				},
			},
			"uid": schema.StringAttribute{
				Description: "A unique identifier for a live input.",
				Computed:    true,
			},
			"recording": schema.SingleNestedAttribute{
				Description: "Records the input to a Cloudflare Stream video. Behavior depends on the mode. In most cases, the video will initially be viewable as a live video and transition to on-demand after a condition is satisfied.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[StreamLiveInputRecordingDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"allowed_origins": schema.ListAttribute{
						Description: "Lists the origins allowed to display videos created with this input. Enter allowed origin domains in an array and use `*` for wildcard subdomains. An empty array allows videos to be viewed on any origin.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"hide_live_viewer_count": schema.BoolAttribute{
						Description: "Disables reporting the number of live viewers when this property is set to `true`.",
						Computed:    true,
					},
					"mode": schema.StringAttribute{
						Description: "Specifies the recording behavior for the live input. Set this value to `off` to prevent a recording. Set the value to `automatic` to begin a recording and transition to on-demand after Stream Live stops receiving input.\nAvailable values: \"off\", \"automatic\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("off", "automatic"),
						},
					},
					"require_signed_urls": schema.BoolAttribute{
						Description: "Indicates if a video using the live input has the `requireSignedURLs` property set. Also enforces access controls on any video recording of the livestream with the live input.",
						Computed:    true,
					},
					"timeout_seconds": schema.Int64Attribute{
						Description: "Determines the amount of time a live input configured in `automatic` mode should wait before a recording transitions from live to on-demand. `0` is recommended for most use cases and indicates the platform default should be used.",
						Computed:    true,
					},
				},
			},
			"rtmps": schema.SingleNestedAttribute{
				Description: "Details for streaming to an live input using RTMPS.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[StreamLiveInputRtmpsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"stream_key": schema.StringAttribute{
						Description: "The secret key to use when streaming via RTMPS to a live input.",
						Computed:    true,
						Sensitive:   true,
					},
					"url": schema.StringAttribute{
						Description: "The RTMPS URL you provide to the broadcaster, which they stream live video to.",
						Computed:    true,
					},
				},
			},
			"rtmps_playback": schema.SingleNestedAttribute{
				Description: "Details for playback from an live input using RTMPS.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[StreamLiveInputRtmpsPlaybackDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"stream_key": schema.StringAttribute{
						Description: "The secret key to use for playback via RTMPS.",
						Computed:    true,
						Sensitive:   true,
					},
					"url": schema.StringAttribute{
						Description: "The URL used to play live video over RTMPS.",
						Computed:    true,
					},
				},
			},
			"srt": schema.SingleNestedAttribute{
				Description: "Details for streaming to a live input using SRT.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[StreamLiveInputSrtDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"passphrase": schema.StringAttribute{
						Description: "The secret key to use when streaming via SRT to a live input.",
						Computed:    true,
						Sensitive:   true,
					},
					"stream_id": schema.StringAttribute{
						Description: "The identifier of the live input to use when streaming via SRT.",
						Computed:    true,
					},
					"url": schema.StringAttribute{
						Description: "The SRT URL you provide to the broadcaster, which they stream live video to.",
						Computed:    true,
					},
				},
			},
			"srt_playback": schema.SingleNestedAttribute{
				Description: "Details for playback from an live input using SRT.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[StreamLiveInputSrtPlaybackDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"passphrase": schema.StringAttribute{
						Description: "The secret key to use for playback via SRT.",
						Computed:    true,
						Sensitive:   true,
					},
					"stream_id": schema.StringAttribute{
						Description: "The identifier of the live input to use for playback via SRT.",
						Computed:    true,
					},
					"url": schema.StringAttribute{
						Description: "The URL used to play live video over SRT.",
						Computed:    true,
					},
				},
			},
			"web_rtc": schema.SingleNestedAttribute{
				Description: "Details for streaming to a live input using WebRTC.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[StreamLiveInputWebRtcDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"url": schema.StringAttribute{
						Description: "The WebRTC URL you provide to the broadcaster, which they stream live video to.",
						Computed:    true,
					},
				},
			},
			"web_rtc_playback": schema.SingleNestedAttribute{
				Description: "Details for playback from a live input using WebRTC.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[StreamLiveInputWebRtcPlaybackDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"url": schema.StringAttribute{
						Description: "The URL used to play live video over WebRTC.",
						Computed:    true,
					},
				},
			},
			"meta": schema.StringAttribute{
				Description: "A user modifiable key-value store used to reference other systems of record for managing live inputs.",
				Computed:    true,
				CustomType:  jsontypes.NormalizedType{},
			},
		},
	}
}

func (d *StreamLiveInputDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *StreamLiveInputDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
