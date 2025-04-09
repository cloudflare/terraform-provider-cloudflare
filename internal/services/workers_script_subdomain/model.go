// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script_subdomain

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersScriptSubdomainModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
ScriptName types.String `tfsdk:"script_name" path:"script_name,required"`
Enabled types.Bool `tfsdk:"enabled" json:"enabled,required"`
PreviewsEnabled types.Bool `tfsdk:"previews_enabled" json:"previews_enabled,optional"`
}

func (m WorkersScriptSubdomainModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m WorkersScriptSubdomainModel) MarshalJSONForUpdate(state WorkersScriptSubdomainModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}
