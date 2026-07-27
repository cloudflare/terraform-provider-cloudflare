// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_rule

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/email_routing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingRulesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[EmailRoutingRulesResultDataSourceModel] `json:"result,computed"`
}

type EmailRoutingRulesDataSourceModel struct {
	AccountID types.String                                                         `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String                                                         `tfsdk:"zone_id" path:"zone_id,optional"`
	Enabled   types.Bool                                                           `tfsdk:"enabled" query:"enabled,optional"`
	MaxItems  types.Int64                                                          `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[EmailRoutingRulesResultDataSourceModel] `tfsdk:"result"`
}

func (m *EmailRoutingRulesDataSourceModel) toListParams(_ context.Context) (params email_routing.RuleListParams, diags diag.Diagnostics) {
	params = email_routing.RuleListParams{}

	if !m.Enabled.IsNull() {
		params.Enabled = cloudflare.F(email_routing.RuleListParamsEnabled(m.Enabled.ValueBool()))
	}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type EmailRoutingRulesResultDataSourceModel struct {
	ID       types.String                                                           `tfsdk:"id" json:"id,computed"`
	Actions  customfield.NestedObjectList[EmailRoutingRulesActionsDataSourceModel]  `tfsdk:"actions" json:"actions,computed"`
	Enabled  types.Bool                                                             `tfsdk:"enabled" json:"enabled,computed"`
	Matchers customfield.NestedObjectList[EmailRoutingRulesMatchersDataSourceModel] `tfsdk:"matchers" json:"matchers,computed"`
	Name     types.String                                                           `tfsdk:"name" json:"name,computed"`
	Priority types.Float64                                                          `tfsdk:"priority" json:"priority,computed"`
	Source   types.String                                                           `tfsdk:"source" json:"source,computed"`
	Tag      types.String                                                           `tfsdk:"tag" json:"tag,computed"`
	Zone     customfield.NestedObject[EmailRoutingRulesZoneDataSourceModel]         `tfsdk:"zone" json:"zone,computed"`
}

type EmailRoutingRulesActionsDataSourceModel struct {
	Type  types.String                   `tfsdk:"type" json:"type,computed"`
	Value customfield.List[types.String] `tfsdk:"value" json:"value,computed"`
}

type EmailRoutingRulesMatchersDataSourceModel struct {
	Type  types.String `tfsdk:"type" json:"type,computed"`
	Field types.String `tfsdk:"field" json:"field,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type EmailRoutingRulesZoneDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
	Tag  types.String `tfsdk:"tag" json:"tag,computed"`
}
