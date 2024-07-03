// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FirewallRulesResultListDataSourceEnvelope struct {
	Result *[]*FirewallRulesItemsDataSourceModel `json:"result,computed"`
}

type FirewallRulesDataSourceModel struct {
	ZoneIdentifier types.String                          `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String                          `tfsdk:"id" query:"id"`
	Action         types.String                          `tfsdk:"action" query:"action"`
	Description    types.String                          `tfsdk:"description" query:"description"`
	Page           types.Float64                         `tfsdk:"page" query:"page"`
	Paused         types.Bool                            `tfsdk:"paused" query:"paused"`
	PerPage        types.Float64                         `tfsdk:"per_page" query:"per_page"`
	MaxItems       types.Int64                           `tfsdk:"max_items"`
	Items          *[]*FirewallRulesItemsDataSourceModel `tfsdk:"items"`
}

type FirewallRulesItemsDataSourceModel struct {
	ID          types.String    `tfsdk:"id" json:"id,computed"`
	Action      types.String    `tfsdk:"action" json:"action,computed"`
	Paused      types.Bool      `tfsdk:"paused" json:"paused,computed"`
	Description types.String    `tfsdk:"description" json:"description,computed"`
	Priority    types.Float64   `tfsdk:"priority" json:"priority,computed"`
	Products    *[]types.String `tfsdk:"products" json:"products,computed"`
	Ref         types.String    `tfsdk:"ref" json:"ref,computed"`
}

type FirewallRulesItemsFilterDataSourceModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Expression  types.String `tfsdk:"expression" json:"expression,computed"`
	Paused      types.Bool   `tfsdk:"paused" json:"paused,computed"`
	Ref         types.String `tfsdk:"ref" json:"ref,computed"`
	Deleted     types.Bool   `tfsdk:"deleted" json:"deleted,computed"`
}
