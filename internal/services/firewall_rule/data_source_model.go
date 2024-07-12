// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FirewallRuleResultDataSourceEnvelope struct {
	Result FirewallRuleDataSourceModel `json:"result,computed"`
}

type FirewallRuleResultListDataSourceEnvelope struct {
	Result *[]*FirewallRuleDataSourceModel `json:"result,computed"`
}

type FirewallRuleDataSourceModel struct {
	PathID         types.String                          `tfsdk:"path_id" path:"id"`
	ZoneIdentifier types.String                          `tfsdk:"zone_identifier" path:"zone_identifier"`
	QueryID        types.String                          `tfsdk:"query_id" query:"id"`
	ID             types.String                          `tfsdk:"id" json:"id,computed"`
	Action         types.String                          `tfsdk:"action" json:"action,computed"`
	Paused         types.Bool                            `tfsdk:"paused" json:"paused,computed"`
	Description    types.String                          `tfsdk:"description" json:"description"`
	Priority       types.Float64                         `tfsdk:"priority" json:"priority"`
	Products       *[]types.String                       `tfsdk:"products" json:"products"`
	Ref            types.String                          `tfsdk:"ref" json:"ref"`
	FindOneBy      *FirewallRuleFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type FirewallRuleFindOneByDataSourceModel struct {
	ZoneIdentifier types.String  `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String  `tfsdk:"id" query:"id"`
	Action         types.String  `tfsdk:"action" query:"action"`
	Description    types.String  `tfsdk:"description" query:"description"`
	Page           types.Float64 `tfsdk:"page" query:"page"`
	Paused         types.Bool    `tfsdk:"paused" query:"paused"`
	PerPage        types.Float64 `tfsdk:"per_page" query:"per_page"`
}
