// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_live_input

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/stream"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamLiveInputResultDataSourceEnvelope struct {
	Result StreamLiveInputDataSourceModel `json:"result,computed"`
}

type StreamLiveInputDataSourceModel struct {
	AccountID                types.String                                  `tfsdk:"account_id" path:"account_id,required"`
	LiveInputIdentifier      types.String                                  `tfsdk:"live_input_identifier" path:"live_input_identifier,required"`
	Created                  timetypes.RFC3339                             `tfsdk:"created" json:"created,optional" format:"date-time"`
	DeleteRecordingAfterDays types.Float64                                 `tfsdk:"delete_recording_after_days" json:"deleteRecordingAfterDays,optional"`
	Modified                 timetypes.RFC3339                             `tfsdk:"modified" json:"modified,optional" format:"date-time"`
	Status                   types.String                                  `tfsdk:"status" json:"status,optional"`
	UID                      types.String                                  `tfsdk:"uid" json:"uid,optional"`
	Recording                *StreamLiveInputRecordingDataSourceModel      `tfsdk:"recording" json:"recording,optional"`
	Rtmps                    *StreamLiveInputRtmpsDataSourceModel          `tfsdk:"rtmps" json:"rtmps,optional"`
	RtmpsPlayback            *StreamLiveInputRtmpsPlaybackDataSourceModel  `tfsdk:"rtmps_playback" json:"rtmpsPlayback,optional"`
	Srt                      *StreamLiveInputSrtDataSourceModel            `tfsdk:"srt" json:"srt,optional"`
	SrtPlayback              *StreamLiveInputSrtPlaybackDataSourceModel    `tfsdk:"srt_playback" json:"srtPlayback,optional"`
	WebRtc                   *StreamLiveInputWebRtcDataSourceModel         `tfsdk:"web_rtc" json:"webRTC,optional"`
	WebRtcPlayback           *StreamLiveInputWebRtcPlaybackDataSourceModel `tfsdk:"web_rtc_playback" json:"webRTCPlayback,optional"`
	Meta                     jsontypes.Normalized                          `tfsdk:"meta" json:"meta,optional"`
}

func (m *StreamLiveInputDataSourceModel) toReadParams(_ context.Context) (params stream.LiveInputGetParams, diags diag.Diagnostics) {
	params = stream.LiveInputGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type StreamLiveInputRecordingDataSourceModel struct {
	AllowedOrigins      customfield.List[types.String] `tfsdk:"allowed_origins" json:"allowedOrigins,computed"`
	HideLiveViewerCount types.Bool                     `tfsdk:"hide_live_viewer_count" json:"hideLiveViewerCount,computed"`
	Mode                types.String                   `tfsdk:"mode" json:"mode,computed"`
	RequireSignedURLs   types.Bool                     `tfsdk:"require_signed_urls" json:"requireSignedURLs,computed"`
	TimeoutSeconds      types.Int64                    `tfsdk:"timeout_seconds" json:"timeoutSeconds,computed"`
}

type StreamLiveInputRtmpsDataSourceModel struct {
	StreamKey types.String `tfsdk:"stream_key" json:"streamKey,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
}

type StreamLiveInputRtmpsPlaybackDataSourceModel struct {
	StreamKey types.String `tfsdk:"stream_key" json:"streamKey,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
}

type StreamLiveInputSrtDataSourceModel struct {
	Passphrase types.String `tfsdk:"passphrase" json:"passphrase,computed"`
	StreamID   types.String `tfsdk:"stream_id" json:"streamId,computed"`
	URL        types.String `tfsdk:"url" json:"url,computed"`
}

type StreamLiveInputSrtPlaybackDataSourceModel struct {
	Passphrase types.String `tfsdk:"passphrase" json:"passphrase,computed"`
	StreamID   types.String `tfsdk:"stream_id" json:"streamId,computed"`
	URL        types.String `tfsdk:"url" json:"url,computed"`
}

type StreamLiveInputWebRtcDataSourceModel struct {
	URL types.String `tfsdk:"url" json:"url,computed"`
}

type StreamLiveInputWebRtcPlaybackDataSourceModel struct {
	URL types.String `tfsdk:"url" json:"url,computed"`
}
