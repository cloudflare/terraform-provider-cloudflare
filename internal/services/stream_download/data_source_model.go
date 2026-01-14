// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_download

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/stream"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamDownloadResultDataSourceEnvelope struct {
	Result StreamDownloadDataSourceModel `json:"result,computed"`
}

type StreamDownloadDataSourceModel struct {
	AccountID  types.String                                                   `tfsdk:"account_id" path:"account_id,required"`
	Identifier types.String                                                   `tfsdk:"identifier" path:"identifier,required"`
	Audio      customfield.NestedObject[StreamDownloadAudioDataSourceModel]   `tfsdk:"audio" json:"audio,computed"`
	Default    customfield.NestedObject[StreamDownloadDefaultDataSourceModel] `tfsdk:"default" json:"default,computed"`
}

func (m *StreamDownloadDataSourceModel) toReadParams(_ context.Context) (params stream.DownloadGetParams, diags diag.Diagnostics) {
	params = stream.DownloadGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type StreamDownloadAudioDataSourceModel struct {
	PercentComplete types.Float64 `tfsdk:"percent_complete" json:"percentComplete,computed"`
	Status          types.String  `tfsdk:"status" json:"status,computed"`
	URL             types.String  `tfsdk:"url" json:"url,computed"`
}

type StreamDownloadDefaultDataSourceModel struct {
	PercentComplete types.Float64 `tfsdk:"percent_complete" json:"percentComplete,computed"`
	Status          types.String  `tfsdk:"status" json:"status,computed"`
	URL             types.String  `tfsdk:"url" json:"url,computed"`
}
