// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_keys_configuration

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessKeysConfigurationResultEnvelope struct {
	Result AccessKeysConfigurationModel `json:"result,computed"`
}

type AccessKeysConfigurationModel struct {
	AccountID               types.String  `tfsdk:"account_id" path:"account_id"`
	KeyRotationIntervalDays types.Float64 `tfsdk:"key_rotation_interval_days" json:"key_rotation_interval_days"`
}
