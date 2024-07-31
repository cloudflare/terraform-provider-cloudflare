// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_agent_blocking_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserAgentBlockingRuleResultEnvelope struct {
	Result UserAgentBlockingRuleModel `json:"result,computed"`
}

type UserAgentBlockingRuleModel struct {
	ZoneIdentifier types.String                             `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String                             `tfsdk:"id" path:"id"`
	Configuration  *UserAgentBlockingRuleConfigurationModel `tfsdk:"configuration" json:"configuration"`
	Mode           types.String                             `tfsdk:"mode" json:"mode"`
}

type UserAgentBlockingRuleConfigurationModel struct {
	Target types.String `tfsdk:"target" json:"target"`
	Value  types.String `tfsdk:"value" json:"value"`
}
