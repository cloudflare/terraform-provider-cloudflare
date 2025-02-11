// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_rule

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingRuleResultEnvelope struct {
	Result EmailRoutingRuleModel `json:"result"`
}

type EmailRoutingRuleModel struct {
	ZoneID         types.String                      `tfsdk:"zone_id" path:"zone_id,required"`
	RuleIdentifier types.String                      `tfsdk:"rule_identifier" path:"rule_identifier,optional"`
	Actions        *[]*EmailRoutingRuleActionsModel  `tfsdk:"actions" json:"actions,required"`
	Matchers       *[]*EmailRoutingRuleMatchersModel `tfsdk:"matchers" json:"matchers,required"`
	Name           types.String                      `tfsdk:"name" json:"name,optional"`
	Enabled        types.Bool                        `tfsdk:"enabled" json:"enabled,computed_optional"`
	Priority       types.Float64                     `tfsdk:"priority" json:"priority,computed_optional"`
}

func (m EmailRoutingRuleModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m EmailRoutingRuleModel) MarshalJSONForUpdate(state EmailRoutingRuleModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type EmailRoutingRuleActionsModel struct {
	Type  types.String    `tfsdk:"type" json:"type,required"`
	Value *[]types.String `tfsdk:"value" json:"value,required"`
}

type EmailRoutingRuleMatchersModel struct {
	Field types.String `tfsdk:"field" json:"field,required"`
	Type  types.String `tfsdk:"type" json:"type,required"`
	Value types.String `tfsdk:"value" json:"value,required"`
}
