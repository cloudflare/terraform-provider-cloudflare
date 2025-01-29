// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/stream"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[StreamsResultDataSourceModel] `json:"result,computed"`
}

type StreamsDataSourceModel struct {
	AccountID     types.String                                               `tfsdk:"account_id" path:"account_id,required"`
	Creator       types.String                                               `tfsdk:"creator" query:"creator,optional"`
	End           timetypes.RFC3339                                          `tfsdk:"end" query:"end,optional" format:"date-time"`
	Search        types.String                                               `tfsdk:"search" query:"search,optional"`
	Start         timetypes.RFC3339                                          `tfsdk:"start" query:"start,optional" format:"date-time"`
	Status        types.String                                               `tfsdk:"status" query:"status,optional"`
	Type          types.String                                               `tfsdk:"type" query:"type,optional"`
	Asc           types.Bool                                                 `tfsdk:"asc" query:"asc,computed_optional"`
	IncludeCounts types.Bool                                                 `tfsdk:"include_counts" query:"include_counts,computed_optional"`
	MaxItems      types.Int64                                                `tfsdk:"max_items"`
	Result        customfield.NestedObjectList[StreamsResultDataSourceModel] `tfsdk:"result"`
}

func (m *StreamsDataSourceModel) toListParams(_ context.Context) (params stream.StreamListParams, diags diag.Diagnostics) {
	mEnd, errs := m.End.ValueRFC3339Time()
	diags.Append(errs...)
	mStart, errs := m.Start.ValueRFC3339Time()
	diags.Append(errs...)

	params = stream.StreamListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Asc.IsNull() {
		params.Asc = cloudflare.F(m.Asc.ValueBool())
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

	return
}

type StreamsResultDataSourceModel struct {
	AllowedOrigins        customfield.List[types.String]                            `tfsdk:"allowed_origins" json:"allowedOrigins,computed"`
	Created               timetypes.RFC3339                                         `tfsdk:"created" json:"created,computed" format:"date-time"`
	Creator               types.String                                              `tfsdk:"creator" json:"creator,computed"`
	Duration              types.Float64                                             `tfsdk:"duration" json:"duration,computed"`
	Input                 customfield.NestedObject[StreamsInputDataSourceModel]     `tfsdk:"input" json:"input,computed"`
	LiveInput             types.String                                              `tfsdk:"live_input" json:"liveInput,computed"`
	MaxDurationSeconds    types.Int64                                               `tfsdk:"max_duration_seconds" json:"maxDurationSeconds,computed"`
	Meta                  jsontypes.Normalized                                      `tfsdk:"meta" json:"meta,computed"`
	Modified              timetypes.RFC3339                                         `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Playback              customfield.NestedObject[StreamsPlaybackDataSourceModel]  `tfsdk:"playback" json:"playback,computed"`
	Preview               types.String                                              `tfsdk:"preview" json:"preview,computed"`
	ReadyToStream         types.Bool                                                `tfsdk:"ready_to_stream" json:"readyToStream,computed"`
	ReadyToStreamAt       timetypes.RFC3339                                         `tfsdk:"ready_to_stream_at" json:"readyToStreamAt,computed" format:"date-time"`
	RequireSignedURLs     types.Bool                                                `tfsdk:"require_signed_urls" json:"requireSignedURLs,computed"`
	ScheduledDeletion     timetypes.RFC3339                                         `tfsdk:"scheduled_deletion" json:"scheduledDeletion,computed" format:"date-time"`
	Size                  types.Float64                                             `tfsdk:"size" json:"size,computed"`
	Status                customfield.NestedObject[StreamsStatusDataSourceModel]    `tfsdk:"status" json:"status,computed"`
	Thumbnail             types.String                                              `tfsdk:"thumbnail" json:"thumbnail,computed"`
	ThumbnailTimestampPct types.Float64                                             `tfsdk:"thumbnail_timestamp_pct" json:"thumbnailTimestampPct,computed"`
	UID                   types.String                                              `tfsdk:"uid" json:"uid,computed"`
	Uploaded              timetypes.RFC3339                                         `tfsdk:"uploaded" json:"uploaded,computed" format:"date-time"`
	UploadExpiry          timetypes.RFC3339                                         `tfsdk:"upload_expiry" json:"uploadExpiry,computed" format:"date-time"`
	Watermark             customfield.NestedObject[StreamsWatermarkDataSourceModel] `tfsdk:"watermark" json:"watermark,computed"`
}

type StreamsInputDataSourceModel struct {
	Height types.Int64 `tfsdk:"height" json:"height,computed"`
	Width  types.Int64 `tfsdk:"width" json:"width,computed"`
}

type StreamsPlaybackDataSourceModel struct {
	Dash types.String `tfsdk:"dash" json:"dash,computed"`
	Hls  types.String `tfsdk:"hls" json:"hls,computed"`
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
