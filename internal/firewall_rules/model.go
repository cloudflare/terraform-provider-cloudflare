// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rules

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FirewallRulesResultEnvelope struct {
	Result FirewallRulesModel `json:"result,computed"`
}

type FirewallRulesModel struct {
	ZoneIdentifier types.String              `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String              `tfsdk:"id" path:"id"`
	PathID         types.String              `tfsdk:"path_id" path:"path_id"`
	Action         types.String              `tfsdk:"action" json:"action"`
	Filter         *FirewallRulesFilterModel `tfsdk:"filter" json:"filter"`
	Paused         types.Bool                `tfsdk:"paused" json:"paused"`
	Description    types.String              `tfsdk:"description" json:"description"`
	Priority       types.Float64             `tfsdk:"priority" json:"priority"`
	Products       []types.String            `tfsdk:"products" json:"products"`
	Ref            types.String              `tfsdk:"ref" json:"ref"`
}

type FirewallRulesFilterModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Description types.String `tfsdk:"description" json:"description"`
	Expression  types.String `tfsdk:"expression" json:"expression"`
	Paused      types.Bool   `tfsdk:"paused" json:"paused"`
	Ref         types.String `tfsdk:"ref" json:"ref"`
	Deleted     types.Bool   `tfsdk:"deleted" json:"deleted"`
}
