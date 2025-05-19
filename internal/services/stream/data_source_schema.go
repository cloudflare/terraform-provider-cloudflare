// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*StreamDataSource)(nil)

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
			"created": schema.StringAttribute{
				Description: "The date and time the media item was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"creator": schema.StringAttribute{
				Description: "A user-defined identifier for the media creator.",
				Computed:    true,
			},
			"duration": schema.Float64Attribute{
				Description: "The duration of the video in seconds. A value of `-1` means the duration is unknown. The duration becomes available after the upload and before the video is ready.",
				Computed:    true,
			},
			"live_input": schema.StringAttribute{
				Description: "The live input ID used to upload a video with Stream Live.",
				Computed:    true,
			},
			"max_duration_seconds": schema.Int64Attribute{
				Description: "The maximum duration in seconds for a video upload. Can be set for a video that is not yet uploaded to limit its duration. Uploads that exceed the specified duration will fail during processing. A value of `-1` means the value is unknown.",
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 21600),
				},
			},
			"modified": schema.StringAttribute{
				Description: "The date and time the media item was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"preview": schema.StringAttribute{
				Description: "The video's preview page URI. This field is omitted until encoding is complete.",
				Computed:    true,
			},
			"ready_to_stream": schema.BoolAttribute{
				Description: "Indicates whether the video is playable. The field is empty if the video is not ready for viewing or the live stream is still in progress.",
				Computed:    true,
			},
			"ready_to_stream_at": schema.StringAttribute{
				Description: "Indicates the time at which the video became playable. The field is empty if the video is not ready for viewing or the live stream is still in progress.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"require_signed_urls": schema.BoolAttribute{
				Description: "Indicates whether the video can be a accessed using the UID. When set to `true`, a signed token must be generated with a signing key to view the video.",
				Computed:    true,
			},
			"scheduled_deletion": schema.StringAttribute{
				Description: "Indicates the date and time at which the video will be deleted. Omit the field to indicate no change, or include with a `null` value to remove an existing scheduled deletion. If specified, must be at least 30 days from upload time.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"size": schema.Float64Attribute{
				Description: "The size of the media item in bytes.",
				Computed:    true,
			},
			"thumbnail": schema.StringAttribute{
				Description: "The media item's thumbnail URI. This field is omitted until encoding is complete.",
				Computed:    true,
			},
			"thumbnail_timestamp_pct": schema.Float64Attribute{
				Description: "The timestamp for a thumbnail image calculated as a percentage value of the video's duration. To convert from a second-wise timestamp to a percentage, divide the desired timestamp by the total duration of the video.  If this value is not set, the default thumbnail image is taken from 0s of the video.",
				Computed:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"uid": schema.StringAttribute{
				Description: "A Cloudflare-generated unique identifier for a media item.",
				Computed:    true,
			},
			"upload_expiry": schema.StringAttribute{
				Description: "The date and time when the video upload URL is no longer valid for direct user uploads.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"uploaded": schema.StringAttribute{
				Description: "The date and time the media item was uploaded.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"allowed_origins": schema.ListAttribute{
				Description: "Lists the origins allowed to display the video. Enter allowed origin domains in an array and use `*` for wildcard subdomains. Empty arrays allow the video to be viewed on any origin.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"input": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[StreamInputDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"height": schema.Int64Attribute{
						Description: "The video height in pixels. A value of `-1` means the height is unknown. The value becomes available after the upload and before the video is ready.",
						Computed:    true,
					},
					"width": schema.Int64Attribute{
						Description: "The video width in pixels. A value of `-1` means the width is unknown. The value becomes available after the upload and before the video is ready.",
						Computed:    true,
					},
				},
			},
			"playback": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[StreamPlaybackDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"dash": schema.StringAttribute{
						Description: "DASH Media Presentation Description for the video.",
						Computed:    true,
					},
					"hls": schema.StringAttribute{
						Description: "The HLS manifest for the video.",
						Computed:    true,
					},
				},
			},
			"status": schema.SingleNestedAttribute{
				Description: "Specifies a detailed status for a video. If the `state` is `inprogress` or `error`, the `step` field returns `encoding` or `manifest`. If the `state` is `inprogress`, `pctComplete` returns a number between 0 and 100 to indicate the approximate percent of completion. If the `state` is `error`, `errorReasonCode` and `errorReasonText` provide additional details.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[StreamStatusDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"error_reason_code": schema.StringAttribute{
						Description: "Specifies why the video failed to encode. This field is empty if the video is not in an `error` state. Preferred for programmatic use.",
						Computed:    true,
					},
					"error_reason_text": schema.StringAttribute{
						Description: "Specifies why the video failed to encode using a human readable error message in English. This field is empty if the video is not in an `error` state.",
						Computed:    true,
					},
					"pct_complete": schema.StringAttribute{
						Description: "Indicates the size of the entire upload in bytes. The value must be a non-negative integer.",
						Computed:    true,
					},
					"state": schema.StringAttribute{
						Description: "Specifies the processing status for all quality levels for a video.\nAvailable values: \"pendingupload\", \"downloading\", \"queued\", \"inprogress\", \"ready\", \"error\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"pendingupload",
								"downloading",
								"queued",
								"inprogress",
								"ready",
								"error",
							),
						},
					},
				},
			},
			"watermark": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[StreamWatermarkDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"created": schema.StringAttribute{
						Description: "The date and a time a watermark profile was created.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"downloaded_from": schema.StringAttribute{
						Description: "The source URL for a downloaded image. If the watermark profile was created via direct upload, this field is null.",
						Computed:    true,
					},
					"height": schema.Int64Attribute{
						Description: "The height of the image in pixels.",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "A short description of the watermark profile.",
						Computed:    true,
					},
					"opacity": schema.Float64Attribute{
						Description: "The translucency of the image. A value of `0.0` makes the image completely transparent, and `1.0` makes the image completely opaque. Note that if the image is already semi-transparent, setting this to `1.0` will not make the image completely opaque.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 1),
						},
					},
					"padding": schema.Float64Attribute{
						Description: "The whitespace between the adjacent edges (determined by position) of the video and the image. `0.0` indicates no padding, and `1.0` indicates a fully padded video width or length, as determined by the algorithm.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 1),
						},
					},
					"position": schema.StringAttribute{
						Description: "The location of the image. Valid positions are: `upperRight`, `upperLeft`, `lowerLeft`, `lowerRight`, and `center`. Note that `center` ignores the `padding` parameter.",
						Computed:    true,
					},
					"scale": schema.Float64Attribute{
						Description: "The size of the image relative to the overall size of the video. This parameter will adapt to horizontal and vertical videos automatically. `0.0` indicates no scaling (use the size of the image as-is), and `1.0 `fills the entire video.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 1),
						},
					},
					"size": schema.Float64Attribute{
						Description: "The size of the image in bytes.",
						Computed:    true,
					},
					"uid": schema.StringAttribute{
						Description: "The unique identifier for a watermark profile.",
						Computed:    true,
					},
					"width": schema.Int64Attribute{
						Description: "The width of the image in pixels.",
						Computed:    true,
					},
				},
			},
			"meta": schema.StringAttribute{
				Description: "A user modifiable key-value store used to reference other systems of record for managing videos.",
				Computed:    true,
				CustomType:  jsontypes.NormalizedType{},
			},
		},
	}
}

func (d *StreamDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *StreamDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
