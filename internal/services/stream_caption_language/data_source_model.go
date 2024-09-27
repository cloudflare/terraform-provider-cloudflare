// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_caption_language

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/stream"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamCaptionLanguageResultDataSourceEnvelope struct {
	Result StreamCaptionLanguageDataSourceModel `json:"result,computed"`
}

type StreamCaptionLanguageDataSourceModel struct {
	AccountID  types.String `tfsdk:"account_id" path:"account_id,required"`
	Identifier types.String `tfsdk:"identifier" path:"identifier,required"`
	Language   types.String `tfsdk:"language" path:"language,computed"`
	Generated  types.Bool   `tfsdk:"generated" json:"generated,optional"`
	Label      types.String `tfsdk:"label" json:"label,optional"`
	Status     types.String `tfsdk:"status" json:"status,optional"`
}

func (m *StreamCaptionLanguageDataSourceModel) toReadParams(_ context.Context) (params stream.CaptionLanguageGetParams, diags diag.Diagnostics) {
	params = stream.CaptionLanguageGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
