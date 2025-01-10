// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_key_configuration

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessKeyConfigurationResultDataSourceEnvelope struct {
	Result ZeroTrustAccessKeyConfigurationDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessKeyConfigurationDataSourceModel struct {
	AccountID               types.String      `tfsdk:"account_id" path:"account_id,required"`
	DaysUntilNextRotation   types.Float64     `tfsdk:"days_until_next_rotation" json:"days_until_next_rotation,computed"`
	KeyRotationIntervalDays types.Float64     `tfsdk:"key_rotation_interval_days" json:"key_rotation_interval_days,computed"`
	LastKeyRotationAt       timetypes.RFC3339 `tfsdk:"last_key_rotation_at" json:"last_key_rotation_at,computed" format:"date-time"`
}

func (m *ZeroTrustAccessKeyConfigurationDataSourceModel) toReadParams(_ context.Context) (params zero_trust.AccessKeyGetParams, diags diag.Diagnostics) {
	params = zero_trust.AccessKeyGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
