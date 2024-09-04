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
	ZoneIdentifier types.String                   `tfsdk:"zone_identifier" path:"zone_identifier"`
	PathID         types.String                   `tfsdk:"path_id" path:"id"`
	ID             types.String                   `tfsdk:"id" path:"id,computed_optional"`
	Action         *FirewallRuleActionModel       `tfsdk:"action" json:"action"`
	Filter         *FirewallRuleFilterModel       `tfsdk:"filter" json:"filter"`
	Description    types.String                   `tfsdk:"description" json:"description,computed"`
	Paused         types.Bool                     `tfsdk:"paused" json:"paused,computed"`
	Priority       types.Float64                  `tfsdk:"priority" json:"priority,computed"`
	Ref            types.String                   `tfsdk:"ref" json:"ref,computed"`
	Products       customfield.List[types.String] `tfsdk:"products" json:"products,computed"`
}

type FirewallRuleActionModel struct {
	Mode     types.String                                              `tfsdk:"mode" json:"mode,computed_optional"`
	Response customfield.NestedObject[FirewallRuleActionResponseModel] `tfsdk:"response" json:"response,computed_optional"`
	Timeout  types.Float64                                             `tfsdk:"timeout" json:"timeout,computed_optional"`
}

type FirewallRuleActionResponseModel struct {
	Body        types.String `tfsdk:"body" json:"body,computed_optional"`
	ContentType types.String `tfsdk:"content_type" json:"content_type,computed_optional"`
}

type FirewallRuleFilterModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Description types.String `tfsdk:"description" json:"description,computed_optional"`
	Expression  types.String `tfsdk:"expression" json:"expression,computed_optional"`
	Paused      types.Bool   `tfsdk:"paused" json:"paused,computed_optional"`
	Ref         types.String `tfsdk:"ref" json:"ref,computed_optional"`
}
