// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FirewallRulesResultListDataSourceEnvelope struct {
	Result *[]*FirewallRulesResultDataSourceModel `json:"result,computed"`
}

type FirewallRulesDataSourceModel struct {
	ZoneIdentifier types.String                           `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String                           `tfsdk:"id" query:"id"`
	Action         types.String                           `tfsdk:"action" query:"action"`
	Description    types.String                           `tfsdk:"description" query:"description"`
	Page           types.Float64                          `tfsdk:"page" query:"page"`
	Paused         types.Bool                             `tfsdk:"paused" query:"paused"`
	PerPage        types.Float64                          `tfsdk:"per_page" query:"per_page"`
	MaxItems       types.Int64                            `tfsdk:"max_items"`
	Result         *[]*FirewallRulesResultDataSourceModel `tfsdk:"result"`
}

type FirewallRulesResultDataSourceModel struct {
	ID          types.String                                                 `tfsdk:"id" json:"id,computed"`
	Action      types.String                                                 `tfsdk:"action" json:"action,computed"`
	Filter      customfield.NestedObject[FirewallRulesFilterDataSourceModel] `tfsdk:"filter" json:"filter,computed"`
	Paused      types.Bool                                                   `tfsdk:"paused" json:"paused,computed"`
	Description types.String                                                 `tfsdk:"description" json:"description"`
	Priority    types.Float64                                                `tfsdk:"priority" json:"priority"`
	Products    *[]types.String                                              `tfsdk:"products" json:"products"`
	Ref         types.String                                                 `tfsdk:"ref" json:"ref"`
}

type FirewallRulesFilterDataSourceModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Description types.String `tfsdk:"description" json:"description"`
	Expression  types.String `tfsdk:"expression" json:"expression"`
	Paused      types.Bool   `tfsdk:"paused" json:"paused"`
	Ref         types.String `tfsdk:"ref" json:"ref"`
	Deleted     types.Bool   `tfsdk:"deleted" json:"deleted"`
}
