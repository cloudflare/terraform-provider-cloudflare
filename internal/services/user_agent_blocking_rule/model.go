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
	ZoneIdentifier types.String                             `tfsdk:"zone_identifier" path:"zone_identifier,required"`
	ID             types.String                             `tfsdk:"id" path:"id,optional"`
	Mode           types.String                             `tfsdk:"mode" json:"mode,required"`
	Configuration  *UserAgentBlockingRuleConfigurationModel `tfsdk:"configuration" json:"configuration,required"`
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
