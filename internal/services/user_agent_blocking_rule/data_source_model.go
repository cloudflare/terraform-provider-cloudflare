// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_agent_blocking_rule

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/firewall"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserAgentBlockingRuleResultDataSourceEnvelope struct {
	Result UserAgentBlockingRuleDataSourceModel `json:"result,computed"`
}

type UserAgentBlockingRuleDataSourceModel struct {
	ID            types.String                                                                `tfsdk:"id" path:"ua_rule_id,computed"`
	UARuleID      types.String                                                                `tfsdk:"ua_rule_id" path:"ua_rule_id,optional"`
	ZoneID        types.String                                                                `tfsdk:"zone_id" path:"zone_id,required"`
	Description   types.String                                                                `tfsdk:"description" json:"description,computed"`
	Mode          types.String                                                                `tfsdk:"mode" json:"mode,computed"`
	Paused        types.Bool                                                                  `tfsdk:"paused" json:"paused,computed"`
	Configuration customfield.NestedObject[UserAgentBlockingRuleConfigurationDataSourceModel] `tfsdk:"configuration" json:"configuration,computed"`
	Filter        *UserAgentBlockingRuleFindOneByDataSourceModel                              `tfsdk:"filter"`
}

func (m *UserAgentBlockingRuleDataSourceModel) toReadParams(_ context.Context) (params firewall.UARuleGetParams, diags diag.Diagnostics) {
	params = firewall.UARuleGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *UserAgentBlockingRuleDataSourceModel) toListParams(_ context.Context) (params firewall.UARuleListParams, diags diag.Diagnostics) {
	params = firewall.UARuleListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if !m.Filter.Description.IsNull() {
		params.Description = cloudflare.F(m.Filter.Description.ValueString())
	}
	if !m.Filter.Paused.IsNull() {
		params.Paused = cloudflare.F(m.Filter.Paused.ValueBool())
	}
	if !m.Filter.UserAgent.IsNull() {
		params.UserAgent = cloudflare.F(m.Filter.UserAgent.ValueString())
	}

	return
}

type UserAgentBlockingRuleConfigurationDataSourceModel struct {
	Target types.String `tfsdk:"target" json:"target,computed"`
	Value  types.String `tfsdk:"value" json:"value,computed"`
}

type UserAgentBlockingRuleFindOneByDataSourceModel struct {
	Description types.String `tfsdk:"description" query:"description,optional"`
	Paused      types.Bool   `tfsdk:"paused" query:"paused,optional"`
	UserAgent   types.String `tfsdk:"user_agent" query:"user_agent,optional"`
}
