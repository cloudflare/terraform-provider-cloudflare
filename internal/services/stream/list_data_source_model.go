// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/stream"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
)

type StreamsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[StreamsResultDataSourceModel] `json:"result,computed"`
}

type StreamsDataSourceModel struct {
	AccountID     types.String                                               `tfsdk:"account_id" path:"account_id,required"`
	After         timetypes.RFC3339                                          `tfsdk:"after" query:"after,optional" format:"date-time"`
	Before        timetypes.RFC3339                                          `tfsdk:"before" query:"before,optional" format:"date-time"`
	Creator       types.String                                               `tfsdk:"creator" query:"creator,optional"`
	End           timetypes.RFC3339                                          `tfsdk:"end" query:"end,optional" format:"date-time"`
	ID            types.String                                               `tfsdk:"id" query:"id,optional"`
	Limit         types.Int64                                                `tfsdk:"limit" query:"limit,optional"`
	LiveInputID   types.String                                               `tfsdk:"live_input_id" query:"live_input_id,optional"`
	Name          types.String                                               `tfsdk:"name" query:"name,optional"`
	Search        types.String                                               `tfsdk:"search" query:"search,optional"`
	Start         timetypes.RFC3339                                          `tfsdk:"start" query:"start,optional" format:"date-time"`
	Status        types.String                                               `tfsdk:"status" query:"status,optional"`
	Type          types.String                                               `tfsdk:"type" query:"type,optional"`
	VideoName     types.String                                               `tfsdk:"video_name" query:"video_name,optional"`
	Asc           types.Bool                                                 `tfsdk:"asc" query:"asc,computed_optional"`
	IncludeCounts types.Bool                                                 `tfsdk:"include_counts" query:"include_counts,computed_optional"`
	MaxItems      types.Int64                                                `tfsdk:"max_items"`
	Result        customfield.NestedObjectList[StreamsResultDataSourceModel] `tfsdk:"result"`
}

func (m *StreamsDataSourceModel) toListParams(_ context.Context) (params stream.StreamListParams, diags diag.Diagnostics) {
	mAfter, errs := m.After.ValueRFC3339Time()
	diags.Append(errs...)
	mBefore, errs := m.Before.ValueRFC3339Time()
	diags.Append(errs...)
	mEnd, errs := m.End.ValueRFC3339Time()
	diags.Append(errs...)
	mStart, errs := m.Start.ValueRFC3339Time()
	diags.Append(errs...)

	params = stream.StreamListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.ID.IsNull() {
		params.ID = cloudflare.F(m.ID.ValueString())
	}
	if !m.After.IsNull() {
		params.After = cloudflare.F(mAfter)
	}
	if !m.Asc.IsNull() {
		params.Asc = cloudflare.F(m.Asc.ValueBool())
	}
	if !m.Before.IsNull() {
		params.Before = cloudflare.F(mBefore)
	}
	if !m.Creator.IsNull() {
		params.Creator = cloudflare.F(m.Creator.ValueString())
	}
	if !m.End.IsNull() {
		params.End = cloudflare.F(mEnd)
	}
	if !m.IncludeCounts.IsNull() {
		params.IncludeCounts = cloudflare.F(m.IncludeCounts.ValueBool())
	}
	if !m.Limit.IsNull() {
		params.Limit = cloudflare.F(m.Limit.ValueInt64())
	}
	if !m.LiveInputID.IsNull() {
		params.LiveInputID = cloudflare.F(m.LiveInputID.ValueString())
	}
	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}
	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
	}
	if !m.Start.IsNull() {
		params.Start = cloudflare.F(mStart)
	}
	if !m.Status.IsNull() {
		params.Status = cloudflare.F(stream.StreamListParamsStatus(m.Status.ValueString()))
	}
	if !m.Type.IsNull() {
		params.Type = cloudflare.F(m.Type.ValueString())
	}
	//if !m.VideoName.IsNull() {
	//	params.VideoName = cloudflare.F(m.VideoName.ValueString())
	//}

	return
}

type StreamsResultDataSourceModel struct {
	AllowedOrigins        customfield.List[types.String]                                `tfsdk:"allowed_origins" json:"allowedOrigins,computed"`
	ClippedFrom           types.String                                                  `tfsdk:"clipped_from" json:"clippedFrom,computed"`
	Created               timetypes.RFC3339                                             `tfsdk:"created" json:"created,computed" format:"date-time"`
	Creator               types.String                                                  `tfsdk:"creator" json:"creator,computed"`
	Duration              types.Float64                                                 `tfsdk:"duration" json:"duration,computed"`
	Input                 customfield.NestedObject[StreamsInputDataSourceModel]         `tfsdk:"input" json:"input,computed"`
	LiveInput             types.String                                                  `tfsdk:"live_input" json:"liveInput,computed"`
	MaxDurationSeconds    types.Int64                                                   `tfsdk:"max_duration_seconds" json:"maxDurationSeconds,computed"`
	MaxSizeBytes          types.Int64                                                   `tfsdk:"max_size_bytes" json:"maxSizeBytes,computed"`
	Meta                  jsontypes.Normalized                                          `tfsdk:"meta" json:"meta,computed"`
	Modified              timetypes.RFC3339                                             `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Playback              customfield.NestedObject[StreamsPlaybackDataSourceModel]      `tfsdk:"playback" json:"playback,computed"`
	Preview               types.String                                                  `tfsdk:"preview" json:"preview,computed"`
	PublicDetails         customfield.NestedObject[StreamsPublicDetailsDataSourceModel] `tfsdk:"public_details" json:"publicDetails,computed"`
	ReadyToStream         types.Bool                                                    `tfsdk:"ready_to_stream" json:"readyToStream,computed"`
	ReadyToStreamAt       timetypes.RFC3339                                             `tfsdk:"ready_to_stream_at" json:"readyToStreamAt,computed" format:"date-time"`
	RequireSignedURLs     types.Bool                                                    `tfsdk:"require_signed_urls" json:"requireSignedURLs,computed"`
	ScheduledDeletion     timetypes.RFC3339                                             `tfsdk:"scheduled_deletion" json:"scheduledDeletion,computed" format:"date-time"`
	Size                  types.Float64                                                 `tfsdk:"size" json:"size,computed"`
	Status                customfield.NestedObject[StreamsStatusDataSourceModel]        `tfsdk:"status" json:"status,computed"`
	Thumbnail             types.String                                                  `tfsdk:"thumbnail" json:"thumbnail,computed"`
	ThumbnailTimestampPct types.Float64                                                 `tfsdk:"thumbnail_timestamp_pct" json:"thumbnailTimestampPct,computed"`
	UID                   types.String                                                  `tfsdk:"uid" json:"uid,computed"`
	Uploaded              timetypes.RFC3339                                             `tfsdk:"uploaded" json:"uploaded,computed" format:"date-time"`
	UploadExpiry          timetypes.RFC3339                                             `tfsdk:"upload_expiry" json:"uploadExpiry,computed" format:"date-time"`
	Watermark             customfield.NestedObject[StreamsWatermarkDataSourceModel]     `tfsdk:"watermark" json:"watermark,computed"`
}

type StreamsInputDataSourceModel struct {
	Height types.Int64 `tfsdk:"height" json:"height,computed"`
	Width  types.Int64 `tfsdk:"width" json:"width,computed"`
}

type StreamsPlaybackDataSourceModel struct {
	Dash types.String `tfsdk:"dash" json:"dash,computed"`
	Hls  types.String `tfsdk:"hls" json:"hls,computed"`
}

type StreamsPublicDetailsDataSourceModel struct {
	ChannelLink types.String `tfsdk:"channel_link" json:"channel_link,computed"`
	Logo        types.String `tfsdk:"logo" json:"logo,computed"`
	MediaID     types.Int64  `tfsdk:"media_id" json:"media_id,computed"`
	ShareLink   types.String `tfsdk:"share_link" json:"share_link,computed"`
	Title       types.String `tfsdk:"title" json:"title,computed"`
}

type StreamsStatusDataSourceModel struct {
	ErrorReasonCode types.String `tfsdk:"error_reason_code" json:"errorReasonCode,computed"`
	ErrorReasonText types.String `tfsdk:"error_reason_text" json:"errorReasonText,computed"`
	PctComplete     types.String `tfsdk:"pct_complete" json:"pctComplete,computed"`
	State           types.String `tfsdk:"state" json:"state,computed"`
}

type StreamsWatermarkDataSourceModel struct {
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
