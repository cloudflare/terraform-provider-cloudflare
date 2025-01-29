// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_watermark

import (
	"bytes"
	"mime/multipart"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apiform"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamWatermarkResultEnvelope struct {
	Result StreamWatermarkModel `json:"result"`
}

type StreamWatermarkModel struct {
	AccountID      types.String      `tfsdk:"account_id" path:"account_id,required"`
	Identifier     types.String      `tfsdk:"identifier" path:"identifier,optional"`
	File           types.String      `tfsdk:"file" json:"file,required"`
	Name           types.String      `tfsdk:"name" json:"name,computed_optional"`
	Opacity        types.Float64     `tfsdk:"opacity" json:"opacity,computed_optional"`
	Padding        types.Float64     `tfsdk:"padding" json:"padding,computed_optional"`
	Position       types.String      `tfsdk:"position" json:"position,computed_optional"`
	Scale          types.Float64     `tfsdk:"scale" json:"scale,computed_optional"`
	Created        timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	DownloadedFrom types.String      `tfsdk:"downloaded_from" json:"downloadedFrom,computed"`
	Height         types.Int64       `tfsdk:"height" json:"height,computed"`
	Size           types.Float64     `tfsdk:"size" json:"size,computed"`
	UID            types.String      `tfsdk:"uid" json:"uid,computed"`
	Width          types.Int64       `tfsdk:"width" json:"width,computed"`
}

func (r StreamWatermarkModel) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
