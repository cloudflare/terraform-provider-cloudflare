// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_key

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessKeyResultEnvelope struct {
	Result AccessKeyModel `json:"result,computed"`
}

type AccessKeyModel struct {
	Identifier              types.String  `tfsdk:"identifier" path:"identifier"`
	KeyRotationIntervalDays types.Float64 `tfsdk:"key_rotation_interval_days" json:"key_rotation_interval_days"`
}
