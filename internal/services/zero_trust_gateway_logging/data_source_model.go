// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_logging

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayLoggingResultDataSourceEnvelope struct {
	Result ZeroTrustGatewayLoggingDataSourceModel `json:"result,computed"`
}

type ZeroTrustGatewayLoggingDataSourceModel struct {
	AccountID          types.String         `tfsdk:"account_id" path:"account_id,required"`
	RedactPii          types.Bool           `tfsdk:"redact_pii" json:"redact_pii,computed"`
	SettingsByRuleType jsontypes.Normalized `tfsdk:"settings_by_rule_type" json:"settings_by_rule_type,computed"`
}

func (m *ZeroTrustGatewayLoggingDataSourceModel) toReadParams(_ context.Context) (params zero_trust.GatewayLoggingGetParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayLoggingGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
