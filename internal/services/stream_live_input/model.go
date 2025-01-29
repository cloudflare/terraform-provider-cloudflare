// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_live_input

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamLiveInputResultEnvelope struct {
	Result StreamLiveInputModel `json:"result"`
}

type StreamLiveInputModel struct {
	AccountID                types.String                                                 `tfsdk:"account_id" path:"account_id,required"`
	LiveInputIdentifier      types.String                                                 `tfsdk:"live_input_identifier" path:"live_input_identifier,optional"`
	DefaultCreator           types.String                                                 `tfsdk:"default_creator" json:"defaultCreator,optional"`
	DeleteRecordingAfterDays types.Float64                                                `tfsdk:"delete_recording_after_days" json:"deleteRecordingAfterDays,optional"`
	Meta                     jsontypes.Normalized                                         `tfsdk:"meta" json:"meta,optional"`
	Recording                customfield.NestedObject[StreamLiveInputRecordingModel]      `tfsdk:"recording" json:"recording,computed_optional"`
	Created                  timetypes.RFC3339                                            `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified                 timetypes.RFC3339                                            `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Status                   types.String                                                 `tfsdk:"status" json:"status,computed"`
	UID                      types.String                                                 `tfsdk:"uid" json:"uid,computed"`
	Rtmps                    customfield.NestedObject[StreamLiveInputRtmpsModel]          `tfsdk:"rtmps" json:"rtmps,computed"`
	RtmpsPlayback            customfield.NestedObject[StreamLiveInputRtmpsPlaybackModel]  `tfsdk:"rtmps_playback" json:"rtmpsPlayback,computed"`
	Srt                      customfield.NestedObject[StreamLiveInputSrtModel]            `tfsdk:"srt" json:"srt,computed"`
	SrtPlayback              customfield.NestedObject[StreamLiveInputSrtPlaybackModel]    `tfsdk:"srt_playback" json:"srtPlayback,computed"`
	WebRtc                   customfield.NestedObject[StreamLiveInputWebRtcModel]         `tfsdk:"web_rtc" json:"webRTC,computed"`
	WebRtcPlayback           customfield.NestedObject[StreamLiveInputWebRtcPlaybackModel] `tfsdk:"web_rtc_playback" json:"webRTCPlayback,computed"`
}

func (m StreamLiveInputModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m StreamLiveInputModel) MarshalJSONForUpdate(state StreamLiveInputModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type StreamLiveInputRecordingModel struct {
	AllowedOrigins      *[]types.String `tfsdk:"allowed_origins" json:"allowedOrigins,optional"`
	HideLiveViewerCount types.Bool      `tfsdk:"hide_live_viewer_count" json:"hideLiveViewerCount,computed_optional"`
	Mode                types.String    `tfsdk:"mode" json:"mode,computed_optional"`
	RequireSignedURLs   types.Bool      `tfsdk:"require_signed_urls" json:"requireSignedURLs,computed_optional"`
	TimeoutSeconds      types.Int64     `tfsdk:"timeout_seconds" json:"timeoutSeconds,computed_optional"`
}

type StreamLiveInputRtmpsModel struct {
	StreamKey types.String `tfsdk:"stream_key" json:"streamKey,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
}

type StreamLiveInputRtmpsPlaybackModel struct {
	StreamKey types.String `tfsdk:"stream_key" json:"streamKey,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
}

type StreamLiveInputSrtModel struct {
	Passphrase types.String `tfsdk:"passphrase" json:"passphrase,computed"`
	StreamID   types.String `tfsdk:"stream_id" json:"streamId,computed"`
	URL        types.String `tfsdk:"url" json:"url,computed"`
}

type StreamLiveInputSrtPlaybackModel struct {
	Passphrase types.String `tfsdk:"passphrase" json:"passphrase,computed"`
	StreamID   types.String `tfsdk:"stream_id" json:"streamId,computed"`
	URL        types.String `tfsdk:"url" json:"url,computed"`
}

type StreamLiveInputWebRtcModel struct {
	URL types.String `tfsdk:"url" json:"url,computed"`
}

type StreamLiveInputWebRtcPlaybackModel struct {
	URL types.String `tfsdk:"url" json:"url,computed"`
}
