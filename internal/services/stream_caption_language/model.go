// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_caption_language

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamCaptionLanguageResultEnvelope struct {
Result StreamCaptionLanguageModel `json:"result"`
}

type StreamCaptionLanguageModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
Identifier types.String `tfsdk:"identifier" path:"identifier,required"`
Language types.String `tfsdk:"language" path:"language,required"`
File types.String `tfsdk:"file" json:"file,optional"`
Generated types.Bool `tfsdk:"generated" json:"generated,computed"`
Label types.String `tfsdk:"label" json:"label,computed"`
Status types.String `tfsdk:"status" json:"status,computed"`
}

func (m StreamCaptionLanguageModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m StreamCaptionLanguageModel) MarshalJSONForUpdate(state StreamCaptionLanguageModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}
