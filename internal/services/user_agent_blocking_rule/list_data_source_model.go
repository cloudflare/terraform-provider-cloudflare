// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_agent_blocking_rule

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/firewall"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserAgentBlockingRulesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[UserAgentBlockingRulesResultDataSourceModel] `json:"result,computed"`
}

type UserAgentBlockingRulesDataSourceModel struct {
	ZoneID            types.String                                                              `tfsdk:"zone_id" path:"zone_id,required"`
	Description       types.String                                                              `tfsdk:"description" query:"description,optional"`
	DescriptionSearch types.String                                                              `tfsdk:"description_search" query:"description_search,optional"`
	UASearch          types.String                                                              `tfsdk:"ua_search" query:"ua_search,optional"`
	MaxItems          types.Int64                                                               `tfsdk:"max_items"`
	Result            customfield.NestedObjectList[UserAgentBlockingRulesResultDataSourceModel] `tfsdk:"result"`
}

func (m *UserAgentBlockingRulesDataSourceModel) toListParams(_ context.Context) (params firewall.UARuleListParams, diags diag.Diagnostics) {
	params = firewall.UARuleListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

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
	ID            types.String                                                                 `tfsdk:"id" json:"id,computed"`
	Configuration customfield.NestedObject[UserAgentBlockingRulesConfigurationDataSourceModel] `tfsdk:"configuration" json:"configuration,computed"`
	Description   types.String                                                                 `tfsdk:"description" json:"description,computed"`
	Mode          types.String                                                                 `tfsdk:"mode" json:"mode,computed"`
	Paused        types.Bool                                                                   `tfsdk:"paused" json:"paused,computed"`
}

type UserAgentBlockingRulesConfigurationDataSourceModel struct {
	Target types.String `tfsdk:"target" json:"target,computed"`
	Value  types.String `tfsdk:"value" json:"value,computed"`
}
