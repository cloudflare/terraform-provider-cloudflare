// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"context"

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
	PathID         types.String                   `tfsdk:"path_id" path:"id,optional"`
	ZoneIdentifier types.String                   `tfsdk:"zone_identifier" path:"zone_identifier,optional"`
	QueryID        types.String                   `tfsdk:"query_id" query:"id,optional"`
	Action         types.String                   `tfsdk:"action" json:"action,computed"`
	Description    types.String                   `tfsdk:"description" json:"description,computed"`
	ID             types.String                   `tfsdk:"id" json:"id,computed"`
	Paused         types.Bool                     `tfsdk:"paused" json:"paused,computed"`
	Priority       types.Float64                  `tfsdk:"priority" json:"priority,computed"`
	Ref            types.String                   `tfsdk:"ref" json:"ref,computed"`
	Products       customfield.List[types.String] `tfsdk:"products" json:"products,computed"`
}

func (m *FirewallRuleDataSourceModel) toReadParams(_ context.Context) (params firewall.RuleGetParams, diags diag.Diagnostics) {
	params = firewall.RuleGetParams{
		PathID: cloudflare.F(m.PathID.ValueString()),
	}

	return
}

func (m *FirewallRuleDataSourceModel) toListParams(_ context.Context) (params firewall.RuleListParams, diags diag.Diagnostics) {
	params = firewall.RuleListParams{}

	return
}
