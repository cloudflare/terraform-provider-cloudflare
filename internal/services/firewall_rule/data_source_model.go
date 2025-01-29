// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/firewall"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FirewallRuleResultDataSourceEnvelope struct {
	Result FirewallRuleDataSourceModel `json:"result,computed"`
}

type FirewallRuleDataSourceModel struct {
	RuleID      types.String                                                `tfsdk:"rule_id" path:"rule_id,required"`
	ZoneID      types.String                                                `tfsdk:"zone_id" path:"zone_id,required"`
	ID          types.String                                                `tfsdk:"id" query:"id,computed_optional"`
	Action      types.String                                                `tfsdk:"action" json:"action,computed"`
	Description types.String                                                `tfsdk:"description" json:"description,computed"`
	Paused      types.Bool                                                  `tfsdk:"paused" json:"paused,computed"`
	Priority    types.Float64                                               `tfsdk:"priority" json:"priority,computed"`
	Ref         types.String                                                `tfsdk:"ref" json:"ref,computed"`
	Products    customfield.List[types.String]                              `tfsdk:"products" json:"products,computed"`
	Filter      customfield.NestedObject[FirewallRuleFilterDataSourceModel] `tfsdk:"filter" json:"filter,computed"`
}

func (m *FirewallRuleDataSourceModel) toReadParams(_ context.Context) (params firewall.RuleGetParams, diags diag.Diagnostics) {
	params = firewall.RuleGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type FirewallRuleFilterDataSourceModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Expression  types.String `tfsdk:"expression" json:"expression,computed"`
	Paused      types.Bool   `tfsdk:"paused" json:"paused,computed"`
	Ref         types.String `tfsdk:"ref" json:"ref,computed"`
	Deleted     types.Bool   `tfsdk:"deleted" json:"deleted,computed"`
}
