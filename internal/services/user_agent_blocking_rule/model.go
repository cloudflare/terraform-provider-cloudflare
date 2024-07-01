// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_agent_blocking_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserAgentBlockingRuleResultEnvelope struct {
	Result UserAgentBlockingRuleModel `json:"result,computed"`
}

type UserAgentBlockingRuleResultDataSourceEnvelope struct {
	Result UserAgentBlockingRuleDataSourceModel `json:"result,computed"`
}

type UserAgentBlockingRulesResultDataSourceEnvelope struct {
	Result UserAgentBlockingRulesDataSourceModel `json:"result,computed"`
}

type UserAgentBlockingRuleModel struct {
	ZoneIdentifier types.String `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String `tfsdk:"id" path:"id"`
}

type UserAgentBlockingRuleDataSourceModel struct {
}

type UserAgentBlockingRulesDataSourceModel struct {
}
