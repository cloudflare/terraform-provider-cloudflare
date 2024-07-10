// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FirewallRuleResultEnvelope struct {
	Result FirewallRuleModel `json:"result,computed"`
}

type FirewallRuleModel struct {
	ZoneIdentifier types.String    `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String    `tfsdk:"id" path:"id"`
	PathID         types.String    `tfsdk:"path_id" path:"id"`
	Action         types.String    `tfsdk:"action" json:"action,computed"`
	Paused         types.Bool      `tfsdk:"paused" json:"paused,computed"`
	Description    types.String    `tfsdk:"description" json:"description,computed"`
	Priority       types.Float64   `tfsdk:"priority" json:"priority,computed"`
	Products       *[]types.String `tfsdk:"products" json:"products,computed"`
	Ref            types.String    `tfsdk:"ref" json:"ref,computed"`
}
