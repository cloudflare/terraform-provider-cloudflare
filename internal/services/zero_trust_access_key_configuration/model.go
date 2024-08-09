// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_key_configuration

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessKeyConfigurationResultEnvelope struct {
	Result ZeroTrustAccessKeyConfigurationModel `json:"result,computed"`
}

type ZeroTrustAccessKeyConfigurationModel struct {
	ID                      types.String  `tfsdk:"id" json:"-,computed"`
	AccountID               types.String  `tfsdk:"account_id" path:"account_id"`
	KeyRotationIntervalDays types.Float64 `tfsdk:"key_rotation_interval_days" json:"key_rotation_interval_days"`
}
