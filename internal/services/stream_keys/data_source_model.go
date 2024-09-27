// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_keys

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/stream"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamKeysResultDataSourceEnvelope struct {
	Result StreamKeysDataSourceModel `json:"result,computed"`
}

type StreamKeysDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}

func (m *StreamKeysDataSourceModel) toReadParams(_ context.Context) (params stream.KeyGetParams, diags diag.Diagnostics) {
	params = stream.KeyGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
