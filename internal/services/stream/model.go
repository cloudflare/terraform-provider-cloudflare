// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamResultEnvelope struct {
	Result StreamModel `json:"result"`
}

type StreamModel struct {
	AccountID             types.String                                   `tfsdk:"account_id" path:"account_id,required"`
	Identifier            types.String                                   `tfsdk:"identifier" path:"identifier,optional"`
	Created               timetypes.RFC3339                              `tfsdk:"created" json:"created,computed" format:"date-time"`
	Creator               types.String                                   `tfsdk:"creator" json:"creator,computed"`
	Duration              types.Float64                                  `tfsdk:"duration" json:"duration,computed"`
	LiveInput             types.String                                   `tfsdk:"live_input" json:"liveInput,computed"`
	MaxDurationSeconds    types.Int64                                    `tfsdk:"max_duration_seconds" json:"maxDurationSeconds,computed"`
	Modified              timetypes.RFC3339                              `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Preview               types.String                                   `tfsdk:"preview" json:"preview,computed"`
	ReadyToStream         types.Bool                                     `tfsdk:"ready_to_stream" json:"readyToStream,computed"`
	ReadyToStreamAt       timetypes.RFC3339                              `tfsdk:"ready_to_stream_at" json:"readyToStreamAt,computed" format:"date-time"`
	RequireSignedURLs     types.Bool                                     `tfsdk:"require_signed_urls" json:"requireSignedURLs,computed"`
	ScheduledDeletion     timetypes.RFC3339                              `tfsdk:"scheduled_deletion" json:"scheduledDeletion,computed" format:"date-time"`
	Size                  types.Float64                                  `tfsdk:"size" json:"size,computed"`
	Thumbnail             types.String                                   `tfsdk:"thumbnail" json:"thumbnail,computed"`
	ThumbnailTimestampPct types.Float64                                  `tfsdk:"thumbnail_timestamp_pct" json:"thumbnailTimestampPct,computed"`
	UID                   types.String                                   `tfsdk:"uid" json:"uid,computed"`
	Uploaded              timetypes.RFC3339                              `tfsdk:"uploaded" json:"uploaded,computed" format:"date-time"`
	UploadExpiry          timetypes.RFC3339                              `tfsdk:"upload_expiry" json:"uploadExpiry,computed" format:"date-time"`
	AllowedOrigins        customfield.List[types.String]                 `tfsdk:"allowed_origins" json:"allowedOrigins,computed"`
	Input                 customfield.NestedObject[StreamInputModel]     `tfsdk:"input" json:"input,computed"`
	Playback              customfield.NestedObject[StreamPlaybackModel]  `tfsdk:"playback" json:"playback,computed"`
	Status                customfield.NestedObject[StreamStatusModel]    `tfsdk:"status" json:"status,computed"`
	Watermark             customfield.NestedObject[StreamWatermarkModel] `tfsdk:"watermark" json:"watermark,computed"`
	Meta                  jsontypes.Normalized                           `tfsdk:"meta" json:"meta,computed"`
}

func (m StreamModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m StreamModel) MarshalJSONForUpdate(state StreamModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type StreamInputModel struct {
	Height types.Int64 `tfsdk:"height" json:"height,computed"`
	Width  types.Int64 `tfsdk:"width" json:"width,computed"`
}

type StreamPlaybackModel struct {
	Dash types.String `tfsdk:"dash" json:"dash,computed"`
	Hls  types.String `tfsdk:"hls" json:"hls,computed"`
}

type StreamStatusModel struct {
	ErrorReasonCode types.String `tfsdk:"error_reason_code" json:"errorReasonCode,computed"`
	ErrorReasonText types.String `tfsdk:"error_reason_text" json:"errorReasonText,computed"`
	PctComplete     types.String `tfsdk:"pct_complete" json:"pctComplete,computed"`
	State           types.String `tfsdk:"state" json:"state,computed"`
}

type StreamWatermarkModel struct {
	Created        timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	DownloadedFrom types.String      `tfsdk:"downloaded_from" json:"downloadedFrom,computed"`
	Height         types.Int64       `tfsdk:"height" json:"height,computed"`
	Name           types.String      `tfsdk:"name" json:"name,computed"`
	Opacity        types.Float64     `tfsdk:"opacity" json:"opacity,computed"`
	Padding        types.Float64     `tfsdk:"padding" json:"padding,computed"`
	Position       types.String      `tfsdk:"position" json:"position,computed"`
	Scale          types.Float64     `tfsdk:"scale" json:"scale,computed"`
	Size           types.Float64     `tfsdk:"size" json:"size,computed"`
	UID            types.String      `tfsdk:"uid" json:"uid,computed"`
	Width          types.Int64       `tfsdk:"width" json:"width,computed"`
}
