// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_agent_blocking_rule

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserAgentBlockingRuleResultEnvelope struct {
	Result UserAgentBlockingRuleModel `json:"result"`
}

type UserAgentBlockingRuleModel struct {
	ZoneID        types.String                             `tfsdk:"zone_id" path:"zone_id,required"`
	UARuleID      types.String                             `tfsdk:"ua_rule_id" path:"ua_rule_id,optional"`
	Mode          types.String                             `tfsdk:"mode" json:"mode,required,no_refresh"`
	Configuration *UserAgentBlockingRuleConfigurationModel `tfsdk:"configuration" json:"configuration,required,no_refresh"`
}

func (m UserAgentBlockingRuleModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m UserAgentBlockingRuleModel) MarshalJSONForUpdate(state UserAgentBlockingRuleModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type UserAgentBlockingRuleConfigurationModel struct {
	Target types.String `tfsdk:"target" json:"target,optional"`
	Value  types.String `tfsdk:"value" json:"value,optional"`
}
