// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_key_configuration

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessKeyConfigurationResultEnvelope struct {
	Result ZeroTrustAccessKeyConfigurationModel `json:"result"`
}

type ZeroTrustAccessKeyConfigurationModel struct {
	ID                      types.String      `tfsdk:"id" json:"-,computed"`
	AccountID               types.String      `tfsdk:"account_id" path:"account_id"`
	KeyRotationIntervalDays types.Float64     `tfsdk:"key_rotation_interval_days" json:"key_rotation_interval_days"`
	DaysUntilNextRotation   types.Float64     `tfsdk:"days_until_next_rotation" json:"days_until_next_rotation,computed"`
	LastKeyRotationAt       timetypes.RFC3339 `tfsdk:"last_key_rotation_at" json:"last_key_rotation_at,computed"`
}
