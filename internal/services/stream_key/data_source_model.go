// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_key

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/stream"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamKeyResultDataSourceEnvelope struct {
	Result StreamKeyDataSourceModel `json:"result,computed"`
}

type StreamKeyDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}

func (m *StreamKeyDataSourceModel) toReadParams(_ context.Context) (params stream.KeyGetParams, diags diag.Diagnostics) {
	params = stream.KeyGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
