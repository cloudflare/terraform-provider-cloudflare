// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_key_configuration

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessKeyConfigurationResultEnvelope struct {
Result ZeroTrustAccessKeyConfigurationModel `json:"result"`
}

type ZeroTrustAccessKeyConfigurationModel struct {
ID types.String `tfsdk:"id" json:"-,computed"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
KeyRotationIntervalDays types.Float64 `tfsdk:"key_rotation_interval_days" json:"key_rotation_interval_days,required"`
DaysUntilNextRotation types.Float64 `tfsdk:"days_until_next_rotation" json:"days_until_next_rotation,computed"`
LastKeyRotationAt timetypes.RFC3339 `tfsdk:"last_key_rotation_at" json:"last_key_rotation_at,computed" format:"date-time"`
}

func (m ZeroTrustAccessKeyConfigurationModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m ZeroTrustAccessKeyConfigurationModel) MarshalJSONForUpdate(state ZeroTrustAccessKeyConfigurationModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}
