// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_rule

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/email_routing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingRuleResultDataSourceEnvelope struct {
	Result EmailRoutingRuleDataSourceModel `json:"result,computed"`
}

type EmailRoutingRuleResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[EmailRoutingRuleDataSourceModel] `json:"result,computed"`
}

type EmailRoutingRuleDataSourceModel struct {
	ID             types.String                                                          `tfsdk:"id" json:"-,computed"`
	RuleIdentifier types.String                                                          `tfsdk:"rule_identifier" path:"rule_identifier,optional"`
	ZoneID         types.String                                                          `tfsdk:"zone_id" path:"zone_id,required"`
	Enabled        types.Bool                                                            `tfsdk:"enabled" json:"enabled,computed"`
	Name           types.String                                                          `tfsdk:"name" json:"name,computed"`
	Priority       types.Float64                                                         `tfsdk:"priority" json:"priority,computed"`
	Tag            types.String                                                          `tfsdk:"tag" json:"tag,computed"`
	Actions        customfield.NestedObjectList[EmailRoutingRuleActionsDataSourceModel]  `tfsdk:"actions" json:"actions,computed"`
	Matchers       customfield.NestedObjectList[EmailRoutingRuleMatchersDataSourceModel] `tfsdk:"matchers" json:"matchers,computed"`
	Filter         *EmailRoutingRuleFindOneByDataSourceModel                             `tfsdk:"filter"`
}

func (m *EmailRoutingRuleDataSourceModel) toReadParams(_ context.Context) (params email_routing.RuleGetParams, diags diag.Diagnostics) {
	params = email_routing.RuleGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *EmailRoutingRuleDataSourceModel) toListParams(_ context.Context) (params email_routing.RuleListParams, diags diag.Diagnostics) {
	params = email_routing.RuleListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if !m.Filter.Enabled.IsNull() {
		params.Enabled = cloudflare.F(email_routing.RuleListParamsEnabled(m.Filter.Enabled.ValueBool()))
	}

	return
}

type EmailRoutingRuleActionsDataSourceModel struct {
	Type  types.String                   `tfsdk:"type" json:"type,computed"`
	Value customfield.List[types.String] `tfsdk:"value" json:"value,computed"`
}

type EmailRoutingRuleMatchersDataSourceModel struct {
	Field types.String `tfsdk:"field" json:"field,computed"`
	Type  types.String `tfsdk:"type" json:"type,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type EmailRoutingRuleFindOneByDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" query:"enabled,optional"`
}
