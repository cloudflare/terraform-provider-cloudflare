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
	PathID         types.String    `tfsdk:"path_id" path:"id"`
	ZoneIdentifier types.String    `tfsdk:"zone_identifier" path:"zone_identifier"`
	QueryID        types.String    `tfsdk:"query_id" query:"id"`
	ID             types.String    `tfsdk:"id" json:"id,computed"`
	Action         types.String    `tfsdk:"action" json:"action,computed"`
	Paused         types.Bool      `tfsdk:"paused" json:"paused,computed"`
	Description    types.String    `tfsdk:"description" json:"description"`
	Priority       types.Float64   `tfsdk:"priority" json:"priority"`
	Products       *[]types.String `tfsdk:"products" json:"products"`
	Ref            types.String    `tfsdk:"ref" json:"ref"`
}
