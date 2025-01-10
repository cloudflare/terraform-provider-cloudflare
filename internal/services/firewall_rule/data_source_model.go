// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/firewall"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FirewallRuleResultDataSourceEnvelope struct {
	Result FirewallRuleDataSourceModel `json:"result,computed"`
}

type FirewallRuleResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[FirewallRuleDataSourceModel] `json:"result,computed"`
}

type FirewallRuleDataSourceModel struct {
	RuleID      types.String                   `tfsdk:"rule_id" path:"rule_id,optional"`
	ZoneID      types.String                   `tfsdk:"zone_id" path:"zone_id,optional"`
	ID          types.String                   `tfsdk:"id" query:"id,computed_optional"`
	Action      types.String                   `tfsdk:"action" json:"action,computed"`
	Description types.String                   `tfsdk:"description" json:"description,computed"`
	Paused      types.Bool                     `tfsdk:"paused" json:"paused,computed"`
	Priority    types.Float64                  `tfsdk:"priority" json:"priority,computed"`
	Ref         types.String                   `tfsdk:"ref" json:"ref,computed"`
	Products    customfield.List[types.String] `tfsdk:"products" json:"products,computed"`
}

func (m *FirewallRuleDataSourceModel) toReadParams(_ context.Context) (params firewall.RuleGetParams, diags diag.Diagnostics) {
	params = firewall.RuleGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *FirewallRuleDataSourceModel) toListParams(_ context.Context) (params firewall.RuleListParams, diags diag.Diagnostics) {
	params = firewall.RuleListParams{}

	return
}
