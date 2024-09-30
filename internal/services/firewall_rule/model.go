// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FirewallRuleResultEnvelope struct {
	Result FirewallRuleModel `json:"result"`
}

type FirewallRuleModel struct {
	ZoneIdentifier types.String                   `tfsdk:"zone_identifier" path:"zone_identifier,required"`
	ID             types.String                   `tfsdk:"id" path:"id,optional"`
	PathID         types.String                   `tfsdk:"path_id" path:"id,optional"`
	Action         *FirewallRuleActionModel       `tfsdk:"action" json:"action,required"`
	Filter         *FirewallRuleFilterModel       `tfsdk:"filter" json:"filter,required"`
	Description    types.String                   `tfsdk:"description" json:"description,computed"`
	Paused         types.Bool                     `tfsdk:"paused" json:"paused,computed"`
	Priority       types.Float64                  `tfsdk:"priority" json:"priority,computed"`
	Ref            types.String                   `tfsdk:"ref" json:"ref,computed"`
	Products       customfield.List[types.String] `tfsdk:"products" json:"products,computed"`
}

type FirewallRuleActionModel struct {
	Mode     types.String                     `tfsdk:"mode" json:"mode,optional"`
	Response *FirewallRuleActionResponseModel `tfsdk:"response" json:"response,optional"`
	Timeout  types.Float64                    `tfsdk:"timeout" json:"timeout,optional"`
}

type FirewallRuleActionResponseModel struct {
	Body        types.String `tfsdk:"body" json:"body,optional"`
	ContentType types.String `tfsdk:"content_type" json:"content_type,optional"`
}

type FirewallRuleFilterModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Description types.String `tfsdk:"description" json:"description,optional"`
	Expression  types.String `tfsdk:"expression" json:"expression,optional"`
	Paused      types.Bool   `tfsdk:"paused" json:"paused,optional"`
	Ref         types.String `tfsdk:"ref" json:"ref,optional"`
}
