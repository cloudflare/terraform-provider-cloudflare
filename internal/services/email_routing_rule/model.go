// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_rule

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingRuleResultEnvelope struct {
	Result EmailRoutingRuleModel `json:"result"`
}

type EmailRoutingRuleModel struct {
	ID       types.String                                                `tfsdk:"id" json:"id,computed"`
	ZoneID   types.String                                                `tfsdk:"zone_id" path:"zone_id,required"`
	Enabled  types.Bool                                                  `tfsdk:"enabled" json:"enabled,computed_optional"`
	Name     types.String                                                `tfsdk:"name" json:"name,computed_optional"`
	Priority types.Float64                                               `tfsdk:"priority" json:"priority,computed_optional"`
	Actions  customfield.NestedObjectList[EmailRoutingRuleActionsModel]  `tfsdk:"actions" json:"actions,computed_optional"`
	Matchers customfield.NestedObjectList[EmailRoutingRuleMatchersModel] `tfsdk:"matchers" json:"matchers,computed_optional"`
	Tag      types.String                                                `tfsdk:"tag" json:"tag,computed"`
}

type EmailRoutingRuleActionsModel struct {
	Type  types.String                   `tfsdk:"type" json:"type,computed_optional"`
	Value customfield.List[types.String] `tfsdk:"value" json:"value,computed_optional"`
}

type EmailRoutingRuleMatchersModel struct {
	Field types.String `tfsdk:"field" json:"field,computed_optional"`
	Type  types.String `tfsdk:"type" json:"type,computed_optional"`
	Value types.String `tfsdk:"value" json:"value,computed_optional"`
}
