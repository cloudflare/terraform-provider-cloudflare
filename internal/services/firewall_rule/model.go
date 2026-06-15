// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FirewallRuleResultEnvelope struct {
	Result FirewallRuleModel `json:"result"`
}

type FirewallRuleModel struct {
	ID          types.String                   `tfsdk:"id" json:"id,computed"`
	ZoneID      types.String                   `tfsdk:"zone_id" path:"zone_id,required"`
	Action      *FirewallRuleActionModel       `tfsdk:"action" json:"action,required,no_refresh"`
	Filter      *FirewallRuleFilterModel       `tfsdk:"filter" json:"filter,required"`
	Description types.String                   `tfsdk:"description" json:"description,computed"`
	Paused      types.Bool                     `tfsdk:"paused" json:"paused,computed"`
	Priority    types.Float64                  `tfsdk:"priority" json:"priority,computed"`
	Ref         types.String                   `tfsdk:"ref" json:"ref,computed"`
	Products    customfield.List[types.String] `tfsdk:"products" json:"products,computed"`
}

func (m FirewallRuleModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m FirewallRuleModel) MarshalJSONForUpdate(state FirewallRuleModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
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
