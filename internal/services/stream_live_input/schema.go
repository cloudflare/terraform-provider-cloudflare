// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_live_input

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*StreamLiveInputResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"live_input_identifier": schema.StringAttribute{
				Description:   "A unique identifier for a live input.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"default_creator": schema.StringAttribute{
				Description: "Sets the creator ID asssociated with this live input.",
				Optional:    true,
			},
			"delete_recording_after_days": schema.Float64Attribute{
				Description: "Indicates the number of days after which the live inputs recordings will be deleted. When a stream completes and the recording is ready, the value is used to calculate a scheduled deletion date for that recording. Omit the field to indicate no change, or include with a `null` value to remove an existing scheduled deletion.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(30),
				},
			},
			"meta": schema.StringAttribute{
				Description: "A user modifiable key-value store used to reference other systems of record for managing live inputs.",
				Optional:    true,
				CustomType:  jsontypes.NormalizedType{},
			},
			"recording": schema.SingleNestedAttribute{
				Description: "Records the input to a Cloudflare Stream video. Behavior depends on the mode. In most cases, the video will initially be viewable as a live video and transition to on-demand after a condition is satisfied.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[StreamLiveInputRecordingModel](ctx),
				Attributes: map[string]schema.Attribute{
					"allowed_origins": schema.ListAttribute{
						Description: "Lists the origins allowed to display videos created with this input. Enter allowed origin domains in an array and use `*` for wildcard subdomains. An empty array allows videos to be viewed on any origin.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"hide_live_viewer_count": schema.BoolAttribute{
						Description: "Disables reporting the number of live viewers when this property is set to `true`.",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
					},
					"mode": schema.StringAttribute{
						Description: "Specifies the recording behavior for the live input. Set this value to `off` to prevent a recording. Set the value to `automatic` to begin a recording and transition to on-demand after Stream Live stops receiving input.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("off", "automatic"),
						},
						Default: stringdefault.StaticString("off"),
					},
					"require_signed_urls": schema.BoolAttribute{
						Description: "Indicates if a video using the live input has the `requireSignedURLs` property set. Also enforces access controls on any video recording of the livestream with the live input.",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
					},
					"timeout_seconds": schema.Int64Attribute{
						Description: "Determines the amount of time a live input configured in `automatic` mode should wait before a recording transitions from live to on-demand. `0` is recommended for most use cases and indicates the platform default should be used.",
						Computed:    true,
						Optional:    true,
						Default:     int64default.StaticInt64(0),
					},
				},
			},
			"created": schema.StringAttribute{
				Description: "The date and time the live input was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified": schema.StringAttribute{
				Description: "The date and time the live input was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"status": schema.StringAttribute{
				Description: "The connection status of a live input.",
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
			"rtmps": schema.SingleNestedAttribute{
				Description: "Details for streaming to an live input using RTMPS.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[StreamLiveInputRtmpsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"stream_key": schema.StringAttribute{
						Description: "The secret key to use when streaming via RTMPS to a live input.",
						Computed:    true,
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
				CustomType:  customfield.NewNestedObjectType[StreamLiveInputRtmpsPlaybackModel](ctx),
				Attributes: map[string]schema.Attribute{
					"stream_key": schema.StringAttribute{
						Description: "The secret key to use for playback via RTMPS.",
						Computed:    true,
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
				CustomType:  customfield.NewNestedObjectType[StreamLiveInputSrtModel](ctx),
				Attributes: map[string]schema.Attribute{
					"passphrase": schema.StringAttribute{
						Description: "The secret key to use when streaming via SRT to a live input.",
						Computed:    true,
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
				CustomType:  customfield.NewNestedObjectType[StreamLiveInputSrtPlaybackModel](ctx),
				Attributes: map[string]schema.Attribute{
					"passphrase": schema.StringAttribute{
						Description: "The secret key to use for playback via SRT.",
						Computed:    true,
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
				CustomType:  customfield.NewNestedObjectType[StreamLiveInputWebRtcModel](ctx),
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
				CustomType:  customfield.NewNestedObjectType[StreamLiveInputWebRtcPlaybackModel](ctx),
				Attributes: map[string]schema.Attribute{
					"url": schema.StringAttribute{
						Description: "The URL used to play live video over WebRTC.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (r *StreamLiveInputResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *StreamLiveInputResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
