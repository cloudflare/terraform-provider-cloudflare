// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_watermark

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/stream"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamWatermarkResultDataSourceEnvelope struct {
	Result StreamWatermarkDataSourceModel `json:"result,computed"`
}

type StreamWatermarkDataSourceModel struct {
	AccountID      types.String      `tfsdk:"account_id" path:"account_id,required"`
	Identifier     types.String      `tfsdk:"identifier" path:"identifier,required"`
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

func (m *StreamWatermarkDataSourceModel) toReadParams(_ context.Context) (params stream.WatermarkGetParams, diags diag.Diagnostics) {
	params = stream.WatermarkGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
