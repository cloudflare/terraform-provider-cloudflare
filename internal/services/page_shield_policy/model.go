// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_policy

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type PageShieldPolicyResultEnvelope struct {
Result PageShieldPolicyModel `json:"result"`
}

type PageShieldPolicyModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
Action types.String `tfsdk:"action" json:"action,required"`
Description types.String `tfsdk:"description" json:"description,required"`
Enabled types.Bool `tfsdk:"enabled" json:"enabled,required"`
Expression types.String `tfsdk:"expression" json:"expression,required"`
Value types.String `tfsdk:"value" json:"value,required"`
}

func (m PageShieldPolicyModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m PageShieldPolicyModel) MarshalJSONForUpdate(state PageShieldPolicyModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}
