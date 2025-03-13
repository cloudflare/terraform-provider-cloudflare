// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_logging

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayLoggingResultEnvelope struct {
	Result ZeroTrustGatewayLoggingModel `json:"result"`
}

type ZeroTrustGatewayLoggingModel struct {
	AccountID          types.String                                                             `tfsdk:"account_id" path:"account_id,required"`
	RedactPii          types.Bool                                                               `tfsdk:"redact_pii" json:"redact_pii,optional"`
	SettingsByRuleType customfield.NestedObject[ZeroTrustGatewayLoggingSettingsByRuleTypeModel] `tfsdk:"settings_by_rule_type" json:"settings_by_rule_type,computed_optional"`
}

func (m ZeroTrustGatewayLoggingModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustGatewayLoggingModel) MarshalJSONForUpdate(state ZeroTrustGatewayLoggingModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustGatewayLoggingSettingsByRuleTypeModel struct {
	DNS  jsontypes.Normalized `tfsdk:"dns" json:"dns,optional"`
	HTTP jsontypes.Normalized `tfsdk:"http" json:"http,optional"`
	L4   jsontypes.Normalized `tfsdk:"l4" json:"l4,optional"`
}
