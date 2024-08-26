// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/firewall"
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
	PathID         types.String    `tfsdk:"path_id" path:"id"`
	ZoneIdentifier types.String    `tfsdk:"zone_identifier" path:"zone_identifier"`
	QueryID        types.String    `tfsdk:"query_id" query:"id"`
	Action         types.String    `tfsdk:"action" json:"action,computed"`
	ID             types.String    `tfsdk:"id" json:"id,computed"`
	Paused         types.Bool      `tfsdk:"paused" json:"paused,computed"`
	Description    types.String    `tfsdk:"description" json:"description,computed_optional"`
	Priority       types.Float64   `tfsdk:"priority" json:"priority,computed_optional"`
	Ref            types.String    `tfsdk:"ref" json:"ref,computed_optional"`
	Products       *[]types.String `tfsdk:"products" json:"products,computed_optional"`
}

func (m *FirewallRuleDataSourceModel) toReadParams() (params firewall.RuleGetParams, diags diag.Diagnostics) {
	params = firewall.RuleGetParams{
		PathID: cloudflare.F(m.PathID.ValueString()),
	}

	return
}

func (m *FirewallRuleDataSourceModel) toListParams() (params firewall.RuleListParams, diags diag.Diagnostics) {
	params = firewall.RuleListParams{}

	return
}
