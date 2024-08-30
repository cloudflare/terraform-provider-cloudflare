// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_agent_blocking_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserAgentBlockingRuleResultEnvelope struct {
	Result UserAgentBlockingRuleModel `json:"result"`
}

type UserAgentBlockingRuleModel struct {
	ZoneIdentifier types.String                             `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String                             `tfsdk:"id" path:"id"`
	Mode           types.String                             `tfsdk:"mode" json:"mode"`
	Configuration  *UserAgentBlockingRuleConfigurationModel `tfsdk:"configuration" json:"configuration"`
}

type UserAgentBlockingRuleConfigurationModel struct {
	Target types.String `tfsdk:"target" json:"target,computed_optional"`
	Value  types.String `tfsdk:"value" json:"value,computed_optional"`
}
