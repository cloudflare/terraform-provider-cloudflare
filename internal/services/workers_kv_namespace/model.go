// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv_namespace

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersKVNamespaceResultEnvelope struct {
Result WorkersKVNamespaceModel `json:"result"`
}

type WorkersKVNamespaceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
Title types.String `tfsdk:"title" json:"title,required"`
SupportsURLEncoding types.Bool `tfsdk:"supports_url_encoding" json:"supports_url_encoding,computed"`
}

func (m WorkersKVNamespaceModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m WorkersKVNamespaceModel) MarshalJSONForUpdate(state WorkersKVNamespaceModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}
