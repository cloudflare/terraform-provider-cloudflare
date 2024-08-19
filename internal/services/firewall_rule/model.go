// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FirewallRuleResultEnvelope struct {
	Result FirewallRuleModel `json:"result,computed"`
}

type FirewallRuleModel struct {
	ZoneIdentifier types.String             `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String             `tfsdk:"id" path:"id"`
	PathID         types.String             `tfsdk:"path_id" path:"id"`
	Action         *FirewallRuleActionModel `tfsdk:"action" json:"action"`
	Filter         *FirewallRuleFilterModel `tfsdk:"filter" json:"filter"`
	Description    types.String             `tfsdk:"description" json:"description,computed"`
	Paused         types.Bool               `tfsdk:"paused" json:"paused,computed"`
	Priority       types.Float64            `tfsdk:"priority" json:"priority,computed"`
	Ref            types.String             `tfsdk:"ref" json:"ref,computed"`
	Products       *[]types.String          `tfsdk:"products" json:"products,computed"`
}

type FirewallRuleActionModel struct {
	Mode     types.String                     `tfsdk:"mode" json:"mode"`
	Response *FirewallRuleActionResponseModel `tfsdk:"response" json:"response"`
	Timeout  types.Float64                    `tfsdk:"timeout" json:"timeout"`
}

type FirewallRuleActionResponseModel struct {
	Body        types.String `tfsdk:"body" json:"body"`
	ContentType types.String `tfsdk:"content_type" json:"content_type"`
}

type FirewallRuleFilterModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Description types.String `tfsdk:"description" json:"description"`
	Expression  types.String `tfsdk:"expression" json:"expression"`
	Paused      types.Bool   `tfsdk:"paused" json:"paused"`
	Ref         types.String `tfsdk:"ref" json:"ref"`
}
