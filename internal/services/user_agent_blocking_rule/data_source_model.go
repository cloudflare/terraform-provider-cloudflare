// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_agent_blocking_rule

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/firewall"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserAgentBlockingRuleResultDataSourceEnvelope struct {
	Result UserAgentBlockingRuleDataSourceModel `json:"result,computed"`
}

type UserAgentBlockingRuleResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[UserAgentBlockingRuleDataSourceModel] `json:"result,computed"`
}

type UserAgentBlockingRuleDataSourceModel struct {
	ZoneIdentifier types.String                                       `tfsdk:"zone_identifier" path:"zone_identifier,optional"`
	ID             types.String                                       `tfsdk:"id" path:"id,computed_optional"`
	Description    types.String                                       `tfsdk:"description" json:"description,optional"`
	Mode           types.String                                       `tfsdk:"mode" json:"mode,optional"`
	Paused         types.Bool                                         `tfsdk:"paused" json:"paused,optional"`
	Configuration  *UserAgentBlockingRuleConfigurationDataSourceModel `tfsdk:"configuration" json:"configuration,optional"`
	Filter         *UserAgentBlockingRuleFindOneByDataSourceModel     `tfsdk:"filter"`
}

func (m *UserAgentBlockingRuleDataSourceModel) toListParams(_ context.Context) (params firewall.UARuleListParams, diags diag.Diagnostics) {
	params = firewall.UARuleListParams{}

	if !m.Filter.Description.IsNull() {
		params.Description = cloudflare.F(m.Filter.Description.ValueString())
	}
	if !m.Filter.DescriptionSearch.IsNull() {
		params.DescriptionSearch = cloudflare.F(m.Filter.DescriptionSearch.ValueString())
	}
	if !m.Filter.UASearch.IsNull() {
		params.UASearch = cloudflare.F(m.Filter.UASearch.ValueString())
	}

	return
}

type UserAgentBlockingRuleConfigurationDataSourceModel struct {
	Target types.String `tfsdk:"target" json:"target,computed"`
	Value  types.String `tfsdk:"value" json:"value,computed"`
}

type UserAgentBlockingRuleFindOneByDataSourceModel struct {
	ZoneIdentifier    types.String `tfsdk:"zone_identifier" path:"zone_identifier,required"`
	Description       types.String `tfsdk:"description" query:"description,optional"`
	DescriptionSearch types.String `tfsdk:"description_search" query:"description_search,optional"`
	UASearch          types.String `tfsdk:"ua_search" query:"ua_search,optional"`
}
