// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_logging

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
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
	DNS  customfield.NestedObject[ZeroTrustGatewayLoggingSettingsByRuleTypeDNSModel]  `tfsdk:"dns" json:"dns,computed_optional"`
	HTTP customfield.NestedObject[ZeroTrustGatewayLoggingSettingsByRuleTypeHTTPModel] `tfsdk:"http" json:"http,computed_optional"`
	L4   customfield.NestedObject[ZeroTrustGatewayLoggingSettingsByRuleTypeL4Model]   `tfsdk:"l4" json:"l4,computed_optional"`
}

type ZeroTrustGatewayLoggingSettingsByRuleTypeDNSModel struct {
	LogAll    types.Bool `tfsdk:"log_all" json:"log_all,optional"`
	LogBlocks types.Bool `tfsdk:"log_blocks" json:"log_blocks,optional"`
}

type ZeroTrustGatewayLoggingSettingsByRuleTypeHTTPModel struct {
	LogAll    types.Bool `tfsdk:"log_all" json:"log_all,optional"`
	LogBlocks types.Bool `tfsdk:"log_blocks" json:"log_blocks,optional"`
}

type ZeroTrustGatewayLoggingSettingsByRuleTypeL4Model struct {
	LogAll    types.Bool `tfsdk:"log_all" json:"log_all,optional"`
	LogBlocks types.Bool `tfsdk:"log_blocks" json:"log_blocks,optional"`
}
