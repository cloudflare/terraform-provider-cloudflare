// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_agent_blocking_rule

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/firewall"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserAgentBlockingRulesResultListDataSourceEnvelope struct {
	Result *[]*UserAgentBlockingRulesResultDataSourceModel `json:"result,computed"`
}

type UserAgentBlockingRulesDataSourceModel struct {
	ZoneIdentifier    types.String                                    `tfsdk:"zone_identifier" path:"zone_identifier"`
	Description       types.String                                    `tfsdk:"description" query:"description"`
	DescriptionSearch types.String                                    `tfsdk:"description_search" query:"description_search"`
	UASearch          types.String                                    `tfsdk:"ua_search" query:"ua_search"`
	MaxItems          types.Int64                                     `tfsdk:"max_items"`
	Result            *[]*UserAgentBlockingRulesResultDataSourceModel `tfsdk:"result"`
}

func (m *UserAgentBlockingRulesDataSourceModel) toListParams() (params firewall.UARuleListParams, diags diag.Diagnostics) {
	params = firewall.UARuleListParams{}

	if !m.Description.IsNull() {
		params.Description = cloudflare.F(m.Description.ValueString())
	}
	if !m.DescriptionSearch.IsNull() {
		params.DescriptionSearch = cloudflare.F(m.DescriptionSearch.ValueString())
	}
	if !m.UASearch.IsNull() {
		params.UASearch = cloudflare.F(m.UASearch.ValueString())
	}

	return
}

type UserAgentBlockingRulesResultDataSourceModel struct {
	ID            types.String                                        `tfsdk:"id" json:"id,computed"`
	Configuration *UserAgentBlockingRulesConfigurationDataSourceModel `tfsdk:"configuration" json:"configuration"`
	Description   types.String                                        `tfsdk:"description" json:"description"`
	Mode          types.String                                        `tfsdk:"mode" json:"mode"`
	Paused        types.Bool                                          `tfsdk:"paused" json:"paused"`
}

type UserAgentBlockingRulesConfigurationDataSourceModel struct {
	Target types.String `tfsdk:"target" json:"target"`
	Value  types.String `tfsdk:"value" json:"value"`
}
