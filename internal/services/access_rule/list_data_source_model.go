// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_rule

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/firewall"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessRulesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AccessRulesResultDataSourceModel] `json:"result,computed"`
}

type AccessRulesDataSourceModel struct {
	AccountID     types.String                                                   `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID        types.String                                                   `tfsdk:"zone_id" path:"zone_id,optional"`
	Direction     types.String                                                   `tfsdk:"direction" query:"direction,optional"`
	Mode          types.String                                                   `tfsdk:"mode" query:"mode,optional"`
	Notes         types.String                                                   `tfsdk:"notes" query:"notes,optional"`
	Order         types.String                                                   `tfsdk:"order" query:"order,optional"`
	Configuration *AccessRulesConfigurationDataSourceModel                       `tfsdk:"configuration" query:"configuration,optional"`
	Match         types.String                                                   `tfsdk:"match" query:"match,computed_optional"`
	MaxItems      types.Int64                                                    `tfsdk:"max_items"`
	Result        customfield.NestedObjectList[AccessRulesResultDataSourceModel] `tfsdk:"result"`
}

func (m *AccessRulesDataSourceModel) toListParams(_ context.Context) (params firewall.AccessRuleListParams, diags diag.Diagnostics) {
	params = firewall.AccessRuleListParams{}

	if m.Configuration != nil {
		paramsConfiguration := firewall.AccessRuleListParamsConfiguration{}
		if !m.Configuration.Target.IsNull() {
			paramsConfiguration.Target = cloudflare.F(firewall.AccessRuleListParamsConfigurationTarget(m.Configuration.Target.ValueString()))
		}
		if !m.Configuration.Value.IsNull() {
			paramsConfiguration.Value = cloudflare.F(m.Configuration.Value.ValueString())
		}
		params.Configuration = cloudflare.F(paramsConfiguration)
	}
	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(firewall.AccessRuleListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Match.IsNull() {
		params.Match = cloudflare.F(firewall.AccessRuleListParamsMatch(m.Match.ValueString()))
	}
	if !m.Mode.IsNull() {
		params.Mode = cloudflare.F(firewall.AccessRuleListParamsMode(m.Mode.ValueString()))
	}
	if !m.Notes.IsNull() {
		params.Notes = cloudflare.F(m.Notes.ValueString())
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(firewall.AccessRuleListParamsOrder(m.Order.ValueString()))
	}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type AccessRulesConfigurationDataSourceModel struct {
	Target types.String `tfsdk:"target" json:"target,computed_optional"`
	Value  types.String `tfsdk:"value" json:"value,computed_optional"`
}

type AccessRulesResultDataSourceModel struct {
}
