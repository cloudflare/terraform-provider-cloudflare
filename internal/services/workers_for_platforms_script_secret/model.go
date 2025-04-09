// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_script_secret

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersForPlatformsScriptSecretResultEnvelope struct {
Result WorkersForPlatformsScriptSecretModel `json:"result"`
}

type WorkersForPlatformsScriptSecretModel struct {
ID types.String `tfsdk:"id" json:"-,computed"`
Name types.String `tfsdk:"name" json:"name,required"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
DispatchNamespace types.String `tfsdk:"dispatch_namespace" path:"dispatch_namespace,required"`
ScriptName types.String `tfsdk:"script_name" path:"script_name,required"`
Text types.String `tfsdk:"text" json:"text,optional"`
Type types.String `tfsdk:"type" json:"type,optional"`
}

func (m WorkersForPlatformsScriptSecretModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m WorkersForPlatformsScriptSecretModel) MarshalJSONForUpdate(state WorkersForPlatformsScriptSecretModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}
