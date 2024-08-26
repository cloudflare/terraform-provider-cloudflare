// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_rule

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/email_routing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingRulesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[EmailRoutingRulesResultDataSourceModel] `json:"result,computed"`
}

type EmailRoutingRulesDataSourceModel struct {
	ZoneIdentifier types.String                                                         `tfsdk:"zone_identifier" path:"zone_identifier"`
	Enabled        types.Bool                                                           `tfsdk:"enabled" query:"enabled"`
	MaxItems       types.Int64                                                          `tfsdk:"max_items"`
	Result         customfield.NestedObjectList[EmailRoutingRulesResultDataSourceModel] `tfsdk:"result"`
}

func (m *EmailRoutingRulesDataSourceModel) toListParams() (params email_routing.RuleListParams, diags diag.Diagnostics) {
	params = email_routing.RuleListParams{}

	if !m.Enabled.IsNull() {
		params.Enabled = cloudflare.F(email_routing.RuleListParamsEnabled(m.Enabled.ValueBool()))
	}

	return
}

type EmailRoutingRulesResultDataSourceModel struct {
	ID       types.String                                 `tfsdk:"id" json:"id,computed"`
	Actions  *[]*EmailRoutingRulesActionsDataSourceModel  `tfsdk:"actions" json:"actions"`
	Enabled  types.Bool                                   `tfsdk:"enabled" json:"enabled,computed"`
	Matchers *[]*EmailRoutingRulesMatchersDataSourceModel `tfsdk:"matchers" json:"matchers"`
	Name     types.String                                 `tfsdk:"name" json:"name"`
	Priority types.Float64                                `tfsdk:"priority" json:"priority,computed"`
	Tag      types.String                                 `tfsdk:"tag" json:"tag,computed"`
}

type EmailRoutingRulesActionsDataSourceModel struct {
	Type  types.String `tfsdk:"type" json:"type,computed"`
	Value types.List   `tfsdk:"value" json:"value,computed"`
}

type EmailRoutingRulesMatchersDataSourceModel struct {
	Field types.String `tfsdk:"field" json:"field,computed"`
	Type  types.String `tfsdk:"type" json:"type,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}
