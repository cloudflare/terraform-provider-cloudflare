// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_download

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/stream"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamDownloadResultDataSourceEnvelope struct {
	Result StreamDownloadDataSourceModel `json:"result,computed"`
}

type StreamDownloadDataSourceModel struct {
	AccountID  types.String `tfsdk:"account_id" path:"account_id,required"`
	Identifier types.String `tfsdk:"identifier" path:"identifier,required"`
}

func (m *StreamDownloadDataSourceModel) toReadParams(_ context.Context) (params stream.DownloadGetParams, diags diag.Diagnostics) {
	params = stream.DownloadGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
