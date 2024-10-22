// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/firewall"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FirewallRulesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[FirewallRulesResultDataSourceModel] `json:"result,computed"`
}

type FirewallRulesDataSourceModel struct {
	ZoneIdentifier types.String                                                     `tfsdk:"zone_identifier" path:"zone_identifier,required"`
	Action         types.String                                                     `tfsdk:"action" query:"action,optional"`
	Description    types.String                                                     `tfsdk:"description" query:"description,optional"`
	ID             types.String                                                     `tfsdk:"id" query:"id,optional"`
	Paused         types.Bool                                                       `tfsdk:"paused" query:"paused,optional"`
	MaxItems       types.Int64                                                      `tfsdk:"max_items"`
	Result         customfield.NestedObjectList[FirewallRulesResultDataSourceModel] `tfsdk:"result"`
}

func (m *FirewallRulesDataSourceModel) toListParams(_ context.Context) (params firewall.RuleListParams, diags diag.Diagnostics) {
	params = firewall.RuleListParams{}

	if !m.ID.IsNull() {
		params.ID = cloudflare.F(m.ID.ValueString())
	}
	if !m.Action.IsNull() {
		params.Action = cloudflare.F(m.Action.ValueString())
	}
	if !m.Description.IsNull() {
		params.Description = cloudflare.F(m.Description.ValueString())
	}
	if !m.Paused.IsNull() {
		params.Paused = cloudflare.F(m.Paused.ValueBool())
	}

	return
}

type FirewallRulesResultDataSourceModel struct {
	ID          types.String                                                 `tfsdk:"id" json:"id,computed"`
	Action      types.String                                                 `tfsdk:"action" json:"action,computed"`
	Filter      customfield.NestedObject[FirewallRulesFilterDataSourceModel] `tfsdk:"filter" json:"filter,computed"`
	Paused      types.Bool                                                   `tfsdk:"paused" json:"paused,computed"`
	Description types.String                                                 `tfsdk:"description" json:"description,computed"`
	Priority    types.Float64                                                `tfsdk:"priority" json:"priority,computed"`
	Products    customfield.List[types.String]                               `tfsdk:"products" json:"products,computed"`
	Ref         types.String                                                 `tfsdk:"ref" json:"ref,computed"`
}

type FirewallRulesFilterDataSourceModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Expression  types.String `tfsdk:"expression" json:"expression,computed"`
	Paused      types.Bool   `tfsdk:"paused" json:"paused,computed"`
	Ref         types.String `tfsdk:"ref" json:"ref,computed"`
	Deleted     types.Bool   `tfsdk:"deleted" json:"deleted,computed"`
}
