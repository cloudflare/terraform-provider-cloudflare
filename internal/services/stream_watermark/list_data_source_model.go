// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_watermark

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/stream"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamWatermarksResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[StreamWatermarksResultDataSourceModel] `json:"result,computed"`
}

type StreamWatermarksDataSourceModel struct {
	AccountID types.String                                                        `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                         `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[StreamWatermarksResultDataSourceModel] `tfsdk:"result"`
}

func (m *StreamWatermarksDataSourceModel) toListParams(_ context.Context) (params stream.WatermarkListParams, diags diag.Diagnostics) {
	params = stream.WatermarkListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type StreamWatermarksResultDataSourceModel struct {
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
